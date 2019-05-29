package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEvent_Table(t *testing.T) {
	table := EventTable()
	assert.Equal(t, tableName, table.Name(), "Unexpected table name")
}
