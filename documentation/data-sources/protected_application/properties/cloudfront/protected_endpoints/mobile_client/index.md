---
page_title: "cloudfront.protected_endpoints.mobile_client"
subcategory: ""
description: "Mobile client configuration OPTIONS."
xcsh_docs: {"aliases": ["cloudfront protected endpoints mobile client"], "body_bytes": 1524, "body_sha256": "sha256:5a6d468e4ad444bfe12e2cd19575b30e00188bc970140eabb7935dadba94c1a6", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:mobile_client:block", "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:mobile_client:continue"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:mobile_client", "parent_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints", "path": "documentation/data-sources/protected_application/properties/cloudfront/protected_endpoints/mobile_client/index.md", "product": "distributed-cloud", "provider_name": "protected_application", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-0300100133111022-1320133032023113-0200312122311310-2330312002011033-0123110113021233-1111122331121201-0322031011211223-1323112133031322", "registry_path": "docs/guides/data-sources--protected_application--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["cloudfront", "protected_endpoints", "mobile_client"], "schema_version": 1, "sections": [{"aliases": ["cloudfront protected endpoints mobile client block"], "anchor": "section", "description": "Block Response.", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:mobile_client:block", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["cloudfront", "protected_endpoints", "mobile_client", "block"], "syntax": "attribute", "type": "object"}, {"aliases": ["cloudfront protected endpoints mobile client continue"], "anchor": "section", "description": "Continue mitigation action.", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:mobile_client:continue", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["cloudfront", "protected_endpoints", "mobile_client", "continue"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protected_application/properties/cloudfront/protected_endpoints/mobile_client/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "Mobile client configuration OPTIONS.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudfront.protected_endpoints.mobile_client

Breadcrumbs:

- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/)
- [cloudfront](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/)
- [cloudfront.protected_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/)
- cloudfront.protected_endpoints.mobile_client

<a id="section"></a>

Type: `"single"`. Computed.

Mobile Client. Mobile client configuration OPTIONS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-mitigation": "[\"block\",\"continue\"]"
}
```

## Direct properties

- [block](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/mobile_client/block/): complete subsection reference.

- [continue](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/mobile_client/continue/): complete subsection reference.
