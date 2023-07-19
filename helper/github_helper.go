package helper

import (
	"log"
	"os"
)

func GetGithubOAuthUrl() string {

	githubOAuthUrl, exists := os.LookupEnv("GITHUB_OAUTH_URL")
	if !exists {
		log.Fatal("Github OAuth URL not found")
	}

	return githubOAuthUrl
}

func GetGithubClientID() string {

	githubClientID, exists := os.LookupEnv("GITHUB_CLIENT_ID")
	if !exists {
		log.Fatal("Github Client ID not found")
	}

	return githubClientID
}

func GetGithubClientSecret() string {

	githubClientSecret, exists := os.LookupEnv("GITHUB_CLIENT_SECRET")
	if !exists {
		log.Fatal("Github Client Secret found")
	}

	return githubClientSecret
}
