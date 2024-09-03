package veiculo

type Veiculo struct {
	Marca       string  `json:"marca"`
	Modelo      string  `json:"modelo"`
	Ano         string  `json:"ano"`
	Cor         string  `json:"cor"`
	Preco       float32 `json:"preco"`
	Flagvendido string  `json:"flagvendido"`
	Id          int     `json:"id"`
}
