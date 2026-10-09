---
page_title: "waf_spec"
subcategory: ""
description: "Configuration for waf_spec."
xcsh_docs: {"aliases": ["waf spec"], "body_bytes": 2059, "body_sha256": "sha256:1a04d21390756a25b5bc3feda4d3521404304dc4d3b86b56b1c9318ad7ab5fbe", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:nginx_instance:properties:waf_spec:blocking_waf_mode", "xcsh-docs:data-sources:nginx_instance:properties:waf_spec:distributed_cloud_policy_management", "xcsh-docs:data-sources:nginx_instance:properties:waf_spec:monitoring_waf_mode", "xcsh-docs:data-sources:nginx_instance:properties:waf_spec:nginx_policy_management", "xcsh-docs:data-sources:nginx_instance:properties:waf_spec:none_waf_mode"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:nginx_instance:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nginx_instance:properties:waf_spec", "parent_id": "xcsh-docs:data-sources:nginx_instance:reference", "path": "documentation/data-sources/nginx_instance/properties/waf_spec/index.md", "product": "distributed-cloud", "provider_name": "nginx_instance", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-1000133302301320-1012132213300332-3020123020312312-3222223313313122-3100212223131120-1020021333003101-3123333122231023-2221001301210100", "registry_path": "docs/guides/data-sources--nginx_instance--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["waf_spec"], "schema_version": 1, "sections": [{"aliases": ["waf spec blocking waf mode"], "anchor": "section", "description": "Enable this option", "document_id": "xcsh-docs:data-sources:nginx_instance:properties:waf_spec:blocking_waf_mode", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["waf_spec", "blocking_waf_mode"], "syntax": "attribute", "type": "object"}, {"aliases": ["waf spec distributed cloud policy management"], "anchor": "section", "description": "Configuration parameter for distributed cloud policy management.", "document_id": "xcsh-docs:data-sources:nginx_instance:properties:waf_spec:distributed_cloud_policy_management", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["waf_spec", "distributed_cloud_policy_management"], "syntax": "attribute", "type": "object"}, {"aliases": ["waf spec monitoring waf mode"], "anchor": "section", "description": "Configuration parameter for monitoring waf mode.", "document_id": "xcsh-docs:data-sources:nginx_instance:properties:waf_spec:monitoring_waf_mode", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["waf_spec", "monitoring_waf_mode"], "syntax": "attribute", "type": "object"}, {"aliases": ["waf spec nginx policy management"], "anchor": "section", "description": "Configuration parameter for nginx policy management.", "document_id": "xcsh-docs:data-sources:nginx_instance:properties:waf_spec:nginx_policy_management", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["waf_spec", "nginx_policy_management"], "syntax": "attribute", "type": "object"}, {"aliases": ["waf spec none waf mode"], "anchor": "section", "description": "Configuration parameter for none waf mode.", "document_id": "xcsh-docs:data-sources:nginx_instance:properties:waf_spec:none_waf_mode", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["waf_spec", "none_waf_mode"], "syntax": "attribute", "type": "object"}, {"aliases": ["waf spec policy file name"], "anchor": "schema-waf_spec--policy_file_name", "description": "WAF Policy File Name. Policy file name for WAF.", "document_id": "xcsh-docs:data-sources:nginx_instance:properties:waf_spec", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["waf_spec", "policy_file_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["waf spec policy name"], "anchor": "schema-waf_spec--policy_name", "description": "WAF Policy Name. Policy name configured for WAF.", "document_id": "xcsh-docs:data-sources:nginx_instance:properties:waf_spec", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["waf_spec", "policy_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["waf spec security log enabled"], "anchor": "schema-waf_spec--security_log_enabled", "description": "Specifies if security logging is enabled.", "document_id": "xcsh-docs:data-sources:nginx_instance:properties:waf_spec", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["waf_spec", "security_log_enabled"], "syntax": "attribute", "type": "bool"}, {"aliases": ["waf spec security log file names"], "anchor": "schema-waf_spec--security_log_file_names", "description": "Specifies the list of security log files specification.", "document_id": "xcsh-docs:data-sources:nginx_instance:properties:waf_spec", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["waf_spec", "security_log_file_names"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nginx_instance/properties/waf_spec/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Configuration for waf_spec.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": [], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# waf_spec

Breadcrumbs:

- [xcsh_nginx_instance](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_instance/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_instance/properties/)
- waf_spec

<a id="section"></a>

Type: `"single"`. Computed.

Configuration for waf\_spec.

## Direct properties

- [blocking_waf_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_instance/properties/waf_spec/blocking_waf_mode/): complete subsection reference.

- [distributed_cloud_policy_management](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_instance/properties/waf_spec/distributed_cloud_policy_management/): complete subsection reference.

- [monitoring_waf_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_instance/properties/waf_spec/monitoring_waf_mode/): complete subsection reference.

- [nginx_policy_management](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_instance/properties/waf_spec/nginx_policy_management/): complete subsection reference.

- [none_waf_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_instance/properties/waf_spec/none_waf_mode/): complete subsection reference.

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
