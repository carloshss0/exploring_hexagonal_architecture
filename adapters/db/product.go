package db

import (
	"database/sql"

	"github.com/carloshss0/exploring_hexagonal_architecture/application"
	_ "github.com/mattn/go-sqlite3"
)

type ProductDb struct {
	db *sql.DB
}

func NewProductDb(db *sql.DB) *ProductDb {
	return &ProductDb{db: db}
}

func (p *ProductDb) Get(id string) (application.ProductInterface, error) {
	var product application.Product
	stmt, err := p.db.Prepare("select id, name, price, status from products where id=?")
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price, &product.Status)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *ProductDb) create(product application.ProductInterface) (application.ProductInterface, error) {

	stmt, err := p.db.Prepare(`insert into products(id, name, price, status) values(?,?,?,?)`)
	if err != nil {
		return nil, err
	}
	_, err = stmt.Exec(
		product.GetID(),
		product.GetName(),
		product.GetPrice(),
		product.GetStatus(),
	)

	if err != nil {
		return nil, err
	}

	err = stmt.Close()

	if err != nil {
		return nil, err
	}
	return product, nil
}


func (p *ProductDb) update(product application.ProductInterface) (application.ProductInterface, error) {

	_, err := p.db.Exec("UPDATE products SET name = ?, price= ?, status= ? WHERE id = ?",
		product.GetName(),
		product.GetPrice(),
		product.GetStatus(),
		product.GetID(),
	)

	if err != nil {
		return nil, err
	}

	return product, nil


}

func (p *ProductDb) Save(product application.ProductInterface) (application.ProductInterface, error) {
	var count int
	err := p.db.QueryRow("Select COUNT(*) from products where id=?", product.GetID()).Scan(&count)

	if err != nil {
		return nil, err
	}

	if count == 0 {
		_, err := p.create(product)
		if err != nil {
			return nil, err
		} 
	} else {
		_, err := p.update(product)
		if err != nil {
			return nil, err
		}
	}
	return product, nil
}


