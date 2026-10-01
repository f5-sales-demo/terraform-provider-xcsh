---
page_title: "where.virtual_network"
subcategory: ""
description: "where.virtual_network for xcsh_advertise_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1730, "body_sha256": "sha256:b1078271a3320dadb234cb12e5ff9ca6c8b38c327b2a7caa9840c045caff71ec", "child_ids": ["xcsh-docs:resources:advertise_policy:properties:where:virtual_network:ref"], "collection_id": "xcsh-docs:resources:advertise_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:advertise_policy:properties:where:virtual_network", "parent_id": "xcsh-docs:resources:advertise_policy:properties:where", "path": "documentation/resources/advertise_policy/properties/where/virtual_network/index.md", "provider_name": "advertise_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["where", "virtual_network"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/advertise_policy/properties/where/virtual_network/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "where.virtual_network for xcsh_advertise_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["advertise_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# where.virtual_network

Breadcrumbs:

- [xcsh_advertise_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/advertise_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/advertise_policy/properties/)
- [where](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/advertise_policy/properties/where/)
- where.virtual_network

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Specifies a direct reference to a network configuration object.

Upstream description:

This specifies a direct reference to a network configuration object.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ref")}
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
virtual_network {
  # Configure direct properties listed below.
}
```

## Direct properties

- [ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/advertise_policy/properties/where/virtual_network/ref/): complete subsection reference.

## Next pages

- [where.virtual_network.ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/advertise_policy/properties/where/virtual_network/ref/)
- [where](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/advertise_policy/properties/where/)
- [xcsh_advertise_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/advertise_policy/)
