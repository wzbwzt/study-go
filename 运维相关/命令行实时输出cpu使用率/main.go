package main

import (
	"fmt"
	"os"
	"time"

	"github.com/olekukonko/tablewriter"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/spf13/cobra"
)

//命令行框架:https://github.com/spf13/cobra/blob

//系统命令第三方工具：https://github.com/shirou/gopsutil

//终端输出表格：https://github.com/olekukonko/tablewriter

var rootCmd = &cobra.Command{
	Use:   "info",
	Short: "get info(cpu,mem) per second",
	Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at http://hugo.spf13.com`,
	Run: func(cmd *cobra.Command, args []string) {

		// file, _ := os.OpenFile("./tableFile.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 666)
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Name", "Count", "Rating"})

		c_pre, _ := cpu.Percent(time.Second, true)
		c_count, _ := cpu.Counts(true)

		table.Append([]string{"CPU", fmt.Sprint(c_count), fmt.Sprintf("%.3f%%", c_pre[0])})

		m, _ := mem.VirtualMemory()
		m_total := m.Total
		m_pre := m.UsedPercent

		table.Append([]string{"Mem", fmt.Sprint(m_total), fmt.Sprintf("%.3f%%", m_pre)})

		table.Render()

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

//##############################################################################
func main() {
	Execute()

}
