package main

import (
	"github.com/lpuig/scopeleccustomerweb/convert"
	"github.com/lpuig/scopeleccustomerweb/custwebrecords"
	"github.com/lpuig/scopeleccustomerweb/recordset"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
	"log"
	"os"
)

const (
	testinfile  string = `C:\Users\Laurent\Golang\src\github.com\lpuig\scopeleccustomerweb\test\exportAll_test.csv`
	testoutfile string = `C:\Users\Laurent\Golang\src\github.com\lpuig\scopeleccustomerweb\test\custweb_test.csv`
)

func main() {
	infile := testinfile
	sourcers, err := NewRecordSetFromFile(infile)
	if err != nil {
		log.Fatal(err)
	}
	of, err := os.Create(testoutfile)
	if err != nil {
		log.Fatal("could not create target file:", err)
	}
	defer of.Close()
	c := convert.NewConverterToCustomerWeb(sourcers)
	c.AddTarget("customerweb", custwebrecords.NewCustomerWebRecords(), c.SqltoCustomerweb, of)
	err = c.Convert()
	if err != nil {
		log.Fatal("could not convert:", err)
	}
	err = c.WriteCSV()
	if err != nil {
		log.Fatal("could not write target file:", err)
	}
}

func NewRecordSetFromFile(file string) (rs *recordset.RecordSet, err error) {
	rs = recordset.NewRecordSet()
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	rs.AddCSVDataFrom(transform.NewReader(f, charmap.Windows1252.NewDecoder()))
	if err != nil {
		return nil, err
	}
	return
}
