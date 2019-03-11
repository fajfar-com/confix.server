package configMng

import (
	"github.com/fajfar-com/confix.server/internal"
	"strings"
)

type ConfigurationManager interface {
	GetApplication(appName string) internal.Application
}

type configurationManager struct {
	applications []internal.Application
}

func GetConfigurationManager() ConfigurationManager {
	return configurationManager{
		applications: []internal.Application{{
			Name: "test",
			Configs: []internal.Configuration{
				{
					Key:"config1",
					Value:"value1",
				},
				{
					Key:"config2",
					Value:"value2",
				},
				{
					Key:"config3",
					Value:"value3",
				},
				{
					Key:"config4",
					Value:"value4",
				},
			},
		}},
	}
}

func (cm configurationManager) GetApplication(appName string) internal.Application {
	for _, a := range cm.applications{
		if strings.Compare(a.Name, appName) == 0{
			return a
		}
	}
	return internal.Application{
		Name: appName,
		Configs: []internal.Configuration{},
	}
}