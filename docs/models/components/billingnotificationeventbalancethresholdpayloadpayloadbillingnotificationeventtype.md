# BillingNotificationEventBalanceThresholdPayloadPayloadBillingNotificationEventType

What the threshold value is measured against.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingNotificationEventBalanceThresholdPayloadPayloadBillingNotificationEventTypeBalanceValue

// Open enum: custom values can be created with a direct type cast
custom := components.BillingNotificationEventBalanceThresholdPayloadPayloadBillingNotificationEventType("custom_value")
```


## Values

| Name                                                                                                | Value                                                                                               |
| --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- |
| `BillingNotificationEventBalanceThresholdPayloadPayloadBillingNotificationEventTypeBalanceValue`    | balance_value                                                                                       |
| `BillingNotificationEventBalanceThresholdPayloadPayloadBillingNotificationEventTypeUsagePercentage` | usage_percentage                                                                                    |
| `BillingNotificationEventBalanceThresholdPayloadPayloadBillingNotificationEventTypeUsageValue`      | usage_value                                                                                         |