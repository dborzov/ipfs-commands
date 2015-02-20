package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/template"

	cmds "github.com/jbenet/go-ipfs/commands"
	commands "github.com/jbenet/go-ipfs/core/commands"
)

const templateFilename = "TEMPLATE.md"

// command name is stored as map ket within the Root.Subcommands
type cmdSpec struct {
	Name string
	Cmd  *cmds.Command
}

func main() {
	t, err := ioutil.ReadFile(templateFilename)
	if err != nil {
		fmt.Printf("Failed to open template file %s: %s \n", templateFilename, err)
	}

	var specs []cmdSpec // convert root command map into array so that one can sort them at the output
	for i := range commands.Root.Subcommands {
		specs = append(specs, cmdSpec{
			Name: i,
			Cmd:  commands.Root.Subcommands[i],
		})
	}

	p, err := template.New("commands").Parse(string(t))
	if err != nil {
		log.Println("Bollocks! Failure parsing template:", err)
	}

	err = p.Execute(os.Stdout, specs)
	if err != nil {
		log.Println("Yarr!!! Failure executing template:", err)
	}
}
