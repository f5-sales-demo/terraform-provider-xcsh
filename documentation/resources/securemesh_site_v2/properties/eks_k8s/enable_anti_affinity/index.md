---
page_title: "eks_k8s.enable_anti_affinity"
subcategory: ""
description: "Configuration for pod anti-affinity scheduling rules. Define multiple rules to control how different applications/components are distributed across your Kubernetes cluster."
xcsh_docs: {"aliases": ["eks k8s enable anti affinity"], "body_bytes": 2019, "body_sha256": "sha256:e5d9dd85401a7a1683c9ee16af804fb7cf1058fd80ac718079896f3a7eee58a4", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:eks_k8s:enable_anti_affinity:rules"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:eks_k8s:enable_anti_affinity", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:eks_k8s", "path": "documentation/resources/securemesh_site_v2/properties/eks_k8s/enable_anti_affinity/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1313030010022021-0221332100201013-0020030200213230-0021013230101031-1131032330001021-0111013012301323-1310322112213013-0102113300201010", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-007.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "eks_k8s.enable_anti_affinity:RequiredObjectAttributes:rules", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:eks_k8s:enable_anti_affinity:rules", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["eks_k8s", "enable_anti_affinity"], "schema_version": 1, "sections": [{"aliases": ["rules"], "anchor": "section", "description": "Define one or more anti-affinity rules. Each rule specifies which pods (by labels) should be distributed across which topology domains. Example: Rule 1 - Distribute VPM pods across nodes, Rule 2 - Distribute Prometheus pods across zones.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:eks_k8s:enable_anti_affinity:rules", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-eks_k8s--enable_anti_affinity--rules--label_key", "enforcement": "provider-schema", "group": "eks_k8s.enable_anti_affinity.rules:RequiredListObjectAttributes:label_key,label_value,topology_keys", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:eks_k8s:enable_anti_affinity:rules", "type": "requires"}, {"anchor": "schema-eks_k8s--enable_anti_affinity--rules--label_value", "enforcement": "provider-schema", "group": "eks_k8s.enable_anti_affinity.rules:RequiredListObjectAttributes:label_key,label_value,topology_keys", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:eks_k8s:enable_anti_affinity:rules", "type": "requires"}, {"anchor": "schema-eks_k8s--enable_anti_affinity--rules--topology_keys", "enforcement": "provider-schema", "group": "eks_k8s.enable_anti_affinity.rules:RequiredListObjectAttributes:label_key,label_value,topology_keys", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:eks_k8s:enable_anti_affinity:rules", "type": "requires"}], "schema_path": ["eks_k8s", "enable_anti_affinity", "rules"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/eks_k8s/enable_anti_affinity/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Configuration for pod anti-affinity scheduling rules. Define multiple rules to control how different applications/components are distributed across your Kubernetes cluster.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# eks_k8s.enable_anti_affinity

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- [eks_k8s](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/eks_k8s/)
- eks_k8s.enable_anti_affinity

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration for pod anti-affinity scheduling rules. Define multiple rules to control how different
applications/components are distributed across your Kubernetes cluster.

Upstream description:

Configuration for pod anti-affinity scheduling rules. Define multiple rules to control how different
applications/components are distributed across your Kubernetes cluster.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("rules")}
```

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
enable_anti_affinity {
  # Configure direct properties listed below.
}
```

## Direct properties

- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/eks_k8s/enable_anti_affinity/rules/): complete subsection reference.

## Next pages

- [eks_k8s.enable_anti_affinity.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/eks_k8s/enable_anti_affinity/rules/)
- [eks_k8s](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/eks_k8s/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
