---
page_title: "rules.ingress_rules.ip_prefix_set"
subcategory: "Security"
description: "rules.ingress_rules.ip_prefix_set for xcsh_network_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1355, "body_sha256": "sha256:9f5110ce98c85d316a907359e27f5d41705e598b0c4c3bb9cb7b34e2cecd994f", "canonical_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:ip_prefix_set", "child_ids": ["xcsh-docs:resources:network_policy:properties:rules:ingress_rules:ip_prefix_set:ref"], "collection_id": "xcsh-docs:resources:network_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:ip_prefix_set", "parent_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules", "path": "docs/guides/resources--network_policy--properties--rules--ingress_rules--ip_prefix_set.md", "provider_name": "network_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rules", "ingress_rules", "ip_prefix_set"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_policy/properties/rules/ingress_rules/ip_prefix_set/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rules.ingress_rules.ip_prefix_set for xcsh_network_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.ingress_rules.ip_prefix_set

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md)
- [Property reference](resources--network_policy--reference.md)
- [rules](resources--network_policy--properties--rules.md)
- [rules.ingress_rules](resources--network_policy--properties--rules--ingress_rules.md)
- rules.ingress_rules.ip_prefix_set

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

- [ref](resources--network_policy--properties--rules--ingress_rules--ip_prefix_set--ref.md): complete subsection reference.

## Next pages

- [rules.ingress_rules.ip_prefix_set.ref](resources--network_policy--properties--rules--ingress_rules--ip_prefix_set--ref.md)
- [rules.ingress_rules](resources--network_policy--properties--rules--ingress_rules.md)
- [xcsh_network_policy](../resources/network_policy.md)
