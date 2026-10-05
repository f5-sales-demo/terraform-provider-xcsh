---
page_title: "local_vrf.sli_config.static_routes.static_routes.node_interface"
subcategory: ""
description: "On multinode site, this type holds the information about per node interfaces."
xcsh_docs: {"aliases": ["local vrf sli config static routes static routes node interface"], "body_bytes": 2257, "body_sha256": "sha256:a10ef4d299ffa28e688027bac2fcf346ec23d8b35a42d7363a4fbefd279949bc", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:securemesh_site_v2:properties:local_vrf:sli_config:static_routes:static_routes:node_interface:list"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:local_vrf:sli_config:static_routes:static_routes:node_interface", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:local_vrf:sli_config:static_routes:static_routes", "path": "documentation/data-sources/securemesh_site_v2/properties/local_vrf/sli_config/static_routes/static_routes/node_interface/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-1033103223101320-3303122123233322-2221100102000133-2223312022121311-2221001312310032-2120003113222020-0223203010302223-0121111102112101", "registry_path": "docs/guides/data-sources--securemesh_site_v2--reference--group-011.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["local_vrf", "sli_config", "static_routes", "static_routes", "node_interface"], "schema_version": 1, "sections": [{"aliases": ["local vrf sli config static routes static routes node interface list"], "anchor": "section", "description": "On a multinode site, this list holds the nodes and corresponding networking_interface.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:local_vrf:sli_config:static_routes:static_routes:node_interface:list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["local_vrf", "sli_config", "static_routes", "static_routes", "node_interface", "list"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/local_vrf/sli_config/static_routes/static_routes/node_interface/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "On multinode site, this type holds the information about per node interfaces.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# local_vrf.sli_config.static_routes.static_routes.node_interface

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/)
- [local_vrf](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/local_vrf/)
- [local_vrf.sli_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/local_vrf/sli_config/)
- [local_vrf.sli_config.static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/local_vrf/sli_config/static_routes/)
- [local_vrf.sli_config.static_routes.static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/local_vrf/sli_config/static_routes/static_routes/)
- local_vrf.sli_config.static_routes.static_routes.node_interface

<a id="section"></a>

Type: `"single"`. Computed.

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

## Direct properties

- [list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/local_vrf/sli_config/static_routes/static_routes/node_interface/list/): complete subsection reference.

## Next pages

- [local_vrf.sli_config.static_routes.static_routes.node_interface.list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/local_vrf/sli_config/static_routes/static_routes/node_interface/list/)
- [local_vrf.sli_config.static_routes.static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/local_vrf/sli_config/static_routes/static_routes/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
