package access_log

import (
	. "launchpad.net/gocheck"
	"github.com/cloudfoundry/gorouter/config"
)

type CreateRunningAccessLoggerSuite struct{}
var _ = Suite(&CreateRunningAccessLoggerSuite{})

func (s *CreateRunningAccessLoggerSuite) TestDoesNotCreateAnAccessLoggerIfNoAccesLogAndNoLoggregatorUrl(c *C) {
	config := config.DefaultConfig()
	c.Assert(CreateRunningAccessLogger(config), IsNil)
}

func (s *CreateRunningAccessLoggerSuite) TestCreatesAnAccessLoggerIfNoAccesLogButLoggregatorUrl(c *C) {
	config := config.DefaultConfig()
	config.LoggregatorConfig.Url = "10.10.3.13:4325"
	config.AccessLog = ""
	c.Assert(CreateRunningAccessLogger(config), NotNil)
}

func (s *CreateRunningAccessLoggerSuite) TestCreatesAnAccessLoggerIfAccesLogButNoLoggregatorUrl(c *C) {
	config := config.DefaultConfig()
	config.AccessLog = "/dev/null"
	c.Assert(CreateRunningAccessLogger(config), NotNil)
}

func (s *CreateRunningAccessLoggerSuite) TestCreatesAnAccessLoggerIfBothAccesLogAndLoggregatorUrl(c *C) {
	config := config.DefaultConfig()
	config.LoggregatorConfig.Url = "10.10.3.13:4325"
	config.AccessLog = "/dev/null"
	c.Assert(CreateRunningAccessLogger(config), NotNil)
}

func (s *CreateRunningAccessLoggerSuite) TestPanicsIfInvalidAccessLogLocation(c *C) {
	config := config.DefaultConfig()
	config.AccessLog = "/this\\should/panic"
	c.Assert(func() {
			CreateRunningAccessLogger(config)
		}, PanicMatches, "open /this\\\\should/panic: no such file or directory")
}