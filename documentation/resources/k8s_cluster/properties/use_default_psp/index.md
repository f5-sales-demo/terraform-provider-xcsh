---
page_title: "use_default_psp"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["use default psp"], "body_bytes": 934, "body_sha256": "sha256:72525ed913044ccca5459d7bb494d787a4185fbbde4c3df157ddc8fcd581e5d5", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:k8s_cluster:collection", "completeness": "complete", "id": "xcsh-docs:resources:k8s_cluster:properties:use_default_psp", "parent_id": "xcsh-docs:resources:k8s_cluster:reference", "path": "documentation/resources/k8s_cluster/properties/use_default_psp/index.md", "product": "distributed-cloud", "provider_name": "k8s_cluster", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-0301302001120301-3120213203003133-3013101332200030-3303203223000211-1100230021000103-3230002012102133-0311030302011303-2322000302222001", "registry_path": "docs/guides/resources--k8s_cluster--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["use_default_psp"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/k8s_cluster/properties/use_default_psp/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["k8s_clusterCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# use_default_psp

Breadcrumbs:

- [xcsh_k8s_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/properties/)
- use_default_psp

<a id="section"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for use default psp. Defaults to \`map\[\]\`. Server applies default when
omitted.

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
use_default_psp = {}
```

This is an empty object or choice marker. It has no direct properties.
