package infrastructure

type Config struct {
	CORS struct {
		AllowOrigins []string
	}
	DB struct {
		Production struct {
			Host     string
			UserName string
			Password string
			DBName   string
		}
	}
	Routing struct {
		Port string
	}
}

func NewConfig() *Config {

	c := new(Config)

	c.CORS.AllowOrigins = []string{""}

	c.DB.Production.Host = "localhost"
	c.DB.Production.UserName = "sgtkuc"
	c.DB.Production.Password = "smthr123"
	c.DB.Production.DBName = "luka"

	c.Routing.Port = ":8080"

	return c
}
