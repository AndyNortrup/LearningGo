// hand-receipt-reader
package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"io"
	"os"
	"regexp"
	"strings"
)

//Row where description of the property book is found
const propertyBookDescriptionRow int = 1
const propertyBookDescriptionCell int = 0

//All of the column numbers where information for a LIN can be found
const firstLINRowNumber int = 8
const colNumberLIN int = 0
const columnNumberLinSubLin int = 1
const columnNumberLinSri int = 2
const columnNumberLinSubErc int = 3
const columnNumberLinNomenclature int = 4
const columnNumberLinAuthDoc int = 9
const columnNumberLinRequired int = 12
const columnNumberLinAuthorized int = 13
const columnNumberLinDI int = 16

//All of the column numbers where information can be found
const columnNumberNSN int = 0
const columnNumberNsnUi int = 2
const columnNumberNsnUp int = 3
const columnNumberNsnNomenclature int = 4
const columnNumberNsnLLC int = 8
const columnNumberNsnECS int = 9
const columnNumberNsnSRRC int = 10
const columnNumberNsnCiic int = 12
const columnNumberNsnDLA int = 13
const columnNumberNsnUiiManaged int = 11
const columnNumberNsnPubData int = 14
const columnNumberNsnOH int = 16

//Serial numbers are listed in columns B, F, J & N
var columnNumberSerialNumbers = [...]int{1, 5, 9, 13}
var descriptions = [...]string{
	"MED COMPONENTS (SUPPLEMENTAL PROPERTY BOOK RECORD FOR EXTERNAL PERIPHERAL COMPONENT OF MEDICAL ASSEMBLIES PER DA PAM 710-2-1).",
	"BASIC LOAD (CLASS I)",
	"ORGANIZATION PROPERTY (DEPLOYABLE)",
	"INSTALLATION PROPERTY (STATION PROPERTY) (NONDEPLOYABLE).",
}

//REGEX to vet a LIN number 6, alphanumeric characters
const regexLIN string = "^\\w{6,6}$"

//REGEX to identify a NSN 13 alphanumeric characters
const regexNSN string = "^\\w{13,13}$"

var stockNumbers []StockNumber
var checkForSerial bool = false

type printable interface {
	print(w *io.Writer)
}

type PBIC struct {
	organization, uic, description string
}

func (pbic PBIC) print(w io.Writer) {
	fmt.Fprintf(w, "INSERT INTO pbic (description, ORANIZATION, uic) VALUES('%v', '%v', '%v');\n",
		pbic.description,
		pbic.organization,
		pbic.uic)
}

type LineNumber struct {
	lin, nomenclature, auth_doc  string
	authorized, required, due_in int
}

//Print a line number for import by SQLLDR
func (lineNumber LineNumber) print(w io.Writer) {

	if 50 <= len(lineNumber.nomenclature) {
		lineNumber.nomenclature = lineNumber.nomenclature[:49]
	}

	fmt.Fprintf(w, "INSERT INTO LINE_NUMBER (LIN, NOMENCLATURE, AUTH_DOC, AUTHORIZED, REQUIRED, DUE_IN, PBIC) VALUES ('%v', '%v', '%v', %v, %v, %v, (SELECT MAX(PBIC_ID) FROM PBIC));\n",
		lineNumber.lin,
		lineNumber.nomenclature,
		lineNumber.auth_doc,
		lineNumber.authorized,
		lineNumber.required,
		lineNumber.due_in)
}

type StockNumber struct {
	nsn, nomenclature, unitOfIssue string
	unitPrice                      float64
}

//Print a stock number for import by SQLLDR
func (stockNumber StockNumber) print(w io.Writer) {

	if 50 <= len(stockNumber.nomenclature) {
		stockNumber.nomenclature = stockNumber.nomenclature[:49]
	}
	fmt.Fprintf(w,
		"INSERT INTO NATIONAL_STOCK_NUMBER (NSN, NOMENCLATURE, UNIT_OF_ISSUE, UNIT_PRICE) VALUES ('%v', '%v', '%v', %v);\n",
		stockNumber.nsn,
		stockNumber.nomenclature,
		stockNumber.unitOfIssue,
		stockNumber.unitPrice)
}

type OnHandStockNumbers struct {
	onHand      int
	stockNumber string
}

func (onHand OnHandStockNumbers) print(w io.Writer) {
	fmt.Fprintf(w,
		"INSERT INTO NSN_ON_HAND (NSN, ON_HAND, LIN_ID) VALUES ('%v', %v, (SELECT MAX(LIN_ID) FROM LINE_NUMBER));\n",
		onHand.stockNumber, onHand.onHand)
}

