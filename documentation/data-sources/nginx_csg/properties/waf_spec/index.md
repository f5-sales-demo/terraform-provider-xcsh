---
page_title: "waf_spec"
subcategory: ""
description: "Configuration for waf_spec."
xcsh_docs: {"aliases": ["waf spec"], "body_bytes": 3044, "body_sha256": "sha256:3d5c2141f37965ed5a00e83a7c3f7f8b94cb3eb1b7ad35b0085b318826e3aac9", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:nginx_csg:properties:waf_spec:blocking_waf_mode", "xcsh-docs:data-sources:nginx_csg:properties:waf_spec:distributed_cloud_policy_management", "xcsh-docs:data-sources:nginx_csg:properties:waf_spec:monitoring_waf_mode", "xcsh-docs:data-sources:nginx_csg:properties:waf_spec:nginx_policy_management", "xcsh-docs:data-sources:nginx_csg:properties:waf_spec:none_waf_mode"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:nginx_csg:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nginx_csg:properties:waf_spec", "parent_id": "xcsh-docs:data-sources:nginx_csg:reference", "path": "documentation/data-sources/nginx_csg/properties/waf_spec/index.md", "product": "distributed-cloud", "provider_name": "nginx_csg", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "data-sources", "registry_anchor": "canonical-0102130320022223-0212231002020230-3202223221211203-2301102230133130-3222220030133133-2312003133032332-3101221121112101-0322232301122131", "registry_path": "docs/guides/data-sources--nginx_csg--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["waf_spec"], "schema_version": 1, "sections": [{"aliases": ["blocking waf mode"], "anchor": "section", "description": "Enable this option", "document_id": "xcsh-docs:data-sources:nginx_csg:properties:waf_spec:blocking_waf_mode", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["waf_spec", "blocking_waf_mode"], "syntax": "attribute", "type": "object"}, {"aliases": ["distributed cloud policy management"], "anchor": "section", "description": "Configuration parameter for distributed cloud policy management.", "document_id": "xcsh-docs:data-sources:nginx_csg:properties:waf_spec:distributed_cloud_policy_management", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["waf_spec", "distributed_cloud_policy_management"], "syntax": "attribute", "type": "object"}, {"aliases": ["monitoring waf mode"], "anchor": "section", "description": "Configuration parameter for monitoring waf mode.", "document_id": "xcsh-docs:data-sources:nginx_csg:properties:waf_spec:monitoring_waf_mode", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["waf_spec", "monitoring_waf_mode"], "syntax": "attribute", "type": "object"}, {"aliases": ["nginx policy management"], "anchor": "section", "description": "Configuration parameter for nginx policy management.", "document_id": "xcsh-docs:data-sources:nginx_csg:properties:waf_spec:nginx_policy_management", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["waf_spec", "nginx_policy_management"], "syntax": "attribute", "type": "object"}, {"aliases": ["none waf mode"], "anchor": "section", "description": "Configuration parameter for none waf mode.", "document_id": "xcsh-docs:data-sources:nginx_csg:properties:waf_spec:none_waf_mode", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["waf_spec", "none_waf_mode"], "syntax": "attribute", "type": "object"}, {"aliases": ["policy file name"], "anchor": "schema-waf_spec--policy_file_name", "description": "WAF Policy File Name. Policy file name for WAF.", "document_id": "xcsh-docs:data-sources:nginx_csg:properties:waf_spec", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["waf_spec", "policy_file_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["policy name"], "anchor": "schema-waf_spec--policy_name", "description": "WAF Policy Name. Policy name configured for WAF.", "document_id": "xcsh-docs:data-sources:nginx_csg:properties:waf_spec", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["waf_spec", "policy_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["security log enabled"], "anchor": "schema-waf_spec--security_log_enabled", "description": "Specifies if security logging is enabled.", "document_id": "xcsh-docs:data-sources:nginx_csg:properties:waf_spec", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["waf_spec", "security_log_enabled"], "syntax": "attribute", "type": "bool"}, {"aliases": ["security log file names"], "anchor": "schema-waf_spec--security_log_file_names", "description": "Specifies the list of security log files specification.", "document_id": "xcsh-docs:data-sources:nginx_csg:properties:waf_spec", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["waf_spec", "security_log_file_names"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nginx_csg/properties/waf_spec/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Configuration for waf_spec.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# waf_spec

Breadcrumbs:

- [xcsh_nginx_csg](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_csg/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_csg/properties/)
- waf_spec

<a id="section"></a>

Type: `"single"`. Computed.

Configuration for waf\_spec.

## Direct properties

- [blocking_waf_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_csg/properties/waf_spec/blocking_waf_mode/): complete subsection reference.

- [distributed_cloud_policy_management](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_csg/properties/waf_spec/distributed_cloud_policy_management/): complete subsection reference.

- [monitoring_waf_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_csg/properties/waf_spec/monitoring_waf_mode/): complete subsection reference.

- [nginx_policy_management](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_csg/properties/waf_spec/nginx_policy_management/): complete subsection reference.

- [none_waf_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_csg/properties/waf_spec/none_waf_mode/): complete subsection reference.

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

- [waf_spec.blocking_waf_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_csg/properties/waf_spec/blocking_waf_mode/)
- [waf_spec.distributed_cloud_policy_management](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_csg/properties/waf_spec/distributed_cloud_policy_management/)
- [waf_spec.monitoring_waf_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_csg/properties/waf_spec/monitoring_waf_mode/)
- [waf_spec.nginx_policy_management](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_csg/properties/waf_spec/nginx_policy_management/)
- [waf_spec.none_waf_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_csg/properties/waf_spec/none_waf_mode/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_csg/properties/)
- [xcsh_nginx_csg](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_csg/)
