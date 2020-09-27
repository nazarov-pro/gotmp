package conf

import (
	"os"
	"path"
	"testing"
)

func TestConfFileDir(t *testing.T) {
	// Current working directory
	dir, _ := os.Getwd()
	os.Setenv(ConfigFilePathKey, path.Dir(path.Dir(dir)))
	InitConfigs()
	expectedAppName := "go-template"
	actualAppName := GetString("app.name")
	if actualAppName != expectedAppName {
		t.Errorf("App name expected like '%s' but actual value is '%s'", expectedAppName, actualAppName)
	}
}

func TestConfFilePath(t *testing.T) {
	// Current working directory
	dir, _ := os.Getwd()
	os.Setenv(ConfigFilePathKey, path.Join(path.Dir(path.Dir(dir)), "configs", "app", "config.yaml"))
	InitConfigs()
	expectedAppName := "go-template"
	actualAppName := GetString("app.name")
	if actualAppName != expectedAppName {
		t.Errorf("App name expected like '%s' but actual value is '%s'", expectedAppName, actualAppName)
	}
}

func TestWithoutConfigFile(t *testing.T) {
	os.Setenv(ConfigFilePathKey, "")
	os.Setenv(ConfigEnvPrefixKey, "ENV")
	InitConfigs()
	expectedAppName := ""
	actualAppName := GetString("app.name")
	if actualAppName != expectedAppName {
		t.Errorf("App name expected like '%s' but actual value is '%s'", expectedAppName, actualAppName)
	}
}

func TestWrongConfigFile(t *testing.T) {
	os.Setenv(ConfigFilePathKey, "/")
	os.Setenv(ConfigEnvPrefixKey, "ENV")
	
	defer func() { 
		recover()
	}()
	
	InitConfigs()

	t.Errorf("The code did not panic")
}
