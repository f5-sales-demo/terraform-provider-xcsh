---
page_title: "rules"
subcategory: ""
description: "rules for xcsh_bgp_routing_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1948, "body_sha256": "sha256:890e6af777b57b9fe68ba02002802b360f5562afc9ce544e1911883dd3296fbd", "canonical_id": "xcsh-docs:data-sources:bgp_routing_policy:properties:rules", "child_ids": ["xcsh-docs:data-sources:bgp_routing_policy:properties:rules:action", "xcsh-docs:data-sources:bgp_routing_policy:properties:rules:match"], "collection_id": "xcsh-docs:data-sources:bgp_routing_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bgp_routing_policy:properties:rules", "parent_id": "xcsh-docs:data-sources:bgp_routing_policy:reference", "path": "docs/guides/data-sources--bgp_routing_policy--properties--rules.md", "provider_name": "bgp_routing_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rules"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bgp_routing_policy/properties/rules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rules for xcsh_bgp_routing_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bgp_routing_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# rules

Breadcrumbs:

- [xcsh_bgp_routing_policy](../data-sources/bgp_routing_policy.md)
- [Property reference](data-sources--bgp_routing_policy--reference.md)
- rules

<a id="section"></a>

Type: `"list"`. Computed.

BGP Routing policy is composed of one or more rules. Note that the order of rules is critical as
rules are applied top to bottom.

Upstream description:

A BGP Routing policy is composed of one or more rules. Note that the order of rules is critical as
rules are applied top to bottom.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1,
    "uniqueItems": false
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

## Direct properties

- [action](data-sources--bgp_routing_policy--properties--rules--action.md): complete subsection reference.

- [match](data-sources--bgp_routing_policy--properties--rules--match.md): complete subsection reference.

## Next pages

- [rules.action](data-sources--bgp_routing_policy--properties--rules--action.md)
- [rules.match](data-sources--bgp_routing_policy--properties--rules--match.md)
- [Property reference](data-sources--bgp_routing_policy--reference.md)
- [xcsh_bgp_routing_policy](../data-sources/bgp_routing_policy.md)
