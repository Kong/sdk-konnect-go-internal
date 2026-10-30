# ListAPIKeysQueryParamFilter

Filter API keys by field. Supported filter fields:
- `name`: exact match or contains. Example: `filter[name]="deploy-key"` or `filter[name][contains]="deploy"`
- `principal_type`: exact match on one or more principal types. Accepted values: `user`, `agent`, `group`. Example: `filter[principal_type]="user"` or `filter[principal_type][oeq]="user,agent,group"`
- `holder`: exact match or contains on the holder's name, matched against the principal's
`display_name` and the names of the groups the principal belongs to. Example: `filter[holder][contains]="jane"`



## Fields

| Field                                                                                     | Type                                                                                      | Required                                                                                  | Description                                                                               |
| ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- |
| `Name`                                                                                    | [*components.LegacyStringFieldFilter](../../models/components/legacystringfieldfilter.md) | :heavy_minus_sign:                                                                        | Filter using **one** of the following operators: `eq`, `contains`                         |
| `PrincipalType`                                                                           | [*operations.PrincipalType](../../models/operations/principaltype.md)                     | :heavy_minus_sign:                                                                        | Filter using **one** of the following operators: `eq`, `oeq`                              |
| `Holder`                                                                                  | [*components.LegacyStringFieldFilter](../../models/components/legacystringfieldfilter.md) | :heavy_minus_sign:                                                                        | Filter using **one** of the following operators: `eq`, `contains`                         |