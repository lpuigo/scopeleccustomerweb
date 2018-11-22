package main

import (
	"fmt"
	"github.com/lpuig/scopeleccustomerweb/custwebrecords"
	"github.com/lpuig/scopeleccustomerweb/recordset"
	"log"
	"os"
	"strings"
)

const (
	testinfile  string = `C:\Users\Laurent\Golang\src\github.com\lpuig\scopeleccustomerweb\test\exportAll_test.csv`
	testoutfile string = `C:\Users\Laurent\Golang\src\github.com\lpuig\scopeleccustomerweb\test\custweb_test.csv`
)

func main() {
	sourcers, err := NewRecordSetFromFile(testinfile)
	if err != nil {
		log.Fatal(err)
	}
	c := NewConverterToCustomerWeb(sourcers)
	err = c.Convert()
	if err != nil {
		log.Fatal("could not convert:", err)
	}

	of, err := os.Create(testoutfile)
	if err != nil {
		log.Fatal("could not create target file:", err)
	}
	defer of.Close()
	err = c.target.WriteCSVTo(of)
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

	rs.AddCSVDataFrom(f)
	if err != nil {
		return nil, err
	}
	return
}

type Converter struct {
	source *recordset.RecordSet
	target *recordset.RecordSet
	index  map[string]int
}

func NewConverterToCustomerWeb(sourceRs *recordset.RecordSet) Converter {
	ind := make(map[string]int)
	for _, col := range sourceRs.GetHeader().GetKeys() {
		colinds, err := sourceRs.GetRecordColNumByName(col)
		if err != nil || len(colinds) < 1 {
			panic("could not create converter index")
		}
		ind[col] = colinds[0]
	}

	return Converter{
		source: sourceRs,
		target: custwebrecords.NewCustomerWebRecords(),
		index:  ind,
	}
}

func (c Converter) Convert() error {
	for i := 0; i < c.source.Len(); i++ {
		cwr, err := c.convertRecord(i)
		if err != nil {
			return err
		}
		c.target.Add(cwr)
	}
	return nil
}

func (c Converter) convertRecord(i int) (cwr recordset.Record, err error) {
	rec := c.source.Get(i)
	cr := c.getSourceCR(i)
	woid := fmt.Sprintf("RECIFS_%05d", i)
	prjcode := "RECETTE IFS"
	addrs := rec[c.index["ADRESSE SITE"]]
	addrss := strings.Split(addrs, ",")
	num_voie := ""
	voie := addrss[0]
	if len(addrss) > 1 {
		num_voie = strings.Trim(addrss[0], " ")
		voie = strings.Trim(addrss[1], " ")
	}
	cp := rec[c.index["CODE POSTAL SITE"]]
	dep := ""
	if len(cp) > 2 {
		dep = cp[0:2]
	}

	cwr = recordset.Record{
		//"Identity.WorkOrderId",
		woid,
		//"Identity.CustomerReference",
		woid,
		//"Identity.ProjectCode",
		prjcode,
		//"Identity.Contract",
		cr["Instgpc"] + cr["Agentetl"] + cr["Cui"] + cr["Cequip"],
		//"Identity.WorkOrderState",
		rec[c.index["STATUT"]],
		//"Planning.PlannedDateOrdered",
		rec[c.index["DATE AU PLUS TOT"]],
		//"Planning.DateRdv",
		rec[c.index["DATE RDV"]],
		//"Planning.Marge",
		cr["marge"],
		//"Planning.ContractualDate",
		cr["DateContractuelle"],
		//"Templating.Criteria1",
		rec[c.index["NOM EQUIPEMENT"]][0:3],
		//"Templating.Criteria2",
		rec[c.index["NOM EQUIPEMENT"]][3:],
		//"Templating.Criteria3",
		cr["Cui"],
		//"Templating.Criteria4",
		"",
		//"Coordinates.StreetNumber",
		num_voie,
		//"Coordinates.NumberComplement",
		"",
		//"Coordinates.StreetName",
		voie,
		//"Coordinates.AddressComplement",
		"",
		//"Coordinates.PostalCode",
		cp,
		//"Coordinates.City",
		rec[c.index["VILLE SITE"]],
		//"Coordinates.Country",
		"FRANCE",
		//"Coordinates.Latitude",
		"",
		//"Coordinates.Longitude",
		"",
		//"Client.ClientFullName",
		rec[c.index["NOM CLIENT"]],
		//"Client.TelMobile",
		rec[c.index["TELEPHONE SITE"]],
		//"Client.TelFix",
		rec[c.index["NOM SITE"]],
		//"Client.MarketType",
		cr["Cegma"],
		//"Client.GTI",
		"",
		//"Client.GTR",
		cr["Engagt"],
		//"Connection.Node",
		cr["Centre"] + cr["Zone"] + cr["Catpc"],
		//"Connection.Information",
		rec[c.index["DESCRIPTION SITE"]],
		//"ParticularConditions.Criteria1",
		cr["securite"],
		//"ParticularConditions.Criteria2",
		"",
		//"ParticularConditions.Criteria3",
		"",
		//"ParticularConditions.Criteria4",
		"",
		//"ParticularConditions.Criteria5",
		"",
		//"ParticularConditions.Criteria6",
		"",
		//"ParticularConditions.Criteria7",
		"",
		//"ParticularConditions.Criteria8",
		"",
		//"ParticularConditions.Criteria9",
		"",
		//"ParticularConditions.Criteria10",
		"",
		//"Routing.Criteria1",
		"",
		//"Routing.Criteria2",
		"",
		//"Routing.Criteria3",
		"",
		//"Routing.Criteria4",
		"",
		//"Dispatching.Criteria1",
		dep,
		//"Dispatching.Criteria2",
		rec[c.index["NOM EQUIPEMENT"]],
		//"Dispatching.Criteria3",
		"",
		//"Dispatching.Criteria4",
		"",
		//"AdditionalInformation",
		rec[c.index["DESCRIPTION SITE"]],
		//"Comment",
		rec[c.index["DESCRIPTION"]],
		//"Attachment.File1",
		"",
		//"Attachment.Name1",
		"",
		//"Attachment.File2",
		"",
		//"Attachment.Name2",
		"",
		//"Attachment.File3",
		"",
		//"Attachment.Name3",
		"",
		//"Extension.Key1",
		"",
		//"Extension.Value1",
		"",
		//"Extension.Key2",
		"",
		//"Extension.Value2",
		"",
		//"Extension.Key3",
		"",
		//"Extension.Value3",
		"",
		//"réponse SPRINT",
		"",
	}
	return
}

func (c Converter) getSourceCol(i, j int) string {
	return c.source.Get(i)[j]
}

func (c Converter) getSourceCR(i int) map[string]string {
	res := make(map[string]string)
	cr := c.getSourceCol(i, c.index["COMPTE-RENDU"])
	pairs := strings.Split(cr, "#")
	for _, p := range pairs {
		vals := strings.Split(p, "=")
		val := ""
		if len(vals) > 1 {
			val = vals[1]
		}
		res[vals[0]] = val
	}
	return res
}
