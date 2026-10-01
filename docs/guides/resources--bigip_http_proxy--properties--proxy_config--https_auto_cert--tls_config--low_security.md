---
page_title: "proxy_config.https_auto_cert.tls_config.low_security"
subcategory: ""
description: "proxy_config.https_auto_cert.tls_config.low_security for xcsh_bigip_http_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 1354, "body_sha256": "sha256:cf417a1219481cbf90f3038dafcb7f29accb1d0e4967055ba5c39bee21bc04ec", "canonical_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https_auto_cert:tls_config:low_security", "child_ids": [], "collection_id": "xcsh-docs:resources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https_auto_cert:tls_config:low_security", "parent_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https_auto_cert:tls_config", "path": "docs/guides/resources--bigip_http_proxy--properties--proxy_config--https_auto_cert--tls_config--low_security.md", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["proxy_config", "https_auto_cert", "tls_config", "low_security"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bigip_http_proxy/properties/proxy_config/https_auto_cert/tls_config/low_security/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "proxy_config.https_auto_cert.tls_config.low_security for xcsh_bigip_http_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# proxy_config.https_auto_cert.tls_config.low_security

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md)
- [Property reference](resources--bigip_http_proxy--reference.md)
- [proxy_config](resources--bigip_http_proxy--properties--proxy_config.md)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--properties--proxy_config--https_auto_cert.md)
- [proxy_config.https_auto_cert.tls_config](resources--bigip_http_proxy--properties--proxy_config--https_auto_cert--tls_config.md)
- proxy_config.https_auto_cert.tls_config.low_security

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
low_security = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [proxy_config.https_auto_cert.tls_config](resources--bigip_http_proxy--properties--proxy_config--https_auto_cert--tls_config.md)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md)
