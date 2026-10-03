---
page_title: "discovery_k8s"
subcategory: ""
description: "Discovery configuration for K8s."
xcsh_docs: {"aliases": ["discovery k8s"], "body_bytes": 2320, "body_sha256": "sha256:c99c912a68a6b4d93a2cdc2544828f950af97dcebd1c620684cfe874e1c39f03", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:discovery:properties:discovery_k8s:access_info", "xcsh-docs:data-sources:discovery:properties:discovery_k8s:default_all", "xcsh-docs:data-sources:discovery:properties:discovery_k8s:namespace_mapping", "xcsh-docs:data-sources:discovery:properties:discovery_k8s:publish_info"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:discovery:properties:discovery_k8s", "parent_id": "xcsh-docs:data-sources:discovery:reference", "path": "documentation/data-sources/discovery/properties/discovery_k8s/index.md", "product": "distributed-cloud", "provider_name": "discovery", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0330013311111202-1322311031122022-3320203133110210-1213102010333012-0022110212111130-2223102233100201-0022000121230323-1321310123011322", "registry_path": "docs/guides/data-sources--discovery--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["discovery_k8s"], "schema_version": 1, "sections": [{"aliases": ["discovery k8s access info"], "anchor": "section", "description": "K8s API server access.", "document_id": "xcsh-docs:data-sources:discovery:properties:discovery_k8s:access_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["discovery_k8s", "access_info"], "syntax": "attribute", "type": "object"}, {"aliases": ["discovery k8s default all"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:discovery:properties:discovery_k8s:default_all", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["discovery_k8s", "default_all"], "syntax": "attribute", "type": "object"}, {"aliases": ["discovery k8s namespace mapping"], "anchor": "section", "description": "Select the mapping between K8s namespaces from which services will be discovered and App Namespace to which the discovered services will be shared.", "document_id": "xcsh-docs:data-sources:discovery:properties:discovery_k8s:namespace_mapping", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["discovery_k8s", "namespace_mapping"], "syntax": "attribute", "type": "object"}, {"aliases": ["discovery k8s publish info"], "anchor": "section", "description": "K8s Configuration to publish VIPs.", "document_id": "xcsh-docs:data-sources:discovery:properties:discovery_k8s:publish_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["discovery_k8s", "publish_info"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/discovery/properties/discovery_k8s/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Discovery configuration for K8s.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["discoveryCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# discovery_k8s

Breadcrumbs:

- [xcsh_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/)
- discovery_k8s

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for discovery k8s.

Upstream description:

Discovery configuration for K8s.

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

## Direct properties

- [access_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_k8s/access_info/): complete subsection reference.

- [default_all](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_k8s/default_all/): complete subsection reference.

- [namespace_mapping](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_k8s/namespace_mapping/): complete subsection reference.

- [publish_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_k8s/publish_info/): complete subsection reference.

## Next pages

- [discovery_k8s.access_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_k8s/access_info/)
- [discovery_k8s.default_all](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_k8s/default_all/)
- [discovery_k8s.namespace_mapping](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_k8s/namespace_mapping/)
- [discovery_k8s.publish_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_k8s/publish_info/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/)
- [xcsh_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/)
