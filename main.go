package main

import ("log"
"github.com/barath1997/go-viper/viper_setup"
)

func main() {

	if err := viper_setup.SetupViper(); err != nil {
		log.Fatal(err)
	}

}