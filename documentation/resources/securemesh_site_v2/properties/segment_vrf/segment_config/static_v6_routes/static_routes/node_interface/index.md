---
page_title: "segment_vrf.segment_config.static_v6_routes.static_routes.node_interface"
subcategory: ""
description: "On multinode site, this type holds the information about per node interfaces."
xcsh_docs: {"aliases": ["segment vrf segment config static v6 routes static routes node interface"], "body_bytes": 2465, "body_sha256": "sha256:a5414f6594ca2cf33e978731a066450c181a5558e7d274256af94be252f4aef5", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:segment_vrf:segment_config:static_v6_routes:static_routes:node_interface:list"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:segment_vrf:segment_config:static_v6_routes:static_routes:node_interface", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:segment_vrf:segment_config:static_v6_routes:static_routes", "path": "documentation/resources/securemesh_site_v2/properties/segment_vrf/segment_config/static_v6_routes/static_routes/node_interface/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2211232301311022-3313233111032330-0111001103333230-2231233131203201-1313200131121321-1203110002132233-2330010020301022-0032021102132233", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-017.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["segment_vrf", "segment_config", "static_v6_routes", "static_routes", "node_interface"], "schema_version": 1, "sections": [{"aliases": ["list"], "anchor": "section", "description": "On a multinode site, this list holds the nodes and corresponding networking_interface.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:segment_vrf:segment_config:static_v6_routes:static_routes:node_interface:list", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["segment_vrf", "segment_config", "static_v6_routes", "static_routes", "node_interface", "list"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/segment_vrf/segment_config/static_v6_routes/static_routes/node_interface/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "On multinode site, this type holds the information about per node interfaces.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# segment_vrf.segment_config.static_v6_routes.static_routes.node_interface

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- [segment_vrf](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/segment_vrf/)
- [segment_vrf.segment_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/segment_vrf/segment_config/)
- [segment_vrf.segment_config.static_v6_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/segment_vrf/segment_config/static_v6_routes/)
- [segment_vrf.segment_config.static_v6_routes.static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/segment_vrf/segment_config/static_v6_routes/static_routes/)
- segment_vrf.segment_config.static_v6_routes.static_routes.node_interface

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

- [list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/segment_vrf/segment_config/static_v6_routes/static_routes/node_interface/list/): complete subsection reference.

## Next pages

- [segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/segment_vrf/segment_config/static_v6_routes/static_routes/node_interface/list/)
- [segment_vrf.segment_config.static_v6_routes.static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/segment_vrf/segment_config/static_v6_routes/static_routes/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
