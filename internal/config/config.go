package config

type Config struct {
	Database Database `json:"database"`
}

type Database struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Name     string `json:"name"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
}
