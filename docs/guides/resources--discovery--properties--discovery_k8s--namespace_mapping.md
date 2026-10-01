---
page_title: "discovery_k8s.namespace_mapping"
subcategory: ""
description: "discovery_k8s.namespace_mapping for xcsh_discovery."
xcsh_docs: {"aliases": [], "body_bytes": 1397, "body_sha256": "sha256:ef649f226062d5ae9d0c08504b27eacaee8439ef7b13b2286e734f3328267917", "canonical_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:namespace_mapping", "child_ids": ["xcsh-docs:resources:discovery:properties:discovery_k8s:namespace_mapping:items"], "collection_id": "xcsh-docs:resources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:resources:discovery:properties:discovery_k8s:namespace_mapping", "parent_id": "xcsh-docs:resources:discovery:properties:discovery_k8s", "path": "docs/guides/resources--discovery--properties--discovery_k8s--namespace_mapping.md", "provider_name": "discovery", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["discovery_k8s", "namespace_mapping"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/discovery/properties/discovery_k8s/namespace_mapping/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "discovery_k8s.namespace_mapping for xcsh_discovery.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["discoveryCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# discovery_k8s.namespace_mapping

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md)
- [Property reference](resources--discovery--reference.md)
- [discovery_k8s](resources--discovery--properties--discovery_k8s.md)
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

- [items](resources--discovery--properties--discovery_k8s--namespace_mapping--items.md): complete subsection reference.

## Next pages

- [discovery_k8s.namespace_mapping.items](resources--discovery--properties--discovery_k8s--namespace_mapping--items.md)
- [discovery_k8s](resources--discovery--properties--discovery_k8s.md)
- [xcsh_discovery](../resources/discovery.md)
