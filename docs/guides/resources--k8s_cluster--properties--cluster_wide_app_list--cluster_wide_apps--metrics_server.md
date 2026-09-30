---
page_title: "cluster_wide_app_list.cluster_wide_apps.metrics_server"
subcategory: ""
description: "cluster_wide_app_list.cluster_wide_apps.metrics_server for xcsh_k8s_cluster."
xcsh_docs: {"aliases": [], "body_bytes": 1184, "body_sha256": "sha256:ce3ff9fc67a96154c402bba55b8c468772d52ba3d6e2e1a73201b071c34bc231", "canonical_id": "xcsh-docs:resources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:metrics_server", "child_ids": [], "collection_id": "xcsh-docs:resources:k8s_cluster:collection", "completeness": "complete", "id": "xcsh-docs:resources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:metrics_server", "parent_id": "xcsh-docs:resources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps", "path": "docs/guides/resources--k8s_cluster--properties--cluster_wide_app_list--cluster_wide_apps--metrics_server.md", "provider_name": "k8s_cluster", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["cluster_wide_app_list", "cluster_wide_apps", "metrics_server"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/k8s_cluster/properties/cluster_wide_app_list/cluster_wide_apps/metrics_server/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cluster_wide_app_list.cluster_wide_apps.metrics_server for xcsh_k8s_cluster.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["k8s_clusterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# cluster_wide_app_list.cluster_wide_apps.metrics_server

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md)
- [Property reference](resources--k8s_cluster--reference.md)
- [cluster_wide_app_list](resources--k8s_cluster--properties--cluster_wide_app_list.md)
- [cluster_wide_app_list.cluster_wide_apps](resources--k8s_cluster--properties--cluster_wide_app_list--cluster_wide_apps.md)
- cluster_wide_app_list.cluster_wide_apps.metrics_server

<a id="section"></a>

Type: `["object", {}]`. Optional.

Description Parameters for Kubernetes Metrics Server application.

Upstream description:

Description Parameters for Kubernetes Metrics Server application.

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
metrics_server = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [cluster_wide_app_list.cluster_wide_apps](resources--k8s_cluster--properties--cluster_wide_app_list--cluster_wide_apps.md)
- [xcsh_k8s_cluster](../resources/k8s_cluster.md)
