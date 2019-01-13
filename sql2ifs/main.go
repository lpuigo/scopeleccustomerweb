package main

import (
	"fmt"
	"github.com/lpuig/scopeleccustomerweb/custwebrecords"
	"github.com/lpuig/scopeleccustomerweb/ifsrecords"
	"log"
)

const (
	testinfile  string = `C:\Users\Laurent\Golang\src\github.com\lpuig\scopeleccustomerweb\test\export andrei v2_20181220.csv`
	testoutfile string = `C:\Users\Laurent\Golang\src\github.com\lpuig\scopeleccustomerweb\test\custweb_test.csv`
)

func main() {
	err := ConvertToIFS(testinfile)
	if err != nil {
		log.Fatal(err)
	}
	err = ConvertToCustomerWeb(testinfile)
	if err != nil {
		log.Fatal(err)
	}
}

func ConvertToIFS(infile string) error {
	c := ifsrecords.NewIFSConverter(infile)
	f := c.AddTarget("Activities", ifsrecords.NewIFSActivitiesRecords(), c.SqlToIFSActivities)
	defer f()
	f = c.AddTarget("ActivitySLAs", ifsrecords.NewIFSActivityslasRecords(), c.SqlToIFSActivityslas)
	defer f()
	f = c.AddTarget("ActivityStatuses", ifsrecords.NewIFSActivitystatusesRecords(), c.SqlToIFSActivitystatuses)
	defer f()
	f = c.AddTarget("Locations", ifsrecords.NewIFSLocationsRecords(), c.SqlToIFSLocations)
	defer f()
	f = c.AddTarget("Ressources", ifsrecords.NewIFSRessourcesRecords(), c.SqlToIFSRessources)
	defer f()

	err := c.Convert()
	if err != nil {
		return fmt.Errorf("could not convert: %s", err.Error())
	}
	err = c.WriteCSV()
	if err != nil {
		return fmt.Errorf("could not write target file: %s", err.Error())
	}
	return nil
}

func ConvertToCustomerWeb(infile string) error {
	c := custwebrecords.NewCWConverter(infile)
	f := c.AddTarget("customerweb", custwebrecords.NewCustomerWebRecords(), c.SqltoCustomerweb)
	defer f()

	err := c.Convert()
	if err != nil {
		return fmt.Errorf("could not convert: %s", err.Error())
	}
	err = c.WriteCSV()
	if err != nil {
		return fmt.Errorf("could not write target file: %s", err.Error())
	}
	return nil
}
