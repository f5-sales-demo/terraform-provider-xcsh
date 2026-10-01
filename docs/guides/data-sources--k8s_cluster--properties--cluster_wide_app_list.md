---
page_title: "cluster_wide_app_list"
subcategory: ""
description: "cluster_wide_app_list for xcsh_k8s_cluster."
xcsh_docs: {"aliases": [], "body_bytes": 1450, "body_sha256": "sha256:754856879cf0f1178754b69c05e029279ce50a1de3117f76767e5811cd5c2f0a", "canonical_id": "xcsh-docs:data-sources:k8s_cluster:properties:cluster_wide_app_list", "child_ids": ["xcsh-docs:data-sources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps"], "collection_id": "xcsh-docs:data-sources:k8s_cluster:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:k8s_cluster:properties:cluster_wide_app_list", "parent_id": "xcsh-docs:data-sources:k8s_cluster:reference", "path": "docs/guides/data-sources--k8s_cluster--properties--cluster_wide_app_list.md", "provider_name": "k8s_cluster", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["cluster_wide_app_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/k8s_cluster/properties/cluster_wide_app_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cluster_wide_app_list for xcsh_k8s_cluster.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["k8s_clusterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cluster_wide_app_list

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md)
- [Property reference](data-sources--k8s_cluster--reference.md)
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

- [cluster_wide_app_list](data-sources--k8s_cluster--properties--cluster_wide_app_list.md#section)
- [no_cluster_wide_apps](data-sources--k8s_cluster--properties--no_cluster_wide_apps.md#section)

Select alternatives according to the provider validators above.

## Direct properties

- [cluster_wide_apps](data-sources--k8s_cluster--properties--cluster_wide_app_list--cluster_wide_apps.md): complete subsection reference.

## Next pages

- [cluster_wide_app_list.cluster_wide_apps](data-sources--k8s_cluster--properties--cluster_wide_app_list--cluster_wide_apps.md)
- [Property reference](data-sources--k8s_cluster--reference.md)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md)
