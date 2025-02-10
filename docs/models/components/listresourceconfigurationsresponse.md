# ListResourceConfigurationsResponse

A paginated list for a collection of resource configurations.


## Fields

| Field                                                                                  | Type                                                                                   | Required                                                                               | Description                                                                            | Example                                                                                |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `Meta`                                                                                 | [components.PaginatedMeta](../../models/components/paginatedmeta.md)                   | :heavy_check_mark:                                                                     | returns the pagination information                                                     | {<br/>"page": {<br/>"number": 2,<br/>"size": 10,<br/>"total": 100<br/>}<br/>}          |
| `Data`                                                                                 | [][components.ResourceConfiguration](../../models/components/resourceconfiguration.md) | :heavy_check_mark:                                                                     | N/A                                                                                    |                                                                                        |