---
page_title: "where.virtual_network"
subcategory: ""
description: "where.virtual_network for xcsh_discovery."
xcsh_docs: {"aliases": [], "body_bytes": 1667, "body_sha256": "sha256:121b99f0b68607781b64ffc09c4d896ec4be5255738b5ce4566fc8c5b92c8e24", "child_ids": ["xcsh-docs:resources:discovery:properties:where:virtual_network:ref"], "collection_id": "xcsh-docs:resources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:resources:discovery:properties:where:virtual_network", "parent_id": "xcsh-docs:resources:discovery:properties:where", "path": "documentation/resources/discovery/properties/where/virtual_network/index.md", "provider_name": "discovery", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["where", "virtual_network"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/discovery/properties/where/virtual_network/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "where.virtual_network for xcsh_discovery.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["discoveryCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# where.virtual_network

Breadcrumbs:

- [xcsh_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/)
- [where](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/where/)
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

- [ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/where/virtual_network/ref/): complete subsection reference.

## Next pages

- [where.virtual_network.ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/where/virtual_network/ref/)
- [where](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/where/)
- [xcsh_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/)
