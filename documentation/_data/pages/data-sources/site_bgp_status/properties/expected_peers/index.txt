---
page_title: "expected_peers"
subcategory: ""
description: "expected_peers for xcsh_site_bgp_status."
xcsh_docs: {"aliases": ["expected peers"], "body_bytes": 1433, "body_sha256": "sha256:656bd8a4a3653047f3876e0f500933c3654d19c91c9b67effd58190d3d5d128e", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_bgp_status:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_bgp_status:properties:expected_peers", "parent_id": "xcsh-docs:data-sources:site_bgp_status:reference", "path": "documentation/data-sources/site_bgp_status/properties/expected_peers/index.md", "product": "distributed-cloud", "provider_name": "site_bgp_status", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-0322310333030033-2101001313021031-3032320333002332-2123332333302302-3122211311112202-0203332223111300-0300021121323320-2331301203303032", "registry_path": "docs/guides/data-sources--site_bgp_status--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["expected_peers"], "schema_version": 1, "sections": [{"aliases": ["expected peers expected imported routes"], "anchor": "schema-expected_peers--expected_imported_routes", "description": "Exact prefixes that must be imported from this remote peer on the expected node.", "document_id": "xcsh-docs:data-sources:site_bgp_status:properties:expected_peers", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["expected_peers", "expected_imported_routes"], "syntax": "attribute", "type": "set"}, {"aliases": ["expected peers mac"], "anchor": "schema-expected_peers--mac", "description": "mac", "document_id": "xcsh-docs:data-sources:site_bgp_status:properties:expected_peers", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["expected_peers", "mac"], "syntax": "attribute", "type": "string"}, {"aliases": ["expected peers node"], "anchor": "schema-expected_peers--node", "description": "node", "document_id": "xcsh-docs:data-sources:site_bgp_status:properties:expected_peers", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["expected_peers", "node"], "syntax": "attribute", "type": "string"}, {"aliases": ["expected peers peer address"], "anchor": "schema-expected_peers--peer_address", "description": "peer address", "document_id": "xcsh-docs:data-sources:site_bgp_status:properties:expected_peers", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["expected_peers", "peer_address"], "syntax": "attribute", "type": "string"}, {"aliases": ["expected peers role"], "anchor": "schema-expected_peers--role", "description": "role", "document_id": "xcsh-docs:data-sources:site_bgp_status:properties:expected_peers", "enum_extraction_complete": true, "enum_validators": [{"case_sensitive": true, "complete": true, "source": "ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf", "validator": "OneOf", "values": ["sli", "slo"], "version": 1}], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["expected_peers", "role"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_bgp_status/properties/expected_peers/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "expected_peers for xcsh_site_bgp_status.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": [], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
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
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["sli","slo"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{stringvalidator.OneOf("slo",
    "sli")}
```
