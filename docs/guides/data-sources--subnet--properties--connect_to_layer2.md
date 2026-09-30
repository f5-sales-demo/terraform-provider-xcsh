---
page_title: "connect_to_layer2"
subcategory: ""
description: "connect_to_layer2 for xcsh_subnet."
xcsh_docs: {"aliases": [], "body_bytes": 1217, "body_sha256": "sha256:15a150e926b0d10b7e7331172610c4708eaf9a4adede81572a98c77ba8b9d7df", "canonical_id": "xcsh-docs:data-sources:subnet:properties:connect_to_layer2", "child_ids": ["xcsh-docs:data-sources:subnet:properties:connect_to_layer2:layer2_intf_ref"], "collection_id": "xcsh-docs:data-sources:subnet:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:subnet:properties:connect_to_layer2", "parent_id": "xcsh-docs:data-sources:subnet:reference", "path": "docs/guides/data-sources--subnet--properties--connect_to_layer2.md", "provider_name": "subnet", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["connect_to_layer2"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/subnet/properties/connect_to_layer2/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "connect_to_layer2 for xcsh_subnet.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["subnetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# connect_to_layer2

Breadcrumbs:

- [xcsh_subnet](../data-sources/subnet.md)
- [Property reference](data-sources--subnet--reference.md)
- connect_to_layer2

<a id="section"></a>

Type: `"single"`. Computed.

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

- [connect_to_layer2](data-sources--subnet--properties--connect_to_layer2.md#section)
- [connect_to_slo](data-sources--subnet--properties--connect_to_slo.md#section)
- [isolated_nw](data-sources--subnet--properties--isolated_nw.md#section)

Select alternatives according to the provider validators above.

## Direct properties

- [layer2_intf_ref](data-sources--subnet--properties--connect_to_layer2--layer2_intf_ref.md): complete subsection reference.

## Next pages

- [connect_to_layer2.layer2_intf_ref](data-sources--subnet--properties--connect_to_layer2--layer2_intf_ref.md)
- [Property reference](data-sources--subnet--reference.md)
- [xcsh_subnet](../data-sources/subnet.md)
