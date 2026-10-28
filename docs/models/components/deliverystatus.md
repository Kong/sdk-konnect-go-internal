# DeliveryStatus

Lifecycle state of a delivery. A scheduled delivery waits for its next attempt, or is being sent. A delivery that reached its expiry time before an attempt settled it is expired. No delivery shows in_flight yet, because the status ignores the lease of a worker, and nothing cancels a delivery, so no delivery is canceled. A filter accepts both values and keeps nothing.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.DeliveryStatusScheduled

// Open enum: custom values can be created with a direct type cast
custom := components.DeliveryStatus("custom_value")
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `DeliveryStatusScheduled` | scheduled                 |
| `DeliveryStatusInFlight`  | in_flight                 |
| `DeliveryStatusSucceeded` | succeeded                 |
| `DeliveryStatusFailed`    | failed                    |
| `DeliveryStatusExpired`   | expired                   |
| `DeliveryStatusCanceled`  | canceled                  |