func main() {
	inFile := openFile()
	stockNumbers = make([]StockNumber, 0)

	for index, sheet := range inFile.Sheets {
		readSheet(sheet, index)
	}

}

func openFile() xlsx.File {
	inFilePath := os.Args[1]
	inFile, err := xlsx.OpenFile(inFilePath)

	if err != nil {
		fmt.Printf("Failed to open file: %v \t %v\n", inFilePath, err)
		os.Exit(1)
	}

	return *inFile
}

func readSheet(sheet *xlsx.Sheet, index int) PBIC {

	//Get PBIC data
	sheetPBIC := newPBIC(sheet)
	sheetPBIC.description = descriptions[index]

	sheetPBIC.print(os.Stdout)

	for x := firstLINRowNumber; x < len(sheet.Rows); x++ {
		if isLineNumber(sheet.Cell(x, colNumberLIN).Value) {
			newLineNumber(sheet, x).print(os.Stdout)
		} else if isStockNumber(sheet.Cell(x, columnNumberNSN).Value) {
			onHand, stockNumber := newStockNumber(sheet, x)
			if !contiainsNSN(stockNumbers, stockNumber) {
				stockNumber.print(os.Stdout)
			}
			onHand.print(os.Stdout)
			checkForSerial = true
		} else if checkForSerial {
			readSerialNumbers(sheet, x, os.Stdout)
		}

	}

	return sheetPBIC
}

func newPBIC(sheet *xlsx.Sheet) PBIC {
	result := new(PBIC)

	organizationSplit := strings.Split(sheet.Cell(
		propertyBookDescriptionRow,
		propertyBookDescriptionCell).Value, "/")

	result.uic = organizationSplit[3]
	result.organization = organizationSplit[4]

	return *result
}

func isLineNumber(value string) bool {
	matches, err := regexp.MatchString(regexLIN, value)
	if err != nil {
		fmt.Fprintf(os.Stdout, "Failed regex: %v\n", err)
	}
	return matches
}

func newLineNumber(sheet *xlsx.Sheet, line int) LineNumber {
	result := new(LineNumber)

	if authorized, _ := sheet.Cell(line, columnNumberLinAuthorized).Int(); authorized > 0 {
		result.authorized = authorized
	}

	result.auth_doc = sheet.Cell(line, columnNumberLinAuthDoc).Value

	if due_in, _ := sheet.Cell(line, columnNumberLinDI).Int(); due_in > 0 {
		result.due_in = due_in
	}
	result.lin = sheet.Cell(line, colNumberLIN).Value
	result.nomenclature = sheet.Cell(line, columnNumberLinNomenclature).Value

	if required, _ := sheet.Cell(line, columnNumberLinRequired).Int(); required > 0 {
		result.required = required
	}

	return *result
}

func isStockNumber(value string) bool {
	matches, err := regexp.MatchString(regexNSN, value)
	if err != nil {
		fmt.Fprintf(os.Stdout, "Failed regex: %v\n", err)
	}
	return matches
}

func newStockNumber(sheet *xlsx.Sheet, line int) (OnHandStockNumbers, StockNumber) {
	stockNumber := new(StockNumber)

	stockNumber.nomenclature = sheet.Cell(line, columnNumberNsnNomenclature).Value
	stockNumber.nsn = sheet.Cell(line, columnNumberNSN).Value
	stockNumber.unitOfIssue = sheet.Cell(line, columnNumberNsnUi).Value
	stockNumber.unitPrice, _ = sheet.Cell(line, columnNumberNsnUp).Float()

	onHand := new(OnHandStockNumbers)
	onHand.onHand, _ = sheet.Cell(line, columnNumberNsnOH).Int()
	onHand.stockNumber = sheet.Cell(line, columnNumberNSN).Value

	return *onHand, *stockNumber
}

func contiainsNSN(nsns []StockNumber, newItem StockNumber) bool {
	for _, value := range nsns {
		if newItem == value {
			return true
		}
	}
	return false
}

func readSerialNumbers(sheet *xlsx.Sheet, row int, w io.Writer) {
	for _, col := range columnNumberSerialNumbers {
		if serial := sheet.Cell(row, col).Value; serial != "" {
			fmt.Fprintf(w, "INSERT INTO ITEM (SERIAL_NUMBER, NSN_ID) VALUES ('%v', (SELECT MAX(NSN_ID) FROM NSN_ON_HAND));\n", serial)
		} else {
			checkForSerial = false
			return
		}
	}
}
