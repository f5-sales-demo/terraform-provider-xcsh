---
page_title: "rules.match.ip_prefixes.prefixes.equal_or_longer_than"
subcategory: ""
description: "rules.match.ip_prefixes.prefixes.equal_or_longer_than for xcsh_bgp_routing_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1341, "body_sha256": "sha256:2400dd4efce598596cb2b3e9283e76a63ba0711f6c799c855b090bb5ee4afd6b", "canonical_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match:ip_prefixes:prefixes:equal_or_longer_than", "child_ids": [], "collection_id": "xcsh-docs:resources:bgp_routing_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match:ip_prefixes:prefixes:equal_or_longer_than", "parent_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match:ip_prefixes:prefixes", "path": "docs/guides/resources--bgp_routing_policy--properties--rules--match--ip_prefixes--prefixes--equal_or_longer_than.md", "provider_name": "bgp_routing_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rules", "match", "ip_prefixes", "prefixes", "equal_or_longer_than"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bgp_routing_policy/properties/rules/match/ip_prefixes/prefixes/equal_or_longer_than/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rules.match.ip_prefixes.prefixes.equal_or_longer_than for xcsh_bgp_routing_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bgp_routing_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# rules.match.ip_prefixes.prefixes.equal_or_longer_than

Breadcrumbs:

- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md)
- [Property reference](resources--bgp_routing_policy--reference.md)
- [rules](resources--bgp_routing_policy--properties--rules.md)
- [rules.match](resources--bgp_routing_policy--properties--rules--match.md)
- [rules.match.ip_prefixes](resources--bgp_routing_policy--properties--rules--match--ip_prefixes.md)
- [rules.match.ip_prefixes.prefixes](resources--bgp_routing_policy--properties--rules--match--ip_prefixes--prefixes.md)
- rules.match.ip_prefixes.prefixes.equal_or_longer_than

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for equal or longer than.

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
equal_or_longer_than = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [rules.match.ip_prefixes.prefixes](resources--bgp_routing_policy--properties--rules--match--ip_prefixes--prefixes.md)
- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md)
