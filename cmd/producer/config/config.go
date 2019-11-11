package config

import (
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/viper"
)

var (
	POSTGRESDRIVER = ""
	POSTGRESURL    = ""
	TIMEOUT        time.Duration
	PORT           = ""
)

func Load() {
	viper.SetConfigFile("./cmd/producer/config.json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		fmt.Println("Running in debug mode")
	}

	postgresDBName := viper.GetString(`postgres.db_name`)
	postgresUser := viper.GetString(`postgres.db_user`)
	postgresPassword := viper.GetString(`postgres.db_password`)
	postgresPort := viper.GetString(`postgres.port`)
	postgresHost := viper.GetString(`postgres.host`)
	POSTGRESDRIVER = viper.GetString(`postgres.driver`)
	POSTGRESURL = fmt.Sprintf("postgres://%s:%s@%s:%s/%s", postgresUser, postgresPassword, postgresHost, postgresPort, postgresDBName)

	if viper.GetString("environment") == "development" {
		POSTGRESURL = fmt.Sprintf("%s?sslmode=disable", POSTGRESURL)
	}

	TIMEOUT = time.Duration(viper.GetInt(`context.timeout`)) * time.Second
	PORT = strconv.Itoa(viper.GetInt(`server.address`))
}
