package repositories

import (
	"github.com/rios0rios0/locallaunch/infrastracture/models"
	"github.com/sirupsen/logrus"
)

type ReadYAMLRepository struct{}

func NewReadYAMLRepository() ReadYAMLRepository {
	return ReadYAMLRepository{}
}

func (itself ReadYAMLRepository) ReadYAML() (models.YamlData, error) {
	logrus.Info("initialized Read YAML")

	model := models.YamlData{}

	return model, nil
}
