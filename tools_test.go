package auth_test

import (
	"log"
	"testing"
)

func TestDB(t *testing.T) {
	result := db{}.ReadObject("table", map[string]string{"id": "1"})
	if len(result) != 3 && result["name"] != "pedro" {
		log.Fatal("se esperaba a pedro", result)
	}
}

func (d db) ReadObject(table_name string, where_fields map[string]string) map[string]string {
	// Variables para almacenar el resultado y el número de coincidencias encontradas
	var result map[string]string
	var count int

	// Iterar sobre los objetos en la base de datos simulada
	for _, obj := range getObjectsFromDB() {
		match := true
		// Verificar si los campos y valores del objeto coinciden con los especificados en where_fields
		for field, value := range where_fields {
			if obj[field] != value {
				match = false
				break
			}
		}

		// Si se encontró una coincidencia, almacenar el objeto y aumentar el contador
		if match {
			result = obj
			count++
		}
	}

	// Si se encontró exactamente una coincidencia, retornar el resultado
	if count == 1 {
		return result
	}

	return nil
}
