package metadata

import "stackoverflow-docker/tools"

type MetaData struct {
	Size   int64
	Digest tools.Hash
}

type Pile struct {
	Content           MetaData
	CompressedContent MetaData
}

func NewMetadata(content []byte) MetaData {
	return MetaData{
		int64(len(content)),
		tools.Digest(content),
	}
}

func New(content, contentGzip []byte) Pile {
	return Pile{NewMetadata(content), NewMetadata(contentGzip)}
}
