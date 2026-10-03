---
page_title: "cluster_wide_app_list.cluster_wide_apps.prometheus"
subcategory: ""
description: "Description Parameters for Prometheus server access."
xcsh_docs: {"aliases": ["cluster wide app list cluster wide apps prometheus"], "body_bytes": 1551, "body_sha256": "sha256:351ca9d8549707ee3a053f115784a4b3981c366d686ed4b006e1cea5fcdbdd28", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:k8s_cluster:collection", "completeness": "complete", "id": "xcsh-docs:resources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:prometheus", "parent_id": "xcsh-docs:resources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps", "path": "documentation/resources/k8s_cluster/properties/cluster_wide_app_list/cluster_wide_apps/prometheus/index.md", "product": "distributed-cloud", "provider_name": "k8s_cluster", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2102233013113032-1120130222210211-2313102122223131-0212311203003303-2121032202130330-3000012030231030-2300301033202210-3023131302301013", "registry_path": "docs/guides/resources--k8s_cluster--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["cluster_wide_app_list", "cluster_wide_apps", "prometheus"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/k8s_cluster/properties/cluster_wide_app_list/cluster_wide_apps/prometheus/index.txt", "spec_pin_digest": "sha256:62f71ec22260bc99f65753ef4581eb9e0dec1b65c506bb0d53099db73a05e19e", "summary": "Description Parameters for Prometheus server access.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.2", "schema_components": ["k8s_clusterCreateRequest"], "target_commit": "1a0b5141f4589ffaf7bb696a4369a16ee74ae2ff"}}
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

Upstream description:

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

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [cluster_wide_app_list.cluster_wide_apps](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/properties/cluster_wide_app_list/cluster_wide_apps/)
- [xcsh_k8s_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/)
