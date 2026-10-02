---
page_title: "segment_vrf.segment_config.static_routes.static_routes.node_interface"
subcategory: ""
description: "On multinode site, this type holds the information about per node interfaces."
xcsh_docs: {"aliases": ["segment vrf segment config static routes static routes node interface"], "body_bytes": 2432, "body_sha256": "sha256:f32ab4e6872f22c6ed2be2c94466eb667ff6069c61e98cf83c28bda4ce56205c", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:segment_vrf:segment_config:static_routes:static_routes:node_interface:list"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:segment_vrf:segment_config:static_routes:static_routes:node_interface", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:segment_vrf:segment_config:static_routes:static_routes", "path": "documentation/resources/securemesh_site_v2/properties/segment_vrf/segment_config/static_routes/static_routes/node_interface/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2103023112031113-2310110120120230-2222310321031013-1100213230000333-2200212222210032-1212001131211320-3320033012321222-3313133011113023", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-017.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["segment_vrf", "segment_config", "static_routes", "static_routes", "node_interface"], "schema_version": 1, "sections": [{"aliases": ["list"], "anchor": "section", "description": "On a multinode site, this list holds the nodes and corresponding networking_interface.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:segment_vrf:segment_config:static_routes:static_routes:node_interface:list", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["segment_vrf", "segment_config", "static_routes", "static_routes", "node_interface", "list"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/segment_vrf/segment_config/static_routes/static_routes/node_interface/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "On multinode site, this type holds the information about per node interfaces.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# segment_vrf.segment_config.static_routes.static_routes.node_interface

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- [segment_vrf](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/segment_vrf/)
- [segment_vrf.segment_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/segment_vrf/segment_config/)
- [segment_vrf.segment_config.static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/segment_vrf/segment_config/static_routes/)
- [segment_vrf.segment_config.static_routes.static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/segment_vrf/segment_config/static_routes/static_routes/)
- segment_vrf.segment_config.static_routes.static_routes.node_interface

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

On multinode site, this type holds the information about per node interfaces.

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
node_interface {
  # Configure direct properties listed below.
}
```

## Direct properties

- [list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/segment_vrf/segment_config/static_routes/static_routes/node_interface/list/): complete subsection reference.

## Next pages

- [segment_vrf.segment_config.static_routes.static_routes.node_interface.list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/segment_vrf/segment_config/static_routes/static_routes/node_interface/list/)
- [segment_vrf.segment_config.static_routes.static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/segment_vrf/segment_config/static_routes/static_routes/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
