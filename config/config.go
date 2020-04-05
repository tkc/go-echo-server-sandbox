package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func getViper() *viper.Viper {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.SetConfigType("yaml")
	vp.AddConfigPath(".")
	vp.AddConfigPath("../")
	vp.AddConfigPath("../../")
	vp.AddConfigPath("../../../")
	vp.AddConfigPath("../../../../")
	vp.AddConfigPath("../../../../../")
	err := vp.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("error: %s", err))
	}
	return vp
}

func IsProd() bool {
	return getViper().Get("app") == "prod"
}

func GetPort() string {
	return getViper().Get("port").(string)
}

func GetDataBaseAccess() string {
	v := getViper()
	connection := fmt.Sprintf("%s:%s@tcp([127.0.0.1]:3306)/%s?charset=utf8&parseTime=True&loc=",
		v.Get("database.user"),
		v.Get("database.password"),
		v.Get("database.name"),
	)

	log.Print(connection)

	connection += "Asia%2FTokyo"
	return connection
}
