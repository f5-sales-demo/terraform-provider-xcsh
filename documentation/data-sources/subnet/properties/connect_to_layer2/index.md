---
page_title: "connect_to_layer2"
subcategory: ""
description: "Configuration parameter for connect to layer2."
xcsh_docs: {"aliases": ["connect to layer2"], "body_bytes": 1394, "body_sha256": "sha256:40112b9c71d6a3d6375d715fab95683bd37299fafd23902ba96f8a8fe887ffca", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:subnet:properties:connect_to_layer2:layer2_intf_ref"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:subnet:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:subnet:properties:connect_to_layer2", "parent_id": "xcsh-docs:data-sources:subnet:reference", "path": "documentation/data-sources/subnet/properties/connect_to_layer2/index.md", "product": "distributed-cloud", "provider_name": "subnet", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-3130121321030111-2013113023133031-1213132113021010-2313220033022320-0122011323332121-3021003231200303-0213001233000220-0202302100030301", "registry_path": "docs/guides/data-sources--subnet--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["connect_to_layer2"], "schema_version": 1, "sections": [{"aliases": ["connect to layer2 layer2 intf ref"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:subnet:properties:connect_to_layer2:layer2_intf_ref", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["connect_to_layer2", "layer2_intf_ref"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/subnet/properties/connect_to_layer2/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Configuration parameter for connect to layer2.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["subnetCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# connect_to_layer2

Breadcrumbs:

- [xcsh_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/subnet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/subnet/properties/)
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

- [connect_to_layer2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/subnet/properties/connect_to_layer2/#section)
- [connect_to_slo](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/subnet/properties/connect_to_slo/#section)
- [isolated_nw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/subnet/properties/isolated_nw/#section)

Select alternatives according to the provider validators above.

## Direct properties

- [layer2_intf_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/subnet/properties/connect_to_layer2/layer2_intf_ref/): complete subsection reference.
