---
page_title: "cluster_wide_app_list.cluster_wide_apps.argo_cd"
subcategory: ""
description: "Description Parameters for Argo Continuous Deployment(CD) application."
xcsh_docs: {"aliases": ["cluster wide app list cluster wide apps argo cd"], "body_bytes": 1381, "body_sha256": "sha256:e1c1b95eb875e1a64e566a2acea44648c39fcd2d2179b5b5374115195d79b53e", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:argo_cd:local_domain"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:k8s_cluster:collection", "completeness": "complete", "id": "xcsh-docs:resources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:argo_cd", "parent_id": "xcsh-docs:resources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps", "path": "documentation/resources/k8s_cluster/properties/cluster_wide_app_list/cluster_wide_apps/argo_cd/index.md", "product": "distributed-cloud", "provider_name": "k8s_cluster", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-0332100103210130-3303021300221232-1002012211332320-2011210220030111-0300222103122203-3300233010303203-0233111303210230-3033131021312230", "registry_path": "docs/guides/resources--k8s_cluster--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["cluster_wide_app_list", "cluster_wide_apps", "argo_cd"], "schema_version": 1, "sections": [{"aliases": ["cluster wide app list cluster wide apps argo cd local domain"], "anchor": "section", "description": "Parameters required to enable local access.", "document_id": "xcsh-docs:resources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:argo_cd:local_domain", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-cluster_wide_app_list--cluster_wide_apps--argo_cd--local_domain--port", "enforcement": "provider-schema", "group": "cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain:ConflictingObjectAttributes:default_port,port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:argo_cd:local_domain", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain:ConflictingObjectAttributes:default_port,port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:argo_cd:local_domain:default_port", "type": "conflicts"}, {"anchor": "schema-cluster_wide_app_list--cluster_wide_apps--argo_cd--local_domain--local_domain", "enforcement": "provider-schema", "group": "cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain:RequiredObjectAttributes:local_domain", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:argo_cd:local_domain", "type": "requires"}], "schema_path": ["cluster_wide_app_list", "cluster_wide_apps", "argo_cd", "local_domain"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/k8s_cluster/properties/cluster_wide_app_list/cluster_wide_apps/argo_cd/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Description Parameters for Argo Continuous Deployment(CD) application.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["k8s_clusterCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cluster_wide_app_list.cluster_wide_apps.argo_cd

Breadcrumbs:

- [xcsh_k8s_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/properties/)
- [cluster_wide_app_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/properties/cluster_wide_app_list/)
- [cluster_wide_app_list.cluster_wide_apps](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/properties/cluster_wide_app_list/cluster_wide_apps/)
- cluster_wide_app_list.cluster_wide_apps.argo_cd

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Description Parameters for Argo Continuous Deployment(CD) application.

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
argo_cd {
  # Configure direct properties listed below.
}
```

## Direct properties

- [local_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/properties/cluster_wide_app_list/cluster_wide_apps/argo_cd/local_domain/): complete subsection reference.
