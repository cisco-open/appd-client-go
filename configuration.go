package appdrest

import (
	"encoding/xml"
	"errors"
	"fmt"
	"strings"
)

var EMPTSTR = ""

type User struct {
	Username, UserId, UserDisplayName, UserRoles, UserPassword, UserEmail string
}

func (u *User) GetParams(create bool) (string, error) {
	params := fmt.Sprintf("?user-name=%s&user-display-name=%s&user-email=%s", u.Username, u.UserDisplayName, u.UserEmail)
	if create {
		if strings.Compare(u.UserRoles, EMPTSTR) != 0 {
			params += fmt.Sprintf("&user-roles=%s", u.UserRoles)
		}

		if strings.Compare(u.UserPassword, EMPTSTR) == 0 {
			// We have an error
			return params, errors.New("Error, the user password is a required parameter for a create.")
		} else {
			params += fmt.Sprintf("&user-password=%s", u.UserPassword)
		}

	} else {
		if strings.Compare(u.UserRoles, EMPTSTR) != 0 {
			params += fmt.Sprintf("&user-roles=%s", u.UserRoles)
		}

		if strings.Compare(u.UserPassword, EMPTSTR) != 0 {
			// We have an error
			params += fmt.Sprintf("&user-password=%s", u.UserPassword)
		}

		if strings.Compare(u.UserId, EMPTSTR) == 0 {
			// We have an error
			return params, errors.New("Error, the user id is a required parameter for an update.")
		} else {
			params += fmt.Sprintf("&user-id=%s", u.UserId)
		}
	}

	return params, nil
}

type CfgBusinessTransactions struct {
	XMLName             xml.Name                 `xml:"business-transactions"`
	BusinessTransaction []CfgBusinessTransaction `xml:"business-transaction"`
}

type CfgBusinessTransaction struct {
	Id int `xml:"id"`
}

// This is going to be the service for configuration
type Configuration service

/*
  -- Have not been tested yet
func (c *Configuration) PostCreateUser(u User) (int, error) {
	url := fmt.Sprintf("controller/rest/user")
	user, err := u.GetParams(true)
	if err != nil {
		return 500, err
	}
	url += user

	err = c.client.Rest("POST", url, nil, nil)
	if err != nil {
		return 500, err
	}

	return 200, nil
}

func (c *Configuration) PostUpdateUser(u User) (int, error) {
	url := fmt.Sprintf("controller/rest/user")
	user, err := u.GetParams(false)
	if err != nil {
		return 500, err
	}
	url += user

	err = c.client.Rest("POST", url, nil, nil)
	if err != nil {
		return 500, err
	}

	return 200, nil
}

func (c *Configuration) ExIncludeBTs(appId string, excl bool, b CfgBusinessTransactions) (int, error) {
	url := fmt.Sprintf("controller/rest/applications/%s/business-transactions?exclude=%t", appId, excl)

	err := c.client.Rest("POST", url, nil, b) // we need to unmarshal the object
	if err != nil {
		return 500, err
	}

	return 200, nil
}
*/
func (c *Configuration) MarkNodeHistorical(nodes string) (int, error) {
	url := fmt.Sprintf("controller/rest/mark-nodes-historical?application-component-node-ids=%s", nodes)
	fmt.Println("The url is ", url)
	err := c.client.Rest("POST", url, nil, nil)
	if err != nil {
		return 500, err
	}

	return 200, nil
}
