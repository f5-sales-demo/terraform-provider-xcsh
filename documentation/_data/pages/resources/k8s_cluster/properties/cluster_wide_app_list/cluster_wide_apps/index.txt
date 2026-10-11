---
page_title: "cluster_wide_app_list.cluster_wide_apps"
subcategory: ""
description: "List of cluster wide applications."
xcsh_docs: {"aliases": ["cluster wide app list cluster wide apps"], "body_bytes": 3031, "body_sha256": "sha256:3d8be38697d449dd65836e5c0211d81507d29d1cfd572c9fb8d13bbaa7b5ab1c", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:argo_cd", "xcsh-docs:resources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:dashboard", "xcsh-docs:resources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:metrics_server", "xcsh-docs:resources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:prometheus"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:k8s_cluster:collection", "completeness": "complete", "id": "xcsh-docs:resources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps", "parent_id": "xcsh-docs:resources:k8s_cluster:properties:cluster_wide_app_list", "path": "documentation/resources/k8s_cluster/properties/cluster_wide_app_list/cluster_wide_apps/index.md", "product": "distributed-cloud", "provider_name": "k8s_cluster", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "resources", "registry_anchor": "canonical-0200001310311212-1121331310012002-3033011130202120-1023002011031303-0200330313203121-0103121103310212-0321232101200102-0333031333312020", "registry_path": "docs/guides/resources--k8s_cluster--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "cluster_wide_app_list.cluster_wide_apps:ConflictingListObjectAttributes:argo_cd,dashboard", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:argo_cd", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cluster_wide_app_list.cluster_wide_apps:ConflictingListObjectAttributes:argo_cd,metrics_server", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:argo_cd", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cluster_wide_app_list.cluster_wide_apps:ConflictingListObjectAttributes:argo_cd,prometheus", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:argo_cd", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cluster_wide_app_list.cluster_wide_apps:ConflictingListObjectAttributes:argo_cd,dashboard", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:dashboard", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cluster_wide_app_list.cluster_wide_apps:ConflictingListObjectAttributes:dashboard,metrics_server", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:dashboard", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cluster_wide_app_list.cluster_wide_apps:ConflictingListObjectAttributes:dashboard,prometheus", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:dashboard", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cluster_wide_app_list.cluster_wide_apps:ConflictingListObjectAttributes:argo_cd,metrics_server", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:metrics_server", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cluster_wide_app_list.cluster_wide_apps:ConflictingListObjectAttributes:dashboard,metrics_server", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:metrics_server", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cluster_wide_app_list.cluster_wide_apps:ConflictingListObjectAttributes:metrics_server,prometheus", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:metrics_server", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cluster_wide_app_list.cluster_wide_apps:ConflictingListObjectAttributes:argo_cd,prometheus", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:prometheus", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cluster_wide_app_list.cluster_wide_apps:ConflictingListObjectAttributes:dashboard,prometheus", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:prometheus", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cluster_wide_app_list.cluster_wide_apps:ConflictingListObjectAttributes:metrics_server,prometheus", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:prometheus", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["cluster_wide_app_list", "cluster_wide_apps"], "schema_version": 1, "sections": [{"aliases": ["cluster wide app list cluster wide apps argo cd"], "anchor": "section", "description": "Description Parameters for Argo Continuous Deployment(CD) application.", "document_id": "xcsh-docs:resources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:argo_cd", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["cluster_wide_app_list", "cluster_wide_apps", "argo_cd"], "syntax": "block", "type": "object"}, {"aliases": ["cluster wide app list cluster wide apps dashboard"], "anchor": "section", "description": "Description Parameters for K8s dashboard.", "document_id": "xcsh-docs:resources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:dashboard", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cluster_wide_app_list", "cluster_wide_apps", "dashboard"], "syntax": "attribute", "type": "object"}, {"aliases": ["cluster wide app list cluster wide apps metrics server"], "anchor": "section", "description": "Description Parameters for Kubernetes Metrics Server application.", "document_id": "xcsh-docs:resources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:metrics_server", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cluster_wide_app_list", "cluster_wide_apps", "metrics_server"], "syntax": "attribute", "type": "object"}, {"aliases": ["cluster wide app list cluster wide apps prometheus"], "anchor": "section", "description": "Description Parameters for Prometheus server access.", "document_id": "xcsh-docs:resources:k8s_cluster:properties:cluster_wide_app_list:cluster_wide_apps:prometheus", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cluster_wide_app_list", "cluster_wide_apps", "prometheus"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/k8s_cluster/properties/cluster_wide_app_list/cluster_wide_apps/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "List of cluster wide applications.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["k8s_clusterCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
