package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/spf13/viper"
)

type databaseRuleRewriter struct {
	db       *gorm.DB
	filePath string
}

type databaseRuleRewriterConfig struct {
	Path string
}

func newDatabaseRuleRewriter(filePath string) *databaseRuleRewriter {
	db, err := gorm.Open("sqlite3", filePath)
	if err != nil {
		panic("failed to connect database")
	}

	if viper.GetBool("app.debug") {
		db.LogMode(true)
	}

	db.AutoMigrate(&StaticRule{})
	//defer db.Close()
	return &(databaseRuleRewriter{
		db:       db,
		filePath: filePath,
	})
}

func (r databaseRuleRewriter) rewrite(path string) *rewriteResult {
	var rule StaticRule
	if err := r.db.Where(&StaticRule{From: path}).First(&rule).Error; err != nil {
		return nil
	}
	return &rewriteResult{rule.To, false}
}

func (r databaseRuleRewriter) String() string {
	var count int
	r.db.Model(&StaticRule{}).Count(&count)
	return fmt.Sprintf("[db] %s. %d rules", r.filePath, count)
}

func (r databaseRuleRewriter) Add(from, to string) {
	r.db.Create(&StaticRule{From: from, To: to})
}

func loadDatabaseRuleRewritersFromConfig(r *rewriteServer) {
	var rules []databaseRuleRewriterConfig
	err := viper.UnmarshalKey("database-rewriters", &rules)

	if err != nil {
		fmt.Printf("Could not parse database-rewriters from config file: %s \n", err)
		return
	}
	for _, rule := range rules {
		r.appendRewriter(newDatabaseRuleRewriter(rule.Path))
	}
}
