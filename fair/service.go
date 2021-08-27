package fair

import (
	"database/sql"

	"github.com/sirupsen/logrus"
)

type UseCase interface {
	Save(f *Fair) (int, error)
	Update(f *Fair) (int64, error)
	Remove(id int) error
	Get(neighborhood string) ([]*Fair, error)
}

type FairService struct {
	db *sql.DB
}

func NewService(db *sql.DB) UseCase {
	return &FairService{
		db,
	}
}

func (fs *FairService) Save(f *Fair) (int, error) {
	stmt, err := fs.db.Prepare(`
	INSERT INTO streetfair (id, long, lat, setcens, areap, coddist, distrito, codsubpref, subprefe, regiao5, regiao8, nome_feira, registro, logradouro, numero, bairro, referencia) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17)
	`)
	if err != nil {
		logrus.Errorf("error preparing to insert street fair query: %v", err)
		return 0, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(f.ID, f.Long, f.Lat, f.Setcens, f.Areap, f.Coddist, f.Distrito, f.Codsubpref, f.Subprefe, f.Regiao5, f.Regiao8, f.NomeFeira, f.Registro, f.Logradouro, f.Numero, f.Bairro, f.Referencia)
	if err != nil {
		logrus.Errorf("error inserting street fair: %v", err)
		return f.ID, err
	}

	return f.ID, nil
}

func (fs *FairService) Update(f *Fair) (int64, error) {
	stmt, err := fs.db.Prepare(`
	UPDATE streetfair SET 
		long = $2, 
		lat = $3, 
		setcens = $4, 
		areap = $5,
		coddist = $6, 
		distrito = $7, 
		codsubpref = $8, 
		subprefe = $9, 
		regiao5 = $10, 
		regiao8 = $11, 
		nome_feira = $12, 
		registro = $13, 
		logradouro = $14, 
		numero = $15, 
		bairro = $16, 
		referencia = $17
	WHERE id = $1
	`)
	if err != nil {
		logrus.Errorf("error preparing to update street fair query: %v", err)
		return 0, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(f.ID, f.Long, f.Lat, f.Setcens, f.Areap, f.Coddist, f.Distrito, f.Codsubpref, f.Subprefe, f.Regiao5, f.Regiao8, f.NomeFeira, f.Registro, f.Logradouro, f.Numero, f.Bairro, f.Referencia)
	if err != nil {
		logrus.Errorf("error updating street fair: %v", err)
		return 0, err
	}

	return res.RowsAffected()
}

func (fs *FairService) Get(neighborhood string) ([]*Fair, error) {
	stmt, err := fs.db.Prepare(`
	SELECT 
		id, long, lat, setcens, areap, coddist, distrito, codsubpref, subprefe, regiao5, regiao8, nome_feira, registro, logradouro, numero, bairro, referencia 
	FROM streetfair WHERE bairro = $1 or bairro ilike '%$1%' OR similarity(bairro, upper(unaccent($1))) > 0.4`)
	if err != nil {
		logrus.Errorf("error preparing to get street fair query: %v", err)
		return nil, err
	}

	defer stmt.Close()

	var fairs []*Fair
	rows, err := stmt.Query(neighborhood)
	if err != nil {
		logrus.Errorf("error getting street fair: %v", err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var f Fair
		err = rows.Scan(&f.ID, &f.Long, &f.Lat, &f.Setcens, &f.Areap, &f.Coddist, &f.Distrito, &f.Codsubpref, &f.Subprefe, &f.Regiao5, &f.Regiao8, &f.NomeFeira, &f.Registro, &f.Logradouro, &f.Numero, &f.Bairro, &f.Referencia)
		if err != nil {
			logrus.Errorf("error scanning street fair: %v", err)
			return nil, err
		}
		fairs = append(fairs, &f)
	}

	return fairs, nil
}

func (fs *FairService) Remove(id int) error {
	stmt, err := fs.db.Prepare("DELETE FROM streetfair WHERE id = $1")
	if err != nil {
		logrus.Errorf("error preparing to delete street fair query: %v", err)
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(id); err != nil {
		logrus.Errorf("error deleting street fair: %v", err)
		return err
	}

	return nil
}
