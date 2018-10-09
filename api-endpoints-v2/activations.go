package apiendpoints

import (
	"fmt"

	"github.com/akamai/AkamaiOPEN-edgegrid-golang/client-v1"
)

type Activation struct {
	Networks               []string `json:"networks"`
	NotificationRecipients []string `json:"notificationRecipients"`
	Notes                  string   `json:"notes"`
}

func ActivateEndpoint(endpointId int, version int, activation *Activation) (*Activation, error) {
	req, err := client.NewJSONRequest(
		Config,
		"POST",
		fmt.Sprintf(
			"/api-definitions/v2/endpoints/%d/versions/%d/activate",
			endpointId,
			version,
		),
		activation,
	)

	if err != nil {
		return nil, err
	}

	res, err := client.Do(Config, req)

	if client.IsError(res) {
		return nil, client.NewAPIError(res)
	}

	return activation, nil
}

func IsActive(endpoint *Endpoint, network string) bool {
	if network == "production" {
		if endpoint.ProductionStatus == "PENDING" || endpoint.ProductionStatus == "ACTIVE" {
			return true
		}
	}

	if network == "staging" {
		if endpoint.StagingStatus == "PENDING" || endpoint.StagingStatus == "ACTIVE" {
			return true
		}
	}

	return false
}
