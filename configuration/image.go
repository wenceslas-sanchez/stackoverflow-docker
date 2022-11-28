package configuration

// Based on https://github.com/philpearl/scratchbuild/blob/master/types.go

import (
	"time"
)

// ImageConfig defines the execution parameters which should be used as a base when running a container using an image.
type ImageConfig struct {
	// User defines the username or UID which the process in the container should run as.
	User string `json:"User,omitempty"`

	// ExposedPorts a set of ports to expose from a container running this image.
	ExposedPorts map[string]struct{} `json:"ExposedPorts,omitempty"`

	// Env is a list of environment variables to be used in a container.
	Env []string `json:"Env,omitempty"`

	// Entrypoint defines a list of arguments to use as the command to execute when the container starts.
	Entrypoint []string `json:"Entrypoint,omitempty"`

	// Cmd defines the default arguments to the entrypoint of the container.
	Cmd []string `json:"Cmd,omitempty"`

	// Volumes is a set of directories describing where the process is likely write data specific to a container instance.
	Volumes map[string]struct{} `json:"Volumes,omitempty"`

	// WorkingDir sets the current working directory of the entrypoint process in the container.
	WorkingDir string `json:"WorkingDir,omitempty"`

	// Labels contains arbitrary metadata for the container.
	Labels map[string]string `json:"Labels,omitempty"`

	// StopSignal contains the system call signal that will be sent to the container to exit.
	StopSignal string `json:"StopSignal,omitempty"`
}

type RootFS struct {
	// Type is the type of the rootfs.
	Type string `json:"type"`
	// DiffIDS is a layer conten hashes # TODO make it an array
	DiffIDS string `json:"diff_ids"`
}

type History struct {
	// Created is the combined date and time at which the layer was created, formatted as defined by RFC 3339, section 5.6.
	Created *time.Time `json:"created,omitempty"`

	// CreatedBy is the command which created the layer.
	CreatedBy string `json:"created_by,omitempty"`

	// Author is the author of the build point.
	Author string `json:"author,omitempty"`

	// Comment is a custom message set when creating the layer.
	Comment string `json:"comment,omitempty"`

	// EmptyLayer is used to mark if the history item created a filesystem diff.
	EmptyLayer bool `json:"empty_layer,omitempty"`
}

type Image struct {
	Create       *time.Time
	Author       string      `json:"author,omitempty"`
	Architecture string      `json:"architecture"`
	OS           string      `json:"os"`
	Config       ImageConfig `json:"config,omitempty"`
	RootFS       RootFS      `json:"rootfs"`
	History      []History   `json:"history,omitempty"`
}
