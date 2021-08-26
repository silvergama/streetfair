package importdb

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/silvergama/unico/fair"
	"github.com/silvergama/unico/repository"
	"github.com/sirupsen/logrus"
)

const MinLenLine = 17

type UseCase interface {
	InportFromCSV(path string) error
}

type importCSV struct {
	db *sql.DB
}

func NewImportCSV() *importCSV {
	return &importCSV{
		repository.GetInstance(),
	}
}

func (is *importCSV) ImportFromCSV(path string) error {
	csvFile, err := os.Open(path)
	if err != nil {
		logrus.Errorln(err)
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	csvr := csv.NewReader(csvFile)
	csvr.FieldsPerRecord = -1
	csvLines, err := csvr.ReadAll()
	if err != nil {
		logrus.WithError(err)
	}

	totalImported := importToDatabase(csvLines, fair.NewService(repository.GetInstance()))
	fmt.Println("Import completed successfully!")
	fmt.Printf("Total of %d imported lines", totalImported)
	logrus.Infof("Total of %d imported lines", totalImported)
	return nil
}

func importToDatabase(CSVLines [][]string, f fair.UseCase) int {
	var count int
	for i, line := range CSVLines {
		if i == 0 {
			continue
		}
		if len(line) < MinLenLine {
			logrus.Infof("the line %d not contain the number of fields required", i)
			continue
		}
		emp := newFair(line)
		id, err := f.Save(emp)
		if err != nil {
			logrus.Warnf("error: %s with ID: %d", err, id)
			continue
		}
		count++
		logrus.Infof("imported line with ID: %d", id)
	}
	return count
}

func newFair(line []string) *fair.Fair {
	id, _ := strconv.Atoi(line[0])
	long, _ := strconv.Atoi(line[1])
	lat, _ := strconv.Atoi(line[2])
	setcens, _ := strconv.Atoi(line[3])
	areap, _ := strconv.Atoi(line[4])
	coddist, _ := strconv.Atoi(line[5])
	distrito := line[6]
	codsubpref, _ := strconv.Atoi(line[7])
	subprefe := line[8]
	regiao5 := line[9]
	regiao8 := line[10]
	nomeFeira := line[11]
	registro := line[12]
	logradouro := line[13]
	numero := line[14]
	bairro := line[15]
	referencia := line[16]

	return &fair.Fair{
		ID:         id,
		Long:       long,
		Lat:        lat,
		Setcens:    setcens,
		Areap:      areap,
		Coddist:    coddist,
		Distrito:   distrito,
		Codsubpref: codsubpref,
		Subprefe:   subprefe,
		Regiao5:    regiao5,
		Regiao8:    regiao8,
		NomeFeira:  nomeFeira,
		Registro:   registro,
		Logradouro: logradouro,
		Numero:     numero,
		Bairro:     bairro,
		Referencia: referencia,
	}
}
