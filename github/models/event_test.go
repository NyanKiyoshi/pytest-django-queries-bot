package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEvent_Table(t *testing.T) {
	event := Event{}
	table := event.Table()
	assert.Equal(t, tableName, table.Name(), "Unexpected table name")
}
