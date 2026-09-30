---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_nginx_instance."
xcsh_docs: {"aliases": [], "body_bytes": 5717, "body_sha256": "sha256:b608dc926e2c89a7bd61959801234caa9dd93832e2543d1c9c238e906dc9b891", "child_ids": ["xcsh-docs:data-sources:nginx_instance:properties:api_discovery_spec", "xcsh-docs:data-sources:nginx_instance:properties:waf_spec"], "collection_id": "xcsh-docs:data-sources:nginx_instance:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nginx_instance:reference", "parent_id": "xcsh-docs:data-sources:nginx_instance:fundamentals", "path": "documentation/data-sources/nginx_instance/properties/index.md", "provider_name": "nginx_instance", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nginx_instance/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_nginx_instance.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

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
