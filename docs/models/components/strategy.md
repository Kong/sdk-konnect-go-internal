# Strategy

How the balancer selects among the eligible targets. `weighted` scores each target on cost and
quality. `margin` selects the cheapest target whose expected quality is close enough to the best
target's. Available while both are evaluated; removed once one is chosen.


## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.StrategyWeighted

// Open enum: custom values can be created with a direct type cast
custom := components.Strategy("custom_value")
```


## Values

| Name               | Value              |
| ------------------ | ------------------ |
| `StrategyWeighted` | weighted           |
| `StrategyMargin`   | margin             |