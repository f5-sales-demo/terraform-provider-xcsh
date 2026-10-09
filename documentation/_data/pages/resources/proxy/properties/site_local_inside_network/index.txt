---
page_title: "site_local_inside_network"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["site local inside network"], "body_bytes": 1300, "body_sha256": "sha256:1bee27f89f54bc702a2470d23bac35932253916157bfc50f24f42030a0057066", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:proxy:properties:site_local_inside_network", "parent_id": "xcsh-docs:resources:proxy:reference", "path": "documentation/resources/proxy/properties/site_local_inside_network/index.md", "product": "distributed-cloud", "provider_name": "proxy", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-2300220130222220-0321012100023110-3312120100301031-2332310123110321-0222103331130303-0122320111223322-2100122213201320-3222032202331000", "registry_path": "docs/guides/resources--proxy--reference--group-005.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["site_local_inside_network"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/proxy/properties/site_local_inside_network/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["proxyCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# site_local_inside_network

Breadcrumbs:

- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/)
- site_local_inside_network

<a id="section"></a>

Type: `["object", {}]`. Optional.

\[OneOf: site\_local\_inside\_network, site\_local\_network\] Enable this option

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

- [site_local_inside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/site_local_inside_network/#section)
- [site_local_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/site_local_network/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
site_local_inside_network = {}
```

This is an empty object or choice marker. It has no direct properties.
