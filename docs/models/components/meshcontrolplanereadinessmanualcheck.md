# MeshControlPlaneReadinessManualCheck

One upgrade item the control plane API cannot answer.


## Fields

| Field                                                                          | Type                                                                           | Required                                                                       | Description                                                                    | Example                                                                        |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `Title`                                                                        | `string`                                                                       | :heavy_check_mark:                                                             | N/A                                                                            | Gateway API resources                                                          |
| `Detail`                                                                       | `*string`                                                                      | :heavy_check_mark:                                                             | N/A                                                                            | Check each Gateway resource by hand; the control plane API does not list them. |
| `Command`                                                                      | `*string`                                                                      | :heavy_check_mark:                                                             | A command that answers the item, and null when there is none to give.          | kubectl get gateways -A                                                        |