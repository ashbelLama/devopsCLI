/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/ashbelLama/cli-golang/controller"
	"github.com/spf13/cobra"
)

// jokeCmd represents the joke command
var jokeCmd = &cobra.Command{
	Use:   "joke",
	Short: "Get a random dad joke",
	Long:  `Its a joke, stop being so serious.`,
	Run: func(cmd *cobra.Command, args []string) {
		controller.JokeController()
	},
}

func init() {
	rootCmd.AddCommand(jokeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// jokeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// jokeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
