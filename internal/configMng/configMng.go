package configMng

import (
	"strings"
)

type Configuration struct {
	Key string
	Value string
}

type Application struct {
	Name string
	Configs []Configuration
}

var applications []Application

func init(){
	applications = []Application{{
		Name: "test",
		Configs: []Configuration{
			{
				Key:"name",
				Value:"Name",
			},
		},
	}}
}

func  GetApplication(appName string) Application {
	for _, a := range applications{
		if strings.Compare(a.Name, appName) == 0{
			return a
		}
	}
	return Application{
		Name: appName,
		Configs: []Configuration{},
	}
}

func (a Application) GetConfigByName(name string) Configuration {
	for _, c := range a.Configs{
		if strings.Compare(name, c.Key) == 0{
			return c
		}
	}
	return Configuration{
		Key:name,
		Value:nil,
	}
}