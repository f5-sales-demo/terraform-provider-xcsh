---
page_title: "discovery_k8s.namespace_mapping"
subcategory: ""
description: "Select the mapping between K8s namespaces from which services will be discovered and App Namespace to which the discovered services will be shared."
xcsh_docs: {"aliases": ["discovery k8s namespace mapping"], "body_bytes": 1503, "body_sha256": "sha256:78a34edd2923e27863e0c598a186cd91af2668edc3f3c26b8be4a9157312408a", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:discovery:properties:discovery_k8s:namespace_mapping:items"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:discovery:properties:discovery_k8s:namespace_mapping", "parent_id": "xcsh-docs:data-sources:discovery:properties:discovery_k8s", "path": "documentation/data-sources/discovery/properties/discovery_k8s/namespace_mapping/index.md", "product": "distributed-cloud", "provider_name": "discovery", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3303013301213222-3332332202321221-2311112031001231-0323332330033132-2230203122010120-1131213122133310-1033023232121000-2320222012200021", "registry_path": "docs/guides/data-sources--discovery--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["discovery_k8s", "namespace_mapping"], "schema_version": 1, "sections": [{"aliases": ["items"], "anchor": "section", "description": "Map K8s namespace(s) to App Namespaces. In Shared Configuration, Discovered Services can only be mapped to a single App Namespace, which is determined by the first matched regex.", "document_id": "xcsh-docs:data-sources:discovery:properties:discovery_k8s:namespace_mapping:items", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["discovery_k8s", "namespace_mapping", "items"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/discovery/properties/discovery_k8s/namespace_mapping/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Select the mapping between K8s namespaces from which services will be discovered and App Namespace to which the discovered services will be shared.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["discoveryCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# discovery_k8s.namespace_mapping

Breadcrumbs:

- [xcsh_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/)
- [discovery_k8s](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_k8s/)
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

- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_k8s/namespace_mapping/items/): complete subsection reference.

## Next pages

- [discovery_k8s.namespace_mapping.items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_k8s/namespace_mapping/items/)
- [discovery_k8s](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_k8s/)
- [xcsh_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/)
