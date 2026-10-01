---
page_title: "default_storage_class"
subcategory: ""
description: "default_storage_class for xcsh_fleet."
xcsh_docs: {"aliases": [], "body_bytes": 1569, "body_sha256": "sha256:b0acebf15e5c314e294e03b12bff2a98137b6e830e08da78a9d69f0744851fb5", "child_ids": [], "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:default_storage_class", "parent_id": "xcsh-docs:resources:fleet:reference", "path": "documentation/resources/fleet/properties/default_storage_class/index.md", "provider_name": "fleet", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["default_storage_class"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/default_storage_class/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "default_storage_class for xcsh_fleet.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_storage_class

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/)
- default_storage_class

<a id="section"></a>

Type: `["object", {}]`. Optional.

\[OneOf: default\_storage\_class, storage\_class\_list; Default: default\_storage\_class\]
Configuration parameter for default storage class.

Upstream description:

This can be used for messages where no values are needed.

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

OneOf alternatives in this subsection:

- [default_storage_class](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/default_storage_class/#section)
- [storage_class_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_class_list/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
default_storage_class = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/)
- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
