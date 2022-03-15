package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var cmds = &cobra.Command{}
var sshcmd = &cobra.Command{
	Use:   "ssh",
	Short: "ssh connect",
	Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at http://hugo.spf13.com`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	cmds.AddCommand(sshcmd)
}

func main() {
	list := make([]int, 10)
	alist := list
	for i := 0; i < 10; i++ {
		alist[i] = i
	}
	fmt.Println(list)
	fmt.Println(alist)

	return
	cmds.Execute()
}
