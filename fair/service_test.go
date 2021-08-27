package fair

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var fair = &Fair{
	ID:         1,
	Long:       -46550164,
	Lat:        -23558733,
	Setcens:    355030885000091,
	Areap:      3550308005040,
	Coddist:    87,
	Distrito:   "VILA FORMOSA",
	Codsubpref: 26,
	Subprefe:   "ARICANDUVA-FORMOSA-CARRAO",
	Regiao5:    "Leste",
	Regiao8:    "Leste 1",
	NomeFeira:  "VILA FORMOSA",
	Registro:   "4041-0",
	Logradouro: "RUA MARAGOJIPE",
	Numero:     "S/N",
	Bairro:     "VL FORMOSA",
	Referencia: "TV RUA PRETORIA",
}

var errorDB = "an error '%s' was not expected when opening a stub database connection"

func TestSave(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.Nilf(t, err, errorDB, err)

	defer db.Close()

	mock.ExpectPrepare("INSERT INTO streetfair").ExpectExec().WithArgs(1,
		-46550164, -23558733, 355030885000091, 3550308005040, 87, "VILA FORMOSA", 26, "ARICANDUVA-FORMOSA-CARRAO", "Leste", "Leste 1", "VILA FORMOSA", "4041-0", "RUA MARAGOJIPE", "S/N", "VL FORMOSA", "TV RUA PRETORIA").WillReturnResult(sqlmock.NewResult(1, 1))

	srv := NewService(db)
	ID, err := srv.Save(fair)
	assert.Nil(t, err)
	assert.Equal(t, 1, ID)
}

func TestUpdate(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.Nilf(t, err, errorDB, err)

	defer db.Close()

	mock.ExpectPrepare("UPDATE streetfair").ExpectExec().WithArgs(1,
		-46550164, -23558733, 355030885000091, 3550308005040, 87, "VILA FORMOSA", 26, "ARICANDUVA-FORMOSA-CARRAO", "Leste", "Leste 1", "VILA FORMOSA", "4041-0", "RUA MARAGOJIPE", "S/N", "VL FORMOSA", "TV RUA PRETORIA").WillReturnResult(sqlmock.NewResult(1, 1))

	srv := NewService(db)
	rowsAffected, err := srv.Update(fair)
	assert.Nil(t, err)
	assert.Equal(t, int64(1), rowsAffected)
}

func TestGet(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.Nilf(t, err, errorDB, err)

	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "long", "lat", "setcens", "areap", "coddist", "distrito", "codsubpref", "subprefe", "regiao5", "regiao8", "nome_feira", "registro", "logradouro", "numero", "bairro", "referencia"}).AddRow(1,
		-46550164, -23558733, 355030885000091, 3550308005040, 87, "VILA FORMOSA", 26, "ARICANDUVA-FORMOSA-CARRAO", "Leste", "Leste 1", "VILA FORMOSA", "4041-0", "RUA MARAGOJIPE", "S/N", "VL FORMOSA", "TV RUA PRETORIA")
	mock.ExpectPrepare("SELECT (.+) FROM streetfair").ExpectQuery().WithArgs("VL FORMOSA").WillReturnRows(rows)

	srv := NewService(db)
	fairs, err := srv.Get("VL FORMOSA")
	assert.Nil(t, err)
	assert.Equal(t, []*Fair{fair}, fairs)
}

func TestDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.Nilf(t, err, errorDB, err)

	defer db.Close()

	mock.ExpectPrepare("DELETE FROM streetfair").ExpectExec().WithArgs(1).WillReturnResult(sqlmock.NewResult(0, 0))

	srv := NewService(db)
	err = srv.Remove(1)
	assert.Nil(t, err)
}
