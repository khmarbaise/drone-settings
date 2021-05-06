package settings

import (
	"encoding/xml"
	"io/ioutil"
	"os"
)

//Settings defines the whole structure of a https://maven.apache.org/settings.html file.
type Settings struct {
	XMLName        xml.Name `xml:"settings"`
	Text           string   `xml:",chardata"`
	Xmlns          string   `xml:"xmlns,attr"`
	Xsi            string   `xml:"xsi,attr"`
	SchemaLocation string   `xml:"schemaLocation,attr"`
	Mirrors        struct {
		Text   string `xml:",chardata"`
		Mirror struct {
			Text     string `xml:",chardata"`
			ID       string `xml:"id"`
			MirrorOf string `xml:"mirrorOf"`
			URL      string `xml:"url"`
		} `xml:"mirror"`
	} `xml:"mirrors"`
	Profiles struct {
		Text    string `xml:",chardata"`
		Profile []struct {
			Text         string `xml:",chardata"`
			ID           string `xml:"id"`
			Repositories struct {
				Text       string `xml:",chardata"`
				Repository []struct {
					Text     string `xml:",chardata"`
					ID       string `xml:"id"`
					URL      string `xml:"url"`
					Releases struct {
						Text           string `xml:",chardata"`
						Enabled        string `xml:"enabled"`
						ChecksumPolicy string `xml:"checksumPolicy"`
					} `xml:"releases"`
					Snapshots struct {
						Text           string `xml:",chardata"`
						Enabled        string `xml:"enabled"`
						ChecksumPolicy string `xml:"checksumPolicy"`
					} `xml:"snapshots"`
				} `xml:"repository"`
			} `xml:"repositories"`
			PluginRepositories struct {
				Text             string `xml:",chardata"`
				PluginRepository []struct {
					Text     string `xml:",chardata"`
					ID       string `xml:"id"`
					URL      string `xml:"url"`
					Releases struct {
						Text           string `xml:",chardata"`
						Enabled        string `xml:"enabled"`
						ChecksumPolicy string `xml:"checksumPolicy"`
					} `xml:"releases"`
					Snapshots struct {
						Text           string `xml:",chardata"`
						Enabled        string `xml:"enabled"`
						ChecksumPolicy string `xml:"checksumPolicy"`
					} `xml:"snapshots"`
				} `xml:"pluginRepository"`
			} `xml:"pluginRepositories"`
			Properties struct {
				Text       string `xml:",chardata"`
				GpgKeyname string `xml:"gpg.keyname"`
			} `xml:"properties"`
		} `xml:"profile"`
	} `xml:"profiles"`
	ActiveProfiles struct {
		Text          string `xml:",chardata"`
		ActiveProfile string `xml:"activeProfile"`
	} `xml:"activeProfiles"`
	Servers struct {
		Text   string `xml:",chardata"`
		Server []struct {
			Text                 string `xml:",chardata"`
			ID                   string `xml:"id"`
			Username             string `xml:"username"`
			PrivateKey           string `xml:"privateKey"`
			Password             string `xml:"password"`
			FilePermissions      string `xml:"filePermissions"`
			DirectoryPermissions string `xml:"directoryPermissions"`
		} `xml:"server"`
	} `xml:"servers"`
}

//ReadSettings Reads the `settings.xml` file to be used for enhancements etc.
func ReadSettings(fileName string) (result Settings, err error) {
	xmlFile, err := os.Open(fileName)
	if err != nil {
		return Settings{}, err
	}

	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	err = xml.Unmarshal(byteValue, &result)
	return result, err
}
