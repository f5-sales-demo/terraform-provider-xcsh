---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_registration_approval."
xcsh_docs: {"aliases": [], "body_bytes": 2760, "body_sha256": "sha256:9ad0484663bbd1dba29b8007c3c1492122c4993efa099ba9a6528eb14ccfac34", "canonical_id": "xcsh-docs:resources:registration_approval:reference", "child_ids": [], "collection_id": "xcsh-docs:resources:registration_approval:collection", "completeness": "complete", "id": "xcsh-docs:resources:registration_approval:reference", "parent_id": "xcsh-docs:resources:registration_approval:fundamentals", "path": "docs/guides/resources--registration_approval--reference.md", "provider_name": "registration_approval", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/registration_approval/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_registration_approval.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_registration_approval](../resources/registration_approval.md)
- Property reference

## Direct properties

<a id="schema-backup_connected_region"></a>

### backup_connected_region property

Type: `"string"`. Optional.

<a id="schema-cluster_size"></a>

### cluster_size property

Type: `"number"`. Required.

Number of nodes in the registration's site cluster. Use 1 for a single-node site and the complete
node count for an HA site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.OneOf(1, 3),
}
```

<a id="schema-connected_region"></a>

### connected_region property

Type: `"string"`. Optional.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NameValidator(),
}
```

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Required.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NamespaceValidator(),
}
```

<a id="schema-preferred_active_re"></a>

### preferred_active_re property

Type: `"string"`. Optional.

<a id="schema-state"></a>

### state property

Type: `"string"`. Optional, Computed.

Provider validators and defaults (from schema source):

```go
Default: stringdefault.StaticString("APPROVED")
```

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `backup_connected_region` | [backup_connected_region](resources--registration_approval--reference.md#schema-backup_connected_region) |
| `cluster_size` | [cluster_size](resources--registration_approval--reference.md#schema-cluster_size) |
| `connected_region` | [connected_region](resources--registration_approval--reference.md#schema-connected_region) |
| `id` | [id](resources--registration_approval--reference.md#schema-id) |
| `name` | [name](resources--registration_approval--reference.md#schema-name) |
| `namespace` | [namespace](resources--registration_approval--reference.md#schema-namespace) |
| `preferred_active_re` | [preferred_active_re](resources--registration_approval--reference.md#schema-preferred_active_re) |
| `state` | [state](resources--registration_approval--reference.md#schema-state) |

## Next pages

- [xcsh_registration_approval](../resources/registration_approval.md)
