# MeshControlPlaneReadinessGroup

Which part of the estate a finding belongs to. These are identifiers rather than labels, so a reader that wants prettier wording supplies its own.


## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.MeshControlPlaneReadinessGroupControlPlane

// Open enum: custom values can be created with a direct type cast
custom := components.MeshControlPlaneReadinessGroup("custom_value")
```


## Values

| Name                                                  | Value                                                 |
| ----------------------------------------------------- | ----------------------------------------------------- |
| `MeshControlPlaneReadinessGroupControlPlane`          | control_plane                                         |
| `MeshControlPlaneReadinessGroupMeshObject`            | mesh_object                                           |
| `MeshControlPlaneReadinessGroupPolicies`              | policies                                              |
| `MeshControlPlaneReadinessGroupRemovedResources`      | removed_resources                                     |
| `MeshControlPlaneReadinessGroupDataPlaneAndWorkloads` | data_plane_and_workloads                              |
| `MeshControlPlaneReadinessGroupOther`                 | other                                                 |