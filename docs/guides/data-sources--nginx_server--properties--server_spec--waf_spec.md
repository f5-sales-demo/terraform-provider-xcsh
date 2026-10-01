---
page_title: "server_spec.waf_spec"
subcategory: ""
description: "server_spec.waf_spec for xcsh_nginx_server."
xcsh_docs: {"aliases": [], "body_bytes": 2729, "body_sha256": "sha256:ac460b52b7093cb5587d2cb3a1dd1eb2d0bc384a3fb7edd2d6fd1e6df798c037", "canonical_id": "xcsh-docs:data-sources:nginx_server:properties:server_spec:waf_spec", "child_ids": ["xcsh-docs:data-sources:nginx_server:properties:server_spec:waf_spec:blocking_waf_mode", "xcsh-docs:data-sources:nginx_server:properties:server_spec:waf_spec:distributed_cloud_policy_management", "xcsh-docs:data-sources:nginx_server:properties:server_spec:waf_spec:monitoring_waf_mode", "xcsh-docs:data-sources:nginx_server:properties:server_spec:waf_spec:nginx_policy_management", "xcsh-docs:data-sources:nginx_server:properties:server_spec:waf_spec:none_waf_mode"], "collection_id": "xcsh-docs:data-sources:nginx_server:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nginx_server:properties:server_spec:waf_spec", "parent_id": "xcsh-docs:data-sources:nginx_server:properties:server_spec", "path": "docs/guides/data-sources--nginx_server--properties--server_spec--waf_spec.md", "provider_name": "nginx_server", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["server_spec", "waf_spec"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nginx_server/properties/server_spec/waf_spec/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "server_spec.waf_spec for xcsh_nginx_server.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# server_spec.waf_spec

Breadcrumbs:

- [xcsh_nginx_server](../data-sources/nginx_server.md)
- [Property reference](data-sources--nginx_server--reference.md)
- [server_spec](data-sources--nginx_server--properties--server_spec.md)
- server_spec.waf_spec

<a id="section"></a>

Type: `"single"`. Computed.

Configuration for waf\_spec.

## Direct properties

- [blocking_waf_mode](data-sources--nginx_server--properties--server_spec--waf_spec--blocking_waf_mode.md): complete subsection reference.

- [distributed_cloud_policy_management](data-sources--nginx_server--properties--server_spec--waf_spec--distributed_cloud_policy_management.md): complete subsection reference.

- [monitoring_waf_mode](data-sources--nginx_server--properties--server_spec--waf_spec--monitoring_waf_mode.md): complete subsection reference.

- [nginx_policy_management](data-sources--nginx_server--properties--server_spec--waf_spec--nginx_policy_management.md): complete subsection reference.

- [none_waf_mode](data-sources--nginx_server--properties--server_spec--waf_spec--none_waf_mode.md): complete subsection reference.

<a id="schema-server_spec--waf_spec--policy_file_name"></a>

### policy_file_name property

Type: `"string"`. Computed.

WAF Policy File Name. Policy file name for WAF.

<a id="schema-server_spec--waf_spec--policy_name"></a>

### policy_name property

Type: `"string"`. Computed.

WAF Policy Name. Policy name configured for WAF.

<a id="schema-server_spec--waf_spec--security_log_enabled"></a>

### security_log_enabled property

Type: `"bool"`. Computed.

Specifies if security logging is enabled.

<a id="schema-server_spec--waf_spec--security_log_file_names"></a>

### security_log_file_names property

Type: `["list", "string"]`. Computed.

Specifies the list of security log files specification.

## Next pages

- [server_spec.waf_spec.blocking_waf_mode](data-sources--nginx_server--properties--server_spec--waf_spec--blocking_waf_mode.md)
- [server_spec.waf_spec.distributed_cloud_policy_management](data-sources--nginx_server--properties--server_spec--waf_spec--distributed_cloud_policy_management.md)
- [server_spec.waf_spec.monitoring_waf_mode](data-sources--nginx_server--properties--server_spec--waf_spec--monitoring_waf_mode.md)
- [server_spec.waf_spec.nginx_policy_management](data-sources--nginx_server--properties--server_spec--waf_spec--nginx_policy_management.md)
- [server_spec.waf_spec.none_waf_mode](data-sources--nginx_server--properties--server_spec--waf_spec--none_waf_mode.md)
- [server_spec](data-sources--nginx_server--properties--server_spec.md)
- [xcsh_nginx_server](../data-sources/nginx_server.md)
