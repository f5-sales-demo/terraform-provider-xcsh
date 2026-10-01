---
page_title: "rules.match"
subcategory: ""
description: "rules.match for xcsh_bgp_routing_policy."
xcsh_docs: {"aliases": [], "body_bytes": 2583, "body_sha256": "sha256:17f3e7ef55a4438b39ec86230f3372ace914b4b496026a9bedb113bfc789b333", "child_ids": ["xcsh-docs:data-sources:bgp_routing_policy:properties:rules:match:community", "xcsh-docs:data-sources:bgp_routing_policy:properties:rules:match:ip_prefixes"], "collection_id": "xcsh-docs:data-sources:bgp_routing_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bgp_routing_policy:properties:rules:match", "parent_id": "xcsh-docs:data-sources:bgp_routing_policy:properties:rules", "path": "documentation/data-sources/bgp_routing_policy/properties/rules/match/index.md", "provider_name": "bgp_routing_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "properties", "schema_path": ["rules", "match"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bgp_routing_policy/properties/rules/match/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rules.match for xcsh_bgp_routing_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bgp_routing_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.match

Breadcrumbs:

- [xcsh_bgp_routing_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp_routing_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp_routing_policy/properties/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp_routing_policy/properties/rules/)
- rules.match

<a id="section"></a>

Type: `"single"`. Computed.

Predicates which have to match information in route for action to be applied.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-type_of_match": "[\"as_path\",\"community\",\"ip_prefixes\"]"
}
```

## Direct properties

<a id="schema-rules--match--as_path"></a>

### as_path property

Type: `"string"`. Computed.

Exclusive with \[community ip\_prefixes\] AS path can also be a regex, which will be matched against
route information.

Upstream description:

Exclusive with \[community ip\_prefixes\] AS path can also be a regex, which will be matched against
route information.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [community](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp_routing_policy/properties/rules/match/community/): complete subsection reference.

- [ip_prefixes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp_routing_policy/properties/rules/match/ip_prefixes/): complete subsection reference.

## Next pages

- [rules.match.community](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp_routing_policy/properties/rules/match/community/)
- [rules.match.ip_prefixes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp_routing_policy/properties/rules/match/ip_prefixes/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp_routing_policy/properties/rules/)
- [xcsh_bgp_routing_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp_routing_policy/)
