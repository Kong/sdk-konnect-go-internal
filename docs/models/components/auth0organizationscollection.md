# Auth0OrganizationsCollection

A paginated list response for a collection of auth0 organizations.


## Fields

| Field                                                                          | Type                                                                           | Required                                                                       | Description                                                                    | Example                                                                        |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `Meta`                                                                         | [*components.PaginatedMeta](../../models/components/paginatedmeta.md)          | :heavy_minus_sign:                                                             | returns the pagination information                                             | {<br/>"page": {<br/>"number": 2,<br/>"size": 10,<br/>"total": 100<br/>}<br/>}  |
| `Data`                                                                         | [][components.Auth0Organization](../../models/components/auth0organization.md) | :heavy_minus_sign:                                                             | N/A                                                                            |                                                                                |