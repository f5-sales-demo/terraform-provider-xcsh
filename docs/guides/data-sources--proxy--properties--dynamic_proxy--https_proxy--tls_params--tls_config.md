---
page_title: "dynamic_proxy.https_proxy.tls_params.tls_config"
subcategory: ""
description: "dynamic_proxy.https_proxy.tls_params.tls_config for xcsh_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 2509, "body_sha256": "sha256:c1d83388aaa8073f04240c8456fb2b8d993804f4c1826e8211bbb9db0013ef42", "canonical_id": "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:https_proxy:tls_params:tls_config", "child_ids": ["xcsh-docs:data-sources:proxy:properties:dynamic_proxy:https_proxy:tls_params:tls_config:custom_security", "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:https_proxy:tls_params:tls_config:default_security", "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:https_proxy:tls_params:tls_config:low_security", "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:https_proxy:tls_params:tls_config:medium_security"], "collection_id": "xcsh-docs:data-sources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:https_proxy:tls_params:tls_config", "parent_id": "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:https_proxy:tls_params", "path": "docs/guides/data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params--tls_config.md", "provider_name": "proxy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["dynamic_proxy", "https_proxy", "tls_params", "tls_config"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/proxy/properties/dynamic_proxy/https_proxy/tls_params/tls_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "dynamic_proxy.https_proxy.tls_params.tls_config for xcsh_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# dynamic_proxy.https_proxy.tls_params.tls_config

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md)
- [Property reference](data-sources--proxy--reference.md)
- [dynamic_proxy](data-sources--proxy--properties--dynamic_proxy.md)
- [dynamic_proxy.https_proxy](data-sources--proxy--properties--dynamic_proxy--https_proxy.md)
- [dynamic_proxy.https_proxy.tls_params](data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params.md)
- dynamic_proxy.https_proxy.tls_params.tls_config

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

- [custom_security](data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params--tls_config--custom_security.md): complete subsection reference.

- [default_security](data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params--tls_config--default_security.md): complete subsection reference.

- [low_security](data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params--tls_config--low_security.md): complete subsection reference.

- [medium_security](data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params--tls_config--medium_security.md): complete subsection reference.

## Next pages

- [dynamic_proxy.https_proxy.tls_params.tls_config.custom_security](data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params--tls_config--custom_security.md)
- [dynamic_proxy.https_proxy.tls_params.tls_config.default_security](data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params--tls_config--default_security.md)
- [dynamic_proxy.https_proxy.tls_params.tls_config.low_security](data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params--tls_config--low_security.md)
- [dynamic_proxy.https_proxy.tls_params.tls_config.medium_security](data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params--tls_config--medium_security.md)
- [dynamic_proxy.https_proxy.tls_params](data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params.md)
- [xcsh_proxy](../data-sources/proxy.md)
