---
page_title: "log_receiver_with_net.use_management_network"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["log receiver with net use management network"], "body_bytes": 1101, "body_sha256": "sha256:f0a857c4a7ed76d54dd8ca7a9bff4f7c88d4d49b79cea120323ac116e0198223", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:log_receiver_with_net:use_management_network", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:log_receiver_with_net", "path": "documentation/resources/securemesh_site_v2/properties/log_receiver_with_net/use_management_network/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-1000230333312100-2312133011203203-2301313000031303-3213013221120231-3211023003322120-2020313121120300-3121213103011210-0023011121111220", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-011.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["log_receiver_with_net", "use_management_network"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/log_receiver_with_net/use_management_network/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# log_receiver_with_net.use_management_network

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- [log_receiver_with_net](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/log_receiver_with_net/)
- log_receiver_with_net.use_management_network

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for use management network.

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
use_management_network = {}
```

This is an empty object or choice marker. It has no direct properties.
