package tests

import (
	"github.com/premsvmm/goapi/controllers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGet(t *testing.T) {
	res := controllers.Hostname
	assert.NotNil(t,res)
}
