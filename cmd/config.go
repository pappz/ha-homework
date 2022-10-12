package main

import (
	"strconv"
	"strings"

	"github.com/spf13/viper"
)

const (
	envPrefix = "ha"

	cfgKeySectorID      = "sector_id"
	cfgKeyListenAddress = "address"
)

type config struct {
	listenAddress string
	sectorId      *int
}

func (c *config) bindEnvs() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix(envPrefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}

func (c *config) setDefaults() {
	viper.SetDefault(cfgKeyListenAddress, ":8080")
	viper.SetDefault(cfgKeySectorID, "")
}

func (c *config) readVariables() {
	c.listenAddress = viper.GetString(cfgKeyListenAddress)
	_, err := strconv.Atoi(viper.GetString(cfgKeySectorID))
	if err == nil {
		id := viper.GetInt(cfgKeySectorID)
		c.sectorId = &id
	}
}

func newConfig() config {
	cfg := config{}
	cfg.bindEnvs()
	cfg.setDefaults()
	cfg.readVariables()
	return cfg
}
