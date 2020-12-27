package settings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ReadSettings_Success(t *testing.T) {
	fileName := "testdata/ExampleSettings.xml"

	settings, err := ReadSettings(fileName)

	assert.Nil(t, err, "There should be no error because we are reading an existing "+
		"file from testdata directory.")
	assert.Equal(t, settings.Mirrors.Mirror.ID, "nexus")
	assert.Equal(t, settings.Mirrors.Mirror.MirrorOf, "*")
	assert.Equal(t, settings.Mirrors.Mirror.URL, "http://localhost:8081/nexus/content/groups/public")
}

func Test_ReadSettings_Failure(t *testing.T) {
	var fileName = "testdata/ExampleSettingsX.xml"

	settings, err := ReadSettings(fileName)

	assert.NotNil(t, err, "There should be an error because we are trying to read a "+
		"file from testdata directory which does not exist.")
	assert.NotNil(t, settings)
	assert.Empty(t, settings.Mirrors.Mirror.ID)
	assert.Empty(t, settings.Mirrors.Mirror.MirrorOf)
	assert.Empty(t, settings.Mirrors.Mirror.URL)
}
