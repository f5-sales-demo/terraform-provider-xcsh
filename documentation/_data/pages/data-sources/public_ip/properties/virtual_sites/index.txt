---
page_title: "virtual_sites"
subcategory: ""
description: "virtual_sites for xcsh_public_ip."
xcsh_docs: {"aliases": [], "body_bytes": 1801, "body_sha256": "sha256:951bec3d36de6fd1e49cb7645c9e1929d1536f9e0bade06a47ba39f41429781c", "child_ids": [], "collection_id": "xcsh-docs:data-sources:public_ip:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:public_ip:properties:virtual_sites", "parent_id": "xcsh-docs:data-sources:public_ip:reference", "path": "documentation/data-sources/public_ip/properties/virtual_sites/index.md", "provider_name": "public_ip", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["virtual_sites"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/public_ip/properties/virtual_sites/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "virtual_sites for xcsh_public_ip.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

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
