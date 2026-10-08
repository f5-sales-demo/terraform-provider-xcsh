---
page_title: "cluster_wide_app_list.cluster_wide_apps"
subcategory: ""
description: "List of cluster wide applications."
xcsh_docs: {"aliases": ["cluster wide app list cluster wide apps"], "body_bytes": 2334, "body_sha256": "sha256:a8b68a61d83c1440eb4132dd737e4da65715a74a91917e4755f94b1341882ec6", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:argo_cd", "xcsh-docs:data-sources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:dashboard", "xcsh-docs:data-sources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:metrics_server", "xcsh-docs:data-sources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:prometheus"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:k8s_cluster:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps", "parent_id": "xcsh-docs:data-sources:k8s_cluster:properties:cluster_wide_app_list", "path": "documentation/data-sources/k8s_cluster/properties/cluster_wide_app_list/cluster_wide_apps/index.md", "product": "distributed-cloud", "provider_name": "k8s_cluster", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-1311132120113231-1220003233130312-3200110311213311-1200212223120103-2102323332111012-3311211121231222-2133100111131320-1333302222203222", "registry_path": "docs/guides/data-sources--k8s_cluster--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["cluster_wide_app_list", "cluster_wide_apps"], "schema_version": 1, "sections": [{"aliases": ["cluster wide app list cluster wide apps argo cd"], "anchor": "section", "description": "Description Parameters for Argo Continuous Deployment(CD) application.", "document_id": "xcsh-docs:data-sources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:argo_cd", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["cluster_wide_app_list", "cluster_wide_apps", "argo_cd"], "syntax": "attribute", "type": "object"}, {"aliases": ["cluster wide app list cluster wide apps dashboard"], "anchor": "section", "description": "Description Parameters for K8s dashboard.", "document_id": "xcsh-docs:data-sources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:dashboard", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cluster_wide_app_list", "cluster_wide_apps", "dashboard"], "syntax": "attribute", "type": "object"}, {"aliases": ["cluster wide app list cluster wide apps metrics server"], "anchor": "section", "description": "Description Parameters for Kubernetes Metrics Server application.", "document_id": "xcsh-docs:data-sources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:metrics_server", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cluster_wide_app_list", "cluster_wide_apps", "metrics_server"], "syntax": "attribute", "type": "object"}, {"aliases": ["cluster wide app list cluster wide apps prometheus"], "anchor": "section", "description": "Description Parameters for Prometheus server access.", "document_id": "xcsh-docs:data-sources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:prometheus", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cluster_wide_app_list", "cluster_wide_apps", "prometheus"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/k8s_cluster/properties/cluster_wide_app_list/cluster_wide_apps/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "List of cluster wide applications.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["k8s_clusterCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cluster_wide_app_list.cluster_wide_apps

Breadcrumbs:

- [xcsh_k8s_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster/properties/)
- [cluster_wide_app_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster/properties/cluster_wide_app_list/)
- cluster_wide_app_list.cluster_wide_apps

<a id="section"></a>

Type: `"list"`. Computed.

Cluster Wide Application List. List of cluster wide applications.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [argo_cd](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster/properties/cluster_wide_app_list/cluster_wide_apps/argo_cd/): complete subsection reference.

- [dashboard](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster/properties/cluster_wide_app_list/cluster_wide_apps/dashboard/): complete subsection reference.

- [metrics_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster/properties/cluster_wide_app_list/cluster_wide_apps/metrics_server/): complete subsection reference.

- [prometheus](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster/properties/cluster_wide_app_list/cluster_wide_apps/prometheus/): complete subsection reference.
