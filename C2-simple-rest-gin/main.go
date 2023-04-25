package main

import "C2-simple-rest-gin/router"

func main() {
	var PORT = ":8080"

	router.StartServer().Run(PORT)

}
