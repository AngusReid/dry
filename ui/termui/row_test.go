package termui

import (
	"testing"

	gtermui "github.com/gizak/termui"
)

type testTable func() []int

func (t testTable) ColumnWidths() []int {
	return t()
}

func TestRowCreation(t *testing.T) {
	row := Row{}
	row.SetWidth(0)
	if len(row.Columns) != 0 {
		t.Errorf("Unexpected number of columns on an empty row, got %d", len(row.Columns))
	}
	if row.Width != 0 {
		t.Errorf("Unexpected width, got %d", row.Width)
	}
	row.SetWidth(10)
	if row.Width != 10 {
		t.Errorf("Unexpected width, got %d", row.Width)
	}
}

func TestSettingRowWidth_RowWithNoTable(t *testing.T) {
	row := Row{}
	c1 := &gtermui.Paragraph{}
	c2 := &gtermui.Paragraph{}
	c3 := &gtermui.Paragraph{}

	row.AddColumn(c1)
	row.AddColumn(c2)
	row.AddColumn(c3)

	rowWidth := 10
	expectedColWidth := rowWidth / len(row.Columns)
	row.SetWidth(rowWidth)
	if row.Width != rowWidth {
		t.Errorf("Unexpected width, got %d", row.Width)
	}
	if c1.Width != expectedColWidth {
		t.Errorf("Unexpected column width: got %d, expected %d", c1.Width, expectedColWidth)
	}
	if c2.Width != expectedColWidth {
		t.Errorf("Unexpected column width: got %d, expected %d", c2.Width, expectedColWidth)
	}
	if c3.Width != expectedColWidth {
		t.Errorf("Unexpected column width: got %d, expected %d", c3.Width, expectedColWidth)
	}

}

func TestSettingRowWidth_RowWithTable(t *testing.T) {
	row := Row{}
	row.Table = testTable(func() []int {
		return []int{2, 2, 2}
	})
	c1 := &gtermui.Paragraph{}
	c2 := &gtermui.Paragraph{}
	c3 := &gtermui.Paragraph{}

	row.AddColumn(c1)
	row.AddColumn(c2)
	row.AddColumn(c3)

	row.SetWidth(10)
	if row.Width != 10 {
		t.Errorf("Unexpected width, got %d", row.Width)
	}
	if c1.Width != 2 {
		t.Errorf("Unexpected column width, got %d", c1.Width)
	}
	if c2.Width != 2 {
		t.Errorf("Unexpected column width, got %d", c2.Width)
	}
	if c3.Width != 2 {
		t.Errorf("Unexpected column width, got %d", c3.Width)
	}

}
