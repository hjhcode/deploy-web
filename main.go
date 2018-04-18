package main

import (
	"flag"
	"github.com/hjhcode/deploy-web/common"
	"github.com/hjhcode/deploy-web/router"
	"github.com/TV4/graceful"
	"net/http"
	"github.com/hjhcode/deploy-web/common/g"
	"fmt"
)

func main() {
	cfgFile := flag.String("c", "cfg/cfg.toml.debug", "set config file")
	flag.Parse()

	common.Init(*cfgFile)
	defer common.Close()

	router := router.GetRouter()

	graceful.LogListenAndServe(&http.Server{
		Addr:    fmt.Sprintf(":%d", g.Conf().Run.HTTPPort),
		Handler: router,
	})
}
