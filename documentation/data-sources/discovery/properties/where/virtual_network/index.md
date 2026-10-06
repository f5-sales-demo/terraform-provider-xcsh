---
page_title: "where.virtual_network"
subcategory: ""
description: "This specifies a direct reference to a network configuration object."
xcsh_docs: {"aliases": ["where virtual network"], "body_bytes": 962, "body_sha256": "sha256:7cf39e47f930dc511c8a6c8231d77e4b4c7030cf3de62d574c558bec514731fa", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:discovery:properties:where:virtual_network:ref"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:discovery:properties:where:virtual_network", "parent_id": "xcsh-docs:data-sources:discovery:properties:where", "path": "documentation/data-sources/discovery/properties/where/virtual_network/index.md", "product": "distributed-cloud", "provider_name": "discovery", "provider_schema_digest": "sha256:057968f86e4ef0ae0087dd4d6097131e285b60998d97655ee1a02c00315f6b0f", "provider_type": "data-sources", "registry_anchor": "canonical-2000100203122133-0301232002323310-3233310132013200-2022221330110313-1210122133103010-3121123023112123-0033000321111322-1210332230303323", "registry_path": "docs/guides/data-sources--discovery--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["where", "virtual_network"], "schema_version": 1, "sections": [{"aliases": ["where virtual network ref"], "anchor": "section", "description": "A virtual network direct reference.", "document_id": "xcsh-docs:data-sources:discovery:properties:where:virtual_network:ref", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["where", "virtual_network", "ref"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/discovery/properties/where/virtual_network/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "This specifies a direct reference to a network configuration object.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["discoveryCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# where.virtual_network

Breadcrumbs:

- [xcsh_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/)
- [where](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/where/)
- where.virtual_network

<a id="section"></a>

Type: `"single"`. Computed.

This specifies a direct reference to a network configuration object.

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

- [ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/where/virtual_network/ref/): complete subsection reference.
