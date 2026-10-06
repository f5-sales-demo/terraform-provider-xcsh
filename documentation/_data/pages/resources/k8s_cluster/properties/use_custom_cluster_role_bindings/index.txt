---
page_title: "use_custom_cluster_role_bindings"
subcategory: ""
description: "List of active cluster role binding list for a K8s cluster."
xcsh_docs: {"aliases": ["use custom cluster role bindings"], "body_bytes": 1795, "body_sha256": "sha256:6cad01bbb110e76105ac8e0bea41dd57f627feaa31075763e5c61e7c58639586", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:k8s_cluster:properties:use_custom_cluster_role_bindings:cluster_role_bindings"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:k8s_cluster:collection", "completeness": "complete", "id": "xcsh-docs:resources:k8s_cluster:properties:use_custom_cluster_role_bindings", "parent_id": "xcsh-docs:resources:k8s_cluster:reference", "path": "documentation/resources/k8s_cluster/properties/use_custom_cluster_role_bindings/index.md", "product": "distributed-cloud", "provider_name": "k8s_cluster", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-0211302322000100-0203322011230121-0312133132101320-1031210001131331-1031230221102200-3221230333033330-3133111123333321-0133230211020211", "registry_path": "docs/guides/resources--k8s_cluster--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "use_custom_cluster_role_bindings:RequiredObjectAttributes:cluster_role_bindings", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:k8s_cluster:properties:use_custom_cluster_role_bindings:cluster_role_bindings", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["use_custom_cluster_role_bindings"], "schema_version": 1, "sections": [{"aliases": ["use custom cluster role bindings cluster role bindings"], "anchor": "section", "description": "List of active cluster role binding list for a K8s cluster.", "document_id": "xcsh-docs:resources:k8s_cluster:properties:use_custom_cluster_role_bindings:cluster_role_bindings", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-use_custom_cluster_role_bindings--cluster_role_bindings--name", "enforcement": "provider-schema", "group": "use_custom_cluster_role_bindings.cluster_role_bindings:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:k8s_cluster:properties:use_custom_cluster_role_bindings:cluster_role_bindings", "type": "requires"}], "schema_path": ["use_custom_cluster_role_bindings", "cluster_role_bindings"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/k8s_cluster/properties/use_custom_cluster_role_bindings/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "List of active cluster role binding list for a K8s cluster.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["k8s_clusterCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# use_custom_cluster_role_bindings

Breadcrumbs:

- [xcsh_k8s_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/properties/)
- use_custom_cluster_role_bindings

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: use\_custom\_cluster\_role\_bindings, use\_default\_cluster\_role\_bindings; Default:
use\_default\_cluster\_role\_bindings\] List of active cluster role binding list for a K8s cluster.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cluster_role_bindings")}
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

OneOf alternatives in this subsection:

- [use_custom_cluster_role_bindings](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/properties/use_custom_cluster_role_bindings/#section)
- [use_default_cluster_role_bindings](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/properties/use_default_cluster_role_bindings/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
use_custom_cluster_role_bindings {
  # Configure direct properties listed below.
}
```

## Direct properties

- [cluster_role_bindings](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/properties/use_custom_cluster_role_bindings/cluster_role_bindings/): complete subsection reference.
