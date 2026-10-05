---
page_title: "api_protection_rules.api_groups_rules.action"
subcategory: "Load Balancing"
description: "The action to take if the input request matches the rule."
xcsh_docs: {"aliases": ["api protection rules api groups rules action"], "body_bytes": 2242, "body_sha256": "sha256:2b2e84c1504a3fe7c0ac52d1d9f908b53a6b0cf43f49f89310478801efb3aca7", "capabilities": ["load-balancing", "security.api-protection"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:action:allow", "xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:action:deny"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:action", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules:api_groups_rules", "path": "documentation/data-sources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/action/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-3020123231322122-3002200301232223-1112021032231130-2323213133100010-2223010022222211-1220022221011102-3312102130121023-2021031100210333", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-005.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["api_protection_rules", "api_groups_rules", "action"], "schema_version": 1, "sections": [{"aliases": ["api protection rules api groups rules action allow"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:action:allow", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_protection_rules", "api_groups_rules", "action", "allow"], "syntax": "attribute", "type": "object"}, {"aliases": ["api protection rules api groups rules action deny"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:action:deny", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_protection_rules", "api_groups_rules", "action", "deny"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/action/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "The action to take if the input request matches the rule.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_protection_rules.api_groups_rules.action

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [api_protection_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_protection_rules/)
- [api_protection_rules.api_groups_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/)
- api_protection_rules.api_groups_rules.action

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

- [allow](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/action/allow/): complete subsection reference.

- [deny](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/action/deny/): complete subsection reference.

## Next pages

- [api_protection_rules.api_groups_rules.action.allow](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/action/allow/)
- [api_protection_rules.api_groups_rules.action.deny](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/action/deny/)
- [api_protection_rules.api_groups_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
