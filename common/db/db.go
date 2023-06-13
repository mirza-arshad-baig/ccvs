package db

type DB interface {
	Connect()
	Close() error
}
