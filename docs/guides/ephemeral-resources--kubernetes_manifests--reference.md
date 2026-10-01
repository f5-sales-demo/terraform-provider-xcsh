---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_kubernetes_manifests."
xcsh_docs: {"aliases": [], "body_bytes": 1874, "body_sha256": "sha256:b7efb6fd0d29f1ce56c25fca4454858682f06f7b399d08ea3d128e8c74bd7366", "canonical_id": "xcsh-docs:ephemeral-resources:kubernetes_manifests:reference", "child_ids": ["xcsh-docs:ephemeral-resources:kubernetes_manifests:properties:extensions"], "collection_id": "xcsh-docs:ephemeral-resources:kubernetes_manifests:collection", "completeness": "complete", "id": "xcsh-docs:ephemeral-resources:kubernetes_manifests:reference", "parent_id": "xcsh-docs:ephemeral-resources:kubernetes_manifests:fundamentals", "path": "docs/guides/ephemeral-resources--kubernetes_manifests--reference.md", "provider_name": "kubernetes_manifests", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "ephemeral-resources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/ephemeral-resources/kubernetes_manifests/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_kubernetes_manifests.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_kubernetes_manifests](../ephemeral-resources/kubernetes_manifests.md)
- Property reference

## Direct properties

<a id="schema-content_type"></a>

### content_type property

Type: `"string"`. Computed, Sensitive.

The HTTP Content-Type header value specifying the content type of the body.

<a id="schema-data"></a>

### data property

Type: `"string"`. Computed, Sensitive.

The HTTP request/response body as raw binary.

- [extensions](ephemeral-resources--kubernetes_manifests--properties--extensions.md): complete subsection reference.

<a id="schema-site"></a>

### site property

Type: `"string"`. Required.

Site Name. Name of the site.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `content_type` | [content_type](ephemeral-resources--kubernetes_manifests--reference.md#schema-content_type) |
| `data` | [data](ephemeral-resources--kubernetes_manifests--reference.md#schema-data) |
| `extensions` | [extensions](ephemeral-resources--kubernetes_manifests--properties--extensions.md#section) |
| `extensions.type_url` | [extensions.type_url](ephemeral-resources--kubernetes_manifests--properties--extensions.md#schema-extensions--type_url) |
| `extensions.value` | [extensions.value](ephemeral-resources--kubernetes_manifests--properties--extensions.md#schema-extensions--value) |
| `site` | [site](ephemeral-resources--kubernetes_manifests--reference.md#schema-site) |

## Next pages

- [extensions](ephemeral-resources--kubernetes_manifests--properties--extensions.md)
- [xcsh_kubernetes_manifests](../ephemeral-resources/kubernetes_manifests.md)
