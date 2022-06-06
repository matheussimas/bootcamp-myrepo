package products

// o meu repository vai fazer as solicitações para o BD

var ps []Product

var lastID int

// a interface vai definir os métodos que buscarão as informações no meu banco de dados
type Repository interface {
	GetAll() ([]Product, error)
	Store(id int, name string, typee string, count int, price float64) (Product, error)
	LastID() (int, error)
}

type repository struct{}

func (repository) GetAll() ([]Product, error) {
	return ps, nil
}

func (repository) LastID() (int, error) {
	return lastID, nil
}

func (repository) Store(id int, name string, typee string, count int, price float64) (Product, error) {
	p := Product{id, name, typee, count, price}
	ps = append(ps, p)
	lastID = p.ID
	return p, nil
}

func NewRepository() Repository {
	return &repository{}
}
