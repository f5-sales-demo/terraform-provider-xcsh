---
page_title: "discovery_k8s.namespace_mapping"
subcategory: ""
description: "Select the mapping between K8s namespaces from which services will be discovered and App Namespace to which the discovered services will be shared."
xcsh_docs: {"aliases": ["discovery k8s namespace mapping"], "body_bytes": 1091, "body_sha256": "sha256:1887354cce6a3e8975ad1820d217959895240be0decc407b53a93a67e536d332", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:discovery:properties:discovery_k8s:namespace_mapping:items"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:discovery:properties:discovery_k8s:namespace_mapping", "parent_id": "xcsh-docs:data-sources:discovery:properties:discovery_k8s", "path": "documentation/data-sources/discovery/properties/discovery_k8s/namespace_mapping/index.md", "product": "distributed-cloud", "provider_name": "discovery", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-3303013301213222-3332332202321221-2311112031001231-0323332330033132-2230203122010120-1131213122133310-1033023232121000-2320222012200021", "registry_path": "docs/guides/data-sources--discovery--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["discovery_k8s", "namespace_mapping"], "schema_version": 1, "sections": [{"aliases": ["discovery k8s namespace mapping items"], "anchor": "section", "description": "Map K8s namespace(s) to App Namespaces. In Shared Configuration, Discovered Services can only be mapped to a single App Namespace, which is determined by the first matched regex.", "document_id": "xcsh-docs:data-sources:discovery:properties:discovery_k8s:namespace_mapping:items", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["discovery_k8s", "namespace_mapping", "items"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/discovery/properties/discovery_k8s/namespace_mapping/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Select the mapping between K8s namespaces from which services will be discovered and App Namespace to which the discovered services will be shared.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["discoveryCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
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
