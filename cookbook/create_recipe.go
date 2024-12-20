package cookbook

import "errors"

func (c cookbook) CreateRecipe(
	recipe Recipe, userContext UserContext,
) (RecipeID, error) {
	userID, err := c.sessionRepository.GetUserID(userContext.SessionToken)
	if err != nil {
		if errors.Is(err, ErrUnknownSessionToken) {
			return EmptyRecipeID, ErrUnknownSessionToken
		}

		return EmptyRecipeID, ErrUnknown
	}

	err = c.userRepository.CanUserPublishInGroups(userID, recipe.Groups)
	if err != nil {
		if errors.Is(err, ErrUnknownUser) {
			return EmptyRecipeID, ErrUnknownUser
		}

		var userIsNotInGroupErr UserIsNotInGroupsError
		if errors.As(err, &userIsNotInGroupErr) {
			return EmptyRecipeID, userIsNotInGroupErr
		}

		var userCannotPublishErr UserCannotPublishInGroupsError
		if errors.As(err, &userCannotPublishErr) {
			return EmptyRecipeID, userCannotPublishErr
		}

		return EmptyRecipeID, ErrUnknown
	}

	recipeID, err := c.recipesRepository.CreateRecipe(recipe)
	if err != nil {
		return EmptyRecipeID, ErrUnknown
	}

	return recipeID, nil
}
