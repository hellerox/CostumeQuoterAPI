package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Material struct {
	Idmaterial       int    `json:"idmaterial"`
	Nombrematerial   string `json:"nombrematerial"`
	Idtipomaterial   int    `json:"idtipomaterial"`
	Idmedidamaterial int    `json:"idmedidamaterial"`
	Costo            int    `json:"costo"`
}

func getMaterials(db *sql.DB) ([]Material, error) {

	defer db.Close()
	rows, err := db.Query("SELECT * FROM material") //WHERE idmaterial = $1", uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	materials := []Material{}

	for rows.Next() {
		var m Material
		if err := rows.Scan(&m.Idmaterial, &m.Nombrematerial, &m.Idtipomaterial, &m.Idmedidamaterial, &m.Costo); err != nil {
			fmt.Println("error", err)
		}
		materials = append(materials, m)

	}
	return materials, nil
}

func (m *Material) createMaterial(db *sql.DB) error {
	_, err := db.Exec("INSERT INTO material (idmaterial, nombre_material,idtipo_material,idmedida_material,costo) VALUES( $1 , $2 , $3 , $4 , $5 )", m.Idmaterial, m.Nombrematerial, m.Idtipomaterial, m.Idmedidamaterial, m.Costo)
	if err != nil {
		return err
	}
	return nil
}
