---
page_title: "adobe_commerce_connector"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["adobe commerce connector"], "body_bytes": 2441, "body_sha256": "sha256:53b00c7d348c869746f55f0c1f93d61f055d29026ef18dc3b12dcd693b78b22c", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:protected_application:properties:adobe_commerce_connector", "parent_id": "xcsh-docs:data-sources:protected_application:reference", "path": "documentation/data-sources/protected_application/properties/adobe_commerce_connector/index.md", "product": "distributed-cloud", "provider_name": "protected_application", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2131122002322200-1211102013312210-0332313001322211-3203310002112222-1102230310021123-1302303130013203-0210310220203232-1201203310213203", "registry_path": "docs/guides/data-sources--protected_application--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["adobe_commerce_connector"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protected_application/properties/adobe_commerce_connector/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# adobe_commerce_connector

Breadcrumbs:

- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/)
- adobe_commerce_connector

<a id="section"></a>

Type: `["object", {}]`. Computed.

\[OneOf: adobe\_commerce\_connector, big\_ip\_iapp, cloudflare, cloudfront, custom\_connector,
f5\_big\_ip, salesforce\_commerce\_connector\] Configuration parameter for adobe commerce connector.

Upstream description:

This can be used for messages where no values are needed.

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

OneOf alternatives in this subsection:

- [adobe_commerce_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/adobe_commerce_connector/#section)
- [big_ip_iapp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/big_ip_iapp/#section)
- [cloudflare](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudflare/#section)
- [cloudfront](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/#section)
- [custom_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/custom_connector/#section)
- [f5_big_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/f5_big_ip/#section)
- [salesforce_commerce_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/salesforce_commerce_connector/#section)

Select alternatives according to the provider validators above.

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/)
- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/)
