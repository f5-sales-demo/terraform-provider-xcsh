---
page_title: "rules.match.ip_prefixes.prefixes.longer_than"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["rules match ip prefixes prefixes longer than"], "body_bytes": 1805, "body_sha256": "sha256:9d76892961d66dc70a2ba1ef6b127d0c2abb59378c8e4f6aa40d91779e9612e2", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bgp_routing_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match:ip_prefixes:prefixes:longer_than", "parent_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match:ip_prefixes:prefixes", "path": "documentation/resources/bgp_routing_policy/properties/rules/match/ip_prefixes/prefixes/longer_than/index.md", "product": "distributed-cloud", "provider_name": "bgp_routing_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3012300123222110-1211201132023330-3133033023001110-1320330202210311-1202131330102203-0303231311330123-0323130113211230-3230100233033010", "registry_path": "docs/guides/resources--bgp_routing_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rules", "match", "ip_prefixes", "prefixes", "longer_than"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bgp_routing_policy/properties/rules/match/ip_prefixes/prefixes/longer_than/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bgp_routing_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.match.ip_prefixes.prefixes.longer_than

Breadcrumbs:

- [xcsh_bgp_routing_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/properties/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/properties/rules/)
- [rules.match](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/properties/rules/match/)
- [rules.match.ip_prefixes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/properties/rules/match/ip_prefixes/)
- [rules.match.ip_prefixes.prefixes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/properties/rules/match/ip_prefixes/prefixes/)
- rules.match.ip_prefixes.prefixes.longer_than

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for longer than.

Upstream description:

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
longer_than = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [rules.match.ip_prefixes.prefixes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/properties/rules/match/ip_prefixes/prefixes/)
- [xcsh_bgp_routing_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/)
