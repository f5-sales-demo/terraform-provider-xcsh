---
page_title: "connected_re_for_config"
subcategory: "Infrastructure"
description: "connected_re_for_config for xcsh_site."
xcsh_docs: {"aliases": [], "body_bytes": 1648, "body_sha256": "sha256:741243c9744cae84536bdf64233f4ec140cbb3066270d3322c1688f1994f8c87", "canonical_id": "xcsh-docs:data-sources:site:properties:connected_re_for_config", "child_ids": [], "collection_id": "xcsh-docs:data-sources:site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site:properties:connected_re_for_config", "parent_id": "xcsh-docs:data-sources:site:reference", "path": "docs/guides/data-sources--site--properties--connected_re_for_config.md", "provider_name": "site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["connected_re_for_config"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site/properties/connected_re_for_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "connected_re_for_config for xcsh_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# connected_re_for_config

Breadcrumbs:

- [xcsh_site](../data-sources/site.md)
- [Property reference](data-sources--site--reference.md)
- connected_re_for_config

<a id="section"></a>

Type: `"list"`. Computed.

Valid only for CE site object List of REs which can send config to this CE site.

## Direct properties

<a id="schema-connected_re_for_config--kind"></a>

### kind property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

<a id="schema-connected_re_for_config--name"></a>

### name property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

<a id="schema-connected_re_for_config--namespace"></a>

### namespace property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

<a id="schema-connected_re_for_config--tenant"></a>

### tenant property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

<a id="schema-connected_re_for_config--uid"></a>

### uid property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

## Next pages

- [Property reference](data-sources--site--reference.md)
- [xcsh_site](../data-sources/site.md)
