# PatchKonnectConnectedApps

Konnect Connected Apps configuration fields to patch


## Fields

| Field                                                                                               | Type                                                                                                | Required                                                                                            | Description                                                                                         |
| --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- |
| `Enabled`                                                                                           | `*bool`                                                                                             | :heavy_minus_sign:                                                                                  | When present, enables or disables Connected Apps for Konnect (same effect as POST /connected-apps). |
| `Dcr`                                                                                               | [*components.PatchDcrSettings](../../models/components/patchdcrsettings.md)                         | :heavy_minus_sign:                                                                                  | Dynamic Client Registration settings to patch                                                       |