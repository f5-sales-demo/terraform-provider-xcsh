---
page_title: "snat_pool"
subcategory: "Networking"
description: "SNAT Pool configuration."
xcsh_docs: {"aliases": ["snat pool"], "body_bytes": 1033, "body_sha256": "sha256:1c219c5cef44370a612ac426c4cb9fad2d6dc737e23d36d1312de56a3b918adc", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:endpoint:properties:snat_pool:no_snat_pool", "xcsh-docs:data-sources:endpoint:properties:snat_pool:snat_pool"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:endpoint:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:endpoint:properties:snat_pool", "parent_id": "xcsh-docs:data-sources:endpoint:reference", "path": "documentation/data-sources/endpoint/properties/snat_pool/index.md", "product": "distributed-cloud", "provider_name": "endpoint", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-1023321130330223-3012133131021311-2313201210232103-2201111323311030-0110100021311312-3133330013320201-0131111221101333-1312112211221301", "registry_path": "docs/guides/data-sources--endpoint--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["snat_pool"], "schema_version": 1, "sections": [{"aliases": ["snat pool no snat pool"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:endpoint:properties:snat_pool:no_snat_pool", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["snat_pool", "no_snat_pool"], "syntax": "attribute", "type": "object"}, {"aliases": ["snat pool snat pool"], "anchor": "section", "description": "List of IPv4 prefixes that represent an endpoint.", "document_id": "xcsh-docs:data-sources:endpoint:properties:snat_pool:snat_pool", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["snat_pool", "snat_pool"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/endpoint/properties/snat_pool/index.txt", "spec_pin_digest": "sha256:2a27992e2c450f8ea5ba537ebc7f5c26b0dfbc946db7ef0e7e86625dba9fb492", "summary": "SNAT Pool configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.4", "schema_components": ["endpointCreateRequest"], "target_commit": "a9149265e7737f6841c289490cb1a3ca946de83e"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# snat_pool

Breadcrumbs:

- [xcsh_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/)
- snat_pool

<a id="section"></a>

Type: `"single"`. Computed.

SNAT Pool. SNAT Pool configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-snat_pool_choice": "[\"no_snat_pool\",\"snat_pool\"]"
}
```

## Direct properties

- [no_snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/snat_pool/no_snat_pool/): complete subsection reference.

- [snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/snat_pool/snat_pool/): complete subsection reference.
