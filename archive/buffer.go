package archive

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"os"
)

// FromBuffer create an archive with the name archive_name which
// contains the file `filename` which contains `content`
func FromBuffer(content []byte, filename, archiveName string) error {
	buffer, err := createFromBuffer(content, filename, 7777)
	if err != nil {
		return err
	}
	file, err := os.Create(archiveName)
	if err != nil {
		return err
	}
	defer func() {
		err = file.Close()
	}()

	if _, err := file.Write(buffer.Bytes()); err != nil {
		return err
	}
	return nil
}

// createFromBuffer a archive file within a buffer. The buffer is surrounded
// by a GZIP writer and a TAR writer
func createFromBuffer(content []byte, filename string, mode int64) (buffer *bytes.Buffer, err error) {
	gw := gzip.NewWriter(buffer)
	defer func() {
		err = gw.Close()
	}()
	tw := tar.NewWriter(gw)
	defer func() {
		err = tw.Close()
	}()

	err = addToArchive(tw, content, filename, mode)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// addToArchive add the content to the archive buffer.
func addToArchive(tw *tar.Writer, content []byte, filename string, mode int64) error {
	header := &tar.Header{
		Name: filename, Mode: mode, Size: int64(len(content)),
	}

	if err := tw.WriteHeader(header); err != nil {
		return err
	}

	if _, err := tw.Write(content); err != nil {
		return err
	}

	return nil
}
