package repositories

import (
	"github.com/rios0rios0/locallaunch/infrastracture/repositories"
	"github.com/rios0rios0/locallaunch/util"
	logger "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type UpCMDRepository struct {
	repositoryReadYAML repositories.ReadYAMLRepository
}

func NewUpCMDRepository(repositoryReadYAML repositories.ReadYAMLRepository) *UpCMDRepository {
	return &UpCMDRepository{repositoryReadYAML}
}
func (itself UpCMDRepository) UpCMD() *cobra.Command {

	upCmd := &cobra.Command{
		Use:   "up",
		Short: "Runs the 'up' commands specified in the yaml file",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			logger.Info("i entered in the function upCmd  ")
			fileReader := &FileReader{filePath: args[0]}

			yamlData, err := itself.repositoryReadYAML.ReadYAML()

			if err != nil {
				logger.Errorf("up cmd - can't reading YAML data. Here the reason: %s", err)
				return
			}
			for _, cmd := range yamlData.Up {
				err = util.ExecCommand(cmd)
				if err != nil {
					logger.Errorf("up cmd  - not range in yamlData. Hete the reason: %s, ", err)
					return
				}
			}
			logger.Info("up cmd - Commands completed successfully")
		},
	}
	return upCmd
}
