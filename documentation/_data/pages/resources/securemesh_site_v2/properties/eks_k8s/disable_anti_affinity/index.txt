---
page_title: "eks_k8s.disable_anti_affinity"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["eks k8s disable anti affinity"], "body_bytes": 1305, "body_sha256": "sha256:4b9d5893571ec7d0c99e64ab60da190b95b28f75afc938bfad5594dd0177562b", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:eks_k8s:disable_anti_affinity", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:eks_k8s", "path": "documentation/resources/securemesh_site_v2/properties/eks_k8s/disable_anti_affinity/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-2302123313002100-2233333033311201-1310100113000030-3230101310313120-2023123110023111-3310203313322012-2221021100323013-3032312333231020", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-007.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["eks_k8s", "disable_anti_affinity"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/eks_k8s/disable_anti_affinity/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# eks_k8s.disable_anti_affinity

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- [eks_k8s](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/eks_k8s/)
- eks_k8s.disable_anti_affinity

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable anti affinity.

Upstream description:

This can be used for messages where no values are needed.

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
disable_anti_affinity = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [eks_k8s](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/eks_k8s/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
