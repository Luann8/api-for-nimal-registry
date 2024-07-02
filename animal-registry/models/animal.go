package models

type Animal struct {
	ID           string  `json:"id"`
	BrincoID     string  `json:"brinco_id"`
	Lote         string  `json:"lote"`
	Manejo       string  `json:"manejo"`
	Nascimento   string  `json:"nascimento"`
	Idade        int     `json:"idade"`
	Raca         string  `json:"raca"`
	Sexo         string  `json:"sexo"`
	Categoria    string  `json:"categoria"`
	Cria         string  `json:"cria"`
	Mae          string  `json:"mae"`
	Pai          string  `json:"pai"`
	PesoAnterior float64 `json:"peso_anterior"`
	PesoAtual    float64 `json:"peso_atual"`
	Medicamentos string  `json:"medicamentos"`
	Observacoes  string  `json:"observacoes"`
}
