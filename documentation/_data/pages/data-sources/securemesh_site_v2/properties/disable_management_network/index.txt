---
page_title: "disable_management_network"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["disable management network"], "body_bytes": 1410, "body_sha256": "sha256:e5965a89d8640bec9f9f2de84bb95af692a36d590f6343f00a02f7dee29db340", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:disable_management_network", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:reference", "path": "documentation/data-sources/securemesh_site_v2/properties/disable_management_network/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-3230123220103000-0033000120201100-0013210112221103-0321230123310232-2021103220331321-3202332010302123-1100131023010011-1121000100310033", "registry_path": "docs/guides/data-sources--securemesh_site_v2--reference--group-007.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["disable_management_network"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/disable_management_network/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# disable_management_network

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/)
- disable_management_network

<a id="section"></a>

Type: `["object", {}]`. Computed.

\[OneOf: disable\_management\_network, enable\_management\_network; Default:
disable\_management\_network\] Configuration parameter for disable management network.

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

OneOf alternatives in this subsection:

- [disable_management_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/disable_management_network/#section)
- [enable_management_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/enable_management_network/#section)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.
