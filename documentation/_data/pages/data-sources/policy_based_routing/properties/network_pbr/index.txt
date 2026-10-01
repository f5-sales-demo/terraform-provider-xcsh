---
page_title: "network_pbr"
subcategory: ""
description: "network_pbr for xcsh_policy_based_routing."
xcsh_docs: {"aliases": [], "body_bytes": 2417, "body_sha256": "sha256:2146d5b79f746cc918893a830d60f91ad33e4aea86c230f2b7f4df622310947d", "child_ids": ["xcsh-docs:data-sources:policy_based_routing:properties:network_pbr:any", "xcsh-docs:data-sources:policy_based_routing:properties:network_pbr:label_selector", "xcsh-docs:data-sources:policy_based_routing:properties:network_pbr:network_pbr_rules", "xcsh-docs:data-sources:policy_based_routing:properties:network_pbr:prefix_list"], "collection_id": "xcsh-docs:data-sources:policy_based_routing:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:policy_based_routing:properties:network_pbr", "parent_id": "xcsh-docs:data-sources:policy_based_routing:reference", "path": "documentation/data-sources/policy_based_routing/properties/network_pbr/index.md", "provider_name": "policy_based_routing", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "properties", "schema_path": ["network_pbr"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/policy_based_routing/properties/network_pbr/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "network_pbr for xcsh_policy_based_routing.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["policy_based_routingCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# network_pbr

Breadcrumbs:

- [xcsh_policy_based_routing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/)
- network_pbr

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for network pbr.

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
  },
  "x-ves-oneof-field-source_choice": "[\"any\",\"label_selector\",\"prefix_list\"]"
}
```

## Direct properties

- [any](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/network_pbr/any/): complete subsection reference.

- [label_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/network_pbr/label_selector/): complete subsection reference.

- [network_pbr_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/network_pbr/network_pbr_rules/): complete subsection reference.

- [prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/network_pbr/prefix_list/): complete subsection reference.

## Next pages

- [network_pbr.any](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/network_pbr/any/)
- [network_pbr.label_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/network_pbr/label_selector/)
- [network_pbr.network_pbr_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/network_pbr/network_pbr_rules/)
- [network_pbr.prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/network_pbr/prefix_list/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/)
- [xcsh_policy_based_routing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/)
