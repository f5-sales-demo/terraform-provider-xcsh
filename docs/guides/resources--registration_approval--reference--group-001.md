---
page_title: "xcsh_registration_approval reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_registration_approval reference."
---

# xcsh_registration_approval reference

<a id="canonical-2950841dcd2dad45aa090a989e2c5989663f4229329496cccde1786e7269c4b0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-68561b1031623123a847b383dd4473edfa134e1ce1a770b1ce7e997f32408d6d"></a>

## Property reference — Property reference / 67833e82b3c1 / 2

Breadcrumbs:

- [xcsh_registration_approval](../resources/registration_approval.md#canonical-7d4fb674027240378875a2202f657eddfb167f6c0f01b29cddfd793bc70315ba)
- Property reference

<a id="canonical-3ab78e3de67b958c291a7c462192cefe1fc78c6c24305fc4a98dffc44e1ca45f"></a>

## Direct properties — Property reference / 67833e82b3c1 / 3

<a id="canonical-9ba7031ebf13f2aad3741fb83a2b23b01a0afbde988701080c9b8e19561c3a63"></a>

<a id="canonical-229cefb120721bfc9cf03a8a062df0dcc838f66d4730d0c462b572b2c083712f"></a>

## backup_connected_region property — Property reference / 67833e82b3c1 / 4

Type: `"string"`. Optional.

<a id="canonical-647b0977f9da554575c9aecdd29b1cf223e17629e975c7bb15a198d4cadb8639"></a>

<a id="canonical-94d9f5fff95ca0cb4fcaaa89220b5f63e0dcf072eb93bb1c65edd3b1606de387"></a>

## cluster_size property — Property reference / 67833e82b3c1 / 5

Type: `"number"`. Required.

Number of nodes in the registration's site cluster. Use 1 for a single-node site and the complete
node count for an HA site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.OneOf(1, 3),
}
```

<a id="canonical-1ce1fb68539eb197efabc67c4fe11dc7eb0060aed85977c17a84c9d83da0d889"></a>

<a id="canonical-7bb8f288ee0924117d74d59f143e8c4e1bc3c06281f0026c40a92d42786c90a1"></a>

## connected_region property — Property reference / 67833e82b3c1 / 6

Type: `"string"`. Optional.

<a id="canonical-320bb803f15bff93de7fec5ecddfc93fa64932c753b92bdc822c000ddf4c9090"></a>

<a id="canonical-0d5c47aa92996552da91a2d90884cd4f5864790a655647fa84e25ed8ca62ce47"></a>

## id property — Property reference / 67833e82b3c1 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-f3288fb78c98926bee5dccf2c071a0b77e094412541bd95d46d76403c94289d6"></a>

<a id="canonical-d6e113ea7d04767bf2b01781b00c80f5919b7c1782311dd38349aa409cb844dc"></a>

## name property — Property reference / 67833e82b3c1 / 8

Type: `"string"`. Required.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NameValidator(),
}
```

<a id="canonical-1c973126eeaf41c31884b475ae51d5d22f5305f76d0978e2dfe5b445c5f16a08"></a>

<a id="canonical-18ae0715f28b360340c6a2afc3814a2d2d4acba16bb0cbf2db8f93003dfdd1f0"></a>

## namespace property — Property reference / 67833e82b3c1 / 9

Type: `"string"`. Required.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NamespaceValidator(),
}
```

<a id="canonical-313afdd2e5ae5f9266875f8c5db9e7dcedf4bb0eaaea3baf4700e2ae05c19b9b"></a>

<a id="canonical-f53b07fa5b65bfaca419f87badde836ea7b66edde81c09707a3ad07239048cd6"></a>

## preferred_active_re property — Property reference / 67833e82b3c1 / 10

Type: `"string"`. Optional.

<a id="canonical-32e06422c43c33eb083a99a362fd8e8af1b2df610ca75498452874b839532b3b"></a>

<a id="canonical-d8e847f8f5f0611144f03c15643b655c4c08de08e88d747b1187c34ba553d37c"></a>

## state property — Property reference / 67833e82b3c1 / 11

Type: `"string"`. Optional, Computed.

Provider validators and defaults (from schema source):

```go
Default: stringdefault.StaticString("APPROVED")
```

<a id="canonical-a0cbd13e8d5b95c613f03ba0bd3240dbfc11a0528090c9aa0e7132f927d66df4"></a>

## All schema paths — Property reference / 67833e82b3c1 / 12

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `backup_connected_region` | [backup_connected_region](resources--registration_approval--reference--group-001.md#canonical-9ba7031ebf13f2aad3741fb83a2b23b01a0afbde988701080c9b8e19561c3a63) |
| `cluster_size` | [cluster_size](resources--registration_approval--reference--group-001.md#canonical-647b0977f9da554575c9aecdd29b1cf223e17629e975c7bb15a198d4cadb8639) |
| `connected_region` | [connected_region](resources--registration_approval--reference--group-001.md#canonical-1ce1fb68539eb197efabc67c4fe11dc7eb0060aed85977c17a84c9d83da0d889) |
| `id` | [id](resources--registration_approval--reference--group-001.md#canonical-320bb803f15bff93de7fec5ecddfc93fa64932c753b92bdc822c000ddf4c9090) |
| `name` | [name](resources--registration_approval--reference--group-001.md#canonical-f3288fb78c98926bee5dccf2c071a0b77e094412541bd95d46d76403c94289d6) |
| `namespace` | [namespace](resources--registration_approval--reference--group-001.md#canonical-1c973126eeaf41c31884b475ae51d5d22f5305f76d0978e2dfe5b445c5f16a08) |
| `preferred_active_re` | [preferred_active_re](resources--registration_approval--reference--group-001.md#canonical-313afdd2e5ae5f9266875f8c5db9e7dcedf4bb0eaaea3baf4700e2ae05c19b9b) |
| `state` | [state](resources--registration_approval--reference--group-001.md#canonical-32e06422c43c33eb083a99a362fd8e8af1b2df610ca75498452874b839532b3b) |

<a id="canonical-d111633310a3c4eda77504f236fe6b645f7cb8306f01d2af3bd2abd2944b3cb9"></a>

## Next pages — Property reference / 67833e82b3c1 / 13

- [xcsh_registration_approval](../resources/registration_approval.md#canonical-7d4fb674027240378875a2202f657eddfb167f6c0f01b29cddfd793bc70315ba)
