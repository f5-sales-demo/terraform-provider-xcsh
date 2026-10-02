---
page_title: "adobe_commerce_connector"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["adobe commerce connector"], "body_bytes": 2475, "body_sha256": "sha256:cc5055cd366991fcf30c84d0585a467020fa8d0364e47c252d8b243131e51896", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:resources:protected_application:properties:adobe_commerce_connector", "parent_id": "xcsh-docs:resources:protected_application:reference", "path": "documentation/resources/protected_application/properties/adobe_commerce_connector/index.md", "product": "distributed-cloud", "provider_name": "protected_application", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3332022203110112-1213232021331323-0313203313032202-0011122310122123-2101032101331003-1001330230023123-1032122322110301-1202320102202323", "registry_path": "docs/guides/resources--protected_application--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["adobe_commerce_connector"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protected_application/properties/adobe_commerce_connector/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# adobe_commerce_connector

Breadcrumbs:

- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/)
- adobe_commerce_connector

<a id="section"></a>

Type: `["object", {}]`. Optional.

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

- [adobe_commerce_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/adobe_commerce_connector/#section)
- [big_ip_iapp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/big_ip_iapp/#section)
- [cloudflare](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/#section)
- [cloudfront](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/#section)
- [custom_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/custom_connector/#section)
- [f5_big_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/f5_big_ip/#section)
- [salesforce_commerce_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/salesforce_commerce_connector/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
adobe_commerce_connector = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/)
- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/)
