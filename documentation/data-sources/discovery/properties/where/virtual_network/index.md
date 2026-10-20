---
page_title: "where.virtual_network"
subcategory: ""
description: "This specifies a direct reference to a network configuration object."
xcsh_docs: {"aliases": ["where virtual network"], "body_bytes": 962, "body_sha256": "sha256:7cf39e47f930dc511c8a6c8231d77e4b4c7030cf3de62d574c558bec514731fa", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:discovery:properties:where:virtual_network:ref"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:discovery:properties:where:virtual_network", "parent_id": "xcsh-docs:data-sources:discovery:properties:where", "path": "documentation/data-sources/discovery/properties/where/virtual_network/index.md", "product": "distributed-cloud", "provider_name": "discovery", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-2000100203122133-0301232002323310-3233310132013200-2022221330110313-1210122133103010-3121123023112123-0033000321111322-1210332230303323", "registry_path": "docs/guides/data-sources--discovery--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["where", "virtual_network"], "schema_version": 1, "sections": [{"aliases": ["where virtual network ref"], "anchor": "section", "description": "A virtual network direct reference.", "document_id": "xcsh-docs:data-sources:discovery:properties:where:virtual_network:ref", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["where", "virtual_network", "ref"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/discovery/properties/where/virtual_network/index.txt", "spec_pin_digest": "sha256:2a27992e2c450f8ea5ba537ebc7f5c26b0dfbc946db7ef0e7e86625dba9fb492", "summary": "This specifies a direct reference to a network configuration object.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.4", "schema_components": ["discoveryCreateRequest"], "target_commit": "a9149265e7737f6841c289490cb1a3ca946de83e"}}
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
