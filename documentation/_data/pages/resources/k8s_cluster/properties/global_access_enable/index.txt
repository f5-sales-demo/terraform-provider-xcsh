---
page_title: "global_access_enable"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["global access enable"], "body_bytes": 1353, "body_sha256": "sha256:0df9a66c050f78de7a3852fff14a75d8494ce41494159ec55f60f35b27b1a0b5", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:k8s_cluster:collection", "completeness": "complete", "id": "xcsh-docs:resources:k8s_cluster:properties:global_access_enable", "parent_id": "xcsh-docs:resources:k8s_cluster:reference", "path": "documentation/resources/k8s_cluster/properties/global_access_enable/index.md", "product": "distributed-cloud", "provider_name": "k8s_cluster", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-1110033101330033-3012032322223221-3012123023320130-1102100010210112-1112321122232211-0033011130032123-1102210111203133-1110010101001011", "registry_path": "docs/guides/resources--k8s_cluster--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["global_access_enable"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/k8s_cluster/properties/global_access_enable/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["k8s_clusterCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# global_access_enable

Breadcrumbs:

- [xcsh_k8s_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/properties/)
- global_access_enable

<a id="section"></a>

Type: `["object", {}]`. Optional.

\[OneOf: global\_access\_enable, no\_global\_access; Default: no\_global\_access\] Configuration
parameter for global access enable.

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

- [global_access_enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/properties/global_access_enable/#section)
- [no_global_access](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/properties/no_global_access/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
global_access_enable = {}
```

This is an empty object or choice marker. It has no direct properties.
