---
page_title: "vip_params_per_az"
subcategory: "Infrastructure"
description: "Optional Publish VIP Parameters Per AZ for public cloud sites. See documentation for 'VIP' in advertise policy to see when Inside VIP or Outside VIP is used. When configured, the VIP(s) defined will be used to publish to external systems like K8s, Consul."
xcsh_docs: {"aliases": ["vip params per az"], "body_bytes": 1885, "body_sha256": "sha256:6126332c3ea39f767a43ab99fbae205963731e127323a4f8fa11df0eaab48154", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site:properties:vip_params_per_az", "parent_id": "xcsh-docs:data-sources:site:reference", "path": "documentation/data-sources/site/properties/vip_params_per_az/index.md", "product": "distributed-cloud", "provider_name": "site", "provider_schema_digest": "sha256:057968f86e4ef0ae0087dd4d6097131e285b60998d97655ee1a02c00315f6b0f", "provider_type": "data-sources", "registry_anchor": "canonical-3133201031323320-1220303013033112-3333333330111220-3202211112310130-3110120331120120-3320030011321231-0132011012111223-3101132110203122", "registry_path": "docs/guides/data-sources--site--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["vip_params_per_az"], "schema_version": 1, "sections": [{"aliases": ["vip params per az az name"], "anchor": "schema-vip_params_per_az--az_name", "description": "AZ Name. Name of the Availability zone.", "document_id": "xcsh-docs:data-sources:site:properties:vip_params_per_az", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vip_params_per_az", "az_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["vip params per az inside vip"], "anchor": "schema-vip_params_per_az--inside_vip", "description": "Inside VIP(s). List of Inside VIPs for an AZ.", "document_id": "xcsh-docs:data-sources:site:properties:vip_params_per_az", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vip_params_per_az", "inside_vip"], "syntax": "attribute", "type": "list"}, {"aliases": ["vip params per az inside vip cname"], "anchor": "schema-vip_params_per_az--inside_vip_cname", "description": "CNAME value for the inside VIP, These are usually public cloud generated CNAME.", "document_id": "xcsh-docs:data-sources:site:properties:vip_params_per_az", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vip_params_per_az", "inside_vip_cname"], "syntax": "attribute", "type": "string"}, {"aliases": ["vip params per az inside vip v6"], "anchor": "schema-vip_params_per_az--inside_vip_v6", "description": "Optional list of Inside IPv6 VIPs for an AZ.", "document_id": "xcsh-docs:data-sources:site:properties:vip_params_per_az", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vip_params_per_az", "inside_vip_v6"], "syntax": "attribute", "type": "list"}, {"aliases": ["vip params per az outside vip"], "anchor": "schema-vip_params_per_az--outside_vip", "description": "Outside VIP(s). List of Outside VIPs for an AZ.", "document_id": "xcsh-docs:data-sources:site:properties:vip_params_per_az", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vip_params_per_az", "outside_vip"], "syntax": "attribute", "type": "list"}, {"aliases": ["vip params per az outside vip cname"], "anchor": "schema-vip_params_per_az--outside_vip_cname", "description": "CNAME value for the outside VIP These are usually public cloud generated CNAME.", "document_id": "xcsh-docs:data-sources:site:properties:vip_params_per_az", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vip_params_per_az", "outside_vip_cname"], "syntax": "attribute", "type": "string"}, {"aliases": ["vip params per az outside vip v6"], "anchor": "schema-vip_params_per_az--outside_vip_v6", "description": "Optional list of Outside IPv6 VIPs for an AZ.", "document_id": "xcsh-docs:data-sources:site:properties:vip_params_per_az", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vip_params_per_az", "outside_vip_v6"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site/properties/vip_params_per_az/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Optional Publish VIP Parameters Per AZ for public cloud sites. See documentation for 'VIP' in advertise policy to see when Inside VIP or Outside VIP is used. When configured, the VIP(s) defined will be used to publish to external systems like K8s, Consul.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": [], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# vip_params_per_az

Breadcrumbs:

- [xcsh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site/properties/)
- vip_params_per_az

<a id="section"></a>

Type: `"list"`. Computed.

Optional Publish VIP Parameters Per AZ for public cloud sites. See documentation for 'VIP' in
advertise policy to see when Inside VIP or Outside VIP is used. When configured, the VIP(s) defined
will be used to publish to external systems like K8s, Consul.

## Direct properties

<a id="schema-vip_params_per_az--az_name"></a>

### az_name property

Type: `"string"`. Computed.

AZ Name. Name of the Availability zone.

<a id="schema-vip_params_per_az--inside_vip"></a>

### inside_vip property

Type: `["list", "string"]`. Computed.

Inside VIP(s). List of Inside VIPs for an AZ.

<a id="schema-vip_params_per_az--inside_vip_cname"></a>

### inside_vip_cname property

Type: `"string"`. Computed.

CNAME value for the inside VIP, These are usually public cloud generated CNAME.

<a id="schema-vip_params_per_az--inside_vip_v6"></a>

### inside_vip_v6 property

Type: `["list", "string"]`. Computed.

Optional list of Inside IPv6 VIPs for an AZ.

<a id="schema-vip_params_per_az--outside_vip"></a>

### outside_vip property

Type: `["list", "string"]`. Computed.

Outside VIP(s). List of Outside VIPs for an AZ.

<a id="schema-vip_params_per_az--outside_vip_cname"></a>

### outside_vip_cname property

Type: `"string"`. Computed.

CNAME value for the outside VIP These are usually public cloud generated CNAME.

<a id="schema-vip_params_per_az--outside_vip_v6"></a>

### outside_vip_v6 property

Type: `["list", "string"]`. Computed.

Optional list of Outside IPv6 VIPs for an AZ.
