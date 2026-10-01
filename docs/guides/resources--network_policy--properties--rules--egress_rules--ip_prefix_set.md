---
page_title: "rules.egress_rules.ip_prefix_set"
subcategory: "Security"
description: "rules.egress_rules.ip_prefix_set for xcsh_network_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1346, "body_sha256": "sha256:f8a870673611cf9836a8dcccfb7e3f279a0b3a0f0374138bbd8877e7f9a7ef95", "canonical_id": "xcsh-docs:resources:network_policy:properties:rules:egress_rules:ip_prefix_set", "child_ids": ["xcsh-docs:resources:network_policy:properties:rules:egress_rules:ip_prefix_set:ref"], "collection_id": "xcsh-docs:resources:network_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_policy:properties:rules:egress_rules:ip_prefix_set", "parent_id": "xcsh-docs:resources:network_policy:properties:rules:egress_rules", "path": "docs/guides/resources--network_policy--properties--rules--egress_rules--ip_prefix_set.md", "provider_name": "network_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rules", "egress_rules", "ip_prefix_set"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_policy/properties/rules/egress_rules/ip_prefix_set/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rules.egress_rules.ip_prefix_set for xcsh_network_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.egress_rules.ip_prefix_set

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md)
- [Property reference](resources--network_policy--reference.md)
- [rules](resources--network_policy--properties--rules.md)
- [rules.egress_rules](resources--network_policy--properties--rules--egress_rules.md)
- rules.egress_rules.ip_prefix_set

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

- [ref](resources--network_policy--properties--rules--egress_rules--ip_prefix_set--ref.md): complete subsection reference.

## Next pages

- [rules.egress_rules.ip_prefix_set.ref](resources--network_policy--properties--rules--egress_rules--ip_prefix_set--ref.md)
- [rules.egress_rules](resources--network_policy--properties--rules--egress_rules.md)
- [xcsh_network_policy](../resources/network_policy.md)
