package cmd

import "github.com/spf13/cobra"

//
var rootCmd = &cobra.Command{}

func Execute() error {
	return rootCmd.Execute()
}

//在下面的init函数中进行相对应的注册
func init() {
	rootCmd.AddCommand(wordCmd)
	rootCmd.AddCommand(timeCmd)
	rootCmd.AddCommand(sqlCmd) //进行json2struct的注册
}
