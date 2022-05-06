package model

type River struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type Alga struct {
	Name        string       `json:"name"`
	Src         string       `json:"src"`
	River       string       `json:"river"`
	Annotations []Annotation `json:"annotations"`
}

type Annotation struct {
	Description string `json:"description"`
	Format      string `json:"format"`
	Url         string `json:"url"`
}

type Operator struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Access   int    `json:"access" binding:"-"`
}
