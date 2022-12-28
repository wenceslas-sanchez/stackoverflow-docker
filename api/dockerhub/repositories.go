package dockerhub

import (
	"encoding/json"
	"fmt"
	"net/http"
	"stackoverflow-docker/tools"
)

const (
	BaseURL = "https://hub.docker.com"
	Version = "/v2"
	// NamespaceURL use as namespace `library` because we are looking for official images
	NamespaceURL    = BaseURL + Version + "/namespaces/library"
	RepositoriesURL = NamespaceURL + "/repositories"
)

// RepositoryTag stores little information about the image.
type RepositoryTag struct {
	FullSize int64 `json:"full_size"`
	Digest   tools.Hash
}

// ReadRepositoryTag request the DockerHub API to fetch the size and the digest of an image.
// It takes a repository and a tag to retrieve from `library` namespaces images metadata.
// Works only for official images.
func ReadRepositoryTag(repository, tag string) (*RepositoryTag, error) {
	client := http.Client{}
	response, err := client.Get(RepositoriesURL + fmt.Sprintf("/%s/tags/%s", repository, tag))
	if err != nil {
		return nil, fmt.Errorf("can't get tag information: %q", err)
	}
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("error while request dockerhub: %q", response.Status)
	}
	defer response.Body.Close()

	resultTag := &RepositoryTag{}
	if err := json.NewDecoder(response.Body).Decode(resultTag); err != nil {
		return resultTag, err
	}

	return resultTag, nil
}
