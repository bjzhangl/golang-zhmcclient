/*
 * =============================================================================
 * IBM Confidential
 * © Copyright IBM Corp. 2020, 2021
 *
 * The source code for this program is not published or otherwise divested of
 * its trade secrets, irrespective of what has been deposited with the
 * U.S. Copyright Office.
 * =============================================================================
 */

package zhmcclient

import (
	"encoding/json"
	"net/http"
	"path"
)

// CpcAPI defines an interface for issuing CPC requests to ZHMC
//go:generate counterfeiter -o fakes/cpc.go --fake-name CpcAPI . CpcAPI
type CpcAPI interface {
	ListCPCs(query map[string]string) ([]CPC, int, *HmcError)
}

type CpcManager struct {
	client ClientAPI
}

func NewCpcManager(client ClientAPI) *CpcManager {
	return &CpcManager{
		client: client,
	}
}

/**
* GET /api/cpcs
* @query is a key, value pairs, currently, supports 'name=$name_reg_expression'
* Return: 200 and CPCs array
*     or: 400
 */
func (m *CpcManager) ListCPCs(query map[string]string) ([]CPC, int, *HmcError) {
	requestUrl := m.client.CloneEndpointURL()
	requestUrl.Path = path.Join(requestUrl.Path, "/api/cpcs")
	requestUrl = BuildUrlFromQuery(requestUrl, query)

	status, responseBody, err := m.client.ExecuteRequest(http.MethodGet, requestUrl, nil, "")
	if err != nil {
		return nil, status, err
	}

	if status == http.StatusOK {
		cpcs := &CpcsArray{}
		err := json.Unmarshal(responseBody, &cpcs)
		if err != nil {
			return nil, status, getHmcErrorFromErr(ERR_CODE_HMC_UNMARSHAL_FAIL, err)
		}
		return cpcs.CPCS, status, nil
	}

	return nil, status, GenerateErrorFromResponse(responseBody)
}
