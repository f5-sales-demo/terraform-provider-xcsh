---
page_title: "expected_peers"
subcategory: ""
description: "expected_peers for xcsh_site_bgp_status."
xcsh_docs: {"aliases": ["expected peers"], "body_bytes": 1434, "body_sha256": "sha256:829feba6e0ad7797f89f3c2eca3d89770fce8c3fd823a711038d47af5be4d394", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_bgp_status:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_bgp_status:properties:expected_peers", "parent_id": "xcsh-docs:data-sources:site_bgp_status:reference", "path": "documentation/data-sources/site_bgp_status/properties/expected_peers/index.md", "product": "distributed-cloud", "provider_name": "site_bgp_status", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0322310333030033-2101001313021031-3032320333002332-2123332333302302-3122211311112202-0203332223111300-0300021121323320-2331301203303032", "registry_path": "docs/guides/data-sources--site_bgp_status--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["expected_peers"], "schema_version": 1, "sections": [{"aliases": ["expected imported routes"], "anchor": "schema-expected_peers--expected_imported_routes", "description": "Exact prefixes that must be imported from this remote peer on the expected node.", "document_id": "xcsh-docs:data-sources:site_bgp_status:properties:expected_peers", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["expected_peers", "expected_imported_routes"], "syntax": "attribute", "type": "set"}, {"aliases": ["mac"], "anchor": "schema-expected_peers--mac", "description": "mac", "document_id": "xcsh-docs:data-sources:site_bgp_status:properties:expected_peers", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["expected_peers", "mac"], "syntax": "attribute", "type": "string"}, {"aliases": ["node"], "anchor": "schema-expected_peers--node", "description": "node", "document_id": "xcsh-docs:data-sources:site_bgp_status:properties:expected_peers", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["expected_peers", "node"], "syntax": "attribute", "type": "string"}, {"aliases": ["peer address"], "anchor": "schema-expected_peers--peer_address", "description": "peer address", "document_id": "xcsh-docs:data-sources:site_bgp_status:properties:expected_peers", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["expected_peers", "peer_address"], "syntax": "attribute", "type": "string"}, {"aliases": ["role"], "anchor": "schema-expected_peers--role", "description": "role", "document_id": "xcsh-docs:data-sources:site_bgp_status:properties:expected_peers", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["expected_peers", "role"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_bgp_status/properties/expected_peers/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "expected_peers for xcsh_site_bgp_status.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# expected_peers

Breadcrumbs:

- [xcsh_site_bgp_status](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_bgp_status/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_bgp_status/properties/)
- expected_peers

<a id="section"></a>

Type: `"map"`. Required.

## Direct properties

<a id="schema-expected_peers--expected_imported_routes"></a>

### expected_imported_routes property

Type: `["set", "string"]`. Required.

Exact prefixes that must be imported from this remote peer on the expected node.

<a id="schema-expected_peers--mac"></a>

### mac property

Type: `"string"`. Required.

<a id="schema-expected_peers--node"></a>

### node property

Type: `"string"`. Required.

<a id="schema-expected_peers--peer_address"></a>

### peer_address property

Type: `"string"`. Required.

<a id="schema-expected_peers--role"></a>

### role property

Type: `"string"`. Required.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{stringvalidator.OneOf("slo",
    "sli")}
```

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_bgp_status/properties/)
- [xcsh_site_bgp_status](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_bgp_status/)
