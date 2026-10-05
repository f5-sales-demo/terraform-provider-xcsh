---
page_title: "cluster_wide_app_list"
subcategory: ""
description: "List of cluster wide applications."
xcsh_docs: {"aliases": ["cluster wide app list"], "body_bytes": 1860, "body_sha256": "sha256:47c1b0a377c5516b9166874fbc4167176d618979ddb67929d4f2bdf87b7682f4", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:k8s_cluster:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:k8s_cluster:properties:cluster_wide_app_list", "parent_id": "xcsh-docs:data-sources:k8s_cluster:reference", "path": "documentation/data-sources/k8s_cluster/properties/cluster_wide_app_list/index.md", "product": "distributed-cloud", "provider_name": "k8s_cluster", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-3012002212023212-0312020122110302-1031000322301012-1133301000011230-1113222032032232-2113012000333012-3031101320033120-2033112013002132", "registry_path": "docs/guides/data-sources--k8s_cluster--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["cluster_wide_app_list"], "schema_version": 1, "sections": [{"aliases": ["cluster wide app list cluster wide apps"], "anchor": "section", "description": "List of cluster wide applications.", "document_id": "xcsh-docs:data-sources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["cluster_wide_app_list", "cluster_wide_apps"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/k8s_cluster/properties/cluster_wide_app_list/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "List of cluster wide applications.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["k8s_clusterCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cluster_wide_app_list

Breadcrumbs:

- [xcsh_k8s_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster/properties/)
- cluster_wide_app_list

<a id="section"></a>

Type: `"single"`. Computed.

\[OneOf: cluster\_wide\_app\_list, no\_cluster\_wide\_apps; Default: no\_cluster\_wide\_apps\]
Cluster Wide Application List. List of cluster wide applications.

Upstream description:

List of cluster wide applications.

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

- [cluster_wide_app_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster/properties/cluster_wide_app_list/#section)
- [no_cluster_wide_apps](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster/properties/no_cluster_wide_apps/#section)

Select alternatives according to the provider validators above.

## Direct properties

- [cluster_wide_apps](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster/properties/cluster_wide_app_list/cluster_wide_apps/): complete subsection reference.

## Next pages

- [cluster_wide_app_list.cluster_wide_apps](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster/properties/cluster_wide_app_list/cluster_wide_apps/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster/properties/)
- [xcsh_k8s_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster/)
