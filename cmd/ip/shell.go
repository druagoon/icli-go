/*
Copyright © 2023 Zonglong Fan <lazyboy.fan@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package ip

import (
	"fmt"
	"net"
	"strings"

	"github.com/spf13/cobra"
)

// shellCmd represents the shell command
var shellCmd = &cobra.Command{
	Use:   "shell",
	Short: "show shell ip prompt",
	Run: func(cmd *cobra.Command, args []string) {
		prompt := []string{}
		interfaces, err := net.Interfaces()
		if err != nil {
			return
		}
		for _, itf := range interfaces {
			ipv4 := GetInterfaceIpv4Addr(itf)
			if ipv4 != "" {
				line := fmt.Sprintf("%s=%s", itf.Name, ipv4)
				prompt = append(prompt, line)
			}
		}
		output := strings.Join(prompt, " ")
		fmt.Println(output)
	},
}

func GetInterfaceIpv4Addr(itf net.Interface) string {
	addrs, err := itf.Addrs()
	if err != nil {
		return ""
	}
	for _, a := range addrs {
		if v, ok := a.(*net.IPNet); ok && !v.IP.IsLoopback() {
			ipv4 := v.IP.To4()
			if ipv4 != nil {
				return ipv4.String()
			}
		}
	}
	return ""
}

func init() {
	ipCmd.AddCommand(shellCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// shellCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// shellCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
