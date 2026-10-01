---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_network_policy_set."
xcsh_docs: {"aliases": [], "body_bytes": 2604, "body_sha256": "sha256:4c08f640ea1e848c51b65e1c1890425a9885fdea2fe361305a5c4b3c09bec43f", "canonical_id": "xcsh-docs:data-sources:network_policy_set:reference", "child_ids": ["xcsh-docs:data-sources:network_policy_set:properties:policies"], "collection_id": "xcsh-docs:data-sources:network_policy_set:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_policy_set:reference", "parent_id": "xcsh-docs:data-sources:network_policy_set:fundamentals", "path": "docs/guides/data-sources--network_policy_set--reference.md", "provider_name": "network_policy_set", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_policy_set/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_network_policy_set.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_network_policy_set](../data-sources/network_policy_set.md)
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

Name of the NetworkPolicySet to look up.

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Required.

Namespace of the NetworkPolicySet.

- [policies](data-sources--network_policy_set--properties--policies.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--network_policy_set--reference.md#schema-annotations) |
| `description` | [description](data-sources--network_policy_set--reference.md#schema-description) |
| `id` | [id](data-sources--network_policy_set--reference.md#schema-id) |
| `labels` | [labels](data-sources--network_policy_set--reference.md#schema-labels) |
| `name` | [name](data-sources--network_policy_set--reference.md#schema-name) |
| `namespace` | [namespace](data-sources--network_policy_set--reference.md#schema-namespace) |
| `policies` | [policies](data-sources--network_policy_set--properties--policies.md#section) |
| `policies.kind` | [policies.kind](data-sources--network_policy_set--properties--policies.md#schema-policies--kind) |
| `policies.name` | [policies.name](data-sources--network_policy_set--properties--policies.md#schema-policies--name) |
| `policies.namespace` | [policies.namespace](data-sources--network_policy_set--properties--policies.md#schema-policies--namespace) |
| `policies.tenant` | [policies.tenant](data-sources--network_policy_set--properties--policies.md#schema-policies--tenant) |
| `policies.uid` | [policies.uid](data-sources--network_policy_set--properties--policies.md#schema-policies--uid) |

## Next pages

- [policies](data-sources--network_policy_set--properties--policies.md)
- [xcsh_network_policy_set](../data-sources/network_policy_set.md)
