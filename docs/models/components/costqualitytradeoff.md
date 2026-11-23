# CostQualityTradeoff

How the balancer weighs cost versus quality when selecting a target. `cost` favors the cheapest available target, `quality` favors the highest-quality target, and `balanced` weighs both.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.CostQualityTradeoffBalanced

// Open enum: custom values can be created with a direct type cast
custom := components.CostQualityTradeoff("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `CostQualityTradeoffBalanced` | balanced                      |
| `CostQualityTradeoffCost`     | cost                          |
| `CostQualityTradeoffQuality`  | quality                       |