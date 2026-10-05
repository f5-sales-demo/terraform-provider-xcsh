---
page_title: "discovery_k8s.namespace_mapping"
subcategory: ""
description: "Select the mapping between K8s namespaces from which services will be discovered and App Namespace to which the discovered services will be shared."
xcsh_docs: {"aliases": ["discovery k8s namespace mapping"], "body_bytes": 1752, "body_sha256": "sha256:086b934776ca4dd31dc248cd493266735d48c4ecc17ab14a14bae6243ea51826", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:discovery:properties:discovery_k8s:namespace_mapping:items"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:resources:discovery:properties:discovery_k8s:namespace_mapping", "parent_id": "xcsh-docs:resources:discovery:properties:discovery_k8s", "path": "documentation/resources/discovery/properties/discovery_k8s/namespace_mapping/index.md", "product": "distributed-cloud", "provider_name": "discovery", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-0321010022030131-3311200301332101-3132022310131230-1233220021130131-3210333300201303-3321221200212203-2130012322120030-1013021110331232", "registry_path": "docs/guides/resources--discovery--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "discovery_k8s.namespace_mapping:RequiredObjectAttributes:items", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:namespace_mapping:items", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["discovery_k8s", "namespace_mapping"], "schema_version": 1, "sections": [{"aliases": ["discovery k8s namespace mapping items"], "anchor": "section", "description": "Map K8s namespace(s) to App Namespaces. In Shared Configuration, Discovered Services can only be mapped to a single App Namespace, which is determined by the first matched regex.", "document_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:namespace_mapping:items", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["discovery_k8s", "namespace_mapping", "items"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/discovery/properties/discovery_k8s/namespace_mapping/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Select the mapping between K8s namespaces from which services will be discovered and App Namespace to which the discovered services will be shared.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["discoveryCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# discovery_k8s.namespace_mapping

Breadcrumbs:

- [xcsh_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/)
- [discovery_k8s](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/)
- discovery_k8s.namespace_mapping

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Select the mapping between K8s namespaces from which services will be discovered and App Namespace
to which the discovered services will be shared.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("items")}
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
namespace_mapping {
  # Configure direct properties listed below.
}
```

## Direct properties

- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/namespace_mapping/items/): complete subsection reference.

## Next pages

- [discovery_k8s.namespace_mapping.items](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/namespace_mapping/items/)
- [discovery_k8s](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/)
- [xcsh_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/)
