---
page_title: "api_protection_rules.api_endpoint_rules.action"
subcategory: "Load Balancing"
description: "The action to take if the input request matches the rule."
xcsh_docs: {"aliases": ["api protection rules api endpoint rules action"], "body_bytes": 2266, "body_sha256": "sha256:8380b405309628f9fddf402ffe0f4f49a9187aaceded756d930142f0392659a8", "capabilities": ["load-balancing", "security.api-protection"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:action:allow", "xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:action:deny"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:action", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules", "path": "documentation/data-sources/http_loadbalancer/properties/api_protection_rules/api_endpoint_rules/action/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-0101112111332320-2210011322232033-2103230203130113-0132012311021110-0203002212031131-0211113333122200-3200122202123212-3033221000020213", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["api_protection_rules", "api_endpoint_rules", "action"], "schema_version": 1, "sections": [{"aliases": ["api protection rules api endpoint rules action allow"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:action:allow", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_protection_rules", "api_endpoint_rules", "action", "allow"], "syntax": "attribute", "type": "object"}, {"aliases": ["api protection rules api endpoint rules action deny"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:action:deny", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_protection_rules", "api_endpoint_rules", "action", "deny"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/api_protection_rules/api_endpoint_rules/action/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "The action to take if the input request matches the rule.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_protection_rules.api_endpoint_rules.action

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [api_protection_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_protection_rules/)
- [api_protection_rules.api_endpoint_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_protection_rules/api_endpoint_rules/)
- api_protection_rules.api_endpoint_rules.action

<a id="section"></a>

Type: `"single"`. Computed.

The action to take if the input request matches the rule.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-action": "[\"allow\",\"deny\"]"
}
```

## Direct properties

- [allow](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_protection_rules/api_endpoint_rules/action/allow/): complete subsection reference.

- [deny](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_protection_rules/api_endpoint_rules/action/deny/): complete subsection reference.

## Next pages

- [api_protection_rules.api_endpoint_rules.action.allow](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_protection_rules/api_endpoint_rules/action/allow/)
- [api_protection_rules.api_endpoint_rules.action.deny](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_protection_rules/api_endpoint_rules/action/deny/)
- [api_protection_rules.api_endpoint_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_protection_rules/api_endpoint_rules/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
