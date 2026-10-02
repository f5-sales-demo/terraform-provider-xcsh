---
page_title: "cloudflare.protected_endpoints.mobile_client"
subcategory: ""
description: "Mobile client configuration OPTIONS."
xcsh_docs: {"aliases": ["cloudflare protected endpoints mobile client"], "body_bytes": 2321, "body_sha256": "sha256:24e2fd7be0e4a5dadefbd5808ce6907e0427f25dd9a1b4f6080feb0c9bf8fef0", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:protected_application:properties:cloudflare:protected_endpoints:mobile_client:block", "xcsh-docs:data-sources:protected_application:properties:cloudflare:protected_endpoints:mobile_client:continue"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:protected_application:properties:cloudflare:protected_endpoints:mobile_client", "parent_id": "xcsh-docs:data-sources:protected_application:properties:cloudflare:protected_endpoints", "path": "documentation/data-sources/protected_application/properties/cloudflare/protected_endpoints/mobile_client/index.md", "product": "distributed-cloud", "provider_name": "protected_application", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2030023030101010-1310113333033102-1013023003332100-3233211203312330-0033001100311212-2030010200302302-1023220213200122-0223130313132223", "registry_path": "docs/guides/data-sources--protected_application--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["cloudflare", "protected_endpoints", "mobile_client"], "schema_version": 1, "sections": [{"aliases": ["block"], "anchor": "section", "description": "Block Response.", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudflare:protected_endpoints:mobile_client:block", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["cloudflare", "protected_endpoints", "mobile_client", "block"], "syntax": "attribute", "type": "object"}, {"aliases": ["continue"], "anchor": "section", "description": "Continue mitigation action.", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudflare:protected_endpoints:mobile_client:continue", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["cloudflare", "protected_endpoints", "mobile_client", "continue"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protected_application/properties/cloudflare/protected_endpoints/mobile_client/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Mobile client configuration OPTIONS.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudflare.protected_endpoints.mobile_client

Breadcrumbs:

- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/)
- [cloudflare](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudflare/)
- [cloudflare.protected_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudflare/protected_endpoints/)
- cloudflare.protected_endpoints.mobile_client

<a id="section"></a>

Type: `"single"`. Computed.

Mobile Client. Mobile client configuration OPTIONS.

Upstream description:

Mobile client configuration OPTIONS.

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

- [block](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudflare/protected_endpoints/mobile_client/block/): complete subsection reference.

- [continue](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudflare/protected_endpoints/mobile_client/continue/): complete subsection reference.

## Next pages

- [cloudflare.protected_endpoints.mobile_client.block](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudflare/protected_endpoints/mobile_client/block/)
- [cloudflare.protected_endpoints.mobile_client.continue](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudflare/protected_endpoints/mobile_client/continue/)
- [cloudflare.protected_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudflare/protected_endpoints/)
- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/)
