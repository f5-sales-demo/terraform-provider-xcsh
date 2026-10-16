---
page_title: "cloudflare.protected_endpoints.web_client"
subcategory: ""
description: "Web client configuration OPTIONS."
xcsh_docs: {"aliases": ["cloudflare protected endpoints web client"], "body_bytes": 1720, "body_sha256": "sha256:a152e3ad7311691c7584dba36b9334c0072dc7695a240b85df1e2ef6c9805750", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:protected_application:properties:cloudflare:protected_endpoints:web_client:block", "xcsh-docs:data-sources:protected_application:properties:cloudflare:protected_endpoints:web_client:continue", "xcsh-docs:data-sources:protected_application:properties:cloudflare:protected_endpoints:web_client:redirect"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:protected_application:properties:cloudflare:protected_endpoints:web_client", "parent_id": "xcsh-docs:data-sources:protected_application:properties:cloudflare:protected_endpoints", "path": "documentation/data-sources/protected_application/properties/cloudflare/protected_endpoints/web_client/index.md", "product": "distributed-cloud", "provider_name": "protected_application", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "data-sources", "registry_anchor": "canonical-3230221231323110-2000013213332023-3011112223321020-1023203303123102-3000330332112310-3132023301231112-0121133032331033-1331200011200311", "registry_path": "docs/guides/data-sources--protected_application--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["cloudflare", "protected_endpoints", "web_client"], "schema_version": 1, "sections": [{"aliases": ["cloudflare protected endpoints web client block"], "anchor": "section", "description": "Block Response.", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudflare:protected_endpoints:web_client:block", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["cloudflare", "protected_endpoints", "web_client", "block"], "syntax": "attribute", "type": "object"}, {"aliases": ["cloudflare protected endpoints web client continue"], "anchor": "section", "description": "Continue mitigation action.", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudflare:protected_endpoints:web_client:continue", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["cloudflare", "protected_endpoints", "web_client", "continue"], "syntax": "attribute", "type": "object"}, {"aliases": ["cloudflare protected endpoints web client redirect"], "anchor": "section", "description": "Redirect.", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudflare:protected_endpoints:web_client:redirect", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["cloudflare", "protected_endpoints", "web_client", "redirect"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protected_application/properties/cloudflare/protected_endpoints/web_client/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "Web client configuration OPTIONS.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudflare.protected_endpoints.web_client

Breadcrumbs:

- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/)
- [cloudflare](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudflare/)
- [cloudflare.protected_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudflare/protected_endpoints/)
- cloudflare.protected_endpoints.web_client

<a id="section"></a>

Type: `"single"`. Computed.

Web Client. Web client configuration OPTIONS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-mitigation": "[\"block\",\"continue\",\"redirect\"]"
}
```

## Direct properties

- [block](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudflare/protected_endpoints/web_client/block/): complete subsection reference.

- [continue](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudflare/protected_endpoints/web_client/continue/): complete subsection reference.

- [redirect](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudflare/protected_endpoints/web_client/redirect/): complete subsection reference.
