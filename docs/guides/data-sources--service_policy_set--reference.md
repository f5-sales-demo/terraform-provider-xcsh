---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_service_policy_set."
xcsh_docs: {"aliases": [], "body_bytes": 2505, "body_sha256": "sha256:42db7e7546fcf0ccfe4f34a4469b55e1acb34a9fe3a6790af7a3398725f582ef", "canonical_id": "xcsh-docs:data-sources:service_policy_set:reference", "child_ids": ["xcsh-docs:data-sources:service_policy_set:properties:policies"], "collection_id": "xcsh-docs:data-sources:service_policy_set:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:service_policy_set:reference", "parent_id": "xcsh-docs:data-sources:service_policy_set:fundamentals", "path": "docs/guides/data-sources--service_policy_set--reference.md", "provider_name": "service_policy_set", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/service_policy_set/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_service_policy_set.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Property reference

Breadcrumbs:

- [xcsh_service_policy_set](../data-sources/service_policy_set.md)
- Property reference

## Direct properties

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Computed.

Annotations.

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

Name of the ServicePolicySet to look up.

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Required.

Namespace of the ServicePolicySet.

- [policies](data-sources--service_policy_set--properties--policies.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--service_policy_set--reference.md#schema-annotations) |
| `description` | [description](data-sources--service_policy_set--reference.md#schema-description) |
| `id` | [id](data-sources--service_policy_set--reference.md#schema-id) |
| `labels` | [labels](data-sources--service_policy_set--reference.md#schema-labels) |
| `name` | [name](data-sources--service_policy_set--reference.md#schema-name) |
| `namespace` | [namespace](data-sources--service_policy_set--reference.md#schema-namespace) |
| `policies` | [policies](data-sources--service_policy_set--properties--policies.md#section) |
| `policies.kind` | [policies.kind](data-sources--service_policy_set--properties--policies.md#schema-policies--kind) |
| `policies.name` | [policies.name](data-sources--service_policy_set--properties--policies.md#schema-policies--name) |
| `policies.namespace` | [policies.namespace](data-sources--service_policy_set--properties--policies.md#schema-policies--namespace) |
| `policies.tenant` | [policies.tenant](data-sources--service_policy_set--properties--policies.md#schema-policies--tenant) |
| `policies.uid` | [policies.uid](data-sources--service_policy_set--properties--policies.md#schema-policies--uid) |

## Next pages

- [policies](data-sources--service_policy_set--properties--policies.md)
- [xcsh_service_policy_set](../data-sources/service_policy_set.md)
