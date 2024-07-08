package modelimplementation

import (
	"context"
	"database/sql"
	"fmt"
)

type PostgresModel struct {
	BaseModel
	Table string
}

func NewPostgresModel(db *sql.DB, tableName string) *PostgresModel {
	return &PostgresModel{
		BaseModel: BaseModel{
			Client: db,
			Name:   tableName,
		},
		Table: tableName,
	}
}

func (p *PostgresModel) FindById(id string) (interface{}, error) {
	var result interface{}
	row := p.Client.(*sql.DB).QueryRowContext(context.Background(), fmt.Sprintf("SELECT * FROM %s WHERE id = $1", p.Table), id)
	err := row.Scan(&result)
	return result, err
}

func (p *PostgresModel) FindOne(criteria interface{}) (interface{}, error) {
	var result interface{}
	row := p.Client.(*sql.DB).QueryRowContext(context.Background(), fmt.Sprintf("SELECT * FROM %s WHERE %v LIMIT 1", p.Table, criteria))
	err := row.Scan(&result)
	return result, err
}

func (p *PostgresModel) FindAll(criteria interface{}) ([]interface{}, error) {
	rows, err := p.Client.(*sql.DB).QueryContext(context.Background(), fmt.Sprintf("SELECT * FROM %s WHERE %v", p.Table, criteria))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var results []interface{}
	for rows.Next() {
		var result interface{}
		err = rows.Scan(&result)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	return results, err
}

func (p *PostgresModel) UpdateOne(criteria interface{}, update interface{}) error {
	_, err := p.Client.(*sql.DB).ExecContext(context.Background(), fmt.Sprintf("UPDATE %s SET %v WHERE %v LIMIT 1", p.Table, update, criteria))
	return err
}

func (p *PostgresModel) UpdateById(id string, update interface{}) error {
	_, err := p.Client.(*sql.DB).ExecContext(context.Background(), fmt.Sprintf("UPDATE %s SET %v WHERE id = $1", p.Table, update), id)
	return err
}

func (p *PostgresModel) UpdateAll(criteria interface{}, update interface{}) error {
	_, err := p.Client.(*sql.DB).ExecContext(context.Background(), fmt.Sprintf("UPDATE %s SET %v WHERE %v", p.Table, update, criteria))
	return err
}

func (p *PostgresModel) DeleteById(id string) error {
	_, err := p.Client.(*sql.DB).ExecContext(context.Background(), fmt.Sprintf("DELETE FROM %s WHERE id = $1", p.Table), id)
	return err
}

func (p *PostgresModel) DeleteOne(criteria interface{}) error {
	_, err := p.Client.(*sql.DB).ExecContext(context.Background(), fmt.Sprintf("DELETE FROM %s WHERE %v LIMIT 1", p.Table, criteria))
	return err
}

func (p *PostgresModel) DeleteAll(criteria interface{}) error {
	_, err := p.Client.(*sql.DB).ExecContext(context.Background(), fmt.Sprintf("DELETE FROM %s WHERE %v", p.Table, criteria))
	return err
}

func (p *PostgresModel) Create(data interface{}) error {
	_, err := p.Client.(*sql.DB).ExecContext(context.Background(), fmt.Sprintf("INSERT INTO %s VALUES (%v)", p.Table, data))
	return err
}
