---
page_title: "aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["aws not managed node list interface list ipv6 auto config router stateful automatic from end"], "body_bytes": 2325, "body_sha256": "sha256:042ffb4559bcc4260630a84f34c8ec84b3484b35593a4cf5363269986bb84662", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:aws:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful:automatic_from_end", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:aws:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful", "path": "documentation/resources/securemesh_site_v2/properties/aws/not_managed/node_list/interface_list/ipv6_auto_config/router/stateful/automatic_from_end/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-0110311100313221-0030311223223302-3032023030021201-3032213032332310-2100223210110101-2322232130001130-2112022320323032-3300132311100123", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["aws", "not_managed", "node_list", "interface_list", "ipv6_auto_config", "router", "stateful", "automatic_from_end"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/aws/not_managed/node_list/interface_list/ipv6_auto_config/router/stateful/automatic_from_end/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- [aws](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/aws/)
- [aws.not_managed](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/aws/not_managed/)
- [aws.not_managed.node_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/aws/not_managed/node_list/)
- [aws.not_managed.node_list.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/aws/not_managed/node_list/interface_list/)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/aws/not_managed/node_list/interface_list/ipv6_auto_config/)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/aws/not_managed/node_list/interface_list/ipv6_auto_config/router/)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/aws/not_managed/node_list/interface_list/ipv6_auto_config/router/stateful/)
- aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for automatic from end.

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
automatic_from_end = {}
```

This is an empty object or choice marker. It has no direct properties.
