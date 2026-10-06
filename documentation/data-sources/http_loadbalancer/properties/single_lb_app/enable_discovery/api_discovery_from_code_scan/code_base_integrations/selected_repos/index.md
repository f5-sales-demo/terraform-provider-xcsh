---
page_title: "single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos"
subcategory: "Load Balancing"
description: "Select which API repositories represent the LB applications."
xcsh_docs: {"aliases": ["single lb app enable discovery api discovery from code scan code base integrations selected repos"], "body_bytes": 2345, "body_sha256": "sha256:2f64e704af8ee09e1fa4a25ead5d3d4de6fe1634af23da9c1d6364af50c92689", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_discovery_from_code_scan:code_base_integrations:selected_repos", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_discovery_from_code_scan:code_base_integrations", "path": "documentation/data-sources/http_loadbalancer/properties/single_lb_app/enable_discovery/api_discovery_from_code_scan/code_base_integrations/selected_repos/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-2331001311312010-1310031003233132-3110032121120130-3123310321110231-2022211132330320-3112132000002312-3303020212201111-3311101320013212", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-026.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["single_lb_app", "enable_discovery", "api_discovery_from_code_scan", "code_base_integrations", "selected_repos"], "schema_version": 1, "sections": [{"aliases": ["single lb app enable discovery api discovery from code scan code base integrations selected repos api code repo"], "anchor": "schema-single_lb_app--enable_discovery--api_discovery_from_code_scan--code_base_integrations--selected_repos--api_code_repo", "description": "Code repository which contain API endpoints.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_discovery_from_code_scan:code_base_integrations:selected_repos", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["single_lb_app", "enable_discovery", "api_discovery_from_code_scan", "code_base_integrations", "selected_repos", "api_code_repo"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/single_lb_app/enable_discovery/api_discovery_from_code_scan/code_base_integrations/selected_repos/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Select which API repositories represent the LB applications.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [single_lb_app](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/single_lb_app/)
- [single_lb_app.enable_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/single_lb_app/enable_discovery/)
- [single_lb_app.enable_discovery.api_discovery_from_code_scan](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/single_lb_app/enable_discovery/api_discovery_from_code_scan/)
- [single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/single_lb_app/enable_discovery/api_discovery_from_code_scan/code_base_integrations/)
- single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos

<a id="section"></a>

Type: `"single"`. Computed.

Select which API repositories represent the LB applications.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

## Direct properties

<a id="schema-single_lb_app--enable_discovery--api_discovery_from_code_scan--code_base_integrations--selected_repos--api_code_repo"></a>

### api_code_repo property

Type: `["list", "string"]`. Computed.

Code repository which contain API endpoints.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```
