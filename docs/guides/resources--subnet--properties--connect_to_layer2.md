---
page_title: "connect_to_layer2"
subcategory: ""
description: "connect_to_layer2 for xcsh_subnet."
xcsh_docs: {"aliases": [], "body_bytes": 1415, "body_sha256": "sha256:fc6a062961f33d57c3d770d08084f3ab300a9b704e5592166692adadc6efbd8b", "canonical_id": "xcsh-docs:resources:subnet:properties:connect_to_layer2", "child_ids": ["xcsh-docs:resources:subnet:properties:connect_to_layer2:layer2_intf_ref"], "collection_id": "xcsh-docs:resources:subnet:collection", "completeness": "complete", "id": "xcsh-docs:resources:subnet:properties:connect_to_layer2", "parent_id": "xcsh-docs:resources:subnet:reference", "path": "docs/guides/resources--subnet--properties--connect_to_layer2.md", "provider_name": "subnet", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["connect_to_layer2"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/subnet/properties/connect_to_layer2/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "connect_to_layer2 for xcsh_subnet.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["subnetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# connect_to_layer2

Breadcrumbs:

- [xcsh_subnet](../resources/subnet.md)
- [Property reference](resources--subnet--reference.md)
- connect_to_layer2

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: connect\_to\_layer2, connect\_to\_slo, isolated\_nw\] Configuration parameter for connect
to layer2.

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

- [connect_to_layer2](resources--subnet--properties--connect_to_layer2.md#section)
- [connect_to_slo](resources--subnet--properties--connect_to_slo.md#section)
- [isolated_nw](resources--subnet--properties--isolated_nw.md#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
connect_to_layer2 {
  # Configure direct properties listed below.
}
```

## Direct properties

- [layer2_intf_ref](resources--subnet--properties--connect_to_layer2--layer2_intf_ref.md): complete subsection reference.

## Next pages

- [connect_to_layer2.layer2_intf_ref](resources--subnet--properties--connect_to_layer2--layer2_intf_ref.md)
- [Property reference](resources--subnet--reference.md)
- [xcsh_subnet](../resources/subnet.md)
