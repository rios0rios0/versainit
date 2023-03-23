package command

import (
	"github.com/rios0rios0/locallaunch/domain/repositories"
	"github.com/spf13/cobra"
)

type UpCMDCommand struct {
	repository repositories.UpCMDRepository
}

func NewUpCMDCommand(repository repositories.UpCMDRepository) *UpCMDCommand {
	return &UpCMDCommand{repository}
}

func (itself UpCMDCommand) ExecuteUp() *cobra.Command {
	return itself.repository.UpCMD()
}
