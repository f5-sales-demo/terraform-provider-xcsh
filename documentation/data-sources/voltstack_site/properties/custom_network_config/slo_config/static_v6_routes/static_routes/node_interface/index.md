---
page_title: "custom_network_config.slo_config.static_v6_routes.static_routes.node_interface"
subcategory: ""
description: "On multinode site, this type holds the information about per node interfaces."
xcsh_docs: {"aliases": ["custom network config slo config static v6 routes static routes node interface"], "body_bytes": 2422, "body_sha256": "sha256:19d7ebf3d15c40062e452e86a54ce27ce1f12dc5695ec582e176e89dc79665d8", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:slo_config:static_v6_routes:static_routes:node_interface:list"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:slo_config:static_v6_routes:static_routes:node_interface", "parent_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:slo_config:static_v6_routes:static_routes", "path": "documentation/data-sources/voltstack_site/properties/custom_network_config/slo_config/static_v6_routes/static_routes/node_interface/index.md", "product": "distributed-cloud", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1011030001223203-3330312030333133-3130223022301223-3232100220110012-1323220113313213-0231130211130312-0310033011031221-3001112032300313", "registry_path": "docs/guides/data-sources--voltstack_site--reference--group-005.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_network_config", "slo_config", "static_v6_routes", "static_routes", "node_interface"], "schema_version": 1, "sections": [{"aliases": ["list"], "anchor": "section", "description": "On a multinode site, this list holds the nodes and corresponding networking_interface.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:slo_config:static_v6_routes:static_routes:node_interface:list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["custom_network_config", "slo_config", "static_v6_routes", "static_routes", "node_interface", "list"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/voltstack_site/properties/custom_network_config/slo_config/static_v6_routes/static_routes/node_interface/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "On multinode site, this type holds the information about per node interfaces.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_network_config.slo_config.static_v6_routes.static_routes.node_interface

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/)
- [custom_network_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/)
- [custom_network_config.slo_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/slo_config/)
- [custom_network_config.slo_config.static_v6_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/slo_config/static_v6_routes/)
- [custom_network_config.slo_config.static_v6_routes.static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/slo_config/static_v6_routes/static_routes/)
- custom_network_config.slo_config.static_v6_routes.static_routes.node_interface

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

- [list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/slo_config/static_v6_routes/static_routes/node_interface/list/): complete subsection reference.

## Next pages

- [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/slo_config/static_v6_routes/static_routes/node_interface/list/)
- [custom_network_config.slo_config.static_v6_routes.static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/slo_config/static_v6_routes/static_routes/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/)
