# ListNetworksResponse

A paginated list for a collection of networks.


## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          | Example                                                              |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `Meta`                                                               | [components.PaginatedMeta](../../models/components/paginatedmeta.md) | :heavy_check_mark:                                                   | Returns pagination information                                       | {<br/>"page": {<br/>"number": 2,<br/>"size": 10,<br/>"total": 100<br/>}<br/>} |
| `Data`                                                               | [][components.Network](../../models/components/network.md)           | :heavy_check_mark:                                                   | N/A                                                                  |                                                                      |