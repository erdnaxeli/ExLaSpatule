package cookbook

type GroupID string
type IngredientID string
type RecipeID string
type SessionToken string
type UserID string

const (
	EmptyGroupID      = GroupID("")
	EmptyIngredientID = IngredientID("")
	EmptyRecipeID     = RecipeID("")
	EmptySessionToken = SessionToken("")
	EmptyUserID       = UserID("")
)

type Ingredient struct {
	Name string
}

type RecipeIngredient struct {
	Ingredient Ingredient
	Quantity   float32
	Unit       string
}

type RecipeStep struct {
	Description string
}

type Recipe struct {
	Name        string
	Ingredients []RecipeIngredient
	Steps       []RecipeStep
	Groups      []GroupID
}

type UserContext struct {
	SessionToken SessionToken
}

type RecipeRepository interface {
	CreateRecipe(recipe Recipe) (RecipeID, error)
	CreateIngredient(ingredient Ingredient) (IngredientID, error)
	GetRecipe(recipe RecipeID) (Recipe, error)
}

type SessionRepository interface {
	GetUserID(token SessionToken) (UserID, error)
}

type UserRepository interface {
	// Return an error if the user cannot publish, nil if they can.
	CanUserPublishInGroups(userID UserID, groups []GroupID) error
}

type cookbook struct {
	recipesRepository RecipeRepository
	userRepository    UserRepository
	sessionRepository SessionRepository
}

func NewCookbook(
	rp RecipeRepository, ur UserRepository, sr SessionRepository,
) cookbook {
	return cookbook{rp, ur, sr}
}
