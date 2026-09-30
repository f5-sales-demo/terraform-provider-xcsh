---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_cloud_region."
xcsh_docs: {"aliases": [], "body_bytes": 2578, "body_sha256": "sha256:e47ee29b9e3d9adfd4f14774234f556eb6b992e019c43c73320615b09238246c", "canonical_id": "xcsh-docs:data-sources:cloud_region:reference", "child_ids": ["xcsh-docs:data-sources:cloud_region:properties:default_policy_group", "xcsh-docs:data-sources:cloud_region:properties:policy_group"], "collection_id": "xcsh-docs:data-sources:cloud_region:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cloud_region:reference", "parent_id": "xcsh-docs:data-sources:cloud_region:fundamentals", "path": "docs/guides/data-sources--cloud_region--reference.md", "provider_name": "cloud_region", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cloud_region/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_cloud_region.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Property reference

Breadcrumbs:

- [xcsh_cloud_region](../data-sources/cloud_region.md)
- Property reference

## Direct properties

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Computed.

Annotations.

- [default_policy_group](data-sources--cloud_region--properties--default_policy_group.md): complete subsection reference.

<a id="schema-description"></a>

### description property

Type: `"string"`. Computed.

Description.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier.

<a id="schema-labels"></a>

### labels property

Type: `["map", "string"]`. Computed.

Labels.

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the CloudRegion to look up.

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Required.

Namespace of the CloudRegion.

- [policy_group](data-sources--cloud_region--properties--policy_group.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--cloud_region--reference.md#schema-annotations) |
| `default_policy_group` | [default_policy_group](data-sources--cloud_region--properties--default_policy_group.md#section) |
| `description` | [description](data-sources--cloud_region--reference.md#schema-description) |
| `id` | [id](data-sources--cloud_region--reference.md#schema-id) |
| `labels` | [labels](data-sources--cloud_region--reference.md#schema-labels) |
| `name` | [name](data-sources--cloud_region--reference.md#schema-name) |
| `namespace` | [namespace](data-sources--cloud_region--reference.md#schema-namespace) |
| `policy_group` | [policy_group](data-sources--cloud_region--properties--policy_group.md#section) |
| `policy_group.name` | [policy_group.name](data-sources--cloud_region--properties--policy_group.md#schema-policy_group--name) |
| `policy_group.namespace` | [policy_group.namespace](data-sources--cloud_region--properties--policy_group.md#schema-policy_group--namespace) |
| `policy_group.tenant` | [policy_group.tenant](data-sources--cloud_region--properties--policy_group.md#schema-policy_group--tenant) |

## Next pages

- [default_policy_group](data-sources--cloud_region--properties--default_policy_group.md)
- [policy_group](data-sources--cloud_region--properties--policy_group.md)
- [xcsh_cloud_region](../data-sources/cloud_region.md)
