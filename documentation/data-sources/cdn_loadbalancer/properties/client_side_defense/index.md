---
page_title: "client_side_defense"
subcategory: "Load Balancing"
description: "This defines various configuration OPTIONS for Client-Side Defense Policy."
xcsh_docs: {"aliases": ["client side defense"], "body_bytes": 1519, "body_sha256": "sha256:46146d08a196727a05204e021ae738a7d92d5e20126e708c0779a2e267d460d3", "capabilities": ["cdn", "security.client-side-defense"], "category": "cdn", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:client_side_defense:policy"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:client_side_defense", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:reference", "path": "documentation/data-sources/cdn_loadbalancer/properties/client_side_defense/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-1031112122022321-3012321000333023-2232001202012132-0321133111111102-1203101100022232-0213233213130113-2121311321122232-1121123211010231", "registry_path": "docs/guides/data-sources--cdn_loadbalancer--reference--group-009.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["client_side_defense"], "schema_version": 1, "sections": [{"aliases": ["client side defense policy"], "anchor": "section", "description": "This defines various configuration OPTIONS for Client-Side Defense policy.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:client_side_defense:policy", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["client_side_defense", "policy"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/client_side_defense/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "This defines various configuration OPTIONS for Client-Side Defense Policy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# client_side_defense

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/)
- client_side_defense

<a id="section"></a>

Type: `"single"`. Computed.

\[OneOf: client\_side\_defense, disable\_client\_side\_defense; Default:
disable\_client\_side\_defense\] Defines various configuration OPTIONS for Client-Side Defense
Policy.

Additional upstream details:

This defines various configuration OPTIONS for Client-Side Defense Policy.

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

- [client_side_defense](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/client_side_defense/#section)
- [disable_client_side_defense](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/disable_client_side_defense/#section)

Select alternatives according to the provider validators above.

## Direct properties

- [policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/client_side_defense/policy/): complete subsection reference.
