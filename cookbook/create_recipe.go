package cookbook

import "errors"

func (c cookbook) CreateRecipe(
	recipe Recipe, userContext UserContext,
) (RecipeID, error) {
	userID, err := c.sessionRepository.GetUserID(userContext.SessionToken)
	if err != nil {
		if errors.Is(err, UnknownSessionTokenError) {
			return EmptyRecipeID, UnknownSessionTokenError
		}

		return EmptyRecipeID, UnknownError
	}

	err = c.userRepository.CanUserPublishInGroups(userID, recipe.Groups)
	if err != nil {
		if errors.Is(err, UnknownUserErr) {
			return EmptyRecipeID, UnknownUserErr
		}

		var userIsNotInGroupErr UserIsNotInGroupsErr
		if errors.As(err, &userIsNotInGroupErr) {
			return EmptyRecipeID, userIsNotInGroupErr
		}

		var userCannotPublishErr UserCannotPublishInGroups
		if errors.As(err, &userCannotPublishErr) {
			return EmptyRecipeID, userCannotPublishErr
		}

		return EmptyRecipeID, UnknownError
	}

	recipeID, err := c.recipesRepository.CreateRecipe(recipe)
	if err != nil {
		return EmptyRecipeID, UnknownError
	}

	return recipeID, nil
}
