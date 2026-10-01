---
page_title: "waf_spec"
subcategory: ""
description: "waf_spec for xcsh_nginx_csg."
xcsh_docs: {"aliases": [], "body_bytes": 2336, "body_sha256": "sha256:24a436b1680e7f009d7a5c3b7c01007c07da0814743b6d30f814dd992321d8ee", "canonical_id": "xcsh-docs:data-sources:nginx_csg:properties:waf_spec", "child_ids": ["xcsh-docs:data-sources:nginx_csg:properties:waf_spec:blocking_waf_mode", "xcsh-docs:data-sources:nginx_csg:properties:waf_spec:distributed_cloud_policy_management", "xcsh-docs:data-sources:nginx_csg:properties:waf_spec:monitoring_waf_mode", "xcsh-docs:data-sources:nginx_csg:properties:waf_spec:nginx_policy_management", "xcsh-docs:data-sources:nginx_csg:properties:waf_spec:none_waf_mode"], "collection_id": "xcsh-docs:data-sources:nginx_csg:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nginx_csg:properties:waf_spec", "parent_id": "xcsh-docs:data-sources:nginx_csg:reference", "path": "docs/guides/data-sources--nginx_csg--properties--waf_spec.md", "provider_name": "nginx_csg", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["waf_spec"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nginx_csg/properties/waf_spec/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "waf_spec for xcsh_nginx_csg.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# waf_spec

Breadcrumbs:

- [xcsh_nginx_csg](../data-sources/nginx_csg.md)
- [Property reference](data-sources--nginx_csg--reference.md)
- waf_spec

<a id="section"></a>

Type: `"single"`. Computed.

Configuration for waf\_spec.

## Direct properties

- [blocking_waf_mode](data-sources--nginx_csg--properties--waf_spec--blocking_waf_mode.md): complete subsection reference.

- [distributed_cloud_policy_management](data-sources--nginx_csg--properties--waf_spec--distributed_cloud_policy_management.md): complete subsection reference.

- [monitoring_waf_mode](data-sources--nginx_csg--properties--waf_spec--monitoring_waf_mode.md): complete subsection reference.

- [nginx_policy_management](data-sources--nginx_csg--properties--waf_spec--nginx_policy_management.md): complete subsection reference.

- [none_waf_mode](data-sources--nginx_csg--properties--waf_spec--none_waf_mode.md): complete subsection reference.

<a id="schema-waf_spec--policy_file_name"></a>

### policy_file_name property

Type: `"string"`. Computed.

WAF Policy File Name. Policy file name for WAF.

<a id="schema-waf_spec--policy_name"></a>

### policy_name property

Type: `"string"`. Computed.

WAF Policy Name. Policy name configured for WAF.

<a id="schema-waf_spec--security_log_enabled"></a>

### security_log_enabled property

Type: `"bool"`. Computed.

Specifies if security logging is enabled.

<a id="schema-waf_spec--security_log_file_names"></a>

### security_log_file_names property

Type: `["list", "string"]`. Computed.

Specifies the list of security log files specification.

## Next pages

- [waf_spec.blocking_waf_mode](data-sources--nginx_csg--properties--waf_spec--blocking_waf_mode.md)
- [waf_spec.distributed_cloud_policy_management](data-sources--nginx_csg--properties--waf_spec--distributed_cloud_policy_management.md)
- [waf_spec.monitoring_waf_mode](data-sources--nginx_csg--properties--waf_spec--monitoring_waf_mode.md)
- [waf_spec.nginx_policy_management](data-sources--nginx_csg--properties--waf_spec--nginx_policy_management.md)
- [waf_spec.none_waf_mode](data-sources--nginx_csg--properties--waf_spec--none_waf_mode.md)
- [Property reference](data-sources--nginx_csg--reference.md)
- [xcsh_nginx_csg](../data-sources/nginx_csg.md)
