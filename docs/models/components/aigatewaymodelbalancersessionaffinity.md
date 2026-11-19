# AIGatewayModelBalancerSessionAffinity

Sends later requests in a conversation to the target that served the earlier turn, if that
target can still take the request. This improves reuse of the provider's prompt cache.
Requests are matched by conversation content and provider response IDs, not by a session
header. Requests with the same conversation content can go to the same target. If the
preferred target is not available, the balancer algorithm picks the target.



## Fields

| Field                                                                                                     | Type                                                                                                      | Required                                                                                                  | Description                                                                                               |
| --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- |
| `TTL`                                                                                                     | `*int64`                                                                                                  | :heavy_minus_sign:                                                                                        | The time in seconds that a conversation's preferred target is kept.                                       |
| `Storage`                                                                                                 | [*components.AIGatewaySessionAffinityStorage](../../models/components/aigatewaysessionaffinitystorage.md) | :heavy_minus_sign:                                                                                        | Where session affinity state is stored.                                                                   |
| `Enabled`                                                                                                 | `*bool`                                                                                                   | :heavy_minus_sign:                                                                                        | Whether session affinity is enabled.                                                                      |