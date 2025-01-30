# TeamGroupMappingCollection

A paginated collection of mappings grouped by team_id.


## Fields

| Field                                                                        | Type                                                                         | Required                                                                     | Description                                                                  | Example                                                                      |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `Meta`                                                                       | [*components.PaginatedMeta](../../models/components/paginatedmeta.md)        | :heavy_minus_sign:                                                           | Returns pagination information                                               | {<br/>"page": {<br/>"number": 2,<br/>"size": 10,<br/>"total": 100<br/>}<br/>} |
| `Data`                                                                       | [][components.TeamGroupMapping](../../models/components/teamgroupmapping.md) | :heavy_minus_sign:                                                           | N/A                                                                          |                                                                              |