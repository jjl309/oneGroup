package oneGroup

import (
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
)

func InitEs() {
	var err error
	cfg := elasticsearch.Config{
		Addresses: []string{
			NaCos.Es.Addr,
		},
		// ...
	}
	Es, err = elasticsearch.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	fmt.Println("elasticsearch connect success")
}
