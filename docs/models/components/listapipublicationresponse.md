# ListAPIPublicationResponse

Paginated list of API publications


## Fields

| Field                                                                                    | Type                                                                                     | Required                                                                                 | Description                                                                              | Example                                                                                  |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `Meta`                                                                                   | [components.PaginatedMeta](../../models/components/paginatedmeta.md)                     | :heavy_check_mark:                                                                       | returns the pagination information                                                       | {<br/>"page": {<br/>"number": 2,<br/>"size": 10,<br/>"total": 100<br/>}<br/>}            |
| `Data`                                                                                   | [][components.APIPublicationListItem](../../models/components/apipublicationlistitem.md) | :heavy_check_mark:                                                                       | N/A                                                                                      |                                                                                          |