# ListIdentitiesQueryParamFilter

Filter identities by field. Supported filter field:
- `type`: exact match on identity type. Accepted values: `oidc`, `consumer`, `custom`, `auth_server_client`.

Example: `filter[type]="oidc"`



## Fields

| Field                                                                                      | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `Type`                                                                                     | [*components.StringFieldOEQFilter](../../models/components/stringfieldoeqfilter.md)        | :heavy_minus_sign:                                                                         | Returns entities that exact match any of the comma-delimited phrases in the filter string. |