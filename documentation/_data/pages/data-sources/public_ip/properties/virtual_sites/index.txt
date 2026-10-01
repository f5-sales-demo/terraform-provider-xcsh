---
page_title: "virtual_sites"
subcategory: ""
description: "virtual_sites for xcsh_public_ip."
xcsh_docs: {"aliases": [], "body_bytes": 1900, "body_sha256": "sha256:2be2d83c850e1e919f1b27689e181d8969144595f58f9d5b435b70919ff1d7ca", "child_ids": [], "collection_id": "xcsh-docs:data-sources:public_ip:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:public_ip:properties:virtual_sites", "parent_id": "xcsh-docs:data-sources:public_ip:reference", "path": "documentation/data-sources/public_ip/properties/virtual_sites/index.md", "provider_name": "public_ip", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "properties", "schema_path": ["virtual_sites"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/public_ip/properties/virtual_sites/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "virtual_sites for xcsh_public_ip.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# virtual_sites

Breadcrumbs:

- [xcsh_public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/public_ip/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/public_ip/properties/)
- virtual_sites

<a id="section"></a>

Type: `"list"`. Computed.

Reference to virtual\_site where this pubic IP will be available.

## Direct properties

<a id="schema-virtual_sites--kind"></a>

### kind property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

<a id="schema-virtual_sites--name"></a>

### name property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

<a id="schema-virtual_sites--namespace"></a>

### namespace property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

<a id="schema-virtual_sites--tenant"></a>

### tenant property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

<a id="schema-virtual_sites--uid"></a>

### uid property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/public_ip/properties/)
- [xcsh_public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/public_ip/)
