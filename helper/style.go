package helper

import (
	"strconv"

	"github.com/xuri/excelize/v2"
)

// AddSpaceBetweenTables menambahkan baris kosong untuk jarak antar tabel
func AddSpaceBetweenTablesProsesBisnis(f *excelize.File, sheet string, afterRow, spaceRows int) error {
	// Menambahkan baris kosong
	for i := 1; i <= spaceRows; i++ {
		row := afterRow + i
		// Menambahkan baris kosong dengan mengisi semua kolom dengan nilai kosong
		for col := 'A'; col <= 'E'; col++ {
			cell := string(col) + strconv.Itoa(row)
			f.SetCellValue(sheet, cell, "")
		}
	}
	return nil
}

// SetColumnWidths mengatur lebar kolom di sheet yang ditentukan
func SetColumnWidths(f *excelize.File, sheet string, widths map[string]float64) error {
	for col, width := range widths {
		// Set lebar untuk setiap kolom
		if err := f.SetColWidth(sheet, col, col, width); err != nil {
			return err
		}
	}
	return nil
}

func CreateHeaderStyle(f *excelize.File) (int, error) {
	style, err := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold:   false,
			Size:   12,
			Family: "Aria",
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
		Border: []excelize.Border{
			{Type: "left", Color: "#000000", Style: 1},
			{Type: "top", Color: "#000000", Style: 1},
			{Type: "right", Color: "#000000", Style: 1},
			{Type: "bottom", Color: "#000000", Style: 1},
		},
		Fill: excelize.Fill{
			Type:    "pattern",
			Pattern: 1,
			Color:   []string{"#B5CFB7"},
		},
	})
	if err != nil {
		return 0, err
	}
	return style, nil
}

func ApplyHeaderStyle(f *excelize.File, sheet string, headerCells []string, styleID int) error {
	for _, cell := range headerCells {
		if err := f.SetCellStyle(sheet, cell, cell, styleID); err != nil {
			return err
		}
	}
	return nil
}

// CreateBodyStyle membuat style untuk body yang memiliki border
func CreateBodyStyle(f *excelize.File) (int, error) {
	style, err := f.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "left", Color: "#000000", Style: 1},
			{Type: "right", Color: "#000000", Style: 1},
			{Type: "top", Color: "#000000", Style: 1},
			{Type: "bottom", Color: "#000000", Style: 1},
		},
	})
	if err != nil {
		return 0, err
	}
	return style, nil
}

// ApplyBodyStyle menerapkan style body ke sel data
func ApplyBodyStyle(f *excelize.File, sheet string, bodyCells []string, styleID int) error {
	for _, cell := range bodyCells {
		if err := f.SetCellStyle(sheet, cell, cell, styleID); err != nil {
			return err
		}
	}
	return nil
}
