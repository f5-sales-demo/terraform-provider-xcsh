---
page_title: "forward_proxy_pbr"
subcategory: ""
description: "forward_proxy_pbr for xcsh_policy_based_routing."
xcsh_docs: {"aliases": [], "body_bytes": 1450, "body_sha256": "sha256:f3495e64e242009fc11b26250ceed1ed5f3dcf3719ace874c1eae71ae98c3770", "canonical_id": "xcsh-docs:data-sources:policy_based_routing:properties:forward_proxy_pbr", "child_ids": ["xcsh-docs:data-sources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules"], "collection_id": "xcsh-docs:data-sources:policy_based_routing:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:policy_based_routing:properties:forward_proxy_pbr", "parent_id": "xcsh-docs:data-sources:policy_based_routing:reference", "path": "docs/guides/data-sources--policy_based_routing--properties--forward_proxy_pbr.md", "provider_name": "policy_based_routing", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["forward_proxy_pbr"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/policy_based_routing/properties/forward_proxy_pbr/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "forward_proxy_pbr for xcsh_policy_based_routing.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["policy_based_routingCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# forward_proxy_pbr

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md)
- [Property reference](data-sources--policy_based_routing--reference.md)
- forward_proxy_pbr

<a id="section"></a>

Type: `"single"`. Computed.

\[OneOf: forward\_proxy\_pbr, network\_pbr\] Configuration parameter for forward proxy pbr.

Upstream description:

Network(L3/L4) routing policy rule.

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

OneOf alternatives in this subsection:

- [forward_proxy_pbr](data-sources--policy_based_routing--properties--forward_proxy_pbr.md#section)
- [network_pbr](data-sources--policy_based_routing--properties--network_pbr.md#section)

Select alternatives according to the provider validators above.

## Direct properties

- [forward_proxy_pbr_rules](data-sources--policy_based_routing--properties--forward_proxy_pbr--forward_proxy_pbr_rules.md): complete subsection reference.

## Next pages

- [forward_proxy_pbr.forward_proxy_pbr_rules](data-sources--policy_based_routing--properties--forward_proxy_pbr--forward_proxy_pbr_rules.md)
- [Property reference](data-sources--policy_based_routing--reference.md)
- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md)
