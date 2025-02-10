# ListAPIDocumentResponse

List of API documents


## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      | Example                                                                          |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `Data`                                                                           | [][components.APIDocumentSummary](../../models/components/apidocumentsummary.md) | :heavy_check_mark:                                                               | N/A                                                                              |                                                                                  |
| `Meta`                                                                           | [components.PaginatedMeta](../../models/components/paginatedmeta.md)             | :heavy_check_mark:                                                               | returns the pagination information                                               | {<br/>"page": {<br/>"number": 2,<br/>"size": 10,<br/>"total": 100<br/>}<br/>}    |