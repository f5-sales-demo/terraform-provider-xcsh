---
page_title: "network_pbr"
subcategory: ""
description: "network_pbr for xcsh_policy_based_routing."
xcsh_docs: {"aliases": [], "body_bytes": 1809, "body_sha256": "sha256:0ac620a07a4ff1b48df01ff13b630041d505361cbf4e63f73b2a6f61b1e9e7e9", "canonical_id": "xcsh-docs:data-sources:policy_based_routing:properties:network_pbr", "child_ids": ["xcsh-docs:data-sources:policy_based_routing:properties:network_pbr:any", "xcsh-docs:data-sources:policy_based_routing:properties:network_pbr:label_selector", "xcsh-docs:data-sources:policy_based_routing:properties:network_pbr:network_pbr_rules", "xcsh-docs:data-sources:policy_based_routing:properties:network_pbr:prefix_list"], "collection_id": "xcsh-docs:data-sources:policy_based_routing:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:policy_based_routing:properties:network_pbr", "parent_id": "xcsh-docs:data-sources:policy_based_routing:reference", "path": "docs/guides/data-sources--policy_based_routing--properties--network_pbr.md", "provider_name": "policy_based_routing", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["network_pbr"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/policy_based_routing/properties/network_pbr/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "network_pbr for xcsh_policy_based_routing.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["policy_based_routingCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# network_pbr

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md)
- [Property reference](data-sources--policy_based_routing--reference.md)
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

- [any](data-sources--policy_based_routing--properties--network_pbr--any.md): complete subsection reference.

- [label_selector](data-sources--policy_based_routing--properties--network_pbr--label_selector.md): complete subsection reference.

- [network_pbr_rules](data-sources--policy_based_routing--properties--network_pbr--network_pbr_rules.md): complete subsection reference.

- [prefix_list](data-sources--policy_based_routing--properties--network_pbr--prefix_list.md): complete subsection reference.

## Next pages

- [network_pbr.any](data-sources--policy_based_routing--properties--network_pbr--any.md)
- [network_pbr.label_selector](data-sources--policy_based_routing--properties--network_pbr--label_selector.md)
- [network_pbr.network_pbr_rules](data-sources--policy_based_routing--properties--network_pbr--network_pbr_rules.md)
- [network_pbr.prefix_list](data-sources--policy_based_routing--properties--network_pbr--prefix_list.md)
- [Property reference](data-sources--policy_based_routing--reference.md)
- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md)
