package utils

import "testing"

func TestConfigFileLoader(t *testing.T) {
	ConfigFileLoader("/Users/will/.rinvokerc", "uname")
}
