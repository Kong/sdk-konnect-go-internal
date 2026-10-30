# ClientJwkItem

The client JWK


## Fields

| Field                                                            | Type                                                             | Required                                                         | Description                                                      | Example                                                          |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `ID`                                                             | `string`                                                         | :heavy_check_mark:                                               | The ID of the client JWK                                         |                                                                  |
| `Kid`                                                            | `string`                                                         | :heavy_check_mark:                                               | The key ID (RFC 7638 SHA-256 thumbprint) identifying this key.   |                                                                  |
| `Alg`                                                            | `string`                                                         | :heavy_check_mark:                                               | The signing algorithm the key is used with, e.g. RS256 or ES256. |                                                                  |
| `Type`                                                           | `string`                                                         | :heavy_check_mark:                                               | The key type, e.g. RSA or EC.                                    |                                                                  |
| `CreatedAt`                                                      | [time.Time](https://pkg.go.dev/time#Time)                        | :heavy_check_mark:                                               | An ISO-8601 timestamp representation of entity creation date.    | 2022-11-04T20:10:06.927Z                                         |
| `UpdatedAt`                                                      | [time.Time](https://pkg.go.dev/time#Time)                        | :heavy_check_mark:                                               | An ISO-8601 timestamp representation of entity update date.      | 2022-11-04T20:10:06.927Z                                         |