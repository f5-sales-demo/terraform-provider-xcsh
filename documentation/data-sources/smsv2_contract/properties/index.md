---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_smsv2_contract."
xcsh_docs: {"aliases": ["smsv2 contract"], "body_bytes": 7064, "body_sha256": "sha256:5c84cf576e1a6d4cf69f5f8f4525490904f83cce163c66f30f0e662ee8760947", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:smsv2_contract:properties:azure_route_server_ebgp_multihop"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:smsv2_contract:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:smsv2_contract:reference", "parent_id": "xcsh-docs:data-sources:smsv2_contract:fundamentals", "path": "documentation/data-sources/smsv2_contract/properties/index.md", "product": "distributed-cloud", "provider_name": "smsv2_contract", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-2232311231011322-0030001303101003-3113022301021321-2022203112032301-3321212110130112-3122331000100210-1110021311120232-0302332010100121", "registry_path": "docs/guides/data-sources--smsv2_contract--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["api release commit"], "anchor": "schema-api_release_commit", "description": "api release commit", "document_id": "xcsh-docs:data-sources:smsv2_contract:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_release_commit"], "syntax": "attribute", "type": "string"}, {"aliases": ["api release tag"], "anchor": "schema-api_release_tag", "description": "api release tag", "document_id": "xcsh-docs:data-sources:smsv2_contract:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_release_tag"], "syntax": "attribute", "type": "string"}, {"aliases": ["aws authorities"], "anchor": "schema-aws_authorities", "description": "aws authorities", "document_id": "xcsh-docs:data-sources:smsv2_contract:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_authorities"], "syntax": "attribute", "type": "list"}, {"aliases": ["aws node configuration"], "anchor": "schema-aws_node_configuration", "description": "Canonical immutable AWS node-configuration contract JSON from the pinned API release.", "document_id": "xcsh-docs:data-sources:smsv2_contract:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_node_configuration"], "syntax": "attribute", "type": "string"}, {"aliases": ["azure route server ebgp multihop"], "anchor": "section", "description": "Authoritative Azure Route Server eBGP multihop availability and immutable source provenance.", "document_id": "xcsh-docs:data-sources:smsv2_contract:properties:azure_route_server_ebgp_multihop", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["azure_route_server_ebgp_multihop"], "syntax": "attribute", "type": "object"}, {"aliases": ["capabilities"], "anchor": "schema-capabilities", "description": "capabilities", "document_id": "xcsh-docs:data-sources:smsv2_contract:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["capabilities"], "syntax": "attribute", "type": "map"}, {"aliases": ["contract id"], "anchor": "schema-contract_id", "description": "contract id", "document_id": "xcsh-docs:data-sources:smsv2_contract:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["contract_id"], "syntax": "attribute", "type": "string"}, {"aliases": ["contract version"], "anchor": "schema-contract_version", "description": "contract version", "document_id": "xcsh-docs:data-sources:smsv2_contract:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["contract_version"], "syntax": "attribute", "type": "string"}, {"aliases": ["f5xc authorities"], "anchor": "schema-f5xc_authorities", "description": "f5xc authorities", "document_id": "xcsh-docs:data-sources:smsv2_contract:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["f5xc_authorities"], "syntax": "attribute", "type": "list"}, {"aliases": ["id"], "anchor": "schema-id", "description": "id", "document_id": "xcsh-docs:data-sources:smsv2_contract:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["kvm image resolution"], "anchor": "schema-kvm_image_resolution", "description": "Complete canonical KVM image-resolution contract JSON, including ownership, validation, and evidence provenance, from the pinned API release.", "document_id": "xcsh-docs:data-sources:smsv2_contract:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["kvm_image_resolution"], "syntax": "attribute", "type": "string"}, {"aliases": ["required capabilities"], "anchor": "schema-required_capabilities", "description": "Capabilities that must be available. A known unavailable capability produces a planning diagnostic before any F5 API request.", "document_id": "xcsh-docs:data-sources:smsv2_contract:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["required_capabilities"], "syntax": "attribute", "type": "set"}, {"aliases": ["telemetry schema id"], "anchor": "schema-telemetry_schema_id", "description": "telemetry schema id", "document_id": "xcsh-docs:data-sources:smsv2_contract:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["telemetry_schema_id"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/smsv2_contract/properties/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Property reference for xcsh_smsv2_contract.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": [], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_smsv2_contract](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_contract/)
- Property reference

## Direct properties

<a id="schema-api_release_commit"></a>

### api_release_commit property

Type: `"string"`. Computed.

<a id="schema-api_release_tag"></a>

### api_release_tag property

Type: `"string"`. Computed.

<a id="schema-aws_authorities"></a>

### aws_authorities property

Type: `["list", "string"]`. Computed.

<a id="schema-aws_node_configuration"></a>

### aws_node_configuration property

Type: `"string"`. Computed.

Canonical immutable AWS node-configuration contract JSON from the pinned API release.

- [azure_route_server_ebgp_multihop](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_contract/properties/azure_route_server_ebgp_multihop/): complete subsection reference.

<a id="schema-capabilities"></a>

### capabilities property

Type: `["map", "string"]`. Computed.

<a id="schema-contract_id"></a>

### contract_id property

Type: `"string"`. Computed.

<a id="schema-contract_version"></a>

### contract_version property

Type: `"string"`. Computed.

<a id="schema-f5xc_authorities"></a>

### f5xc_authorities property

Type: `["list", "string"]`. Computed.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

<a id="schema-kvm_image_resolution"></a>

### kvm_image_resolution property

Type: `"string"`. Computed.

Complete canonical KVM image-resolution contract JSON, including ownership, validation, and evidence
provenance, from the pinned API release.

<a id="schema-required_capabilities"></a>

### required_capabilities property

Type: `["set", "string"]`. Optional.

Capabilities that must be available. A known unavailable capability produces a planning diagnostic
before any F5 API request.

<a id="schema-telemetry_schema_id"></a>

### telemetry_schema_id property

Type: `"string"`. Computed.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `api_release_commit` | [api_release_commit](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_contract/properties/#schema-api_release_commit) |
| `api_release_tag` | [api_release_tag](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_contract/properties/#schema-api_release_tag) |
| `aws_authorities` | [aws_authorities](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_contract/properties/#schema-aws_authorities) |
| `aws_node_configuration` | [aws_node_configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_contract/properties/#schema-aws_node_configuration) |
| `azure_route_server_ebgp_multihop` | [azure_route_server_ebgp_multihop](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_contract/properties/azure_route_server_ebgp_multihop/#section) |
| `azure_route_server_ebgp_multihop.availability` | [azure_route_server_ebgp_multihop.availability](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_contract/properties/azure_route_server_ebgp_multihop/#schema-azure_route_server_ebgp_multihop--availability) |
| `azure_route_server_ebgp_multihop.enforcement` | [azure_route_server_ebgp_multihop.enforcement](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_contract/properties/azure_route_server_ebgp_multihop/#schema-azure_route_server_ebgp_multihop--enforcement) |
| `azure_route_server_ebgp_multihop.reason` | [azure_route_server_ebgp_multihop.reason](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_contract/properties/azure_route_server_ebgp_multihop/#schema-azure_route_server_ebgp_multihop--reason) |
| `azure_route_server_ebgp_multihop.source` | [azure_route_server_ebgp_multihop.source](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_contract/properties/azure_route_server_ebgp_multihop/source/#section) |
| `azure_route_server_ebgp_multihop.source.asset_path` | [azure_route_server_ebgp_multihop.source.asset_path](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_contract/properties/azure_route_server_ebgp_multihop/source/#schema-azure_route_server_ebgp_multihop--source--asset_path) |
| `azure_route_server_ebgp_multihop.source.asset_sha256` | [azure_route_server_ebgp_multihop.source.asset_sha256](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_contract/properties/azure_route_server_ebgp_multihop/source/#schema-azure_route_server_ebgp_multihop--source--asset_sha256) |
| `azure_route_server_ebgp_multihop.source.commit` | [azure_route_server_ebgp_multihop.source.commit](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_contract/properties/azure_route_server_ebgp_multihop/source/#schema-azure_route_server_ebgp_multihop--source--commit) |
| `azure_route_server_ebgp_multihop.source.repository` | [azure_route_server_ebgp_multihop.source.repository](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_contract/properties/azure_route_server_ebgp_multihop/source/#schema-azure_route_server_ebgp_multihop--source--repository) |
| `azure_route_server_ebgp_multihop.source.schema_paths` | [azure_route_server_ebgp_multihop.source.schema_paths](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_contract/properties/azure_route_server_ebgp_multihop/source/#schema-azure_route_server_ebgp_multihop--source--schema_paths) |
| `capabilities` | [capabilities](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_contract/properties/#schema-capabilities) |
| `contract_id` | [contract_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_contract/properties/#schema-contract_id) |
| `contract_version` | [contract_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_contract/properties/#schema-contract_version) |
| `f5xc_authorities` | [f5xc_authorities](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_contract/properties/#schema-f5xc_authorities) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_contract/properties/#schema-id) |
| `kvm_image_resolution` | [kvm_image_resolution](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_contract/properties/#schema-kvm_image_resolution) |
| `required_capabilities` | [required_capabilities](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_contract/properties/#schema-required_capabilities) |
| `telemetry_schema_id` | [telemetry_schema_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_contract/properties/#schema-telemetry_schema_id) |
