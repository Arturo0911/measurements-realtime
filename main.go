package main

import (
	. "github.com/Arturo0911/measurements-realtime/connection"
	. "github.com/Arturo0911/measurements-realtime/server"
)

func main() {
	LoadEnv()
	Serve()
	//fmt.Println("")
	//DataBaseConnection()
	//Connection()
}
