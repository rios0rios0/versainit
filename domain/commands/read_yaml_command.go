package command

import (
	"github.com/rios0rios0/locallaunch/domain/repositories"
	"github.com/rios0rios0/locallaunch/infrastracture/models"
)

type ReadYAMLCommand struct {
	repository repositories.IdentifyYAMLRepository
}

func NewReadYAMLCommand(repository repositories.IdentifyYAMLRepository) *ReadYAMLCommand {
	return &ReadYAMLCommand{repository}
}

func (itself ReadYAMLCommand) Excute() (models.YamlData, error) {
	return itself.repository.ReadYAML()
}
