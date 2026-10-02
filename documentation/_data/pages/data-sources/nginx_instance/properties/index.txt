---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_nginx_instance."
xcsh_docs: {"aliases": ["nginx instance"], "body_bytes": 5816, "body_sha256": "sha256:e0a068e33006a16b2a466a1b30f11cc09526f8dd0217d7f8590ca26d4022922b", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:nginx_instance:properties:api_discovery_spec", "xcsh-docs:data-sources:nginx_instance:properties:waf_spec"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:nginx_instance:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nginx_instance:reference", "parent_id": "xcsh-docs:data-sources:nginx_instance:fundamentals", "path": "documentation/data-sources/nginx_instance/properties/index.md", "product": "distributed-cloud", "provider_name": "nginx_instance", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "data-sources", "registry_anchor": "canonical-3131110011123322-2133331020011210-3000121001113220-1122300010232133-0220211102033013-2121023120220133-1222111233001023-2003132120300201", "registry_path": "docs/guides/data-sources--nginx_instance--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations.", "document_id": "xcsh-docs:data-sources:nginx_instance:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["api discovery spec"], "anchor": "section", "description": "Configuration for api_discovery_spec.", "document_id": "xcsh-docs:data-sources:nginx_instance:properties:api_discovery_spec", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_discovery_spec"], "syntax": "attribute", "type": "object"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Description.", "document_id": "xcsh-docs:data-sources:nginx_instance:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["host name"], "anchor": "schema-host_name", "description": "Dataplane identifier name for instance in NGINX One.", "document_id": "xcsh-docs:data-sources:nginx_instance:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["host_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier.", "document_id": "xcsh-docs:data-sources:nginx_instance:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Labels.", "document_id": "xcsh-docs:data-sources:nginx_instance:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["name"], "anchor": "schema-name", "description": "Name of the NginxInstance to look up.", "document_id": "xcsh-docs:data-sources:nginx_instance:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "Namespace of the NginxInstance.", "document_id": "xcsh-docs:data-sources:nginx_instance:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["object id"], "anchor": "schema-object_id", "description": "Identifier for individual instance in NGINX One.", "document_id": "xcsh-docs:data-sources:nginx_instance:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["object_id"], "syntax": "attribute", "type": "string"}, {"aliases": ["waf spec"], "anchor": "section", "description": "Configuration for waf_spec.", "document_id": "xcsh-docs:data-sources:nginx_instance:properties:waf_spec", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["waf_spec"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nginx_instance/properties/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Property reference for xcsh_nginx_instance.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

## Next pages

- [api_discovery_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_instance/properties/api_discovery_spec/)
- [waf_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_instance/properties/waf_spec/)
- [xcsh_nginx_instance](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_instance/)
