---
page_title: "snat_pool"
subcategory: "Networking"
description: "SNAT Pool configuration."
xcsh_docs: {"aliases": ["snat pool"], "body_bytes": 1586, "body_sha256": "sha256:00994d20baf11dc9f8b5b15fe619a223c41a6b8360bdbc8b277040fa6b576e18", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:endpoint:properties:snat_pool:no_snat_pool", "xcsh-docs:data-sources:endpoint:properties:snat_pool:snat_pool"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:endpoint:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:endpoint:properties:snat_pool", "parent_id": "xcsh-docs:data-sources:endpoint:reference", "path": "documentation/data-sources/endpoint/properties/snat_pool/index.md", "product": "distributed-cloud", "provider_name": "endpoint", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-1023321130330223-3012133131021311-2313201210232103-2201111323311030-0110100021311312-3133330013320201-0131111221101333-1312112211221301", "registry_path": "docs/guides/data-sources--endpoint--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["snat_pool"], "schema_version": 1, "sections": [{"aliases": ["snat pool no snat pool"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:endpoint:properties:snat_pool:no_snat_pool", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["snat_pool", "no_snat_pool"], "syntax": "attribute", "type": "object"}, {"aliases": ["snat pool snat pool"], "anchor": "section", "description": "List of IPv4 prefixes that represent an endpoint.", "document_id": "xcsh-docs:data-sources:endpoint:properties:snat_pool:snat_pool", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["snat_pool", "snat_pool"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/endpoint/properties/snat_pool/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "SNAT Pool configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["endpointCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
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

Upstream description:

SNAT Pool configuration.

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

## Next pages

- [snat_pool.no_snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/snat_pool/no_snat_pool/)
- [snat_pool.snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/snat_pool/snat_pool/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/)
- [xcsh_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/)
