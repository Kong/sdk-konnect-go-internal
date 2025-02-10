# ListApplicationsResponse

A paginated list of applications in a portal.


## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          | Example                                                              |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `Meta`                                                               | [components.PaginatedMeta](../../models/components/paginatedmeta.md) | :heavy_check_mark:                                                   | returns the pagination information                                   | {<br/>"page": {<br/>"number": 2,<br/>"size": 10,<br/>"total": 100<br/>}<br/>} |
| `Data`                                                               | [][components.Application](../../models/components/application.md)   | :heavy_check_mark:                                                   | N/A                                                                  |                                                                      |