# UpdateDcrSettings

Dynamic Client Registration settings for the auth server


## Fields

| Field                                                                                                     | Type                                                                                                      | Required                                                                                                  | Description                                                                                               |
| --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- |
| `RedirectURIAllowlist`                                                                                    | []`string`                                                                                                | :heavy_minus_sign:                                                                                        | Redirect URIs allowed for clients registered via open Dynamic Client Registration                         |
| `DefaultAccessTokenDuration`                                                                              | `*int64`                                                                                                  | :heavy_minus_sign:                                                                                        | The default access token duration, in seconds, applied to DCR clients registered against this auth server |