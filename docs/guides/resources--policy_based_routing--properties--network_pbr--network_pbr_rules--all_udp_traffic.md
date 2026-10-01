---
page_title: "network_pbr.network_pbr_rules.all_udp_traffic"
subcategory: ""
description: "network_pbr.network_pbr_rules.all_udp_traffic for xcsh_policy_based_routing."
xcsh_docs: {"aliases": [], "body_bytes": 1249, "body_sha256": "sha256:d3e4c5fa35ded4379abdf869b7e4a73127b6c7c18fdf492bc3e9b8f5da59f33e", "canonical_id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:all_udp_traffic", "child_ids": [], "collection_id": "xcsh-docs:resources:policy_based_routing:collection", "completeness": "complete", "id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:all_udp_traffic", "parent_id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules", "path": "docs/guides/resources--policy_based_routing--properties--network_pbr--network_pbr_rules--all_udp_traffic.md", "provider_name": "policy_based_routing", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["network_pbr", "network_pbr_rules", "all_udp_traffic"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/policy_based_routing/properties/network_pbr/network_pbr_rules/all_udp_traffic/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "network_pbr.network_pbr_rules.all_udp_traffic for xcsh_policy_based_routing.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["policy_based_routingCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# network_pbr.network_pbr_rules.all_udp_traffic

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md)
- [Property reference](resources--policy_based_routing--reference.md)
- [network_pbr](resources--policy_based_routing--properties--network_pbr.md)
- [network_pbr.network_pbr_rules](resources--policy_based_routing--properties--network_pbr--network_pbr_rules.md)
- network_pbr.network_pbr_rules.all_udp_traffic

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for all udp traffic.

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
all_udp_traffic = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [network_pbr.network_pbr_rules](resources--policy_based_routing--properties--network_pbr--network_pbr_rules.md)
- [xcsh_policy_based_routing](../resources/policy_based_routing.md)
