---
page_title: "cache_profile.disable_cache_profile"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["cache profile disable cache profile"], "body_bytes": 1278, "body_sha256": "sha256:0258855d0d0cab7a9a06885ed84f23c38a6e0acad3545da3f1a067caae82e7a0", "capabilities": ["dns"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:dns_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_proxy:properties:cache_profile:disable_cache_profile", "parent_id": "xcsh-docs:resources:dns_proxy:properties:cache_profile", "path": "documentation/resources/dns_proxy/properties/cache_profile/disable_cache_profile/index.md", "product": "distributed-cloud", "provider_name": "dns_proxy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0212030122111310-0001322001032333-3111220300130012-1112202203333311-3001120122013032-2311032113110123-1110302030132313-0301000310012230", "registry_path": "docs/guides/resources--dns_proxy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["cache_profile", "disable_cache_profile"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_proxy/properties/cache_profile/disable_cache_profile/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["dns_proxyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cache_profile.disable_cache_profile

Breadcrumbs:

- [xcsh_dns_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/)
- [cache_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/cache_profile/)
- cache_profile.disable_cache_profile

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable cache profile.

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
disable_cache_profile = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [cache_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/cache_profile/)
- [xcsh_dns_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/)
