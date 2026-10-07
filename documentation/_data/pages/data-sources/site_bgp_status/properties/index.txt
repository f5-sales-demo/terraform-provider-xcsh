---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_site_bgp_status."
xcsh_docs: {"aliases": ["site bgp status"], "body_bytes": 7390, "body_sha256": "sha256:1e55eaa984a4d2140f035988051db0da3334384a5d0e6091fad58a1a6c63d6a3", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:site_bgp_status:properties:expected_peers", "xcsh-docs:data-sources:site_bgp_status:properties:peers"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_bgp_status:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_bgp_status:reference", "parent_id": "xcsh-docs:data-sources:site_bgp_status:fundamentals", "path": "documentation/data-sources/site_bgp_status/properties/index.md", "product": "distributed-cloud", "provider_name": "site_bgp_status", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-2303311322002032-0001001012023033-2023213203123302-2111022211132330-1321230133100322-0231121310002212-0031102223011210-2123231130230213", "registry_path": "docs/guides/data-sources--site_bgp_status--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["bgp routes json"], "anchor": "schema-bgp_routes_json", "description": "bgp routes json", "document_id": "xcsh-docs:data-sources:site_bgp_status:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bgp_routes_json"], "syntax": "attribute", "type": "string"}, {"aliases": ["converged"], "anchor": "schema-converged", "description": "converged", "document_id": "xcsh-docs:data-sources:site_bgp_status:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["converged"], "syntax": "attribute", "type": "bool"}, {"aliases": ["expected exported routes"], "anchor": "schema-expected_exported_routes", "description": "Exact prefixes that every expected node must export and carry in the selected SLO or SLI route view.", "document_id": "xcsh-docs:data-sources:site_bgp_status:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["expected_exported_routes"], "syntax": "attribute", "type": "set"}, {"aliases": ["expected peers"], "anchor": "section", "description": "expected peers", "document_id": "xcsh-docs:data-sources:site_bgp_status:properties:expected_peers", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": "map", "relationships": [], "schema_path": ["expected_peers"], "syntax": "attribute", "type": "object"}, {"aliases": ["id"], "anchor": "schema-id", "description": "id", "document_id": "xcsh-docs:data-sources:site_bgp_status:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "namespace", "document_id": "xcsh-docs:data-sources:site_bgp_status:reference", "enum_extraction_complete": true, "enum_validators": [{"case_sensitive": true, "complete": true, "source": "ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf", "validator": "OneOf", "values": ["system"], "version": 1}], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["peers"], "anchor": "section", "description": "peers", "document_id": "xcsh-docs:data-sources:site_bgp_status:properties:peers", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "map", "relationships": [], "schema_path": ["peers"], "syntax": "attribute", "type": "object"}, {"aliases": ["poll interval seconds"], "anchor": "schema-poll_interval_seconds", "description": "poll interval seconds", "document_id": "xcsh-docs:data-sources:site_bgp_status:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["poll_interval_seconds"], "syntax": "attribute", "type": "number"}, {"aliases": ["site"], "anchor": "schema-site", "description": "site", "document_id": "xcsh-docs:data-sources:site_bgp_status:reference", "enum_extraction_complete": true, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["site"], "syntax": "attribute", "type": "string"}, {"aliases": ["sli routes json"], "anchor": "schema-sli_routes_json", "description": "sli routes json", "document_id": "xcsh-docs:data-sources:site_bgp_status:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["sli_routes_json"], "syntax": "attribute", "type": "string"}, {"aliases": ["slo routes json"], "anchor": "schema-slo_routes_json", "description": "slo routes json", "document_id": "xcsh-docs:data-sources:site_bgp_status:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["slo_routes_json"], "syntax": "attribute", "type": "string"}, {"aliases": ["duration", "timeout seconds"], "anchor": "schema-timeout_seconds", "description": "timeout seconds", "document_id": "xcsh-docs:data-sources:site_bgp_status:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["timeout_seconds"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_bgp_status/properties/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Property reference for xcsh_site_bgp_status.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": [], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_site_bgp_status](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_bgp_status/)
- Property reference

## Direct properties

<a id="schema-bgp_routes_json"></a>

### bgp_routes_json property

Type: `"string"`. Computed.

<a id="schema-converged"></a>

### converged property

Type: `"bool"`. Computed.

<a id="schema-expected_exported_routes"></a>

### expected_exported_routes property

Type: `["set", "string"]`. Required.

Exact prefixes that every expected node must export and carry in the selected SLO or SLI route view.

- [expected_peers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_bgp_status/properties/expected_peers/): complete subsection reference.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Required.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["system"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{stringvalidator.OneOf("system")}
```

- [peers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_bgp_status/properties/peers/): complete subsection reference.

<a id="schema-poll_interval_seconds"></a>

### poll_interval_seconds property

Type: `"number"`. Optional, Computed.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{int64validator.Between(1, 60)}
```

<a id="schema-site"></a>

### site property

Type: `"string"`. Required.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{stringvalidator.LengthAtLeast(1)}
```

<a id="schema-sli_routes_json"></a>

### sli_routes_json property

Type: `"string"`. Computed.

<a id="schema-slo_routes_json"></a>

### slo_routes_json property

Type: `"string"`. Computed.

<a id="schema-timeout_seconds"></a>

### timeout_seconds property

Type: `"number"`. Optional, Computed.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{int64validator.Between(1, 1800)}
```

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `bgp_routes_json` | [bgp_routes_json](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_bgp_status/properties/#schema-bgp_routes_json) |
| `converged` | [converged](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_bgp_status/properties/#schema-converged) |
| `expected_exported_routes` | [expected_exported_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_bgp_status/properties/#schema-expected_exported_routes) |
| `expected_peers` | [expected_peers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_bgp_status/properties/expected_peers/#section) |
| `expected_peers.expected_imported_routes` | [expected_peers.expected_imported_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_bgp_status/properties/expected_peers/#schema-expected_peers--expected_imported_routes) |
| `expected_peers.mac` | [expected_peers.mac](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_bgp_status/properties/expected_peers/#schema-expected_peers--mac) |
| `expected_peers.node` | [expected_peers.node](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_bgp_status/properties/expected_peers/#schema-expected_peers--node) |
| `expected_peers.peer_address` | [expected_peers.peer_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_bgp_status/properties/expected_peers/#schema-expected_peers--peer_address) |
| `expected_peers.role` | [expected_peers.role](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_bgp_status/properties/expected_peers/#schema-expected_peers--role) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_bgp_status/properties/#schema-id) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_bgp_status/properties/#schema-namespace) |
| `peers` | [peers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_bgp_status/properties/peers/#section) |
| `peers.advertised_prefix_count` | [peers.advertised_prefix_count](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_bgp_status/properties/peers/#schema-peers--advertised_prefix_count) |
| `peers.established` | [peers.established](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_bgp_status/properties/peers/#schema-peers--established) |
| `peers.interface_name` | [peers.interface_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_bgp_status/properties/peers/#schema-peers--interface_name) |
| `peers.mac` | [peers.mac](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_bgp_status/properties/peers/#schema-peers--mac) |
| `peers.node` | [peers.node](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_bgp_status/properties/peers/#schema-peers--node) |
| `peers.peer_address` | [peers.peer_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_bgp_status/properties/peers/#schema-peers--peer_address) |
| `peers.received_prefix_count` | [peers.received_prefix_count](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_bgp_status/properties/peers/#schema-peers--received_prefix_count) |
| `peers.role` | [peers.role](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_bgp_status/properties/peers/#schema-peers--role) |
| `peers.state` | [peers.state](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_bgp_status/properties/peers/#schema-peers--state) |
| `peers.state_changed_at` | [peers.state_changed_at](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_bgp_status/properties/peers/#schema-peers--state_changed_at) |
| `poll_interval_seconds` | [poll_interval_seconds](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_bgp_status/properties/#schema-poll_interval_seconds) |
| `site` | [site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_bgp_status/properties/#schema-site) |
| `sli_routes_json` | [sli_routes_json](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_bgp_status/properties/#schema-sli_routes_json) |
| `slo_routes_json` | [slo_routes_json](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_bgp_status/properties/#schema-slo_routes_json) |
| `timeout_seconds` | [timeout_seconds](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_bgp_status/properties/#schema-timeout_seconds) |
