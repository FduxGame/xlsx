package xlsx

import (
	"strconv"
	"time"
)

// StreamCell holds the data, style and type of cell for streaming.
type StreamCell struct {
	cellData  string
	cellStyle StreamStyle
	cellType  CellType
	HMerge    int
	VMerge    int
}

// Merge with other cells, horizontally and/or vertically. The merged cell should not extend past the last column of the
// streamed sheet as specified by the first row. This does not verify that merged cells do not overlap nor does it
// verify that vertical merges are within the bounds of the number of rows in the sheet.
func (c *StreamCell) Merge(hcells, vcells int) {
	c.HMerge = hcells
	c.VMerge = vcells
}

// NewStreamCell creates a new cell containing the given data with the given style and type.
func NewStreamCell(cellData string, cellStyle StreamStyle, cellType CellType) StreamCell {
	return StreamCell{
		cellData:  cellData,
		cellStyle: cellStyle,
		cellType:  cellType,
	}
}

// NewStringStreamCell creates a new cell that holds string data, is of type string and uses general formatting.
func NewStringStreamCell(cellData string) StreamCell {
	return NewStreamCell(cellData, StreamStyleDefaultString, CellTypeString)
}

// NewStyledStringStreamCell creates a new cell that holds a string and is styled according to the given style.
func NewStyledStringStreamCell(cellData string, cellStyle StreamStyle) StreamCell {
	return NewStreamCell(cellData, cellStyle, CellTypeString)
}

// NewIntegerStreamCell creates a new cell that holds an integer value (represented as string),
// is formatted as a standard integer and is of type numeric.
func NewIntegerStreamCell(cellData int) StreamCell {
	return NewStreamCell(strconv.Itoa(cellData), StreamStyleDefaultInteger, CellTypeNumeric)
}

// NewStyledIntegerStreamCell creates a new cell that holds an integer value (represented as string)
// and is styled according to the given style.
func NewStyledIntegerStreamCell(cellData int, cellStyle StreamStyle) StreamCell {
	return NewStreamCell(strconv.Itoa(cellData), cellStyle, CellTypeNumeric)
}

// NewDateStreamCell creates a new cell that holds a date value and is formatted as dd-mm-yyyy
// and is of type numeric.
func NewDateStreamCell(t time.Time) StreamCell {
	excelTime := TimeToExcelTime(t, false)
	return NewStreamCell(strconv.Itoa(int(excelTime)), StreamStyleDefaultDate, CellTypeNumeric)
}
