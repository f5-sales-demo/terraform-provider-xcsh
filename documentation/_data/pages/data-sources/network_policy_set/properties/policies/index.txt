---
page_title: "policies"
subcategory: ""
description: "Ordered list of references to the network policy that make up this Network policy set."
xcsh_docs: {"aliases": ["policies"], "body_bytes": 1940, "body_sha256": "sha256:60164f1240470dafa2d0e60ea0475b922278c8e8832bd993fc2d7403483cd296", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:network_policy_set:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_policy_set:properties:policies", "parent_id": "xcsh-docs:data-sources:network_policy_set:reference", "path": "documentation/data-sources/network_policy_set/properties/policies/index.md", "product": "distributed-cloud", "provider_name": "network_policy_set", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2311000312202303-1102231001200212-1323332221030211-0011012123233133-0003030100232023-3231212213030232-2133123020320300-1011233223011120", "registry_path": "docs/guides/data-sources--network_policy_set--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["policies"], "schema_version": 1, "sections": [{"aliases": ["policies kind"], "anchor": "schema-policies--kind", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. 'route').", "document_id": "xcsh-docs:data-sources:network_policy_set:properties:policies", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["policies", "kind"], "syntax": "attribute", "type": "string"}, {"aliases": ["policies name"], "anchor": "schema-policies--name", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. Route's) name.", "document_id": "xcsh-docs:data-sources:network_policy_set:properties:policies", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["policies", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["policies namespace"], "anchor": "schema-policies--namespace", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. Route's) namespace.", "document_id": "xcsh-docs:data-sources:network_policy_set:properties:policies", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["policies", "namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["policies tenant"], "anchor": "schema-policies--tenant", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. Route's) tenant.", "document_id": "xcsh-docs:data-sources:network_policy_set:properties:policies", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["policies", "tenant"], "syntax": "attribute", "type": "string"}, {"aliases": ["policies uid"], "anchor": "schema-policies--uid", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. Route's) uid.", "document_id": "xcsh-docs:data-sources:network_policy_set:properties:policies", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["policies", "uid"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_policy_set/properties/policies/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Ordered list of references to the network policy that make up this Network policy set.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": [], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# policies

Breadcrumbs:

- [xcsh_network_policy_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_set/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_set/properties/)
- policies

<a id="section"></a>

Type: `"list"`. Computed.

Ordered list of references to the network policy that make up this Network policy set.

## Direct properties

<a id="schema-policies--kind"></a>

### kind property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

<a id="schema-policies--name"></a>

### name property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

<a id="schema-policies--namespace"></a>

### namespace property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

<a id="schema-policies--tenant"></a>

### tenant property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

<a id="schema-policies--uid"></a>

### uid property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_set/properties/)
- [xcsh_network_policy_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_set/)
