---
page_title: "rules.virtual_network"
subcategory: ""
description: "rules.virtual_network for xcsh_nat_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1068, "body_sha256": "sha256:3221001ad1682c3de9c424f636f7e91b6fa5a8057bc8d5afb50a7707a97e9249", "canonical_id": "xcsh-docs:resources:nat_policy:properties:rules:virtual_network", "child_ids": ["xcsh-docs:resources:nat_policy:properties:rules:virtual_network:refs"], "collection_id": "xcsh-docs:resources:nat_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:nat_policy:properties:rules:virtual_network", "parent_id": "xcsh-docs:resources:nat_policy:properties:rules", "path": "docs/guides/resources--nat_policy--properties--rules--virtual_network.md", "provider_name": "nat_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rules", "virtual_network"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nat_policy/properties/rules/virtual_network/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rules.virtual_network for xcsh_nat_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nat_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.virtual_network

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md)
- [Property reference](resources--nat_policy--reference.md)
- [rules](resources--nat_policy--properties--rules.md)
- rules.virtual_network

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Carries the reference to virtual network.

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
virtual_network {
  # Configure direct properties listed below.
}
```

## Direct properties

- [refs](resources--nat_policy--properties--rules--virtual_network--refs.md): complete subsection reference.

## Next pages

- [rules.virtual_network.refs](resources--nat_policy--properties--rules--virtual_network--refs.md)
- [rules](resources--nat_policy--properties--rules.md)
- [xcsh_nat_policy](../resources/nat_policy.md)
