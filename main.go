package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/go-github/v38/github"
	"golang.org/x/oauth2"
)

func main() {
	kustomizePath := os.Getenv("INPUT_KUSTOMIZE_PATH")
	if kustomizePath == "" {
		log.Fatal("INPUT_KUSTOMIZE_PATH is not set")
	}
	token := os.Getenv("INPUT_GITHUB_TOKEN")
	if token == "" {
		log.Fatal("INPUT_GITHUB_TOKEN is not set")
	}
	targetRepoPath := os.Getenv("INPUT_TARGET_REPO")
	if targetRepoPath == "" {
		log.Fatal("INPUT_TARGET_REPO is not set")
	}
	filePath := os.Getenv("INPUT_FILE_PATH")
	if filePath == "" {
		log.Fatal("INPUT_FILE_PATH is not set")
	}
	imageBase := os.Getenv("INPUT_IMAGE_BASE")
	if imageBase == "" {
		log.Fatal("INPUT_IMAGE_BASE is not set")
	}
	namespace := os.Getenv("INPUT_NAMESPACE")
	if namespace == "" {
		namespace = strings.Split(os.Getenv("GITHUB_REPOSITORY"), "/")[1]
	}
	
	imageTag := os.Getenv("INPUT_IMAGE_TAG")
	if imageTag == "" {
		fullSHA := os.Getenv("GITHUB_SHA")
		if len(fullSHA) >= 7 {
			imageTag = fullSHA[:7]  // default to the first 7 characters of SHA if not provided
		} else {
			log.Fatal("GITHUB_SHA is unexpectedly short")
		}
	}
	
	splits := strings.Split(targetRepoPath, "/")
	owner := splits[0]
	repo := splits[1]

	if namespace == "" {
		namespace = strings.Split(os.Getenv("GITHUB_REPOSITORY"), "/")[1]
	}

	client := createClient(token)

	// Get default branch for the repository
	repoInfo, _, err := client.Repositories.Get(context.Background(), owner, repo)
	if err != nil {
		panic(err)
	}
	defaultBranch := repoInfo.GetDefaultBranch()

	// Check if the file already exists
	_, _, _, err = client.Repositories.GetContents(context.Background(), owner, repo, filePath, nil)
	if err == nil {
		fmt.Println("File already exists, skipping creation.")
		return
	}

	// Create the YAML content
	content := fmt.Sprintf(`
	kustomizePath: %s
	image: %s%s
	imageTag: '%s'
	namespace: %s`, kustomizePath, imageBase, namespace, imageTag, namespace)

	opts := &github.RepositoryContentFileOptions{
		Message:   github.String("Generated k8s YAML file by GH Action"),
		Content:   []byte(content),
		Branch:    github.String(defaultBranch),
		Committer: &github.CommitAuthor{Name: github.String("GitHub Action"), Email: github.String("action@github.com")},
	}
	_, _, err = client.Repositories.CreateFile(context.Background(), owner, repo, filePath, opts)
	if err != nil {
		panic(err)
	}

	fmt.Println("YAML file created successfully!")
}

func createClient(token string) *github.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(context.Background(), ts)

	return github.NewClient(tc)
}
