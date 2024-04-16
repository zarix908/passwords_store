package storage

type Storage interface {
	AddKey(data []byte) error
	GetKey() ([]byte, error)
}
