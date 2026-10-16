---
page_title: "rules.action.allow"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["rules action allow"], "body_bytes": 1095, "body_sha256": "sha256:f344c814cf885c89e6f7c8de40c865fbcda6d5fa46648764acf793449f187485", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bgp_routing_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action:allow", "parent_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action", "path": "documentation/resources/bgp_routing_policy/properties/rules/action/allow/index.md", "product": "distributed-cloud", "provider_name": "bgp_routing_policy", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "resources", "registry_anchor": "canonical-0021130232302232-0313010111011031-2212010021233102-0302133233321200-1212213303100312-3023331320222022-0220232323330300-3201112022203213", "registry_path": "docs/guides/resources--bgp_routing_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rules", "action", "allow"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bgp_routing_policy/properties/rules/action/allow/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": ["bgp_routing_policyCreateRequest"], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.action.allow

Breadcrumbs:

- [xcsh_bgp_routing_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/properties/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/properties/rules/)
- [rules.action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/properties/rules/action/)
- rules.action.allow

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

This can be used for messages where no values are needed.

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

Terraform syntax:

```terraform
allow = {}
```

This is an empty object or choice marker. It has no direct properties.
