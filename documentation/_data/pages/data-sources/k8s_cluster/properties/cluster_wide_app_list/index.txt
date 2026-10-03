---
page_title: "cluster_wide_app_list"
subcategory: ""
description: "List of cluster wide applications."
xcsh_docs: {"aliases": ["cluster wide app list"], "body_bytes": 1860, "body_sha256": "sha256:47c1b0a377c5516b9166874fbc4167176d618979ddb67929d4f2bdf87b7682f4", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:k8s_cluster:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:k8s_cluster:properties:cluster_wide_app_list", "parent_id": "xcsh-docs:data-sources:k8s_cluster:reference", "path": "documentation/data-sources/k8s_cluster/properties/cluster_wide_app_list/index.md", "product": "distributed-cloud", "provider_name": "k8s_cluster", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3012002212023212-0312020122110302-1031000322301012-1133301000011230-1113222032032232-2113012000333012-3031101320033120-2033112013002132", "registry_path": "docs/guides/data-sources--k8s_cluster--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["cluster_wide_app_list"], "schema_version": 1, "sections": [{"aliases": ["cluster wide app list cluster wide apps"], "anchor": "section", "description": "List of cluster wide applications.", "document_id": "xcsh-docs:data-sources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["cluster_wide_app_list", "cluster_wide_apps"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/k8s_cluster/properties/cluster_wide_app_list/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "List of cluster wide applications.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["k8s_clusterCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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
