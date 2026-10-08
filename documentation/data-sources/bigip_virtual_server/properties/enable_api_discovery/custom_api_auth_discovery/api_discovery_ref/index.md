---
page_title: "enable_api_discovery.custom_api_auth_discovery.api_discovery_ref"
subcategory: ""
description: "Type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name."
xcsh_docs: {"aliases": ["enable api discovery custom api auth discovery api discovery ref"], "body_bytes": 1940, "body_sha256": "sha256:62aac359397cc4859105524df6c3a3c6397e1c01bb131301b58f63d977026266", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bigip_virtual_server:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bigip_virtual_server:properties:enable_api_discovery:custom_api_auth_discovery:api_discovery_ref", "parent_id": "xcsh-docs:data-sources:bigip_virtual_server:properties:enable_api_discovery:custom_api_auth_discovery", "path": "documentation/data-sources/bigip_virtual_server/properties/enable_api_discovery/custom_api_auth_discovery/api_discovery_ref/index.md", "product": "distributed-cloud", "provider_name": "bigip_virtual_server", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-2303021133000112-3322130333023121-0302132032131033-3232121233223032-0301030222131003-0201012121210212-3221113011332302-3110232123020302", "registry_path": "docs/guides/data-sources--bigip_virtual_server--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["enable_api_discovery", "custom_api_auth_discovery", "api_discovery_ref"], "schema_version": 1, "sections": [{"aliases": ["enable api discovery custom api auth discovery api discovery ref name"], "anchor": "schema-enable_api_discovery--custom_api_auth_discovery--api_discovery_ref--name", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. Route's) name.", "document_id": "xcsh-docs:data-sources:bigip_virtual_server:properties:enable_api_discovery:custom_api_auth_discovery:api_discovery_ref", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_api_discovery", "custom_api_auth_discovery", "api_discovery_ref", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["enable api discovery custom api auth discovery api discovery ref namespace"], "anchor": "schema-enable_api_discovery--custom_api_auth_discovery--api_discovery_ref--namespace", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. Route's) namespace.", "document_id": "xcsh-docs:data-sources:bigip_virtual_server:properties:enable_api_discovery:custom_api_auth_discovery:api_discovery_ref", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_api_discovery", "custom_api_auth_discovery", "api_discovery_ref", "namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["enable api discovery custom api auth discovery api discovery ref tenant"], "anchor": "schema-enable_api_discovery--custom_api_auth_discovery--api_discovery_ref--tenant", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. Route's) tenant.", "document_id": "xcsh-docs:data-sources:bigip_virtual_server:properties:enable_api_discovery:custom_api_auth_discovery:api_discovery_ref", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_api_discovery", "custom_api_auth_discovery", "api_discovery_ref", "tenant"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bigip_virtual_server/properties/enable_api_discovery/custom_api_auth_discovery/api_discovery_ref/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": [], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_api_discovery.custom_api_auth_discovery.api_discovery_ref

Breadcrumbs:

- [xcsh_bigip_virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/)
- [enable_api_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/enable_api_discovery/)
- [enable_api_discovery.custom_api_auth_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/enable_api_discovery/custom_api_auth_discovery/)
- enable_api_discovery.custom_api_auth_discovery.api_discovery_ref

<a id="section"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

## Direct properties

<a id="schema-enable_api_discovery--custom_api_auth_discovery--api_discovery_ref--name"></a>

### name property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

<a id="schema-enable_api_discovery--custom_api_auth_discovery--api_discovery_ref--namespace"></a>

### namespace property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

<a id="schema-enable_api_discovery--custom_api_auth_discovery--api_discovery_ref--tenant"></a>

### tenant property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.
