package oneGroup

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
)

func InitConsul() {
	sul := NaCos.Consul
	grp := NaCos.Group
	client, err := api.NewClient(&api.Config{
		Address: fmt.Sprintf("%s:%d", sul.Hort, sul.Port),
	})
	if err != nil {
		return
	}
	id := uuid.New().String()
	err = client.Agent().ServiceRegister(&api.AgentServiceRegistration{
		ID:      id,
		Name:    "shop_server",
		Tags:    []string{"GRPC"},
		Port:    grp.Port,
		Address: grp.Host,
		Check: &api.AgentServiceCheck{
			Interval:                       "5s",
			Timeout:                        "5s",
			GRPC:                           fmt.Sprintf("%s:%d", grp.Host, grp.Port),
			DeregisterCriticalServiceAfter: "30s",
		},
	})
	if err != nil {
		return
	}
	fmt.Println("consul connect success")
}
