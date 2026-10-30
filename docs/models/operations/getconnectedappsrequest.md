# GetConnectedAppsRequest


## Fields

| Field                                                   | Type                                                    | Required                                                | Description                                             | Example                                                 |
| ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- |
| `Domain`                                                | [operations.Domain](../../models/operations/domain.md)  | :heavy_check_mark:                                      | Domain for Connected Apps (konnect or portal)           |                                                         |
| `PortalID`                                              | `*string`                                               | :heavy_minus_sign:                                      | The ID of the portal. Required when domain is 'portal'. | d32d905a-ed33-46a3-a093-d8f536af9a8a                    |