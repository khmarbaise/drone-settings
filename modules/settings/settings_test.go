package settings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ReadSettings_Success(t *testing.T) {
	fileName := "testdata/ExampleSettings.xml"

	_, err := ReadSettings(fileName)

	assert.Nil(t, err, "There should be no error because we are reading an existing "+
		"file from testdata directory.")
}

func Test_ReadSettings_MirrorOf(t *testing.T) {
	fileName := "testdata/ExampleSettings.xml"

	settings, err := ReadSettings(fileName)

	assert.Nil(t, err, "There should be no error because we are reading an existing "+
		"file from testdata directory.")
	assert.Equal(t, "nexus", settings.Mirrors.Mirror.ID)
	assert.Equal(t, "*", settings.Mirrors.Mirror.MirrorOf)
	assert.Equal(t, "http://localhost:8081/nexus/content/groups/public", settings.Mirrors.Mirror.URL)
}

func Test_ReadSettings_Profiles(t *testing.T) {
	fileName := "testdata/ExampleSettings.xml"

	settings, err := ReadSettings(fileName)

	assert.Nil(t, err, "There should be no error because we are reading an existing "+
		"file from testdata directory.")
	assert.Equal(t, 2, len(settings.Profiles.Profile))
	assert.Equal(t, "nexus", settings.Profiles.Profile[0].ID)
	assert.Equal(t, 1, len(settings.Profiles.Profile[0].Repositories.Repository))
	assert.Equal(t, 1, len(settings.Profiles.Profile[0].PluginRepositories.PluginRepository))
	assert.Equal(t, "http://localhost:8081/nexus/content/groups/public", settings.Mirrors.Mirror.URL)
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
