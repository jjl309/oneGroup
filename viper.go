package oneGroup

import (
	"fmt"
	"github.com/spf13/viper"
)

type AppConfig struct {
	NaCosAppFig
}
type NaCosAppFig struct {
	NamespaceId string
	IpAddr      string
	Port        uint64
	DataId      string
	Group       string
}

var Config AppConfig

func InitViper(devFilePath string) {
	viper.SetConfigFile(devFilePath)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&Config)
	if err != nil {
		panic(err)
	}
	fmt.Println("viper连接成功")
}
