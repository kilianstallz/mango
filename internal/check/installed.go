package check

import (
	"github.com/spf13/viper"
	"log"
	"mango/config"
	"os"
	"regexp"
	"sort"
)

func ListInstalledVersions() ([]string, error) {
	// get config from viper
	installRoot := viper.GetString(config.BinRoot)

	// get contents of installRoot folder
	entries, err := os.ReadDir(installRoot)
	if err != nil {
		log.Fatal(err)
	}

	var versions []string
	for _, entry := range entries {
		// check if entry is a directory with a name that starts with "go*.x.x"
		if entry.IsDir() && useRegex(entry.Name()) {
			versions = append(versions, entry.Name())
		}
	}
	return versions, nil
}

func GetLatestVersion(versions []string) string {
	list := versions
	// if the list is empty, return an empty string
	if len(list) == 0 {
		return ""
	}
	// if the list has only one element, return that element
	if len(list) == 1 {
		return list[0]
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i] < list[j]
	})
	last := list[len(list)-1]
	return last
}

func IsUpdateAvailable() (string, error) {
	versions, err := ListInstalledVersions()
	if err != nil {
		return "", err
	}

	latest := GetLatestVersion(versions)

	latestRelease, err := QueryLatestGoVersion()
	if err != nil {
		return "", err
	}

	if latest != latestRelease {
		return latestRelease, nil
	}
	return "", nil
}

func useRegex(s string) bool {
	re := regexp.MustCompile("([A-Za-z0-9]+(\\.[A-Za-z0-9]+)+)")
	return re.MatchString(s)
}
