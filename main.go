package main

import (
	"fmt"
	"os"
	"os/exec"

	toml "github.com/pelletier/go-toml/v2"
)

func main() {
	args := os.Args[1:]

	cmds, err := os.ReadFile("./commands.toml")
	if err != nil {
		fmt.Println("Missing commands toml file.")
		return
	}

	var cmds_go map[string]interface{}

	toml_err := toml.Unmarshal(cmds, &cmds_go)
	if toml_err != nil || cmds_go == nil {
		fmt.Println("Unable to parse commands toml file.")
		return
	}

	for key, cmd := range cmds_go {
		for _, arg := range args {
			if key != arg {
				continue
			}
			cmd := exec.Command("bash", "-c", cmd.(string))
			out, err := cmd.Output()
			if err != nil {
				fmt.Println(err.Error())
				break
			}
			fmt.Println(string(out))
			break
		}
	}
}
