---
page_title: "peers.external.family_inet.enable.aggregation.options.summary_only"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["peers external family inet enable aggregation options summary only"], "body_bytes": 1833, "body_sha256": "sha256:638a8c7128270a93a816cb897e0ad80a6b285e60088b6dad987cdf1a827e7637", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bgp:collection", "completeness": "complete", "id": "xcsh-docs:resources:bgp:properties:peers:external:family_inet:enable:aggregation:options:summary_only", "parent_id": "xcsh-docs:resources:bgp:properties:peers:external:family_inet:enable:aggregation:options", "path": "documentation/resources/bgp/properties/peers/external/family_inet/enable/aggregation/options/summary_only/index.md", "product": "distributed-cloud", "provider_name": "bgp", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "resources", "registry_anchor": "canonical-2131301300010020-3333123330223310-2121203323032233-1233320303222232-2003113320200011-2201320112301012-2002121303102130-1131232211003200", "registry_path": "docs/guides/resources--bgp--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["peers", "external", "family_inet", "enable", "aggregation", "options", "summary_only"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bgp/properties/peers/external/family_inet/enable/aggregation/options/summary_only/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": ["bgpCreateRequest"], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# peers.external.family_inet.enable.aggregation.options.summary_only

Breadcrumbs:

- [xcsh_bgp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/)
- [peers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/)
- [peers.external](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/external/)
- [peers.external.family_inet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/external/family_inet/)
- [peers.external.family_inet.enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/external/family_inet/enable/)
- [peers.external.family_inet.enable.aggregation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/external/family_inet/enable/aggregation/)
- [peers.external.family_inet.enable.aggregation.options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/external/family_inet/enable/aggregation/options/)
- peers.external.family_inet.enable.aggregation.options.summary_only

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for summary only.

Additional upstream details:

This can be used for messages where no values are needed.

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
summary_only {}
```

This is an empty object or choice marker. It has no direct properties.
