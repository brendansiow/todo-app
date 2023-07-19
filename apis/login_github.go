package apis

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/brendansiow/todo-app/core"
	"github.com/brendansiow/todo-app/helper"
	"github.com/brendansiow/todo-app/models"
	"github.com/gin-gonic/gin"
)

const loginType = "github"

func BindLoginApi(router *gin.RouterGroup) {
	router.GET("/login/github", loginGithub)
	router.GET("/login/github/callback", loginGithubCallback)
	router.GET("/login/github/success", loginGithubSuccess)
	router.GET("/login/github/failed", loginGithubFailed)
}

func loginGithub(c *gin.Context) {
	// Get the environment variable
	githubClientID := helper.GetGithubClientID()
	githubOAuthUrl := helper.GetGithubOAuthUrl()

	// Create the dynamic redirect URL for login
	callbackUrl := url.URL{Scheme: "http", Host: c.Request.Host, Path: "/login/github/callback"}

	redirectURL := fmt.Sprintf(
		githubOAuthUrl,
		githubClientID,
		callbackUrl.String(),
	)

	c.Redirect(http.StatusFound, redirectURL)
}

func loginGithubCallback(c *gin.Context) {
	code := c.Query("code")

	githubAccessToken := getGithubAccessToken(code)
	githubData := getGithubData(githubAccessToken)

	loginGithubCallbackHandler(c, githubData)
}

func loginGithubCallbackHandler(c *gin.Context, githubData string) {
	var prettyJSON bytes.Buffer
	parserr := json.Indent(&prettyJSON, []byte(githubData), "", "\t")
	if parserr != nil {
		log.Panic("JSON parse error")
	}

	githubUser := models.GithubUser{}
	json.Unmarshal(prettyJSON.Bytes(), &githubUser)

	if githubUser.Email == "" {
		// Unauthorized users get an unauthorized message
		redirectURL := url.URL{Scheme: "http", Host: c.Request.Host, Path: "/login/github/failed"}
		c.Redirect(http.StatusFound, redirectURL.String())
		return
	}
	token := loginOrRegisterGithubUser(githubUser)
	redirectURL := url.URL{Scheme: "http", Host: c.Request.Host, Path: "/login/github/success", RawQuery: "token=" + token}
	c.Redirect(http.StatusFound, redirectURL.String())
}

func loginGithubSuccess(c *gin.Context) {
	token := c.Query("token")
	c.JSON(http.StatusOK, gin.H{"data": token})
}

func loginGithubFailed(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{"message": "Github login failed"})
}

func getGithubAccessToken(code string) string {
	clientID := helper.GetGithubClientID()
	clientSecret := helper.GetGithubClientSecret()

	// Set us the request body as JSON
	requestBodyMap := map[string]string{
		"client_id":     clientID,
		"client_secret": clientSecret,
		"code":          code,
	}
	requestJSON, _ := json.Marshal(requestBodyMap)

	// POST request to set URL
	req, reqErr := http.NewRequest(
		"POST",
		"https://github.com/login/oauth/access_token",
		bytes.NewBuffer(requestJSON),
	)
	if reqErr != nil {
		log.Panic("Request creation failed")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	// Get the response
	resp, respErr := http.DefaultClient.Do(req)
	if respErr != nil {
		log.Panic("Request failed")
	}
	// Response body converted to stringified JSON
	responseBody, _ := io.ReadAll(resp.Body)

	// Convert stringified JSON to a struct object of type githubAccessTokenResponse
	var githubResponse models.GithubAccessTokenResponse
	json.Unmarshal(responseBody, &githubResponse)

	// Return the access token (as the rest of the
	// details are relatively unnecessary for us)
	return githubResponse.AccessToken
}

func getGithubData(accessToken string) string {
	// Get request to a set URL
	req, reqerr := http.NewRequest(
		"GET",
		"https://api.github.com/user",
		nil,
	)
	if reqerr != nil {
		log.Panic("API Request creation failed")
	}

	authorizationHeaderValue := fmt.Sprintf("token %s", accessToken)
	req.Header.Set("Authorization", authorizationHeaderValue)

	// Make the request
	resp, resperr := http.DefaultClient.Do(req)
	if resperr != nil {
		log.Panic("Request failed")
	}

	// Read the response as a byte slice
	respbody, _ := io.ReadAll(resp.Body)

	// Convert byte slice to string and return
	return string(respbody)

}

func loginOrRegisterGithubUser(githubUser models.GithubUser) string {
	user := models.User{
		LoginType:   loginType,
		LoginTypeId: githubUser.ID,
	}
	userResult := core.DB.Find(&user)
	if userResult.Error != nil {
		log.Fatal(userResult.Error.Error())
	}
	if userResult.RowsAffected == 0 {
		githubUserResult := core.DB.Create(&githubUser)
		if githubUserResult.Error != nil {
			return ""
		}
		user.LoginType = loginType
		userResult := core.DB.Create(&user)
		if userResult.Error != nil {
			return ""
		}
		return helper.GenerateJwtToken(user.ID)
	} else {
		if user.LoginType == loginType {
			//Update  user
			user.Email = githubUser.Email
			core.DB.Save(&githubUser)
			core.DB.Save(&user)
			return helper.GenerateJwtToken(user.ID)
		} else {
			//TODO: handle for other login
			return ""
		}
	}
}
