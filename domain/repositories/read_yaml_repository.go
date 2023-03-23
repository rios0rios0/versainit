package repositories

import "github.com/rios0rios0/locallaunch/infrastracture/models"

type IdentifyYAMLRepository interface {
	ReadYAML() (models.YamlData, error)
}
