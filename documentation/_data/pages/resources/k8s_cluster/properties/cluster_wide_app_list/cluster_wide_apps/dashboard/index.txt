---
page_title: "cluster_wide_app_list.cluster_wide_apps.dashboard"
subcategory: ""
description: "Description Parameters for K8s dashboard."
xcsh_docs: {"aliases": ["cluster wide app list cluster wide apps dashboard"], "body_bytes": 1146, "body_sha256": "sha256:c694d30e0ccd7376a356085147b7c51b5605586ed6580fd35d25b2fc5e50819e", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:k8s_cluster:collection", "completeness": "complete", "id": "xcsh-docs:resources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:dashboard", "parent_id": "xcsh-docs:resources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps", "path": "documentation/resources/k8s_cluster/properties/cluster_wide_app_list/cluster_wide_apps/dashboard/index.md", "product": "distributed-cloud", "provider_name": "k8s_cluster", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-3111130013310203-2021310113320211-1112100020113323-3123101301130332-1301211221103223-0123110011020110-0002210120232023-3010031222033000", "registry_path": "docs/guides/resources--k8s_cluster--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["cluster_wide_app_list", "cluster_wide_apps", "dashboard"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/k8s_cluster/properties/cluster_wide_app_list/cluster_wide_apps/dashboard/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "Description Parameters for K8s dashboard.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["k8s_clusterCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cluster_wide_app_list.cluster_wide_apps.dashboard

Breadcrumbs:

- [xcsh_k8s_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/properties/)
- [cluster_wide_app_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/properties/cluster_wide_app_list/)
- [cluster_wide_app_list.cluster_wide_apps](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/properties/cluster_wide_app_list/cluster_wide_apps/)
- cluster_wide_app_list.cluster_wide_apps.dashboard

<a id="section"></a>

Type: `["object", {}]`. Optional.

Description Parameters for K8s dashboard.

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
dashboard = {}
```

This is an empty object or choice marker. It has no direct properties.
