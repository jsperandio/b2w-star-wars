package config_test

import (
	"os"
	"testing"

	_cfg "github.com/jsperandio/b2w-star-wars/planet/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ConfigSuite struct {
	config *_cfg.Config
	suite.Suite
}

// The SetupSuite method will be run by testify once, at the very
// start of the testing suite, before any tests are run.
func (c *ConfigSuite) SetupSuite() {
	os.Setenv("test_sucess_int", "11")
	os.Setenv("test_sucess_string", "onze")
}

// The SetupTest method will be run before every test in the suite.
func (c *ConfigSuite) SetupTest() {
	var cnfg _cfg.Config
	c.config = &cnfg
}

// The TearDownSuite method will be run by testify once, at the very
// end of the testing suite, after all tests have been run.
func (c *ConfigSuite) TearDownSuite() {
	os.Unsetenv("test_sucess_int")
	os.Unsetenv("test_sucess_string")
}

// TestInitConfig will check if the default values are ok in the init
func (c *ConfigSuite) TestInitConfig() {
	c.config.InitConfig()

	assert := assert.New(c.T())

	assert.Equal("mongodb://localhost:27017", c.config.Mongodb_connetion_string)
	assert.Equal(2, c.config.Rest_max_retry)
	assert.Equal(2, c.config.Rest_wait_sec)
	assert.Equal(10, c.config.Rest_max_wait_sec)
}

// TestFallback will test if the Fallback is correct on return
func (c *ConfigSuite) TestFallback() {
	assert := assert.New(c.T())

	assert.Equal(10, c.config.GetEnvAsIntOrFallback("test_int", 10))
	assert.Equal("dez", c.config.GetEnvAsStringOrFallback("test_string", "dez"))
}

// TestGetEnvSucess will check if the Env is load correctly
func (c *ConfigSuite) TestGetEnvSucess() {
	assert := assert.New(c.T())

	assert.Equal(11, c.config.GetEnvAsIntOrFallback("test_sucess_int", 10))
	assert.Equal("onze", c.config.GetEnvAsStringOrFallback("test_sucess_string", "dez"))
}

// TestConfigSuite is the function to kick off the test suite
func TestConfigSuite(t *testing.T) {
	suite.Run(t, new(ConfigSuite))
}
