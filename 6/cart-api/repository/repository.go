package repository

import (
	"errors"
	"richard-here/haioo-api/cart-api/database"
	"richard-here/haioo-api/cart-api/model"

	"github.com/google/uuid"
)

type Repo struct {
	Db database.DBInstance
}

func CreateRepository(db database.DBInstance) Repo {
	return Repo{Db: db}
}

func (r *Repo) AddProductToCart(p *model.Product) error {
	err := r.Db.Db.Create(&p).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) UpdateProductInCart(p *model.Product) error {
	err := r.Db.Db.Save(&p).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) GetProductsInCart() ([]model.Product, error) {
	var ps []model.Product
	err := r.Db.Db.Find(&ps).Error
	if err != nil {
		return nil, err
	}
	return ps, err
}

func (r *Repo) DeleteProductFromCart(code string) error {
	var p model.Product
	err := r.Db.Db.Find(&p, "code = ?", code).Error
	if err != nil {
		return err
	}

	if p.Code == uuid.Nil {
		return errors.New("Product with given code not found")
	}

	err = r.Db.Db.Delete(&p).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) VerifyProductExists(name string) (*model.Product, error) {
	var p *model.Product
	err := r.Db.Db.Find(&p, "name = ?", name).Error
	if err != nil {
		return nil, err
	}
	if p.Code == uuid.Nil {
		return nil, nil
	}
	return p, nil
}
