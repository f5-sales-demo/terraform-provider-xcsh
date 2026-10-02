---
page_title: "connect_to_layer2"
subcategory: ""
description: "Configuration parameter for connect to layer2."
xcsh_docs: {"aliases": ["connect to layer2"], "body_bytes": 1876, "body_sha256": "sha256:76ba1a4fbc0d45686adfa7586a0595dcc2b61b6e4949e49bf4d4622540918a0b", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:subnet:properties:connect_to_layer2:layer2_intf_ref"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:subnet:collection", "completeness": "complete", "id": "xcsh-docs:resources:subnet:properties:connect_to_layer2", "parent_id": "xcsh-docs:resources:subnet:reference", "path": "documentation/resources/subnet/properties/connect_to_layer2/index.md", "product": "distributed-cloud", "provider_name": "subnet", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3211320301332132-0311221202303332-0010032231210300-2000132130330101-1303202203222131-1231323330211010-0321233002313030-1020202233221222", "registry_path": "docs/guides/resources--subnet--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["connect_to_layer2"], "schema_version": 1, "sections": [{"aliases": ["layer2 intf ref"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:subnet:properties:connect_to_layer2:layer2_intf_ref", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-connect_to_layer2--layer2_intf_ref--name", "enforcement": "provider-schema", "group": "connect_to_layer2.layer2_intf_ref:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:subnet:properties:connect_to_layer2:layer2_intf_ref", "type": "requires"}], "schema_path": ["connect_to_layer2", "layer2_intf_ref"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/subnet/properties/connect_to_layer2/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Configuration parameter for connect to layer2.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["subnetCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# connect_to_layer2

Breadcrumbs:

- [xcsh_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/subnet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/subnet/properties/)
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

- [connect_to_layer2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/subnet/properties/connect_to_layer2/#section)
- [connect_to_slo](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/subnet/properties/connect_to_slo/#section)
- [isolated_nw](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/subnet/properties/isolated_nw/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
connect_to_layer2 {
  # Configure direct properties listed below.
}
```

## Direct properties

- [layer2_intf_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/subnet/properties/connect_to_layer2/layer2_intf_ref/): complete subsection reference.

## Next pages

- [connect_to_layer2.layer2_intf_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/subnet/properties/connect_to_layer2/layer2_intf_ref/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/subnet/properties/)
- [xcsh_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/subnet/)
