---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_nginx_csg."
xcsh_docs: {"aliases": ["nginx csg"], "body_bytes": 5273, "body_sha256": "sha256:efc995680bfb270248471429b9e7b211964991aa2e7904df4d85fe9aa68bc95b", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:nginx_csg:properties:api_discovery_spec", "xcsh-docs:data-sources:nginx_csg:properties:waf_spec"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:nginx_csg:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nginx_csg:reference", "parent_id": "xcsh-docs:data-sources:nginx_csg:fundamentals", "path": "documentation/data-sources/nginx_csg/properties/index.md", "product": "distributed-cloud", "provider_name": "nginx_csg", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-3110131130130312-1330331133222230-1102230130033011-1120202332122213-1330122020313232-2111301013301332-1220122313011023-2213323101223000", "registry_path": "docs/guides/data-sources--nginx_csg--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations.", "document_id": "xcsh-docs:data-sources:nginx_csg:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["api discovery spec"], "anchor": "section", "description": "Configuration for api_discovery_spec.", "document_id": "xcsh-docs:data-sources:nginx_csg:properties:api_discovery_spec", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_discovery_spec"], "syntax": "attribute", "type": "object"}, {"aliases": ["csg name"], "anchor": "schema-csg_name", "description": "CSGName. Name for CSG in NGINX One.", "document_id": "xcsh-docs:data-sources:nginx_csg:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["csg_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Description.", "document_id": "xcsh-docs:data-sources:nginx_csg:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier.", "document_id": "xcsh-docs:data-sources:nginx_csg:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Labels.", "document_id": "xcsh-docs:data-sources:nginx_csg:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["name"], "anchor": "schema-name", "description": "Name of the NginxCsg to look up.", "document_id": "xcsh-docs:data-sources:nginx_csg:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "Namespace of the NginxCsg.", "document_id": "xcsh-docs:data-sources:nginx_csg:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["object id"], "anchor": "schema-object_id", "description": "Identifier for config sync group in NGINX One.", "document_id": "xcsh-docs:data-sources:nginx_csg:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["object_id"], "syntax": "attribute", "type": "string"}, {"aliases": ["waf spec"], "anchor": "section", "description": "Configuration for waf_spec.", "document_id": "xcsh-docs:data-sources:nginx_csg:properties:waf_spec", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["waf_spec"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nginx_csg/properties/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Property reference for xcsh_nginx_csg.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": [], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_nginx_csg](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_csg/)
- Property reference

## Direct properties

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Computed.

Annotations.

- [api_discovery_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_csg/properties/api_discovery_spec/): complete subsection reference.

<a id="schema-csg_name"></a>

### csg_name property

Type: `"string"`. Computed.

CSGName. Name for CSG in NGINX One.

<a id="schema-description"></a>

### description property

Type: `"string"`. Computed.

Description.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier.

<a id="schema-labels"></a>

### labels property

Type: `["map", "string"]`. Computed.

Labels.

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the NginxCsg to look up.

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Required.

Namespace of the NginxCsg.

<a id="schema-object_id"></a>

### object_id property

Type: `"string"`. Computed.

Identifier for config sync group in NGINX One.

- [waf_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_csg/properties/waf_spec/): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_csg/properties/#schema-annotations) |
| `api_discovery_spec` | [api_discovery_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_csg/properties/api_discovery_spec/#section) |
| `api_discovery_spec.disabled` | [api_discovery_spec.disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_csg/properties/api_discovery_spec/disabled/#section) |
| `api_discovery_spec.enabled` | [api_discovery_spec.enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_csg/properties/api_discovery_spec/enabled/#section) |
| `csg_name` | [csg_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_csg/properties/#schema-csg_name) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_csg/properties/#schema-description) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_csg/properties/#schema-id) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_csg/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_csg/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_csg/properties/#schema-namespace) |
| `object_id` | [object_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_csg/properties/#schema-object_id) |
| `waf_spec` | [waf_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_csg/properties/waf_spec/#section) |
| `waf_spec.blocking_waf_mode` | [waf_spec.blocking_waf_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_csg/properties/waf_spec/blocking_waf_mode/#section) |
| `waf_spec.distributed_cloud_policy_management` | [waf_spec.distributed_cloud_policy_management](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_csg/properties/waf_spec/distributed_cloud_policy_management/#section) |
| `waf_spec.monitoring_waf_mode` | [waf_spec.monitoring_waf_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_csg/properties/waf_spec/monitoring_waf_mode/#section) |
| `waf_spec.nginx_policy_management` | [waf_spec.nginx_policy_management](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_csg/properties/waf_spec/nginx_policy_management/#section) |
| `waf_spec.none_waf_mode` | [waf_spec.none_waf_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_csg/properties/waf_spec/none_waf_mode/#section) |
| `waf_spec.policy_file_name` | [waf_spec.policy_file_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_csg/properties/waf_spec/#schema-waf_spec--policy_file_name) |
| `waf_spec.policy_name` | [waf_spec.policy_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_csg/properties/waf_spec/#schema-waf_spec--policy_name) |
| `waf_spec.security_log_enabled` | [waf_spec.security_log_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_csg/properties/waf_spec/#schema-waf_spec--security_log_enabled) |
| `waf_spec.security_log_file_names` | [waf_spec.security_log_file_names](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_csg/properties/waf_spec/#schema-waf_spec--security_log_file_names) |
