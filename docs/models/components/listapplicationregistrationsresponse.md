# ListApplicationRegistrationsResponse

A paginated list of application registrations.


## Fields

| Field                                                                                      | Type                                                                                       | Required                                                                                   | Description                                                                                | Example                                                                                    |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `Meta`                                                                                     | [components.PaginatedMeta](../../models/components/paginatedmeta.md)                       | :heavy_check_mark:                                                                         | returns the pagination information                                                         | {<br/>"page": {<br/>"number": 2,<br/>"size": 10,<br/>"total": 100<br/>}<br/>}              |
| `Data`                                                                                     | [][components.ApplicationRegistration](../../models/components/applicationregistration.md) | :heavy_check_mark:                                                                         | N/A                                                                                        |                                                                                            |