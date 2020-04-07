package models

import (
	"github.com/NyanKiyoshi/pytest-django-queries-bot/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEvent_Table(t *testing.T) {
	table := EventTable()
	assert.Equal(t, config.DynamoEventsTableName, table.Name(), "Unexpected table name")
}
