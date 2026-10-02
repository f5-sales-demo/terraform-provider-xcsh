---
page_title: "cluster_wide_app_list.cluster_wide_apps"
subcategory: ""
description: "List of cluster wide applications."
xcsh_docs: {"aliases": ["cluster wide app list cluster wide apps"], "body_bytes": 3455, "body_sha256": "sha256:27cd83f7be11058ce79c27d60f0d0d767d0291384e5be1bad44362ddf7b96e69", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:argo_cd", "xcsh-docs:data-sources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:dashboard", "xcsh-docs:data-sources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:metrics_server", "xcsh-docs:data-sources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:prometheus"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:k8s_cluster:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps", "parent_id": "xcsh-docs:data-sources:k8s_cluster:properties:cluster_wide_app_list", "path": "documentation/data-sources/k8s_cluster/properties/cluster_wide_app_list/cluster_wide_apps/index.md", "product": "distributed-cloud", "provider_name": "k8s_cluster", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1311132120113231-1220003233130312-3200110311213311-1200212223120103-2102323332111012-3311211121231222-2133100111131320-1333302222203222", "registry_path": "docs/guides/data-sources--k8s_cluster--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["cluster_wide_app_list", "cluster_wide_apps"], "schema_version": 1, "sections": [{"aliases": ["argo cd"], "anchor": "section", "description": "Description Parameters for Argo Continuous Deployment(CD) application.", "document_id": "xcsh-docs:data-sources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:argo_cd", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["cluster_wide_app_list", "cluster_wide_apps", "argo_cd"], "syntax": "attribute", "type": "object"}, {"aliases": ["dashboard"], "anchor": "section", "description": "Description Parameters for K8s dashboard.", "document_id": "xcsh-docs:data-sources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:dashboard", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cluster_wide_app_list", "cluster_wide_apps", "dashboard"], "syntax": "attribute", "type": "object"}, {"aliases": ["metrics server"], "anchor": "section", "description": "Description Parameters for Kubernetes Metrics Server application.", "document_id": "xcsh-docs:data-sources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:metrics_server", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cluster_wide_app_list", "cluster_wide_apps", "metrics_server"], "syntax": "attribute", "type": "object"}, {"aliases": ["prometheus"], "anchor": "section", "description": "Description Parameters for Prometheus server access.", "document_id": "xcsh-docs:data-sources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:prometheus", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cluster_wide_app_list", "cluster_wide_apps", "prometheus"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/k8s_cluster/properties/cluster_wide_app_list/cluster_wide_apps/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "List of cluster wide applications.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["k8s_clusterCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

- [argo_cd](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster/properties/cluster_wide_app_list/cluster_wide_apps/argo_cd/): complete subsection reference.

- [dashboard](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster/properties/cluster_wide_app_list/cluster_wide_apps/dashboard/): complete subsection reference.

- [metrics_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster/properties/cluster_wide_app_list/cluster_wide_apps/metrics_server/): complete subsection reference.

- [prometheus](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster/properties/cluster_wide_app_list/cluster_wide_apps/prometheus/): complete subsection reference.

## Next pages

- [cluster_wide_app_list.cluster_wide_apps.argo_cd](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster/properties/cluster_wide_app_list/cluster_wide_apps/argo_cd/)
- [cluster_wide_app_list.cluster_wide_apps.dashboard](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster/properties/cluster_wide_app_list/cluster_wide_apps/dashboard/)
- [cluster_wide_app_list.cluster_wide_apps.metrics_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster/properties/cluster_wide_app_list/cluster_wide_apps/metrics_server/)
- [cluster_wide_app_list.cluster_wide_apps.prometheus](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster/properties/cluster_wide_app_list/cluster_wide_apps/prometheus/)
- [cluster_wide_app_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster/properties/cluster_wide_app_list/)
- [xcsh_k8s_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster/)
