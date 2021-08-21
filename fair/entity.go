package fair

type Fair struct {
	ID         int    `json:"id,omitempty"`
	Long       int    `json:"long,omitempty"`
	Lat        int    `json:"lat,omitempty"`
	Setcens    int    `json:"setcens,omitempty"`
	Areap      int    `json:"areap,omitempty"`
	Coddist    int    `json:"coddist,omitempty"`
	Distrito   string `json:"distrito,omitempty"`
	Codsubpref int    `json:"codsubpref,omitempty"`
	Subprefe   string `json:"subprefe,omitempty"`
	Regiao5    string `json:"regiao_5,omitempty"`
	Regiao8    string `json:"regiao_8,omitempty"`
	NomeFeira  string `json:"nome_feira,omitempty"`
	Registro   string `json:"registro,omitempty"`
	Logradouro string `json:"logradouro,omitempty"`
	Numero     string `json:"numero,omitempty"`
	Bairro     string `json:"bairro,omitempty"`
	Referencia string `json:"referencia,omitempty"`
}
