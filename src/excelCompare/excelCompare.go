package main

import "fmt"
import "github.com/tealeg/xlsx"


func main () {
	fmt.Println("Demo of excel sheets reading and comparing!")

	//select * from party where party.id in (select ep.party from EXTERNAL_PARTY_ID ep where ep.end_date is not null);
	//13 records out of 869798 are marked for pending hard delete
	//var external_pids_file = "/Users/pagart/Downloads/ANZ-Projects/BDS/PRODUCTION DATA/externalPartyId.xlsx"

	//select id from party;
	var party_ids_file = "/Users/pagart/Downloads/ANZ-Projects/BDS/PRODUCTION DATA/party_id.xlsx"
	
	//select distinct(owning_party) from service_case;
	var so_ids_file = "/Users/pagart/Downloads/ANZ-Projects/BDS/PRODUCTION DATA/service_case.xlsx"

	//load party ids in a set of integers = a
	//load owner parties in a set of integers = b
	//find C = b - a = so owner parties which are not in the party ids set
	// this set C is missing parties from party db, means as many orphaned cases at least

	partyIds := loadFile(party_ids_file)
	fmt.Println("\n\n Size of array read from excel = %v \n", len(partyIds))

	ownerParties := loadFile(so_ids_file)
	fmt.Println("\n\n Size of array read from excel = %v \n", len(ownerParties))

	missingParties := Difference(ownerParties, partyIds)
	
	printParties(missingParties)
}

func printParties(data []int) {
	for _, item := range data {
		fmt.Println(item)
	}

	fmt.Printf("\n\n Total diff count = %v \n", len(data))
}

func Difference(a, b []int) (diff []int) {
	m := make(map[int]bool)

	for _, item := range b {
			m[item] = true
	}

	for _, item := range a {
			if _, ok := m[item]; !ok {
					diff = append(diff, item)
			}
	}
	return
}

func loadFile(excelFileName string) []int {
	fmt.Println("GOING TO READ FILE WITH PATH : " + excelFileName)
	var data = make([]int, 10000000)
	xlFile, err := xlsx.OpenFile(excelFileName)
    if err != nil {
		fmt.Println("Cannot read the file: " + excelFileName)
		return data
	}
	sheet := xlFile.Sheets[0]
	
	i := 0
    for _, row := range sheet.Rows {
        for _, cell := range row.Cells {
			val, errObj := cell.Int()
			if (errObj != nil) {
				//skip this row
				fmt.Printf("<error>")
			} else {
				data = append(data, val)
			}	
			i = i + 1
        }
	}
	return data
}

