# ReportCollection

A paginated list response for a collection of reports.


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           | Example                                                               |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `Meta`                                                                | [*components.PaginatedMeta](../../models/components/paginatedmeta.md) | :heavy_minus_sign:                                                    | Returns pagination information                                        | {<br/>"page": {<br/>"number": 2,<br/>"size": 10,<br/>"total": 100<br/>}<br/>} |
| `Data`                                                                | [][components.Report](../../models/components/report.md)              | :heavy_minus_sign:                                                    | N/A                                                                   |                                                                       |