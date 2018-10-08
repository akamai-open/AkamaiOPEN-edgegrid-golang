package apiendpoints

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/akamai/AkamaiOPEN-edgegrid-golang/client-v1"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cast"
)

type Resources []ResourceBaseInfo

type Resource struct {
	APIResourceID           int     `json:"apiResourceId"`
	APIResourceName         string  `json:"apiResourceName"`
	ResourcePath            string  `json:"resourcePath"`
	Description             string  `json:"description"`
	LockVersion             int     `json:"lockVersion"`
	APIResourceClonedFromID *int    `json:"apiResourceClonedFromId"`
	APIResourceLogicID      int     `json:"apiResourceLogicId"`
	CreatedBy               string  `json:"createdBy"`
	CreateDate              string  `json:"createDate"`
	UpdatedBy               string  `json:"updatedBy"`
	UpdateDate              string  `json:"updateDate"`
	APIResourceMethods      Methods `json:"apiResourceMethods"`
}

type ResourceBaseInfo struct {
	APIResourceClonedFromID     *int     `json:"apiResourceClonedFromId"`
	APIResourceID               int      `json:"apiResourceId"`
	APIResourceLogicID          int      `json:"apiResourceLogicId"`
	APIResourceName             string   `json:"apiResourceName"`
	CreateDate                  string   `json:"createDate"`
	CreatedBy                   string   `json:"createdBy"`
	Description                 *string  `json:"description"`
	Link                        *string  `json:"link"`
	LockVersion                 int      `json:"lockVersion"`
	Private                     bool     `json:"private"`
	ResourcePath                string   `json:"resourcePath"`
	UpdateDate                  string   `json:"updateDate"`
	UpdatedBy                   string   `json:"updatedBy"`
	APIResourceMethods          Methods  `json:"apiResourceMethods"`
	APIResourceMethodsNameLists []string `json:"apiResourceMethodNameLists,omitempty"`
}

type ResourceSettings struct {
	Path                 string   `json:"path"`
	Methods              []string `json:"methods"`
	InheritsFromEndpoint bool     `json:"inheritsFromEndpoint"`
}

func GetResources(endpointId int, version int) (*Resources, error) {
	req, err := client.NewJSONRequest(
		Config,
		"GET",
		fmt.Sprintf(
			"/api-definitions/v2/endpoints/%d/versions/%d/resources",
			endpointId,
			version,
		),
		nil,
	)

	if err != nil {
		return nil, err
	}

	res, err := client.Do(Config, req)

	if client.IsError(res) {
		return nil, client.NewAPIError(res)
	}

	rep := &Resources{}
	if err = client.BodyJSON(res, rep); err != nil {
		return nil, err
	}

	return rep, nil
}

func (r *Resources) ToTable() *tablewriter.Table {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name", "Path", "Methods", "Private"})

	for _, resource := range *r {
		table.Append([]string{
			cast.ToString(resource.APIResourceID),
			cast.ToString(resource.APIResourceName),
			cast.ToString(resource.ResourcePath),
			cast.ToString(strings.Join(resource.APIResourceMethodsNameLists[:], ",")),
			cast.ToString(resource.Private),
		})
	}
	return table
}

func GetResource(endpointId int, resource string, version int) (*ResourceBaseInfo, error) {
	resources, err := GetResources(endpointId, version)
	if err != nil {
		return nil, err
	}

	id, err := strconv.Atoi(resource)
	if err == nil {
		for _, r := range *resources {
			if id == r.APIResourceID {
				return &r, nil
			}
		}
	} else {
		for _, r := range *resources {
			matched1, err := regexp.MatchString(resource, r.APIResourceName)
			if err != nil {
				return nil, err
			}

			matched2, err := regexp.MatchString(resource, r.ResourcePath)
			if err != nil {
				return nil, err
			}

			if matched1 || matched2 {
				return &r, nil
			}
		}
	}

	return nil, errors.New("Resource not found.")
}

func GetResourceMulti(endpointId int, resource string, version int) ([]ResourceBaseInfo, error) {
	resources, err := GetResources(endpointId, version)
	if err != nil {
		return nil, err
	}

	id, err := strconv.Atoi(resource)
	if err == nil {
		for _, r := range *resources {
			if id == r.APIResourceID {
				return []ResourceBaseInfo{r}, nil
			}
		}
	} else {
		if resource[len(resource):] == "*" {
			resource = resource[:len(resource)-1] + ".*"
		}
		ret := []ResourceBaseInfo{}
		for _, r := range *resources {
			matched1, err := regexp.MatchString(resource, r.APIResourceName)
			if err != nil {
				return nil, err
			}

			matched2, err := regexp.MatchString(resource, r.ResourcePath)
			if err != nil {
				return nil, err
			}

			if matched1 || matched2 {
				ret = append(ret, r)
			}
		}
		if len(ret) > 0 {
			return ret, nil
		}
	}

	return nil, errors.New("Resource not found.")
}
