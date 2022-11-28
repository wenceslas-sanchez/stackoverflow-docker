package archive

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"os"
)

// NewFile create an archive with the name archive_name which
// contains the file `filename` which contains `content`
func NewFile(content []byte, filename, archive_name string) error {
	buffer, err := Create(content, filename, 7777)
	if err != nil {
		return err
	}
	file, err := os.Create(archive_name)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.Write(buffer.Bytes()); err != nil {
		return err
	}
	return nil
}

// Create a archive file within a buffer. The buffer is surrounded
// by a GZIP writer and a TAR writer
func Create(content []byte, filename string, mode int64) (*bytes.Buffer, error) {
	var buffer bytes.Buffer

	gw := gzip.NewWriter(&buffer)
	defer gw.Close()
	tw := tar.NewWriter(gw)
	defer tw.Close()

	err := addToArchive(tw, content, filename, mode)
	if err != nil {
		return nil, err
	}

	return &buffer, nil
}

// addToArchive add the content to the archive buffer.
func addToArchive(tw *tar.Writer, content []byte, filename string, mode int64) error {
	header := &tar.Header{
		Name: filename, Mode: mode, Size: int64(len(content)),
	}

	err := tw.WriteHeader(header)
	if err != nil {
		return err
	}

	if _, err := tw.Write(content); err != nil {
		return err
	}

	return nil
}
