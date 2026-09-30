---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_registration_approval."
xcsh_docs: {"aliases": [], "body_bytes": 3187, "body_sha256": "sha256:28e58591b485a667f38f7964b6cbb6cca58ee37b1a54deb63716faa45cabdc93", "child_ids": [], "collection_id": "xcsh-docs:resources:registration_approval:collection", "completeness": "complete", "id": "xcsh-docs:resources:registration_approval:reference", "parent_id": "xcsh-docs:resources:registration_approval:fundamentals", "path": "documentation/resources/registration_approval/properties/index.md", "provider_name": "registration_approval", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/registration_approval/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_registration_approval.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Property reference

Breadcrumbs:

- [xcsh_registration_approval](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration_approval/)
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
| `backup_connected_region` | [backup_connected_region](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration_approval/properties/#schema-backup_connected_region) |
| `cluster_size` | [cluster_size](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration_approval/properties/#schema-cluster_size) |
| `connected_region` | [connected_region](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration_approval/properties/#schema-connected_region) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration_approval/properties/#schema-id) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration_approval/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration_approval/properties/#schema-namespace) |
| `preferred_active_re` | [preferred_active_re](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration_approval/properties/#schema-preferred_active_re) |
| `state` | [state](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration_approval/properties/#schema-state) |

## Next pages

- [xcsh_registration_approval](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration_approval/)
