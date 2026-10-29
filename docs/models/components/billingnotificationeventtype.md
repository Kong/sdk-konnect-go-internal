# BillingNotificationEventType

The type of the event.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingNotificationEventTypeEntitlementsBalanceThreshold

// Open enum: custom values can be created with a direct type cast
custom := components.BillingNotificationEventType("custom_value")
```


## Values

| Name                                                       | Value                                                      |
| ---------------------------------------------------------- | ---------------------------------------------------------- |
| `BillingNotificationEventTypeEntitlementsBalanceThreshold` | entitlements.balance.threshold                             |
| `BillingNotificationEventTypeEntitlementsReset`            | entitlements.reset                                         |
| `BillingNotificationEventTypeInvoiceCreated`               | invoice.created                                            |
| `BillingNotificationEventTypeInvoiceUpdated`               | invoice.updated                                            |