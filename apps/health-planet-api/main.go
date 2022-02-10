package main

import (
	"flag"
	"health-planet-api/config"
	"health-planet-api/infrastracture/web"
)

var (
	path = flag.String("c", "config.localhost.yaml", "config file path")
)

func main() {
	flag.Parse()

	con, err := config.LoadFile(*path)
	if err != nil {
		panic(err)
	}

	con.LoadEnvPassword()
	web.LoadRouter(con).Start(":8001")
}
