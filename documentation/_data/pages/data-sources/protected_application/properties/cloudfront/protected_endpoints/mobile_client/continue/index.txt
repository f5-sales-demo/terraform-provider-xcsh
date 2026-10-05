---
page_title: "cloudfront.protected_endpoints.mobile_client.continue"
subcategory: ""
description: "Continue mitigation action."
xcsh_docs: {"aliases": ["cloudfront protected endpoints mobile client continue"], "body_bytes": 2662, "body_sha256": "sha256:a348e81a47cdabc22fbcc595a31878f94e46290765d086a198fe86adda9cf64b", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:mobile_client:continue:add_header", "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:mobile_client:continue:no_header"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:mobile_client:continue", "parent_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:mobile_client", "path": "documentation/data-sources/protected_application/properties/cloudfront/protected_endpoints/mobile_client/continue/index.md", "product": "distributed-cloud", "provider_name": "protected_application", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-0023212102233030-1310113233310133-1001111232030332-0031222031303032-1012003301313330-3122230011323003-1120310200220230-0010330121001203", "registry_path": "docs/guides/data-sources--protected_application--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["cloudfront", "protected_endpoints", "mobile_client", "continue"], "schema_version": 1, "sections": [{"aliases": ["cloudfront protected endpoints mobile client continue add header"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:mobile_client:continue:add_header", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cloudfront", "protected_endpoints", "mobile_client", "continue", "add_header"], "syntax": "attribute", "type": "object"}, {"aliases": ["cloudfront protected endpoints mobile client continue no header"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:mobile_client:continue:no_header", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cloudfront", "protected_endpoints", "mobile_client", "continue", "no_header"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protected_application/properties/cloudfront/protected_endpoints/mobile_client/continue/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Continue mitigation action.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudfront.protected_endpoints.mobile_client.continue

Breadcrumbs:

- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/)
- [cloudfront](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/)
- [cloudfront.protected_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/)
- [cloudfront.protected_endpoints.mobile_client](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/mobile_client/)
- cloudfront.protected_endpoints.mobile_client.continue

<a id="section"></a>

Type: `"single"`. Computed.

Select Continue Bot Mitigation Action. Continue mitigation action.

Upstream description:

Continue mitigation action.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-add_header_choice": "[\"add_header\",\"no_header\"]"
}
```

## Direct properties

- [add_header](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/mobile_client/continue/add_header/): complete subsection reference.

- [no_header](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/mobile_client/continue/no_header/): complete subsection reference.

## Next pages

- [cloudfront.protected_endpoints.mobile_client.continue.add_header](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/mobile_client/continue/add_header/)
- [cloudfront.protected_endpoints.mobile_client.continue.no_header](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/mobile_client/continue/no_header/)
- [cloudfront.protected_endpoints.mobile_client](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/mobile_client/)
- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/)
