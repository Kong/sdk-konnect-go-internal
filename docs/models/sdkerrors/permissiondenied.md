# PermissionDenied

The error object.


## Fields

| Field                                                   | Type                                                    | Required                                                | Description                                             | Example                                                 |
| ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- |
| `Type`                                                  | **string*                                               | :heavy_minus_sign:                                      | Documentation for this error.                           |                                                         |
| `Status`                                                | **int64*                                                | :heavy_minus_sign:                                      | HTTP status code.                                       | 403                                                     |
| `Title`                                                 | **string*                                               | :heavy_minus_sign:                                      | HTTP status code                                        | Bad Request                                             |
| `Instance`                                              | **string*                                               | :heavy_minus_sign:                                      | Konnect traceback error code.                           | konnect:trace:3674017986744198214                       |
| `Detail`                                                | **string*                                               | :heavy_minus_sign:                                      | Information about the error response.                   | Only org admins or analytics admins can create reports. |