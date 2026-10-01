---
page_title: "network_pbr.network_pbr_rules.all_tcp_traffic"
subcategory: ""
description: "network_pbr.network_pbr_rules.all_tcp_traffic for xcsh_policy_based_routing."
xcsh_docs: {"aliases": [], "body_bytes": 1249, "body_sha256": "sha256:dfd2b9bfb76c95c36ed310249aba52f05db1bc5ad54152cbcff9f0ce38ede16b", "canonical_id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:all_tcp_traffic", "child_ids": [], "collection_id": "xcsh-docs:resources:policy_based_routing:collection", "completeness": "complete", "id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:all_tcp_traffic", "parent_id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules", "path": "docs/guides/resources--policy_based_routing--properties--network_pbr--network_pbr_rules--all_tcp_traffic.md", "provider_name": "policy_based_routing", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["network_pbr", "network_pbr_rules", "all_tcp_traffic"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/policy_based_routing/properties/network_pbr/network_pbr_rules/all_tcp_traffic/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "network_pbr.network_pbr_rules.all_tcp_traffic for xcsh_policy_based_routing.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["policy_based_routingCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# network_pbr.network_pbr_rules.all_tcp_traffic

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md)
- [Property reference](resources--policy_based_routing--reference.md)
- [network_pbr](resources--policy_based_routing--properties--network_pbr.md)
- [network_pbr.network_pbr_rules](resources--policy_based_routing--properties--network_pbr--network_pbr_rules.md)
- network_pbr.network_pbr_rules.all_tcp_traffic

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for all tcp traffic.

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
all_tcp_traffic = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [network_pbr.network_pbr_rules](resources--policy_based_routing--properties--network_pbr--network_pbr_rules.md)
- [xcsh_policy_based_routing](../resources/policy_based_routing.md)
