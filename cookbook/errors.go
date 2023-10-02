package cookbook

import (
	"errors"
)

var UnknownSessionTokenError = errors.New("Unknown session token")
var UnknownUserErr = errors.New("Unknown user")
var UnknownError = errors.New("An unknown error happened")

type UserIsNotInGroupsErr struct {
	Groups []GroupID
}

func (UserIsNotInGroupsErr) Error() string {
	return "User is not in groups"
}

type UserCannotPublishInGroups struct {
	Groups []GroupID
}

func (UserCannotPublishInGroups) Error() string {
	return "User cannot publish in groups"
}
