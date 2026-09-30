---
page_title: "rules.match.ip_prefixes"
subcategory: ""
description: "rules.match.ip_prefixes for xcsh_bgp_routing_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1075, "body_sha256": "sha256:31c3414cf6634c9f19ed0566bf48d807ea106c2156c3d56866b01057d372c03e", "canonical_id": "xcsh-docs:data-sources:bgp_routing_policy:properties:rules:match:ip_prefixes", "child_ids": ["xcsh-docs:data-sources:bgp_routing_policy:properties:rules:match:ip_prefixes:prefixes"], "collection_id": "xcsh-docs:data-sources:bgp_routing_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bgp_routing_policy:properties:rules:match:ip_prefixes", "parent_id": "xcsh-docs:data-sources:bgp_routing_policy:properties:rules:match", "path": "docs/guides/data-sources--bgp_routing_policy--properties--rules--match--ip_prefixes.md", "provider_name": "bgp_routing_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rules", "match", "ip_prefixes"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bgp_routing_policy/properties/rules/match/ip_prefixes/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rules.match.ip_prefixes for xcsh_bgp_routing_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bgp_routing_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# rules.match.ip_prefixes

Breadcrumbs:

- [xcsh_bgp_routing_policy](../data-sources/bgp_routing_policy.md)
- [Property reference](data-sources--bgp_routing_policy--reference.md)
- [rules](data-sources--bgp_routing_policy--properties--rules.md)
- [rules.match](data-sources--bgp_routing_policy--properties--rules--match.md)
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

- [prefixes](data-sources--bgp_routing_policy--properties--rules--match--ip_prefixes--prefixes.md): complete subsection reference.

## Next pages

- [rules.match.ip_prefixes.prefixes](data-sources--bgp_routing_policy--properties--rules--match--ip_prefixes--prefixes.md)
- [rules.match](data-sources--bgp_routing_policy--properties--rules--match.md)
- [xcsh_bgp_routing_policy](../data-sources/bgp_routing_policy.md)
