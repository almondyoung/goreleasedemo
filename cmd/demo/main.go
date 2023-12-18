package main

import (
	"fmt"
	"github.com/almondyoung/goreleasedemo/cmd"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "myapp",
		Short: "A simple command-line application",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello, World!")
		},
	}

	greetCmd := &cobra.Command{
		Use:   "greet",
		Short: "Greet someone",
		Run: func(cmd *cobra.Command, args []string) {
			name, _ := cmd.Flags().GetString("name")
			fmt.Printf("Hello, %s!\n", name)
		},
	}

	greetCmd.Flags().StringP("name", "n", "", "Name of the person to greet")

	rootCmd.AddCommand(greetCmd)
	rootCmd.AddCommand(cmd.NewVersionCmd())

	rootCmd.Execute()
}
