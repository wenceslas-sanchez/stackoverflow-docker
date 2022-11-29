package archive

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
)

func FromFolder(path, archiveName string) (err error) {
	buffer, err := createFromFolder(path)
	if err != nil {
		return err
	}
	file, err := os.Create(archiveName)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.Write(buffer.Bytes()); err != nil {
		return err
	}
	return nil
}

func createFromFolder(path string) (*bytes.Buffer, error) {
	var buffer bytes.Buffer

	gw := gzip.NewWriter(&buffer)
	defer gw.Close()
	tw := tar.NewWriter(gw)
	defer tw.Close()

	walker := func(file string, fi os.FileInfo, err error) error {
		header, err := tar.FileInfoHeader(fi, file)
		if err != nil {
			return nil
		}
		if err := tw.WriteHeader(header); err != nil {
			return err
		}
		header.Name = filepath.ToSlash(file)

		if !fi.IsDir() {
			content, err := os.Open(file)
			if err != nil {
				return err
			}

			if _, err := io.Copy(tw, content); err != nil {
				return err
			}
		}

		return nil
	}
	if err := filepath.Walk(path, walker); err != nil {
		return nil, err
	}

	return &buffer, nil
}
