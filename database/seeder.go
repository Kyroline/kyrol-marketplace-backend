package database

import (
	model "kyrol-marketplace-backend/models"
)

func Seeder() error {

	users := []*model.User{
		&model.User{
			ID:        "US1001",
			FirstName: "Valerica",
			LastName:  "Liende",
			Username:  "Valiende",
			Email:     "valerica@liende.com",
			Password:  "password",
		},
		&model.User{
			ID:        "US1002",
			FirstName: "Valencia",
			LastName:  "Liende",
			Username:  "Valeniende",
			Email:     "valencia@liende.com",
			Password:  "password",
		},
	}

	products := []*model.Product{
		&model.Product{
			ID:          "P1001",
			Name:        "Sample Product #1",
			Description: "Sample Blah blah blah",
			Price:       300000,
			Stock:       250,
		},
		&model.Product{
			ID:          "P1002",
			Name:        "Sample Product #2",
			Description: "Sample Blah blah blah",
			Price:       310000,
			Stock:       55,
		},
		&model.Product{
			ID:          "P1003",
			Name:        "Sample Product #3",
			Description: "Sample Blah blah blah",
			Price:       100000,
			Stock:       40,
		},
		&model.Product{
			ID:          "P1004",
			Name:        "Sample Product #4",
			Description: "Sample Blah blah blah",
			Price:       300000,
			Stock:       70,
		},
		&model.Product{
			ID:          "P1005",
			Name:        "Sample Product #5",
			Description: "Sample Blah blah blah",
			Price:       400000,
			Stock:       5,
		},
	}

	roles := []*model.Role{
		&model.Role{
			ID:   1,
			Name: "Admin",
		},
		&model.Role{
			ID:   2,
			Name: "User",
		},
	}

	permissions := []*model.Permission{
		&model.Permission{
			ID:   1,
			Name: "product_create",
		},
		&model.Permission{
			ID:   2,
			Name: "product_read",
		},
		&model.Permission{
			ID:   3,
			Name: "product_update",
		},
		&model.Permission{
			ID:   4,
			Name: "product_delete",
		},
		&model.Permission{
			ID:   5,
			Name: "cart_create",
		},
		&model.Permission{
			ID:   6,
			Name: "cart_read",
		},
		&model.Permission{
			ID:   7,
			Name: "cart_update",
		},
		&model.Permission{
			ID:   8,
			Name: "cart_delete",
		},
		&model.Permission{
			ID:   9,
			Name: "category_create",
		},
		&model.Permission{
			ID:   10,
			Name: "category_read",
		},
		&model.Permission{
			ID:   11,
			Name: "category_update",
		},
		&model.Permission{
			ID:   12,
			Name: "category_delete",
		},
		&model.Permission{
			ID:   13,
			Name: "invoice_create",
		},
		&model.Permission{
			ID:   14,
			Name: "invoice_read",
		},
		&model.Permission{
			ID:   15,
			Name: "invoice_update",
		},
		&model.Permission{
			ID:   16,
			Name: "invoice_delete",
		},
	}

	categories := []*model.Category{
		&model.Category{
			ID:   1,
			Name: "Categoryy A",
		},
		&model.Category{
			ID:   2,
			Name: "Category B",
		},
	}

	DB.Create(&users)
	DB.Create(&permissions)
	DB.Create(&roles)
	DB.Create(&categories)
	DB.Create(&products)

	user := model.User{}
	role := []model.Role{}
	DB.Model(&model.User{}).Where("id = ?", "US1001").Take(&user)
	DB.Model(&model.Role{}).Where("id IN ?", [1]uint{1}).Find(&role)
	DB.Model(&user).Association("Roles").Replace(&role)

	user = model.User{}
	role = []model.Role{}
	DB.Model(&model.User{}).Where("id = ?", "US1002").Take(&user)
	DB.Model(&model.Role{}).Where("id IN ?", [1]uint{2}).Find(&role)
	DB.Model(&user).Association("Roles").Replace(&role)

	product := model.Product{}
	category := []model.Category{}
	DB.Model(&model.Product{}).Where("id = ?", "P1003").Take(&product)
	DB.Model(&[]model.Category{}).Where("id IN ?", [2]uint{1, 2}).Find(&category)
	DB.Model(&product).Association("Categories").Replace(&category)

	product = model.Product{}
	category = []model.Category{}
	DB.Model(&model.Product{}).Where("id = ?", "P1004").Take(&product)
	DB.Model(&[]model.Category{}).Where("id IN ?", [1]uint{2}).Find(&category)
	DB.Model(&product).Association("Categories").Replace(&category)

	product = model.Product{}
	category = []model.Category{}
	DB.Model(&model.Product{}).Where("id = ?", "P1005").Take(&product)
	DB.Model(&[]model.Category{}).Where("id IN ?", [1]uint{2}).Find(&category)
	DB.Model(&product).Association("Categories").Replace(&category)

	role1 := model.Role{}
	permission := []model.Permission{}
	DB.Model(&model.Role{}).Where("id = ?", 1).Take(&role1)
	DB.Model(&[]model.Permission{}).Where("id IN ?", [16]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}).Find(&permission)
	DB.Model(&role1).Association("Permissions").Replace(&permission)

	role1 = model.Role{}
	permission = []model.Permission{}
	DB.Model(&model.Role{}).Where("id = ?", 2).Take(&role1)
	DB.Model(&[]model.Permission{}).Where("id IN ?", [5]uint{2, 6, 10, 13, 14}).Find(&permission)
	DB.Model(&role1).Association("Permissions").Replace(&permission)

	return nil
}
