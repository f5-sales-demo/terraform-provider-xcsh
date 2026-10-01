---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_smsv2_contract."
xcsh_docs: {"aliases": [], "body_bytes": 6013, "body_sha256": "sha256:49068a16eba327abd87e1de6ce69e3622e6a43d35519f7370adc578419b9cbea", "canonical_id": "xcsh-docs:data-sources:smsv2_contract:reference", "child_ids": ["xcsh-docs:data-sources:smsv2_contract:properties:azure_route_server_ebgp_multihop"], "collection_id": "xcsh-docs:data-sources:smsv2_contract:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:smsv2_contract:reference", "parent_id": "xcsh-docs:data-sources:smsv2_contract:fundamentals", "path": "docs/guides/data-sources--smsv2_contract--reference.md", "provider_name": "smsv2_contract", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/smsv2_contract/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_smsv2_contract.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_smsv2_contract](../data-sources/smsv2_contract.md)
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

- [azure_route_server_ebgp_multihop](data-sources--smsv2_contract--properties--azure_route_server_ebgp_multihop.md): complete subsection reference.

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
| `api_release_commit` | [api_release_commit](data-sources--smsv2_contract--reference.md#schema-api_release_commit) |
| `api_release_tag` | [api_release_tag](data-sources--smsv2_contract--reference.md#schema-api_release_tag) |
| `aws_authorities` | [aws_authorities](data-sources--smsv2_contract--reference.md#schema-aws_authorities) |
| `aws_node_configuration` | [aws_node_configuration](data-sources--smsv2_contract--reference.md#schema-aws_node_configuration) |
| `azure_route_server_ebgp_multihop` | [azure_route_server_ebgp_multihop](data-sources--smsv2_contract--properties--azure_route_server_ebgp_multihop.md#section) |
| `azure_route_server_ebgp_multihop.availability` | [azure_route_server_ebgp_multihop.availability](data-sources--smsv2_contract--properties--azure_route_server_ebgp_multihop.md#schema-azure_route_server_ebgp_multihop--availability) |
| `azure_route_server_ebgp_multihop.enforcement` | [azure_route_server_ebgp_multihop.enforcement](data-sources--smsv2_contract--properties--azure_route_server_ebgp_multihop.md#schema-azure_route_server_ebgp_multihop--enforcement) |
| `azure_route_server_ebgp_multihop.reason` | [azure_route_server_ebgp_multihop.reason](data-sources--smsv2_contract--properties--azure_route_server_ebgp_multihop.md#schema-azure_route_server_ebgp_multihop--reason) |
| `azure_route_server_ebgp_multihop.source` | [azure_route_server_ebgp_multihop.source](data-sources--smsv2_contract--properties--azure_route_server_ebgp_multihop--source.md#section) |
| `azure_route_server_ebgp_multihop.source.asset_path` | [azure_route_server_ebgp_multihop.source.asset_path](data-sources--smsv2_contract--properties--azure_route_server_ebgp_multihop--source.md#schema-azure_route_server_ebgp_multihop--source--asset_path) |
| `azure_route_server_ebgp_multihop.source.asset_sha256` | [azure_route_server_ebgp_multihop.source.asset_sha256](data-sources--smsv2_contract--properties--azure_route_server_ebgp_multihop--source.md#schema-azure_route_server_ebgp_multihop--source--asset_sha256) |
| `azure_route_server_ebgp_multihop.source.commit` | [azure_route_server_ebgp_multihop.source.commit](data-sources--smsv2_contract--properties--azure_route_server_ebgp_multihop--source.md#schema-azure_route_server_ebgp_multihop--source--commit) |
| `azure_route_server_ebgp_multihop.source.repository` | [azure_route_server_ebgp_multihop.source.repository](data-sources--smsv2_contract--properties--azure_route_server_ebgp_multihop--source.md#schema-azure_route_server_ebgp_multihop--source--repository) |
| `azure_route_server_ebgp_multihop.source.schema_paths` | [azure_route_server_ebgp_multihop.source.schema_paths](data-sources--smsv2_contract--properties--azure_route_server_ebgp_multihop--source.md#schema-azure_route_server_ebgp_multihop--source--schema_paths) |
| `capabilities` | [capabilities](data-sources--smsv2_contract--reference.md#schema-capabilities) |
| `contract_id` | [contract_id](data-sources--smsv2_contract--reference.md#schema-contract_id) |
| `contract_version` | [contract_version](data-sources--smsv2_contract--reference.md#schema-contract_version) |
| `f5xc_authorities` | [f5xc_authorities](data-sources--smsv2_contract--reference.md#schema-f5xc_authorities) |
| `id` | [id](data-sources--smsv2_contract--reference.md#schema-id) |
| `kvm_image_resolution` | [kvm_image_resolution](data-sources--smsv2_contract--reference.md#schema-kvm_image_resolution) |
| `required_capabilities` | [required_capabilities](data-sources--smsv2_contract--reference.md#schema-required_capabilities) |
| `telemetry_schema_id` | [telemetry_schema_id](data-sources--smsv2_contract--reference.md#schema-telemetry_schema_id) |

## Next pages

- [azure_route_server_ebgp_multihop](data-sources--smsv2_contract--properties--azure_route_server_ebgp_multihop.md)
- [xcsh_smsv2_contract](../data-sources/smsv2_contract.md)
