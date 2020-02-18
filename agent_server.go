package thousandeyes

import (
	"errors"
	"fmt"
)

type AgentServer struct {
	Agents                []Agent        `json:"agents,omitempty"`
	AlertsEnabled         int            `json:"alertsEnabled,omitempty"`
	AlertRules            []AlertRule    `json:"alertRules,omitempty"`
	ApiLinks              []ApiLink      `json:"apiLinks,omitempty"`
	CreatedBy             string         `json:"createdBy,omitempty"`
	CreatedDate           string         `json:"createdDate,omitempty"`
	Description           string         `json:"description,omitempty"`
	Enabled               int            `json:"enabled,omitempty"`
	Groups                []GroupLabel   `json:"groups,omitempty"`
	LiveShare             int            `json:"liveShare,omitempty"`
	ModifiedBy            string         `json:"modifiedBy,omitempty"`
	ModifiedDate          string         `json:"modifiedDate,omitempty"`
	SavedEvent            int            `json:"savedEvent,omitempty"`
	SharedWithAccounts    []AccountGroup `json:"sharedWithAccounts,omitempty"`
	TestId                int            `json:"testId,omitempty"`
	TestName              string         `json:"testName,omitempty"`
	Type                  string         `json:"type,omitempty"`
	BandwidthMeasurements int            `json:"bandwidthMeasurements,omitempty"`
	BgpMeasurements       int            `json:"bgpMeasurements,omitempty"`
	BgpMonitors           []Monitor      `json:"bgpMonitors,omitempty"`
	Interval              int            `json:"interval,omitempty"`
	MtuMeasurements       int            `json:"mtuMeasurements,omitempty"`
	NumPathTraces         int            `json:"numPathTraces,omitempty"`
	Port                  int            `json:"port,omitempty"`
	ProbeMode             string         `json:"probeMode,omitempty"`
	Protocol              string         `json:"protocol,omitempty"`
	Server                string         `json:"server,omitempty"`
}

func (t *AgentServer) AddAgent(id int) {
	agent := Agent{AgentId: id}
	t.Agents = append(t.Agents, agent)
}

func (c *Client) GetAgentServer(id int) (*AgentServer, error) {
	resp, err := c.get(fmt.Sprintf("/tests/%d", id))
	if err != nil {
		return &AgentServer{}, err
	}
	var target map[string][]AgentServer
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}

func (c Client) CreateAgentServer(t AgentServer) (*AgentServer, error) {
	resp, err := c.post("/tests/agent-to-server/new", t, nil)
	if err != nil {
		return &t, err
	}
	if resp.StatusCode != 201 {
		return &t, errors.New(fmt.Sprintf("failed to create agent server, response code %d", resp.StatusCode))
	}
	var target map[string][]AgentServer
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}

func (c *Client) DeleteAgentServer(id int) error {
	resp, err := c.post(fmt.Sprintf("/tests/agent-to-server/%d/delete", id), nil, nil)
	if err != nil {
		return err
	}
	if resp.StatusCode != 204 {
		return errors.New(fmt.Sprintf("failed to delete agent server, response code %d", resp.StatusCode))
	}
	return nil
}

func (c *Client) UpdateAgentServer(id int, t AgentServer) (*AgentServer, error) {
	resp, err := c.post(fmt.Sprintf("/tests/agent-to-server/%d/update", id), t, nil)
	if err != nil {
		return &t, err
	}
	if resp.StatusCode != 200 {
		return &t, errors.New(fmt.Sprintf("failed to update agent server, response code %d", resp.StatusCode))
	}
	var target map[string][]AgentServer
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}