package kubernetes

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func doTag(imageName string, version string, environment string, appName string) error {
	return reTagImageOnDockerHub(imageName, version, environment)
}

type auth struct {
	Token string `json:"token"`
}

func reTagImageOnDockerHub(repo string, version string, environment string) error {
	client := http.DefaultClient
	authReq, err := http.NewRequest("GET", "https://auth.docker.io/token", nil)
	if err != nil {
		return err
	}

	authReq.SetBasicAuth("sanalmarket", "Migros2017")
	query := authReq.URL.Query()
	query.Add("service", "registry.docker.io")
	query.Add("scope", fmt.Sprintf("repository:%s:pull,push", repo))
	authReq.URL.RawQuery = query.Encode()
	do, err := client.Do(authReq)
	if err != nil {
		return err
	}

	token := auth{}
	json.NewDecoder(do.Body).Decode(&token)

	manifestPath := "https://registry-1.docker.io/v2/%s/manifests/%s"
	manifestReq, err := http.NewRequest("GET", fmt.Sprintf(manifestPath, repo, version), nil)
	if err != nil {
		return err
	}

	manifestReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token.Token))
	manifestReq.Header.Set("Accept", "application/vnd.docker.distribution.manifest.v2+json")
	manifestResp, err := client.Do(manifestReq)
	if err != nil {
		return err
	}

	taggingReq, err := http.NewRequest("PUT", fmt.Sprintf(manifestPath, repo, "latest-in-"+environment), manifestResp.Body)
	if err != nil {
		return err
	}

	taggingReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token.Token))
	taggingReq.Header.Set("Content-type", "application/vnd.docker.distribution.manifest.v2+json")
	taggingResp, err := client.Do(taggingReq)
	if taggingResp.StatusCode != 201 {
		return err
	}
	return nil
}
