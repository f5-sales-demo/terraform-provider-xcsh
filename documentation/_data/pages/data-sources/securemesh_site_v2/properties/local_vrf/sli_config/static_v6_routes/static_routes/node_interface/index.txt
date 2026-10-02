---
page_title: "local_vrf.sli_config.static_v6_routes.static_routes.node_interface"
subcategory: ""
description: "On multinode site, this type holds the information about per node interfaces."
xcsh_docs: {"aliases": ["local vrf sli config static v6 routes static routes node interface"], "body_bytes": 2290, "body_sha256": "sha256:ebac6ecfaac33388325bb6376d7631137531cbd86eac3baaaadfa2df22250a2b", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:securemesh_site_v2:properties:local_vrf:sli_config:static_v6_routes:static_routes:node_interface:list"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:local_vrf:sli_config:static_v6_routes:static_routes:node_interface", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:local_vrf:sli_config:static_v6_routes:static_routes", "path": "documentation/data-sources/securemesh_site_v2/properties/local_vrf/sli_config/static_v6_routes/static_routes/node_interface/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2213120302211010-3213221121001311-0133133230123103-1212022203001332-2100120320210010-1132102323211000-3001221020000330-1321132030203132", "registry_path": "docs/guides/data-sources--securemesh_site_v2--reference--group-011.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["local_vrf", "sli_config", "static_v6_routes", "static_routes", "node_interface"], "schema_version": 1, "sections": [{"aliases": ["list"], "anchor": "section", "description": "On a multinode site, this list holds the nodes and corresponding networking_interface.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:local_vrf:sli_config:static_v6_routes:static_routes:node_interface:list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["local_vrf", "sli_config", "static_v6_routes", "static_routes", "node_interface", "list"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/local_vrf/sli_config/static_v6_routes/static_routes/node_interface/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "On multinode site, this type holds the information about per node interfaces.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# local_vrf.sli_config.static_v6_routes.static_routes.node_interface

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/)
- [local_vrf](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/local_vrf/)
- [local_vrf.sli_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/local_vrf/sli_config/)
- [local_vrf.sli_config.static_v6_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/local_vrf/sli_config/static_v6_routes/)
- [local_vrf.sli_config.static_v6_routes.static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/local_vrf/sli_config/static_v6_routes/static_routes/)
- local_vrf.sli_config.static_v6_routes.static_routes.node_interface

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

- [list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/local_vrf/sli_config/static_v6_routes/static_routes/node_interface/list/): complete subsection reference.

## Next pages

- [local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/local_vrf/sli_config/static_v6_routes/static_routes/node_interface/list/)
- [local_vrf.sli_config.static_v6_routes.static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/local_vrf/sli_config/static_v6_routes/static_routes/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
