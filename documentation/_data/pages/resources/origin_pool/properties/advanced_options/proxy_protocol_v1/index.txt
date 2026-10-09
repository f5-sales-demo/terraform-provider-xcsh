---
page_title: "advanced_options.proxy_protocol_v1"
subcategory: "Load Balancing"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["advanced options proxy protocol v1"], "body_bytes": 1033, "body_sha256": "sha256:7b0763669805b518924831f4cdceaf57d83de6df06c86fc9158831b3abc805a4", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:resources:origin_pool:properties:advanced_options:proxy_protocol_v1", "parent_id": "xcsh-docs:resources:origin_pool:properties:advanced_options", "path": "documentation/resources/origin_pool/properties/advanced_options/proxy_protocol_v1/index.md", "product": "distributed-cloud", "provider_name": "origin_pool", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-2231302013311332-2302020300210101-3013021032120300-1322020323111001-1021312211133123-0001210223111212-2210221303330123-2132302330031012", "registry_path": "docs/guides/resources--origin_pool--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["advanced_options", "proxy_protocol_v1"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/origin_pool/properties/advanced_options/proxy_protocol_v1/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["origin_poolCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# advanced_options.proxy_protocol_v1

Breadcrumbs:

- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/)
- [advanced_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/)
- advanced_options.proxy_protocol_v1

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for proxy protocol v1.

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
proxy_protocol_v1 = {}
```

This is an empty object or choice marker. It has no direct properties.
