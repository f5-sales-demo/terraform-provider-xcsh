---
page_title: "cloudfront.protected_endpoints.web_mobile_client.continue_web"
subcategory: ""
description: "Continue mitigation action."
xcsh_docs: {"aliases": ["cloudfront protected endpoints web mobile client continue web"], "body_bytes": 1838, "body_sha256": "sha256:b638180c7396039152d28cba18fcadf5d2f59478eba0e155f8e18ee86853ff17", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:web_mobile_client:continue_web:add_header", "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:web_mobile_client:continue_web:no_header"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:web_mobile_client:continue_web", "parent_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:web_mobile_client", "path": "documentation/data-sources/protected_application/properties/cloudfront/protected_endpoints/web_mobile_client/continue_web/index.md", "product": "distributed-cloud", "provider_name": "protected_application", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-3003233231331133-0211000100110230-0203330230231230-2200323233330120-1313222223212200-0003000201210031-1000122130123023-1322012210211030", "registry_path": "docs/guides/data-sources--protected_application--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["cloudfront", "protected_endpoints", "web_mobile_client", "continue_web"], "schema_version": 1, "sections": [{"aliases": ["cloudfront protected endpoints web mobile client continue web add header"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:web_mobile_client:continue_web:add_header", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cloudfront", "protected_endpoints", "web_mobile_client", "continue_web", "add_header"], "syntax": "attribute", "type": "object"}, {"aliases": ["cloudfront protected endpoints web mobile client continue web no header"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:web_mobile_client:continue_web:no_header", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cloudfront", "protected_endpoints", "web_mobile_client", "continue_web", "no_header"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protected_application/properties/cloudfront/protected_endpoints/web_mobile_client/continue_web/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Continue mitigation action.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudfront.protected_endpoints.web_mobile_client.continue_web

Breadcrumbs:

- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/)
- [cloudfront](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/)
- [cloudfront.protected_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/)
- [cloudfront.protected_endpoints.web_mobile_client](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/web_mobile_client/)
- cloudfront.protected_endpoints.web_mobile_client.continue_web

<a id="section"></a>

Type: `"single"`. Computed.

Select Continue Bot Mitigation Action. Continue mitigation action.

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

- [add_header](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/web_mobile_client/continue_web/add_header/): complete subsection reference.

- [no_header](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/web_mobile_client/continue_web/no_header/): complete subsection reference.
