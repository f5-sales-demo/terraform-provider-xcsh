---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_nginx_server."
xcsh_docs: {"aliases": [], "body_bytes": 10342, "body_sha256": "sha256:73d3cdaad82954a1126166179162c2f46b10814e8e7a904bae48411f155babbb", "canonical_id": "xcsh-docs:data-sources:nginx_server:reference", "child_ids": ["xcsh-docs:data-sources:nginx_server:properties:dataplane_ref", "xcsh-docs:data-sources:nginx_server:properties:server_spec"], "collection_id": "xcsh-docs:data-sources:nginx_server:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nginx_server:reference", "parent_id": "xcsh-docs:data-sources:nginx_server:fundamentals", "path": "docs/guides/data-sources--nginx_server--reference.md", "provider_name": "nginx_server", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nginx_server/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_nginx_server.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_nginx_server](../data-sources/nginx_server.md)
- Property reference

## Direct properties

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Computed.

Annotations.

- [dataplane_ref](data-sources--nginx_server--properties--dataplane_ref.md): complete subsection reference.

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

- [server_spec](data-sources--nginx_server--properties--server_spec.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--nginx_server--reference.md#schema-annotations) |
| `dataplane_ref` | [dataplane_ref](data-sources--nginx_server--properties--dataplane_ref.md#section) |
| `dataplane_ref.nginx_csg` | [dataplane_ref.nginx_csg](data-sources--nginx_server--properties--dataplane_ref--nginx_csg.md#section) |
| `dataplane_ref.nginx_csg.name` | [dataplane_ref.nginx_csg.name](data-sources--nginx_server--properties--dataplane_ref--nginx_csg.md#schema-dataplane_ref--nginx_csg--name) |
| `dataplane_ref.nginx_csg.namespace` | [dataplane_ref.nginx_csg.namespace](data-sources--nginx_server--properties--dataplane_ref--nginx_csg.md#schema-dataplane_ref--nginx_csg--namespace) |
| `dataplane_ref.nginx_csg.tenant` | [dataplane_ref.nginx_csg.tenant](data-sources--nginx_server--properties--dataplane_ref--nginx_csg.md#schema-dataplane_ref--nginx_csg--tenant) |
| `dataplane_ref.nginx_instance` | [dataplane_ref.nginx_instance](data-sources--nginx_server--properties--dataplane_ref--nginx_instance.md#section) |
| `dataplane_ref.nginx_instance.name` | [dataplane_ref.nginx_instance.name](data-sources--nginx_server--properties--dataplane_ref--nginx_instance.md#schema-dataplane_ref--nginx_instance--name) |
| `dataplane_ref.nginx_instance.namespace` | [dataplane_ref.nginx_instance.namespace](data-sources--nginx_server--properties--dataplane_ref--nginx_instance.md#schema-dataplane_ref--nginx_instance--namespace) |
| `dataplane_ref.nginx_instance.tenant` | [dataplane_ref.nginx_instance.tenant](data-sources--nginx_server--properties--dataplane_ref--nginx_instance.md#schema-dataplane_ref--nginx_instance--tenant) |
| `description` | [description](data-sources--nginx_server--reference.md#schema-description) |
| `id` | [id](data-sources--nginx_server--reference.md#schema-id) |
| `labels` | [labels](data-sources--nginx_server--reference.md#schema-labels) |
| `name` | [name](data-sources--nginx_server--reference.md#schema-name) |
| `namespace` | [namespace](data-sources--nginx_server--reference.md#schema-namespace) |
| `server_spec` | [server_spec](data-sources--nginx_server--properties--server_spec.md#section) |
| `server_spec.api_discovery_spec` | [server_spec.api_discovery_spec](data-sources--nginx_server--properties--server_spec--api_discovery_spec.md#section) |
| `server_spec.api_discovery_spec.disabled` | [server_spec.api_discovery_spec.disabled](data-sources--nginx_server--properties--server_spec--api_discovery_spec--disabled.md#section) |
| `server_spec.api_discovery_spec.enabled` | [server_spec.api_discovery_spec.enabled](data-sources--nginx_server--properties--server_spec--api_discovery_spec--enabled.md#section) |
| `server_spec.domains` | [server_spec.domains](data-sources--nginx_server--properties--server_spec.md#schema-server_spec--domains) |
| `server_spec.locations` | [server_spec.locations](data-sources--nginx_server--properties--server_spec--locations.md#section) |
| `server_spec.locations.api_discovery_spec` | [server_spec.locations.api_discovery_spec](data-sources--nginx_server--properties--server_spec--locations--api_discovery_spec.md#section) |
| `server_spec.locations.api_discovery_spec.disabled` | [server_spec.locations.api_discovery_spec.disabled](data-sources--nginx_server--properties--server_spec--locations--api_discovery_spec--disabled.md#section) |
| `server_spec.locations.api_discovery_spec.enabled` | [server_spec.locations.api_discovery_spec.enabled](data-sources--nginx_server--properties--server_spec--locations--api_discovery_spec--enabled.md#section) |
| `server_spec.locations.definition` | [server_spec.locations.definition](data-sources--nginx_server--properties--server_spec--locations.md#schema-server_spec--locations--definition) |
| `server_spec.locations.name` | [server_spec.locations.name](data-sources--nginx_server--properties--server_spec--locations.md#schema-server_spec--locations--name) |
| `server_spec.locations.waf_spec` | [server_spec.locations.waf_spec](data-sources--nginx_server--properties--server_spec--locations--waf_spec.md#section) |
| `server_spec.locations.waf_spec.blocking_waf_mode` | [server_spec.locations.waf_spec.blocking_waf_mode](data-sources--nginx_server--properties--server_spec--locations--waf_spec--blocking_waf_mode.md#section) |
| `server_spec.locations.waf_spec.distributed_cloud_policy_management` | [server_spec.locations.waf_spec.distributed_cloud_policy_management](data-sources--nginx_server--properties--server_spec--locations--waf_spec--distributed_cloud_policy_management.md#section) |
| `server_spec.locations.waf_spec.monitoring_waf_mode` | [server_spec.locations.waf_spec.monitoring_waf_mode](data-sources--nginx_server--properties--server_spec--locations--waf_spec--monitoring_waf_mode.md#section) |
| `server_spec.locations.waf_spec.nginx_policy_management` | [server_spec.locations.waf_spec.nginx_policy_management](data-sources--nginx_server--properties--server_spec--locations--waf_spec--nginx_policy_management.md#section) |
| `server_spec.locations.waf_spec.none_waf_mode` | [server_spec.locations.waf_spec.none_waf_mode](data-sources--nginx_server--properties--server_spec--locations--waf_spec--none_waf_mode.md#section) |
| `server_spec.locations.waf_spec.policy_file_name` | [server_spec.locations.waf_spec.policy_file_name](data-sources--nginx_server--properties--server_spec--locations--waf_spec.md#schema-server_spec--locations--waf_spec--policy_file_name) |
| `server_spec.locations.waf_spec.policy_name` | [server_spec.locations.waf_spec.policy_name](data-sources--nginx_server--properties--server_spec--locations--waf_spec.md#schema-server_spec--locations--waf_spec--policy_name) |
| `server_spec.locations.waf_spec.security_log_enabled` | [server_spec.locations.waf_spec.security_log_enabled](data-sources--nginx_server--properties--server_spec--locations--waf_spec.md#schema-server_spec--locations--waf_spec--security_log_enabled) |
| `server_spec.locations.waf_spec.security_log_file_names` | [server_spec.locations.waf_spec.security_log_file_names](data-sources--nginx_server--properties--server_spec--locations--waf_spec.md#schema-server_spec--locations--waf_spec--security_log_file_names) |
| `server_spec.nginx_one_object_id` | [server_spec.nginx_one_object_id](data-sources--nginx_server--properties--server_spec.md#schema-server_spec--nginx_one_object_id) |
| `server_spec.nginx_one_object_name` | [server_spec.nginx_one_object_name](data-sources--nginx_server--properties--server_spec.md#schema-server_spec--nginx_one_object_name) |
| `server_spec.port` | [server_spec.port](data-sources--nginx_server--properties--server_spec.md#schema-server_spec--port) |
| `server_spec.server_name` | [server_spec.server_name](data-sources--nginx_server--properties--server_spec.md#schema-server_spec--server_name) |
| `server_spec.total_routes` | [server_spec.total_routes](data-sources--nginx_server--properties--server_spec.md#schema-server_spec--total_routes) |
| `server_spec.waf_spec` | [server_spec.waf_spec](data-sources--nginx_server--properties--server_spec--waf_spec.md#section) |
| `server_spec.waf_spec.blocking_waf_mode` | [server_spec.waf_spec.blocking_waf_mode](data-sources--nginx_server--properties--server_spec--waf_spec--blocking_waf_mode.md#section) |
| `server_spec.waf_spec.distributed_cloud_policy_management` | [server_spec.waf_spec.distributed_cloud_policy_management](data-sources--nginx_server--properties--server_spec--waf_spec--distributed_cloud_policy_management.md#section) |
| `server_spec.waf_spec.monitoring_waf_mode` | [server_spec.waf_spec.monitoring_waf_mode](data-sources--nginx_server--properties--server_spec--waf_spec--monitoring_waf_mode.md#section) |
| `server_spec.waf_spec.nginx_policy_management` | [server_spec.waf_spec.nginx_policy_management](data-sources--nginx_server--properties--server_spec--waf_spec--nginx_policy_management.md#section) |
| `server_spec.waf_spec.none_waf_mode` | [server_spec.waf_spec.none_waf_mode](data-sources--nginx_server--properties--server_spec--waf_spec--none_waf_mode.md#section) |
| `server_spec.waf_spec.policy_file_name` | [server_spec.waf_spec.policy_file_name](data-sources--nginx_server--properties--server_spec--waf_spec.md#schema-server_spec--waf_spec--policy_file_name) |
| `server_spec.waf_spec.policy_name` | [server_spec.waf_spec.policy_name](data-sources--nginx_server--properties--server_spec--waf_spec.md#schema-server_spec--waf_spec--policy_name) |
| `server_spec.waf_spec.security_log_enabled` | [server_spec.waf_spec.security_log_enabled](data-sources--nginx_server--properties--server_spec--waf_spec.md#schema-server_spec--waf_spec--security_log_enabled) |
| `server_spec.waf_spec.security_log_file_names` | [server_spec.waf_spec.security_log_file_names](data-sources--nginx_server--properties--server_spec--waf_spec.md#schema-server_spec--waf_spec--security_log_file_names) |

## Next pages

- [dataplane_ref](data-sources--nginx_server--properties--dataplane_ref.md)
- [server_spec](data-sources--nginx_server--properties--server_spec.md)
- [xcsh_nginx_server](../data-sources/nginx_server.md)
