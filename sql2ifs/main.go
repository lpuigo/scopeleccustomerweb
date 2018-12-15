package main

import (
	"github.com/lpuig/scopeleccustomerweb/ifsrecords"
	"log"
)

const (
	testinfile  string = `C:\Users\Laurent\Golang\src\github.com\lpuig\scopeleccustomerweb\test\sql_test.csv`
	testoutfile string = `C:\Users\Laurent\Golang\src\github.com\lpuig\scopeleccustomerweb\test\custweb_test.csv`
)

func main() {
	c := ifsrecords.NewIFSConverter(testinfile)
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
		log.Fatal("could not convert:", err)
	}
	err = c.WriteCSV()
	if err != nil {
		log.Fatal("could not write target file:", err)
	}
}
