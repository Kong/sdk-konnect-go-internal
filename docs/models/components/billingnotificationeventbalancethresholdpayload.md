# BillingNotificationEventBalanceThresholdPayload

A balance threshold notification event payload.


## Fields

| Field                                                            | Type                                                             | Required                                                         | Description                                                      | Example                                                          |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `ID`                                                             | `string`                                                         | :heavy_check_mark:                                               | The identifier of the event the payload belongs to.              | 01G65Z755AFWAKHE12NY0CQ9FH                                       |
| `Type`                                                           | [components.PayloadType](../../models/components/payloadtype.md) | :heavy_check_mark:                                               | The type of the event.                                           |                                                                  |
| `Timestamp`                                                      | [time.Time](https://pkg.go.dev/time#Time)                        | :heavy_check_mark:                                               | When the event was generated.                                    | 2023-01-01T01:01:01.001Z                                         |
| `Data`                                                           | [components.PayloadData](../../models/components/payloaddata.md) | :heavy_check_mark:                                               | The entities and threshold the event refers to.                  |                                                                  |