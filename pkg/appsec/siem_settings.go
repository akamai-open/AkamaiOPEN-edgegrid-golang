package appsec

import (
	"context"
	"fmt"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// SiemSettings represents a collection of SiemSettings
//
// See: SiemSettings.GetSiemSettings()
// API Docs: // appsec v1
//
// https://developer.akamai.com/api/cloud_security/application_security/v1.html

type (
	// SiemSettings  contains operations available on SiemSettings  resource
	// See: // appsec v1
	//
	// https://developer.akamai.com/api/cloud_security/application_security/v1.html#getsiemsettings
	SiemSettings interface {
		GetSiemSettings(ctx context.Context, params GetSiemSettingsRequest) (*GetSiemSettingsResponse, error)
		UpdateSiemSettings(ctx context.Context, params UpdateSiemSettingsRequest) (*UpdateSiemSettingsResponse, error)
		RemoveSiemSettings(ctx context.Context, params RemoveSiemSettingsRequest) (*RemoveSiemSettingsResponse, error)
	}

	GetSiemSettingsResponse struct {
		EnableForAllPolicies    bool     `json:"enableForAllPolicies"`
		EnableSiem              bool     `json:"enableSiem"`
		EnabledBotmanSiemEvents bool     `json:"enabledBotmanSiemEvents"`
		SiemDefinitionID        int      `json:"siemDefinitionId"`
		FirewallPolicyIds       []string `json:"firewallPolicyIds"`
	}

	GetSiemSettingsRequest struct {
		ConfigID int `json:"-"`
		Version  int `json:"-"`
	}

	GetSiemSettingResponse struct {
		EnableForAllPolicies    bool     `json:"enableForAllPolicies"`
		EnableSiem              bool     `json:"enableSiem"`
		EnabledBotmanSiemEvents bool     `json:"enabledBotmanSiemEvents"`
		SiemDefinitionID        int      `json:"siemDefinitionId"`
		FirewallPolicyIds       []string `json:"firewallPolicyIds"`
	}

	GetSiemSettingRequest struct {
		ConfigID int `json:"-"`
		Version  int `json:"-"`
	}

	UpdateSiemSettingsResponse struct {
		EnableForAllPolicies    bool     `json:"enableForAllPolicies"`
		EnableSiem              bool     `json:"enableSiem"`
		EnabledBotmanSiemEvents bool     `json:"enabledBotmanSiemEvents"`
		SiemDefinitionID        int      `json:"siemDefinitionId"`
		FirewallPolicyIds       []string `json:"firewallPolicyIds"`
	}

	UpdateSiemSettingsRequest struct {
		ConfigID                int      `json:"-"`
		Version                 int      `json:"-"`
		EnableForAllPolicies    bool     `json:"enableForAllPolicies"`
		EnableSiem              bool     `json:"enableSiem"`
		EnabledBotmanSiemEvents bool     `json:"enabledBotmanSiemEvents"`
		SiemDefinitionID        int      `json:"siemDefinitionId"`
		FirewallPolicyIds       []string `json:"firewallPolicyIds"`
	}

	RemoveSiemSettingsResponse struct {
		EnableForAllPolicies    bool     `json:"enableForAllPolicies"`
		EnableSiem              bool     `json:"enableSiem"`
		EnabledBotmanSiemEvents bool     `json:"enabledBotmanSiemEvents"`
		SiemDefinitionID        int      `json:"siemDefinitionId"`
		FirewallPolicyIds       []string `json:"firewallPolicyIds"`
	}

	RemoveSiemSettingsRequest struct {
		ConfigID                int      `json:"-"`
		Version                 int      `json:"-"`
		EnableForAllPolicies    bool     `json:"-"`
		EnableSiem              bool     `json:"enableSiem"`
		EnabledBotmanSiemEvents bool     `json:"-"`
		SiemDefinitionID        int      `json:"-"`
		FirewallPolicyIds       []string `json:"-"`
	}
)

// Validate validates GetSiemSettingsRequest
func (v GetSiemSettingsRequest) Validate() error {
	return validation.Errors{
		"ConfigID": validation.Validate(v.ConfigID, validation.Required),
		"Version":  validation.Validate(v.Version, validation.Required),
	}.Filter()
}

// Validate validates UpdateSiemSettingsRequest
func (v UpdateSiemSettingsRequest) Validate() error {
	return validation.Errors{
		"ConfigID": validation.Validate(v.ConfigID, validation.Required),
		"Version":  validation.Validate(v.Version, validation.Required),
	}.Filter()
}

// Validate validates UpdateSiemSettingsRequest
func (v RemoveSiemSettingsRequest) Validate() error {
	return validation.Errors{
		"ConfigID": validation.Validate(v.ConfigID, validation.Required),
		"Version":  validation.Validate(v.Version, validation.Required),
	}.Filter()
}

func (p *appsec) GetSiemSettings(ctx context.Context, params GetSiemSettingsRequest) (*GetSiemSettingsResponse, error) {
	if err := params.Validate(); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrStructValidation, err.Error())
	}

	logger := p.Log(ctx)
	logger.Debug("GetSiemSettings")

	var rval GetSiemSettingsResponse

	uri := fmt.Sprintf(
		"/appsec/v1/configs/%d/versions/%d/siem",
		params.ConfigID,
		params.Version,
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create getsiemsettings request: %w", err)
	}

	resp, err := p.Exec(req, &rval)
	if err != nil {
		return nil, fmt.Errorf("getsiemsettings  request failed: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, p.Error(resp)
	}

	return &rval, nil

}

// Update will update a SiemSettings.
//
// API Docs: // appsec v1
//
// https://developer.akamai.com/api/cloud_security/application_security/v1.html#putsiemsettings

func (p *appsec) UpdateSiemSettings(ctx context.Context, params UpdateSiemSettingsRequest) (*UpdateSiemSettingsResponse, error) {
	if err := params.Validate(); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrStructValidation, err.Error())
	}

	logger := p.Log(ctx)
	logger.Debug("UpdateSiemSettings")

	putURL := fmt.Sprintf(
		"/appsec/v1/configs/%d/versions/%d/siem",
		params.ConfigID,
		params.Version,
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, putURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create create SiemSettingsrequest: %w", err)
	}

	var rval UpdateSiemSettingsResponse
	resp, err := p.Exec(req, &rval, params)
	if err != nil {
		return nil, fmt.Errorf("create SiemSettings request failed: %w", err)
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, p.Error(resp)
	}

	return &rval, nil
}

// Remove will Remove a SiemSettings.
//
// API Docs: // appsec v1
//
// https://developer.akamai.com/api/cloud_security/application_security/v1.html#putsiemsettings

func (p *appsec) RemoveSiemSettings(ctx context.Context, params RemoveSiemSettingsRequest) (*RemoveSiemSettingsResponse, error) {
	if err := params.Validate(); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrStructValidation, err.Error())
	}

	logger := p.Log(ctx)
	logger.Debug("UpdateSiemSettings")

	putURL := fmt.Sprintf(
		"/appsec/v1/configs/%d/versions/%d/siem",
		params.ConfigID,
		params.Version,
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, putURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create Remove SiemSettingsrequest: %w", err)
	}

	var rval RemoveSiemSettingsResponse
	resp, err := p.Exec(req, &rval, params)
	if err != nil {
		return nil, fmt.Errorf("remove SiemSettings request failed: %w", err)
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, p.Error(resp)
	}

	return &rval, nil
}