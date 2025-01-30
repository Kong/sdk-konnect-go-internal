# ListControlPlanesResponse

A paginated list response for a collection of control planes.


## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          | Example                                                              |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `Meta`                                                               | [components.PaginatedMeta](../../models/components/paginatedmeta.md) | :heavy_check_mark:                                                   | Returns pagination information                                       | {<br/>"page": {<br/>"number": 2,<br/>"size": 10,<br/>"total": 100<br/>}<br/>} |
| `Data`                                                               | [][components.ControlPlane](../../models/components/controlplane.md) | :heavy_check_mark:                                                   | N/A                                                                  |                                                                      |