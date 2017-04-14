package main

import (
	"fmt"

	"github.com/spf13/viper"
	"gopkg.in/gin-gonic/gin.v1"
)

type rewriteResult struct {
	url             string
	showLandingPage bool
}

type rewriter interface {
	rewrite(string) *rewriteResult
	String() string
}

func main() {
	loadConfiguration()
	s := rewriteServer{}
	loadStaticRulesFromConfig(&s)
	loadReplaceRulesFromConfig(&s)
	loadDatabaseRuleRewritersFromConfig(&s)
	s.listRewriters()

	r := gin.Default()
	r.LoadHTMLFiles("landingpage.html")
	r.Use(s.handleRedirect)
	portNumber := viper.GetInt("app.port")
	serverURI := fmt.Sprintf(":%v", portNumber)
	r.Run(serverURI)
}

func loadConfiguration() {
	viper.SetDefault("app.debug", false)
	viper.SetDefault("app.port", "8080")
	viper.SetDefault("static-rules", []StaticRule{})

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
