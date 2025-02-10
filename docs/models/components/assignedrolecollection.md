# AssignedRoleCollection

A paginated list response for a collection of assigned roles.


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           | Example                                                               |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `Meta`                                                                | [*components.PaginatedMeta](../../models/components/paginatedmeta.md) | :heavy_minus_sign:                                                    | returns the pagination information                                    | {<br/>"page": {<br/>"number": 2,<br/>"size": 10,<br/>"total": 100<br/>}<br/>} |
| `Data`                                                                | [][components.AssignedRole](../../models/components/assignedrole.md)  | :heavy_minus_sign:                                                    | An Array                                                              |                                                                       |