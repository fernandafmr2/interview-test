package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCon(t *testing.T) {
	db := SetUpDatabaseConnection(New("../env.test"))
	assert.NotNil(t, db)
}