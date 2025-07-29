package repositories

import (
	"database/sql"
	"go-api/models"
)

// Retorna todas as pessoas
func GetAll(db *sql.DB) ([]models.Pessoa, error) {
	rows, err := db.Query("SELECT Id, Nome, Idade FROM Pessoas")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pessoas []models.Pessoa
	for rows.Next() {
		var p models.Pessoa
		rows.Scan(&p.ID, &p.Nome, &p.Idade)
		pessoas = append(pessoas, p)
	}
	return pessoas, nil
}

// Retorna pessoa por ID
func GetByID(db *sql.DB, id int) (models.Pessoa, error) {
	var p models.Pessoa
	err := db.QueryRow("SELECT Id, Nome, Idade FROM Pessoas WHERE Id = @p1", id).Scan(&p.ID, &p.Nome, &p.Idade)
	return p, err
}

// Cria nova pessoa e retorna com ID
func Create(db *sql.DB, p *models.Pessoa) error {
	return db.QueryRow("INSERT INTO Pessoas (Nome, Idade) OUTPUT INSERTED.Id VALUES (@p1, @p2)", p.Nome, p.Idade).Scan(&p.ID)
}

// Atualiza pessoa existente
func Update(db *sql.DB, p models.Pessoa) error {
	_, err := db.Exec("UPDATE Pessoas SET Nome=@p1, Idade=@p2 WHERE Id=@p3", p.Nome, p.Idade, p.ID)
	return err
}

// Exclui pessoa por ID
func Delete(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM Pessoas WHERE Id=@p1", id)
	return err
}
