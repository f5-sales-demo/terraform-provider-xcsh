---
page_title: "eks_k8s.enable_anti_affinity"
subcategory: ""
description: "Configuration for pod anti-affinity scheduling rules. Define multiple rules to control how different applications/components are distributed across your Kubernetes cluster."
xcsh_docs: {"aliases": ["eks k8s enable anti affinity"], "body_bytes": 1140, "body_sha256": "sha256:27b4a2828c69388482f31c5798a8b5bdc303f2caaa7d71d60caa8fd40b75d2f1", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:securemesh_site_v2:properties:eks_k8s:enable_anti_affinity:rules"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:eks_k8s:enable_anti_affinity", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:eks_k8s", "path": "documentation/data-sources/securemesh_site_v2/properties/eks_k8s/enable_anti_affinity/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "data-sources", "registry_anchor": "canonical-3132013131321030-0230231111323330-0110133210202030-1033301212132032-0031211122101322-2000233033010320-0233132211021221-0322023230121200", "registry_path": "docs/guides/data-sources--securemesh_site_v2--reference--group-007.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["eks_k8s", "enable_anti_affinity"], "schema_version": 1, "sections": [{"aliases": ["eks k8s enable anti affinity rules"], "anchor": "section", "description": "Define one or more anti-affinity rules. Each rule specifies which pods (by labels) should be distributed across which topology domains. Example: Rule 1 - Distribute VPM pods across nodes, Rule 2 - Distribute Prometheus pods across zones.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:eks_k8s:enable_anti_affinity:rules", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["eks_k8s", "enable_anti_affinity", "rules"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/eks_k8s/enable_anti_affinity/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "Configuration for pod anti-affinity scheduling rules. Define multiple rules to control how different applications/components are distributed across your Kubernetes cluster.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# eks_k8s.enable_anti_affinity

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/)
- [eks_k8s](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/eks_k8s/)
- eks_k8s.enable_anti_affinity

<a id="section"></a>

Type: `"single"`. Computed.

Configuration for pod anti-affinity scheduling rules. Define multiple rules to control how different
applications/components are distributed across your Kubernetes cluster.

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

## Direct properties

- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/eks_k8s/enable_anti_affinity/rules/): complete subsection reference.
