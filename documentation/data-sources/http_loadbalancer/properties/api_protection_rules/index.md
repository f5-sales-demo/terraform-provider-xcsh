---
page_title: "api_protection_rules"
subcategory: "Load Balancing"
description: "API Protection Rules."
xcsh_docs: {"aliases": ["api protection rules"], "body_bytes": 1080, "body_sha256": "sha256:f877201ff8d74351ebe8d5da415549be1bebe8c60e71f4c54105af6a8cb3e22e", "capabilities": ["load-balancing", "security.api-protection"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules", "xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules:api_groups_rules"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:reference", "path": "documentation/data-sources/http_loadbalancer/properties/api_protection_rules/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-005.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["api_protection_rules"], "schema_version": 1, "sections": [{"aliases": ["api protection rules api endpoint rules"], "anchor": "section", "description": "This category defines specific rules per API endpoints. If request matches any of these rules, skipping second category rules.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["api_protection_rules", "api_endpoint_rules"], "syntax": "attribute", "type": "object"}, {"aliases": ["api protection rules api groups rules"], "anchor": "section", "description": "This category includes rules per API group or Server URL. For API groups, refer to API Definition which includes API groups derived from uploaded swaggers.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules:api_groups_rules", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["api_protection_rules", "api_groups_rules"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/api_protection_rules/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "API Protection Rules.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_protection_rules

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- api_protection_rules

<a id="section"></a>

Type: `"single"`. Computed.

API Protection Rules. API Protection Rules.

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

- [api_endpoint_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_protection_rules/api_endpoint_rules/): complete subsection reference.

- [api_groups_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/): complete subsection reference.
