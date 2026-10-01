---
page_title: "where.virtual_network"
subcategory: "Networking"
description: "where.virtual_network for xcsh_endpoint."
xcsh_docs: {"aliases": [], "body_bytes": 1658, "body_sha256": "sha256:b48a6dd65583446c7eb055af5fcefbe29d0480193112141a73712fb307aeacde", "child_ids": ["xcsh-docs:resources:endpoint:properties:where:virtual_network:ref"], "collection_id": "xcsh-docs:resources:endpoint:collection", "completeness": "complete", "id": "xcsh-docs:resources:endpoint:properties:where:virtual_network", "parent_id": "xcsh-docs:resources:endpoint:properties:where", "path": "documentation/resources/endpoint/properties/where/virtual_network/index.md", "provider_name": "endpoint", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["where", "virtual_network"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/endpoint/properties/where/virtual_network/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "where.virtual_network for xcsh_endpoint.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["endpointCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# where.virtual_network

Breadcrumbs:

- [xcsh_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/)
- [where](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/where/)
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

- [ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/where/virtual_network/ref/): complete subsection reference.

## Next pages

- [where.virtual_network.ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/where/virtual_network/ref/)
- [where](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/where/)
- [xcsh_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/)
