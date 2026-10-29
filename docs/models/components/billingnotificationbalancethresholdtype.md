# BillingNotificationBalanceThresholdType

What the threshold value is measured against.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingNotificationBalanceThresholdTypeBalanceValue

// Open enum: custom values can be created with a direct type cast
custom := components.BillingNotificationBalanceThresholdType("custom_value")
```


## Values

| Name                                                     | Value                                                    |
| -------------------------------------------------------- | -------------------------------------------------------- |
| `BillingNotificationBalanceThresholdTypeBalanceValue`    | balance_value                                            |
| `BillingNotificationBalanceThresholdTypeUsagePercentage` | usage_percentage                                         |
| `BillingNotificationBalanceThresholdTypeUsageValue`      | usage_value                                              |