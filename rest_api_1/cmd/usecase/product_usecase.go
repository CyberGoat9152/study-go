package usecase

import (
	"restapi/1/cmd/model"
	"restapi/1/cmd/repository"
)

type ProductUsecase struct {
	repository repository.ProductRepository
}

func NewProductUserCase(repo repository.ProductRepository) ProductUsecase {
	return ProductUsecase{
		repository: repo,
	}
}

func (pu *ProductUsecase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}

func (pu *ProductUsecase) GetProductById(product_id int) (*model.Product, error) {

	product, err := pu.repository.GetProductById(product_id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (pu *ProductUsecase) CreateProduct(product model.Product) (model.Product, error) {
	productId, err := pu.repository.CreateProduct(product)
	if err != nil {
		return model.Product{}, err
	}

	product.ID = productId

	return product, nil
}

func (pu *ProductUsecase) UpdateProduct(product_id int, product model.Product) error {
	err := pu.repository.UpdateProduct(product_id, product)

	if err != nil {
		return err
	}

	return nil
}

func (pu *ProductUsecase) DeleteProduct(product_id int) error {
	err := pu.repository.DeleteProduct(product_id)

	if err != nil {
		return err
	}

	return nil
}
