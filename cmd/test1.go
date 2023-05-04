/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// test1Cmd represents the test1 command
var test1Cmd = &cobra.Command{
	Use:   "test1",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("test1 called")
		test1()
	},
}

func init() {
	rootCmd.AddCommand(test1Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// test1Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// test1Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func test1() {
	// Define the contents of the config file
	config := `username: YOUR_USERNAME
password: YOUR_PASSWORD
server: YOUR_SERVER`

	// Create a file at ~/.deliveryhalf.yaml
	file, err := os.Create(os.Getenv("HOME") + "/.deliveryhalf.yaml")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Write the contents of the config file to the file
	_, err = file.WriteString(config)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Config file created successfully!")
}
