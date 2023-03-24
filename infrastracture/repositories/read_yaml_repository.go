package repositories

import (
	"github.com/rios0rios0/locallaunch/infrastracture/models"
	logger "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type ReadYAMLRepository struct {
	filePath string
}

func NewReadYAMLRepository() ReadYAMLRepository {
	return ReadYAMLRepository{}
}

func (itself ReadYAMLRepository) ReadYAML() (models.YamlData, error) {
	logger.Info("i entered in the function readyaml")

	data, err := ioutil.ReadFile(itself.filePath)
	if err != nil {
		logger.Errorf("read yaml - can't reading file. Here the reason: %s", err)
		return models.YamlData{}, err
	}
	var yamlData models.YamlData
	err = yaml.Unmarshal(data, &yamlData)
	if err != nil {
		logger.Errorf("Error parsing YAML %s", err)
		return models.YamlData{}, err
	}
	return yamlData, nil

	return yamlData, nil
}
