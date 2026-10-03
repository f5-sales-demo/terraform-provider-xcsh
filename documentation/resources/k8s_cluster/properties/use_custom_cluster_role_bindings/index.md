---
page_title: "use_custom_cluster_role_bindings"
subcategory: ""
description: "List of active cluster role binding list for a K8s cluster."
xcsh_docs: {"aliases": ["use custom cluster role bindings"], "body_bytes": 2315, "body_sha256": "sha256:b47a152558dedcc78abce801cc6758eeaa5153b16e2d74af9a9d2feb13a16cab", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:k8s_cluster:properties:use_custom_cluster_role_bindings:cluster_role_bindings"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:k8s_cluster:collection", "completeness": "complete", "id": "xcsh-docs:resources:k8s_cluster:properties:use_custom_cluster_role_bindings", "parent_id": "xcsh-docs:resources:k8s_cluster:reference", "path": "documentation/resources/k8s_cluster/properties/use_custom_cluster_role_bindings/index.md", "product": "distributed-cloud", "provider_name": "k8s_cluster", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0211302322000100-0203322011230121-0312133132101320-1031210001131331-1031230221102200-3221230333033330-3133111123333321-0133230211020211", "registry_path": "docs/guides/resources--k8s_cluster--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "use_custom_cluster_role_bindings:RequiredObjectAttributes:cluster_role_bindings", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:k8s_cluster:properties:use_custom_cluster_role_bindings:cluster_role_bindings", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["use_custom_cluster_role_bindings"], "schema_version": 1, "sections": [{"aliases": ["use custom cluster role bindings cluster role bindings"], "anchor": "section", "description": "List of active cluster role binding list for a K8s cluster.", "document_id": "xcsh-docs:resources:k8s_cluster:properties:use_custom_cluster_role_bindings:cluster_role_bindings", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-use_custom_cluster_role_bindings--cluster_role_bindings--name", "enforcement": "provider-schema", "group": "use_custom_cluster_role_bindings.cluster_role_bindings:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:k8s_cluster:properties:use_custom_cluster_role_bindings:cluster_role_bindings", "type": "requires"}], "schema_path": ["use_custom_cluster_role_bindings", "cluster_role_bindings"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/k8s_cluster/properties/use_custom_cluster_role_bindings/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "List of active cluster role binding list for a K8s cluster.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["k8s_clusterCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

Upstream description:

List of active cluster role binding list for a K8s cluster.

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

## Next pages

- [use_custom_cluster_role_bindings.cluster_role_bindings](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/properties/use_custom_cluster_role_bindings/cluster_role_bindings/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/properties/)
- [xcsh_k8s_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/)
