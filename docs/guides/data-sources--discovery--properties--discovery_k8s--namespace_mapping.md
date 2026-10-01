---
page_title: "discovery_k8s.namespace_mapping"
subcategory: ""
description: "discovery_k8s.namespace_mapping for xcsh_discovery."
xcsh_docs: {"aliases": [], "body_bytes": 1148, "body_sha256": "sha256:065a198509724280611253be21e21c3c2756de5e832a83cb46bfb609a6ba90a6", "canonical_id": "xcsh-docs:data-sources:discovery:properties:discovery_k8s:namespace_mapping", "child_ids": ["xcsh-docs:data-sources:discovery:properties:discovery_k8s:namespace_mapping:items"], "collection_id": "xcsh-docs:data-sources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:discovery:properties:discovery_k8s:namespace_mapping", "parent_id": "xcsh-docs:data-sources:discovery:properties:discovery_k8s", "path": "docs/guides/data-sources--discovery--properties--discovery_k8s--namespace_mapping.md", "provider_name": "discovery", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["discovery_k8s", "namespace_mapping"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/discovery/properties/discovery_k8s/namespace_mapping/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "discovery_k8s.namespace_mapping for xcsh_discovery.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["discoveryCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# discovery_k8s.namespace_mapping

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md)
- [Property reference](data-sources--discovery--reference.md)
- [discovery_k8s](data-sources--discovery--properties--discovery_k8s.md)
- discovery_k8s.namespace_mapping

<a id="section"></a>

Type: `"single"`. Computed.

Select the mapping between K8s namespaces from which services will be discovered and App Namespace
to which the discovered services will be shared.

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

- [items](data-sources--discovery--properties--discovery_k8s--namespace_mapping--items.md): complete subsection reference.

## Next pages

- [discovery_k8s.namespace_mapping.items](data-sources--discovery--properties--discovery_k8s--namespace_mapping--items.md)
- [discovery_k8s](data-sources--discovery--properties--discovery_k8s.md)
- [xcsh_discovery](../data-sources/discovery.md)
