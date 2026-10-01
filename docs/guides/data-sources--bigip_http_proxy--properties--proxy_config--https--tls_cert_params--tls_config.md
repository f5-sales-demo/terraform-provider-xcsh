---
page_title: "proxy_config.https.tls_cert_params.tls_config"
subcategory: ""
description: "proxy_config.https.tls_cert_params.tls_config for xcsh_bigip_http_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 2743, "body_sha256": "sha256:a7cc7ebfb197eb8e7b623de82892822da38cb46856691b8ffde6521ba5a15009", "canonical_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params:tls_config", "child_ids": ["xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params:tls_config:custom_security", "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params:tls_config:default_security", "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params:tls_config:low_security", "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params:tls_config:medium_security"], "collection_id": "xcsh-docs:data-sources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params:tls_config", "parent_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params", "path": "docs/guides/data-sources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params--tls_config.md", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["proxy_config", "https", "tls_cert_params", "tls_config"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_cert_params/tls_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "proxy_config.https.tls_cert_params.tls_config for xcsh_bigip_http_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# proxy_config.https.tls_cert_params.tls_config

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md)
- [Property reference](data-sources--bigip_http_proxy--reference.md)
- [proxy_config](data-sources--bigip_http_proxy--properties--proxy_config.md)
- [proxy_config.https](data-sources--bigip_http_proxy--properties--proxy_config--https.md)
- [proxy_config.https.tls_cert_params](data-sources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params.md)
- proxy_config.https.tls_cert_params.tls_config

<a id="section"></a>

Type: `"single"`. Computed.

Defines various OPTIONS to configure TLS configuration parameters.

Upstream description:

This defines various OPTIONS to configure TLS configuration parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"custom_security\",\"default_security\",\"low_security\",\"medium_security\"]"
}
```

## Direct properties

- [custom_security](data-sources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params--tls_config--custom_security.md): complete subsection reference.

- [default_security](data-sources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params--tls_config--default_security.md): complete subsection reference.

- [low_security](data-sources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params--tls_config--low_security.md): complete subsection reference.

- [medium_security](data-sources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params--tls_config--medium_security.md): complete subsection reference.

## Next pages

- [proxy_config.https.tls_cert_params.tls_config.custom_security](data-sources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params--tls_config--custom_security.md)
- [proxy_config.https.tls_cert_params.tls_config.default_security](data-sources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params--tls_config--default_security.md)
- [proxy_config.https.tls_cert_params.tls_config.low_security](data-sources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params--tls_config--low_security.md)
- [proxy_config.https.tls_cert_params.tls_config.medium_security](data-sources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params--tls_config--medium_security.md)
- [proxy_config.https.tls_cert_params](data-sources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params.md)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md)
