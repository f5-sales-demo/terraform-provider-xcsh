---
page_title: "server_spec.locations.waf_spec"
subcategory: ""
description: "Configuration for waf_spec."
xcsh_docs: {"aliases": ["server spec locations waf spec"], "body_bytes": 2559, "body_sha256": "sha256:5dae8b625b9ef076e55d22eb4b21e3e191b635be237f128c78409c06c2e871e2", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:nginx_server:properties:server_spec:locations:waf_spec:blocking_waf_mode", "xcsh-docs:data-sources:nginx_server:properties:server_spec:locations:waf_spec:distributed_cloud_policy_management", "xcsh-docs:data-sources:nginx_server:properties:server_spec:locations:waf_spec:monitoring_waf_mode", "xcsh-docs:data-sources:nginx_server:properties:server_spec:locations:waf_spec:nginx_policy_management", "xcsh-docs:data-sources:nginx_server:properties:server_spec:locations:waf_spec:none_waf_mode"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:nginx_server:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nginx_server:properties:server_spec:locations:waf_spec", "parent_id": "xcsh-docs:data-sources:nginx_server:properties:server_spec:locations", "path": "documentation/data-sources/nginx_server/properties/server_spec/locations/waf_spec/index.md", "product": "distributed-cloud", "provider_name": "nginx_server", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "data-sources", "registry_anchor": "canonical-3100323131320030-0311003231231131-2022123311132211-3002313310223121-0011022332010330-0121230133121212-2222120301230210-1230013312011200", "registry_path": "docs/guides/data-sources--nginx_server--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["server_spec", "locations", "waf_spec"], "schema_version": 1, "sections": [{"aliases": ["server spec locations waf spec blocking waf mode"], "anchor": "section", "description": "Enable this option", "document_id": "xcsh-docs:data-sources:nginx_server:properties:server_spec:locations:waf_spec:blocking_waf_mode", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["server_spec", "locations", "waf_spec", "blocking_waf_mode"], "syntax": "attribute", "type": "object"}, {"aliases": ["server spec locations waf spec distributed cloud policy management"], "anchor": "section", "description": "Configuration parameter for distributed cloud policy management.", "document_id": "xcsh-docs:data-sources:nginx_server:properties:server_spec:locations:waf_spec:distributed_cloud_policy_management", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["server_spec", "locations", "waf_spec", "distributed_cloud_policy_management"], "syntax": "attribute", "type": "object"}, {"aliases": ["server spec locations waf spec monitoring waf mode"], "anchor": "section", "description": "Configuration parameter for monitoring waf mode.", "document_id": "xcsh-docs:data-sources:nginx_server:properties:server_spec:locations:waf_spec:monitoring_waf_mode", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["server_spec", "locations", "waf_spec", "monitoring_waf_mode"], "syntax": "attribute", "type": "object"}, {"aliases": ["server spec locations waf spec nginx policy management"], "anchor": "section", "description": "Configuration parameter for nginx policy management.", "document_id": "xcsh-docs:data-sources:nginx_server:properties:server_spec:locations:waf_spec:nginx_policy_management", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["server_spec", "locations", "waf_spec", "nginx_policy_management"], "syntax": "attribute", "type": "object"}, {"aliases": ["server spec locations waf spec none waf mode"], "anchor": "section", "description": "Configuration parameter for none waf mode.", "document_id": "xcsh-docs:data-sources:nginx_server:properties:server_spec:locations:waf_spec:none_waf_mode", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["server_spec", "locations", "waf_spec", "none_waf_mode"], "syntax": "attribute", "type": "object"}, {"aliases": ["server spec locations waf spec policy file name"], "anchor": "schema-server_spec--locations--waf_spec--policy_file_name", "description": "WAF Policy File Name. Policy file name for WAF.", "document_id": "xcsh-docs:data-sources:nginx_server:properties:server_spec:locations:waf_spec", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["server_spec", "locations", "waf_spec", "policy_file_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["server spec locations waf spec policy name"], "anchor": "schema-server_spec--locations--waf_spec--policy_name", "description": "WAF Policy Name. Policy name configured for WAF.", "document_id": "xcsh-docs:data-sources:nginx_server:properties:server_spec:locations:waf_spec", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["server_spec", "locations", "waf_spec", "policy_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["server spec locations waf spec security log enabled"], "anchor": "schema-server_spec--locations--waf_spec--security_log_enabled", "description": "Specifies if security logging is enabled.", "document_id": "xcsh-docs:data-sources:nginx_server:properties:server_spec:locations:waf_spec", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["server_spec", "locations", "waf_spec", "security_log_enabled"], "syntax": "attribute", "type": "bool"}, {"aliases": ["server spec locations waf spec security log file names"], "anchor": "schema-server_spec--locations--waf_spec--security_log_file_names", "description": "Specifies the list of security log files specification.", "document_id": "xcsh-docs:data-sources:nginx_server:properties:server_spec:locations:waf_spec", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["server_spec", "locations", "waf_spec", "security_log_file_names"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nginx_server/properties/server_spec/locations/waf_spec/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Configuration for waf_spec.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": [], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# server_spec.locations.waf_spec

Breadcrumbs:

- [xcsh_nginx_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/)
- [server_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/server_spec/)
- [server_spec.locations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/server_spec/locations/)
- server_spec.locations.waf_spec

<a id="section"></a>

Type: `"single"`. Computed.

Configuration for waf\_spec.

## Direct properties

- [blocking_waf_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/server_spec/locations/waf_spec/blocking_waf_mode/): complete subsection reference.

- [distributed_cloud_policy_management](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/server_spec/locations/waf_spec/distributed_cloud_policy_management/): complete subsection reference.

- [monitoring_waf_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/server_spec/locations/waf_spec/monitoring_waf_mode/): complete subsection reference.

- [nginx_policy_management](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/server_spec/locations/waf_spec/nginx_policy_management/): complete subsection reference.

- [none_waf_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/server_spec/locations/waf_spec/none_waf_mode/): complete subsection reference.

<a id="schema-server_spec--locations--waf_spec--policy_file_name"></a>

### policy_file_name property

Type: `"string"`. Computed.

WAF Policy File Name. Policy file name for WAF.

<a id="schema-server_spec--locations--waf_spec--policy_name"></a>

### policy_name property

Type: `"string"`. Computed.

WAF Policy Name. Policy name configured for WAF.

<a id="schema-server_spec--locations--waf_spec--security_log_enabled"></a>

### security_log_enabled property

Type: `"bool"`. Computed.

Specifies if security logging is enabled.

<a id="schema-server_spec--locations--waf_spec--security_log_file_names"></a>

### security_log_file_names property

Type: `["list", "string"]`. Computed.

Specifies the list of security log files specification.
