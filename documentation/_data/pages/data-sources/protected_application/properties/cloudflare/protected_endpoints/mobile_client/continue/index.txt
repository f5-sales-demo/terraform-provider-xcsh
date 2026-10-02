---
page_title: "cloudflare.protected_endpoints.mobile_client.continue"
subcategory: ""
description: "Continue mitigation action."
xcsh_docs: {"aliases": ["cloudflare protected endpoints mobile client continue"], "body_bytes": 2662, "body_sha256": "sha256:9e7d61db04dc7b553c7ae73a456316cb9e12a7ba53cb84a1a2549908bd243d4a", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:protected_application:properties:cloudflare:protected_endpoints:mobile_client:continue:add_header", "xcsh-docs:data-sources:protected_application:properties:cloudflare:protected_endpoints:mobile_client:continue:no_header"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:protected_application:properties:cloudflare:protected_endpoints:mobile_client:continue", "parent_id": "xcsh-docs:data-sources:protected_application:properties:cloudflare:protected_endpoints:mobile_client", "path": "documentation/data-sources/protected_application/properties/cloudflare/protected_endpoints/mobile_client/continue/index.md", "product": "distributed-cloud", "provider_name": "protected_application", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0311320122200022-2213032123021112-1102223311223231-0210303203321020-0001233023123033-3002220212021321-3332311010210313-0313201122120000", "registry_path": "docs/guides/data-sources--protected_application--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["cloudflare", "protected_endpoints", "mobile_client", "continue"], "schema_version": 1, "sections": [{"aliases": ["add header"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudflare:protected_endpoints:mobile_client:continue:add_header", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cloudflare", "protected_endpoints", "mobile_client", "continue", "add_header"], "syntax": "attribute", "type": "object"}, {"aliases": ["no header"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudflare:protected_endpoints:mobile_client:continue:no_header", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cloudflare", "protected_endpoints", "mobile_client", "continue", "no_header"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protected_application/properties/cloudflare/protected_endpoints/mobile_client/continue/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Continue mitigation action.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudflare.protected_endpoints.mobile_client.continue

Breadcrumbs:

- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/)
- [cloudflare](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudflare/)
- [cloudflare.protected_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudflare/protected_endpoints/)
- [cloudflare.protected_endpoints.mobile_client](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudflare/protected_endpoints/mobile_client/)
- cloudflare.protected_endpoints.mobile_client.continue

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

- [add_header](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudflare/protected_endpoints/mobile_client/continue/add_header/): complete subsection reference.

- [no_header](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudflare/protected_endpoints/mobile_client/continue/no_header/): complete subsection reference.

## Next pages

- [cloudflare.protected_endpoints.mobile_client.continue.add_header](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudflare/protected_endpoints/mobile_client/continue/add_header/)
- [cloudflare.protected_endpoints.mobile_client.continue.no_header](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudflare/protected_endpoints/mobile_client/continue/no_header/)
- [cloudflare.protected_endpoints.mobile_client](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudflare/protected_endpoints/mobile_client/)
- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/)
