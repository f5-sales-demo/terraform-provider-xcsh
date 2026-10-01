---
page_title: "cluster_wide_app_list.cluster_wide_apps"
subcategory: ""
description: "cluster_wide_app_list.cluster_wide_apps for xcsh_k8s_cluster."
xcsh_docs: {"aliases": [], "body_bytes": 2806, "body_sha256": "sha256:a40d31658dc2ca16c94a5a557a47338ab1cc100a7cb1d512124347dbf51e7b46", "canonical_id": "xcsh-docs:data-sources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps", "child_ids": ["xcsh-docs:data-sources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:argo_cd", "xcsh-docs:data-sources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:dashboard", "xcsh-docs:data-sources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:metrics_server", "xcsh-docs:data-sources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:prometheus"], "collection_id": "xcsh-docs:data-sources:k8s_cluster:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps", "parent_id": "xcsh-docs:data-sources:k8s_cluster:properties:cluster_wide_app_list", "path": "docs/guides/data-sources--k8s_cluster--properties--cluster_wide_app_list--cluster_wide_apps.md", "provider_name": "k8s_cluster", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["cluster_wide_app_list", "cluster_wide_apps"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/k8s_cluster/properties/cluster_wide_app_list/cluster_wide_apps/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cluster_wide_app_list.cluster_wide_apps for xcsh_k8s_cluster.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["k8s_clusterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cluster_wide_app_list.cluster_wide_apps

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md)
- [Property reference](data-sources--k8s_cluster--reference.md)
- [cluster_wide_app_list](data-sources--k8s_cluster--properties--cluster_wide_app_list.md)
- cluster_wide_app_list.cluster_wide_apps

<a id="section"></a>

Type: `"list"`. Computed.

Cluster Wide Application List. List of cluster wide applications.

Upstream description:

List of cluster wide applications.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 5,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 5,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

## Direct properties

- [argo_cd](data-sources--k8s_cluster--properties--cluster_wide_app_list--cluster_wide_apps--argo_cd.md): complete subsection reference.

- [dashboard](data-sources--k8s_cluster--properties--cluster_wide_app_list--cluster_wide_apps--dashboard.md): complete subsection reference.

- [metrics_server](data-sources--k8s_cluster--properties--cluster_wide_app_list--cluster_wide_apps--metrics_server.md): complete subsection reference.

- [prometheus](data-sources--k8s_cluster--properties--cluster_wide_app_list--cluster_wide_apps--prometheus.md): complete subsection reference.

## Next pages

- [cluster_wide_app_list.cluster_wide_apps.argo_cd](data-sources--k8s_cluster--properties--cluster_wide_app_list--cluster_wide_apps--argo_cd.md)
- [cluster_wide_app_list.cluster_wide_apps.dashboard](data-sources--k8s_cluster--properties--cluster_wide_app_list--cluster_wide_apps--dashboard.md)
- [cluster_wide_app_list.cluster_wide_apps.metrics_server](data-sources--k8s_cluster--properties--cluster_wide_app_list--cluster_wide_apps--metrics_server.md)
- [cluster_wide_app_list.cluster_wide_apps.prometheus](data-sources--k8s_cluster--properties--cluster_wide_app_list--cluster_wide_apps--prometheus.md)
- [cluster_wide_app_list](data-sources--k8s_cluster--properties--cluster_wide_app_list.md)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md)
