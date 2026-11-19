# AIGatewaySessionAffinityStorageLocal

Stores session affinity state in the memory of each data plane node. A request that
reaches a different node does not find the earlier target, so the balancer algorithm
picks the target.



## Fields

| Field              | Type               | Required           | Description        |
| ------------------ | ------------------ | ------------------ | ------------------ |
| `Type`             | `string`           | :heavy_check_mark: | N/A                |