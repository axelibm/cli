package ccv3

import (
	"net/url"

	"code.cloudfoundry.org/cli/api/cloudcontroller"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv3/internal"
)

// Application represents a Cloud Controller V3 Application.
type Application struct {
	Name string `json:"name"`
	GUID string `json:"guid"`
}

// GetApplications lists applications with optional filters.
func (client *Client) GetApplications(query url.Values) ([]Application, Warnings, error) {
	request, err := client.newHTTPRequest(requestOptions{
		RequestName: internal.GetAppsRequest,
		Query:       query,
	})
	if err != nil {
		return nil, nil, err
	}

	var fullAppsList []Application
	warnings, err := client.paginate(request, Application{}, func(item interface{}) error {
		if app, ok := item.(Application); ok {
			fullAppsList = append(fullAppsList, app)
		} else {
			return cloudcontroller.UnknownObjectInListError{
				Expected:   Application{},
				Unexpected: item,
			}
		}
		return nil
	})

	return fullAppsList, warnings, err
}
