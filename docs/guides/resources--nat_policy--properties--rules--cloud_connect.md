---
page_title: "rules.cloud_connect"
subcategory: ""
description: "rules.cloud_connect for xcsh_nat_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1259, "body_sha256": "sha256:5b0db4a390e1d1abb695841ab246b0932b8085b340c744736da529335e5cd8a4", "canonical_id": "xcsh-docs:resources:nat_policy:properties:rules:cloud_connect", "child_ids": ["xcsh-docs:resources:nat_policy:properties:rules:cloud_connect:refs"], "collection_id": "xcsh-docs:resources:nat_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:nat_policy:properties:rules:cloud_connect", "parent_id": "xcsh-docs:resources:nat_policy:properties:rules", "path": "docs/guides/resources--nat_policy--properties--rules--cloud_connect.md", "provider_name": "nat_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rules", "cloud_connect"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nat_policy/properties/rules/cloud_connect/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rules.cloud_connect for xcsh_nat_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nat_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.cloud_connect

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md)
- [Property reference](resources--nat_policy--reference.md)
- [rules](resources--nat_policy--properties--rules.md)
- rules.cloud_connect

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for cloud connect.

Upstream description:

Reference to Cloud connect Object.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("refs")}
```

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
cloud_connect {
  # Configure direct properties listed below.
}
```

## Direct properties

- [refs](resources--nat_policy--properties--rules--cloud_connect--refs.md): complete subsection reference.

## Next pages

- [rules.cloud_connect.refs](resources--nat_policy--properties--rules--cloud_connect--refs.md)
- [rules](resources--nat_policy--properties--rules.md)
- [xcsh_nat_policy](../resources/nat_policy.md)
