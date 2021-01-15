package appsec

import (
	"context"
	"fmt"
	"net/http"

	"time"
)

// ExportConfiguration represents a collection of ExportConfiguration
//
// See: ExportConfiguration.GetExportConfiguration()
// API Docs: // appsec v1
//
// https://developer.akamai.com/api/cloud_security/application_security/v1.html

type (
	// ExportConfiguration  contains operations available on ExportConfiguration  resource
	// See: // appsec v1
	//
	// https://developer.akamai.com/api/cloud_security/application_security/v1.html#getexportconfiguration
	ExportConfiguration interface {
		GetExportConfigurations(ctx context.Context, params GetExportConfigurationsRequest) (*GetExportConfigurationsResponse, error)
	}

	GetExportConfigurationsRequest struct {
		ConfigID int `json:"configId"`
		Version  int `json:"version"`
	}

	GetExportConfigurationsResponse struct {
		ConfigID   int    `json:"configId"`
		ConfigName string `json:"configName"`
		Version    int    `json:"version"`
		BasedOn    int    `json:"basedOn"`
		Staging    struct {
			Status string `json:"status"`
		} `json:"staging"`
		Production struct {
			Status string `json:"status"`
		} `json:"production"`
		CreateDate      time.Time `json:"createDate"`
		CreatedBy       string    `json:"createdBy"`
		SelectedHosts   []string  `json:"selectedHosts"`
		SelectableHosts []string  `json:"selectableHosts"`
		RatePolicies    []struct {
			AdditionalMatchOptions []struct {
				PositiveMatch bool     `json:"positiveMatch"`
				Type          string   `json:"type"`
				Values        []string `json:"values"`
			} `json:"additionalMatchOptions"`
			AllTraffic       bool      `json:"allTraffic"`
			AverageThreshold int       `json:"averageThreshold"`
			BurstThreshold   int       `json:"burstThreshold"`
			ClientIdentifier string    `json:"clientIdentifier"`
			CreateDate       time.Time `json:"createDate"`
			Description      string    `json:"description"`
			FileExtensions   struct {
				PositiveMatch bool     `json:"positiveMatch"`
				Values        []string `json:"values"`
			} `json:"fileExtensions"`
			ID        int    `json:"id"`
			MatchType string `json:"matchType"`
			Name      string `json:"name"`
			Path      struct {
				PositiveMatch bool     `json:"positiveMatch"`
				Values        []string `json:"values"`
			} `json:"path"`
			PathMatchType        string `json:"pathMatchType"`
			PathURIPositiveMatch bool   `json:"pathUriPositiveMatch"`
			QueryParameters      []struct {
				Name          string   `json:"name"`
				PositiveMatch bool     `json:"positiveMatch"`
				ValueInRange  bool     `json:"valueInRange"`
				Values        []string `json:"values"`
			} `json:"queryParameters"`
			RequestType           string    `json:"requestType"`
			SameActionOnIpv6      bool      `json:"sameActionOnIpv6"`
			Type                  string    `json:"type"`
			UpdateDate            time.Time `json:"updateDate"`
			UseXForwardForHeaders bool      `json:"useXForwardForHeaders"`
			Used                  bool      `json:"used"`
		} `json:"ratePolicies"`
		ReputationProfiles []struct {
			Context         string `json:"context"`
			ContextReadable string `json:"contextReadable"`
			Enabled         bool   `json:"enabled"`
			ID              int    `json:"id"`
			Name            string `json:"name"`
			Threshold       int    `json:"threshold"`
		} `json:"reputationProfiles"`
		CustomRules []struct {
			Conditions []struct {
				Type          string   `json:"type"`
				PositiveMatch bool     `json:"positiveMatch"`
				Value         []string `json:"value"`
				ValueCase     bool     `json:"valueCase,omitempty"`
				ValueWildcard bool     `json:"valueWildcard,omitempty"`
			} `json:"conditions"`
			Description   string   `json:"description"`
			ID            int      `json:"id"`
			Name          string   `json:"name"`
			RuleActivated bool     `json:"ruleActivated"`
			Structured    bool     `json:"structured"`
			Tag           []string `json:"tag"`
			Version       int      `json:"version"`
		} `json:"customRules"`
		Rulesets []struct {
			ID               int       `json:"id"`
			RulesetVersionID int       `json:"rulesetVersionId"`
			Type             string    `json:"type"`
			ReleaseDate      time.Time `json:"releaseDate"`
			Rules            []struct {
				ID                  int    `json:"id"`
				InspectRequestBody  bool   `json:"inspectRequestBody"`
				InspectResponseBody bool   `json:"inspectResponseBody"`
				Outdated            bool   `json:"outdated"`
				RuleVersion         int    `json:"ruleVersion"`
				Score               int    `json:"score"`
				Tag                 string `json:"tag"`
				Title               string `json:"title"`
			} `json:"rules"`
		} `json:"rulesets"`
		MatchTargets struct {
			WebsiteTargets []struct {
				Type                      string `json:"type"`
				DefaultFile               string `json:"defaultFile"`
				EffectiveSecurityControls struct {
					ApplyApplicationLayerControls bool `json:"applyApplicationLayerControls"`
					ApplyNetworkLayerControls     bool `json:"applyNetworkLayerControls"`
					ApplyRateControls             bool `json:"applyRateControls"`
					ApplyReputationControls       bool `json:"applyReputationControls"`
					ApplySlowPostControls         bool `json:"applySlowPostControls"`
				} `json:"effectiveSecurityControls"`
				FilePaths                    []string `json:"filePaths"`
				ID                           int      `json:"id"`
				IsNegativeFileExtensionMatch bool     `json:"isNegativeFileExtensionMatch"`
				IsNegativePathMatch          bool     `json:"isNegativePathMatch"`
				SecurityPolicy               struct {
					PolicyID string `json:"policyId"`
				} `json:"securityPolicy"`
				Sequence int `json:"sequence"`
			} `json:"websiteTargets"`
		} `json:"matchTargets"`
		SecurityPolicies []struct {
			ID                      string `json:"id"`
			Name                    string `json:"name"`
			HasRatePolicyWithAPIKey bool   `json:"hasRatePolicyWithApiKey"`
			SecurityControls        struct {
				ApplyAPIConstraints           bool `json:"applyApiConstraints"`
				ApplyApplicationLayerControls bool `json:"applyApplicationLayerControls"`
				ApplyBotmanControls           bool `json:"applyBotmanControls"`
				ApplyNetworkLayerControls     bool `json:"applyNetworkLayerControls"`
				ApplyRateControls             bool `json:"applyRateControls"`
				ApplyReputationControls       bool `json:"applyReputationControls"`
				ApplySlowPostControls         bool `json:"applySlowPostControls"`
			} `json:"securityControls"`
			WebApplicationFirewall struct {
				RuleActions []struct {
					Action           string `json:"action"`
					ID               int    `json:"id"`
					RulesetVersionID int    `json:"rulesetVersionId"`
				} `json:"ruleActions"`
				AttackGroupActions []struct {
					Action           string `json:"action"`
					Group            string `json:"group"`
					RulesetVersionID int    `json:"rulesetVersionId"`
				} `json:"attackGroupActions"`
			} `json:"webApplicationFirewall"`
			CustomRuleActions []struct {
				Action string `json:"action"`
				ID     int    `json:"id"`
			} `json:"customRuleActions"`
			APIRequestConstraints struct {
				Action string `json:"action"`
			} `json:"apiRequestConstraints"`
			ClientReputation struct {
				ReputationProfileActions []struct {
					Action string `json:"action"`
					ID     int    `json:"id"`
				} `json:"reputationProfileActions"`
			} `json:"clientReputation"`
			RatePolicyActions []struct {
				ID         int    `json:"id"`
				Ipv4Action string `json:"ipv4Action"`
				Ipv6Action string `json:"ipv6Action"`
			} `json:"ratePolicyActions"`
			IPGeoFirewall struct {
				Block string `json:"block"`
			} `json:"ipGeoFirewall"`
			SlowPost struct {
				Action            string `json:"action"`
				SlowRateThreshold struct {
					Period int `json:"period"`
					Rate   int `json:"rate"`
				} `json:"slowRateThreshold"`
				DurationThreshold struct {
					Timeout int `json:"timeout"`
				} `json:"durationThreshold"`
			} `json:"slowPost"`
		} `json:"securityPolicies"`
		AdvancedOptions struct {
			Logging struct {
				AllowSampling bool `json:"allowSampling"`
				Cookies       struct {
					Type string `json:"type"`
				} `json:"cookies"`
				CustomHeaders struct {
					Type string `json:"type"`
				} `json:"customHeaders"`
				StandardHeaders struct {
					Type string `json:"type"`
				} `json:"standardHeaders"`
			} `json:"logging"`
			Prefetch struct {
				AllExtensions      bool     `json:"allExtensions"`
				EnableAppLayer     bool     `json:"enableAppLayer"`
				EnableRateControls bool     `json:"enableRateControls"`
				Extensions         []string `json:"extensions"`
			} `json:"prefetch"`
		} `json:"advancedOptions"`
	}
)

func (p *appsec) GetExportConfigurations(ctx context.Context, params GetExportConfigurationsRequest) (*GetExportConfigurationsResponse, error) {

	logger := p.Log(ctx)
	logger.Debug("GetExportConfigurations")

	var rval GetExportConfigurationsResponse

	uri := fmt.Sprintf(
		"/appsec/v1/export/configs/%d/versions/%d",
		params.ConfigID,
		params.Version)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create getexportconfigurations request: %w", err)
	}

	resp, err := p.Exec(req, &rval)
	if err != nil {
		return nil, fmt.Errorf("getexportconfigurations request failed: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, p.Error(resp)
	}

	return &rval, nil

}