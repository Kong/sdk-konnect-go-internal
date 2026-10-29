# BillingNotificationEventRuleType

The type of event the rule generates.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingNotificationEventRuleTypeEntitlementsBalanceThreshold

// Open enum: custom values can be created with a direct type cast
custom := components.BillingNotificationEventRuleType("custom_value")
```


## Values

| Name                                                           | Value                                                          |
| -------------------------------------------------------------- | -------------------------------------------------------------- |
| `BillingNotificationEventRuleTypeEntitlementsBalanceThreshold` | entitlements.balance.threshold                                 |
| `BillingNotificationEventRuleTypeEntitlementsReset`            | entitlements.reset                                             |
| `BillingNotificationEventRuleTypeInvoiceCreated`               | invoice.created                                                |
| `BillingNotificationEventRuleTypeInvoiceUpdated`               | invoice.updated                                                |