/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/rocky114/craftsman/internal/pkg/excel"
	"github.com/rocky114/craftsman/internal/storage"
	"github.com/spf13/cobra"
)

// execlCmd represents the execl command
var execlCmd = &cobra.Command{
	Use:   "excel",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		for _, school := range excel.GetSchools(fmt.Sprintf("%s/%s", rootDir, "assets/W020211027623974108131.xlsx")) {
			res, err := storage.GetQueries().CreateSchool(ctx, school)
			if err != nil {
				log.Fatalf("insert school err: %v", err)
			}
			fmt.Println(res.LastInsertId())
		}
	},
}

func init() {
	rootCmd.AddCommand(execlCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// execlCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// execlCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
