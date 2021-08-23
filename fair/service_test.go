package fair

import (
	"log"
	"testing"

	"github.com/silvergama/unico/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

var (
	fair, fairTwo *Fair
)

type fairSuiteTest struct {
	suite.Suite
	service *FairPostgresService
}

func (s *fairSuiteTest) SetupSuite() {
	log.Print("Starting fair test suite...")

	err := repository.Setup()
	require.NoError(s.T(), err)
	s.service = NewFairPostgresService()
	fair = &Fair{
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
	fairTwo = &Fair{
		ID:         2,
		Long:       -46550161,
		Lat:        -23558731,
		Setcens:    355030885000091,
		Areap:      3550308005041,
		Coddist:    81,
		Distrito:   "VILA FORMOSA",
		Codsubpref: 21,
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
}

func (s *fairSuiteTest) TearDownSuite() {
	log.Print("Finishing fair test suite...")
}

func (s *fairSuiteTest) TearDownTest() {
	if err := s.service.Remove(fair.ID); err != nil {
		s.T().FailNow()
	}
	if err := s.service.Remove(fairTwo.ID); err != nil {
		s.T().FailNow()
	}
}

func (s *fairSuiteTest) SetupTest() {
	if _, err := s.service.Save(fair); err != nil {
		s.T().FailNow()
	}
}

func TestFairSuite(t *testing.T) {
	suite.Run(t, &fairSuiteTest{})
}

func (s *fairSuiteTest) TestSave() {
	id, err := s.service.Save(fairTwo)
	require.NoError(s.T(), err)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), id, fairTwo.ID)
	assert.NotEmpty(s.T(), id)
}

func (s *fairSuiteTest) TestUpdate() {
	fair.Bairro = "VILA PRUDENTE"
	rowsAffected, err := s.service.Update(fair)
	if err != nil {
		s.T().FailNow()
	}
	assert.Equal(s.T(), int64(1), rowsAffected)
}

func (s *fairSuiteTest) TestGet() {

	f, err := s.service.Get(fair.Bairro)
	if err != nil {
		s.T().FailNow()
	}
	assert.Equal(s.T(), 1, len(f))
}

func (s *fairSuiteTest) TestRemove() {
	if err := s.service.Remove(fairTwo.ID); err != nil {
		s.T().FailNow()
	}
}
