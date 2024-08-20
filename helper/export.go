package helper

import (
	"log"
	"net/http"

	"github.com/xuri/excelize/v2"
)

func SendExcelFile(writer http.ResponseWriter, f *excelize.File, filename string) error {
	// Set the proper headers to indicate file download
	writer.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	writer.Header().Set("Content-Disposition", "attachment;filename="+filename)

	// Write the file to the response
	err := f.Write(writer)
	if err != nil {
		log.Println("Error writing Excel file:", err)
		http.Error(writer, "Gagal menulis file", http.StatusInternalServerError)
		return err
	}

	return nil
}
