# BillingNotificationEventDeliveryAttemptState

The delivery state reached by this attempt.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingNotificationEventDeliveryAttemptStateSuccess

// Open enum: custom values can be created with a direct type cast
custom := components.BillingNotificationEventDeliveryAttemptState("custom_value")
```


## Values

| Name                                                    | Value                                                   |
| ------------------------------------------------------- | ------------------------------------------------------- |
| `BillingNotificationEventDeliveryAttemptStateSuccess`   | success                                                 |
| `BillingNotificationEventDeliveryAttemptStateFailed`    | failed                                                  |
| `BillingNotificationEventDeliveryAttemptStateSending`   | sending                                                 |
| `BillingNotificationEventDeliveryAttemptStatePending`   | pending                                                 |
| `BillingNotificationEventDeliveryAttemptStateResending` | resending                                               |