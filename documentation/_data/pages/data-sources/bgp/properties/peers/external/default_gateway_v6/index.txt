---
page_title: "peers.external.default_gateway_v6"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["peers external default gateway v6"], "body_bytes": 1047, "body_sha256": "sha256:4a8b8e0d4ed1a2800886ff5d87c7507bc45f76cc2887fcb673bbb40a6c1ad4d4", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bgp:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bgp:properties:peers:external:default_gateway_v6", "parent_id": "xcsh-docs:data-sources:bgp:properties:peers:external", "path": "documentation/data-sources/bgp/properties/peers/external/default_gateway_v6/index.md", "product": "distributed-cloud", "provider_name": "bgp", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-2002320231121031-3100100212113030-0230121333202013-1323223233323022-0131021013202001-0010023003120312-2003102220233200-2321100213130002", "registry_path": "docs/guides/data-sources--bgp--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["peers", "external", "default_gateway_v6"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bgp/properties/peers/external/default_gateway_v6/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["bgpCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# peers.external.default_gateway_v6

Breadcrumbs:

- [xcsh_bgp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/)
- [peers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/)
- [peers.external](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/external/)
- peers.external.default_gateway_v6

<a id="section"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default gateway v6.

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

This is an empty object or choice marker. It has no direct properties.
