---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_nginx_csg."
xcsh_docs: {"aliases": [], "body_bytes": 4251, "body_sha256": "sha256:9d6911f669ed3fb7f2d569a163d79b6dfc4b0d5956dbe2a0bf0553598deb4fda", "canonical_id": "xcsh-docs:data-sources:nginx_csg:reference", "child_ids": ["xcsh-docs:data-sources:nginx_csg:properties:api_discovery_spec", "xcsh-docs:data-sources:nginx_csg:properties:waf_spec"], "collection_id": "xcsh-docs:data-sources:nginx_csg:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nginx_csg:reference", "parent_id": "xcsh-docs:data-sources:nginx_csg:fundamentals", "path": "docs/guides/data-sources--nginx_csg--reference.md", "provider_name": "nginx_csg", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nginx_csg/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_nginx_csg.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_nginx_csg](../data-sources/nginx_csg.md)
- Property reference

## Direct properties

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Computed.

Annotations.

- [api_discovery_spec](data-sources--nginx_csg--properties--api_discovery_spec.md): complete subsection reference.

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

- [waf_spec](data-sources--nginx_csg--properties--waf_spec.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--nginx_csg--reference.md#schema-annotations) |
| `api_discovery_spec` | [api_discovery_spec](data-sources--nginx_csg--properties--api_discovery_spec.md#section) |
| `api_discovery_spec.disabled` | [api_discovery_spec.disabled](data-sources--nginx_csg--properties--api_discovery_spec--disabled.md#section) |
| `api_discovery_spec.enabled` | [api_discovery_spec.enabled](data-sources--nginx_csg--properties--api_discovery_spec--enabled.md#section) |
| `csg_name` | [csg_name](data-sources--nginx_csg--reference.md#schema-csg_name) |
| `description` | [description](data-sources--nginx_csg--reference.md#schema-description) |
| `id` | [id](data-sources--nginx_csg--reference.md#schema-id) |
| `labels` | [labels](data-sources--nginx_csg--reference.md#schema-labels) |
| `name` | [name](data-sources--nginx_csg--reference.md#schema-name) |
| `namespace` | [namespace](data-sources--nginx_csg--reference.md#schema-namespace) |
| `object_id` | [object_id](data-sources--nginx_csg--reference.md#schema-object_id) |
| `waf_spec` | [waf_spec](data-sources--nginx_csg--properties--waf_spec.md#section) |
| `waf_spec.blocking_waf_mode` | [waf_spec.blocking_waf_mode](data-sources--nginx_csg--properties--waf_spec--blocking_waf_mode.md#section) |
| `waf_spec.distributed_cloud_policy_management` | [waf_spec.distributed_cloud_policy_management](data-sources--nginx_csg--properties--waf_spec--distributed_cloud_policy_management.md#section) |
| `waf_spec.monitoring_waf_mode` | [waf_spec.monitoring_waf_mode](data-sources--nginx_csg--properties--waf_spec--monitoring_waf_mode.md#section) |
| `waf_spec.nginx_policy_management` | [waf_spec.nginx_policy_management](data-sources--nginx_csg--properties--waf_spec--nginx_policy_management.md#section) |
| `waf_spec.none_waf_mode` | [waf_spec.none_waf_mode](data-sources--nginx_csg--properties--waf_spec--none_waf_mode.md#section) |
| `waf_spec.policy_file_name` | [waf_spec.policy_file_name](data-sources--nginx_csg--properties--waf_spec.md#schema-waf_spec--policy_file_name) |
| `waf_spec.policy_name` | [waf_spec.policy_name](data-sources--nginx_csg--properties--waf_spec.md#schema-waf_spec--policy_name) |
| `waf_spec.security_log_enabled` | [waf_spec.security_log_enabled](data-sources--nginx_csg--properties--waf_spec.md#schema-waf_spec--security_log_enabled) |
| `waf_spec.security_log_file_names` | [waf_spec.security_log_file_names](data-sources--nginx_csg--properties--waf_spec.md#schema-waf_spec--security_log_file_names) |

## Next pages

- [api_discovery_spec](data-sources--nginx_csg--properties--api_discovery_spec.md)
- [waf_spec](data-sources--nginx_csg--properties--waf_spec.md)
- [xcsh_nginx_csg](../data-sources/nginx_csg.md)
