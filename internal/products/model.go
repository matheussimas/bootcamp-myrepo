package products

// o meu repository vai fazer as solicitações para o BD

// começo declarando o modelo do dominio que será trabalhado
type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Type  string  `json:"type"`
	Count int     `json:"count"`
	Price float64 `json:"price"`
}
