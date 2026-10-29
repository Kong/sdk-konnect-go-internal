# ListNotificationEventsParamsFilterSubjectKeyStringFieldNEQFilter

Filter by the key of the subject the event refers to. Events without a subject
never match, not even with `neq`.


## Fields

| Field              | Type               | Required           | Description        |
| ------------------ | ------------------ | ------------------ | ------------------ |
| `Eq`               | `*string`          | :heavy_minus_sign: | N/A                |
| `Oeq`              | `string`           | :heavy_check_mark: | N/A                |
| `Neq`              | `string`           | :heavy_check_mark: | N/A                |