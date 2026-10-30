# PatchIdentityQuotaRequest


## Fields

| Field                                                                        | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `QuotaType`                                                                  | [components.IdentityQuotaType](../../models/components/identityquotatype.md) | :heavy_check_mark:                                                           | The type of resource the quota applies to                                    |
| `QuotaPatchBody`                                                             | [components.QuotaPatchBody](../../models/components/quotapatchbody.md)       | :heavy_check_mark:                                                           | Quota fields to partially update                                             |