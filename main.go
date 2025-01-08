package main

import (
	"flag"
	"fmt"

	"githib.com/dharmendrashaw/csi-driver/pkg/driver"
)

func main() {
	var (
		endpoint = flag.String("endpoint", "defaultValue", "Endpoint of gRPC server will run at")
		name     = flag.String("name", "defaultValue", "Name of the csi driver")
		token    = flag.String("token", "defaultValue", "Token for authenticating storage")
		region   = flag.String("region", "ams3", "region of volume going to be provisioned")
	)

	flag.Parse()

	fmt.Println(*endpoint, *name, *token, *region)

	//create driver
	driver := driver.NewDriver(driver.InputParams{
		Name:     *name,
		Endpoint: *endpoint,
		Region:   *region,
		Token:    *token,
	})

	//run driver

	if err := driver.Run(); err != nil {
		fmt.Printf("Error %s, running driver\n", err.Error())
	}

}
