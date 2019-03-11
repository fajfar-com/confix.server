package internal

import "strings"

type Configuration struct {
	Key string
	Value string
}

type Application struct {
	Name string
	Configs []Configuration
}

func (a Application) GetConfigByName(name string) Configuration {
	for _, c := range a.Configs{
		if strings.Compare(name, c.Key) == 0{
			return c
		}
	}
	return Configuration{
		Key:name,
		Value:"",
	}
}