package config

type Config struct {
	Mysql struct {
		User     string
		Password string
		Host     string
		Port     int
		Database string
	} `json:"mysql"`
	Redis struct {
		Host     string
		Password string
	} `json:"redis"`
	Es struct {
		Address string
	} `json:"es"`
}
