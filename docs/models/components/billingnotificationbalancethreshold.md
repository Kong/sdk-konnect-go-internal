# BillingNotificationBalanceThreshold

A balance threshold of a notification rule. Crossing it generates an
`entitlements.balance.threshold` event.


## Fields

| Field                                                                                                                    | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `Type`                                                                                                                   | [components.BillingNotificationBalanceThresholdType](../../models/components/billingnotificationbalancethresholdtype.md) | :heavy_check_mark:                                                                                                       | What the threshold value is measured against.                                                                            |
| `Value`                                                                                                                  | `float64`                                                                                                                | :heavy_check_mark:                                                                                                       | The threshold value.                                                                                                     |