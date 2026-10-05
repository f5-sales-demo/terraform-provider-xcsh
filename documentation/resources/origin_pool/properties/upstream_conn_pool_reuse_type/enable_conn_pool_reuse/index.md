---
page_title: "upstream_conn_pool_reuse_type.enable_conn_pool_reuse"
subcategory: "Load Balancing"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["upstream conn pool reuse type enable conn pool reuse"], "body_bytes": 1392, "body_sha256": "sha256:ff884624aeba86a3a496bd930baf1e7745b3ffad3f415b2fa65cb6a97ccd9b58", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:resources:origin_pool:properties:upstream_conn_pool_reuse_type:enable_conn_pool_reuse", "parent_id": "xcsh-docs:resources:origin_pool:properties:upstream_conn_pool_reuse_type", "path": "documentation/resources/origin_pool/properties/upstream_conn_pool_reuse_type/enable_conn_pool_reuse/index.md", "product": "distributed-cloud", "provider_name": "origin_pool", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-3131222111103001-1022113210301001-2003220322321311-2101301230133010-3001313222333012-1033323302303222-3113132123323330-0121133313120301", "registry_path": "docs/guides/resources--origin_pool--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["upstream_conn_pool_reuse_type", "enable_conn_pool_reuse"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/origin_pool/properties/upstream_conn_pool_reuse_type/enable_conn_pool_reuse/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# upstream_conn_pool_reuse_type.enable_conn_pool_reuse

Breadcrumbs:

- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/)
- [upstream_conn_pool_reuse_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/upstream_conn_pool_reuse_type/)
- upstream_conn_pool_reuse_type.enable_conn_pool_reuse

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable conn pool reuse.

Upstream description:

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

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [upstream_conn_pool_reuse_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/upstream_conn_pool_reuse_type/)
- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/)
