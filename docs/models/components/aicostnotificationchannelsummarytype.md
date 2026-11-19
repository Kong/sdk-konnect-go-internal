# AICostNotificationChannelSummaryType

Whether the channel sends email or posts to Slack.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AICostNotificationChannelSummaryTypeEmail

// Open enum: custom values can be created with a direct type cast
custom := components.AICostNotificationChannelSummaryType("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `AICostNotificationChannelSummaryTypeEmail` | email                                       |
| `AICostNotificationChannelSummaryTypeSlack` | slack                                       |