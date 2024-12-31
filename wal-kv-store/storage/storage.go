package storage

type Storage interface {
	Load() ([]byte, error)
	Save(data []byte) error
}
