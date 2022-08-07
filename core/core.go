package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"

	"leexsh/file_sys/core/internal/config"
	"leexsh/file_sys/core/internal/handler"
	"leexsh/file_sys/core/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/core-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
