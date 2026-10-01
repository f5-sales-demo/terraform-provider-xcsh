---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_nginx_instance."
xcsh_docs: {"aliases": [], "body_bytes": 4430, "body_sha256": "sha256:ddf107b5908467dfe8a0aa8db7d980c6ca19865284638f4e649df5e63ba69862", "canonical_id": "xcsh-docs:data-sources:nginx_instance:reference", "child_ids": ["xcsh-docs:data-sources:nginx_instance:properties:api_discovery_spec", "xcsh-docs:data-sources:nginx_instance:properties:waf_spec"], "collection_id": "xcsh-docs:data-sources:nginx_instance:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nginx_instance:reference", "parent_id": "xcsh-docs:data-sources:nginx_instance:fundamentals", "path": "docs/guides/data-sources--nginx_instance--reference.md", "provider_name": "nginx_instance", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nginx_instance/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_nginx_instance.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_nginx_instance](../data-sources/nginx_instance.md)
- Property reference

## Direct properties

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Computed.

Annotations.

- [api_discovery_spec](data-sources--nginx_instance--properties--api_discovery_spec.md): complete subsection reference.

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

- [waf_spec](data-sources--nginx_instance--properties--waf_spec.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--nginx_instance--reference.md#schema-annotations) |
| `api_discovery_spec` | [api_discovery_spec](data-sources--nginx_instance--properties--api_discovery_spec.md#section) |
| `api_discovery_spec.disabled` | [api_discovery_spec.disabled](data-sources--nginx_instance--properties--api_discovery_spec--disabled.md#section) |
| `api_discovery_spec.enabled` | [api_discovery_spec.enabled](data-sources--nginx_instance--properties--api_discovery_spec--enabled.md#section) |
| `description` | [description](data-sources--nginx_instance--reference.md#schema-description) |
| `host_name` | [host_name](data-sources--nginx_instance--reference.md#schema-host_name) |
| `id` | [id](data-sources--nginx_instance--reference.md#schema-id) |
| `labels` | [labels](data-sources--nginx_instance--reference.md#schema-labels) |
| `name` | [name](data-sources--nginx_instance--reference.md#schema-name) |
| `namespace` | [namespace](data-sources--nginx_instance--reference.md#schema-namespace) |
| `object_id` | [object_id](data-sources--nginx_instance--reference.md#schema-object_id) |
| `waf_spec` | [waf_spec](data-sources--nginx_instance--properties--waf_spec.md#section) |
| `waf_spec.blocking_waf_mode` | [waf_spec.blocking_waf_mode](data-sources--nginx_instance--properties--waf_spec--blocking_waf_mode.md#section) |
| `waf_spec.distributed_cloud_policy_management` | [waf_spec.distributed_cloud_policy_management](data-sources--nginx_instance--properties--waf_spec--distributed_cloud_policy_management.md#section) |
| `waf_spec.monitoring_waf_mode` | [waf_spec.monitoring_waf_mode](data-sources--nginx_instance--properties--waf_spec--monitoring_waf_mode.md#section) |
| `waf_spec.nginx_policy_management` | [waf_spec.nginx_policy_management](data-sources--nginx_instance--properties--waf_spec--nginx_policy_management.md#section) |
| `waf_spec.none_waf_mode` | [waf_spec.none_waf_mode](data-sources--nginx_instance--properties--waf_spec--none_waf_mode.md#section) |
| `waf_spec.policy_file_name` | [waf_spec.policy_file_name](data-sources--nginx_instance--properties--waf_spec.md#schema-waf_spec--policy_file_name) |
| `waf_spec.policy_name` | [waf_spec.policy_name](data-sources--nginx_instance--properties--waf_spec.md#schema-waf_spec--policy_name) |
| `waf_spec.security_log_enabled` | [waf_spec.security_log_enabled](data-sources--nginx_instance--properties--waf_spec.md#schema-waf_spec--security_log_enabled) |
| `waf_spec.security_log_file_names` | [waf_spec.security_log_file_names](data-sources--nginx_instance--properties--waf_spec.md#schema-waf_spec--security_log_file_names) |

## Next pages

- [api_discovery_spec](data-sources--nginx_instance--properties--api_discovery_spec.md)
- [waf_spec](data-sources--nginx_instance--properties--waf_spec.md)
- [xcsh_nginx_instance](../data-sources/nginx_instance.md)
