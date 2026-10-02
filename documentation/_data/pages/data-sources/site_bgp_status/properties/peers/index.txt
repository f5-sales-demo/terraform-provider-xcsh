---
page_title: "peers"
subcategory: ""
description: "peers for xcsh_site_bgp_status."
xcsh_docs: {"aliases": ["peers"], "body_bytes": 1636, "body_sha256": "sha256:5dcb12685c3b893e3d7e389c60f854c6b70e31f804effb8aa975c857953a305a", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_bgp_status:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_bgp_status:properties:peers", "parent_id": "xcsh-docs:data-sources:site_bgp_status:reference", "path": "documentation/data-sources/site_bgp_status/properties/peers/index.md", "product": "distributed-cloud", "provider_name": "site_bgp_status", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "data-sources", "registry_anchor": "canonical-3312332133002221-1122310103103321-1333013223330010-2021001000100231-0212110020211312-2223302022233111-0202230313323310-0110212231112123", "registry_path": "docs/guides/data-sources--site_bgp_status--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["peers"], "schema_version": 1, "sections": [{"aliases": ["advertised prefix count"], "anchor": "schema-peers--advertised_prefix_count", "description": "advertised prefix count", "document_id": "xcsh-docs:data-sources:site_bgp_status:properties:peers", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["peers", "advertised_prefix_count"], "syntax": "attribute", "type": "number"}, {"aliases": ["established"], "anchor": "schema-peers--established", "description": "established", "document_id": "xcsh-docs:data-sources:site_bgp_status:properties:peers", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["peers", "established"], "syntax": "attribute", "type": "bool"}, {"aliases": ["interface name"], "anchor": "schema-peers--interface_name", "description": "interface name", "document_id": "xcsh-docs:data-sources:site_bgp_status:properties:peers", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["peers", "interface_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["mac"], "anchor": "schema-peers--mac", "description": "mac", "document_id": "xcsh-docs:data-sources:site_bgp_status:properties:peers", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["peers", "mac"], "syntax": "attribute", "type": "string"}, {"aliases": ["node"], "anchor": "schema-peers--node", "description": "node", "document_id": "xcsh-docs:data-sources:site_bgp_status:properties:peers", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["peers", "node"], "syntax": "attribute", "type": "string"}, {"aliases": ["peer address"], "anchor": "schema-peers--peer_address", "description": "peer address", "document_id": "xcsh-docs:data-sources:site_bgp_status:properties:peers", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["peers", "peer_address"], "syntax": "attribute", "type": "string"}, {"aliases": ["received prefix count"], "anchor": "schema-peers--received_prefix_count", "description": "received prefix count", "document_id": "xcsh-docs:data-sources:site_bgp_status:properties:peers", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["peers", "received_prefix_count"], "syntax": "attribute", "type": "number"}, {"aliases": ["role"], "anchor": "schema-peers--role", "description": "role", "document_id": "xcsh-docs:data-sources:site_bgp_status:properties:peers", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["peers", "role"], "syntax": "attribute", "type": "string"}, {"aliases": ["state"], "anchor": "schema-peers--state", "description": "state", "document_id": "xcsh-docs:data-sources:site_bgp_status:properties:peers", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["peers", "state"], "syntax": "attribute", "type": "string"}, {"aliases": ["state changed at"], "anchor": "schema-peers--state_changed_at", "description": "state changed at", "document_id": "xcsh-docs:data-sources:site_bgp_status:properties:peers", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["peers", "state_changed_at"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_bgp_status/properties/peers/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "peers for xcsh_site_bgp_status.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# peers

Breadcrumbs:

- [xcsh_site_bgp_status](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_bgp_status/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_bgp_status/properties/)
- peers

<a id="section"></a>

Type: `"map"`. Computed.

## Direct properties

<a id="schema-peers--advertised_prefix_count"></a>

### advertised_prefix_count property

Type: `"number"`. Computed.

<a id="schema-peers--established"></a>

### established property

Type: `"bool"`. Computed.

<a id="schema-peers--interface_name"></a>

### interface_name property

Type: `"string"`. Computed.

<a id="schema-peers--mac"></a>

### mac property

Type: `"string"`. Computed.

<a id="schema-peers--node"></a>

### node property

Type: `"string"`. Computed.

<a id="schema-peers--peer_address"></a>

### peer_address property

Type: `"string"`. Computed.

<a id="schema-peers--received_prefix_count"></a>

### received_prefix_count property

Type: `"number"`. Computed.

<a id="schema-peers--role"></a>

### role property

Type: `"string"`. Computed.

<a id="schema-peers--state"></a>

### state property

Type: `"string"`. Computed.

<a id="schema-peers--state_changed_at"></a>

### state_changed_at property

Type: `"string"`. Computed.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_bgp_status/properties/)
- [xcsh_site_bgp_status](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_bgp_status/)
