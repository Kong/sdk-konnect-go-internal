# OrgConfigurationNamespace

Namespace of an org scoped event. It lists the same values as NotificationNamespace, internal ones included, because the platform API dev spec cannot reference a schema that has internal values.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.OrgConfigurationNamespacePlanAndUsage

// Open enum: custom values can be created with a direct type cast
custom := components.OrgConfigurationNamespace("custom_value")
```


## Values

| Name                                              | Value                                             |
| ------------------------------------------------- | ------------------------------------------------- |
| `OrgConfigurationNamespacePlanAndUsage`           | plan-and-usage                                    |
| `OrgConfigurationNamespaceOrganization`           | organization                                      |
| `OrgConfigurationNamespaceDevPortal`              | dev-portal                                        |
| `OrgConfigurationNamespaceCloudGateways`          | cloud-gateways                                    |
| `OrgConfigurationNamespaceRegionalEventNamespace` | regional-event-namespace                          |
| `OrgConfigurationNamespaceGatewayManager`         | gateway-manager                                   |
| `OrgConfigurationNamespaceAdvancedAnalytics`      | advanced-analytics                                |
| `OrgConfigurationNamespaceAiManagement`           | ai-management                                     |