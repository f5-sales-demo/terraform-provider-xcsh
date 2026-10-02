---
page_title: "sli_to_slo_snat"
subcategory: "Networking"
description: "X-example: \"\" description."
xcsh_docs: {"aliases": ["sli to slo snat"], "body_bytes": 1803, "body_sha256": "sha256:678dd91e4f578757ca17ea048a5042f0c2e4f75675b141546bd5f2176e4ea295", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:network_connector:properties:sli_to_slo_snat:default_gw_snat", "xcsh-docs:data-sources:network_connector:properties:sli_to_slo_snat:interface_ip"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:network_connector:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_connector:properties:sli_to_slo_snat", "parent_id": "xcsh-docs:data-sources:network_connector:reference", "path": "documentation/data-sources/network_connector/properties/sli_to_slo_snat/index.md", "product": "distributed-cloud", "provider_name": "network_connector", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0121123320323123-3213322110002213-2100333022000003-2120210113312111-1331133121020232-0212331020331110-1031122203113102-0201102222130031", "registry_path": "docs/guides/data-sources--network_connector--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["sli_to_slo_snat"], "schema_version": 1, "sections": [{"aliases": ["default gw snat"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:network_connector:properties:sli_to_slo_snat:default_gw_snat", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["sli_to_slo_snat", "default_gw_snat"], "syntax": "attribute", "type": "object"}, {"aliases": ["interface ip"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:network_connector:properties:sli_to_slo_snat:interface_ip", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["sli_to_slo_snat", "interface_ip"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_connector/properties/sli_to_slo_snat/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "X-example: \"\" description.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["network_connectorCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# sli_to_slo_snat

Breadcrumbs:

- [xcsh_network_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/properties/)
- sli_to_slo_snat

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for sli to slo snat.

Upstream description:

X-example: "" description.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-pool_choice": "[\"interface_ip\"]",
  "x-ves-oneof-field-routing_choice": "[\"default_gw_snat\"]"
}
```

## Direct properties

- [default_gw_snat](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/properties/sli_to_slo_snat/default_gw_snat/): complete subsection reference.

- [interface_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/properties/sli_to_slo_snat/interface_ip/): complete subsection reference.

## Next pages

- [sli_to_slo_snat.default_gw_snat](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/properties/sli_to_slo_snat/default_gw_snat/)
- [sli_to_slo_snat.interface_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/properties/sli_to_slo_snat/interface_ip/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/properties/)
- [xcsh_network_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/)
