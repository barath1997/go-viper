package viper_setup

import (
	"fmt"

	"github.com/spf13/viper"
)

func SetupViper() error {
   viper.AddConfigPath("../go-viper/viper_setup/config")
   viper.SetConfigName("appconfig")
   viper.SetConfigType("yaml")

   if err := viper.ReadInConfig(); err != nil {
	  return err
   }
   fmt.Println("val : ",viper.Get("port"))
   return nil
}