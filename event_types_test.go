package footballdata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_UpdateDescriptor_Fields(t *testing.T) {
	testValue := UpdateDescriptor("status|IN_PLAY -> FINISHED")
	fields := testValue.Fields()

	assert.Len(t, fields, 1)

	assert.Equal(t, "status", fields[0].Name)
	assert.NotNil(t, fields[0].Values)
	assert.Equal(t, "IN_PLAY", fields[0].Values.OldValue)
	assert.Equal(t, "FINISHED", fields[0].Values.NewValue)
}

func Test_UpdateDescriptor_Fields_Multiple(t *testing.T) {
	testValue := UpdateDescriptor("status|IN_PLAY -> FINISHED;score|1:0 -> 1:1")
	fields := testValue.Fields()

	assert.Len(t, fields, 2)

	assert.Equal(t, "status", fields[0].Name)
	assert.NotNil(t, fields[0].Values)
	assert.Equal(t, "IN_PLAY", fields[0].Values.OldValue)
	assert.Equal(t, "FINISHED", fields[0].Values.NewValue)

	assert.Equal(t, "score", fields[1].Name)
	assert.NotNil(t, fields[1].Values)
	assert.Equal(t, "1:0", fields[1].Values.OldValue)
	assert.Equal(t, "1:1", fields[1].Values.NewValue)
}

func Test_UpdateDescriptor_Fields_NameOnly(t *testing.T) {
	testValue := UpdateDescriptor("status")
	fields := testValue.Fields()

	assert.Len(t, fields, 1)
	assert.Equal(t, "status", fields[0].Name)
	assert.Nil(t, fields[0].Values)
}

func Test_UpdateDescriptor_Fields_NameOnly_Multiple(t *testing.T) {
	testValue := UpdateDescriptor("status;leagueTable")
	fields := testValue.Fields()

	assert.Len(t, fields, 2)

	assert.Equal(t, "status", fields[0].Name)
	assert.Nil(t, fields[0].Values)

	assert.Equal(t, "leagueTable", fields[1].Name)
	assert.Nil(t, fields[1].Values)
}
