---
page_title: "upstream_conn_pool_reuse_type"
subcategory: ""
description: "Select upstream connection pool reuse state for every downstream connection. This configuration choice is for HTTP(S) LB only."
xcsh_docs: {"aliases": ["upstream conn pool reuse type"], "body_bytes": 1923, "body_sha256": "sha256:5330f1bf957d7f11931566e7a4f595ab6c8b6ad6c4178f43350e7f44faa07878", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:cluster:properties:upstream_conn_pool_reuse_type:disable_conn_pool_reuse", "xcsh-docs:data-sources:cluster:properties:upstream_conn_pool_reuse_type:enable_conn_pool_reuse"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:cluster:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cluster:properties:upstream_conn_pool_reuse_type", "parent_id": "xcsh-docs:data-sources:cluster:reference", "path": "documentation/data-sources/cluster/properties/upstream_conn_pool_reuse_type/index.md", "product": "distributed-cloud", "provider_name": "cluster", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-0313001211201301-0220311332132120-1031132020332300-1331101230303030-2110211232132032-3032132311122313-0201333030230221-2110023012110012", "registry_path": "docs/guides/data-sources--cluster--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["upstream_conn_pool_reuse_type"], "schema_version": 1, "sections": [{"aliases": ["upstream conn pool reuse type disable conn pool reuse"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:cluster:properties:upstream_conn_pool_reuse_type:disable_conn_pool_reuse", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["upstream_conn_pool_reuse_type", "disable_conn_pool_reuse"], "syntax": "attribute", "type": "object"}, {"aliases": ["upstream conn pool reuse type enable conn pool reuse"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:cluster:properties:upstream_conn_pool_reuse_type:enable_conn_pool_reuse", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["upstream_conn_pool_reuse_type", "enable_conn_pool_reuse"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cluster/properties/upstream_conn_pool_reuse_type/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Select upstream connection pool reuse state for every downstream connection. This configuration choice is for HTTP(S) LB only.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["clusterCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# upstream_conn_pool_reuse_type

Breadcrumbs:

- [xcsh_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/properties/)
- upstream_conn_pool_reuse_type

<a id="section"></a>

Type: `"single"`. Computed.

Select upstream connection pool reuse state for every downstream connection. This configuration
choice is for HTTP(S) LB only.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-map_downstream_to_upstream_conn_pool_type": "[\"disable_conn_pool_reuse\",\"enable_conn_pool_reuse\"]"
}
```

## Direct properties

- [disable_conn_pool_reuse](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/properties/upstream_conn_pool_reuse_type/disable_conn_pool_reuse/): complete subsection reference.

- [enable_conn_pool_reuse](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/properties/upstream_conn_pool_reuse_type/enable_conn_pool_reuse/): complete subsection reference.

## Next pages

- [upstream_conn_pool_reuse_type.disable_conn_pool_reuse](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/properties/upstream_conn_pool_reuse_type/disable_conn_pool_reuse/)
- [upstream_conn_pool_reuse_type.enable_conn_pool_reuse](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/properties/upstream_conn_pool_reuse_type/enable_conn_pool_reuse/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/properties/)
- [xcsh_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/)
