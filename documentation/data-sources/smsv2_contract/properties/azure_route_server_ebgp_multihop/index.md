---
page_title: "azure_route_server_ebgp_multihop"
subcategory: ""
description: "Authoritative Azure Route Server eBGP multihop availability and immutable source provenance."
xcsh_docs: {"aliases": ["azure route server ebgp multihop"], "body_bytes": 1120, "body_sha256": "sha256:4440aa8b9486e1b581801f736a429c303d5a53537f58f8ff94fff2d2ce6dc2fc", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:smsv2_contract:properties:azure_route_server_ebgp_multihop:source"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:smsv2_contract:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:smsv2_contract:properties:azure_route_server_ebgp_multihop", "parent_id": "xcsh-docs:data-sources:smsv2_contract:reference", "path": "documentation/data-sources/smsv2_contract/properties/azure_route_server_ebgp_multihop/index.md", "product": "distributed-cloud", "provider_name": "smsv2_contract", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-3203021021112331-3310333011000121-2322111013331311-3231330230221022-1231021022021133-0332300030130202-2002222110012321-0222000332031302", "registry_path": "docs/guides/data-sources--smsv2_contract--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["azure_route_server_ebgp_multihop"], "schema_version": 1, "sections": [{"aliases": ["azure route server ebgp multihop availability"], "anchor": "schema-azure_route_server_ebgp_multihop--availability", "description": "availability", "document_id": "xcsh-docs:data-sources:smsv2_contract:properties:azure_route_server_ebgp_multihop", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["azure_route_server_ebgp_multihop", "availability"], "syntax": "attribute", "type": "string"}, {"aliases": ["azure route server ebgp multihop enforcement"], "anchor": "schema-azure_route_server_ebgp_multihop--enforcement", "description": "enforcement", "document_id": "xcsh-docs:data-sources:smsv2_contract:properties:azure_route_server_ebgp_multihop", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["azure_route_server_ebgp_multihop", "enforcement"], "syntax": "attribute", "type": "string"}, {"aliases": ["azure route server ebgp multihop reason"], "anchor": "schema-azure_route_server_ebgp_multihop--reason", "description": "reason", "document_id": "xcsh-docs:data-sources:smsv2_contract:properties:azure_route_server_ebgp_multihop", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["azure_route_server_ebgp_multihop", "reason"], "syntax": "attribute", "type": "string"}, {"aliases": ["azure route server ebgp multihop source"], "anchor": "section", "description": "source", "document_id": "xcsh-docs:data-sources:smsv2_contract:properties:azure_route_server_ebgp_multihop:source", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["azure_route_server_ebgp_multihop", "source"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/smsv2_contract/properties/azure_route_server_ebgp_multihop/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Authoritative Azure Route Server eBGP multihop availability and immutable source provenance.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": [], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# azure_route_server_ebgp_multihop

Breadcrumbs:

- [xcsh_smsv2_contract](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_contract/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_contract/properties/)
- azure_route_server_ebgp_multihop

<a id="section"></a>

Type: `"single"`. Computed.

Authoritative Azure Route Server eBGP multihop availability and immutable source provenance.

## Direct properties

<a id="schema-azure_route_server_ebgp_multihop--availability"></a>

### availability property

Type: `"string"`. Computed.

<a id="schema-azure_route_server_ebgp_multihop--enforcement"></a>

### enforcement property

Type: `"string"`. Computed.

<a id="schema-azure_route_server_ebgp_multihop--reason"></a>

### reason property

Type: `"string"`. Computed.

- [source](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_contract/properties/azure_route_server_ebgp_multihop/source/): complete subsection reference.
