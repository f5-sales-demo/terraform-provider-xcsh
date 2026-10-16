---
page_title: "lb_algorithm.round_robin"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["lb algorithm round robin"], "body_bytes": 1026, "body_sha256": "sha256:997b7369207d30788fceeaa7b41ec781cc29a5e8d7181702045606174e251ff2", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:bigip_http_proxy:properties:lb_algorithm:round_robin", "parent_id": "xcsh-docs:resources:bigip_http_proxy:properties:lb_algorithm", "path": "documentation/resources/bigip_http_proxy/properties/lb_algorithm/round_robin/index.md", "product": "distributed-cloud", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "resources", "registry_anchor": "canonical-1032322331030110-3003003122311201-2203010003020020-0033132123302200-2213122211132021-0133023330113310-3102121331113002-0221302122230131", "registry_path": "docs/guides/resources--bigip_http_proxy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["lb_algorithm", "round_robin"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bigip_http_proxy/properties/lb_algorithm/round_robin/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# lb_algorithm.round_robin

Breadcrumbs:

- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/)
- [lb_algorithm](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/lb_algorithm/)
- lb_algorithm.round_robin

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for round robin.

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
round_robin {}
```

This is an empty object or choice marker. It has no direct properties.
