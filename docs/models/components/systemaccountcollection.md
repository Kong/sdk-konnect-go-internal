# SystemAccountCollection

A paginated list response for a collection of system accounts.


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            | Example                                                                |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `Meta`                                                                 | [*components.PaginatedMeta](../../models/components/paginatedmeta.md)  | :heavy_minus_sign:                                                     | returns the pagination information                                     | {<br/>"page": {<br/>"number": 2,<br/>"size": 10,<br/>"total": 100<br/>}<br/>} |
| `Data`                                                                 | [][components.SystemAccount](../../models/components/systemaccount.md) | :heavy_minus_sign:                                                     | N/A                                                                    |                                                                        |