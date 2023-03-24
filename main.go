package main

import (
	"github.com/rios0rios0/locallaunch/util"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var log = logrus.New()

// FileReader is a struct that will handle the file reading
type FileReader struct {
	filePath string
}

// ExecCommand executes a command in the operating system

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Runs the 'up' commands specified in the yaml file",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("i entered in the function upCmd  ")
		fileReader := &FileReader{filePath: args[0]}
		yamlData, err := fileReader.ReadYAML()

		if err != nil {
			log.WithError(err).Error("Error reading YAML data")
			return
		}
		for _, cmd := range yamlData.Up {
			err = util.ExecCommand(cmd)
			if err != nil {
				log.WithError(err).Error("Error running command")
				return
			}
		}
		log.Info("Commands completed successfully")
	},
}

var downCmd = &cobra.Command{
	Use:   "down",
	Short: "Runs the 'down' commands specified in the yaml file",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("i entered in the function downCmd ")
		fileReader := &FileReader{filePath: args[0]}
		yamlData, err := fileReader.ReadYAML()
		if err != nil {
			log.WithError(err).Error("Error reading YAML data")
			return
		}
		for _, cmd := range yamlData.Down {
			err = util.ExecCommand(cmd)
			if err != nil {
				log.WithError(err).Error("Error running command")
				return
			}
		}
		log.Info("Commands completed successfully")
	},
}

func main() {

	var rootCmd = &cobra.Command{
		Use:   "lol [README.md]",
		Short: "LocalLaunch is a CLI to read a YAML slice of code inside a README.md file",
	}
	rootCmd.AddCommand(upCmd)
	rootCmd.AddCommand(downCmd)
	rootCmd.Execute()
}
