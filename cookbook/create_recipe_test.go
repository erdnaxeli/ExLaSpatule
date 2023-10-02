package cookbook_test

import (
	"errors"
	"testing"

	"github.com/erdnaxeli/ExLaSpatule/cookbook"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateRecipe_UnknownSessionToken(t *testing.T) {
	// setup
	token := cookbook.SessionToken("abcdef")

	rr := &testRecipeRepository{}
	ur := &testUserRepository{}
	sr := &testSessionRepository{}
	sr.On("GetUserID", token).Return(cookbook.EmptyUserID, cookbook.UnknownSessionTokenError)

	cb := cookbook.NewCookbook(rr, ur, sr)

	// test
	recipeID, err := cb.CreateRecipe(cookbook.Recipe{}, cookbook.UserContext{SessionToken: token})

	// assertions
	assert.Equal(t, cookbook.EmptyRecipeID, recipeID)
	assert.ErrorIs(t, err, cookbook.UnknownSessionTokenError)

	rr.AssertExpectations(t)
	ur.AssertExpectations(t)
	sr.AssertExpectations(t)
}

func TestCreateRecipe_SessionRepositoryUnknownError(t *testing.T) {
	// setup
	token := cookbook.SessionToken("abcdef")

	rr := &testRecipeRepository{}
	ur := &testUserRepository{}
	sr := &testSessionRepository{}
	sr.On("GetUserID", token).Return(cookbook.EmptyUserID, errors.New("Unknown error"))

	cb := cookbook.NewCookbook(rr, ur, sr)

	// test
	recipeID, err := cb.CreateRecipe(cookbook.Recipe{}, cookbook.UserContext{SessionToken: token})

	// assertions
	assert.Equal(t, cookbook.EmptyRecipeID, recipeID)
	assert.ErrorIs(t, err, cookbook.UnknownError)

	rr.AssertExpectations(t)
	ur.AssertExpectations(t)
	sr.AssertExpectations(t)
}

func TestCreateRecipe_UnknownUser(t *testing.T) {
	// setup
	token := cookbook.SessionToken("abcdef")
	userID := cookbook.UserID("someid")
	groups := []cookbook.GroupID{"group1", "group2"}

	rr := &testRecipeRepository{}
	ur := &testUserRepository{}
	sr := &testSessionRepository{}
	sr.On("GetUserID", token).Return(userID, nil)
	ur.On("CanUserPublishInGroups", userID, groups).Return(cookbook.UnknownUserErr)

	cb := cookbook.NewCookbook(rr, ur, sr)

	// test
	recipeID, err := cb.CreateRecipe(cookbook.Recipe{Groups: groups}, cookbook.UserContext{SessionToken: token})

	// assertions
	assert.Equal(t, cookbook.RecipeID(""), recipeID)
	assert.ErrorIs(t, err, cookbook.UnknownUserErr)

	rr.AssertExpectations(t)
	ur.AssertExpectations(t)
	sr.AssertExpectations(t)
}

func TestCreateRecipe_UserNotInGroups(t *testing.T) {
	// setup
	token := cookbook.SessionToken("abcdef")
	userID := cookbook.UserID("someid")
	groups := []cookbook.GroupID{"group1", "group2"}
	rErr := cookbook.UserIsNotInGroupsErr{Groups: groups[:1]}

	rr := &testRecipeRepository{}
	ur := &testUserRepository{}
	sr := &testSessionRepository{}
	sr.On("GetUserID", token).Return(userID, nil)
	ur.On("CanUserPublishInGroups", userID, groups).Return(rErr)

	cb := cookbook.NewCookbook(rr, ur, sr)

	// test
	recipeID, err := cb.CreateRecipe(cookbook.Recipe{Groups: groups}, cookbook.UserContext{SessionToken: token})

	// assertions
	assert.Equal(t, cookbook.RecipeID(""), recipeID)
	var tErr cookbook.UserIsNotInGroupsErr
	require.ErrorAs(t, err, &tErr)
	assert.Equal(t, rErr.Groups, tErr.Groups)

	rr.AssertExpectations(t)
	ur.AssertExpectations(t)
	sr.AssertExpectations(t)
}

func TestCreateRecipe_UserCannotPublish(t *testing.T) {
	// setup
	token := cookbook.SessionToken("abcdef")
	userID := cookbook.UserID("someid")
	groups := []cookbook.GroupID{"group1", "group2"}
	rErr := cookbook.UserCannotPublishInGroups{Groups: groups[:1]}

	rr := &testRecipeRepository{}
	ur := &testUserRepository{}
	sr := &testSessionRepository{}
	sr.On("GetUserID", token).Return(userID, nil)
	ur.On("CanUserPublishInGroups", userID, groups).Return(rErr)

	cb := cookbook.NewCookbook(rr, ur, sr)

	// test
	recipeID, err := cb.CreateRecipe(cookbook.Recipe{Groups: groups}, cookbook.UserContext{SessionToken: token})

	// assertions
	assert.Equal(t, cookbook.RecipeID(""), recipeID)
	var tErr cookbook.UserCannotPublishInGroups
	require.ErrorAs(t, err, &tErr)
	assert.Equal(t, rErr.Groups, tErr.Groups)

	rr.AssertExpectations(t)
	ur.AssertExpectations(t)
	sr.AssertExpectations(t)
}

func TestCreateRecipe_UserRepositoryUnknownError(t *testing.T) {
	// setup
	token := cookbook.SessionToken("abcdef")
	userID := cookbook.UserID("someid")
	groups := []cookbook.GroupID{"group1", "group2"}

	rr := &testRecipeRepository{}
	ur := &testUserRepository{}
	sr := &testSessionRepository{}
	sr.On("GetUserID", token).Return(userID, nil)
	ur.On("CanUserPublishInGroups", userID, groups).Return(errors.New("Unknown error"))

	cb := cookbook.NewCookbook(rr, ur, sr)

	// test
	recipeID, err := cb.CreateRecipe(cookbook.Recipe{Groups: groups}, cookbook.UserContext{SessionToken: token})

	// assertions
	assert.Equal(t, cookbook.RecipeID(""), recipeID)
	assert.ErrorIs(t, err, cookbook.UnknownError)

	rr.AssertExpectations(t)
	ur.AssertExpectations(t)
	sr.AssertExpectations(t)
}

func TestCreateRecipe_RecipeRepositoryUnknownError(t *testing.T) {
	// setup
	token := cookbook.SessionToken("abcdef")
	userID := cookbook.UserID("someid")
	groups := []cookbook.GroupID{"group1", "group2"}
	recipe := cookbook.Recipe{Groups: groups}

	rr := &testRecipeRepository{}
	ur := &testUserRepository{}
	sr := &testSessionRepository{}
	sr.On("GetUserID", token).Return(userID, nil)
	ur.On("CanUserPublishInGroups", userID, groups).Return(nil)
	rr.On("CreateRecipe", recipe).Return(cookbook.EmptyRecipeID, errors.New("Unknown error"))

	cb := cookbook.NewCookbook(rr, ur, sr)

	// test
	recipeID, err := cb.CreateRecipe(recipe, cookbook.UserContext{SessionToken: token})

	// assertions
	assert.Equal(t, cookbook.RecipeID(""), recipeID)
	assert.ErrorIs(t, err, cookbook.UnknownError)

	rr.AssertExpectations(t)
	ur.AssertExpectations(t)
	sr.AssertExpectations(t)
}

func TestCreateRecipe_ok(t *testing.T) {
	// setup
	token := cookbook.SessionToken("abcdef")
	userID := cookbook.UserID("someid")
	recipeID := cookbook.RecipeID("someotherid")
	groups := []cookbook.GroupID{"group1", "group2"}
	recipe := cookbook.Recipe{Groups: groups}

	rr := &testRecipeRepository{}
	ur := &testUserRepository{}
	sr := &testSessionRepository{}
	sr.On("GetUserID", token).Return(userID, nil)
	ur.On("CanUserPublishInGroups", userID, groups).Return(nil)
	rr.On("CreateRecipe", recipe).Return(recipeID, nil)

	cb := cookbook.NewCookbook(rr, ur, sr)

	// test
	result, err := cb.CreateRecipe(recipe, cookbook.UserContext{SessionToken: token})

	// assertions
	assert.Equal(t, recipeID, result)
	assert.Nil(t, err)

	rr.AssertExpectations(t)
	ur.AssertExpectations(t)
	sr.AssertExpectations(t)
}
