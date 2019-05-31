package models

import (
	"github.com/stretchr/testify/assert"
	"pytest-queries-bot/pytestqueries/generated"
	"testing"
)

func TestEvent_Table(t *testing.T) {
	table := EventTable()
	assert.Equal(t, generated.DynamoEventsTableName, table.Name(), "Unexpected table name")
}
