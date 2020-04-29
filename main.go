package main

import (

	//"github.com/gin-gonic/gin"
	"net/http"

	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/web"

	//	"github.com/micro/go-plugins/config/source/grpc"
	_ "test/local"

	"github.com/siddontang/go/log"
)

func main() {

	reg := etcd.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"http://127.0.0.1:2379",
		}
	})
	service := web.NewService(
		web.Name("webir:ad:http"),
		web.Version("0.0.1"),
		web.Registry(reg),
		web.Address(":9002"),
	)
	//client.InitApiService()
	//model.Init()
	//r := gin.Default()
	//r := routers.NewRouter()
	//service.Handle("/", r)
	service.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Hello!!!"))
	})

	err := service.Run()
	if err != nil {
		log.Info(err)
	}
}
