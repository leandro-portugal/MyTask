package settings

import (
	"github.com/spf13/viper"
)

var setup *settings

type settings struct {
	API APISettings
	DB  DBSettings
}

type APISettings struct {
	Port string
}

type DBSettings struct {
	Hostname string
	Port     string
	Username string
	Password string
	Database string
}

func init() {

	viper.SetDefault("api.port", "9000")
	viper.SetDefault("database.hostname", "localhost")
	viper.SetDefault("database.port", "5432")

}

func load() error {

	viper.SetConfigName("settings")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()

	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	setup = new(settings)

	setup.API = APISettings{
		Port: viper.GetString("api.port"),
	}

	setup.DB = DBSettings{
		Hostname: viper.GetString("db.hostname"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		Database: viper.GetString("db.database"),
	}

	return nil
}

func GetDataBase() DBSettings {

	return setup.DB
}

func GetServerPort() string {

	return setup.API.Port
}
