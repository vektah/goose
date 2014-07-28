package cmd

import (
	"bitbucket.org/liamstask/goose/lib/goose"
	"fmt"
	"log"
)

var dbVersionCmd = &Command{
	Name:    "dbversion",
	Usage:   "",
	Summary: "Print the current version of the database",
	Help:    `dbversion extended help here...`,
	Run:     dbVersionRun,
}

func dbVersionRun(cmd *Command, conf *goose.DBConf, args ...string) {

	current, err := goose.GetDBVersion(conf)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("goose: dbversion %v\n", current)
}
