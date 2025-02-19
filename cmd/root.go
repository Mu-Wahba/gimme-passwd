package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/mu-wahba/gimme-passwd/utils"
	"github.com/spf13/cobra"
)

var length int
var description string
var line string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gimme-pass",
	Short: "Generate Random password with fixed lenght of 12 and coan be configured",
	Long:  `Generate password and save it with a note into a file and can be uploaded to s3 as a backup`,
	Run: func(cmd *cobra.Command, args []string) {
		if length < 20 && length >= 8 {
			pass := utils.GeneratePasswd(length)
			green := color.New(color.FgGreen).SprintFunc() // Green color formatter

			fmt.Println("Password -->", green(pass))
			// fmt.Println("Password -->", pass)

			//append to file
			timezone, err := time.LoadLocation(os.Getenv("TIME_ZONE"))
			if err != nil {
				timezone = time.UTC
			}
			line = fmt.Sprintf("Password:  %s  ,Added at [ %s ]\n", pass, time.Now().In(timezone).Format("2006-01-02 15:04:05"))
			if description != "" {
				line = fmt.Sprintf("Password:  %s  ,Added at [%s] , Description: [%s] \n", pass, time.Now().In(timezone).Format("2006-01-02 15:04:05"), description)
			}
			utils.AppendToFile(os.Getenv("PASSWD_FILE"), line)

		} else {
			red := "\033[31m"
			reset := "\033[0m" // Resets the color
			fmt.Println(string(red), "Error: The allowed value for --length or -l is between 8 and 20.", string(reset))
			os.Exit(1)

		}
		// utils.AppendToFile

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().IntVarP(&length, "length", "l", 8, "Password length default is 8 and max is 20")
	rootCmd.Flags().StringVarP(&description, "descr", "d", "", "password description")
}
