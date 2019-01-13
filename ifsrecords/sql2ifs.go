package ifsrecords

import (
	"github.com/lpuig/scopeleccustomerweb/convert"
	"github.com/lpuig/scopeleccustomerweb/recordset"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
	"io"
)

type IFSConverter struct {
	convert.FileConverter
}

func NewIFSConverter(infile string) IFSConverter {
	cf := func(r io.Reader) io.Reader {
		return transform.NewReader(r, charmap.Windows1252.NewDecoder())
	}
	return IFSConverter{FileConverter: convert.NewFileConverter(infile, cf)}
}

func NewIFSActivitiesRecords() *recordset.RecordSet {
	rs := recordset.NewRecordSet()
	rs.AddHeader(recordset.Record{
		"id",
		"activity class",
		"external ref",
		"date time created",
		"base value",
		"project",
		"description",
		"duration",
		"resource 2 required (Y/N)",
		"manual split control (Y/N)",
		"report",
		"archived (Y/N)",
		"do not archive (Y/N)",
		"customer id",
		"contract id",
		"activity type id",
		"customer site id",
		"name",
		"post code zip",
		"address line 1",
		"address line 2",
		"city",
		"state",
		"latitude",
		"longitude",
		"add time",
		"qualified address",
		"qualified post code zip",
		"qualified city",
		"qualified state",
		"similarity level",
		"monday",
		"tuesday",
		"wednesday",
		"thursday",
		"friday",
		"saturday",
		"sunday",
		"region id list",
		"date time start",
		"date time end",
		"skill id list",
		"resource id list",
		"resource preference list",
		"resource type id list",
		"resource type preference list",
		"part id list",
		"usage list",
		"organisation id",
		"shift id",
	})
	return rs
}

func (c IFSConverter) SqlToIFSActivities(i int) (cwr recordset.Record, err error) {
	rec := c.Source.Get(i)
	cwr = recordset.Record{
		//"id",
		rec[c.Index["NUMINT"]],
		//"activity class",
		"CALL",
		//"external ref",
		"",
		//"date time created",
		rec[c.Index["Date de creation"]],
		//"base value",
		"",
		//"project",
		"",
		//"description",
		rec[c.Index["Activite"]] + rec[c.Index["Produit"]],
		//"duration",
		"01:30",
		//"resource 2 required (Y/N)",
		"",
		//"manual split control (Y/N)",
		"",
		//"report",
		"",
		//"archived (Y/N)",
		"",
		//"do not archive (Y/N)",
		"",
		//"customer id",
		"",
		//"contract id",
		"",
		//"activity type id",
		"",
		//"customer site id",
		"",
		//"name",
		rec[c.Index["Nom client"]],
		//"post code zip",
		rec[c.Index["Code postal site"]],
		//"address line 1",
		rec[c.Index["Adresse site"]],
		//"address line 2",
		"",
		//"city",
		rec[c.Index["Ville site"]],
		//"state",
		"France",
		//"latitude",
		"",
		//"longitude",
		"",
		//"add time",
		"",
		//"qualified address",
		"",
		//"qualified post code zip",
		"",
		//"qualified city",
		"",
		//"qualified state",
		"",
		//"similarity level",
		"",
		//"monday",
		"",
		//"tuesday",
		"",
		//"wednesday",
		"",
		//"thursday",
		"",
		//"friday",
		"",
		//"saturday",
		"",
		//"sunday",
		"",
		//"region id list",
		"",
		//"date time start",
		"",
		//"date time end",
		"",
		//"skill id list",
		"Comp_Test_IFS",
		//"resource id list",
		"",
		//"resource preference list",
		"",
		//"resource type id list",
		"",
		//"resource type preference list",
		"",
		//"part id list",
		"",
		//"usage list",
		"",
		//"organisation id",
		"",
		//"shift id",
		"",
	}
	return
}

func NewIFSActivityslasRecords() *recordset.RecordSet {
	rs := recordset.NewRecordSet()
	rs.AddHeader(recordset.Record{
		"activity id",
		"priority",
		"SLA type id",
		"start based (Y/N)",
		"date time start",
		"date time end",
	})
	return rs
}

func (c IFSConverter) SqlToIFSActivityslas(i int) (cwr recordset.Record, err error) {
	rec := c.Source.Get(i)
	cwr = recordset.Record{
		//"activity id",
		rec[c.Index["NUMINT"]],
		//"priority",
		"1",
		//"SLA type id",
		"Default",
		//"start based (Y/N)",
		"N",
		//"date time start",
		rec[c.Index["Date de réalisation"]],
		//"date time end",
		rec[c.Index["Date de fin intervention"]],
	}
	return cwr, nil
}

