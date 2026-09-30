---
page_title: "network_pbr"
subcategory: ""
description: "network_pbr for xcsh_policy_based_routing."
xcsh_docs: {"aliases": [], "body_bytes": 1710, "body_sha256": "sha256:e5485b1516306179d2126fc247e757b30d5e101a4e75935580b50fc9bda55333", "canonical_id": "xcsh-docs:data-sources:policy_based_routing:properties:network_pbr", "child_ids": ["xcsh-docs:data-sources:policy_based_routing:properties:network_pbr:any", "xcsh-docs:data-sources:policy_based_routing:properties:network_pbr:label_selector", "xcsh-docs:data-sources:policy_based_routing:properties:network_pbr:network_pbr_rules", "xcsh-docs:data-sources:policy_based_routing:properties:network_pbr:prefix_list"], "collection_id": "xcsh-docs:data-sources:policy_based_routing:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:policy_based_routing:properties:network_pbr", "parent_id": "xcsh-docs:data-sources:policy_based_routing:reference", "path": "docs/guides/data-sources--policy_based_routing--properties--network_pbr.md", "provider_name": "policy_based_routing", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["network_pbr"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/policy_based_routing/properties/network_pbr/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "network_pbr for xcsh_policy_based_routing.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["policy_based_routingCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

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
