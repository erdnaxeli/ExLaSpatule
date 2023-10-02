package cookbook_test

import (
	"github.com/erdnaxeli/ExLaSpatule/cookbook"
	"github.com/stretchr/testify/mock"
)

type testRecipeRepository struct {
	mock.Mock
}

func (r *testRecipeRepository) CreateRecipe(recipe cookbook.Recipe) (cookbook.RecipeID, error) {
	args := r.Called(recipe)
	return args.Get(0).(cookbook.RecipeID), args.Error(1)
}

func (r *testRecipeRepository) CreateIngredient(ingredient cookbook.Ingredient) (cookbook.IngredientID, error) {
	panic("not implemented") // TODO: Implement
}

func (r *testRecipeRepository) GetRecipe(recipe cookbook.RecipeID) (cookbook.Recipe, error) {
	panic("not implemented") // TODO: Implement
}

type testUserRepository struct {
	mock.Mock
}

func (u *testUserRepository) CanUserPublishInGroups(userID cookbook.UserID, groups []cookbook.GroupID) error {
	args := u.Called(userID, groups)
	return args.Error(0)
}

type testSessionRepository struct {
	mock.Mock
}

func (s *testSessionRepository) GetUserID(token cookbook.SessionToken) (cookbook.UserID, error) {
	args := s.Called(token)
	return args.Get(0).(cookbook.UserID), args.Error(1)
}
