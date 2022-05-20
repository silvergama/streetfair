package app

import (
	"github.com/olxbr/ligeiro/envcfg"
)

var Config = envcfg.Load(envcfg.Map{
	"APPLICATION": "street fairs",
	"API_PORT":    "9000",
	"DOCS_PATH":   "./docs",

	"DB_URL": "host=localhost port=5432 user=streetfair password=123456 dbname=streetfair sslmode=disable connect_timeout=2 statement_timeout=2s",
})
