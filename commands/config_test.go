package commands

import "testing"

var testPathToConfig string

func TestFindConfig(t *testing.T) {
	testPathToConfig = FindConfig()
	t.Logf("Config found in %s...", testPathToConfig)
}

func TestLoadConfigFromFile(t *testing.T) {
	cfg, err := LoadConfigFromFile(testPathToConfig)
	if err != nil {
		t.Errorf("%s : while reading config from %s", err, testPathToConfig)
	}
	for k, v := range cfg {
		t.Logf("Sink %s found for bot token %s and chat %v", k, v.Token, v.ChatID)
	}
}
