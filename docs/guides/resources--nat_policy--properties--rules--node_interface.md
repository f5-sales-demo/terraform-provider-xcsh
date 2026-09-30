---
page_title: "rules.node_interface"
subcategory: ""
description: "rules.node_interface for xcsh_nat_policy."
xcsh_docs: {"aliases": [], "body_bytes": 999, "body_sha256": "sha256:48c68aff76727cf8004601fe2299d89107851bc41d2528ba6791feff9da2dd50", "canonical_id": "xcsh-docs:resources:nat_policy:properties:rules:node_interface", "child_ids": ["xcsh-docs:resources:nat_policy:properties:rules:node_interface:list"], "collection_id": "xcsh-docs:resources:nat_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:nat_policy:properties:rules:node_interface", "parent_id": "xcsh-docs:resources:nat_policy:properties:rules", "path": "docs/guides/resources--nat_policy--properties--rules--node_interface.md", "provider_name": "nat_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rules", "node_interface"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nat_policy/properties/rules/node_interface/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rules.node_interface for xcsh_nat_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nat_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# rules.node_interface

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md)
- [Property reference](resources--nat_policy--reference.md)
- [rules](resources--nat_policy--properties--rules.md)
- rules.node_interface

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

On multinode site, this type holds the information about per node interfaces.

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
node_interface {
  # Configure direct properties listed below.
}
```

## Direct properties

- [list](resources--nat_policy--properties--rules--node_interface--list.md): complete subsection reference.

## Next pages

- [rules.node_interface.list](resources--nat_policy--properties--rules--node_interface--list.md)
- [rules](resources--nat_policy--properties--rules.md)
- [xcsh_nat_policy](../resources/nat_policy.md)
