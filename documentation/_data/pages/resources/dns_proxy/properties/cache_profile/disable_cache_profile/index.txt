---
page_title: "cache_profile.disable_cache_profile"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["cache profile disable cache profile"], "body_bytes": 1029, "body_sha256": "sha256:d0078611a3be83e1bf72f2a2744278cba857e8f0c9c24b79844884247597429a", "capabilities": ["dns"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:dns_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_proxy:properties:cache_profile:disable_cache_profile", "parent_id": "xcsh-docs:resources:dns_proxy:properties:cache_profile", "path": "documentation/resources/dns_proxy/properties/cache_profile/disable_cache_profile/index.md", "product": "distributed-cloud", "provider_name": "dns_proxy", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-0212030122111310-0001322001032333-3111220300130012-1112202203333311-3001120122013032-2311032113110123-1110302030132313-0301000310012230", "registry_path": "docs/guides/resources--dns_proxy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["cache_profile", "disable_cache_profile"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_proxy/properties/cache_profile/disable_cache_profile/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["dns_proxyCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
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

Terraform syntax:

```terraform
disable_cache_profile = {}
```

This is an empty object or choice marker. It has no direct properties.