func NewIFSActivitystatusesRecords() *recordset.RecordSet {
	rs := recordset.NewRecordSet()
	rs.AddHeader(recordset.Record{
		"activity id",
		"date time status",
		"status id",
		"visit id",
		"date time stamp",
		"resource id",
		"resource2 id",
		"date time fixed",
		"duration",
		"user id",
		"reason",
		"commit sort value",
		"date time earliest",
		"email sent (Y/N)",
	})
	return rs
}

func (c IFSConverter) SqlToIFSActivitystatuses(i int) (cwr recordset.Record, err error) {
	rec := c.Source.Get(i)
	cwr = recordset.Record{
		//"activity id",
		rec[c.Index["NUMINT"]],
		//"date time status",
		rec[c.Index["Date Importation"]],
		//"status id",
		"0",
		//"visit id",
		"1",
		//"date time stamp",
		rec[c.Index["Date Importation"]],
		//"resource id",
		"",
		//"resource2 id",
		"",
		//"date time fixed",
		"",
		//"duration",
		"",
		//"user id",
		"",
		//"reason",
		"",
		//"commit sort value",
		"",
		//"date time earliest",
		"",
		//"email sent (Y/N)",
		"",
	}
	return cwr, nil
}

func NewIFSLocationsRecords() *recordset.RecordSet {
	rs := recordset.NewRecordSet()
	rs.AddHeader(recordset.Record{
		"id",
		"name",
		"post code zip",
		"address line 1",
		"address line 2",
		"city",
		"state",
		"latitude",
		"longitude",
		"add time",
		"qualified address",
		"qualified post code zip",
		"qualified city",
		"qualified state",
		"similarity level",
	})
	return rs
}

func (c IFSConverter) SqlToIFSLocations(i int) (cwr recordset.Record, err error) {
	rec := c.Source.Get(i)
	cwr = recordset.Record{
		//"id",
		rec[c.Index["Code Technicien"]],
		//"name",
		rec[c.Index["Nom du technicien"]],
		//"post code zip",
		"",
		//"address line 1",
		"",
		//"address line 2",
		"",
		//"city",
		"",
		//"state",
		"",
		//"latitude",
		"",
		//"longitude",
		"",
		//"add time",
		"",
		//"qualified address",
		"",
		//"qualified post code zip",
		"",
		//"qualified city",
		"",
		//"qualified state",
		"",
		//"similarity level",
		"",
	}
	return cwr, nil
}

func NewIFSRessourcesRecords() *recordset.RecordSet {
	rs := recordset.NewRecordSet()
	rs.AddHeader(recordset.Record{
		"id",
		"first name",
		"surname",
		"memo",
		"max travel",
		"travel to",
		"travel from",
		"cost per km",
		"cost per hour",
		"cost per hour overtime",
		"utilisation",
		"shift cost",
		"out of region multiplier",
		"speed factor",
		"travel with resource id",
		"resource type id",
		"organisation id",
		"start location id",
		"end location id",
		"skill id list",
		"proficiency list",
		"region id list",
		"within region multiplier list",
	})
	return rs
}

func (c IFSConverter) SqlToIFSRessources(i int) (cwr recordset.Record, err error) {
	rec := c.Source.Get(i)
	cwr = recordset.Record{
		//"id",
		rec[c.Index["Code Technicien"]],
		//"first name",
		rec[c.Index["Nom du technicien"]],
		//"surname",
		"",
		//"memo",
		"",
		//"max travel",
		"",
		//"travel to",
		"",
		//"travel from",
		"",
		//"cost per km",
		"",
		//"cost per hour",
		"",
		//"cost per hour overtime",
		"",
		//"utilisation",
		"",
		//"shift cost",
		"",
		//"out of region multiplier",
		"",
		//"speed factor",
		"",
		//"travel with resource id",
		"",
		//"resource type id",
		"Tech",
		//"organisation id",
		"",
		//"start location id",
		rec[c.Index["Code Technicien"]],
		//"end location id",
		rec[c.Index["Code Technicien"]],
		//"skill id list",
		"Comp_Test_IFS",
		//"proficiency list",
		"",
		//"region id list",
		"",
		//"within region multiplier list",
		"",
	}
	return cwr, nil
}
