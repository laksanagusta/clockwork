package config

type Config struct {
	Server   Server
	Redis    Redis
	MySql    MySql
	Midtrans Midtrans
}

type Server struct {
	Host string
	Port string
}

type Redis struct {
	Password string
}

type MySql struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type Midtrans struct {
	Key          string
	IsProduction bool
}
