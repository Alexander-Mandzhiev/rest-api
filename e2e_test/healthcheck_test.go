package e2e_test

import (
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ExampleTestSuite struct {
	suite.Suite
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ExampleTestSuite))
}

func (s *ExampleTestSuite) TestHappyHealthCheck() {
	c := http.Client{}

	res, _ := c.Get("http://localhost:6000/health-check")

	s.Equal(http.StatusOK, res.StatusCode)
	bodyRes, _ := io.ReadAll(res.Body)
	s.JSONEq(`{"status":"OK", "messages":[]}`, string(bodyRes))
}
