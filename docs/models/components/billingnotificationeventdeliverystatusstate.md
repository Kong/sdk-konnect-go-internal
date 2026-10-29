# BillingNotificationEventDeliveryStatusState

The current delivery state.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingNotificationEventDeliveryStatusStateSuccess

// Open enum: custom values can be created with a direct type cast
custom := components.BillingNotificationEventDeliveryStatusState("custom_value")
```


## Values

| Name                                                   | Value                                                  |
| ------------------------------------------------------ | ------------------------------------------------------ |
| `BillingNotificationEventDeliveryStatusStateSuccess`   | success                                                |
| `BillingNotificationEventDeliveryStatusStateFailed`    | failed                                                 |
| `BillingNotificationEventDeliveryStatusStateSending`   | sending                                                |
| `BillingNotificationEventDeliveryStatusStatePending`   | pending                                                |
| `BillingNotificationEventDeliveryStatusStateResending` | resending                                              |