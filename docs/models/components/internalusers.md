# InternalUsers

A get action paginated response of users and organization IDs.


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           | Example                                                               |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `Meta`                                                                | [*components.PaginatedMeta](../../models/components/paginatedmeta.md) | :heavy_minus_sign:                                                    | returns the pagination information                                    | {<br/>"page": {<br/>"number": 2,<br/>"size": 10,<br/>"total": 100<br/>}<br/>} |
| `Data`                                                                | [][components.UserOrgIDs](../../models/components/userorgids.md)      | :heavy_minus_sign:                                                    | N/A                                                                   |                                                                       |