package repositories

import (
	"fmt"

	"github.com/go-rest-api/db"
	"github.com/go-rest-api/models"
)

func GetProducts() ([]models.Product, error) {
	query := "select * from Products"
	rows, err := db.SQL_Query(query)
	if err != nil {
		return nil, err
	}

	var products []models.Product
	for rows.Next() {
		product := models.Product{}
		err := rows.Scan(
			&product.Id,
			&product.ProductName,
			&product.Price,
			&product.Description,
			&product.CategoryId)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	rows.Close()
	return products, nil
}

func GetProductById(id string) (*models.Product, error) {
	query := "select * from Products WHERE Id = ?"
	rows, err := db.SQL_Query(query, id)
	if err != nil {
		return nil, err
	}

	product := models.Product{}
	count := 0
	for rows.Next() {
		err := rows.Scan(
			&product.Id,
			&product.ProductName,
			&product.Price,
			&product.Description,
			&product.CategoryId)

		if err != nil {
			return nil, err
		}
		count++
	}

	rows.Close()

	if count == 0 {
		return nil, fmt.Errorf("product not found")
	}
	return &product, err
}

func GetProductByCategoryId(categoryId string) ([]models.Product, error) {
	query := "select * from Products WHERE CategoryId = ?"
	rows, err := db.SQL_Query(query, categoryId)
	if err != nil {
		return nil, err
	}

	count := 0
	var products []models.Product
	for rows.Next() {
		product := models.Product{}
		err := rows.Scan(
			&product.Id,
			&product.ProductName,
			&product.Price,
			&product.Description,
			&product.CategoryId)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
		count++
	}

	rows.Close()
	
	if count == 0 {
		return nil, fmt.Errorf("product not found")
	}
	return products, err
}
