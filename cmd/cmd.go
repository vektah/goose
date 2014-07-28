package cmd

import (
	"bitbucket.org/liamstask/goose/lib/goose"
	"flag"
)

// shamelessly snagged from the go tool
// each command gets its own set of args,
// defines its own entry point, and provides its own help
type Command struct {
	Run  func(cmd *Command, conf *goose.DBConf, args ...string)
	Flag flag.FlagSet

	Name  string
	Usage string

	Summary string
	Help    string
}

func (c *Command) Exec(conf *goose.DBConf, args []string) {
	c.Flag.Usage = func() {
		// helpFunc(c, c.Name)
	}
	c.Flag.Parse(args)
	c.Run(c, conf, c.Flag.Args()...)
}

var commands = []*Command{
	upCmd,
	downCmd,
	redoCmd,
	statusCmd,
	createCmd,
	dbVersionCmd,
}

func Commands() []*Command {
	return commands
}
