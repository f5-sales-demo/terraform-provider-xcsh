---
page_title: "server_spec.waf_spec"
subcategory: ""
description: "Configuration for waf_spec."
xcsh_docs: {"aliases": ["server spec waf spec"], "body_bytes": 3476, "body_sha256": "sha256:b21572d1db85f06c2bb837e36250339ee107c02789080dcb2d717b918fe8b860", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:nginx_server:properties:server_spec:waf_spec:blocking_waf_mode", "xcsh-docs:data-sources:nginx_server:properties:server_spec:waf_spec:distributed_cloud_policy_management", "xcsh-docs:data-sources:nginx_server:properties:server_spec:waf_spec:monitoring_waf_mode", "xcsh-docs:data-sources:nginx_server:properties:server_spec:waf_spec:nginx_policy_management", "xcsh-docs:data-sources:nginx_server:properties:server_spec:waf_spec:none_waf_mode"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:nginx_server:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nginx_server:properties:server_spec:waf_spec", "parent_id": "xcsh-docs:data-sources:nginx_server:properties:server_spec", "path": "documentation/data-sources/nginx_server/properties/server_spec/waf_spec/index.md", "product": "distributed-cloud", "provider_name": "nginx_server", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "data-sources", "registry_anchor": "canonical-2133233233002331-2321212110002331-3230230001313032-2231310333323203-2302213001222100-1122302111002033-0001200213302303-2222333001023232", "registry_path": "docs/guides/data-sources--nginx_server--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["server_spec", "waf_spec"], "schema_version": 1, "sections": [{"aliases": ["blocking waf mode"], "anchor": "section", "description": "Enable this option", "document_id": "xcsh-docs:data-sources:nginx_server:properties:server_spec:waf_spec:blocking_waf_mode", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["server_spec", "waf_spec", "blocking_waf_mode"], "syntax": "attribute", "type": "object"}, {"aliases": ["distributed cloud policy management"], "anchor": "section", "description": "Configuration parameter for distributed cloud policy management.", "document_id": "xcsh-docs:data-sources:nginx_server:properties:server_spec:waf_spec:distributed_cloud_policy_management", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["server_spec", "waf_spec", "distributed_cloud_policy_management"], "syntax": "attribute", "type": "object"}, {"aliases": ["monitoring waf mode"], "anchor": "section", "description": "Configuration parameter for monitoring waf mode.", "document_id": "xcsh-docs:data-sources:nginx_server:properties:server_spec:waf_spec:monitoring_waf_mode", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["server_spec", "waf_spec", "monitoring_waf_mode"], "syntax": "attribute", "type": "object"}, {"aliases": ["nginx policy management"], "anchor": "section", "description": "Configuration parameter for nginx policy management.", "document_id": "xcsh-docs:data-sources:nginx_server:properties:server_spec:waf_spec:nginx_policy_management", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["server_spec", "waf_spec", "nginx_policy_management"], "syntax": "attribute", "type": "object"}, {"aliases": ["none waf mode"], "anchor": "section", "description": "Configuration parameter for none waf mode.", "document_id": "xcsh-docs:data-sources:nginx_server:properties:server_spec:waf_spec:none_waf_mode", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["server_spec", "waf_spec", "none_waf_mode"], "syntax": "attribute", "type": "object"}, {"aliases": ["policy file name"], "anchor": "schema-server_spec--waf_spec--policy_file_name", "description": "WAF Policy File Name. Policy file name for WAF.", "document_id": "xcsh-docs:data-sources:nginx_server:properties:server_spec:waf_spec", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["server_spec", "waf_spec", "policy_file_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["policy name"], "anchor": "schema-server_spec--waf_spec--policy_name", "description": "WAF Policy Name. Policy name configured for WAF.", "document_id": "xcsh-docs:data-sources:nginx_server:properties:server_spec:waf_spec", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["server_spec", "waf_spec", "policy_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["security log enabled"], "anchor": "schema-server_spec--waf_spec--security_log_enabled", "description": "Specifies if security logging is enabled.", "document_id": "xcsh-docs:data-sources:nginx_server:properties:server_spec:waf_spec", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["server_spec", "waf_spec", "security_log_enabled"], "syntax": "attribute", "type": "bool"}, {"aliases": ["security log file names"], "anchor": "schema-server_spec--waf_spec--security_log_file_names", "description": "Specifies the list of security log files specification.", "document_id": "xcsh-docs:data-sources:nginx_server:properties:server_spec:waf_spec", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["server_spec", "waf_spec", "security_log_file_names"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nginx_server/properties/server_spec/waf_spec/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Configuration for waf_spec.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# server_spec.waf_spec

Breadcrumbs:

- [xcsh_nginx_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/)
- [server_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/server_spec/)
- server_spec.waf_spec

<a id="section"></a>

Type: `"single"`. Computed.

Configuration for waf\_spec.

## Direct properties

- [blocking_waf_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/server_spec/waf_spec/blocking_waf_mode/): complete subsection reference.

- [distributed_cloud_policy_management](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/server_spec/waf_spec/distributed_cloud_policy_management/): complete subsection reference.

- [monitoring_waf_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/server_spec/waf_spec/monitoring_waf_mode/): complete subsection reference.

- [nginx_policy_management](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/server_spec/waf_spec/nginx_policy_management/): complete subsection reference.

- [none_waf_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/server_spec/waf_spec/none_waf_mode/): complete subsection reference.

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

- [server_spec.waf_spec.blocking_waf_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/server_spec/waf_spec/blocking_waf_mode/)
- [server_spec.waf_spec.distributed_cloud_policy_management](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/server_spec/waf_spec/distributed_cloud_policy_management/)
- [server_spec.waf_spec.monitoring_waf_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/server_spec/waf_spec/monitoring_waf_mode/)
- [server_spec.waf_spec.nginx_policy_management](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/server_spec/waf_spec/nginx_policy_management/)
- [server_spec.waf_spec.none_waf_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/server_spec/waf_spec/none_waf_mode/)
- [server_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/server_spec/)
- [xcsh_nginx_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/)
