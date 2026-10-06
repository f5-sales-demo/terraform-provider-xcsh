---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_nginx_server."
xcsh_docs: {"aliases": ["nginx server"], "body_bytes": 12825, "body_sha256": "sha256:568995272719278911c1c1e97b93f0eaaaa0014b0ff6cab5604debf74ef6f9de", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:nginx_server:properties:dataplane_ref", "xcsh-docs:data-sources:nginx_server:properties:server_spec"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:nginx_server:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nginx_server:reference", "parent_id": "xcsh-docs:data-sources:nginx_server:fundamentals", "path": "documentation/data-sources/nginx_server/properties/index.md", "product": "distributed-cloud", "provider_name": "nginx_server", "provider_schema_digest": "sha256:057968f86e4ef0ae0087dd4d6097131e285b60998d97655ee1a02c00315f6b0f", "provider_type": "data-sources", "registry_anchor": "canonical-1023323233213133-0302133111112113-0100222132220233-3300220100330002-1122012022213202-1333110201013021-3010020312201133-3123122020002010", "registry_path": "docs/guides/data-sources--nginx_server--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations.", "document_id": "xcsh-docs:data-sources:nginx_server:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["dataplane ref"], "anchor": "section", "description": "DataplaneReference.", "document_id": "xcsh-docs:data-sources:nginx_server:properties:dataplane_ref", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["dataplane_ref"], "syntax": "attribute", "type": "object"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Description.", "document_id": "xcsh-docs:data-sources:nginx_server:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier.", "document_id": "xcsh-docs:data-sources:nginx_server:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Labels.", "document_id": "xcsh-docs:data-sources:nginx_server:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["name"], "anchor": "schema-name", "description": "Name of the NginxServer to look up.", "document_id": "xcsh-docs:data-sources:nginx_server:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "Namespace of the NginxServer.", "document_id": "xcsh-docs:data-sources:nginx_server:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["server spec"], "anchor": "section", "description": "Configuration for server_spec.", "document_id": "xcsh-docs:data-sources:nginx_server:properties:server_spec", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["server_spec"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nginx_server/properties/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Property reference for xcsh_nginx_server.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": [], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_nginx_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/)
- Property reference

## Direct properties

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Computed.

Annotations.

- [dataplane_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/dataplane_ref/): complete subsection reference.

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

Name of the NginxServer to look up.

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Required.

Namespace of the NginxServer.

