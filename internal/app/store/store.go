package store

// Инициализируем репозитории
//go:generate mockery --name=Store --structname=Store
type Store interface {
	File() FileRepository
	Region() RegionRepository
}
