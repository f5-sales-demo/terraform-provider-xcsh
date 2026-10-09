---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_nginx_instance."
xcsh_docs: {"aliases": ["nginx instance"], "body_bytes": 5432, "body_sha256": "sha256:7ca7c49613108bca3c2011c202a9bd04e93f4f7d4abf02d914364b7426c30646", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:nginx_instance:properties:api_discovery_spec", "xcsh-docs:data-sources:nginx_instance:properties:waf_spec"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:nginx_instance:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nginx_instance:reference", "parent_id": "xcsh-docs:data-sources:nginx_instance:fundamentals", "path": "documentation/data-sources/nginx_instance/properties/index.md", "product": "distributed-cloud", "provider_name": "nginx_instance", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-3131110011123322-2133331020011210-3000121001113220-1122300010232133-0220211102033013-2121023120220133-1222111233001023-2003132120300201", "registry_path": "docs/guides/data-sources--nginx_instance--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations.", "document_id": "xcsh-docs:data-sources:nginx_instance:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["api discovery spec"], "anchor": "section", "description": "Configuration for api_discovery_spec.", "document_id": "xcsh-docs:data-sources:nginx_instance:properties:api_discovery_spec", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_discovery_spec"], "syntax": "attribute", "type": "object"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Description.", "document_id": "xcsh-docs:data-sources:nginx_instance:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["host name"], "anchor": "schema-host_name", "description": "Dataplane identifier name for instance in NGINX One.", "document_id": "xcsh-docs:data-sources:nginx_instance:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["host_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier.", "document_id": "xcsh-docs:data-sources:nginx_instance:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Labels.", "document_id": "xcsh-docs:data-sources:nginx_instance:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["name"], "anchor": "schema-name", "description": "Name of the NginxInstance to look up.", "document_id": "xcsh-docs:data-sources:nginx_instance:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "Namespace of the NginxInstance.", "document_id": "xcsh-docs:data-sources:nginx_instance:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["object id"], "anchor": "schema-object_id", "description": "Identifier for individual instance in NGINX One.", "document_id": "xcsh-docs:data-sources:nginx_instance:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["object_id"], "syntax": "attribute", "type": "string"}, {"aliases": ["waf spec"], "anchor": "section", "description": "Configuration for waf_spec.", "document_id": "xcsh-docs:data-sources:nginx_instance:properties:waf_spec", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["waf_spec"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nginx_instance/properties/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Property reference for xcsh_nginx_instance.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": [], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_nginx_instance](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_instance/)
- Property reference

## Direct properties

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Computed.

Annotations.

- [api_discovery_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_instance/properties/api_discovery_spec/): complete subsection reference.

<a id="schema-description"></a>

### description property

Type: `"string"`. Computed.

Description.

<a id="schema-host_name"></a>

### host_name property

Type: `"string"`. Computed.

Dataplane identifier name for instance in NGINX One.

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

Name of the NginxInstance to look up.

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Required.

Namespace of the NginxInstance.

<a id="schema-object_id"></a>

### object_id property

Type: `"string"`. Computed.

Identifier for individual instance in NGINX One.

- [waf_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_instance/properties/waf_spec/): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_instance/properties/#schema-annotations) |
| `api_discovery_spec` | [api_discovery_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_instance/properties/api_discovery_spec/#section) |
| `api_discovery_spec.disabled` | [api_discovery_spec.disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_instance/properties/api_discovery_spec/disabled/#section) |
| `api_discovery_spec.enabled` | [api_discovery_spec.enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_instance/properties/api_discovery_spec/enabled/#section) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_instance/properties/#schema-description) |
| `host_name` | [host_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_instance/properties/#schema-host_name) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_instance/properties/#schema-id) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_instance/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_instance/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_instance/properties/#schema-namespace) |
| `object_id` | [object_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_instance/properties/#schema-object_id) |
| `waf_spec` | [waf_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_instance/properties/waf_spec/#section) |
| `waf_spec.blocking_waf_mode` | [waf_spec.blocking_waf_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_instance/properties/waf_spec/blocking_waf_mode/#section) |
| `waf_spec.distributed_cloud_policy_management` | [waf_spec.distributed_cloud_policy_management](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_instance/properties/waf_spec/distributed_cloud_policy_management/#section) |
| `waf_spec.monitoring_waf_mode` | [waf_spec.monitoring_waf_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_instance/properties/waf_spec/monitoring_waf_mode/#section) |
| `waf_spec.nginx_policy_management` | [waf_spec.nginx_policy_management](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_instance/properties/waf_spec/nginx_policy_management/#section) |
| `waf_spec.none_waf_mode` | [waf_spec.none_waf_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_instance/properties/waf_spec/none_waf_mode/#section) |
| `waf_spec.policy_file_name` | [waf_spec.policy_file_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_instance/properties/waf_spec/#schema-waf_spec--policy_file_name) |
| `waf_spec.policy_name` | [waf_spec.policy_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_instance/properties/waf_spec/#schema-waf_spec--policy_name) |
| `waf_spec.security_log_enabled` | [waf_spec.security_log_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_instance/properties/waf_spec/#schema-waf_spec--security_log_enabled) |
| `waf_spec.security_log_file_names` | [waf_spec.security_log_file_names](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_instance/properties/waf_spec/#schema-waf_spec--security_log_file_names) |
