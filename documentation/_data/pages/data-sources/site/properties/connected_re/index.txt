---
page_title: "connected_re"
subcategory: "Infrastructure"
description: "Following fields are only for customer edge sites List of REs to which to which this CE initiates IPsec/SSL connection to."
xcsh_docs: {"aliases": ["connected re"], "body_bytes": 1704, "body_sha256": "sha256:fcf37370ddf7232a017defc2afc1c7d5f1e6b19eb48511cdbb20416eccb7b7fb", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site:properties:connected_re", "parent_id": "xcsh-docs:data-sources:site:reference", "path": "documentation/data-sources/site/properties/connected_re/index.md", "product": "distributed-cloud", "provider_name": "site", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-2011130012313201-2232312123022113-2022011032000300-3032313132002332-1003030100100321-0231122220201001-3222131223230231-3323322222112021", "registry_path": "docs/guides/data-sources--site--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["connected_re"], "schema_version": 1, "sections": [{"aliases": ["connected re kind"], "anchor": "schema-connected_re--kind", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. 'route').", "document_id": "xcsh-docs:data-sources:site:properties:connected_re", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["connected_re", "kind"], "syntax": "attribute", "type": "string"}, {"aliases": ["connected re name"], "anchor": "schema-connected_re--name", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. Route's) name.", "document_id": "xcsh-docs:data-sources:site:properties:connected_re", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["connected_re", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["connected re namespace"], "anchor": "schema-connected_re--namespace", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. Route's) namespace.", "document_id": "xcsh-docs:data-sources:site:properties:connected_re", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["connected_re", "namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["connected re tenant"], "anchor": "schema-connected_re--tenant", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. Route's) tenant.", "document_id": "xcsh-docs:data-sources:site:properties:connected_re", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["connected_re", "tenant"], "syntax": "attribute", "type": "string"}, {"aliases": ["connected re uid"], "anchor": "schema-connected_re--uid", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. Route's) uid.", "document_id": "xcsh-docs:data-sources:site:properties:connected_re", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["connected_re", "uid"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site/properties/connected_re/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Following fields are only for customer edge sites List of REs to which to which this CE initiates IPsec/SSL connection to.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": [], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# connected_re

Breadcrumbs:

- [xcsh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site/properties/)
- connected_re

<a id="section"></a>

Type: `"list"`. Computed.

Following fields are only for customer edge sites List of REs to which to which this CE initiates
IPsec/SSL connection to.

## Direct properties

<a id="schema-connected_re--kind"></a>

### kind property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

<a id="schema-connected_re--name"></a>

### name property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

<a id="schema-connected_re--namespace"></a>

### namespace property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

<a id="schema-connected_re--tenant"></a>

### tenant property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

<a id="schema-connected_re--uid"></a>

### uid property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.
