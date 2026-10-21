---
page_title: "cluster_wide_app_list.cluster_wide_apps.prometheus"
subcategory: ""
description: "Description Parameters for Prometheus server access."
xcsh_docs: {"aliases": ["cluster wide app list cluster wide apps prometheus"], "body_bytes": 1160, "body_sha256": "sha256:00e440d74c69649214b58704b316a103eb8e3b094abeba7fff24b2301c6b275b", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:k8s_cluster:collection", "completeness": "complete", "id": "xcsh-docs:resources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:prometheus", "parent_id": "xcsh-docs:resources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps", "path": "documentation/resources/k8s_cluster/properties/cluster_wide_app_list/cluster_wide_apps/prometheus/index.md", "product": "distributed-cloud", "provider_name": "k8s_cluster", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-2102233013113032-1120130222210211-2313102122223131-0212311203003303-2121032202130330-3000012030231030-2300301033202210-3023131302301013", "registry_path": "docs/guides/resources--k8s_cluster--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["cluster_wide_app_list", "cluster_wide_apps", "prometheus"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/k8s_cluster/properties/cluster_wide_app_list/cluster_wide_apps/prometheus/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "Description Parameters for Prometheus server access.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["k8s_clusterCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cluster_wide_app_list.cluster_wide_apps.prometheus

Breadcrumbs:

- [xcsh_k8s_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/properties/)
- [cluster_wide_app_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/properties/cluster_wide_app_list/)
- [cluster_wide_app_list.cluster_wide_apps](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/properties/cluster_wide_app_list/cluster_wide_apps/)
- cluster_wide_app_list.cluster_wide_apps.prometheus

<a id="section"></a>

Type: `["object", {}]`. Optional.

Description Parameters for Prometheus server access.

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
prometheus = {}
```

This is an empty object or choice marker. It has no direct properties.
