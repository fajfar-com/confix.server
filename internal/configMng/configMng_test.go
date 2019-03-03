package configMng

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetApplication(t *testing.T) {
	app := GetApplication("test")

	assert.Equal(t,"test", app.Name)

	assert.Contains(t, app.Configs, Configuration{
		Key:"name",
		Value:"Name",
	})
}