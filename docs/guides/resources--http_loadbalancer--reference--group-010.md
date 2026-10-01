---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-4d5baf9d5458be02410b796428adf3e47e4a291f6db7808cf5aa2c8dce622ca1"></a>

## api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / 656a27444800 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_specification](resources--http_loadbalancer--reference--group-008.md#canonical-e39f46ee17d4d8075f3d741ef024d557ae4090d0ab94a0b4063a8cce9d955c10)
- [api_specification.validation_custom_list](resources--http_loadbalancer--reference--group-009.md#canonical-d0b19f04e97ec7b3a15988db3291da92c5fec8ada37dc4e08e29ad4a8768bce4)
- [api_specification.validation_custom_list.open_api_validation_rules](resources--http_loadbalancer--reference--group-009.md#canonical-871283c10740c5f020e9ec61c8e18d845ce5a2124b0cf332c843dc80bf985601)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](resources--http_loadbalancer--reference--group-009.md#canonical-7092af96e2bbe1a0001f3798a1a4c0196c698be541608d84309a038923d194bf)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active

<a id="canonical-f7ea33ee0993fd8f998ada8b837a00eadfe25af5508de6436e81045421637696"></a>

Type: `"object"`. single nested block, Optional.

Enable OpenAPI validation and explicitly select enforcement\_report to allow and log invalid
traffic, or enforcement\_block to reject invalid requests with HTTP 403.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("request_validation_properties"),
  validators.ConflictingObjectAttributes("enforcement_block",
    "enforcement_report")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-validation_enforcement_type": "[\"enforcement_block\",\"enforcement_report\"]"
}
```

Terraform syntax:

```terraform
validation_mode_active {
  # Configure direct properties listed below.
}
```

<a id="canonical-4d9451a777bdd59c34b1402554f4a9869d45cd38724964fb5b659032bfc0b9ab"></a>

## Direct properties — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / 656a27444800 / 3

- [enforcement_block](resources--http_loadbalancer--reference--group-010.md#canonical-e350895e423a355d3675baf3035d71e9e6051202d9a730908c73127b32a9faef): complete subsection reference.

- [enforcement_report](resources--http_loadbalancer--reference--group-010.md#canonical-f5c22323778d3245e4bab3b0a6ede74d278e011578dd6ea4e349ce9fd8808020): complete subsection reference.

<a id="canonical-0c555852911178fb575967c6bf93fe1afc054a4518e301a32374de0143e5fb74"></a>

<a id="canonical-d5e4a890827df074331a3f04799ec82b6c39933bb4d48af8172e3625499ce951"></a>

## request_validation_properties property — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / 656a27444800 / 4

Type: `["list", "string"]`. Optional.

\[Enum:
PROPERTY\_QUERY\_PARAMETERS|PROPERTY\_PATH\_PARAMETERS|PROPERTY\_CONTENT\_TYPE|PROPERTY\_COOKIE\_PARAMETERS|PROPERTY\_HTTP\_HEADERS|PROPERTY\_HTTP\_BODY|PROPERTY\_SECURITY\_SCHEMA|PROPERTY\_RESPONSE\_CODE\]
List of properties of the request to validate according to the OpenAPI specification file (a.k.a.
Swagger). Possible values are \`PROPERTY\_QUERY\_PARAMETERS\`, \`PROPERTY\_PATH\_PARAMETERS\`,
\`PROPERTY\_CONTENT\_TYPE\`, \`PROPERTY\_COOKIE\_PARAMETERS\`, \`PROPERTY\_HTTP\_HEADERS\`,
\`PROPERTY\_HTTP\_BODY\`, \`PROPERTY\_SECURITY\_SCHEMA\`, \`PROPERTY\_RESPONSE\_CODE\`. Defaults to
\`PROPERTY\_QUERY\_PARAMETERS\`.

Upstream description:

List of properties of the request to validate according to the OpenAPI specification file (a.k.a.
Swagger)

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtLeast(1),
}
```

Receipt-pinned upstream constraints:

```json
{
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[7]",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[7]",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-74f36715c91bc1e7c1d142a61e759e51e57bdc26e980a9820255f593393ed60a"></a>

## Next pages — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / 656a27444800 / 5

- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active.enforcement_block](resources--http_loadbalancer--reference--group-010.md#canonical-e350895e423a355d3675baf3035d71e9e6051202d9a730908c73127b32a9faef)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active.enforcement_report](resources--http_loadbalancer--reference--group-010.md#canonical-f5c22323778d3245e4bab3b0a6ede74d278e011578dd6ea4e349ce9fd8808020)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](resources--http_loadbalancer--reference--group-009.md#canonical-7092af96e2bbe1a0001f3798a1a4c0196c698be541608d84309a038923d194bf)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-e350895e423a355d3675baf3035d71e9e6051202d9a730908c73127b32a9faef"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b954d2e697f46b8d3fa4dae53762f2a67d47c9bf05669e0085e8a4e23c903ed2"></a>

## api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active.enforcement_block — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / c823db54de41 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_specification](resources--http_loadbalancer--reference--group-008.md#canonical-e39f46ee17d4d8075f3d741ef024d557ae4090d0ab94a0b4063a8cce9d955c10)
- [api_specification.validation_custom_list](resources--http_loadbalancer--reference--group-009.md#canonical-d0b19f04e97ec7b3a15988db3291da92c5fec8ada37dc4e08e29ad4a8768bce4)
- [api_specification.validation_custom_list.open_api_validation_rules](resources--http_loadbalancer--reference--group-009.md#canonical-871283c10740c5f020e9ec61c8e18d845ce5a2124b0cf332c843dc80bf985601)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](resources--http_loadbalancer--reference--group-009.md#canonical-7092af96e2bbe1a0001f3798a1a4c0196c698be541608d84309a038923d194bf)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active](resources--http_loadbalancer--reference--group-009.md#canonical-dc6a3e18b1c100d758cfe54d98983bac9119a64fd12b057253c5e470994f9dd8)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active.enforcement_block

<a id="canonical-92302800317a96245651256d1411db20de7bcf697388dd9ecfb3ea8d8f419eae"></a>

Type: `["object", {}]`. Optional.

Blocking validation: reject traffic that violates the selected OpenAPI validation properties.
Invalid requests are returned as HTTP 403.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
enforcement_block = {}
```

<a id="canonical-fc74cb85d4936a22cc8a84480c20cdd7548d9f82fe447cf0f7ea68d5c52f8363"></a>

## Direct properties — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / c823db54de41 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e88e2a1c619dd5ba53a361b5d5d3543cb402487ccf2f613556a314e3b461adce"></a>

## Next pages — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / c823db54de41 / 4

- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active](resources--http_loadbalancer--reference--group-009.md#canonical-dc6a3e18b1c100d758cfe54d98983bac9119a64fd12b057253c5e470994f9dd8)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-f5c22323778d3245e4bab3b0a6ede74d278e011578dd6ea4e349ce9fd8808020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0cc6e7b2709dda5a02ae781fa9cda3271fa29239e67f7a4481c21290d5712a1d"></a>

## api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active.enforcement_report — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / aaee39a2d364 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_specification](resources--http_loadbalancer--reference--group-008.md#canonical-e39f46ee17d4d8075f3d741ef024d557ae4090d0ab94a0b4063a8cce9d955c10)
- [api_specification.validation_custom_list](resources--http_loadbalancer--reference--group-009.md#canonical-d0b19f04e97ec7b3a15988db3291da92c5fec8ada37dc4e08e29ad4a8768bce4)
- [api_specification.validation_custom_list.open_api_validation_rules](resources--http_loadbalancer--reference--group-009.md#canonical-871283c10740c5f020e9ec61c8e18d845ce5a2124b0cf332c843dc80bf985601)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](resources--http_loadbalancer--reference--group-009.md#canonical-7092af96e2bbe1a0001f3798a1a4c0196c698be541608d84309a038923d194bf)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active](resources--http_loadbalancer--reference--group-009.md#canonical-dc6a3e18b1c100d758cfe54d98983bac9119a64fd12b057253c5e470994f9dd8)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active.enforcement_report

<a id="canonical-ccabd88deac1bef73afdf6c7fd4b4647541ddaba701a7452053609d63cea33eb"></a>

Type: `["object", {}]`. Optional.

Report-only validation: record OpenAPI violations while allowing the request or response to
continue.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
enforcement_report = {}
```

<a id="canonical-7c2605a7900b9de40acd5442a025983c3fa6c2f0c1b26441e6b23eaf09f9378b"></a>

## Direct properties — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / aaee39a2d364 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5a9fe1ae2439d8dbb51602bbecfddc97fb6f51e01149dd750526fe23ad169a9f"></a>

## Next pages — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / aaee39a2d364 / 4

- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active](resources--http_loadbalancer--reference--group-009.md#canonical-dc6a3e18b1c100d758cfe54d98983bac9119a64fd12b057253c5e470994f9dd8)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-983c0f7b8b65130a581f82668eae61ae86c32f3cb977721f3ee830baad3f8e2b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a26096d21ee7a117945029b9d91d746e7a085c637aa34893a85e2970add8b85e"></a>

## api_specification.validation_custom_list.settings — api_specification.validation_custom_list.settings / 465e4cd2f734 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_specification](resources--http_loadbalancer--reference--group-008.md#canonical-e39f46ee17d4d8075f3d741ef024d557ae4090d0ab94a0b4063a8cce9d955c10)
- [api_specification.validation_custom_list](resources--http_loadbalancer--reference--group-009.md#canonical-d0b19f04e97ec7b3a15988db3291da92c5fec8ada37dc4e08e29ad4a8768bce4)
- api_specification.validation_custom_list.settings

<a id="canonical-113609f22beb4b08a35b1d48e2ecd550ce57aee6ca08ffdb535706650ad5e726"></a>

Type: `"object"`. single nested block, Optional.

OpenAPI specification validation settings relevant for 'API Inventory' enforcement and for 'Custom
list' enforcement.

Upstream description:

OpenAPI specification validation settings relevant for "API Inventory" enforcement and for "Custom
list" enforcement.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("oversized_body_fail_validation",
    "oversized_body_skip_validation"),
  validators.ConflictingObjectAttributes("property_validation_settings_custom",
    "property_validation_settings_default")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-fail_configuration": "[]",
  "x-ves-oneof-field-oversized_body_choice": "[\"oversized_body_fail_validation\",\"oversized_body_skip_validation\"]",
  "x-ves-oneof-field-property_validation_settings_choice": "[\"property_validation_settings_custom\",\"property_validation_settings_default\"]"
}
```

Terraform syntax:

```terraform
settings {
  # Configure direct properties listed below.
}
```

<a id="canonical-119141bb857e858a4751803442d4d5047098b890fefa47bedfcd9cf49aff59c6"></a>

## Direct properties — api_specification.validation_custom_list.settings / 465e4cd2f734 / 3

- [oversized_body_fail_validation](resources--http_loadbalancer--reference--group-010.md#canonical-d745dc9e89b055234e652c311835cc067423886015727f8c3ca3c49d2677d7f0): complete subsection reference.

- [oversized_body_skip_validation](resources--http_loadbalancer--reference--group-010.md#canonical-23f2d3a01cda963f2ac765e5a71d74e494c8ebf86ae29cbfda30c7d4a03e1429): complete subsection reference.

- [property_validation_settings_custom](resources--http_loadbalancer--reference--group-010.md#canonical-480f314b3310e6fa40e0fc2201a789115dc6cbbceccc96517e4e2b1387badce6): complete subsection reference.

- [property_validation_settings_default](resources--http_loadbalancer--reference--group-010.md#canonical-e375936ed81ce8b3f0e8c4228cc983e477e09d3e86b5b1c4a83f91e23b1e30ed): complete subsection reference.

<a id="canonical-217e289fdf97074ff46eddf01902f504cde47e36c56251e78ce08308071416f4"></a>

## Next pages — api_specification.validation_custom_list.settings / 465e4cd2f734 / 4

- [api_specification.validation_custom_list.settings.oversized_body_fail_validation](resources--http_loadbalancer--reference--group-010.md#canonical-d745dc9e89b055234e652c311835cc067423886015727f8c3ca3c49d2677d7f0)
- [api_specification.validation_custom_list.settings.oversized_body_skip_validation](resources--http_loadbalancer--reference--group-010.md#canonical-23f2d3a01cda963f2ac765e5a71d74e494c8ebf86ae29cbfda30c7d4a03e1429)
- [api_specification.validation_custom_list.settings.property_validation_settings_custom](resources--http_loadbalancer--reference--group-010.md#canonical-480f314b3310e6fa40e0fc2201a789115dc6cbbceccc96517e4e2b1387badce6)
- [api_specification.validation_custom_list.settings.property_validation_settings_default](resources--http_loadbalancer--reference--group-010.md#canonical-e375936ed81ce8b3f0e8c4228cc983e477e09d3e86b5b1c4a83f91e23b1e30ed)
- [api_specification.validation_custom_list](resources--http_loadbalancer--reference--group-009.md#canonical-d0b19f04e97ec7b3a15988db3291da92c5fec8ada37dc4e08e29ad4a8768bce4)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-d745dc9e89b055234e652c311835cc067423886015727f8c3ca3c49d2677d7f0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-60eba72dd99656d0d8528ec68527d41caba87af2c8d57173916e84e63c893758"></a>

## api_specification.validation_custom_list.settings.oversized_body_fail_validation — api_specification.validation_custom_list.settings.oversized_body_fail_validation / 3d908874fad5 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_specification](resources--http_loadbalancer--reference--group-008.md#canonical-e39f46ee17d4d8075f3d741ef024d557ae4090d0ab94a0b4063a8cce9d955c10)
- [api_specification.validation_custom_list](resources--http_loadbalancer--reference--group-009.md#canonical-d0b19f04e97ec7b3a15988db3291da92c5fec8ada37dc4e08e29ad4a8768bce4)
- [api_specification.validation_custom_list.settings](resources--http_loadbalancer--reference--group-010.md#canonical-983c0f7b8b65130a581f82668eae61ae86c32f3cb977721f3ee830baad3f8e2b)
- api_specification.validation_custom_list.settings.oversized_body_fail_validation

<a id="canonical-7bc87332c70e84d58443a768e46c412e6e682fac6626f2e613c6bc0be8dea02a"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
oversized_body_fail_validation = {}
```

<a id="canonical-e793c3e198b96ba8a4ed65f97561abe77c9c4daf20492cd47cfe57f06ce5cf55"></a>

## Direct properties — api_specification.validation_custom_list.settings.oversized_body_fail_validation / 3d908874fad5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-622468cd40cf558269f5c56f99e4e77aa9c3dec26cf35a24df7b34ed810d88b5"></a>

## Next pages — api_specification.validation_custom_list.settings.oversized_body_fail_validation / 3d908874fad5 / 4

- [api_specification.validation_custom_list.settings](resources--http_loadbalancer--reference--group-010.md#canonical-983c0f7b8b65130a581f82668eae61ae86c32f3cb977721f3ee830baad3f8e2b)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-23f2d3a01cda963f2ac765e5a71d74e494c8ebf86ae29cbfda30c7d4a03e1429"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-751dc82fa7138206274a430cbbb941341e5f0bd033a06859d53306e8de6b41e9"></a>

## api_specification.validation_custom_list.settings.oversized_body_skip_validation — api_specification.validation_custom_list.settings.oversized_body_skip_validation / 9eb8248af5c2 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_specification](resources--http_loadbalancer--reference--group-008.md#canonical-e39f46ee17d4d8075f3d741ef024d557ae4090d0ab94a0b4063a8cce9d955c10)
- [api_specification.validation_custom_list](resources--http_loadbalancer--reference--group-009.md#canonical-d0b19f04e97ec7b3a15988db3291da92c5fec8ada37dc4e08e29ad4a8768bce4)
- [api_specification.validation_custom_list.settings](resources--http_loadbalancer--reference--group-010.md#canonical-983c0f7b8b65130a581f82668eae61ae86c32f3cb977721f3ee830baad3f8e2b)
- api_specification.validation_custom_list.settings.oversized_body_skip_validation

<a id="canonical-e7708ed02bbee915e792c4d153a56908553e1554ca73dbab1284b35927967079"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
oversized_body_skip_validation = {}
```

<a id="canonical-1984b1f5ac23e0e6d648637478c4a786090139b30a5aaa9f64698e40960a756b"></a>

## Direct properties — api_specification.validation_custom_list.settings.oversized_body_skip_validation / 9eb8248af5c2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9d3073886795277c534b3e739d57ca1fd945d9a93a6088e69c375c0fa6c240db"></a>

## Next pages — api_specification.validation_custom_list.settings.oversized_body_skip_validation / 9eb8248af5c2 / 4

- [api_specification.validation_custom_list.settings](resources--http_loadbalancer--reference--group-010.md#canonical-983c0f7b8b65130a581f82668eae61ae86c32f3cb977721f3ee830baad3f8e2b)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-480f314b3310e6fa40e0fc2201a789115dc6cbbceccc96517e4e2b1387badce6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7523b0e9c421f6145e052cf035a54ca8d81a697b8c715441bb395ce002a641c3"></a>

## api_specification.validation_custom_list.settings.property_validation_settings_custom — api_specification.validation_custom_list.settings.property_validation_settings_c / a85325b01740 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_specification](resources--http_loadbalancer--reference--group-008.md#canonical-e39f46ee17d4d8075f3d741ef024d557ae4090d0ab94a0b4063a8cce9d955c10)
- [api_specification.validation_custom_list](resources--http_loadbalancer--reference--group-009.md#canonical-d0b19f04e97ec7b3a15988db3291da92c5fec8ada37dc4e08e29ad4a8768bce4)
- [api_specification.validation_custom_list.settings](resources--http_loadbalancer--reference--group-010.md#canonical-983c0f7b8b65130a581f82668eae61ae86c32f3cb977721f3ee830baad3f8e2b)
- api_specification.validation_custom_list.settings.property_validation_settings_custom

<a id="canonical-925dae0436093d5300d57f93ed0adff1437546786787b4b688504b35c55ff0e9"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for property validation settings custom.

Upstream description:

Custom property validation settings.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
property_validation_settings_custom {
  # Configure direct properties listed below.
}
```

<a id="canonical-c42797d30d1bd9f72ab177a117ba5654de648be497046fb789910b0c68667ba3"></a>

## Direct properties — api_specification.validation_custom_list.settings.property_validation_settings_c / a85325b01740 / 3

- [query_parameters](resources--http_loadbalancer--reference--group-010.md#canonical-c93388cf2b1cb41fa78ba883a3ced2d1e31fc23d4cbf1987edb39c8fbf09335a): complete subsection reference.

<a id="canonical-b654c7ff7001a9ff4448342ea783747c745abd5c0f1752bfe103a3e4124812e6"></a>

## Next pages — api_specification.validation_custom_list.settings.property_validation_settings_c / a85325b01740 / 4

- [api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters](resources--http_loadbalancer--reference--group-010.md#canonical-c93388cf2b1cb41fa78ba883a3ced2d1e31fc23d4cbf1987edb39c8fbf09335a)
- [api_specification.validation_custom_list.settings](resources--http_loadbalancer--reference--group-010.md#canonical-983c0f7b8b65130a581f82668eae61ae86c32f3cb977721f3ee830baad3f8e2b)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-c93388cf2b1cb41fa78ba883a3ced2d1e31fc23d4cbf1987edb39c8fbf09335a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e2b7c12b90da5467a796cef64925d27cc216a1df81f1651862ae4f915750f9a6"></a>

## api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters — api_specification.validation_custom_list.settings.property_validation_settings_c / 375eaec81e6b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_specification](resources--http_loadbalancer--reference--group-008.md#canonical-e39f46ee17d4d8075f3d741ef024d557ae4090d0ab94a0b4063a8cce9d955c10)
- [api_specification.validation_custom_list](resources--http_loadbalancer--reference--group-009.md#canonical-d0b19f04e97ec7b3a15988db3291da92c5fec8ada37dc4e08e29ad4a8768bce4)
- [api_specification.validation_custom_list.settings](resources--http_loadbalancer--reference--group-010.md#canonical-983c0f7b8b65130a581f82668eae61ae86c32f3cb977721f3ee830baad3f8e2b)
- [api_specification.validation_custom_list.settings.property_validation_settings_custom](resources--http_loadbalancer--reference--group-010.md#canonical-480f314b3310e6fa40e0fc2201a789115dc6cbbceccc96517e4e2b1387badce6)
- api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters

<a id="canonical-3d9ba41315f63264bebe16a9349bf825f6cdd3d8ce972337ca6d7c6ee2576fa2"></a>

Type: `"object"`. single nested block, Optional.

Custom settings for query parameters validation.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("allow_additional_parameters",
    "disallow_additional_parameters")}
```

Terraform syntax:

```terraform
query_parameters {
  # Configure direct properties listed below.
}
```

<a id="canonical-92dc00ee750b55e7b3eb888ad1bc4328cf700749e7eb2af4866beb4d31e40d98"></a>

## Direct properties — api_specification.validation_custom_list.settings.property_validation_settings_c / 375eaec81e6b / 3

- [allow_additional_parameters](resources--http_loadbalancer--reference--group-010.md#canonical-955d6de145b6845fea91743227367a55ffd976bcfc70494c6afe71abb8a23e81): complete subsection reference.

- [disallow_additional_parameters](resources--http_loadbalancer--reference--group-010.md#canonical-eac2a060491bb983f39aca164ec345ab9fb54ef9e6a378317f4ef0ef0fdcdbbf): complete subsection reference.

<a id="canonical-104452712aad96b8992f2cd4b3b49d1b13bab3ae020cd821a7540f8bd4114438"></a>

## Next pages — api_specification.validation_custom_list.settings.property_validation_settings_c / 375eaec81e6b / 4

- [api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters.allow_additional_parameters](resources--http_loadbalancer--reference--group-010.md#canonical-955d6de145b6845fea91743227367a55ffd976bcfc70494c6afe71abb8a23e81)
- [api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters.disallow_additional_parameters](resources--http_loadbalancer--reference--group-010.md#canonical-eac2a060491bb983f39aca164ec345ab9fb54ef9e6a378317f4ef0ef0fdcdbbf)
- [api_specification.validation_custom_list.settings.property_validation_settings_custom](resources--http_loadbalancer--reference--group-010.md#canonical-480f314b3310e6fa40e0fc2201a789115dc6cbbceccc96517e4e2b1387badce6)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-955d6de145b6845fea91743227367a55ffd976bcfc70494c6afe71abb8a23e81"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-36a9641ac0848345295c484d225ad9c90302300362ccb2460f62e03796bedbbd"></a>

## api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters.allow_additional_parameters — api_specification.validation_custom_list.settings.property_validation_settings_c / 5adb94c1348c / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_specification](resources--http_loadbalancer--reference--group-008.md#canonical-e39f46ee17d4d8075f3d741ef024d557ae4090d0ab94a0b4063a8cce9d955c10)
- [api_specification.validation_custom_list](resources--http_loadbalancer--reference--group-009.md#canonical-d0b19f04e97ec7b3a15988db3291da92c5fec8ada37dc4e08e29ad4a8768bce4)
- [api_specification.validation_custom_list.settings](resources--http_loadbalancer--reference--group-010.md#canonical-983c0f7b8b65130a581f82668eae61ae86c32f3cb977721f3ee830baad3f8e2b)
- [api_specification.validation_custom_list.settings.property_validation_settings_custom](resources--http_loadbalancer--reference--group-010.md#canonical-480f314b3310e6fa40e0fc2201a789115dc6cbbceccc96517e4e2b1387badce6)
- [api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters](resources--http_loadbalancer--reference--group-010.md#canonical-c93388cf2b1cb41fa78ba883a3ced2d1e31fc23d4cbf1987edb39c8fbf09335a)
- api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters.allow_additional_parameters

<a id="canonical-7bf970724a67e004c0f4ccf3419c6a65c598276868569c48ef90efba0365f859"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for allow additional parameters.

Terraform syntax:

```terraform
allow_additional_parameters = {}
```

<a id="canonical-be84a217718951f5e1d2b6c75026d4d87fc90b11372630a0643203715099fbf9"></a>

## Direct properties — api_specification.validation_custom_list.settings.property_validation_settings_c / 5adb94c1348c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-507105975f31ce257a5f938c475ee8c209fe3aa1d1f2a8cd2b3372c6d6f6b23b"></a>

## Next pages — api_specification.validation_custom_list.settings.property_validation_settings_c / 5adb94c1348c / 4

- [api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters](resources--http_loadbalancer--reference--group-010.md#canonical-c93388cf2b1cb41fa78ba883a3ced2d1e31fc23d4cbf1987edb39c8fbf09335a)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-eac2a060491bb983f39aca164ec345ab9fb54ef9e6a378317f4ef0ef0fdcdbbf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9c446a7c86867392f6f99d065866090f9009e8d9e1aa794fdf10517513753683"></a>

## api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters.disallow_additional_parameters — api_specification.validation_custom_list.settings.property_validation_settings_c / 86ccdce17066 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_specification](resources--http_loadbalancer--reference--group-008.md#canonical-e39f46ee17d4d8075f3d741ef024d557ae4090d0ab94a0b4063a8cce9d955c10)
- [api_specification.validation_custom_list](resources--http_loadbalancer--reference--group-009.md#canonical-d0b19f04e97ec7b3a15988db3291da92c5fec8ada37dc4e08e29ad4a8768bce4)
- [api_specification.validation_custom_list.settings](resources--http_loadbalancer--reference--group-010.md#canonical-983c0f7b8b65130a581f82668eae61ae86c32f3cb977721f3ee830baad3f8e2b)
- [api_specification.validation_custom_list.settings.property_validation_settings_custom](resources--http_loadbalancer--reference--group-010.md#canonical-480f314b3310e6fa40e0fc2201a789115dc6cbbceccc96517e4e2b1387badce6)
- [api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters](resources--http_loadbalancer--reference--group-010.md#canonical-c93388cf2b1cb41fa78ba883a3ced2d1e31fc23d4cbf1987edb39c8fbf09335a)
- api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters.disallow_additional_parameters

<a id="canonical-7446d13951d53b3b6dc3b3e3044541cf25b1d3b6c9152e867719f51435ab58c2"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disallow additional parameters.

Terraform syntax:

```terraform
disallow_additional_parameters = {}
```

<a id="canonical-a78b8d3db8f496790c6baaedf4ef35936acbda35f45d018e8feaedded3d84c18"></a>

## Direct properties — api_specification.validation_custom_list.settings.property_validation_settings_c / 86ccdce17066 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-40ba3ccb05e929895cc9fd662155426cbca0afab346278ce66377d467fe5b77b"></a>

## Next pages — api_specification.validation_custom_list.settings.property_validation_settings_c / 86ccdce17066 / 4

- [api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters](resources--http_loadbalancer--reference--group-010.md#canonical-c93388cf2b1cb41fa78ba883a3ced2d1e31fc23d4cbf1987edb39c8fbf09335a)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-e375936ed81ce8b3f0e8c4228cc983e477e09d3e86b5b1c4a83f91e23b1e30ed"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a82ca729ee5d2029d879fe78eab47c36a54185b70a7df96c6889e934783e96f5"></a>

## api_specification.validation_custom_list.settings.property_validation_settings_default — api_specification.validation_custom_list.settings.property_validation_settings_d / aceff7ff9619 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_specification](resources--http_loadbalancer--reference--group-008.md#canonical-e39f46ee17d4d8075f3d741ef024d557ae4090d0ab94a0b4063a8cce9d955c10)
- [api_specification.validation_custom_list](resources--http_loadbalancer--reference--group-009.md#canonical-d0b19f04e97ec7b3a15988db3291da92c5fec8ada37dc4e08e29ad4a8768bce4)
- [api_specification.validation_custom_list.settings](resources--http_loadbalancer--reference--group-010.md#canonical-983c0f7b8b65130a581f82668eae61ae86c32f3cb977721f3ee830baad3f8e2b)
- api_specification.validation_custom_list.settings.property_validation_settings_default

<a id="canonical-a24b96f46a9a5ec77f91e78412ecd9c001d2576b47785b3161dbf078dde944f0"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for property validation settings default.

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
property_validation_settings_default = {}
```

<a id="canonical-6f385c89d2bd22e4c661c57fe7497fdbaee8830c7dae9412fb27da1d708a2624"></a>

## Direct properties — api_specification.validation_custom_list.settings.property_validation_settings_d / aceff7ff9619 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c089123e216aae74f32b86fb6c82ce813901ebadf9bdd1fa7bcd3f74d13cdef9"></a>

## Next pages — api_specification.validation_custom_list.settings.property_validation_settings_d / aceff7ff9619 / 4

- [api_specification.validation_custom_list.settings](resources--http_loadbalancer--reference--group-010.md#canonical-983c0f7b8b65130a581f82668eae61ae86c32f3cb977721f3ee830baad3f8e2b)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-37e8d4458cfcaef85a98370e5f94e240f8bd8191af346c1e16886e800d749ece"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6a49439323c2a633eaebe4c96fabcc3c50975ca19df26cf2139f483f3169ef8a"></a>

## api_specification.validation_disabled — api_specification.validation_disabled / 31fc0350e6a8 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_specification](resources--http_loadbalancer--reference--group-008.md#canonical-e39f46ee17d4d8075f3d741ef024d557ae4090d0ab94a0b4063a8cce9d955c10)
- api_specification.validation_disabled

<a id="canonical-9b69275d003ac1160a4bf129b7044bb83674b7c2f5a56e6d5b2f11aac7ec3d25"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
validation_disabled = {}
```

<a id="canonical-1cf917510f6761098b571f8bec3b6a4a0ba2c676012ecf0c0e627b8837e62a1c"></a>

## Direct properties — api_specification.validation_disabled / 31fc0350e6a8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d5e579646086483c6734402aeffe30ad55992c442980b6e25b30cbe74c4291b6"></a>

## Next pages — api_specification.validation_disabled / 31fc0350e6a8 / 4

- [api_specification](resources--http_loadbalancer--reference--group-008.md#canonical-e39f46ee17d4d8075f3d741ef024d557ae4090d0ab94a0b4063a8cce9d955c10)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-699d08bedad7c6a3db855143b0dfa0f5040e8efa4d08cd58dccbf39afde96a71"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-13865affa20f3273433b0fbec2f6305866f329f89c3be4cfd15c0b617ed8df29"></a>

## api_testing — api_testing / 392862aa2e4f / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- api_testing

<a id="canonical-eeef58c1e470002deefd08411cd8129df347d5e8a7f2df5a2ec84396f86db68b"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: api\_testing, disable\_api\_testing; Default: disable\_api\_testing\] API Testing.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("domains"),
  validators.ConflictingObjectAttributes("every_day",
    "every_month"),
  validators.ConflictingObjectAttributes("every_day",
    "every_week"),
  validators.ConflictingObjectAttributes("every_month",
    "every_week")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-frequency_choice": "[\"every_day\",\"every_month\",\"every_week\"]"
}
```

OneOf alternatives in this subsection:

- [api_testing](resources--http_loadbalancer--reference--group-010.md#canonical-eeef58c1e470002deefd08411cd8129df347d5e8a7f2df5a2ec84396f86db68b)
- [disable_api_testing](resources--http_loadbalancer--reference--group-017.md#canonical-92fdffa0592551285acab640901a3e81c0a77f5c572709339b273f36a12809fc)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
api_testing {
  # Configure direct properties listed below.
}
```

<a id="canonical-f8a1dac3d4856cecd1f7af397f3e3e50adde104e961d92e0045cbb30ce8fe869"></a>

## Direct properties — api_testing / 392862aa2e4f / 3

<a id="canonical-437d6ad71ff8b03e45203916aadc3767a2f2b4855128a4e76c861dc0317980d2"></a>

<a id="canonical-aab33e24d1c846b031f8ea33a93ee18f2b3f165e50c17c818d6932a3f4f4a7fe"></a>

## custom_header_value property — api_testing / 392862aa2e4f / 4

Type: `"string"`. Optional.

Add x-F5-API-testing-identifier header value to prevent security flags on API testing traffic.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

- [domains](resources--http_loadbalancer--reference--group-010.md#canonical-687b0526d20593ad003d8453ffecc5d3b2c352a7eaf26c5ec5b05ee3acc553ca): complete subsection reference.

- [every_day](resources--http_loadbalancer--reference--group-010.md#canonical-6f5a123aaa33c1ff4b16bdb712f9d78af4cdae144c66528c2fd0e7b95f34650d): complete subsection reference.

- [every_month](resources--http_loadbalancer--reference--group-010.md#canonical-7df86bfc68fe31586cd8cd007d1b3882017c016ccef11fe919977a1226c5ca90): complete subsection reference.

- [every_week](resources--http_loadbalancer--reference--group-010.md#canonical-0e2a352cfeac99dc8c252e6360e205819f252fddd47683e704ceed7790c20ff6): complete subsection reference.

<a id="canonical-09e2151953e058043d528685f98843e9261a47008d27cf9062730f38e5e05472"></a>

## Next pages — api_testing / 392862aa2e4f / 5

- [api_testing.domains](resources--http_loadbalancer--reference--group-010.md#canonical-687b0526d20593ad003d8453ffecc5d3b2c352a7eaf26c5ec5b05ee3acc553ca)
- [api_testing.every_day](resources--http_loadbalancer--reference--group-010.md#canonical-6f5a123aaa33c1ff4b16bdb712f9d78af4cdae144c66528c2fd0e7b95f34650d)
- [api_testing.every_month](resources--http_loadbalancer--reference--group-010.md#canonical-7df86bfc68fe31586cd8cd007d1b3882017c016ccef11fe919977a1226c5ca90)
- [api_testing.every_week](resources--http_loadbalancer--reference--group-010.md#canonical-0e2a352cfeac99dc8c252e6360e205819f252fddd47683e704ceed7790c20ff6)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-687b0526d20593ad003d8453ffecc5d3b2c352a7eaf26c5ec5b05ee3acc553ca"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-739181a47e4a4b42a9d8837693dbf556eea1396bcd48a86bd5e48ab24a085c4c"></a>

## api_testing.domains — api_testing.domains / c6d79efa8c27 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_testing](resources--http_loadbalancer--reference--group-010.md#canonical-699d08bedad7c6a3db855143b0dfa0f5040e8efa4d08cd58dccbf39afde96a71)
- api_testing.domains

<a id="canonical-dabdbaf0df7bfeec357ea126e5cd6b0b94ca47c326f38554c54b4c28e9d812c8"></a>

Type: `"object"`. list nested block, Optional.

Add and configure testing domains and credentials.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("credentials",
    "domain")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32"
  }
}
```

Terraform syntax:

```terraform
domains {
  # Configure direct properties listed below.
}
```

<a id="canonical-7b45b53c20c0a25fcf5a5717da9dd5ce1b540412c1ac9666edd9d911ac42245f"></a>

## Direct properties — api_testing.domains / c6d79efa8c27 / 3

<a id="canonical-de7e7b7830d79705e9c1bdac6903fef9a6d85d0dd3423f419b8453de64bb92fd"></a>

<a id="canonical-75383fab4f5c26d2d5789df641b9d6774abf3fc4865a5040384bb14b4b8cc50e"></a>

## allow_destructive_methods property — api_testing.domains / c6d79efa8c27 / 4

Type: `"bool"`. Optional.

Enable to allow API Testing to execute against destructive methods. Use with caution as these may
modify or DELETE data.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [credentials](resources--http_loadbalancer--reference--group-010.md#canonical-b610213d02bd2bc48360d13b69ca5c914988db8ebbb077cfa1edf5a349059f36): complete subsection reference.

<a id="canonical-7d9003456fe515cc4483464d93cc446ef06f1216cc85bd5816fa6b4235579d52"></a>

<a id="canonical-cc3e9dae144a075dfbfb2a05980bcea418d926968aae4fe3f016db984662e44f"></a>

## domain property — api_testing.domains / c6d79efa8c27 / 5

Type: `"string"`. Optional.

Add your testing environment domain. Be aware that running tests on a production domain can impact
live applications, as API testing cannot distinguish between production and testing environments.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.vh_domain": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.vh_domain": "true"
  }
}
```

<a id="canonical-72d1b0b6585da3d600bd457d12dc0a914f3e2d3d5990407a5b5a3722447ca7a2"></a>

## Next pages — api_testing.domains / c6d79efa8c27 / 6

- [api_testing.domains.credentials](resources--http_loadbalancer--reference--group-010.md#canonical-b610213d02bd2bc48360d13b69ca5c914988db8ebbb077cfa1edf5a349059f36)
- [api_testing](resources--http_loadbalancer--reference--group-010.md#canonical-699d08bedad7c6a3db855143b0dfa0f5040e8efa4d08cd58dccbf39afde96a71)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-b610213d02bd2bc48360d13b69ca5c914988db8ebbb077cfa1edf5a349059f36"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-24179ed5ca3c77a06fce1434fd51071a6ad18800a425b44b103ed2da45733c90"></a>

## api_testing.domains.credentials — api_testing.domains.credentials / 30ed209a0fde / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_testing](resources--http_loadbalancer--reference--group-010.md#canonical-699d08bedad7c6a3db855143b0dfa0f5040e8efa4d08cd58dccbf39afde96a71)
- [api_testing.domains](resources--http_loadbalancer--reference--group-010.md#canonical-687b0526d20593ad003d8453ffecc5d3b2c352a7eaf26c5ec5b05ee3acc553ca)
- api_testing.domains.credentials

<a id="canonical-6c959de0280e5189cda63ab0ec9036bcfdd8ae81161f41821983000f8b30d327"></a>

Type: `"object"`. list nested block, Optional.

Add credentials for API testing to use in the selected environment.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("credential_name"),
  validators.ConflictingListObjectAttributes("admin",
    "standard"),
  validators.ConflictingListObjectAttributes("api_key",
    "basic_auth"),
  validators.ConflictingListObjectAttributes("api_key",
    "bearer_token"),
  validators.ConflictingListObjectAttributes("api_key",
    "login_endpoint"),
  validators.ConflictingListObjectAttributes("basic_auth",
    "bearer_token"),
  validators.ConflictingListObjectAttributes("basic_auth",
    "login_endpoint"),
  validators.ConflictingListObjectAttributes("bearer_token",
    "login_endpoint")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

Terraform syntax:

```terraform
credentials {
  # Configure direct properties listed below.
}
```

<a id="canonical-158fbfbfcb0b56c074d092a0870659385394ddeb3fd76d3bfab8e88a2e816b8c"></a>

## Direct properties — api_testing.domains.credentials / 30ed209a0fde / 3

- [admin](resources--http_loadbalancer--reference--group-010.md#canonical-7a172f7eab2281c028bbf38c8461f2327fbb07d344cd5bf251c478dd01532e45): complete subsection reference.

- [api_key](resources--http_loadbalancer--reference--group-010.md#canonical-133753a73ec1db28898c1e436e98dde126bc7489d1ccacd34ea71506d4749b5c): complete subsection reference.

- [basic_auth](resources--http_loadbalancer--reference--group-010.md#canonical-ba6ae20ae6e213b12c2649468bd7833395743cd8b07901eb48f1673fe04c0dd8): complete subsection reference.

- [bearer_token](resources--http_loadbalancer--reference--group-010.md#canonical-6633e13b98f83d6d632c1bc8b169ea978c0c926699633f813b8e5517d3705b5c): complete subsection reference.

<a id="canonical-dc7289aa5ddebe0dbc19a9c01acd6ef6452e613a2dd50448f1413a8471d6521e"></a>

<a id="canonical-8a9430c67cc0eb262b3f831ad4d74f1050990ae379f672c418cab52c01ec272c"></a>

## credential_name property — api_testing.domains.credentials / 30ed209a0fde / 4

Type: `"string"`. Optional.

Enter a unique name for the credentials used in API testing.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

- [login_endpoint](resources--http_loadbalancer--reference--group-010.md#canonical-5c162cb70a77a04a40b0c7666a4b3f299a6a0815866bc29413b5748c87bdbd51): complete subsection reference.

- [standard](resources--http_loadbalancer--reference--group-010.md#canonical-022b9c587fad201b235813190250e458695f4c8fa89ebe6944a69960dbacee82): complete subsection reference.

<a id="canonical-2bddf822c5a9983ecd1eaf011a8cee2cb115b6e8499fb6cb0e31137dfb0cbb6a"></a>

## Next pages — api_testing.domains.credentials / 30ed209a0fde / 5

- [api_testing.domains.credentials.admin](resources--http_loadbalancer--reference--group-010.md#canonical-7a172f7eab2281c028bbf38c8461f2327fbb07d344cd5bf251c478dd01532e45)
- [api_testing.domains.credentials.api_key](resources--http_loadbalancer--reference--group-010.md#canonical-133753a73ec1db28898c1e436e98dde126bc7489d1ccacd34ea71506d4749b5c)
- [api_testing.domains.credentials.basic_auth](resources--http_loadbalancer--reference--group-010.md#canonical-ba6ae20ae6e213b12c2649468bd7833395743cd8b07901eb48f1673fe04c0dd8)
- [api_testing.domains.credentials.bearer_token](resources--http_loadbalancer--reference--group-010.md#canonical-6633e13b98f83d6d632c1bc8b169ea978c0c926699633f813b8e5517d3705b5c)
- [api_testing.domains.credentials.login_endpoint](resources--http_loadbalancer--reference--group-010.md#canonical-5c162cb70a77a04a40b0c7666a4b3f299a6a0815866bc29413b5748c87bdbd51)
- [api_testing.domains.credentials.standard](resources--http_loadbalancer--reference--group-010.md#canonical-022b9c587fad201b235813190250e458695f4c8fa89ebe6944a69960dbacee82)
- [api_testing.domains](resources--http_loadbalancer--reference--group-010.md#canonical-687b0526d20593ad003d8453ffecc5d3b2c352a7eaf26c5ec5b05ee3acc553ca)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-7a172f7eab2281c028bbf38c8461f2327fbb07d344cd5bf251c478dd01532e45"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-db1b42265ea0e3b1f34eb7d52c00fccf11f5475080637980dd3cbae8148f0dbb"></a>

## api_testing.domains.credentials.admin — api_testing.domains.credentials.admin / 8c9a22da23bb / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_testing](resources--http_loadbalancer--reference--group-010.md#canonical-699d08bedad7c6a3db855143b0dfa0f5040e8efa4d08cd58dccbf39afde96a71)
- [api_testing.domains](resources--http_loadbalancer--reference--group-010.md#canonical-687b0526d20593ad003d8453ffecc5d3b2c352a7eaf26c5ec5b05ee3acc553ca)
- [api_testing.domains.credentials](resources--http_loadbalancer--reference--group-010.md#canonical-b610213d02bd2bc48360d13b69ca5c914988db8ebbb077cfa1edf5a349059f36)
- api_testing.domains.credentials.admin

<a id="canonical-864b264d81a722118316c3f25db48fae1128eeec7201c382c7f6a0360b9a2e8f"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
admin = {}
```

<a id="canonical-626768a8d723d915d7473670ba4458bfef2df1f793e26f9230e5117265175a84"></a>

## Direct properties — api_testing.domains.credentials.admin / 8c9a22da23bb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7dc274ff8ecd8abd71c8595efc57d22793a2d83ba23f7b2487421fb2f60b5aeb"></a>

## Next pages — api_testing.domains.credentials.admin / 8c9a22da23bb / 4

- [api_testing.domains.credentials](resources--http_loadbalancer--reference--group-010.md#canonical-b610213d02bd2bc48360d13b69ca5c914988db8ebbb077cfa1edf5a349059f36)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-133753a73ec1db28898c1e436e98dde126bc7489d1ccacd34ea71506d4749b5c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-59f42efc692927cc59df21aba893025da6e5b9f3fc21518584557219b432106f"></a>

## api_testing.domains.credentials.api_key — api_testing.domains.credentials.api_key / 1fd5c5ae0e9c / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_testing](resources--http_loadbalancer--reference--group-010.md#canonical-699d08bedad7c6a3db855143b0dfa0f5040e8efa4d08cd58dccbf39afde96a71)
- [api_testing.domains](resources--http_loadbalancer--reference--group-010.md#canonical-687b0526d20593ad003d8453ffecc5d3b2c352a7eaf26c5ec5b05ee3acc553ca)
- [api_testing.domains.credentials](resources--http_loadbalancer--reference--group-010.md#canonical-b610213d02bd2bc48360d13b69ca5c914988db8ebbb077cfa1edf5a349059f36)
- api_testing.domains.credentials.api_key

<a id="canonical-bf44fb8c1c781a5152dfc734ee0c33573ab1af422f6b785206bfe7c33c0ea870"></a>

Type: `"object"`. single nested block, Optional.

API Key

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("key")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
api_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-abed3369862b087c85a965593952b480fa9a28187d784b968788058cdea14170"></a>

## Direct properties — api_testing.domains.credentials.api_key / 1fd5c5ae0e9c / 3

<a id="canonical-b22cfcedb8f2b380487abb81e576d305fd03360109722d5226ae59d68b8f0973"></a>

<a id="canonical-0ee82a23ef8c9c7328602c435cd36305ecec1b3928a95362c6af5ab751c56df3"></a>

## key property — api_testing.domains.credentials.api_key / 1fd5c5ae0e9c / 4

Type: `"string"`. Optional.

Key. Cryptographic key material

Upstream description:

Cryptographic key material

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

- [value](resources--http_loadbalancer--reference--group-010.md#canonical-bafb015e30697032add2d1280d2cd13ea316041dd664b652e3c523059edcaab9): complete subsection reference.

<a id="canonical-dbaa750f69c9c3c771e37bf5d36b36f27c4aa27772a74e9b9e01faf0ba65c034"></a>

## Next pages — api_testing.domains.credentials.api_key / 1fd5c5ae0e9c / 5

- [api_testing.domains.credentials.api_key.value](resources--http_loadbalancer--reference--group-010.md#canonical-bafb015e30697032add2d1280d2cd13ea316041dd664b652e3c523059edcaab9)
- [api_testing.domains.credentials](resources--http_loadbalancer--reference--group-010.md#canonical-b610213d02bd2bc48360d13b69ca5c914988db8ebbb077cfa1edf5a349059f36)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-bafb015e30697032add2d1280d2cd13ea316041dd664b652e3c523059edcaab9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3437712c5d13a47117065b441dde03c4f04494a704405a08a0b166e60243aa60"></a>

## api_testing.domains.credentials.api_key.value — api_testing.domains.credentials.api_key.value / 370fdd7b7a59 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_testing](resources--http_loadbalancer--reference--group-010.md#canonical-699d08bedad7c6a3db855143b0dfa0f5040e8efa4d08cd58dccbf39afde96a71)
- [api_testing.domains](resources--http_loadbalancer--reference--group-010.md#canonical-687b0526d20593ad003d8453ffecc5d3b2c352a7eaf26c5ec5b05ee3acc553ca)
- [api_testing.domains.credentials](resources--http_loadbalancer--reference--group-010.md#canonical-b610213d02bd2bc48360d13b69ca5c914988db8ebbb077cfa1edf5a349059f36)
- [api_testing.domains.credentials.api_key](resources--http_loadbalancer--reference--group-010.md#canonical-133753a73ec1db28898c1e436e98dde126bc7489d1ccacd34ea71506d4749b5c)
- api_testing.domains.credentials.api_key.value

<a id="canonical-476d048db78073dcf758fbc84b319bb0e1a659c592e9cb7f1f6037925cbd3d0c"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
value {
  # Configure direct properties listed below.
}
```

<a id="canonical-887d37258406b569c4cbaffa60ff818868858aa3bf735fa84ded332528acc780"></a>

## Direct properties — api_testing.domains.credentials.api_key.value / 370fdd7b7a59 / 3

- [blindfold_secret_info](resources--http_loadbalancer--reference--group-010.md#canonical-df8d7891a3afb2845e4a71cb7cdfd4a543145f3587a0735baf060b03849c5c76): complete subsection reference.

- [clear_secret_info](resources--http_loadbalancer--reference--group-010.md#canonical-f981b54f19b03767dab4f2c2cf23d2d8993c1db385b7bbc8941d4ad19f0f3aae): complete subsection reference.

<a id="canonical-dc3eb4d5b6719f0e11b1909521b076eb2ab69d81dd6c844eccb61384fa90f5e7"></a>

## Next pages — api_testing.domains.credentials.api_key.value / 370fdd7b7a59 / 4

- [api_testing.domains.credentials.api_key.value.blindfold_secret_info](resources--http_loadbalancer--reference--group-010.md#canonical-df8d7891a3afb2845e4a71cb7cdfd4a543145f3587a0735baf060b03849c5c76)
- [api_testing.domains.credentials.api_key.value.clear_secret_info](resources--http_loadbalancer--reference--group-010.md#canonical-f981b54f19b03767dab4f2c2cf23d2d8993c1db385b7bbc8941d4ad19f0f3aae)
- [api_testing.domains.credentials.api_key](resources--http_loadbalancer--reference--group-010.md#canonical-133753a73ec1db28898c1e436e98dde126bc7489d1ccacd34ea71506d4749b5c)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-df8d7891a3afb2845e4a71cb7cdfd4a543145f3587a0735baf060b03849c5c76"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-02c10e06cc957956485b0cc546414b759644f61bb6ff53285adf7476cd521a56"></a>

## api_testing.domains.credentials.api_key.value.blindfold_secret_info — api_testing.domains.credentials.api_key.value.blindfold_secret_info / acfe4066b8e3 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_testing](resources--http_loadbalancer--reference--group-010.md#canonical-699d08bedad7c6a3db855143b0dfa0f5040e8efa4d08cd58dccbf39afde96a71)
- [api_testing.domains](resources--http_loadbalancer--reference--group-010.md#canonical-687b0526d20593ad003d8453ffecc5d3b2c352a7eaf26c5ec5b05ee3acc553ca)
- [api_testing.domains.credentials](resources--http_loadbalancer--reference--group-010.md#canonical-b610213d02bd2bc48360d13b69ca5c914988db8ebbb077cfa1edf5a349059f36)
- [api_testing.domains.credentials.api_key](resources--http_loadbalancer--reference--group-010.md#canonical-133753a73ec1db28898c1e436e98dde126bc7489d1ccacd34ea71506d4749b5c)
- [api_testing.domains.credentials.api_key.value](resources--http_loadbalancer--reference--group-010.md#canonical-bafb015e30697032add2d1280d2cd13ea316041dd664b652e3c523059edcaab9)
- api_testing.domains.credentials.api_key.value.blindfold_secret_info

<a id="canonical-79bcb3ef40d34957a32613ec5e949013bc78685250176f48b17ba88cfb660dc5"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("location")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-f7f1bca840264ce868ba38b95f0442f5e64130fa31ce472238f8fb0356840e87"></a>

## Direct properties — api_testing.domains.credentials.api_key.value.blindfold_secret_info / acfe4066b8e3 / 3

<a id="canonical-0dfe39fd36ec392b7c53484b6979ef1409b1df07d3402364360b8703a1109b53"></a>

<a id="canonical-e81a4073b8cb129c165534f1841049c6c708ce77c240ea5aca4805876708a2c1"></a>

## decryption_provider property — api_testing.domains.credentials.api_key.value.blindfold_secret_info / acfe4066b8e3 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the backend Secret
Management service.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-741c6beb7e6e47cafae26512ee683a92122362572306f64956877ad80acf23e5"></a>

<a id="canonical-978b8fbb88a615e41068095747f5bfd869578f4ac16288c689ba7e603e8003c0"></a>

## location property — api_testing.domains.credentials.api_key.value.blindfold_secret_info / acfe4066b8e3 / 5

Type: `"string"`. Optional, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
    "metadata": {
      "category": "content",
      "confidence": 1.0,
      "note": "Blindfold envelope encryption (AES-256-GCM + RSA-OAEP) of an RSA-2048 TLS private key produces ~3700 char string:/// URL. 128KB max secret size = ~175KB base64. Discovery reported 1024 which is incorrect.",
      "source": "manual-override",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 4
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-65308ccaad6768ca13e9bf77aca2386c31b836722676917c839611b2babb7ef1"></a>

<a id="canonical-94e2371ca6d889618c4aa8685292e0ea4b88beaa647a521402c04cdb9e298eec"></a>

## store_provider property — api_testing.domains.credentials.api_key.value.blindfold_secret_info / acfe4066b8e3 / 6

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-f90a6d224bf08dd987e526ae506f1bccbb2d238f7693711ecf4542d9e1fc339e"></a>

## Next pages — api_testing.domains.credentials.api_key.value.blindfold_secret_info / acfe4066b8e3 / 7

- [api_testing.domains.credentials.api_key.value](resources--http_loadbalancer--reference--group-010.md#canonical-bafb015e30697032add2d1280d2cd13ea316041dd664b652e3c523059edcaab9)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-f981b54f19b03767dab4f2c2cf23d2d8993c1db385b7bbc8941d4ad19f0f3aae"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-de094b96b7532a2e1c4dd08d6d9b3553f6934166b434ee0d51d8ab2d1864252a"></a>

## api_testing.domains.credentials.api_key.value.clear_secret_info — api_testing.domains.credentials.api_key.value.clear_secret_info / 9f3c04fbfcae / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_testing](resources--http_loadbalancer--reference--group-010.md#canonical-699d08bedad7c6a3db855143b0dfa0f5040e8efa4d08cd58dccbf39afde96a71)
- [api_testing.domains](resources--http_loadbalancer--reference--group-010.md#canonical-687b0526d20593ad003d8453ffecc5d3b2c352a7eaf26c5ec5b05ee3acc553ca)
- [api_testing.domains.credentials](resources--http_loadbalancer--reference--group-010.md#canonical-b610213d02bd2bc48360d13b69ca5c914988db8ebbb077cfa1edf5a349059f36)
- [api_testing.domains.credentials.api_key](resources--http_loadbalancer--reference--group-010.md#canonical-133753a73ec1db28898c1e436e98dde126bc7489d1ccacd34ea71506d4749b5c)
- [api_testing.domains.credentials.api_key.value](resources--http_loadbalancer--reference--group-010.md#canonical-bafb015e30697032add2d1280d2cd13ea316041dd664b652e3c523059edcaab9)
- api_testing.domains.credentials.api_key.value.clear_secret_info

<a id="canonical-27a977aab0e6504c782738e45668bfbd77e5fdeb0d53547ec78066aec835c829"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-9a24286c646758329ed5f18d8d1691e00b5579f8c9efd46b6e78fe7b71f429fb"></a>

## Direct properties — api_testing.domains.credentials.api_key.value.clear_secret_info / 9f3c04fbfcae / 3

<a id="canonical-4aa73b3d85b09f9c1d07d0dff207419fff44f4fa358789f2e373d679aae2d041"></a>

<a id="canonical-a2be10f7fbdabb22e2343ec4b66051998166c30c9027b0747d59cc2db28f555a"></a>

## provider_ref property — api_testing.domains.credentials.api_key.value.clear_secret_info / 9f3c04fbfcae / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-697ee3ee232267fa7e301f4e8271939b89a39592fdae63b6380f851632680451"></a>

<a id="canonical-31a89af29b233065193c01a106b35b6a3aa11567e2ed427c8105072ed572d920"></a>

## url property — api_testing.domains.credentials.api_key.value.clear_secret_info / 9f3c04fbfcae / 5

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-0117556c40139fd2d5e4b9c8f3a1edfa9e2ca4dea89e24f8eef062251148647b"></a>

## Next pages — api_testing.domains.credentials.api_key.value.clear_secret_info / 9f3c04fbfcae / 6

- [api_testing.domains.credentials.api_key.value](resources--http_loadbalancer--reference--group-010.md#canonical-bafb015e30697032add2d1280d2cd13ea316041dd664b652e3c523059edcaab9)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-ba6ae20ae6e213b12c2649468bd7833395743cd8b07901eb48f1673fe04c0dd8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7e54fc326088a9a3e235c9c824d4f258a3cbbfe18c1c115a0502f444673529ae"></a>

## api_testing.domains.credentials.basic_auth — api_testing.domains.credentials.basic_auth / 9a5fec21a74f / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_testing](resources--http_loadbalancer--reference--group-010.md#canonical-699d08bedad7c6a3db855143b0dfa0f5040e8efa4d08cd58dccbf39afde96a71)
- [api_testing.domains](resources--http_loadbalancer--reference--group-010.md#canonical-687b0526d20593ad003d8453ffecc5d3b2c352a7eaf26c5ec5b05ee3acc553ca)
- [api_testing.domains.credentials](resources--http_loadbalancer--reference--group-010.md#canonical-b610213d02bd2bc48360d13b69ca5c914988db8ebbb077cfa1edf5a349059f36)
- api_testing.domains.credentials.basic_auth

<a id="canonical-e9b2e42643a2629489f7fc8959f3d7f1f9bb8a23901a413d1c83e100ee9b1661"></a>

Type: `"object"`. single nested block, Optional.

Basic Authentication.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("user")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
basic_auth {
  # Configure direct properties listed below.
}
```

<a id="canonical-5137bb55cc1a6c77c8d42818b7ba27ea3c61a7af31d534cedcb0b023f7af4fdc"></a>

## Direct properties — api_testing.domains.credentials.basic_auth / 9a5fec21a74f / 3

- [password](resources--http_loadbalancer--reference--group-010.md#canonical-72bc8de3bc123a167d4d173e0522602a12fde966bd4bc7953d3a18bc3dc70931): complete subsection reference.

<a id="canonical-505880ae3e49950c4d9016334377ea0d062f11222078636786f20fc374e50c65"></a>

<a id="canonical-1513e0b8bd8d4cb52125e64e0ee2f38a7420fb8648fbfa0ccfd21c539bd509e0"></a>

## user property — api_testing.domains.credentials.basic_auth / 9a5fec21a74f / 4

Type: `"string"`. Optional.

User. Configuration parameter for user

Upstream description:

Configuration parameter for user

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-zA-Z0-9_.-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-55cf0ef8c02089059efa23c98836e992680f2e001a5dc00170c1eb3b95b23c38"></a>

## Next pages — api_testing.domains.credentials.basic_auth / 9a5fec21a74f / 5

- [api_testing.domains.credentials.basic_auth.password](resources--http_loadbalancer--reference--group-010.md#canonical-72bc8de3bc123a167d4d173e0522602a12fde966bd4bc7953d3a18bc3dc70931)
- [api_testing.domains.credentials](resources--http_loadbalancer--reference--group-010.md#canonical-b610213d02bd2bc48360d13b69ca5c914988db8ebbb077cfa1edf5a349059f36)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-72bc8de3bc123a167d4d173e0522602a12fde966bd4bc7953d3a18bc3dc70931"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0e63b4e3df1c579027a289a237276191bccaa81a4042f70410be2ac1b0d67a2a"></a>

## api_testing.domains.credentials.basic_auth.password — api_testing.domains.credentials.basic_auth.password / 6945805b0b95 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_testing](resources--http_loadbalancer--reference--group-010.md#canonical-699d08bedad7c6a3db855143b0dfa0f5040e8efa4d08cd58dccbf39afde96a71)
- [api_testing.domains](resources--http_loadbalancer--reference--group-010.md#canonical-687b0526d20593ad003d8453ffecc5d3b2c352a7eaf26c5ec5b05ee3acc553ca)
- [api_testing.domains.credentials](resources--http_loadbalancer--reference--group-010.md#canonical-b610213d02bd2bc48360d13b69ca5c914988db8ebbb077cfa1edf5a349059f36)
- [api_testing.domains.credentials.basic_auth](resources--http_loadbalancer--reference--group-010.md#canonical-ba6ae20ae6e213b12c2649468bd7833395743cd8b07901eb48f1673fe04c0dd8)
- api_testing.domains.credentials.basic_auth.password

<a id="canonical-52e8f49054f87b7c84044455fd6fccde2c4f7e83a7af04826018f9bae4f953f0"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
password {
  # Configure direct properties listed below.
}
```

<a id="canonical-411959253b298e9da6f5e5f4213301f1c7c51859ac49b9055c67f8e4db279f44"></a>

## Direct properties — api_testing.domains.credentials.basic_auth.password / 6945805b0b95 / 3

- [blindfold_secret_info](resources--http_loadbalancer--reference--group-010.md#canonical-61a13545f4416072fe5d2bd69d3c61c90537680cef53a7753291d3cd980149cf): complete subsection reference.

- [clear_secret_info](resources--http_loadbalancer--reference--group-010.md#canonical-2b2e207f6100788f249c89fb84077d533a64d420e7a41ef7ef843e7e902e8294): complete subsection reference.

<a id="canonical-80f9679e8b244b802a696515c5d8b43b570e105532b3f8e8ff5fd535286b03ce"></a>

## Next pages — api_testing.domains.credentials.basic_auth.password / 6945805b0b95 / 4

- [api_testing.domains.credentials.basic_auth.password.blindfold_secret_info](resources--http_loadbalancer--reference--group-010.md#canonical-61a13545f4416072fe5d2bd69d3c61c90537680cef53a7753291d3cd980149cf)
- [api_testing.domains.credentials.basic_auth.password.clear_secret_info](resources--http_loadbalancer--reference--group-010.md#canonical-2b2e207f6100788f249c89fb84077d533a64d420e7a41ef7ef843e7e902e8294)
- [api_testing.domains.credentials.basic_auth](resources--http_loadbalancer--reference--group-010.md#canonical-ba6ae20ae6e213b12c2649468bd7833395743cd8b07901eb48f1673fe04c0dd8)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-61a13545f4416072fe5d2bd69d3c61c90537680cef53a7753291d3cd980149cf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b5e7fc1062edcfcd7b01068cff8880e363d13807b66eb11612b2fb98953317f8"></a>

## api_testing.domains.credentials.basic_auth.password.blindfold_secret_info — api_testing.domains.credentials.basic_auth.password.blindfold_secret_info / 42dbf1410b29 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_testing](resources--http_loadbalancer--reference--group-010.md#canonical-699d08bedad7c6a3db855143b0dfa0f5040e8efa4d08cd58dccbf39afde96a71)
- [api_testing.domains](resources--http_loadbalancer--reference--group-010.md#canonical-687b0526d20593ad003d8453ffecc5d3b2c352a7eaf26c5ec5b05ee3acc553ca)
- [api_testing.domains.credentials](resources--http_loadbalancer--reference--group-010.md#canonical-b610213d02bd2bc48360d13b69ca5c914988db8ebbb077cfa1edf5a349059f36)
- [api_testing.domains.credentials.basic_auth](resources--http_loadbalancer--reference--group-010.md#canonical-ba6ae20ae6e213b12c2649468bd7833395743cd8b07901eb48f1673fe04c0dd8)
- [api_testing.domains.credentials.basic_auth.password](resources--http_loadbalancer--reference--group-010.md#canonical-72bc8de3bc123a167d4d173e0522602a12fde966bd4bc7953d3a18bc3dc70931)
- api_testing.domains.credentials.basic_auth.password.blindfold_secret_info

<a id="canonical-35f7afbed9232c007f7b452d143d4bbbb8aa532d3e91c0d416876b8f48609093"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("location")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-ea3773ff2b9f9090cfd999b2fdbfff9e252811d969cf7915b762461b0ca070b1"></a>

## Direct properties — api_testing.domains.credentials.basic_auth.password.blindfold_secret_info / 42dbf1410b29 / 3

<a id="canonical-e902248e7061a2bc18dbf956392c0bfd584157b0cb2e543514d534d675661191"></a>

<a id="canonical-ed35a0f890c6fcfc8226fe011787d3e3fdadec3b6e904e1a76ea8a188e4384bf"></a>

## decryption_provider property — api_testing.domains.credentials.basic_auth.password.blindfold_secret_info / 42dbf1410b29 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the backend Secret
Management service.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-25b8041882ad8e903d5568bdc94010ccceb359c22d8521b3ece60423bb3f024a"></a>

<a id="canonical-0f6dd8deb663bb97095d7e70b90de355538ac0b96a944e4acc42780acaa248e7"></a>

## location property — api_testing.domains.credentials.basic_auth.password.blindfold_secret_info / 42dbf1410b29 / 5

Type: `"string"`. Optional, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
    "metadata": {
      "category": "content",
      "confidence": 1.0,
      "note": "Blindfold envelope encryption (AES-256-GCM + RSA-OAEP) of an RSA-2048 TLS private key produces ~3700 char string:/// URL. 128KB max secret size = ~175KB base64. Discovery reported 1024 which is incorrect.",
      "source": "manual-override",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 4
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-dc070a7fc40b6d4c95469d51143c082933308e8810149842f7dc8edcdbce4e13"></a>

<a id="canonical-39383b0e6e776418ef28083df19ef3c04c9c2a7f6f7b1c89949ede39690d72cc"></a>

## store_provider property — api_testing.domains.credentials.basic_auth.password.blindfold_secret_info / 42dbf1410b29 / 6

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-4200609886053cd8fef7e538dcde669dc6a1665d8f9360e6ce697e8d7a0514c1"></a>

## Next pages — api_testing.domains.credentials.basic_auth.password.blindfold_secret_info / 42dbf1410b29 / 7

- [api_testing.domains.credentials.basic_auth.password](resources--http_loadbalancer--reference--group-010.md#canonical-72bc8de3bc123a167d4d173e0522602a12fde966bd4bc7953d3a18bc3dc70931)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-2b2e207f6100788f249c89fb84077d533a64d420e7a41ef7ef843e7e902e8294"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f5094fa78b55c3d2447af7cd80323b99b8b29908828b73c80ec33961fa13f749"></a>

## api_testing.domains.credentials.basic_auth.password.clear_secret_info — api_testing.domains.credentials.basic_auth.password.clear_secret_info / 81a25c0e7e02 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_testing](resources--http_loadbalancer--reference--group-010.md#canonical-699d08bedad7c6a3db855143b0dfa0f5040e8efa4d08cd58dccbf39afde96a71)
- [api_testing.domains](resources--http_loadbalancer--reference--group-010.md#canonical-687b0526d20593ad003d8453ffecc5d3b2c352a7eaf26c5ec5b05ee3acc553ca)
- [api_testing.domains.credentials](resources--http_loadbalancer--reference--group-010.md#canonical-b610213d02bd2bc48360d13b69ca5c914988db8ebbb077cfa1edf5a349059f36)
- [api_testing.domains.credentials.basic_auth](resources--http_loadbalancer--reference--group-010.md#canonical-ba6ae20ae6e213b12c2649468bd7833395743cd8b07901eb48f1673fe04c0dd8)
- [api_testing.domains.credentials.basic_auth.password](resources--http_loadbalancer--reference--group-010.md#canonical-72bc8de3bc123a167d4d173e0522602a12fde966bd4bc7953d3a18bc3dc70931)
- api_testing.domains.credentials.basic_auth.password.clear_secret_info

<a id="canonical-1c3b112f18ad4449d933524ca5d82c535aa8fbfab021c4262de9988dba029acf"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-1ecf8693b3700ca2fd79de8c67a53c2795703acbf3764e3275e7e98e6c86e217"></a>

## Direct properties — api_testing.domains.credentials.basic_auth.password.clear_secret_info / 81a25c0e7e02 / 3

<a id="canonical-ea7532e752e4fc98d3a0990a9afe9737f996f1c6f2b760439b43b24cafddca7b"></a>

<a id="canonical-e0fb51063809704425c691b7703c8acd91384171a514d9001e2a0f6c0e8b1074"></a>

## provider_ref property — api_testing.domains.credentials.basic_auth.password.clear_secret_info / 81a25c0e7e02 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-a357752a5a326f993274a6a0eb9b6f38d83402669873b779a0f79dc9b482f97d"></a>

<a id="canonical-5c2e4bfec659d5bc42b3dd533ccebabf08ffa689ae0dae21443464645fc30d1d"></a>

## url property — api_testing.domains.credentials.basic_auth.password.clear_secret_info / 81a25c0e7e02 / 5

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-3be8583eca998d2720778ba871895f5802835283622a3fbcc945ee9dc9d6c4d3"></a>

## Next pages — api_testing.domains.credentials.basic_auth.password.clear_secret_info / 81a25c0e7e02 / 6

- [api_testing.domains.credentials.basic_auth.password](resources--http_loadbalancer--reference--group-010.md#canonical-72bc8de3bc123a167d4d173e0522602a12fde966bd4bc7953d3a18bc3dc70931)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-6633e13b98f83d6d632c1bc8b169ea978c0c926699633f813b8e5517d3705b5c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5d16de27c4e90997a598d125d2e13496ba6e1448ce29fb545f7b433156369d11"></a>

## api_testing.domains.credentials.bearer_token — api_testing.domains.credentials.bearer_token / d3aefafe4ded / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_testing](resources--http_loadbalancer--reference--group-010.md#canonical-699d08bedad7c6a3db855143b0dfa0f5040e8efa4d08cd58dccbf39afde96a71)
- [api_testing.domains](resources--http_loadbalancer--reference--group-010.md#canonical-687b0526d20593ad003d8453ffecc5d3b2c352a7eaf26c5ec5b05ee3acc553ca)
- [api_testing.domains.credentials](resources--http_loadbalancer--reference--group-010.md#canonical-b610213d02bd2bc48360d13b69ca5c914988db8ebbb077cfa1edf5a349059f36)
- api_testing.domains.credentials.bearer_token

<a id="canonical-cf77db0c8744dd344db6d9819eea787bac2eae549e159bce2941cf7afba30024"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for bearer token.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
bearer_token {
  # Configure direct properties listed below.
}
```

<a id="canonical-58ec95f833021f7389269d62524becf7254bf9175eee6ff01ddd3cb7d0801953"></a>

## Direct properties — api_testing.domains.credentials.bearer_token / d3aefafe4ded / 3

- [token](resources--http_loadbalancer--reference--group-010.md#canonical-eb8ce56cbea7f7c506d4f7ebc435b71a95885bee33c5adefc86ca300e88f4f12): complete subsection reference.

<a id="canonical-bb00c154b0c06d8bd5769ec99130cf938bff6c019d4e34a6d6cccdec37943b88"></a>

## Next pages — api_testing.domains.credentials.bearer_token / d3aefafe4ded / 4

- [api_testing.domains.credentials.bearer_token.token](resources--http_loadbalancer--reference--group-010.md#canonical-eb8ce56cbea7f7c506d4f7ebc435b71a95885bee33c5adefc86ca300e88f4f12)
- [api_testing.domains.credentials](resources--http_loadbalancer--reference--group-010.md#canonical-b610213d02bd2bc48360d13b69ca5c914988db8ebbb077cfa1edf5a349059f36)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-eb8ce56cbea7f7c506d4f7ebc435b71a95885bee33c5adefc86ca300e88f4f12"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-42068a06ff1a01e89fa6a434d024e9a7c26b98b7013b4fab8c0fd587466d30f2"></a>

## api_testing.domains.credentials.bearer_token.token — api_testing.domains.credentials.bearer_token.token / 10669fe961fd / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_testing](resources--http_loadbalancer--reference--group-010.md#canonical-699d08bedad7c6a3db855143b0dfa0f5040e8efa4d08cd58dccbf39afde96a71)
- [api_testing.domains](resources--http_loadbalancer--reference--group-010.md#canonical-687b0526d20593ad003d8453ffecc5d3b2c352a7eaf26c5ec5b05ee3acc553ca)
- [api_testing.domains.credentials](resources--http_loadbalancer--reference--group-010.md#canonical-b610213d02bd2bc48360d13b69ca5c914988db8ebbb077cfa1edf5a349059f36)
- [api_testing.domains.credentials.bearer_token](resources--http_loadbalancer--reference--group-010.md#canonical-6633e13b98f83d6d632c1bc8b169ea978c0c926699633f813b8e5517d3705b5c)
- api_testing.domains.credentials.bearer_token.token

<a id="canonical-3fc9a0b3e9613f5f594fd529e07f7e0de9431fa8458318d743edb342c0855540"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
token {
  # Configure direct properties listed below.
}
```

<a id="canonical-0785bbe565fe42427ddad426401e6962465b5816950a5ccd66bba33882462d05"></a>

## Direct properties — api_testing.domains.credentials.bearer_token.token / 10669fe961fd / 3

- [blindfold_secret_info](resources--http_loadbalancer--reference--group-010.md#canonical-6c6eb02335a0acbb496810530ffcdb873f489419a1b4f5d9c95fc3e985fc733d): complete subsection reference.

- [clear_secret_info](resources--http_loadbalancer--reference--group-010.md#canonical-8b821552e3d665ac6c82acabf98ef560017fbb9824548f36150e5fa81e7a67d8): complete subsection reference.

<a id="canonical-65dd56957c8d3508c565c3dc11b96cc5cefc1c5a987d0af02dd382d2422329e2"></a>

## Next pages — api_testing.domains.credentials.bearer_token.token / 10669fe961fd / 4

- [api_testing.domains.credentials.bearer_token.token.blindfold_secret_info](resources--http_loadbalancer--reference--group-010.md#canonical-6c6eb02335a0acbb496810530ffcdb873f489419a1b4f5d9c95fc3e985fc733d)
- [api_testing.domains.credentials.bearer_token.token.clear_secret_info](resources--http_loadbalancer--reference--group-010.md#canonical-8b821552e3d665ac6c82acabf98ef560017fbb9824548f36150e5fa81e7a67d8)
- [api_testing.domains.credentials.bearer_token](resources--http_loadbalancer--reference--group-010.md#canonical-6633e13b98f83d6d632c1bc8b169ea978c0c926699633f813b8e5517d3705b5c)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-6c6eb02335a0acbb496810530ffcdb873f489419a1b4f5d9c95fc3e985fc733d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a0c1caa9c73313159bdb4c164aee9312ec3cc3857ee7f9eac0bf5ce3c7e5812a"></a>

## api_testing.domains.credentials.bearer_token.token.blindfold_secret_info — api_testing.domains.credentials.bearer_token.token.blindfold_secret_info / 3de798d920fe / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_testing](resources--http_loadbalancer--reference--group-010.md#canonical-699d08bedad7c6a3db855143b0dfa0f5040e8efa4d08cd58dccbf39afde96a71)
- [api_testing.domains](resources--http_loadbalancer--reference--group-010.md#canonical-687b0526d20593ad003d8453ffecc5d3b2c352a7eaf26c5ec5b05ee3acc553ca)
- [api_testing.domains.credentials](resources--http_loadbalancer--reference--group-010.md#canonical-b610213d02bd2bc48360d13b69ca5c914988db8ebbb077cfa1edf5a349059f36)
- [api_testing.domains.credentials.bearer_token](resources--http_loadbalancer--reference--group-010.md#canonical-6633e13b98f83d6d632c1bc8b169ea978c0c926699633f813b8e5517d3705b5c)
- [api_testing.domains.credentials.bearer_token.token](resources--http_loadbalancer--reference--group-010.md#canonical-eb8ce56cbea7f7c506d4f7ebc435b71a95885bee33c5adefc86ca300e88f4f12)
- api_testing.domains.credentials.bearer_token.token.blindfold_secret_info

<a id="canonical-5a22db33ae49247fb0d6dfec754fb72354c210aba74612165074dc612f3e3a90"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("location")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-e69b302a834efb4856bd2b36750acff79684e881fd9984c4b4e3efc1840cc9d9"></a>

## Direct properties — api_testing.domains.credentials.bearer_token.token.blindfold_secret_info / 3de798d920fe / 3

<a id="canonical-67463f391387b88b05ea61828389405ecde49a453f71df52a547bd2349c37fbb"></a>

<a id="canonical-45f802ea681f42bfbf85270753aa430b3a145e9636aa0fb1b793fa1068f505d8"></a>

## decryption_provider property — api_testing.domains.credentials.bearer_token.token.blindfold_secret_info / 3de798d920fe / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the backend Secret
Management service.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-a8788bfe4a499639005eea7e3af6c6895b0f9e2d87da4d8592d339ff41842f17"></a>

<a id="canonical-f555e9f3ccf15ed4f704c0049de6523fef18cbeab7fa2555f79aa7ceeaf19ee9"></a>

## location property — api_testing.domains.credentials.bearer_token.token.blindfold_secret_info / 3de798d920fe / 5

Type: `"string"`. Optional, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
    "metadata": {
      "category": "content",
      "confidence": 1.0,
      "note": "Blindfold envelope encryption (AES-256-GCM + RSA-OAEP) of an RSA-2048 TLS private key produces ~3700 char string:/// URL. 128KB max secret size = ~175KB base64. Discovery reported 1024 which is incorrect.",
      "source": "manual-override",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 4
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-cf8a9c9ed23b0d197a7b10790f6ee62dda025fcaeda8c54c46bb938628d941ef"></a>

<a id="canonical-1f0d652a28efa6f37ee40acfd3abaf9e73f8a1a2d043a53b01f844c9e5651e7f"></a>

## store_provider property — api_testing.domains.credentials.bearer_token.token.blindfold_secret_info / 3de798d920fe / 6

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-9cecf9d03ca8909095c086d3cf57d1fd56fe9236ea8ad3cbf1fda8c5d3143a4d"></a>

## Next pages — api_testing.domains.credentials.bearer_token.token.blindfold_secret_info / 3de798d920fe / 7

- [api_testing.domains.credentials.bearer_token.token](resources--http_loadbalancer--reference--group-010.md#canonical-eb8ce56cbea7f7c506d4f7ebc435b71a95885bee33c5adefc86ca300e88f4f12)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-8b821552e3d665ac6c82acabf98ef560017fbb9824548f36150e5fa81e7a67d8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-59e71724a11ff61347a551b8140e3d3e2da2bacb1159889e4f1031ac29b17e4e"></a>

## api_testing.domains.credentials.bearer_token.token.clear_secret_info — api_testing.domains.credentials.bearer_token.token.clear_secret_info / 094559bdb911 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_testing](resources--http_loadbalancer--reference--group-010.md#canonical-699d08bedad7c6a3db855143b0dfa0f5040e8efa4d08cd58dccbf39afde96a71)
- [api_testing.domains](resources--http_loadbalancer--reference--group-010.md#canonical-687b0526d20593ad003d8453ffecc5d3b2c352a7eaf26c5ec5b05ee3acc553ca)
- [api_testing.domains.credentials](resources--http_loadbalancer--reference--group-010.md#canonical-b610213d02bd2bc48360d13b69ca5c914988db8ebbb077cfa1edf5a349059f36)
- [api_testing.domains.credentials.bearer_token](resources--http_loadbalancer--reference--group-010.md#canonical-6633e13b98f83d6d632c1bc8b169ea978c0c926699633f813b8e5517d3705b5c)
- [api_testing.domains.credentials.bearer_token.token](resources--http_loadbalancer--reference--group-010.md#canonical-eb8ce56cbea7f7c506d4f7ebc435b71a95885bee33c5adefc86ca300e88f4f12)
- api_testing.domains.credentials.bearer_token.token.clear_secret_info

<a id="canonical-1f2c604479dc0dddf27c99a2c0bb6720c9e8b8c0b01a2c95fb1920163353ab0e"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-d4a25137498d9111d693ab0bd1275626b7986b7e02418cc3db2e1434a9c33659"></a>

## Direct properties — api_testing.domains.credentials.bearer_token.token.clear_secret_info / 094559bdb911 / 3

<a id="canonical-4b8eaa28d07d3e5746099400f8beb49ed658dbd24420a74b76c908a3b728d36a"></a>

<a id="canonical-d8932f39c7b31edc0f0d6c4d6fe8cc9835cacbbb2d48f25472c2e704e92db0e9"></a>

## provider_ref property — api_testing.domains.credentials.bearer_token.token.clear_secret_info / 094559bdb911 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-9a4552a6e6a2ea339fb38fb893da3e37d150720385e4808be7816453f706c91b"></a>

<a id="canonical-74bfe1a324f9169d17583a31f148e45e2b85a1a44f1b2e0c9b3f475e0189b30e"></a>

## url property — api_testing.domains.credentials.bearer_token.token.clear_secret_info / 094559bdb911 / 5

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-82392ead7fb1defcdf640cf63399b5e60d6f754c69462df0e9dba1a30fa6284c"></a>

## Next pages — api_testing.domains.credentials.bearer_token.token.clear_secret_info / 094559bdb911 / 6

- [api_testing.domains.credentials.bearer_token.token](resources--http_loadbalancer--reference--group-010.md#canonical-eb8ce56cbea7f7c506d4f7ebc435b71a95885bee33c5adefc86ca300e88f4f12)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-5c162cb70a77a04a40b0c7666a4b3f299a6a0815866bc29413b5748c87bdbd51"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-604288c900d7ac6f668cbc9b2b4957fed1ed644bc8ff7066efa7f6f4f79fa70a"></a>

## api_testing.domains.credentials.login_endpoint — api_testing.domains.credentials.login_endpoint / eb31c141dc79 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_testing](resources--http_loadbalancer--reference--group-010.md#canonical-699d08bedad7c6a3db855143b0dfa0f5040e8efa4d08cd58dccbf39afde96a71)
- [api_testing.domains](resources--http_loadbalancer--reference--group-010.md#canonical-687b0526d20593ad003d8453ffecc5d3b2c352a7eaf26c5ec5b05ee3acc553ca)
- [api_testing.domains.credentials](resources--http_loadbalancer--reference--group-010.md#canonical-b610213d02bd2bc48360d13b69ca5c914988db8ebbb077cfa1edf5a349059f36)
- api_testing.domains.credentials.login_endpoint

<a id="canonical-6e143a43e983316de6237a17e313a156a1f62f41c5c00b5e46ba017b657bdd8e"></a>

Type: `"object"`. single nested block, Optional.

Login Endpoint.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("path",
    "token_response_key")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
login_endpoint {
  # Configure direct properties listed below.
}
```

<a id="canonical-7a028eb18e4bff56b4f5c842af4fcf880ce46bd93cfa6c5efffaa69a92f3f7fe"></a>

## Direct properties — api_testing.domains.credentials.login_endpoint / eb31c141dc79 / 3

- [json_payload](resources--http_loadbalancer--reference--group-010.md#canonical-e971da389b0bf5e5e9c2c2448f15488c668ba80da1e442cf06e1646726a355b3): complete subsection reference.

<a id="canonical-ebfdea445bcbd061b3e03551c3585136e00728790203d20ccbba2eeddb30b480"></a>

<a id="canonical-9a2589e9bb1c91e65c68359a4d97f01359d71cfaf533b760a93f0f32cb959c81"></a>

## method property — api_testing.domains.credentials.login_endpoint / eb31c141dc79 / 4

Type: `"string"`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Specifies the HTTP method
used to access a resource. Any HTTP Method. Possible values are \`ANY\`, \`GET\`, \`HEAD\`,
\`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to
\`ANY\`.

Upstream description:

Specifies the HTTP method used to access a resource.

Any HTTP Method.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ANY",
    "GET",
    "HEAD",
    "POST",
    "PUT",
    "DELETE",
    "CONNECT",
    "OPTIONS",
    "TRACE",
    "PATCH",
    "COPY"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "ANY",
  "enum": [
    "ANY",
    "GET",
    "HEAD",
    "POST",
    "PUT",
    "DELETE",
    "CONNECT",
    "OPTIONS",
    "TRACE",
    "PATCH",
    "COPY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-fe5339efb36c5d08fc73633f3fc433f1d9abc9d34eeab51abf2346c691d87a66"></a>

<a id="canonical-9a803d687c611eb8141bd9a375369a138dbf8d3dd96a6e2a8a6128011923ec68"></a>

## path property — api_testing.domains.credentials.login_endpoint / eb31c141dc79 / 5

Type: `"string"`. Optional.

Path. URL path for the endpoint

Upstream description:

URL path for the endpoint

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "1024"
  }
}
```

<a id="canonical-94e9d6b86962dd780c7db76b8d80f1b5fdf8a5fbbe33773c572397b85f344c87"></a>

<a id="canonical-5d22c4a100a8f976f1b54a6a54f6cec70ec3a7d91f85f06bf3f0c30a67227fd1"></a>

## token_response_key property — api_testing.domains.credentials.login_endpoint / eb31c141dc79 / 6

Type: `"string"`. Optional.

Specifies the key name used to extract the authentication token from the login response, such as
token or access\_token.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-49644d2ec58d0cfb2345da6794ef7dbabb503160e96478b78e60b10b799a0fec"></a>

## Next pages — api_testing.domains.credentials.login_endpoint / eb31c141dc79 / 7

- [api_testing.domains.credentials.login_endpoint.json_payload](resources--http_loadbalancer--reference--group-010.md#canonical-e971da389b0bf5e5e9c2c2448f15488c668ba80da1e442cf06e1646726a355b3)
- [api_testing.domains.credentials](resources--http_loadbalancer--reference--group-010.md#canonical-b610213d02bd2bc48360d13b69ca5c914988db8ebbb077cfa1edf5a349059f36)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-e971da389b0bf5e5e9c2c2448f15488c668ba80da1e442cf06e1646726a355b3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-112ce53bc13200509c51e7e681ac9111e2a7ec502e0dec0f0d5d57ea3e5df221"></a>

## api_testing.domains.credentials.login_endpoint.json_payload — api_testing.domains.credentials.login_endpoint.json_payload / 762d6884307d / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_testing](resources--http_loadbalancer--reference--group-010.md#canonical-699d08bedad7c6a3db855143b0dfa0f5040e8efa4d08cd58dccbf39afde96a71)
- [api_testing.domains](resources--http_loadbalancer--reference--group-010.md#canonical-687b0526d20593ad003d8453ffecc5d3b2c352a7eaf26c5ec5b05ee3acc553ca)
- [api_testing.domains.credentials](resources--http_loadbalancer--reference--group-010.md#canonical-b610213d02bd2bc48360d13b69ca5c914988db8ebbb077cfa1edf5a349059f36)
- [api_testing.domains.credentials.login_endpoint](resources--http_loadbalancer--reference--group-010.md#canonical-5c162cb70a77a04a40b0c7666a4b3f299a6a0815866bc29413b5748c87bdbd51)
- api_testing.domains.credentials.login_endpoint.json_payload

<a id="canonical-b35dad35b926cf19de1f1f0bc4af8155d3f37aa854312c4c75cfe1ec018ac060"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
json_payload {
  # Configure direct properties listed below.
}
```

<a id="canonical-a3b37480ec0e21fc2b182ddf4dcbe903b2484b13cb4586c913977604b3d143c2"></a>

## Direct properties — api_testing.domains.credentials.login_endpoint.json_payload / 762d6884307d / 3

- [blindfold_secret_info](resources--http_loadbalancer--reference--group-010.md#canonical-7c445aa96d19ecbf6979f4991d07c7ed9f7d6e177399121b88eb36a017b1191e): complete subsection reference.

- [clear_secret_info](resources--http_loadbalancer--reference--group-010.md#canonical-ba81bf3ad06bff644d88c26f9aa4a637829017ea3bf124a11a927f49207522f1): complete subsection reference.

<a id="canonical-2ed2e45f21801bda4f6219845b5198cf346d875f480e5a44caa87e47c80ceea5"></a>

## Next pages — api_testing.domains.credentials.login_endpoint.json_payload / 762d6884307d / 4

- [api_testing.domains.credentials.login_endpoint.json_payload.blindfold_secret_info](resources--http_loadbalancer--reference--group-010.md#canonical-7c445aa96d19ecbf6979f4991d07c7ed9f7d6e177399121b88eb36a017b1191e)
- [api_testing.domains.credentials.login_endpoint.json_payload.clear_secret_info](resources--http_loadbalancer--reference--group-010.md#canonical-ba81bf3ad06bff644d88c26f9aa4a637829017ea3bf124a11a927f49207522f1)
- [api_testing.domains.credentials.login_endpoint](resources--http_loadbalancer--reference--group-010.md#canonical-5c162cb70a77a04a40b0c7666a4b3f299a6a0815866bc29413b5748c87bdbd51)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-7c445aa96d19ecbf6979f4991d07c7ed9f7d6e177399121b88eb36a017b1191e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ee71619421275f552cb08cab09e6f0c52d20a8f0a1b728adb933ff04fbd976e3"></a>

## api_testing.domains.credentials.login_endpoint.json_payload.blindfold_secret_info — api_testing.domains.credentials.login_endpoint.json_payload.blindfold_secret_inf / e74b3d26fe23 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_testing](resources--http_loadbalancer--reference--group-010.md#canonical-699d08bedad7c6a3db855143b0dfa0f5040e8efa4d08cd58dccbf39afde96a71)
- [api_testing.domains](resources--http_loadbalancer--reference--group-010.md#canonical-687b0526d20593ad003d8453ffecc5d3b2c352a7eaf26c5ec5b05ee3acc553ca)
- [api_testing.domains.credentials](resources--http_loadbalancer--reference--group-010.md#canonical-b610213d02bd2bc48360d13b69ca5c914988db8ebbb077cfa1edf5a349059f36)
- [api_testing.domains.credentials.login_endpoint](resources--http_loadbalancer--reference--group-010.md#canonical-5c162cb70a77a04a40b0c7666a4b3f299a6a0815866bc29413b5748c87bdbd51)
- [api_testing.domains.credentials.login_endpoint.json_payload](resources--http_loadbalancer--reference--group-010.md#canonical-e971da389b0bf5e5e9c2c2448f15488c668ba80da1e442cf06e1646726a355b3)
- api_testing.domains.credentials.login_endpoint.json_payload.blindfold_secret_info

<a id="canonical-548dfccc9a6df294da5b388431158b06e857354192e9b7a90864e0efa6f90007"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("location")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-290c7d4a90909b82c5439a51559a1fa18e12283906e7e576e15386ac8ff4ff00"></a>

## Direct properties — api_testing.domains.credentials.login_endpoint.json_payload.blindfold_secret_inf / e74b3d26fe23 / 3

<a id="canonical-c22e0e1c55ab0a5508e1ddfd8319c2e15f6146f61676171317864379f9840aed"></a>

<a id="canonical-5a0ebad86d546350ad51a62e75ca8c8d57a95cbe0ffd678af0f733c01b260bdd"></a>

## decryption_provider property — api_testing.domains.credentials.login_endpoint.json_payload.blindfold_secret_inf / e74b3d26fe23 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the backend Secret
Management service.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-fb8815a963a40b1b4e8b81f3ad2746ccaa5202ed0a83f36c42e21e6099f42dc1"></a>

<a id="canonical-038cf9836efdf19f621b91928d21cdf5728fee3ed5cea6f34b2d27eb780ebdee"></a>

## location property — api_testing.domains.credentials.login_endpoint.json_payload.blindfold_secret_inf / e74b3d26fe23 / 5

Type: `"string"`. Optional, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
    "metadata": {
      "category": "content",
      "confidence": 1.0,
      "note": "Blindfold envelope encryption (AES-256-GCM + RSA-OAEP) of an RSA-2048 TLS private key produces ~3700 char string:/// URL. 128KB max secret size = ~175KB base64. Discovery reported 1024 which is incorrect.",
      "source": "manual-override",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 4
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-63f6f6be7bd13cfd98d34c16efdda6a2b51a947d290c79e869e59ef69f1127e7"></a>

<a id="canonical-c1bd4963c6fd9f5d47eedfd49f2b7c2b080df09ec5bb2ebce8a770e53c4202f4"></a>

## store_provider property — api_testing.domains.credentials.login_endpoint.json_payload.blindfold_secret_inf / e74b3d26fe23 / 6

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-ce5fc1f49e3a3432bb5651a77639abab025b55fb8f05dad684e4beee6ba9a78d"></a>

## Next pages — api_testing.domains.credentials.login_endpoint.json_payload.blindfold_secret_inf / e74b3d26fe23 / 7

- [api_testing.domains.credentials.login_endpoint.json_payload](resources--http_loadbalancer--reference--group-010.md#canonical-e971da389b0bf5e5e9c2c2448f15488c668ba80da1e442cf06e1646726a355b3)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-ba81bf3ad06bff644d88c26f9aa4a637829017ea3bf124a11a927f49207522f1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d52023e34901f79ad98b39603323f6879a0623c454226b86e0bb6b3ff89662d6"></a>

## api_testing.domains.credentials.login_endpoint.json_payload.clear_secret_info — api_testing.domains.credentials.login_endpoint.json_payload.clear_secret_info / 147c9e7f76b5 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_testing](resources--http_loadbalancer--reference--group-010.md#canonical-699d08bedad7c6a3db855143b0dfa0f5040e8efa4d08cd58dccbf39afde96a71)
- [api_testing.domains](resources--http_loadbalancer--reference--group-010.md#canonical-687b0526d20593ad003d8453ffecc5d3b2c352a7eaf26c5ec5b05ee3acc553ca)
- [api_testing.domains.credentials](resources--http_loadbalancer--reference--group-010.md#canonical-b610213d02bd2bc48360d13b69ca5c914988db8ebbb077cfa1edf5a349059f36)
- [api_testing.domains.credentials.login_endpoint](resources--http_loadbalancer--reference--group-010.md#canonical-5c162cb70a77a04a40b0c7666a4b3f299a6a0815866bc29413b5748c87bdbd51)
- [api_testing.domains.credentials.login_endpoint.json_payload](resources--http_loadbalancer--reference--group-010.md#canonical-e971da389b0bf5e5e9c2c2448f15488c668ba80da1e442cf06e1646726a355b3)
- api_testing.domains.credentials.login_endpoint.json_payload.clear_secret_info

<a id="canonical-16b1928c1b3db9e2fec56fb08650e9f8adb303d666ec6b3249a72bb5dfeb575d"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-c90ed2180e6d4642152fe855db8bf8fd2a22ca0f80ef95d012c2628af7d1e386"></a>

## Direct properties — api_testing.domains.credentials.login_endpoint.json_payload.clear_secret_info / 147c9e7f76b5 / 3

<a id="canonical-5a76b5546e265fa30f12c5fca260a09e843a7d484c789e2f83294bf5d3b26287"></a>

<a id="canonical-496174bc5008e4808b68b9bfbf663c9e402e0e4743131259e6ef9908df164e70"></a>

## provider_ref property — api_testing.domains.credentials.login_endpoint.json_payload.clear_secret_info / 147c9e7f76b5 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-dfe49fd9cb7e9ed9f4490304428e7cc3ef84e470e3fefa66092cafb3794fa432"></a>

<a id="canonical-70d6c230537610f596bcd673facb642902fd6058a2a65a53f7805e676adb126b"></a>

## url property — api_testing.domains.credentials.login_endpoint.json_payload.clear_secret_info / 147c9e7f76b5 / 5

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-d1e312dbba1767027310c776e9e132fe194aff145d9ed2d72b245100fb5c668f"></a>

## Next pages — api_testing.domains.credentials.login_endpoint.json_payload.clear_secret_info / 147c9e7f76b5 / 6

- [api_testing.domains.credentials.login_endpoint.json_payload](resources--http_loadbalancer--reference--group-010.md#canonical-e971da389b0bf5e5e9c2c2448f15488c668ba80da1e442cf06e1646726a355b3)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-022b9c587fad201b235813190250e458695f4c8fa89ebe6944a69960dbacee82"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ecc9d0a9eadbf1510cccd0753c97e2406c7a5fa4a652e6f56017e1ed177f2df2"></a>

## api_testing.domains.credentials.standard — api_testing.domains.credentials.standard / af3fe11ec204 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_testing](resources--http_loadbalancer--reference--group-010.md#canonical-699d08bedad7c6a3db855143b0dfa0f5040e8efa4d08cd58dccbf39afde96a71)
- [api_testing.domains](resources--http_loadbalancer--reference--group-010.md#canonical-687b0526d20593ad003d8453ffecc5d3b2c352a7eaf26c5ec5b05ee3acc553ca)
- [api_testing.domains.credentials](resources--http_loadbalancer--reference--group-010.md#canonical-b610213d02bd2bc48360d13b69ca5c914988db8ebbb077cfa1edf5a349059f36)
- api_testing.domains.credentials.standard

<a id="canonical-1bff33045f1aa23af900779b5ad1f1fa12a83ef96e9958c256ed38a7f12d1e81"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
standard = {}
```

<a id="canonical-e0e588b1a6dd1863dbbd42ee6a7de36ca8e23a79b478e99342a6394d75b2fc74"></a>

## Direct properties — api_testing.domains.credentials.standard / af3fe11ec204 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-31257dd70c9e2ba4379e95bd8f126972aa859611fb93176b4f6233c4e6ac0254"></a>

## Next pages — api_testing.domains.credentials.standard / af3fe11ec204 / 4

- [api_testing.domains.credentials](resources--http_loadbalancer--reference--group-010.md#canonical-b610213d02bd2bc48360d13b69ca5c914988db8ebbb077cfa1edf5a349059f36)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-6f5a123aaa33c1ff4b16bdb712f9d78af4cdae144c66528c2fd0e7b95f34650d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-41024d88417cdc8d69ea7c345021d912a31ad5451a50187619168643181781e3"></a>

## api_testing.every_day — api_testing.every_day / 67e779e01afa / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_testing](resources--http_loadbalancer--reference--group-010.md#canonical-699d08bedad7c6a3db855143b0dfa0f5040e8efa4d08cd58dccbf39afde96a71)
- api_testing.every_day

<a id="canonical-f13453cffbe243876de59c845362a878527d7af30ca0f9f11ddfdece6c83bd9d"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
every_day = {}
```

<a id="canonical-79c6393b2645c8d81941e2ea7c3139f92afc95f5900f67d8e76f116831c489db"></a>

## Direct properties — api_testing.every_day / 67e779e01afa / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-59b2be0e582c50f4d010d2f047c108f576b17f873dd5400737032fec6630821b"></a>

## Next pages — api_testing.every_day / 67e779e01afa / 4

- [api_testing](resources--http_loadbalancer--reference--group-010.md#canonical-699d08bedad7c6a3db855143b0dfa0f5040e8efa4d08cd58dccbf39afde96a71)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-7df86bfc68fe31586cd8cd007d1b3882017c016ccef11fe919977a1226c5ca90"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-752b47ff6debc9f23f8ef4e93dfe7d3a42dc4d29d557a76fbf5f3f59aeecd063"></a>

## api_testing.every_month — api_testing.every_month / 1fd1351b8bba / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_testing](resources--http_loadbalancer--reference--group-010.md#canonical-699d08bedad7c6a3db855143b0dfa0f5040e8efa4d08cd58dccbf39afde96a71)
- api_testing.every_month

<a id="canonical-43df27e8557c67ca6dcd2a829c5dde0ac694cd66f7ee438e06afe9c10d50f477"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for every month.

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
every_month = {}
```

<a id="canonical-3bf555686afdfdcd81bcfcb5ee97cd29a6c562bbbe72e5c3e0aa639632d83beb"></a>

## Direct properties — api_testing.every_month / 1fd1351b8bba / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-59de304887812ce11635130db6ad7e73911ec5529db70897061454ec953189c6"></a>

## Next pages — api_testing.every_month / 1fd1351b8bba / 4

- [api_testing](resources--http_loadbalancer--reference--group-010.md#canonical-699d08bedad7c6a3db855143b0dfa0f5040e8efa4d08cd58dccbf39afde96a71)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-0e2a352cfeac99dc8c252e6360e205819f252fddd47683e704ceed7790c20ff6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-61789fcad0bd0eca1a32ea2460943bc16372a788dd8c442e78c09e417326e684"></a>

## api_testing.every_week — api_testing.every_week / 8fa2332365f0 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_testing](resources--http_loadbalancer--reference--group-010.md#canonical-699d08bedad7c6a3db855143b0dfa0f5040e8efa4d08cd58dccbf39afde96a71)
- api_testing.every_week

<a id="canonical-97229c9a9526cb64b8a6c3fbcf1d93b5a6f6580522d50a4cbabdead016b9685a"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
every_week = {}
```

<a id="canonical-a9109919b6c1a6c5a6beb441ae1f541b3fb4ced9fc774747a846af9c176a094c"></a>

## Direct properties — api_testing.every_week / 8fa2332365f0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-eb7b7f507e11d6aade0966c81ebe291027239932c2f820e70bc5241531a450c5"></a>

## Next pages — api_testing.every_week / 8fa2332365f0 / 4

- [api_testing](resources--http_loadbalancer--reference--group-010.md#canonical-699d08bedad7c6a3db855143b0dfa0f5040e8efa4d08cd58dccbf39afde96a71)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-bff9954daed615e67cb6a3932ebf85a6050831ca3225573c0776c74dd5eb1f8f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a6046070408cd9c5f5da7917d372a420aaea66134794dc699b528bbd4cccc76b"></a>

## app_firewall — app_firewall / 245cd078d0a4 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- app_firewall

<a id="canonical-4f28115a42776dcf0b06835c67b25f6d7bbf4fc41c7c55a94ce67222c5aa0b51"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: app\_firewall, disable\_waf; Default: disable\_waf\] Type establishes a direct reference
from one object(the referrer) to another(the referred). Such a reference is in form of
tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

OneOf alternatives in this subsection:

- [app_firewall](resources--http_loadbalancer--reference--group-010.md#canonical-4f28115a42776dcf0b06835c67b25f6d7bbf4fc41c7c55a94ce67222c5aa0b51)
- [disable_waf](resources--http_loadbalancer--reference--group-017.md#canonical-09d5ef0dc74c89796b7c2991846a1462d1dfd53787270c8da7a0fde3c05fe341)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
app_firewall {
  # Configure direct properties listed below.
}
```

<a id="canonical-e9e25ba5087be32fd4262e601e1771874912225e7130a613696216f9e1e88c35"></a>

## Direct properties — app_firewall / 245cd078d0a4 / 3

<a id="canonical-f128e641d700b387cdad0b9048671b30b2cf70851a6eb71df4705290c8adadbe"></a>

<a id="canonical-acd92da7980ba94ed7b3196777b3ab5d43d57544779e7ee4e67f514add6ef34b"></a>

## name property — app_firewall / 245cd078d0a4 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-cab26d3f0b6dedaf849ab49f54a342ebcff9c65aa46c599e5eba08c9227421d9"></a>

<a id="canonical-d28c4489b23a789e24f6d7f7e523bc8dba46f589807bed9a8ac9444f3faf1fff"></a>

## namespace property — app_firewall / 245cd078d0a4 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-c22759a60d0e3fe327fa7ce1d8ea98fa7f3dd44f135f7c0c6c927e7afdcc345b"></a>

<a id="canonical-a0b8e45389eaa4dc2cf57e026168a6deae29a8e13eeb5d4448e3f8f3a49e1dc3"></a>

## tenant property — app_firewall / 245cd078d0a4 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-b8af0af678134f3d11fc80afdae68b29fc77506c514e7314f65b58a2a2bd492d"></a>

## Next pages — app_firewall / 245cd078d0a4 / 7

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-3a2e92765af87c3d276d5f20a62319b47649d669158cdc32c4fb394dd3035716"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d0e844c8f7e48bf3fb0c5a117ca57f9fc2c83b8ba319d61fc09cbe92f94ed6dc"></a>

## blocked_clients — blocked_clients / 25a982fd3f5f / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- blocked_clients

<a id="canonical-92118415c9cc3cb14b92a809ca6ca5d9dc675bba0cf04a69ad5a2721c7a07a28"></a>

Type: `"object"`. list nested block, Optional.

Define rules to block IP Prefixes or AS numbers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("actions"),
  validators.ConflictingListObjectAttributes("as_number",
    "http_header"),
  validators.ConflictingListObjectAttributes("as_number",
    "ip_prefix"),
  validators.ConflictingListObjectAttributes("as_number",
    "ipv6_prefix"),
  validators.ConflictingListObjectAttributes("as_number",
    "user_identifier"),
  validators.ConflictingListObjectAttributes("bot_skip_processing",
    "skip_processing"),
  validators.ConflictingListObjectAttributes("bot_skip_processing",
    "waf_skip_processing"),
  validators.ConflictingListObjectAttributes("http_header",
    "ip_prefix"),
  validators.ConflictingListObjectAttributes("http_header",
    "ipv6_prefix"),
  validators.ConflictingListObjectAttributes("http_header",
    "user_identifier"),
  validators.ConflictingListObjectAttributes("ip_prefix",
    "ipv6_prefix"),
  validators.ConflictingListObjectAttributes("ip_prefix",
    "user_identifier"),
  validators.ConflictingListObjectAttributes("ipv6_prefix",
    "user_identifier"),
  validators.ConflictingListObjectAttributes("skip_processing",
    "waf_skip_processing")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

Terraform syntax:

```terraform
blocked_clients {
  # Configure direct properties listed below.
}
```

<a id="canonical-6e49336d6a76fc29cac96ab28db6651dbe3ba96b7826525fe9095c7a1b5acf68"></a>

## Direct properties — blocked_clients / 25a982fd3f5f / 3

<a id="canonical-0de55b0a61ce52c8808124da403fd7941f1dd8f18aecd3e862cf39a0ba26cb12"></a>

<a id="canonical-64f0ecd4a8b48a6029584aef755165d0aa8aff8d1c4c17be5a9db7dc37e2435c"></a>

## actions property — blocked_clients / 25a982fd3f5f / 4

Type: `["list", "string"]`. Optional.

\[Enum:
SKIP\_PROCESSING\_WAF|SKIP\_PROCESSING\_BOT|SKIP\_PROCESSING\_MUM|SKIP\_PROCESSING\_IP\_REPUTATION|SKIP\_PROCESSING\_API\_PROTECTION|SKIP\_PROCESSING\_OAS\_VALIDATION|SKIP\_PROCESSING\_DDOS\_PROTECTION|SKIP\_PROCESSING\_THREAT\_MESH|SKIP\_PROCESSING\_MALWARE\_PROTECTION\]
Actions that should be taken when client identifier matches the rule. Possible values are
\`SKIP\_PROCESSING\_WAF\`, \`SKIP\_PROCESSING\_BOT\`, \`SKIP\_PROCESSING\_MUM\`,
\`SKIP\_PROCESSING\_IP\_REPUTATION\`, \`SKIP\_PROCESSING\_API\_PROTECTION\`,
\`SKIP\_PROCESSING\_OAS\_VALIDATION\`, \`SKIP\_PROCESSING\_DDOS\_PROTECTION\`,
\`SKIP\_PROCESSING\_THREAT\_MESH\`, \`SKIP\_PROCESSING\_MALWARE\_PROTECTION\`. Defaults to
\`SKIP\_PROCESSING\_WAF\`.

Upstream description:

Actions that should be taken when client identifier matches the rule.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(10),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 10,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 10,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "10",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "10",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-02d4145efec15c7574e1622cbe6cb811cb62067d960fd7f1576b1bc7f75c1d8f"></a>

<a id="canonical-8e9bba8f06205f62a508abae0c3bcc1971212c5382a285ac7b60e2cdc8892b65"></a>

## as_number property — blocked_clients / 25a982fd3f5f / 5

Type: `"number"`. Optional.

Exclusive with \[http\_header ip\_prefix ipv6\_prefix user\_identifier\] RFC 6793 defined 4-byte AS
number.

Upstream description:

Exclusive with \[http\_header ip\_prefix ipv6\_prefix user\_identifier\] RFC 6793 defined 4-byte AS
number.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 401308),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 401308,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "401308"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "401308"
  }
}
```

- [bot_skip_processing](resources--http_loadbalancer--reference--group-010.md#canonical-95e02bd97e3aefd12867c55712f8667506d11e49f7891b7b0bf4c9251bdc66e7): complete subsection reference.

<a id="canonical-b885cf9870a71776eb39a6da08eb80dd053f6118f33b979047627ab87ad3ff9c"></a>

<a id="canonical-0c3ba06fbac16ce489a9dd0d4849011498083fe1cf9659999761be7e47d7783f"></a>

## expiration_timestamp property — blocked_clients / 25a982fd3f5f / 6

Type: `"string"`. Optional.

Specifies expiration\_timestamp the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

Upstream description:

The expiration\_timestamp is the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "date-time",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.timestamp.within.seconds": "31536000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.timestamp.within.seconds": "31536000"
  }
}
```

- [http_header](resources--http_loadbalancer--reference--group-010.md#canonical-2ca77ca6b0be0a0935adc422da07f2a57c6135e8b5ef26e08226e7e8434003c0): complete subsection reference.

<a id="canonical-7d2e30633490f09bb9efe42e15242c9ce6119a3527ddbbb61a7f9e76cd6657c0"></a>

<a id="canonical-70a7a4c86189914eccd6c8e82f8d4192d9201a0b41491ae03948c57c70de1f63"></a>

## ip_prefix property — blocked_clients / 25a982fd3f5f / 7

Type: `"string"`. Optional.

Exclusive with \[as\_number http\_header ipv6\_prefix user\_identifier\] IPv4 prefix string.

Upstream description:

Exclusive with \[as\_number http\_header ipv6\_prefix user\_identifier\] IPv4 prefix string.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  }
}
```

<a id="canonical-227537f53e232d5412352795f3cc88067cf79c5646a2126d3b0a14cc2a87917e"></a>

<a id="canonical-2d5b53c479213d27378523527934724d851c04b9368f29a4eab0d1ff4a49b92e"></a>

## ipv6_prefix property — blocked_clients / 25a982fd3f5f / 8

Type: `"string"`. Optional.

Exclusive with \[as\_number http\_header ip\_prefix user\_identifier\] IPv6 prefix string.

Upstream description:

Exclusive with \[as\_number http\_header ip\_prefix user\_identifier\] IPv6 prefix string.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv6_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6_prefix": "true"
  }
}
```

- [metadata](resources--http_loadbalancer--reference--group-010.md#canonical-d49fa845136c542bd65f7f174b08cc150ff3188dfe78a027aa262e846d1fdc42): complete subsection reference.

- [skip_processing](resources--http_loadbalancer--reference--group-010.md#canonical-1cdd8c1fe1621e4699aeb7c38ca7b05d43be9653fc39052ac6e9ac8b8609a2f6): complete subsection reference.

<a id="canonical-95a400b2942314d58799ee33b95250a467c4e380765321594aed9933d3a5a0da"></a>

<a id="canonical-45bb157aaca3e1205ea76ebe96bd12a3087fda9fb7040c8dff742abbe48a7997"></a>

## user_identifier property — blocked_clients / 25a982fd3f5f / 9

Type: `"string"`. Optional.

Exclusive with \[as\_number http\_header ip\_prefix ipv6\_prefix\] Identify user based on user
identifier. User identifier value needs to be copied from security event.

Upstream description:

Exclusive with \[as\_number http\_header ip\_prefix ipv6\_prefix\] Identify user based on user
identifier. User identifier value needs to be copied from security event.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [waf_skip_processing](resources--http_loadbalancer--reference--group-010.md#canonical-d12d8940e71d6722389385312e21d6921a72f09de78f2e839085cec22daecb9d): complete subsection reference.

<a id="canonical-4378a56ba08301a1ae6d94d147d4ab4f6325dd838fa064a417f96ee28b6759a2"></a>

## Next pages — blocked_clients / 25a982fd3f5f / 10

- [blocked_clients.bot_skip_processing](resources--http_loadbalancer--reference--group-010.md#canonical-95e02bd97e3aefd12867c55712f8667506d11e49f7891b7b0bf4c9251bdc66e7)
- [blocked_clients.http_header](resources--http_loadbalancer--reference--group-010.md#canonical-2ca77ca6b0be0a0935adc422da07f2a57c6135e8b5ef26e08226e7e8434003c0)
- [blocked_clients.metadata](resources--http_loadbalancer--reference--group-010.md#canonical-d49fa845136c542bd65f7f174b08cc150ff3188dfe78a027aa262e846d1fdc42)
- [blocked_clients.skip_processing](resources--http_loadbalancer--reference--group-010.md#canonical-1cdd8c1fe1621e4699aeb7c38ca7b05d43be9653fc39052ac6e9ac8b8609a2f6)
- [blocked_clients.waf_skip_processing](resources--http_loadbalancer--reference--group-010.md#canonical-d12d8940e71d6722389385312e21d6921a72f09de78f2e839085cec22daecb9d)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-95e02bd97e3aefd12867c55712f8667506d11e49f7891b7b0bf4c9251bdc66e7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4a4c0660630d34a8e7fa20fc4227fb7b08ffb2ddd0f30db9d995f5c6ad2d6370"></a>

## blocked_clients.bot_skip_processing — blocked_clients.bot_skip_processing / 726daa16c667 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [blocked_clients](resources--http_loadbalancer--reference--group-010.md#canonical-3a2e92765af87c3d276d5f20a62319b47649d669158cdc32c4fb394dd3035716)
- blocked_clients.bot_skip_processing

<a id="canonical-166c0221d276b3f052a576e1d106de3abcc3163a1bdd6b4f267e2c45a9a769a9"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
bot_skip_processing = {}
```

<a id="canonical-e0702ac0b02f8c5e4480755d0913ad046115a03def97a805152bced706384e7c"></a>

## Direct properties — blocked_clients.bot_skip_processing / 726daa16c667 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-11e7f96dba98de937e1c08db44d7d537796ab3e6e2f0a07dfcf11ea47cf00cd5"></a>

## Next pages — blocked_clients.bot_skip_processing / 726daa16c667 / 4

- [blocked_clients](resources--http_loadbalancer--reference--group-010.md#canonical-3a2e92765af87c3d276d5f20a62319b47649d669158cdc32c4fb394dd3035716)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-2ca77ca6b0be0a0935adc422da07f2a57c6135e8b5ef26e08226e7e8434003c0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7b30d8e2d34bf1dab66e1b56d2b10d0239cfd0264d81a9902af7f2bac5612ccb"></a>

## blocked_clients.http_header — blocked_clients.http_header / 246c195231b1 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [blocked_clients](resources--http_loadbalancer--reference--group-010.md#canonical-3a2e92765af87c3d276d5f20a62319b47649d669158cdc32c4fb394dd3035716)
- blocked_clients.http_header

<a id="canonical-01ff20a439dfef609d4c5188adf3250c9c736e1e02560e47ecd0ac09192b4809"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for http header.

Upstream description:

Request header name and value pairs.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("headers")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
http_header {
  # Configure direct properties listed below.
}
```

<a id="canonical-1aaf5677be34ed0ba1274202d95b2cda3ce3bfbd3c9e13745cebb2dd7da8953e"></a>

## Direct properties — blocked_clients.http_header / 246c195231b1 / 3

- [headers](resources--http_loadbalancer--reference--group-010.md#canonical-ebff6f7c1f750d8b6a8b22fb484b9ac5c7478d6b64b334058b85ec809059f1ae): complete subsection reference.

<a id="canonical-7549ddb01bd20d7ad5084501b1995adc061acc194ba43f8d18375b19cf83647f"></a>

## Next pages — blocked_clients.http_header / 246c195231b1 / 4

- [blocked_clients.http_header.headers](resources--http_loadbalancer--reference--group-010.md#canonical-ebff6f7c1f750d8b6a8b22fb484b9ac5c7478d6b64b334058b85ec809059f1ae)
- [blocked_clients](resources--http_loadbalancer--reference--group-010.md#canonical-3a2e92765af87c3d276d5f20a62319b47649d669158cdc32c4fb394dd3035716)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-ebff6f7c1f750d8b6a8b22fb484b9ac5c7478d6b64b334058b85ec809059f1ae"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0f99328058a4edb9379ff222733210fba47fd3c352b69ffb9e53dd41413c5b2c"></a>

## blocked_clients.http_header.headers — blocked_clients.http_header.headers / b61b99f6891e / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [blocked_clients](resources--http_loadbalancer--reference--group-010.md#canonical-3a2e92765af87c3d276d5f20a62319b47649d669158cdc32c4fb394dd3035716)
- [blocked_clients.http_header](resources--http_loadbalancer--reference--group-010.md#canonical-2ca77ca6b0be0a0935adc422da07f2a57c6135e8b5ef26e08226e7e8434003c0)
- blocked_clients.http_header.headers

<a id="canonical-4fb5685dcb70456acfb3143a52d230a9466d0115af27267e76872dd6e2bed2a3"></a>

Type: `"object"`. list nested block, Optional.

List of HTTP header name and value pairs.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("exact",
    "presence"),
  validators.ConflictingListObjectAttributes("exact",
    "regex"),
  validators.ConflictingListObjectAttributes("presence",
    "regex")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 0,
    "uniqueItems": false
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

Terraform syntax:

```terraform
headers {
  # Configure direct properties listed below.
}
```

<a id="canonical-615b3e4f1652ce81bbfa247f416c704e12c3c043cd36d370135fc79217b2b8f1"></a>

## Direct properties — blocked_clients.http_header.headers / b61b99f6891e / 3

<a id="canonical-3074490a8475b07c86248d7f529e76a380bf862377115eb552bff6c7f182ffd0"></a>

<a id="canonical-718df94ec6a1118cb51ecf43642ad6aa2053b788df5fd5a3bb00edee73ac38ac"></a>

## exact property — blocked_clients.http_header.headers / b61b99f6891e / 4

Type: `"string"`. Optional.

Exclusive with \[presence regex\] Header value to match exactly.

Upstream description:

Exclusive with \[presence regex\] Header value to match exactly.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  }
}
```

<a id="canonical-88fd3fe91bf25b60e353cdb783d1f3ccb95281809871214fc3d838f469263001"></a>

<a id="canonical-d043d619ca83c28badb266921997826f5d050bf91735c52a8a8700e067b194b1"></a>

## invert_match property — blocked_clients.http_header.headers / b61b99f6891e / 5

Type: `"bool"`. Optional.

Invert the result of the match to detect missing header or non-matching value.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3d21e2c5aa11d8d95a48632c11d0b3ade5fa7a382baee5d385aee7bbd82dc5d1"></a>

<a id="canonical-ae88290f5c0b7d7c7dbcac5cf2c1364b8608099e63117ba9f9a72d7e1596a622"></a>

## name property — blocked_clients.http_header.headers / b61b99f6891e / 6

Type: `"string"`. Optional.

Name. Name of the header.

Upstream description:

Name of the header.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
    },
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-b7b4ba1d201a147994398665cd82997a5dfda196023ea61baadff21bcff39322"></a>

<a id="canonical-064623f7ea259368b254252e1d4aa8429fb46450f36f25a49c8159eff1078737"></a>

## presence property — blocked_clients.http_header.headers / b61b99f6891e / 7

Type: `"bool"`. Optional.

Exclusive with \[exact regex\] If true, check for presence of header.

Upstream description:

Exclusive with \[exact regex\] If true, check for presence of header.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-9eedc70b3235430f5c82f09e234392a3390cc6a46cd9661b63c9b16c46ef6a05"></a>

<a id="canonical-1415a0bcacc34425bb44d85b927c793cf9a6cb8abcba5f43350f73b99d490eaf"></a>

## regex property — blocked_clients.http_header.headers / b61b99f6891e / 8

Type: `"string"`. Optional.

Exclusive with \[exact presence\] Regex match of the header value in re2 format.

Upstream description:

Exclusive with \[exact presence\] Regex match of the header value in re2 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-2a73a93634dad8915a9249eb69738ed9857b51921453ecaae337b6e7a4284628"></a>

## Next pages — blocked_clients.http_header.headers / b61b99f6891e / 9

- [blocked_clients.http_header](resources--http_loadbalancer--reference--group-010.md#canonical-2ca77ca6b0be0a0935adc422da07f2a57c6135e8b5ef26e08226e7e8434003c0)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-d49fa845136c542bd65f7f174b08cc150ff3188dfe78a027aa262e846d1fdc42"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e32554d9e5e81c91da53f98c1589bc06d4ff869c8bcd26c72cb380c15817f0df"></a>

## blocked_clients.metadata — blocked_clients.metadata / fa8bb4cbf23f / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [blocked_clients](resources--http_loadbalancer--reference--group-010.md#canonical-3a2e92765af87c3d276d5f20a62319b47649d669158cdc32c4fb394dd3035716)
- blocked_clients.metadata

<a id="canonical-b7f52cb69a773d8a446ab762c54a34a15188ff68669c8470cfe496093e42813e"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-703adc489d6152405d33019a57c7f45da01f2d524bdda17d8ae0f8efb0c603e8"></a>

## Direct properties — blocked_clients.metadata / fa8bb4cbf23f / 3

<a id="canonical-283e64342e8c84801401d7d3e6edf70ef8d1cc3c23ea4fc9aafd62aed1fa64d5"></a>

<a id="canonical-a949466da1a88a755c143e4bee1292ff70a2a77bc9bfb855b233453292378d5e"></a>

## description_spec property — blocked_clients.metadata / fa8bb4cbf23f / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-93a6bec41e5880441d560a7e166f25bd401165f9b7923fd4fc99ef7e620ca080"></a>

<a id="canonical-ba935f044879268a7c6fcfbad859a0a6d1e9386ce98ecc9053662f175b1a8fdb"></a>

## name property — blocked_clients.metadata / fa8bb4cbf23f / 5

Type: `"string"`. Optional.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-30a15203a53dc02f323443e0338d0167403b770ac0b0fafa2e5404b17d0da756"></a>

## Next pages — blocked_clients.metadata / fa8bb4cbf23f / 6

- [blocked_clients](resources--http_loadbalancer--reference--group-010.md#canonical-3a2e92765af87c3d276d5f20a62319b47649d669158cdc32c4fb394dd3035716)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-1cdd8c1fe1621e4699aeb7c38ca7b05d43be9653fc39052ac6e9ac8b8609a2f6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-38283111a1e8babf25be038ecbe24a657d19df8f507e91aa10599a196010df30"></a>

## blocked_clients.skip_processing — blocked_clients.skip_processing / b5d629f691cd / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [blocked_clients](resources--http_loadbalancer--reference--group-010.md#canonical-3a2e92765af87c3d276d5f20a62319b47649d669158cdc32c4fb394dd3035716)
- blocked_clients.skip_processing

<a id="canonical-8957b6e3a7db54ab576ca3472c28c65a4fe138c4593a8ab5a7930fd24f94193b"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
skip_processing = {}
```

<a id="canonical-11ad68904b9c9c28f2e3e7cfce4b1b9b8ead47caf53db7d9ab5e29b202113592"></a>

## Direct properties — blocked_clients.skip_processing / b5d629f691cd / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2e879d51306e6ef5c32f50445b437164e73b408a2c8f4d531bfcdcaa0a4b7e58"></a>

## Next pages — blocked_clients.skip_processing / b5d629f691cd / 4

- [blocked_clients](resources--http_loadbalancer--reference--group-010.md#canonical-3a2e92765af87c3d276d5f20a62319b47649d669158cdc32c4fb394dd3035716)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-d12d8940e71d6722389385312e21d6921a72f09de78f2e839085cec22daecb9d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eb2619df4a2426baf6bf5863f2a9315ab16af43f16af6c18c652405af786c16a"></a>

## blocked_clients.waf_skip_processing — blocked_clients.waf_skip_processing / 101ae069752e / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [blocked_clients](resources--http_loadbalancer--reference--group-010.md#canonical-3a2e92765af87c3d276d5f20a62319b47649d669158cdc32c4fb394dd3035716)
- blocked_clients.waf_skip_processing

<a id="canonical-1868ba1e5c7b5f8a3a41743a0c0a7f78adbbaa872e2fbec5634103aed062f3b0"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
waf_skip_processing = {}
```

<a id="canonical-10edd96f7421515ede00b58332744d406b5ea9d5f6375a5c9fad864a3293b16d"></a>

## Direct properties — blocked_clients.waf_skip_processing / 101ae069752e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-166fcd7bd6e06bf04c2233845985fa18a7095c995a5370b7b405438befbdb89c"></a>

## Next pages — blocked_clients.waf_skip_processing / 101ae069752e / 4

- [blocked_clients](resources--http_loadbalancer--reference--group-010.md#canonical-3a2e92765af87c3d276d5f20a62319b47649d669158cdc32c4fb394dd3035716)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-58380aa004270508aab70b4873132bb3a9bf4d50c4b694dc88c1fc754d557ffc"></a>

## bot_defense — bot_defense / 0fdcc13e9b11 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- bot_defense

<a id="canonical-373b88b6dc3a6cb6d8f33f39467e5c90ed3f2213a20286c5ad746bf6b40bcc90"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: bot\_defense, bot\_defense\_advanced\_protection, disable\_bot\_defense; Default:
disable\_bot\_defense\] Defines various configuration OPTIONS for Bot Defense Policy.

Upstream description:

This defines various configuration OPTIONS for Bot Defense Policy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_cors_support",
    "enable_cors_support")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-cors_support_choice": "[\"disable_cors_support\",\"enable_cors_support\"]"
}
```

OneOf alternatives in this subsection:

- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-373b88b6dc3a6cb6d8f33f39467e5c90ed3f2213a20286c5ad746bf6b40bcc90)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-012.md#canonical-7d2fbebc18d0457e5b88baf2edf2236a96be148f95e0b8f95831cb57e63378ca)
- [disable_bot_defense](resources--http_loadbalancer--reference--group-017.md#canonical-0dea985b2422bf8166d9f9cfcc2f6abf7e8192e24f50ef1844e9974ffd7ea15f)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
bot_defense {
  # Configure direct properties listed below.
}
```

<a id="canonical-7de4dee4d3a60093432502b504f4975073465b697c6727fee5ef7055bc1328ef"></a>

## Direct properties — bot_defense / 0fdcc13e9b11 / 3

- [disable_cors_support](resources--http_loadbalancer--reference--group-010.md#canonical-486feb57884ea4179aba9f9965566272c2ad62db9c69c514ef9401b3403b8d93): complete subsection reference.

- [enable_cors_support](resources--http_loadbalancer--reference--group-010.md#canonical-802ccfa490089b22f4425c43005ddd6d1c2fdc4febee6a50ce7d74a5a5577fc3): complete subsection reference.

- [policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047): complete subsection reference.

<a id="canonical-23eeff803945684bc041739a03d0b7dcefb42d0e9088fd5d1ff96ad3d4118778"></a>

<a id="canonical-5eefe54472b7117614dfe4c27e3c19c1364647870eee4615d7550e1f2db40810"></a>

## regional_endpoint property — bot_defense / 0fdcc13e9b11 / 4

Type: `"string"`. Optional.

\[Enum: AUTO|US|EU|ASIA\] Defines a selection for Bot Defense region - AUTO: AUTO Automatic
selection based on client IP address - US: US US region - EU: EU European Union region - ASIA: ASIA
Asia region. Possible values are \`AUTO\`, \`US\`, \`EU\`, \`ASIA\`. Defaults to \`AUTO\`.

Upstream description:

Defines a selection for Bot Defense region

&#8203;- AUTO: AUTO

Automatic selection based on client IP address &#8203;- US: US

US region &#8203;- EU: EU

European Union region &#8203;- ASIA: ASIA

Asia region.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("AUTO",
    "US",
    "EU",
    "ASIA"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "AUTO",
  "enum": [
    "AUTO",
    "US",
    "EU",
    "ASIA"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-b8b3fc492e1d6ade134e118b5cdc1424ed54c37e9d32047b5e69b16bf8b926a9"></a>

<a id="canonical-5e13da9e13d14f50a16cc8a1959828b5f91ab473beee20c62ee465cfbd34454b"></a>

## timeout property — bot_defense / 0fdcc13e9b11 / 5

Type: `"number"`. Optional.

The timeout for the inference check, in milliseconds.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 60000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 60000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "60000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "60000"
  }
}
```

<a id="canonical-799d589aa1ee52d589d30d270065fa89a1efc1d6749fce7c77b25ce2b716d138"></a>

## Next pages — bot_defense / 0fdcc13e9b11 / 6

- [bot_defense.disable_cors_support](resources--http_loadbalancer--reference--group-010.md#canonical-486feb57884ea4179aba9f9965566272c2ad62db9c69c514ef9401b3403b8d93)
- [bot_defense.enable_cors_support](resources--http_loadbalancer--reference--group-010.md#canonical-802ccfa490089b22f4425c43005ddd6d1c2fdc4febee6a50ce7d74a5a5577fc3)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-486feb57884ea4179aba9f9965566272c2ad62db9c69c514ef9401b3403b8d93"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b2e362e20f4d18f4999b5c39b281ef664f9cbf03930545430dfa944077265b9f"></a>

## bot_defense.disable_cors_support — bot_defense.disable_cors_support / d124979240c9 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- bot_defense.disable_cors_support

<a id="canonical-481b422efdb35368ab5d8b0197e0bca820c97ef49da9e96c79745200a0381d74"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
disable_cors_support = {}
```

<a id="canonical-89644994591ce47b63940129cff489951564d034d671e8ec00fe74c1ab98c345"></a>

## Direct properties — bot_defense.disable_cors_support / d124979240c9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7a658fee2f71296cf6aff724f359087b54ead470b9efaec4b43a81eb05e2aefc"></a>

## Next pages — bot_defense.disable_cors_support / d124979240c9 / 4

- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-802ccfa490089b22f4425c43005ddd6d1c2fdc4febee6a50ce7d74a5a5577fc3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5018ac82088dd94eacecd5aa9cf93911a82d34c7b2bb22578bcedf7667ce1f94"></a>

## bot_defense.enable_cors_support — bot_defense.enable_cors_support / 34d19ed39527 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- bot_defense.enable_cors_support

<a id="canonical-9a154c28d943cff5125ad698b6c401e4b7d0079ac2a01b91698fe52d5be22adf"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
enable_cors_support = {}
```

<a id="canonical-39b11ca449a950772f6f5c7bff5a4d4e647f023ff264fd521396e2961fe0b10f"></a>

## Direct properties — bot_defense.enable_cors_support / 34d19ed39527 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-157b14b456789ad39be17bbb076adaedd0d1b69071492199c642673da2c16072"></a>

## Next pages — bot_defense.enable_cors_support / 34d19ed39527 / 4

- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-813d2f24cee0e8b6116d4110d568238355c50a1e68a20fcbb8fab349620792ab"></a>

## bot_defense.policy — bot_defense.policy / fff3f496f4aa / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- bot_defense.policy

<a id="canonical-aa750bd0a8089c3537d8c918d08b1498c4349a213e5e7bb68a65a4666607897b"></a>

Type: `"object"`. single nested block, Optional.

Defines various configuration OPTIONS for Bot Defense policy.

Upstream description:

This defines various configuration OPTIONS for Bot Defense policy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("protected_app_endpoints"),
  validators.ConflictingObjectAttributes("disable_js_insert",
    "js_insert_all_pages"),
  validators.ConflictingObjectAttributes("disable_js_insert",
    "js_insert_all_pages_except"),
  validators.ConflictingObjectAttributes("disable_js_insert",
    "js_insertion_rules"),
  validators.ConflictingObjectAttributes("disable_mobile_sdk",
    "mobile_sdk_config"),
  validators.ConflictingObjectAttributes("js_insert_all_pages",
    "js_insert_all_pages_except"),
  validators.ConflictingObjectAttributes("js_insert_all_pages",
    "js_insertion_rules"),
  validators.ConflictingObjectAttributes("js_insert_all_pages_except",
    "js_insertion_rules")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-java_script_choice": "[\"disable_js_insert\",\"js_insert_all_pages\",\"js_insert_all_pages_except\",\"js_insertion_rules\"]",
  "x-ves-oneof-field-mobile_sdk_choice": "[\"disable_mobile_sdk\",\"mobile_sdk_config\"]"
}
```

Terraform syntax:

```terraform
policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-681c5b3c8b412c984463ec9a27401a418b0710561671f278270575d6458a7491"></a>

## Direct properties — bot_defense.policy / fff3f496f4aa / 3

- [disable_js_insert](resources--http_loadbalancer--reference--group-010.md#canonical-4b600b6558eb98697410048b34a52c3e045eb97c490eea3caafc956343cd61a8): complete subsection reference.

- [disable_mobile_sdk](resources--http_loadbalancer--reference--group-010.md#canonical-acf05630d1991a2a68ba1da45646b67920f1a12d7ba61e603d1ab3bae049b173): complete subsection reference.

<a id="canonical-449928cb07d47de82fc357445f7b1c2c526698ad1551ff5f3f538cdcb5262d46"></a>

<a id="canonical-a1eb201600ff4cd0202b622f561cf8a460bd2a0c6f8c65ac249a7665d4f2d0b2"></a>

## javascript_mode property — bot_defense.policy / fff3f496f4aa / 4

Type: `"string"`. Optional.

\[Enum: ASYNC\_JS\_NO\_CACHING|ASYNC\_JS\_CACHING|SYNC\_JS\_NO\_CACHING|SYNC\_JS\_CACHING\] Web
Client JavaScript Mode. Bot Defense JavaScript for telemetry collection is requested asynchronously,
and it is non-cacheable Bot Defense JavaScript for telemetry collection is requested asynchronously,
and it is cacheable Bot Defense JavaScript for telemetry collection is requested.. Possible values
are \`ASYNC\_JS\_NO\_CACHING\`, \`ASYNC\_JS\_CACHING\`, \`SYNC\_JS\_NO\_CACHING\`,
\`SYNC\_JS\_CACHING\`. Defaults to \`ASYNC\_JS\_NO\_CACHING\`.

Upstream description:

Web Client JavaScript Mode.

Bot Defense JavaScript for telemetry collection is requested asynchronously, and it is non-cacheable
Bot Defense JavaScript for telemetry collection is requested asynchronously, and it is cacheable Bot
Defense JavaScript for telemetry collection is requested synchronously, and it is non-cacheable Bot
Defense JavaScript for telemetry collection is requested synchronously, and it is cacheable.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ASYNC_JS_NO_CACHING",
    "ASYNC_JS_CACHING",
    "SYNC_JS_NO_CACHING",
    "SYNC_JS_CACHING"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "ASYNC_JS_NO_CACHING",
  "enum": [
    "ASYNC_JS_NO_CACHING",
    "ASYNC_JS_CACHING",
    "SYNC_JS_NO_CACHING",
    "SYNC_JS_CACHING"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-5a33d2b4247a52ec8ccc9532954dd9fb8f1bf246b56c00661f15f176de36c261"></a>

<a id="canonical-b88f957281f36999317c3b6e11ce1eea420b8e798a1bf9d0e9cfc8d4308da18e"></a>

## js_download_path property — bot_defense.policy / fff3f496f4aa / 5

Type: `"string"`. Optional.

Customize Bot Defense Client JavaScript path. If not specified, default

Upstream description:

Customize Bot Defense Client JavaScript path. If not specified, default \`/common.js\`

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true"
  }
}
```

- [js_insert_all_pages](resources--http_loadbalancer--reference--group-010.md#canonical-c6f252b1c5846fc9ba690c371bac4a6d6015c56bbc2ab1f8e9ecc69de77bafca): complete subsection reference.

- [js_insert_all_pages_except](resources--http_loadbalancer--reference--group-010.md#canonical-035b1a4ae3260f37d29db3939a2013a1b761d7484d73425eef373a4961cb85e5): complete subsection reference.

- [js_insertion_rules](resources--http_loadbalancer--reference--group-011.md#canonical-c4bd537840118a268474298cc46c412b1c1b0a4f99c30d98fef38546b031777b): complete subsection reference.

- [mobile_sdk_config](resources--http_loadbalancer--reference--group-011.md#canonical-9b456501acaee5f5e09376a1cb9063698f3bb87e2f63196ff20cbe7ecd77188c): complete subsection reference.

- [protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9): complete subsection reference.

<a id="canonical-92eaa90b5fe25a4f88fe57bff0ac8b65c6499a2d7f4bc4abc1a0efac8fcd3507"></a>

## Next pages — bot_defense.policy / fff3f496f4aa / 6

- [bot_defense.policy.disable_js_insert](resources--http_loadbalancer--reference--group-010.md#canonical-4b600b6558eb98697410048b34a52c3e045eb97c490eea3caafc956343cd61a8)
- [bot_defense.policy.disable_mobile_sdk](resources--http_loadbalancer--reference--group-010.md#canonical-acf05630d1991a2a68ba1da45646b67920f1a12d7ba61e603d1ab3bae049b173)
- [bot_defense.policy.js_insert_all_pages](resources--http_loadbalancer--reference--group-010.md#canonical-c6f252b1c5846fc9ba690c371bac4a6d6015c56bbc2ab1f8e9ecc69de77bafca)
- [bot_defense.policy.js_insert_all_pages_except](resources--http_loadbalancer--reference--group-010.md#canonical-035b1a4ae3260f37d29db3939a2013a1b761d7484d73425eef373a4961cb85e5)
- [bot_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-011.md#canonical-c4bd537840118a268474298cc46c412b1c1b0a4f99c30d98fef38546b031777b)
- [bot_defense.policy.mobile_sdk_config](resources--http_loadbalancer--reference--group-011.md#canonical-9b456501acaee5f5e09376a1cb9063698f3bb87e2f63196ff20cbe7ecd77188c)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-4b600b6558eb98697410048b34a52c3e045eb97c490eea3caafc956343cd61a8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-739f86e43ce0e7b9063da7446e52b46301053bf50f919519757d54293db9fd81"></a>

## bot_defense.policy.disable_js_insert — bot_defense.policy.disable_js_insert / cf8a267400a2 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- bot_defense.policy.disable_js_insert

<a id="canonical-a1a8b1d58ae1a5fa136da3f1b236456ab4f7c85279bf18673e256e2b78513008"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable js insert.

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
disable_js_insert = {}
```

<a id="canonical-a49ae730dc17d12e8b709f04d861e170b4ac1fd9c350ecb03d7993d9204dd569"></a>

## Direct properties — bot_defense.policy.disable_js_insert / cf8a267400a2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ff8897b786752546c2980058052393a6f58aaf558cb906b9eeca509b36091d75"></a>

## Next pages — bot_defense.policy.disable_js_insert / cf8a267400a2 / 4

- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-acf05630d1991a2a68ba1da45646b67920f1a12d7ba61e603d1ab3bae049b173"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a8bdffd3e7fa97362a6c524a4b89047131b0fb279f83f0cda942cfbb78059a5f"></a>

## bot_defense.policy.disable_mobile_sdk — bot_defense.policy.disable_mobile_sdk / 1f4eb6d6d294 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- bot_defense.policy.disable_mobile_sdk

<a id="canonical-0896ad28b0b536d1f6c5845ffca91f1448f8945bb40a18693b247fb6d662a1cc"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
disable_mobile_sdk = {}
```

<a id="canonical-ace8d85f2f047cb7a802c2e8e6e1c58a565f1d0f5a2c1c3a02a394d77bceb832"></a>

## Direct properties — bot_defense.policy.disable_mobile_sdk / 1f4eb6d6d294 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d86d91b0d70bade051c45c4c357a49659db3d6a37248c6d8223b182989fdebb9"></a>

## Next pages — bot_defense.policy.disable_mobile_sdk / 1f4eb6d6d294 / 4

- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-c6f252b1c5846fc9ba690c371bac4a6d6015c56bbc2ab1f8e9ecc69de77bafca"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9274529b01f51325927e63eeae95920fd51c4a51991f00e04245a974b64e2a8d"></a>

## bot_defense.policy.js_insert_all_pages — bot_defense.policy.js_insert_all_pages / b8c845087c1f / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- bot_defense.policy.js_insert_all_pages

<a id="canonical-c4ad467ca51995e16ebef4c349d162f9b06396cb9fcf044796e8c106e99c8f55"></a>

Type: `"object"`. single nested block, Optional.

Insert Bot Defense JavaScript in all pages.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
js_insert_all_pages {
  # Configure direct properties listed below.
}
```

<a id="canonical-4b46490c05192a9da69af7add9f92cb95cdeaa1a8f4760c51f97fcad2d8977ec"></a>

## Direct properties — bot_defense.policy.js_insert_all_pages / b8c845087c1f / 3

<a id="canonical-31aeb77524e89310ca7df453b9a9eb9f0b7a9064037854c5fd68ad1c2b612ad0"></a>

<a id="canonical-19de7e5902bf0801d9414b7d170cc6d71efbde112ef01990d2dab0652f78a0ad"></a>

## javascript_location property — bot_defense.policy.js_insert_all_pages / b8c845087c1f / 4

Type: `"string"`. Optional.

\[Enum: AFTER\_HEAD|AFTER\_TITLE\_END|BEFORE\_SCRIPT\] All inside networks. Insert JavaScript after
&lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert JavaScript before first tag.
Possible values are \`AFTER\_HEAD\`, \`AFTER\_TITLE\_END\`, \`BEFORE\_SCRIPT\`. Defaults to
\`AFTER\_HEAD\`.

Upstream description:

All inside networks.

Insert JavaScript after &lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert
JavaScript before first tag.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("AFTER_HEAD",
    "AFTER_TITLE_END",
    "BEFORE_SCRIPT"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "AFTER_HEAD",
  "enum": [
    "AFTER_HEAD",
    "AFTER_TITLE_END",
    "BEFORE_SCRIPT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-b154a715bc47adc02d1f59c42148cd570c48101bc668fadae63609597229a898"></a>

## Next pages — bot_defense.policy.js_insert_all_pages / b8c845087c1f / 5

- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-035b1a4ae3260f37d29db3939a2013a1b761d7484d73425eef373a4961cb85e5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
