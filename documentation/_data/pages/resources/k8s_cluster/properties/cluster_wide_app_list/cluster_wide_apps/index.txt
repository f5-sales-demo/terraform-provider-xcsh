---
page_title: "cluster_wide_app_list.cluster_wide_apps"
subcategory: ""
description: "cluster_wide_app_list.cluster_wide_apps for xcsh_k8s_cluster."
xcsh_docs: {"aliases": [], "body_bytes": 4104, "body_sha256": "sha256:acb7a4c75b745898b7fe5b40e6154a039ebc606d83b8baaea63d4efc423e9221", "child_ids": ["xcsh-docs:resources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:argo_cd", "xcsh-docs:resources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:dashboard", "xcsh-docs:resources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:metrics_server", "xcsh-docs:resources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:prometheus"], "collection_id": "xcsh-docs:resources:k8s_cluster:collection", "completeness": "complete", "id": "xcsh-docs:resources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps", "parent_id": "xcsh-docs:resources:k8s_cluster:properties:cluster_wide_app_list", "path": "documentation/resources/k8s_cluster/properties/cluster_wide_app_list/cluster_wide_apps/index.md", "provider_name": "k8s_cluster", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["cluster_wide_app_list", "cluster_wide_apps"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/k8s_cluster/properties/cluster_wide_app_list/cluster_wide_apps/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cluster_wide_app_list.cluster_wide_apps for xcsh_k8s_cluster.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["k8s_clusterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cluster_wide_app_list.cluster_wide_apps

Breadcrumbs:

- [xcsh_k8s_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/properties/)
- [cluster_wide_app_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/properties/cluster_wide_app_list/)
- cluster_wide_app_list.cluster_wide_apps

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Cluster Wide Application List. List of cluster wide applications.

Upstream description:

List of cluster wide applications.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("argo_cd",
    "dashboard"),
  validators.ConflictingListObjectAttributes("argo_cd",
    "metrics_server"),
  validators.ConflictingListObjectAttributes("argo_cd",
    "prometheus"),
  validators.ConflictingListObjectAttributes("dashboard",
    "metrics_server"),
  validators.ConflictingListObjectAttributes("dashboard",
    "prometheus"),
  validators.ConflictingListObjectAttributes("metrics_server",
    "prometheus")}
```

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

Terraform syntax:

```terraform
cluster_wide_apps {
  # Configure direct properties listed below.
}
```

## Direct properties

- [argo_cd](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/properties/cluster_wide_app_list/cluster_wide_apps/argo_cd/): complete subsection reference.

- [dashboard](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/properties/cluster_wide_app_list/cluster_wide_apps/dashboard/): complete subsection reference.

- [metrics_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/properties/cluster_wide_app_list/cluster_wide_apps/metrics_server/): complete subsection reference.

- [prometheus](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/properties/cluster_wide_app_list/cluster_wide_apps/prometheus/): complete subsection reference.

## Next pages

- [cluster_wide_app_list.cluster_wide_apps.argo_cd](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/properties/cluster_wide_app_list/cluster_wide_apps/argo_cd/)
- [cluster_wide_app_list.cluster_wide_apps.dashboard](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/properties/cluster_wide_app_list/cluster_wide_apps/dashboard/)
- [cluster_wide_app_list.cluster_wide_apps.metrics_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/properties/cluster_wide_app_list/cluster_wide_apps/metrics_server/)
- [cluster_wide_app_list.cluster_wide_apps.prometheus](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/properties/cluster_wide_app_list/cluster_wide_apps/prometheus/)
- [cluster_wide_app_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/properties/cluster_wide_app_list/)
- [xcsh_k8s_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/)
