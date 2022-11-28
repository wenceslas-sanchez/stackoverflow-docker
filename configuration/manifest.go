package configuration

import (
	"encoding/json"
	"fmt"
)

// Based on https://github.com/philpearl/scratchbuild/blob/master/types.go

// Descriptor contains a reference to a blob
type Descriptor struct {
	// MediaType describe the type of the content. All text based formats are
	// encoded as utf-8.
	MediaType string `json:"mediaType,omitempty"`

	// Size in bytes of content.
	Size int64 `json:"size,omitempty"`

	// Digest uniquely identifies the content. A byte stream can be verified
	// against this digest.
	Digest string `json:"digest,omitempty"`

	// URLs contains the source URLs of this content.
	URLs []string `json:"urls,omitempty"`
}

// Manifest describes a container image
type Manifest struct {
	Versioned

	// Config references the image configuration as a blob.
	Config Descriptor `json:"config"`

	// Layers lists descriptors for the layers referenced by the
	// configuration.
	Layers []Descriptor `json:"layers"`
}

func (m *Manifest) ToString() (string, error) {
	out, err := json.Marshal(&m)
	if err != nil {
		return "", fmt.Errorf("cannot marshalize the manifest: %q", err)
	}
	return string(out), nil
}
