package main

import (
	"go.arpabet.com/cligo"
	"go.arpabet.com/glue"
)

type Run struct {
	Parent cligo.CliGroup `cli:"group=cli"`
	Name   string         `cli:"argument=name"`
	Print  bool           `cli:"option=print,default=false,help=prints it"`
}

func (cmd *Run) Command() string {
	return "run"
}

func (cmd *Run) Help() (string, string) {
	return "Runs the program.", `This command runs the program by using cligo framework.
This framework is a game changer in the CLI tools.`
}

func (cmd *Run) Run(ctx glue.Context) error {
	if cmd.Print {
		cligo.Echo("Invokes '%s' command with name '%s' argument", cmd.Command(), cmd.Name)
	}
	return nil
}

func main() {
	cligo.Main(cligo.Beans(&Run{}))
}
