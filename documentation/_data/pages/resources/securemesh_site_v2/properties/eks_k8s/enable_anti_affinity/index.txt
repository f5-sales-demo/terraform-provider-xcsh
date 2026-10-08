---
page_title: "eks_k8s.enable_anti_affinity"
subcategory: ""
description: "Configuration for pod anti-affinity scheduling rules. Define multiple rules to control how different applications/components are distributed across your Kubernetes cluster."
xcsh_docs: {"aliases": ["eks k8s enable anti affinity"], "body_bytes": 1431, "body_sha256": "sha256:398e20f468ecd4afe33b1259937657b9e68385b79c307ea4de62e32c86081947", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:eks_k8s:enable_anti_affinity:rules"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:eks_k8s:enable_anti_affinity", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:eks_k8s", "path": "documentation/resources/securemesh_site_v2/properties/eks_k8s/enable_anti_affinity/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-1313030010022021-0221332100201013-0020030200213230-0021013230101031-1131032330001021-0111013012301323-1310322112213013-0102113300201010", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-006.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "eks_k8s.enable_anti_affinity:RequiredObjectAttributes:rules", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:eks_k8s:enable_anti_affinity:rules", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["eks_k8s", "enable_anti_affinity"], "schema_version": 1, "sections": [{"aliases": ["eks k8s enable anti affinity rules"], "anchor": "section", "description": "Define one or more anti-affinity rules. Each rule specifies which pods (by labels) should be distributed across which topology domains. Example: Rule 1 - Distribute VPM pods across nodes, Rule 2 - Distribute Prometheus pods across zones.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:eks_k8s:enable_anti_affinity:rules", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-eks_k8s--enable_anti_affinity--rules--label_key", "enforcement": "provider-schema", "group": "eks_k8s.enable_anti_affinity.rules:RequiredListObjectAttributes:label_key,label_value,topology_keys", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:eks_k8s:enable_anti_affinity:rules", "type": "requires"}, {"anchor": "schema-eks_k8s--enable_anti_affinity--rules--label_value", "enforcement": "provider-schema", "group": "eks_k8s.enable_anti_affinity.rules:RequiredListObjectAttributes:label_key,label_value,topology_keys", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:eks_k8s:enable_anti_affinity:rules", "type": "requires"}, {"anchor": "schema-eks_k8s--enable_anti_affinity--rules--topology_keys", "enforcement": "provider-schema", "group": "eks_k8s.enable_anti_affinity.rules:RequiredListObjectAttributes:label_key,label_value,topology_keys", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:eks_k8s:enable_anti_affinity:rules", "type": "requires"}], "schema_path": ["eks_k8s", "enable_anti_affinity", "rules"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/eks_k8s/enable_anti_affinity/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Configuration for pod anti-affinity scheduling rules. Define multiple rules to control how different applications/components are distributed across your Kubernetes cluster.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
