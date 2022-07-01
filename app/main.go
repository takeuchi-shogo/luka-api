package main

import "github.com/takeuchi-shogo/luka-api/src/infrastructure"

func main() {
	config := infrastructure.NewConfig()
	// db := infrastructure.NewDB(config)

	r := infrastructure.NewRouting(config)

	r.Run(r.Port)
}
