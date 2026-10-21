---
page_title: "gre.gre_parameters.segment"
subcategory: ""
description: "Reference to Segment Object."
xcsh_docs: {"aliases": ["gre gre parameters segment"], "body_bytes": 1147, "body_sha256": "sha256:caa279b24d8a0c2a745837725b9a00e972bdd679edb94dbda492289867346c1e", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:external_connector:properties:gre:gre_parameters:segment:refs"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:external_connector:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:external_connector:properties:gre:gre_parameters:segment", "parent_id": "xcsh-docs:data-sources:external_connector:properties:gre:gre_parameters", "path": "documentation/data-sources/external_connector/properties/gre/gre_parameters/segment/index.md", "product": "distributed-cloud", "provider_name": "external_connector", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-3010322332310332-1103122330310323-3313213220203120-1332330202322012-1112112013111101-2003331112332000-2031133300111100-2013202322300102", "registry_path": "docs/guides/data-sources--external_connector--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["gre", "gre_parameters", "segment"], "schema_version": 1, "sections": [{"aliases": ["gre gre parameters segment refs"], "anchor": "section", "description": "Reference to Segment Object.", "document_id": "xcsh-docs:data-sources:external_connector:properties:gre:gre_parameters:segment:refs", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["gre", "gre_parameters", "segment", "refs"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/external_connector/properties/gre/gre_parameters/segment/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "Reference to Segment Object.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["external_connectorCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# gre.gre_parameters.segment

Breadcrumbs:

- [xcsh_external_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/)
- [gre](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/gre/)
- [gre.gre_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/gre/gre_parameters/)
- gre.gre_parameters.segment

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

- [refs](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/gre/gre_parameters/segment/refs/): complete subsection reference.
