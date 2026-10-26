# MeshControlPlaneReadinessSummary

Counts of what the scan found.


## Fields

| Field                                                      | Type                                                       | Required                                                   | Description                                                | Example                                                    |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `Blockers`                                                 | `int64`                                                    | :heavy_check_mark:                                         | Findings that stop the upgrade.                            | 2                                                          |
| `Info`                                                     | `int64`                                                    | :heavy_check_mark:                                         | Findings that do not block the upgrade.                    | 3                                                          |
| `SystemFindings`                                           | `int64`                                                    | :heavy_check_mark:                                         | Findings about resources Kong owns rather than the reader. | 0                                                          |