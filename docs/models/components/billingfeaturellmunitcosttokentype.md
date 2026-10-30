# BillingFeatureLLMUnitCostTokenType

Static token type value. Use this when the feature tracks a single token type
(e.g., only input tokens). `request` is an alias for `input`, `response` is an
alias for `output`. Mutually exclusive with `token_type_property`.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingFeatureLLMUnitCostTokenTypeInput

// Open enum: custom values can be created with a direct type cast
custom := components.BillingFeatureLLMUnitCostTokenType("custom_value")
```


## Values

| Name                                           | Value                                          |
| ---------------------------------------------- | ---------------------------------------------- |
| `BillingFeatureLLMUnitCostTokenTypeInput`      | input                                          |
| `BillingFeatureLLMUnitCostTokenTypeOutput`     | output                                         |
| `BillingFeatureLLMUnitCostTokenTypeCacheRead`  | cache_read                                     |
| `BillingFeatureLLMUnitCostTokenTypeCacheWrite` | cache_write                                    |
| `BillingFeatureLLMUnitCostTokenTypeReasoning`  | reasoning                                      |
| `BillingFeatureLLMUnitCostTokenTypeRequest`    | request                                        |
| `BillingFeatureLLMUnitCostTokenTypeResponse`   | response                                       |