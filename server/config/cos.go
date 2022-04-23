package config

type TencentCOS struct {
	Bucket     string `json:"bucket"`
	Region     string `json:"region"`
	SecretID   string `json:"secretID"`
	SecretKey  string `json:"secretKey"`
	BaseURL    string `json:"baseURL"`
	PathPrefix string `json:"pathPrefix"`
}
