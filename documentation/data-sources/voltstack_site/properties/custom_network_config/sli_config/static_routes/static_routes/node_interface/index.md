---
page_title: "custom_network_config.sli_config.static_routes.static_routes.node_interface"
subcategory: ""
description: "On multinode site, this type holds the information about per node interfaces."
xcsh_docs: {"aliases": ["custom network config sli config static routes static routes node interface"], "body_bytes": 2389, "body_sha256": "sha256:32ea04b4a816403ace3a58f35fe62295a348186c2f438ebb1350605551455ca0", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:sli_config:static_routes:static_routes:node_interface:list"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:sli_config:static_routes:static_routes:node_interface", "parent_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:sli_config:static_routes:static_routes", "path": "documentation/data-sources/voltstack_site/properties/custom_network_config/sli_config/static_routes/static_routes/node_interface/index.md", "product": "distributed-cloud", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1020322103333333-0201000202013023-2032212303332131-0331330302011230-3003010111231231-1133311000333203-0331110210121331-1303312000330221", "registry_path": "docs/guides/data-sources--voltstack_site--reference--group-005.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_network_config", "sli_config", "static_routes", "static_routes", "node_interface"], "schema_version": 1, "sections": [{"aliases": ["list"], "anchor": "section", "description": "On a multinode site, this list holds the nodes and corresponding networking_interface.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:sli_config:static_routes:static_routes:node_interface:list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["custom_network_config", "sli_config", "static_routes", "static_routes", "node_interface", "list"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/voltstack_site/properties/custom_network_config/sli_config/static_routes/static_routes/node_interface/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "On multinode site, this type holds the information about per node interfaces.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_network_config.sli_config.static_routes.static_routes.node_interface

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/)
- [custom_network_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/)
- [custom_network_config.sli_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/sli_config/)
- [custom_network_config.sli_config.static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/sli_config/static_routes/)
- [custom_network_config.sli_config.static_routes.static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/sli_config/static_routes/static_routes/)
- custom_network_config.sli_config.static_routes.static_routes.node_interface

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

- [list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/sli_config/static_routes/static_routes/node_interface/list/): complete subsection reference.

## Next pages

- [custom_network_config.sli_config.static_routes.static_routes.node_interface.list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/sli_config/static_routes/static_routes/node_interface/list/)
- [custom_network_config.sli_config.static_routes.static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/sli_config/static_routes/static_routes/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/)
