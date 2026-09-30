---
page_title: "rules.virtual_network"
subcategory: ""
description: "rules.virtual_network for xcsh_nat_policy."
xcsh_docs: {"aliases": [], "body_bytes": 969, "body_sha256": "sha256:b4ac818545c8ae33ad8499a72d69722bbed757c7d16d5c8d6a6c9a849993bf21", "canonical_id": "xcsh-docs:resources:nat_policy:properties:rules:virtual_network", "child_ids": ["xcsh-docs:resources:nat_policy:properties:rules:virtual_network:refs"], "collection_id": "xcsh-docs:resources:nat_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:nat_policy:properties:rules:virtual_network", "parent_id": "xcsh-docs:resources:nat_policy:properties:rules", "path": "docs/guides/resources--nat_policy--properties--rules--virtual_network.md", "provider_name": "nat_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rules", "virtual_network"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nat_policy/properties/rules/virtual_network/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rules.virtual_network for xcsh_nat_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nat_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

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
