package configuration

type File interface {
	ToString() (string, error)
	ToBytes() ([]byte, error)
}
