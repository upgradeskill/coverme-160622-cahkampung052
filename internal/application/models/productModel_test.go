package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateDataDummy(t *testing.T) {
	GenerateDataDummy()
}

func TestStorage(t *testing.T) {
	products := Storage()
	assert.GreaterOrEqual(t, len(*products), 1, "Total product seharusnya lebih dari 1")
}
