# Response

The response returned by the recipient.


## Fields

| Field                                                                                 | Type                                                                                  | Required                                                                              | Description                                                                           |
| ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- |
| `StatusCode`                                                                          | `*int`                                                                                | :heavy_minus_sign:                                                                    | The HTTP status code returned by the recipient. Absent when no response was<br/>received. |
| `Body`                                                                                | `string`                                                                              | :heavy_check_mark:                                                                    | The response body returned by the recipient. Empty when no body was received.         |
| `DurationMs`                                                                          | `int64`                                                                               | :heavy_check_mark:                                                                    | How long the delivery attempt took, in milliseconds.                                  |
| `URL`                                                                                 | `*string`                                                                             | :heavy_minus_sign:                                                                    | The URL the event was delivered to. Only set for webhook channels.                    |