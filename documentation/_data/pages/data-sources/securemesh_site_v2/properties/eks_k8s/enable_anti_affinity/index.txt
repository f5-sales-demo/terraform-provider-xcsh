---
page_title: "eks_k8s.enable_anti_affinity"
subcategory: ""
description: "Configuration for pod anti-affinity scheduling rules. Define multiple rules to control how different applications/components are distributed across your Kubernetes cluster."
xcsh_docs: {"aliases": ["eks k8s enable anti affinity"], "body_bytes": 1767, "body_sha256": "sha256:b3b8818639404ba6688f9459f6d079f8d03767fbcc0caaebebcab5f4d3d53552", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:securemesh_site_v2:properties:eks_k8s:enable_anti_affinity:rules"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:eks_k8s:enable_anti_affinity", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:eks_k8s", "path": "documentation/data-sources/securemesh_site_v2/properties/eks_k8s/enable_anti_affinity/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3132013131321030-0230231111323330-0110133210202030-1033301212132032-0031211122101322-2000233033010320-0233132211021221-0322023230121200", "registry_path": "docs/guides/data-sources--securemesh_site_v2--reference--group-007.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["eks_k8s", "enable_anti_affinity"], "schema_version": 1, "sections": [{"aliases": ["rules"], "anchor": "section", "description": "Define one or more anti-affinity rules. Each rule specifies which pods (by labels) should be distributed across which topology domains. Example: Rule 1 - Distribute VPM pods across nodes, Rule 2 - Distribute Prometheus pods across zones.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:eks_k8s:enable_anti_affinity:rules", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["eks_k8s", "enable_anti_affinity", "rules"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/eks_k8s/enable_anti_affinity/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Configuration for pod anti-affinity scheduling rules. Define multiple rules to control how different applications/components are distributed across your Kubernetes cluster.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

Upstream description:

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

## Next pages

- [eks_k8s.enable_anti_affinity.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/eks_k8s/enable_anti_affinity/rules/)
- [eks_k8s](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/eks_k8s/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
