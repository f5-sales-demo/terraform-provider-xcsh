---
page_title: "peers.ebgp_multihop_disabled"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["peers ebgp multihop disabled"], "body_bytes": 944, "body_sha256": "sha256:e8408a5d259c2bfdc3f823f8447252bf7b43d13a5a7ab406b9d54bc715a8fb50", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bgp:collection", "completeness": "complete", "id": "xcsh-docs:resources:bgp:properties:peers:ebgp_multihop_disabled", "parent_id": "xcsh-docs:resources:bgp:properties:peers", "path": "documentation/resources/bgp/properties/peers/ebgp_multihop_disabled/index.md", "product": "distributed-cloud", "provider_name": "bgp", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "resources", "registry_anchor": "canonical-0113332302222133-0010030322122122-1123031002202320-2021122212101331-2200032002111131-3023210221322221-1002000223112121-2202021112303120", "registry_path": "docs/guides/resources--bgp--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["peers", "ebgp_multihop_disabled"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bgp/properties/peers/ebgp_multihop_disabled/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": ["bgpCreateRequest"], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# peers.ebgp_multihop_disabled

Breadcrumbs:

- [xcsh_bgp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/)
- [peers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/)
- peers.ebgp_multihop_disabled

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

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
ebgp_multihop_disabled = {}
```

This is an empty object or choice marker. It has no direct properties.
