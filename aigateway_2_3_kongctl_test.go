package sdkkonnectgo_test

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"

	sdk "github.com/Kong/sdk-konnect-go"
	c "github.com/Kong/sdk-konnect-go/models/components"
	"github.com/stretchr/testify/require"
)

type costTagsCapture struct {
	method, path string
	body         []byte
}

func (c *costTagsCapture) Do(req *http.Request) (*http.Response, error) {
	c.method, c.path = req.Method, req.URL.Path
	if req.Body != nil {
		var err error
		c.body, err = io.ReadAll(req.Body)
		if err != nil {
			return nil, err
		}
	}
	return &http.Response{
		StatusCode: http.StatusOK,
		Header:     http.Header{"Content-Type": {"application/json"}},
		Body:       io.NopCloser(strings.NewReader(`{"headers":[],"prefixes":[]}`)),
		Request:    req,
	}, nil
}

func TestAIGateway23CostTagCapture(t *testing.T) {
	capture := &costTagsCapture{}
	client := sdk.New(sdk.WithClient(capture), sdk.WithServerURL("https://probe.invalid"))
	_, err := client.AIGatewayCostManagementTags.GetCostManagementTags(t.Context())
	require.NoError(t, err)
	require.Equal(t, http.MethodGet, capture.method)
	require.Equal(t, "/v1/ai-cost-management/tag-capture", capture.path)

	var request c.UpdateCostManagementTags
	input := `{"headers":[{"name":"x-tenant-id","tag_key":"tenant"}],"prefixes":[{"name":"my-company"}]}`
	require.NoError(t, json.Unmarshal([]byte(input), &request))
	_, err = client.AIGatewayCostManagementTags.UpdateCostManagementTags(t.Context(), request)
	require.NoError(t, err)
	require.Equal(t, http.MethodPut, capture.method)
	require.Equal(t, "/v1/ai-cost-management/tag-capture", capture.path)
	require.JSONEq(t, input, string(capture.body))
}

// Exercise request decoding and encoding: fields absent from the generated SDK
// would otherwise be silently dropped when declarative input is decoded.
func TestAIGateway23RequestFields(t *testing.T) {
	for _, tc := range []struct {
		name    string
		request any
		input   string
		want    string
	}{
		{
			name:    "baseten provider",
			request: &c.CreateAIGatewayModelProviderRequest{},
			input:   `{"type":"baseten","name":"baseten","display_name":"Baseten","config":{"auth":{"type":"basic"}}}`,
			want:    `{"type":"baseten","name":"baseten"}`,
		},
		{
			name:    "bedrock options",
			request: &c.AIGatewayTargetBedrockConfig{},
			input:   `{"type":"bedrock","endpoint_type":"mantle","vpc_endpoint":"https://bedrock.example"}`,
			want:    `{"endpoint_type":"mantle","vpc_endpoint":"https://bedrock.example"}`,
		},
		{
			name:    "failover criteria",
			request: &c.AIGatewayModelBalancerRoundRobinConfig{},
			input:   `{"algorithm":"round-robin","failover_criteria":["http_400","http_529","http_5xx"]}`,
			want:    `{"failover_criteria":["http_400","http_529","http_5xx"]}`,
		},
		{
			name:    "token vault datastore",
			request: &c.CreateAIGatewayDatastoreRequest{},
			input:   `{"type":"token-vault","name":"tokens","display_name":"Tokens","config":{"directory":"directory","current_version":"v1","keys":[{"version":"v1","key":"{vault://hcv/encryption-key}"}]}}`,
			want:    `{"type":"token-vault","config":{"directory":"directory","current_version":"v1","keys":[{"version":"v1","key":"{vault://hcv/encryption-key}"}]}}`,
		},
		{
			name:    "oidc credential claims and token vault",
			request: &c.CreateAIGatewayAuthStrategyRequest{},
			input:   `{"type":"openid-connect","name":"oidc","display_name":"OIDC","config":{"issuer":"https://issuer.example","cache_tokens_salt":"salt","credential_claims":[["account","id"],["sub"]]},"datastores":[{"name":"tokens"}],"token_vault":{"enable_credential_manager":true}}`,
			want:    `{"config":{"credential_claims":[["account","id"],["sub"]]},"datastores":[{"name":"tokens"}],"token_vault":{"enable_credential_manager":true}}`,
		},
		{
			name:    "mcp limits tls and token vault",
			request: &c.CreateAIGatewayMCPServerRequest{},
			input:   `{"type":"passthrough-listener","name":"tools","display_name":"Tools","config":{"url":"https://mcp.example","certificates":["client-cert"],"ca_certificates":["root-ca"]},"rate_limiting":{"limits":[{"method":"tools/call","limit":100,"window_size":60}]},"datastores":[{"name":"tokens"}],"token_vault":{"provider":"upstream"}}`,
			want:    `{"config":{"certificates":["client-cert"],"ca_certificates":["root-ca"]},"rate_limiting":{"limits":[{"method":"tools/call","limit":100,"window_size":60}]},"datastores":[{"name":"tokens"}],"token_vault":{"provider":"upstream"}}`,
		},
		{
			name:    "agent tls",
			request: &c.CreateAIGatewayAgentRequest{},
			input:   `{"name":"agent","display_name":"Agent","config":{"url":"https://agent.example","certificates":["client-cert"],"ca_certificates":["root-ca"]}}`,
			want:    `{"config":{"certificates":["client-cert"],"ca_certificates":["root-ca"]}}`,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			require.NoError(t, json.Unmarshal([]byte(tc.input), tc.request))
			encoded, err := json.Marshal(tc.request)
			require.NoError(t, err)
			var got, want any
			require.NoError(t, json.Unmarshal(encoded, &got))
			require.NoError(t, json.Unmarshal([]byte(tc.want), &want))
			requireJSONSubset(t, want, got)
		})
	}
}

