package main

import (
	"answers_app/routes"
)

func main() {

	r := routes.NewRouter()
	r.Run(":9090")
}
