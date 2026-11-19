# AIGatewayAvailablePolicy


## Fields

| Field                                                             | Type                                                              | Required                                                          | Description                                                       | Example                                                           |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `Name`                                                            | `string`                                                          | :heavy_check_mark:                                                | N/A                                                               | ai-sanitizer                                                      |
| `Scopes`                                                          | [][components.Scopes](../../models/components/scopes.md)          | :heavy_check_mark:                                                | N/A                                                               |                                                                   |
| `MinRuntimeVersion`                                               | `*string`                                                         | :heavy_minus_sign:                                                | The minimum AI Gateway runtime version that supports this policy. | 2.3                                                               |