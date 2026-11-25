# AIGatewayModelBalancerIntelligentConfigFailoverCriteria

**`http_400` requires a minimum runtime version of `2.3`**.

**`http_401` requires a minimum runtime version of `2.3`**.

**`http_402` requires a minimum runtime version of `2.3`**.

**`http_405` requires a minimum runtime version of `2.3`**.

**`http_408` requires a minimum runtime version of `2.3`**.

**`http_409` requires a minimum runtime version of `2.3`**.

**`http_413` requires a minimum runtime version of `2.3`**.

**`http_415` requires a minimum runtime version of `2.3`**.

**`http_422` requires a minimum runtime version of `2.3`**.

**`http_522` requires a minimum runtime version of `2.3`**.

**`http_529` requires a minimum runtime version of `2.3`**.

**`http_5xx` requires a minimum runtime version of `2.3`**.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayModelBalancerIntelligentConfigFailoverCriteriaError

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayModelBalancerIntelligentConfigFailoverCriteria("custom_value")
```


## Values

| Name                                                                   | Value                                                                  |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `AIGatewayModelBalancerIntelligentConfigFailoverCriteriaError`         | error                                                                  |
| `AIGatewayModelBalancerIntelligentConfigFailoverCriteriaHttp400`       | http_400                                                               |
| `AIGatewayModelBalancerIntelligentConfigFailoverCriteriaHttp401`       | http_401                                                               |
| `AIGatewayModelBalancerIntelligentConfigFailoverCriteriaHttp402`       | http_402                                                               |
| `AIGatewayModelBalancerIntelligentConfigFailoverCriteriaHttp403`       | http_403                                                               |
| `AIGatewayModelBalancerIntelligentConfigFailoverCriteriaHttp404`       | http_404                                                               |
| `AIGatewayModelBalancerIntelligentConfigFailoverCriteriaHttp405`       | http_405                                                               |
| `AIGatewayModelBalancerIntelligentConfigFailoverCriteriaHttp408`       | http_408                                                               |
| `AIGatewayModelBalancerIntelligentConfigFailoverCriteriaHttp409`       | http_409                                                               |
| `AIGatewayModelBalancerIntelligentConfigFailoverCriteriaHttp413`       | http_413                                                               |
| `AIGatewayModelBalancerIntelligentConfigFailoverCriteriaHttp415`       | http_415                                                               |
| `AIGatewayModelBalancerIntelligentConfigFailoverCriteriaHttp422`       | http_422                                                               |
| `AIGatewayModelBalancerIntelligentConfigFailoverCriteriaHttp429`       | http_429                                                               |
| `AIGatewayModelBalancerIntelligentConfigFailoverCriteriaHttp500`       | http_500                                                               |
| `AIGatewayModelBalancerIntelligentConfigFailoverCriteriaHttp502`       | http_502                                                               |
| `AIGatewayModelBalancerIntelligentConfigFailoverCriteriaHttp503`       | http_503                                                               |
| `AIGatewayModelBalancerIntelligentConfigFailoverCriteriaHttp504`       | http_504                                                               |
| `AIGatewayModelBalancerIntelligentConfigFailoverCriteriaHttp522`       | http_522                                                               |
| `AIGatewayModelBalancerIntelligentConfigFailoverCriteriaHttp529`       | http_529                                                               |
| `AIGatewayModelBalancerIntelligentConfigFailoverCriteriaHttp5xx`       | http_5xx                                                               |
| `AIGatewayModelBalancerIntelligentConfigFailoverCriteriaInvalidHeader` | invalid_header                                                         |
| `AIGatewayModelBalancerIntelligentConfigFailoverCriteriaNonIdempotent` | non_idempotent                                                         |
| `AIGatewayModelBalancerIntelligentConfigFailoverCriteriaTimeout`       | timeout                                                                |