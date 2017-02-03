package main

import (
	"fmt"
	"net/http"

	"github.com/spf13/viper"
)

type rewriter interface {
	rewrite(string) *string
	String() string
}

func main() {
	loadConfiguration()
	rewriteServer := rewriteServer{}
	loadStaticRulesFromConfig(&rewriteServer)

	rewriteServer.listRewriters()

	http.HandleFunc("/", rewriteServer.ServeHTTP)
	portNumber := viper.GetInt("app.port")
	serverURI := fmt.Sprintf(":%v", portNumber)
	fmt.Printf("Listening on %v", serverURI)
	http.ListenAndServe(serverURI, nil)
}

func loadConfiguration() {
	viper.SetDefault("app.debug", false)
	viper.SetDefault("app.port", "8080")
	viper.SetDefault("static-rules", []staticRule{})

	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/go-there/")
	viper.AddConfigPath("$HOME/.go-there")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Could not find a config file:")
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Using configuration file: %s\n", viper.ConfigFileUsed())
	}
	viper.AutomaticEnv()
}

func debug(msg string, vars ...interface{}) {
	if viper.GetBool("app.debug") {
		fmt.Printf(msg, vars...)
	}
}
