package archive

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"stackoverflow-docker/tools"
)

type Compressed struct {
	Digest  tools.Hash
	DigestC tools.Hash
}

func FromFolder(path, archiveName string) (*Compressed, error) {
	tarBuffer := bytes.Buffer{}
	gzipBuffer := bytes.Buffer{}

	if err := buildTar(path, &tarBuffer); err != nil {
		return nil, err
	}

	content := tarBuffer.Bytes()
	if err := compressToGzip(content, &gzipBuffer); err != nil {
		return nil, err
	}

	file, err := os.Create(archiveName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	contentGzip := gzipBuffer.Bytes()
	if _, err := file.Write(contentGzip); err != nil {
		return nil, err
	}

	return &Compressed{tools.Digest(content), tools.Digest(contentGzip)}, nil
}

func buildTar(path string, w io.Writer) error {
	tw := tar.NewWriter(w)
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
		return err
	}

	return nil
}

func compressToGzip(content []byte, w io.Writer) error {
	gw := gzip.NewWriter(w)
	if _, err := gw.Write(content); err != nil {
		return fmt.Errorf("can't compress content: %w", err)
	}
	if err := gw.Close(); err != nil {
		return fmt.Errorf("can't compress content: %w", err)
	}

	return nil
}
