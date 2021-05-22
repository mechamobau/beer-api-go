package beer

type UseCase interface {
	GetAll() ([]*Beer, error)
	Get(ID int64) (*Beer, error)
	Store(b *Beer) error
	Update(b *Beer) error
	Remove(ID int64) error
}

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetAll() ([]*Beer, error) {
	return nil, nil
}

func (s *Service) Get(ID int64) (*Beer, error) {
	return nil, nil
}

func (s *Service) Store(b *Beer) error {
	return nil
}

func (s *Service) Update(b *Beer) error {
	return nil
}

func (s *Service) Remove(ID int64) error {
	return nil
}
