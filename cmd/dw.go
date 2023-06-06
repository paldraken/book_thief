package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/paldraken/book_thief/internal/download"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var WorkUrl string
var UserName string
var Password string
var OutputDir string

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(dwCmd)
	dwCmd.Flags().StringVarP(&WorkUrl, "work", "w", "", "Link to the book page")

	dwCmd.Flags().StringVarP(&OutputDir, "output", "o", "", "The output directory is where the downloaded book will be stored.")
	viper.BindPFlag("output", dwCmd.PersistentFlags().Lookup("output"))

	dwCmd.PersistentFlags().StringVarP(&UserName, "author_today.username", "u", "", "User name")
	viper.BindPFlag("author_today.username", dwCmd.PersistentFlags().Lookup("author_today.username"))

}

var dwCmd = &cobra.Command{
	Use:   "dw",
	Short: "Download book console command",
	Run: func(cmd *cobra.Command, args []string) {

		if OutputDir == "" {
			path, err := os.Getwd()
			if err != nil {
				log.Println(err)
			}
			OutputDir = path
			viper.Set("output", path)
		}

		download.Console(WorkUrl)
	},
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName("bookthief.yaml")
	}

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("config error", err)
	} else {
		// fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
