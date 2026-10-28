# DeliveryDetailResponse

A single delivery and all its attempts.


## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `Data`                                                                     | [components.Delivery](../../models/components/delivery.md)                 | :heavy_check_mark:                                                         | Properties of a delivery. The payload is present only on the detail view.  |
| `Attempts`                                                                 | [][components.DeliveryAttempt](../../models/components/deliveryattempt.md) | :heavy_minus_sign:                                                         | N/A                                                                        |