# AIGatewayModelBalancerIntelligentClassifierAuthBasic

Sends one header, such as an API key, with each request to the classifier service.

**Requires a minimum runtime version of `2.5`**.


## Fields

| Field                                                        | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `Type`                                                       | `string`                                                     | :heavy_check_mark:                                           | N/A                                                          |
| `Headers`                                                    | [][components.Headers](../../models/components/headers.md)   | :heavy_minus_sign:                                           | The header sent with each request to the classifier service. |