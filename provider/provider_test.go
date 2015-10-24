package provider

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func getData() []map[string]string {
	return []map[string]string{
		map[string]string{
			"id":   "a",
			"desc": "a",
		},

		map[string]string{
			"id":   "b",
			"desc": "bb",
		},

		map[string]string{
			"id":   "cc",
			"desc": "ccc",
		},
	}
}

func getColumns() []string {
	return []string{
		"id",
		"desc",
	}
}

//==================================================
// Tests
//==================================================
func Test_generateCellsLength(t *testing.T) {
	pvd := &Provider{}
	ret := pvd.generateCellsLength(getData(), getColumns())
	assert.Equal(t, ret["id"], 2)
	assert.Equal(t, ret["desc"], 3)
}

func Test_generateRowOutput(t *testing.T) {
	// should generate default row output
	pvd := &Provider{Delimiter: " | "}
	cellsLength := pvd.generateCellsLength(getData(), getColumns())
	ret := pvd.generateRowOutput(getData()[0], getColumns(), cellsLength)
	assert.Equal(t, ret, "Id: a  | Description: a  ")

	ret = pvd.generateRowOutput(getData()[1], getColumns(), cellsLength)
	assert.Equal(t, ret, "Id: b  | Description: bb ")

	ret = pvd.generateRowOutput(getData()[2], getColumns(), cellsLength)
	assert.Equal(t, ret, "Id: cc | Description: ccc")

	// should generate row output without columns
	pvd = &Provider{Delimiter: " || ", OmitColumns: true}
	cellsLength = pvd.generateCellsLength(getData(), getColumns())
	ret = pvd.generateRowOutput(getData()[0], getColumns(), cellsLength)
	assert.Equal(t, ret, "a  || a  ")

	ret = pvd.generateRowOutput(getData()[1], getColumns(), cellsLength)
	assert.Equal(t, ret, "b  || bb ")

	ret = pvd.generateRowOutput(getData()[2], getColumns(), cellsLength)
	assert.Equal(t, ret, "cc || ccc")

	// should generate row output with selected columns
	newColumns := []string{"id"}
	pvd = &Provider{Delimiter: " || ", Fields: newColumns}
	cellsLength = pvd.generateCellsLength(getData(), newColumns)
	ret = pvd.generateRowOutput(getData()[0], newColumns, cellsLength)
	assert.Equal(t, ret, "Id: a ")

	ret = pvd.generateRowOutput(getData()[1], newColumns, cellsLength)
	assert.Equal(t, ret, "Id: b ")

	ret = pvd.generateRowOutput(getData()[2], newColumns, cellsLength)
	assert.Equal(t, ret, "Id: cc")
}
