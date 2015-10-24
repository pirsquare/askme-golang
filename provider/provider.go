package provider

import (
	"fmt"
	"github.com/codegangsta/cli"
	"strings"
)

type Provider struct {
	Delimiter   string
	OmitColumns bool
	Fields      []string
}

func NewProviderWithCliContext(c *cli.Context) *Provider {
	// Use all columns by default
	fieldList := columnKeys
	if c.String("fields") != "" {
		fieldList = strings.Split(c.String("fields"), ",")
	}

	return &Provider{
		Delimiter:   c.String("delimiter"),
		OmitColumns: c.Bool("omit-columns"),
		Fields:      fieldList,
	}

}

func (pvd *Provider) generateCellsLength(data []map[string]string, columns []string) map[string]int {
	cellsLength := map[string]int{}

	for _, row := range data {
		for _, column := range columns {
			newLength := len(row[column])

			if existingLength, ok := cellsLength[column]; ok {
				// For given column, if current cell length is greater than existing cellsLength
				if newLength > existingLength {
					cellsLength[column] = newLength
				}
			} else {
				// No existing cellsLength for column
				cellsLength[column] = newLength
			}
		}
	}

	return cellsLength
}

func (pvd *Provider) render(data []map[string]string) {
	cellsLength := pvd.generateCellsLength(data, pvd.Fields)
	for _, row := range data {
		fmt.Println(pvd.generateRowOutput(row, pvd.Fields, cellsLength))
	}
}

func (pvd *Provider) generateRowOutput(
	row map[string]string,
	columns []string,
	cellsLength map[string]int) string {

	output := ""
	for idx, column := range columns {
		// Column output
		if !pvd.OmitColumns {
			output += columnsMapper[column] + ": "
		}

		// Field output
		// To ensure evenly spaced cells we will add empty spaces to each cells until
		// it has the same length as the biggest cell for a given column
		output += row[column] + strings.Repeat(" ", (cellsLength[column]-len(row[column])))

		// Delimiter output. don't output delimiter for last column
		if (idx + 1) != len(columns) {
			output += pvd.Delimiter
		}
	}

	return output
}
