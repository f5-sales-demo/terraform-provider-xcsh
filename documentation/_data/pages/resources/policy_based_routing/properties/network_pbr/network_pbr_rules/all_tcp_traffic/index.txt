---
page_title: "network_pbr.network_pbr_rules.all_tcp_traffic"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["network pbr network pbr rules all tcp traffic"], "body_bytes": 1241, "body_sha256": "sha256:f66751d6a0fd7baa6f638265ae94b4985a4ca34b198156d11453436108bb6881", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:policy_based_routing:collection", "completeness": "complete", "id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:all_tcp_traffic", "parent_id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules", "path": "documentation/resources/policy_based_routing/properties/network_pbr/network_pbr_rules/all_tcp_traffic/index.md", "product": "distributed-cloud", "provider_name": "policy_based_routing", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-1121223032020001-2302301003103121-1220032100221132-0021211113110122-0020020121030330-0132100212321122-1130031010230300-3233123332003331", "registry_path": "docs/guides/resources--policy_based_routing--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["network_pbr", "network_pbr_rules", "all_tcp_traffic"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/policy_based_routing/properties/network_pbr/network_pbr_rules/all_tcp_traffic/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["policy_based_routingCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# network_pbr.network_pbr_rules.all_tcp_traffic

Breadcrumbs:

- [xcsh_policy_based_routing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/properties/)
- [network_pbr](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/properties/network_pbr/)
- [network_pbr.network_pbr_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/properties/network_pbr/network_pbr_rules/)
- network_pbr.network_pbr_rules.all_tcp_traffic

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for all tcp traffic.

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
all_tcp_traffic = {}
```

This is an empty object or choice marker. It has no direct properties.
