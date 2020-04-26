package config
type Config struct {
	DB *DBConfig
}
type DBConfig struct {
	Dialect  string
	Username string
	Password string
	Host string
	Name     string
	Charset  string
}
func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "mysql",
			Username: "root",
			Password: "root",
			Host:	"localhost",
			Name:     "mydb",
			Charset:  "utf8",
		},
	}
}
