---
page_title: "limits.disabled"
subcategory: "Security"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["limits disabled"], "body_bytes": 1186, "body_sha256": "sha256:e61f3371c2155faefd5be2415fa6e6697e4786bd08bd195f9327fbfb3aba6fb2", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["receipt-pinned-upstream"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:rate_limiter:collection", "completeness": "complete", "id": "xcsh-docs:resources:rate_limiter:properties:limits:disabled", "parent_id": "xcsh-docs:resources:rate_limiter:properties:limits", "path": "documentation/resources/rate_limiter/properties/limits/disabled/index.md", "product": "distributed-cloud", "provider_name": "rate_limiter", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2100333132122033-3311323310131120-0232112103120101-0232121211203013-2132200113113312-1323303211112203-3100011121020032-2231102302303300", "registry_path": "docs/guides/resources--rate_limiter--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["limits", "disabled"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/rate_limiter/properties/limits/disabled/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["rate_limiterCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# limits.disabled

Breadcrumbs:

- [xcsh_rate_limiter](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/properties/)
- [limits](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/properties/limits/)
- limits.disabled

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

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

Terraform syntax:

```terraform
disabled = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [limits](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/properties/limits/)
- [xcsh_rate_limiter](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/)
