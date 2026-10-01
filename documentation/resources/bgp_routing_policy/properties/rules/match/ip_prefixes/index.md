---
page_title: "rules.match.ip_prefixes"
subcategory: ""
description: "rules.match.ip_prefixes for xcsh_bgp_routing_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1819, "body_sha256": "sha256:f3b3f635517dfbc771ff9b32ea23d7fca3f36c7f0d6d8b02b17302f034511c93", "child_ids": ["xcsh-docs:resources:bgp_routing_policy:properties:rules:match:ip_prefixes:prefixes"], "collection_id": "xcsh-docs:resources:bgp_routing_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match:ip_prefixes", "parent_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match", "path": "documentation/resources/bgp_routing_policy/properties/rules/match/ip_prefixes/index.md", "provider_name": "bgp_routing_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["rules", "match", "ip_prefixes"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bgp_routing_policy/properties/rules/match/ip_prefixes/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rules.match.ip_prefixes for xcsh_bgp_routing_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bgp_routing_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.match.ip_prefixes

Breadcrumbs:

- [xcsh_bgp_routing_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/properties/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/properties/rules/)
- [rules.match](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/properties/rules/match/)
- rules.match.ip_prefixes

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

List of IP prefix and prefix length range match condition.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("prefixes")}
```

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
ip_prefixes {
  # Configure direct properties listed below.
}
```

## Direct properties

- [prefixes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/properties/rules/match/ip_prefixes/prefixes/): complete subsection reference.

## Next pages

- [rules.match.ip_prefixes.prefixes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/properties/rules/match/ip_prefixes/prefixes/)
- [rules.match](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/properties/rules/match/)
- [xcsh_bgp_routing_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/)
