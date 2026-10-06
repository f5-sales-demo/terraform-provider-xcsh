---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_public_ip."
xcsh_docs: {"aliases": ["public ip"], "body_bytes": 3711, "body_sha256": "sha256:9b509e6ccb02802a3dc49e919ec7fe4e1ef0fabb75b3be3d04c9e267df88a894", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:public_ip:properties:virtual_sites"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:public_ip:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:public_ip:reference", "parent_id": "xcsh-docs:data-sources:public_ip:fundamentals", "path": "documentation/data-sources/public_ip/properties/index.md", "product": "distributed-cloud", "provider_name": "public_ip", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-1320331200230230-1110321111312310-1320320001202310-3322211120133022-3333120310113022-0300120302121212-2210030111311112-1221332130113230", "registry_path": "docs/guides/data-sources--public_ip--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations.", "document_id": "xcsh-docs:data-sources:public_ip:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Description.", "document_id": "xcsh-docs:data-sources:public_ip:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier.", "document_id": "xcsh-docs:data-sources:public_ip:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["ip"], "anchor": "schema-ip", "description": "IPv4 address for this object. An empty string indicates no IPv4 address is configured.", "document_id": "xcsh-docs:data-sources:public_ip:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ip"], "syntax": "attribute", "type": "string"}, {"aliases": ["ipv6"], "anchor": "schema-ipv6", "description": "IPv6 address for this object. An empty string indicates no IPv6 address is configured.", "document_id": "xcsh-docs:data-sources:public_ip:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ipv6"], "syntax": "attribute", "type": "string"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Labels.", "document_id": "xcsh-docs:data-sources:public_ip:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["name"], "anchor": "schema-name", "description": "Name of the PublicIP to look up.", "document_id": "xcsh-docs:data-sources:public_ip:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "Namespace of the PublicIP.", "document_id": "xcsh-docs:data-sources:public_ip:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["virtual sites"], "anchor": "section", "description": "Reference to virtual_site where this pubic IP will be available.", "document_id": "xcsh-docs:data-sources:public_ip:properties:virtual_sites", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["virtual_sites"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/public_ip/properties/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Property reference for xcsh_public_ip.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": [], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/public_ip/)
- Property reference

## Direct properties

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Computed.

Annotations.

<a id="schema-description"></a>

### description property

Type: `"string"`. Computed.

Description.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier.

<a id="schema-ip"></a>

### ip property

Type: `"string"`. Computed.

IPv4 address for this object. An empty string indicates no IPv4 address is configured.

<a id="schema-ipv6"></a>

### ipv6 property

Type: `"string"`. Computed.

IPv6 address for this object. An empty string indicates no IPv6 address is configured.

<a id="schema-labels"></a>

### labels property

Type: `["map", "string"]`. Computed.

Labels.

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the PublicIP to look up.

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Required.

Namespace of the PublicIP.

- [virtual_sites](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/public_ip/properties/virtual_sites/): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/public_ip/properties/#schema-annotations) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/public_ip/properties/#schema-description) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/public_ip/properties/#schema-id) |
| `ip` | [ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/public_ip/properties/#schema-ip) |
| `ipv6` | [ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/public_ip/properties/#schema-ipv6) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/public_ip/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/public_ip/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/public_ip/properties/#schema-namespace) |
| `virtual_sites` | [virtual_sites](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/public_ip/properties/virtual_sites/#section) |
| `virtual_sites.kind` | [virtual_sites.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/public_ip/properties/virtual_sites/#schema-virtual_sites--kind) |
| `virtual_sites.name` | [virtual_sites.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/public_ip/properties/virtual_sites/#schema-virtual_sites--name) |
| `virtual_sites.namespace` | [virtual_sites.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/public_ip/properties/virtual_sites/#schema-virtual_sites--namespace) |
| `virtual_sites.tenant` | [virtual_sites.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/public_ip/properties/virtual_sites/#schema-virtual_sites--tenant) |
| `virtual_sites.uid` | [virtual_sites.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/public_ip/properties/virtual_sites/#schema-virtual_sites--uid) |
