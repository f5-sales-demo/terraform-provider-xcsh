---
page_title: "vip_params_per_az"
subcategory: "Infrastructure"
description: "Optional Publish VIP Parameters Per AZ for public cloud sites. See documentation for 'VIP' in advertise policy to see when Inside VIP or Outside VIP is used. When configured, the VIP(s) defined will be used to publish to external systems like K8s, Consul."
xcsh_docs: {"aliases": ["vip params per az"], "body_bytes": 2101, "body_sha256": "sha256:ac498fdae85c959e740b7fd306cfad5826a15acc2b1d4691b07cc9a16b27ead0", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site:properties:vip_params_per_az", "parent_id": "xcsh-docs:data-sources:site:reference", "path": "documentation/data-sources/site/properties/vip_params_per_az/index.md", "product": "distributed-cloud", "provider_name": "site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3133201031323320-1220303013033112-3333333330111220-3202211112310130-3110120331120120-3320030011321231-0132011012111223-3101132110203122", "registry_path": "docs/guides/data-sources--site--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["vip_params_per_az"], "schema_version": 1, "sections": [{"aliases": ["az name"], "anchor": "schema-vip_params_per_az--az_name", "description": "AZ Name. Name of the Availability zone.", "document_id": "xcsh-docs:data-sources:site:properties:vip_params_per_az", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vip_params_per_az", "az_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["inside vip"], "anchor": "schema-vip_params_per_az--inside_vip", "description": "Inside VIP(s). List of Inside VIPs for an AZ.", "document_id": "xcsh-docs:data-sources:site:properties:vip_params_per_az", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vip_params_per_az", "inside_vip"], "syntax": "attribute", "type": "list"}, {"aliases": ["inside vip cname"], "anchor": "schema-vip_params_per_az--inside_vip_cname", "description": "CNAME value for the inside VIP, These are usually public cloud generated CNAME.", "document_id": "xcsh-docs:data-sources:site:properties:vip_params_per_az", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vip_params_per_az", "inside_vip_cname"], "syntax": "attribute", "type": "string"}, {"aliases": ["inside vip v6"], "anchor": "schema-vip_params_per_az--inside_vip_v6", "description": "Optional list of Inside IPv6 VIPs for an AZ.", "document_id": "xcsh-docs:data-sources:site:properties:vip_params_per_az", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vip_params_per_az", "inside_vip_v6"], "syntax": "attribute", "type": "list"}, {"aliases": ["outside vip"], "anchor": "schema-vip_params_per_az--outside_vip", "description": "Outside VIP(s). List of Outside VIPs for an AZ.", "document_id": "xcsh-docs:data-sources:site:properties:vip_params_per_az", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vip_params_per_az", "outside_vip"], "syntax": "attribute", "type": "list"}, {"aliases": ["outside vip cname"], "anchor": "schema-vip_params_per_az--outside_vip_cname", "description": "CNAME value for the outside VIP These are usually public cloud generated CNAME.", "document_id": "xcsh-docs:data-sources:site:properties:vip_params_per_az", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vip_params_per_az", "outside_vip_cname"], "syntax": "attribute", "type": "string"}, {"aliases": ["outside vip v6"], "anchor": "schema-vip_params_per_az--outside_vip_v6", "description": "Optional list of Outside IPv6 VIPs for an AZ.", "document_id": "xcsh-docs:data-sources:site:properties:vip_params_per_az", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vip_params_per_az", "outside_vip_v6"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site/properties/vip_params_per_az/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Optional Publish VIP Parameters Per AZ for public cloud sites. See documentation for 'VIP' in advertise policy to see when Inside VIP or Outside VIP is used. When configured, the VIP(s) defined will be used to publish to external systems like K8s, Consul.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site/properties/)
- [xcsh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site/)
