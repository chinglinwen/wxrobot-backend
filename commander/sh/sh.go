package basic

import (
	"fmt"
	"os/exec"

	"github.com/chinglinwen/wechat-commander/commander"
)

type Sh struct {
}

func (*Sh) Command(cmd string) (out string, err error) {
	c := exec.Command("bash", "-l", "-c", cmd)

	output, err := c.CombinedOutput()
	if err != nil {
		err = fmt.Errorf("execute cmds: %v\noutput: %v", err, string(output))
		return
	}
	out = string(output)

	return
}

func (*Sh) Help() string {
	return "execute shell commands"
}

func init() {
	commander.Register("sh", &Sh{})
}
