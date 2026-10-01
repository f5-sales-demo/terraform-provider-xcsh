---
page_title: "discovery_k8s"
subcategory: ""
description: "discovery_k8s for xcsh_discovery."
xcsh_docs: {"aliases": [], "body_bytes": 1976, "body_sha256": "sha256:119fb5b7895c5ea23a7c3ec7e9794d89262d00baa0331be94b8cea32ee2941b4", "canonical_id": "xcsh-docs:resources:discovery:properties:discovery_k8s", "child_ids": ["xcsh-docs:resources:discovery:properties:discovery_k8s:access_info", "xcsh-docs:resources:discovery:properties:discovery_k8s:default_all", "xcsh-docs:resources:discovery:properties:discovery_k8s:namespace_mapping", "xcsh-docs:resources:discovery:properties:discovery_k8s:publish_info"], "collection_id": "xcsh-docs:resources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:resources:discovery:properties:discovery_k8s", "parent_id": "xcsh-docs:resources:discovery:reference", "path": "docs/guides/resources--discovery--properties--discovery_k8s.md", "provider_name": "discovery", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["discovery_k8s"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/discovery/properties/discovery_k8s/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "discovery_k8s for xcsh_discovery.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["discoveryCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# discovery_k8s

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md)
- [Property reference](resources--discovery--reference.md)
- discovery_k8s

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for discovery k8s.

Upstream description:

Discovery configuration for K8s.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_all",
    "namespace_mapping")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-namespace_mapping_choice": "[\"default_all\",\"namespace_mapping\"]"
}
```

Terraform syntax:

```terraform
discovery_k8s {
  # Configure direct properties listed below.
}
```

## Direct properties

- [access_info](resources--discovery--properties--discovery_k8s--access_info.md): complete subsection reference.

- [default_all](resources--discovery--properties--discovery_k8s--default_all.md): complete subsection reference.

- [namespace_mapping](resources--discovery--properties--discovery_k8s--namespace_mapping.md): complete subsection reference.

- [publish_info](resources--discovery--properties--discovery_k8s--publish_info.md): complete subsection reference.

## Next pages

- [discovery_k8s.access_info](resources--discovery--properties--discovery_k8s--access_info.md)
- [discovery_k8s.default_all](resources--discovery--properties--discovery_k8s--default_all.md)
- [discovery_k8s.namespace_mapping](resources--discovery--properties--discovery_k8s--namespace_mapping.md)
- [discovery_k8s.publish_info](resources--discovery--properties--discovery_k8s--publish_info.md)
- [Property reference](resources--discovery--reference.md)
- [xcsh_discovery](../resources/discovery.md)
