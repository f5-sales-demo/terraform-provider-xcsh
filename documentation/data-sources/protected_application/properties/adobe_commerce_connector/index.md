---
page_title: "adobe_commerce_connector"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["adobe commerce connector"], "body_bytes": 2159, "body_sha256": "sha256:1d47094bea706b6558526f4ad25aa542f331bd68a218c584d9f478f7a8833bc9", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:protected_application:properties:adobe_commerce_connector", "parent_id": "xcsh-docs:data-sources:protected_application:reference", "path": "documentation/data-sources/protected_application/properties/adobe_commerce_connector/index.md", "product": "distributed-cloud", "provider_name": "protected_application", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-2131122002322200-1211102013312210-0332313001322211-3203310002112222-1102230310021123-1302303130013203-0210310220203232-1201203310213203", "registry_path": "docs/guides/data-sources--protected_application--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["adobe_commerce_connector"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protected_application/properties/adobe_commerce_connector/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
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

Additional upstream details:

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

This is an empty object or choice marker. It has no direct properties.
