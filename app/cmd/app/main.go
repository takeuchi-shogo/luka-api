package main

import "github.com/takeuchi-shogo/luka-api/src/infrastructure"

func main() {
	config := infrastructure.NewConfig()
	db := infrastructure.NewDB(config)
	// _ = infrastructure.NewRedis()
	infrastructure.NewElasticSearch()

	r := infrastructure.NewRouting(config, db)

	r.Run(r.Port)
}
