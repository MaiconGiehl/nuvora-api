package entity

type BusRepositoryInterface interface {
	Save(bus *Bus) error
}
