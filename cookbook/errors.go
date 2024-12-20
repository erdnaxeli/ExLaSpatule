package cookbook

import (
	"errors"
)

var (
	ErrUnknownSessionToken = errors.New("Unknown session token")
	ErrUnknownUser         = errors.New("Unknown user")
	ErrUnknown             = errors.New("An unknown error happened")
)

type UserIsNotInGroupsError struct {
	Groups []GroupID
}

func (UserIsNotInGroupsError) Error() string {
	return "User is not in groups"
}

type UserCannotPublishInGroupsError struct {
	Groups []GroupID
}

func (UserCannotPublishInGroupsError) Error() string {
	return "User cannot publish in groups"
}
