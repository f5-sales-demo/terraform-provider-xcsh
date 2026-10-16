---
page_title: "upstream_conn_pool_reuse_type.enable_conn_pool_reuse"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["upstream conn pool reuse type enable conn pool reuse"], "body_bytes": 1089, "body_sha256": "sha256:98e563ab4887badcdc2e9adb9286ab54e4860974bf0b70c1321a2dde525f4afd", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cluster:collection", "completeness": "complete", "id": "xcsh-docs:resources:cluster:properties:upstream_conn_pool_reuse_type:enable_conn_pool_reuse", "parent_id": "xcsh-docs:resources:cluster:properties:upstream_conn_pool_reuse_type", "path": "documentation/resources/cluster/properties/upstream_conn_pool_reuse_type/enable_conn_pool_reuse/index.md", "product": "distributed-cloud", "provider_name": "cluster", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "resources", "registry_anchor": "canonical-2303230031132101-0221312032202231-3131333202122000-0321210220013310-2012022020012310-2203211202111212-1223313210332112-0131300233211203", "registry_path": "docs/guides/resources--cluster--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["upstream_conn_pool_reuse_type", "enable_conn_pool_reuse"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cluster/properties/upstream_conn_pool_reuse_type/enable_conn_pool_reuse/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["clusterCreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# upstream_conn_pool_reuse_type.enable_conn_pool_reuse

Breadcrumbs:

- [xcsh_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/)
- [upstream_conn_pool_reuse_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/upstream_conn_pool_reuse_type/)
- upstream_conn_pool_reuse_type.enable_conn_pool_reuse

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable conn pool reuse.

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
enable_conn_pool_reuse = {}
```

This is an empty object or choice marker. It has no direct properties.
