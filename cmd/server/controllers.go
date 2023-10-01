package main

import "github.com/labstack/echo/v4"

type controllers struct{}

func newControllers() controllers {
	return controllers{}
}

// (POST /ingredients)
func (c controllers) CreateIngredient(ctx echo.Context) error {
	panic("not implemented") // TODO: Implement
}

// (POST /recipes)
func (c controllers) CreateRecipes(ctx echo.Context) error {
	panic("not implemented") // TODO: Implement
}

// (GET /recipes/{id})
func (c controllers) GetRecipe(ctx echo.Context, id string) error {
	panic("not implemented") // TODO: Implement
}

// (POST /user/login)
func (c controllers) UserLogIn(ctx echo.Context) error {
	panic("not implemented") // TODO: Implement
}

// (GET /user/logout)
func (c controllers) UserLogOut(ctx echo.Context) error {
	panic("not implemented") // TODO: Implement
}
