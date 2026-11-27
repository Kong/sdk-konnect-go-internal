# PrincipalsCache

Shared Redis cache for Kong Identity principal lookups.

**Requires a minimum runtime version of `2.4`**.


## Fields

| Field                                                                                     | Type                                                                                      | Required                                                                                  | Description                                                                               |
| ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- |
| `Pepper`                                                                                  | `*string`                                                                                 | :heavy_minus_sign:                                                                        | The secret key for the HMAC-SHA256 that names the entries in the shared principal cache.<br/> |