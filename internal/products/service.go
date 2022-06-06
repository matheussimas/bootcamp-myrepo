package products

type Service interface {
	GetAll() ([]Product, error)
	Store(name, typee string, count int, price float64) (Product, error)
}

type service struct {
	repository Repository
}

func (s service) GetAll() ([]Product, error) {
	ps, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return ps, nil
}

func (s service) Store(name, typee string, count int, price float64) (Product, error) {

	lastID, err := s.repository.LastID()

	if err != nil {
		return Product{}, err
	}

	lastID++

	product, err := s.repository.Store(lastID, name, typee, count, price)

	if err != nil {
		return Product{}, err
	}

	return product, nil

}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}
