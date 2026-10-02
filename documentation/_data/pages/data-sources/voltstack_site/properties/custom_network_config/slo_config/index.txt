---
page_title: "custom_network_config.slo_config"
subcategory: ""
description: "Site local network configuration."
xcsh_docs: {"aliases": ["custom network config slo config"], "body_bytes": 4271, "body_sha256": "sha256:066afd49b597a2c26fa4b39cb68c804ffb418d18810572c409aa4575c7cbc12c", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:slo_config:dc_cluster_group", "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:slo_config:labels", "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:slo_config:no_dc_cluster_group", "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:slo_config:no_static_routes", "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:slo_config:no_static_v6_routes", "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:slo_config:static_routes", "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:slo_config:static_v6_routes"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:slo_config", "parent_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config", "path": "documentation/data-sources/voltstack_site/properties/custom_network_config/slo_config/index.md", "product": "distributed-cloud", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2012302033131313-0112313333203013-0222312103011303-1202231112003210-0030001133223102-1233101312030000-1112113001223201-3231100201223111", "registry_path": "docs/guides/data-sources--voltstack_site--reference--group-005.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_network_config", "slo_config"], "schema_version": 1, "sections": [{"aliases": ["dc cluster group"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:slo_config:dc_cluster_group", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["custom_network_config", "slo_config", "dc_cluster_group"], "syntax": "attribute", "type": "object"}, {"aliases": ["labels"], "anchor": "section", "description": "Add Labels for this network, these labels can be used in firewall policy.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:slo_config:labels", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["custom_network_config", "slo_config", "labels"], "syntax": "attribute", "type": "object"}, {"aliases": ["no dc cluster group"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:slo_config:no_dc_cluster_group", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_network_config", "slo_config", "no_dc_cluster_group"], "syntax": "attribute", "type": "object"}, {"aliases": ["no static routes"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:slo_config:no_static_routes", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_network_config", "slo_config", "no_static_routes"], "syntax": "attribute", "type": "object"}, {"aliases": ["no static v6 routes"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:slo_config:no_static_v6_routes", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_network_config", "slo_config", "no_static_v6_routes"], "syntax": "attribute", "type": "object"}, {"aliases": ["static routes"], "anchor": "section", "description": "List of static routes.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:slo_config:static_routes", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["custom_network_config", "slo_config", "static_routes"], "syntax": "attribute", "type": "object"}, {"aliases": ["static v6 routes"], "anchor": "section", "description": "List of IPv6 static routes.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:slo_config:static_v6_routes", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["custom_network_config", "slo_config", "static_v6_routes"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/voltstack_site/properties/custom_network_config/slo_config/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Site local network configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_network_config.slo_config

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/)
- [custom_network_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/)
- custom_network_config.slo_config

<a id="section"></a>

Type: `"single"`. Computed.

Site Local Network Configuration. Site local network configuration.

Upstream description:

Site local network configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-dc_cluster_group_choice": "[\"dc_cluster_group\",\"no_dc_cluster_group\"]",
  "x-ves-oneof-field-static_route_choice": "[\"no_static_routes\",\"static_routes\"]",
  "x-ves-oneof-field-static_v6_route_choice": "[\"no_static_v6_routes\",\"static_v6_routes\"]"
}
```

## Direct properties

- [dc_cluster_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/slo_config/dc_cluster_group/): complete subsection reference.

- [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/slo_config/labels/): complete subsection reference.

- [no_dc_cluster_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/slo_config/no_dc_cluster_group/): complete subsection reference.

- [no_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/slo_config/no_static_routes/): complete subsection reference.

- [no_static_v6_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/slo_config/no_static_v6_routes/): complete subsection reference.

- [static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/slo_config/static_routes/): complete subsection reference.

- [static_v6_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/slo_config/static_v6_routes/): complete subsection reference.

## Next pages

- [custom_network_config.slo_config.dc_cluster_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/slo_config/dc_cluster_group/)
- [custom_network_config.slo_config.labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/slo_config/labels/)
- [custom_network_config.slo_config.no_dc_cluster_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/slo_config/no_dc_cluster_group/)
- [custom_network_config.slo_config.no_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/slo_config/no_static_routes/)
- [custom_network_config.slo_config.no_static_v6_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/slo_config/no_static_v6_routes/)
- [custom_network_config.slo_config.static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/slo_config/static_routes/)
- [custom_network_config.slo_config.static_v6_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/slo_config/static_v6_routes/)
- [custom_network_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/)
