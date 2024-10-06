package repository

import (
	"database/sql"
	"fmt"
	"restapi/1/cmd/model"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	query := "SELECT id, product_name, price FROM product"
	rows, err := pr.connection.Query(query)

	if err != nil {
		fmt.Println(err)
		return []model.Product{}, err
	}

	var productList []model.Product
	var productObj model.Product

	for rows.Next() {
		err := rows.Scan(
			&productObj.ID,
			&productObj.Name,
			&productObj.Price,
		)

		if err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}

		productList = append(productList, productObj)
	}

	rows.Close()

	return productList, nil
}

func (pr *ProductRepository) CreateProduct(product model.Product) (int, error) {
	var id int
	query, err := pr.connection.Prepare("INSERT INTO product (product_name, price)" +
		" VALUES ($1, $2) RETURNING id;")

	if err != nil {
		fmt.Println(err)
		return -1, err
	}

	err = query.QueryRow(product.Name, product.Price).Scan(&id)

	if err != nil {
		fmt.Println(err)
		return -1, err
	}

	query.Close()
	return id, nil
}

func (pr *ProductRepository) GetProductById(product_id int) (*model.Product, error) {
	var product model.Product
	query, err := pr.connection.Prepare("SELECT id, product_name, price FROM product WHERE id=$1;")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	err = query.QueryRow(product_id).Scan(
		&product.ID,
		&product.Name,
		&product.Price,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		fmt.Println(err)
		return nil, err
	}

	query.Close()
	return &product, nil
}

func (pr *ProductRepository) UpdateProduct(product_id int, new_product model.Product) error {
	query, err := pr.connection.Prepare("UPDATE product SET product_name=$1, price=$2 WHERE id=$3;")

	if err != nil {
		fmt.Println(err)
		return err
	}

	err = query.QueryRow(new_product.Name, new_product.Price, product_id).Scan()

	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		fmt.Println(err)
		return err
	}

	query.Close()
	return nil
}

func (pr *ProductRepository) DeleteProduct(product_id int) error {
	query, err := pr.connection.Prepare("DELETE FROM product WHERE id=$1;")

	if err != nil {
		fmt.Println(err)
		return err
	}

	err = query.QueryRow(product_id).Scan()

	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		fmt.Println(err)
		return err
	}

	query.Close()
	return nil
}
