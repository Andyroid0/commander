package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

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
		arg := args[0]
		cmd_str := args[1:]

		if key != arg {
			continue
		}

		cmd := exec.Command("bash", "-c", strings.Join(append([]string{cmd.(string)}, cmd_str...), " "))
		out, err := cmd.Output()
		if err != nil {
			fmt.Println(err.Error())
			break
		}

		fmt.Println(string(out))
		break
	}
}
