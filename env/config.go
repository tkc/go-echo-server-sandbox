package env

import (
	"fmt"
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

func getAppEnv() string {
	v := getViper()
	app := v.Get("app")
	return app.(string)
}

func IsProd() bool {
	if (getAppEnv() == "prod") {
		return true
	} else {
		return false
	}
}

func GetPort() string {
	v := getViper()
	return v.Get("port").(string)
}

func GetDataBaseAccess() string {
	v := getViper()
	connection := fmt.Sprintf("%s:%s@tcp([127.0.0.1]:3306)/%s?charset=utf8&parseTime=True&loc=",
		v.Get("database.user_name"),
		v.Get("database.password"),
		v.Get("database.name"),
	)
	connection += "Asia%2FTokyo"
	return connection
}