func TestAIGateway23BedrockEndpointTypeOmitted(t *testing.T) {
	var request c.AIGatewayTargetBedrockConfig
	require.NoError(t, json.Unmarshal([]byte(`{"type":"bedrock"}`), &request))
	encoded, err := json.Marshal(request)
	require.NoError(t, err)
	var body map[string]any
	require.NoError(t, json.Unmarshal(encoded, &body))
	require.NotContains(t, body, "endpoint_type")
}

func TestAIGateway23WildcardEnums(t *testing.T) {
	format := c.AIGatewayModelFormatType("*")
	modelCapability := c.AIGatewayModelModelCapabilities("*")
	apiCapability := c.Capabilities("*")
	require.True(t, format.IsExact(), "wildcard must be a known format")
	require.True(t, modelCapability.IsExact(), "wildcard must be a known model capability")
	require.True(t, apiCapability.IsExact(), "wildcard must be a known API capability")
}

func TestAIGateway23WildcardRequests(t *testing.T) {
	for _, modelType := range []string{"model", "api"} {
		for _, operation := range []string{"create", "update"} {
			t.Run(modelType+"/"+operation, func(t *testing.T) {
				input := `{"type":"` + modelType + `","name":"wildcard","display_name":"Wildcard",` +
					`"formats":[{"type":"*"}],"capabilities":["*"],"config":{"route":{"paths":["/llm"]}},` +
					`"targets":[{"name":"upstream","provider":"provider","config":{"type":"openai"}}]}`
				var request any = &c.CreateAIGatewayModelRequest{}
				if operation == "update" {
					request = &c.UpdateAIGatewayModelRequest{}
				}
				require.NoError(t, json.Unmarshal([]byte(input), request))
				encoded, err := json.Marshal(request)
				require.NoError(t, err)
				var body map[string]any
				require.NoError(t, json.Unmarshal(encoded, &body))
				require.Equal(t, []any{map[string]any{"type": "*"}}, body["formats"])
				require.Equal(t, []any{"*"}, body["capabilities"])
			})
		}
	}
}

func requireJSONSubset(t *testing.T, want, got any) {
	t.Helper()
	switch expected := want.(type) {
	case map[string]any:
		actual, ok := got.(map[string]any)
		require.True(t, ok, "expected object, got %T", got)
		for key, value := range expected {
			require.Contains(t, actual, key)
			requireJSONSubset(t, value, actual[key])
		}
	case []any:
		actual, ok := got.([]any)
		require.True(t, ok, "expected array, got %T", got)
		require.Len(t, actual, len(expected))
		for i := range expected {
			requireJSONSubset(t, expected[i], actual[i])
		}
	default:
		require.Equal(t, want, got)
	}
}
