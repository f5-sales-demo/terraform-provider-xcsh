---
page_title: "ipsec.ipsec_tunnel_parameters.segment"
subcategory: ""
description: "Reference to Segment Object."
xcsh_docs: {"aliases": ["ipsec ipsec tunnel parameters segment"], "body_bytes": 1206, "body_sha256": "sha256:1bb677782e7ab58a43d9938b0ef4d4a133127219187fc6166eb26eea9c6d98cc", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:external_connector:properties:ipsec:ipsec_tunnel_parameters:segment:refs"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:external_connector:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:external_connector:properties:ipsec:ipsec_tunnel_parameters:segment", "parent_id": "xcsh-docs:data-sources:external_connector:properties:ipsec:ipsec_tunnel_parameters", "path": "documentation/data-sources/external_connector/properties/ipsec/ipsec_tunnel_parameters/segment/index.md", "product": "distributed-cloud", "provider_name": "external_connector", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-1010001120002030-1111212020010231-3000003302003103-2301312322120220-1322311223231031-3233202013113121-0223211101232333-1222132123123001", "registry_path": "docs/guides/data-sources--external_connector--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ipsec", "ipsec_tunnel_parameters", "segment"], "schema_version": 1, "sections": [{"aliases": ["ipsec ipsec tunnel parameters segment refs"], "anchor": "section", "description": "Reference to Segment Object.", "document_id": "xcsh-docs:data-sources:external_connector:properties:ipsec:ipsec_tunnel_parameters:segment:refs", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["ipsec", "ipsec_tunnel_parameters", "segment", "refs"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/external_connector/properties/ipsec/ipsec_tunnel_parameters/segment/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Reference to Segment Object.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["external_connectorCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ipsec.ipsec_tunnel_parameters.segment

Breadcrumbs:

- [xcsh_external_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/)
- [ipsec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/)
- [ipsec.ipsec_tunnel_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ipsec_tunnel_parameters/)
- ipsec.ipsec_tunnel_parameters.segment

<a id="section"></a>

Type: `"single"`. Computed.

Segment Reference Type. Reference to Segment Object.

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

- [refs](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ipsec_tunnel_parameters/segment/refs/): complete subsection reference.
