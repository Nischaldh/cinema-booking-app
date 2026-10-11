package booking

type Service struct{
	store BookingStore
}

func NewSerive(store BookingStore) *Service{
	return &Service{
		store: store,
	}
}