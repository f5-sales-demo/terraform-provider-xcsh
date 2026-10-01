---
page_title: "network_pbr.network_pbr_rules.ip_prefix_set"
subcategory: ""
description: "network_pbr.network_pbr_rules.ip_prefix_set for xcsh_policy_based_routing."
xcsh_docs: {"aliases": [], "body_bytes": 1517, "body_sha256": "sha256:19b28d22d12a66b45cd980c9dd13ec79a7a2257fbc049958c0b4868ff6eb21a3", "canonical_id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:ip_prefix_set", "child_ids": ["xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:ip_prefix_set:ref"], "collection_id": "xcsh-docs:resources:policy_based_routing:collection", "completeness": "complete", "id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:ip_prefix_set", "parent_id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules", "path": "docs/guides/resources--policy_based_routing--properties--network_pbr--network_pbr_rules--ip_prefix_set.md", "provider_name": "policy_based_routing", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["network_pbr", "network_pbr_rules", "ip_prefix_set"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/policy_based_routing/properties/network_pbr/network_pbr_rules/ip_prefix_set/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "network_pbr.network_pbr_rules.ip_prefix_set for xcsh_policy_based_routing.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["policy_based_routingCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# network_pbr.network_pbr_rules.ip_prefix_set

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md)
- [Property reference](resources--policy_based_routing--reference.md)
- [network_pbr](resources--policy_based_routing--properties--network_pbr.md)
- [network_pbr.network_pbr_rules](resources--policy_based_routing--properties--network_pbr--network_pbr_rules.md)
- network_pbr.network_pbr_rules.ip_prefix_set

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

List of references to ip\_prefix\_set objects.

Upstream description:

A list of references to ip\_prefix\_set objects.

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
ip_prefix_set {
  # Configure direct properties listed below.
}
```

## Direct properties

- [ref](resources--policy_based_routing--properties--network_pbr--network_pbr_rules--ip_prefix_set--ref.md): complete subsection reference.

## Next pages

- [network_pbr.network_pbr_rules.ip_prefix_set.ref](resources--policy_based_routing--properties--network_pbr--network_pbr_rules--ip_prefix_set--ref.md)
- [network_pbr.network_pbr_rules](resources--policy_based_routing--properties--network_pbr--network_pbr_rules.md)
- [xcsh_policy_based_routing](../resources/policy_based_routing.md)
