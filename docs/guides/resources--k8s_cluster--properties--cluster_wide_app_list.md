---
page_title: "cluster_wide_app_list"
subcategory: ""
description: "cluster_wide_app_list for xcsh_k8s_cluster."
xcsh_docs: {"aliases": [], "body_bytes": 1712, "body_sha256": "sha256:f9679efc32f034d0288576a5c152164ff9e3b0774c2cced0607801d53cbdc7bc", "canonical_id": "xcsh-docs:resources:k8s_cluster:properties:cluster_wide_app_list", "child_ids": ["xcsh-docs:resources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps"], "collection_id": "xcsh-docs:resources:k8s_cluster:collection", "completeness": "complete", "id": "xcsh-docs:resources:k8s_cluster:properties:cluster_wide_app_list", "parent_id": "xcsh-docs:resources:k8s_cluster:reference", "path": "docs/guides/resources--k8s_cluster--properties--cluster_wide_app_list.md", "provider_name": "k8s_cluster", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["cluster_wide_app_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/k8s_cluster/properties/cluster_wide_app_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cluster_wide_app_list for xcsh_k8s_cluster.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["k8s_clusterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cluster_wide_app_list

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md)
- [Property reference](resources--k8s_cluster--reference.md)
- cluster_wide_app_list

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: cluster\_wide\_app\_list, no\_cluster\_wide\_apps; Default: no\_cluster\_wide\_apps\]
Cluster Wide Application List. List of cluster wide applications.

Upstream description:

List of cluster wide applications.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cluster_wide_apps")}
```

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

- [cluster_wide_app_list](resources--k8s_cluster--properties--cluster_wide_app_list.md#section)
- [no_cluster_wide_apps](resources--k8s_cluster--properties--no_cluster_wide_apps.md#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
cluster_wide_app_list {
  # Configure direct properties listed below.
}
```

## Direct properties

- [cluster_wide_apps](resources--k8s_cluster--properties--cluster_wide_app_list--cluster_wide_apps.md): complete subsection reference.

## Next pages

- [cluster_wide_app_list.cluster_wide_apps](resources--k8s_cluster--properties--cluster_wide_app_list--cluster_wide_apps.md)
- [Property reference](resources--k8s_cluster--reference.md)
- [xcsh_k8s_cluster](../resources/k8s_cluster.md)
