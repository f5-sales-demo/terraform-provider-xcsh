---
page_title: "proxy_config.https_auto_cert.coalescing_options"
subcategory: ""
description: "proxy_config.https_auto_cert.coalescing_options for xcsh_bigip_http_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 1847, "body_sha256": "sha256:e18c6ed679b00b69dc508dfbd5ed641786eb5ff652b2a30a7bfc826ba9d8f5e2", "canonical_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https_auto_cert:coalescing_options", "child_ids": ["xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https_auto_cert:coalescing_options:default_coalescing", "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https_auto_cert:coalescing_options:strict_coalescing"], "collection_id": "xcsh-docs:data-sources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https_auto_cert:coalescing_options", "parent_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https_auto_cert", "path": "docs/guides/data-sources--bigip_http_proxy--properties--proxy_config--https_auto_cert--coalescing_options.md", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["proxy_config", "https_auto_cert", "coalescing_options"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/coalescing_options/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "proxy_config.https_auto_cert.coalescing_options for xcsh_bigip_http_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# proxy_config.https_auto_cert.coalescing_options

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md)
- [Property reference](data-sources--bigip_http_proxy--reference.md)
- [proxy_config](data-sources--bigip_http_proxy--properties--proxy_config.md)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--properties--proxy_config--https_auto_cert.md)
- proxy_config.https_auto_cert.coalescing_options

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

- [default_coalescing](data-sources--bigip_http_proxy--properties--proxy_config--https_auto_cert--coalescing_options--default_coalescing.md): complete subsection reference.

- [strict_coalescing](data-sources--bigip_http_proxy--properties--proxy_config--https_auto_cert--coalescing_options--strict_coalescing.md): complete subsection reference.

## Next pages

- [proxy_config.https_auto_cert.coalescing_options.default_coalescing](data-sources--bigip_http_proxy--properties--proxy_config--https_auto_cert--coalescing_options--default_coalescing.md)
- [proxy_config.https_auto_cert.coalescing_options.strict_coalescing](data-sources--bigip_http_proxy--properties--proxy_config--https_auto_cert--coalescing_options--strict_coalescing.md)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--properties--proxy_config--https_auto_cert.md)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md)
