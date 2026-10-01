---
page_title: "cache_profile"
subcategory: ""
description: "cache_profile for xcsh_dns_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 2270, "body_sha256": "sha256:ca5126e28393882759de6e75dc2f4dce5838426c7a6a534dc16e1d584f227841", "child_ids": ["xcsh-docs:data-sources:dns_proxy:properties:cache_profile:disable_cache_profile"], "collection_id": "xcsh-docs:data-sources:dns_proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_proxy:properties:cache_profile", "parent_id": "xcsh-docs:data-sources:dns_proxy:reference", "path": "documentation/data-sources/dns_proxy/properties/cache_profile/index.md", "provider_name": "dns_proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "properties", "schema_path": ["cache_profile"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_proxy/properties/cache_profile/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cache_profile for xcsh_dns_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cache_profile

Breadcrumbs:

- [xcsh_dns_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/)
- cache_profile

<a id="section"></a>

Type: `"single"`. Computed.

DNS Cache specifies cache configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-cache_profile_choice": "[\"cache_size\",\"disable_cache_profile\"]"
}
```

## Direct properties

<a id="schema-cache_profile--cache_size"></a>

### cache_size property

Type: `"number"`. Computed.

Exclusive with \[disable\_cache\_profile\] cache size.

Upstream description:

Exclusive with \[disable\_cache\_profile\] cache size.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 10240,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "10240"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "10240"
  }
}
```

- [disable_cache_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/cache_profile/disable_cache_profile/): complete subsection reference.

## Next pages

- [cache_profile.disable_cache_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/cache_profile/disable_cache_profile/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/)
- [xcsh_dns_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/)
