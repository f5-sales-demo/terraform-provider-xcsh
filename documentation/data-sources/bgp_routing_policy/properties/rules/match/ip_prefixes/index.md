---
page_title: "rules.match.ip_prefixes"
subcategory: ""
description: "List of IP prefix and prefix length range match condition."
xcsh_docs: {"aliases": ["rules match ip prefixes"], "body_bytes": 1576, "body_sha256": "sha256:dbe4998ae133a73e28fda7343e360dd2c67a49426ca5affd09125be50008f572", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:bgp_routing_policy:properties:rules:match:ip_prefixes:prefixes"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bgp_routing_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bgp_routing_policy:properties:rules:match:ip_prefixes", "parent_id": "xcsh-docs:data-sources:bgp_routing_policy:properties:rules:match", "path": "documentation/data-sources/bgp_routing_policy/properties/rules/match/ip_prefixes/index.md", "product": "distributed-cloud", "provider_name": "bgp_routing_policy", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-3232120320131320-1221111231000331-2033031200002022-2231123202131301-3023212210322112-3131202030202103-2111032111021232-0012100132103323", "registry_path": "docs/guides/data-sources--bgp_routing_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rules", "match", "ip_prefixes"], "schema_version": 1, "sections": [{"aliases": ["rules match ip prefixes prefixes"], "anchor": "section", "description": "List of IP prefix.", "document_id": "xcsh-docs:data-sources:bgp_routing_policy:properties:rules:match:ip_prefixes:prefixes", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["rules", "match", "ip_prefixes", "prefixes"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bgp_routing_policy/properties/rules/match/ip_prefixes/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "List of IP prefix and prefix length range match condition.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["bgp_routing_policyCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.match.ip_prefixes

Breadcrumbs:

- [xcsh_bgp_routing_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp_routing_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp_routing_policy/properties/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp_routing_policy/properties/rules/)
- [rules.match](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp_routing_policy/properties/rules/match/)
- rules.match.ip_prefixes

<a id="section"></a>

Type: `"single"`. Computed.

List of IP prefix and prefix length range match condition.

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

- [prefixes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp_routing_policy/properties/rules/match/ip_prefixes/prefixes/): complete subsection reference.

## Next pages

- [rules.match.ip_prefixes.prefixes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp_routing_policy/properties/rules/match/ip_prefixes/prefixes/)
- [rules.match](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp_routing_policy/properties/rules/match/)
- [xcsh_bgp_routing_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp_routing_policy/)
