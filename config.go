package oneGroup

type T struct {
	Mysql struct {
		User     string `json:"user"`
		Password string `json:"password"`
		Host     string `json:"host"`
		Port     int    `json:"port"`
		Database string `json:"database"`
	} `json:"mysql"`
	Redis struct {
		Addr     string `json:"addr"`
		Password string `json:"password"`
		Db       int    `json:"db"`
	} `json:"redis"`
	Consul struct {
		Hort string `json:"Hort"`
		Port int    `json:"port"`
	} `json:"consul"`
	Group struct {
		Host string `json:"host"`
		Port int    `json:"port"`
	} `json:"group"`
	Aliyun struct {
		AccessKeyID     string `json:"AccessKeyID"`
		AccessKeySecret string `json:"AccessKeySecret"`
		Bucket          string `json:"bucket"`
	} `json:"aliyun"`
	Alipay struct {
		PrivateKey string `json:"privateKey"`
		AppId      string `json:"appId"`
		NotifyURL  string `json:"NotifyURL"`
		ReturnURL  string `json:"ReturnURL"`
	} `json:"alipay"`
	Es struct {
		Addr string `json:"addr"`
	} `json:"Es"`
	Mongodb struct {
		User       string `json:"user"`
		Passwd     string `json:"passwd"`
		Hort       string `json:"Hort"`
		Port       int    `json:"port"`
		Database   string `json:"Database"`
		Collection string `json:"Collection"`
	} `json:"mongodb"`
	Txy struct {
		AccessKeyID     string `json:"AccessKeyID"`
		AccessKeySecret string `json:"AccessKeySecret"`
	} `json:"txy"`
	GaoDe struct {
		Key string `json:"key"`
	} `json:"gaoDe"`
}

var NaCos T
