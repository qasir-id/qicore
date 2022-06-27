package main

import (
	"fmt"
	"os"

	"gitlab.com/qasir/gateway-fajri/route"
	"gitlab.com/qasir/gateway-fajri/util"
	"gitlab.com/qasir/web/project/qasircore.git"
)

func init() {
	qasircore.Env("./")
}

func main() {
	e := route.Init()
	data, err := util.Json.MarshalIndent(e.Routes(), "", "  ")
	if err != nil {
		panic(fmt.Sprint(err))
	}

	fmt.Println(string(data))
	e.Logger.Fatal(e.Start(":" + os.Getenv("APP_ENDPOINT")))
}
