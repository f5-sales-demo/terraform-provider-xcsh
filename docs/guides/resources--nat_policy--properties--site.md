---
page_title: "site"
subcategory: ""
description: "site for xcsh_nat_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1118, "body_sha256": "sha256:61d5030adbbd47251a313fcc4adc86fe2a50d2172c33d61a551a07b4c17d7e25", "canonical_id": "xcsh-docs:resources:nat_policy:properties:site", "child_ids": ["xcsh-docs:resources:nat_policy:properties:site:refs"], "collection_id": "xcsh-docs:resources:nat_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:nat_policy:properties:site", "parent_id": "xcsh-docs:resources:nat_policy:reference", "path": "docs/guides/resources--nat_policy--properties--site.md", "provider_name": "nat_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["site"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nat_policy/properties/site/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "site for xcsh_nat_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nat_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# site

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md)
- [Property reference](resources--nat_policy--reference.md)
- site

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Site Reference Type. Reference to Site Object.

Upstream description:

Reference to Site Object.

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
site {
  # Configure direct properties listed below.
}
```

## Direct properties

- [refs](resources--nat_policy--properties--site--refs.md): complete subsection reference.

## Next pages

- [site.refs](resources--nat_policy--properties--site--refs.md)
- [Property reference](resources--nat_policy--reference.md)
- [xcsh_nat_policy](../resources/nat_policy.md)
