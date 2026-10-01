---
page_title: "proxy_config.https.coalescing_options"
subcategory: ""
description: "proxy_config.https.coalescing_options for xcsh_bigip_http_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 1826, "body_sha256": "sha256:b6e314d13f84659bde27443be268de46b33dfcd52002c00b5f9bd8d9f9efff3a", "canonical_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https:coalescing_options", "child_ids": ["xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https:coalescing_options:default_coalescing", "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https:coalescing_options:strict_coalescing"], "collection_id": "xcsh-docs:data-sources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https:coalescing_options", "parent_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https", "path": "docs/guides/data-sources--bigip_http_proxy--properties--proxy_config--https--coalescing_options.md", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["proxy_config", "https", "coalescing_options"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bigip_http_proxy/properties/proxy_config/https/coalescing_options/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "proxy_config.https.coalescing_options for xcsh_bigip_http_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# proxy_config.https.coalescing_options

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md)
- [Property reference](data-sources--bigip_http_proxy--reference.md)
- [proxy_config](data-sources--bigip_http_proxy--properties--proxy_config.md)
- [proxy_config.https](data-sources--bigip_http_proxy--properties--proxy_config--https.md)
- proxy_config.https.coalescing_options

<a id="section"></a>

Type: `"single"`. Computed.

TLS connection coalescing configuration (not compatible with mTLS).

Upstream description:

TLS connection coalescing configuration (not compatible with mTLS)

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-coalescing_choice": "[\"default_coalescing\",\"strict_coalescing\"]"
}
```

## Direct properties

- [default_coalescing](data-sources--bigip_http_proxy--properties--proxy_config--https--coalescing_options--default_coalescing.md): complete subsection reference.

- [strict_coalescing](data-sources--bigip_http_proxy--properties--proxy_config--https--coalescing_options--strict_coalescing.md): complete subsection reference.

## Next pages

- [proxy_config.https.coalescing_options.default_coalescing](data-sources--bigip_http_proxy--properties--proxy_config--https--coalescing_options--default_coalescing.md)
- [proxy_config.https.coalescing_options.strict_coalescing](data-sources--bigip_http_proxy--properties--proxy_config--https--coalescing_options--strict_coalescing.md)
- [proxy_config.https](data-sources--bigip_http_proxy--properties--proxy_config--https.md)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md)