- [server_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/server_spec/): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/#schema-annotations) |
| `dataplane_ref` | [dataplane_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/dataplane_ref/#section) |
| `dataplane_ref.nginx_csg` | [dataplane_ref.nginx_csg](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/dataplane_ref/nginx_csg/#section) |
| `dataplane_ref.nginx_csg.name` | [dataplane_ref.nginx_csg.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/dataplane_ref/nginx_csg/#schema-dataplane_ref--nginx_csg--name) |
| `dataplane_ref.nginx_csg.namespace` | [dataplane_ref.nginx_csg.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/dataplane_ref/nginx_csg/#schema-dataplane_ref--nginx_csg--namespace) |
| `dataplane_ref.nginx_csg.tenant` | [dataplane_ref.nginx_csg.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/dataplane_ref/nginx_csg/#schema-dataplane_ref--nginx_csg--tenant) |
| `dataplane_ref.nginx_instance` | [dataplane_ref.nginx_instance](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/dataplane_ref/nginx_instance/#section) |
| `dataplane_ref.nginx_instance.name` | [dataplane_ref.nginx_instance.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/dataplane_ref/nginx_instance/#schema-dataplane_ref--nginx_instance--name) |
| `dataplane_ref.nginx_instance.namespace` | [dataplane_ref.nginx_instance.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/dataplane_ref/nginx_instance/#schema-dataplane_ref--nginx_instance--namespace) |
| `dataplane_ref.nginx_instance.tenant` | [dataplane_ref.nginx_instance.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/dataplane_ref/nginx_instance/#schema-dataplane_ref--nginx_instance--tenant) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/#schema-description) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/#schema-id) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/#schema-namespace) |
| `server_spec` | [server_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/server_spec/#section) |
| `server_spec.api_discovery_spec` | [server_spec.api_discovery_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/server_spec/api_discovery_spec/#section) |
| `server_spec.api_discovery_spec.disabled` | [server_spec.api_discovery_spec.disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/server_spec/api_discovery_spec/disabled/#section) |
| `server_spec.api_discovery_spec.enabled` | [server_spec.api_discovery_spec.enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/server_spec/api_discovery_spec/enabled/#section) |
| `server_spec.domains` | [server_spec.domains](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/server_spec/#schema-server_spec--domains) |
| `server_spec.locations` | [server_spec.locations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/server_spec/locations/#section) |
| `server_spec.locations.api_discovery_spec` | [server_spec.locations.api_discovery_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/server_spec/locations/api_discovery_spec/#section) |
| `server_spec.locations.api_discovery_spec.disabled` | [server_spec.locations.api_discovery_spec.disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/server_spec/locations/api_discovery_spec/disabled/#section) |
| `server_spec.locations.api_discovery_spec.enabled` | [server_spec.locations.api_discovery_spec.enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/server_spec/locations/api_discovery_spec/enabled/#section) |
| `server_spec.locations.definition` | [server_spec.locations.definition](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/server_spec/locations/#schema-server_spec--locations--definition) |
| `server_spec.locations.name` | [server_spec.locations.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/server_spec/locations/#schema-server_spec--locations--name) |
| `server_spec.locations.waf_spec` | [server_spec.locations.waf_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/server_spec/locations/waf_spec/#section) |
| `server_spec.locations.waf_spec.blocking_waf_mode` | [server_spec.locations.waf_spec.blocking_waf_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/server_spec/locations/waf_spec/blocking_waf_mode/#section) |
| `server_spec.locations.waf_spec.distributed_cloud_policy_management` | [server_spec.locations.waf_spec.distributed_cloud_policy_management](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/server_spec/locations/waf_spec/distributed_cloud_policy_management/#section) |
| `server_spec.locations.waf_spec.monitoring_waf_mode` | [server_spec.locations.waf_spec.monitoring_waf_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/server_spec/locations/waf_spec/monitoring_waf_mode/#section) |
| `server_spec.locations.waf_spec.nginx_policy_management` | [server_spec.locations.waf_spec.nginx_policy_management](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/server_spec/locations/waf_spec/nginx_policy_management/#section) |
| `server_spec.locations.waf_spec.none_waf_mode` | [server_spec.locations.waf_spec.none_waf_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/server_spec/locations/waf_spec/none_waf_mode/#section) |
| `server_spec.locations.waf_spec.policy_file_name` | [server_spec.locations.waf_spec.policy_file_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/server_spec/locations/waf_spec/#schema-server_spec--locations--waf_spec--policy_file_name) |
| `server_spec.locations.waf_spec.policy_name` | [server_spec.locations.waf_spec.policy_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/server_spec/locations/waf_spec/#schema-server_spec--locations--waf_spec--policy_name) |
| `server_spec.locations.waf_spec.security_log_enabled` | [server_spec.locations.waf_spec.security_log_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/server_spec/locations/waf_spec/#schema-server_spec--locations--waf_spec--security_log_enabled) |
| `server_spec.locations.waf_spec.security_log_file_names` | [server_spec.locations.waf_spec.security_log_file_names](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/server_spec/locations/waf_spec/#schema-server_spec--locations--waf_spec--security_log_file_names) |
| `server_spec.nginx_one_object_id` | [server_spec.nginx_one_object_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/server_spec/#schema-server_spec--nginx_one_object_id) |
| `server_spec.nginx_one_object_name` | [server_spec.nginx_one_object_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/server_spec/#schema-server_spec--nginx_one_object_name) |
| `server_spec.port` | [server_spec.port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/server_spec/#schema-server_spec--port) |
| `server_spec.server_name` | [server_spec.server_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/server_spec/#schema-server_spec--server_name) |
| `server_spec.total_routes` | [server_spec.total_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/server_spec/#schema-server_spec--total_routes) |
| `server_spec.waf_spec` | [server_spec.waf_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/server_spec/waf_spec/#section) |
| `server_spec.waf_spec.blocking_waf_mode` | [server_spec.waf_spec.blocking_waf_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/server_spec/waf_spec/blocking_waf_mode/#section) |
| `server_spec.waf_spec.distributed_cloud_policy_management` | [server_spec.waf_spec.distributed_cloud_policy_management](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/server_spec/waf_spec/distributed_cloud_policy_management/#section) |
| `server_spec.waf_spec.monitoring_waf_mode` | [server_spec.waf_spec.monitoring_waf_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/server_spec/waf_spec/monitoring_waf_mode/#section) |
| `server_spec.waf_spec.nginx_policy_management` | [server_spec.waf_spec.nginx_policy_management](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/server_spec/waf_spec/nginx_policy_management/#section) |
| `server_spec.waf_spec.none_waf_mode` | [server_spec.waf_spec.none_waf_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/server_spec/waf_spec/none_waf_mode/#section) |
| `server_spec.waf_spec.policy_file_name` | [server_spec.waf_spec.policy_file_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/server_spec/waf_spec/#schema-server_spec--waf_spec--policy_file_name) |
| `server_spec.waf_spec.policy_name` | [server_spec.waf_spec.policy_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/server_spec/waf_spec/#schema-server_spec--waf_spec--policy_name) |
| `server_spec.waf_spec.security_log_enabled` | [server_spec.waf_spec.security_log_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/server_spec/waf_spec/#schema-server_spec--waf_spec--security_log_enabled) |
| `server_spec.waf_spec.security_log_file_names` | [server_spec.waf_spec.security_log_file_names](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/server_spec/waf_spec/#schema-server_spec--waf_spec--security_log_file_names) |
