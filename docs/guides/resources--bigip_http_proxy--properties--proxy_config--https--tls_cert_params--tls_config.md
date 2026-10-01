---
page_title: "proxy_config.https.tls_cert_params.tls_config"
subcategory: ""
description: "proxy_config.https.tls_cert_params.tls_config for xcsh_bigip_http_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 3411, "body_sha256": "sha256:f9a1db7a0fd4b0d6686e8d8432daba88ef5a97569d22c40f0e6a1e8379760fcd", "canonical_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params:tls_config", "child_ids": ["xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params:tls_config:custom_security", "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params:tls_config:default_security", "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params:tls_config:low_security", "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params:tls_config:medium_security"], "collection_id": "xcsh-docs:resources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params:tls_config", "parent_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params", "path": "docs/guides/resources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params--tls_config.md", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["proxy_config", "https", "tls_cert_params", "tls_config"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bigip_http_proxy/properties/proxy_config/https/tls_cert_params/tls_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "proxy_config.https.tls_cert_params.tls_config for xcsh_bigip_http_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# proxy_config.https.tls_cert_params.tls_config

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md)
- [Property reference](resources--bigip_http_proxy--reference.md)
- [proxy_config](resources--bigip_http_proxy--properties--proxy_config.md)
- [proxy_config.https](resources--bigip_http_proxy--properties--proxy_config--https.md)
- [proxy_config.https.tls_cert_params](resources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params.md)
- proxy_config.https.tls_cert_params.tls_config

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Defines various OPTIONS to configure TLS configuration parameters.

Upstream description:

This defines various OPTIONS to configure TLS configuration parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_security",
    "default_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "low_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("default_security",
    "low_security"),
  validators.ConflictingObjectAttributes("default_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("low_security",
    "medium_security")}
```

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

Terraform syntax:

```terraform
tls_config {
  # Configure direct properties listed below.
}
```

## Direct properties

- [custom_security](resources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params--tls_config--custom_security.md): complete subsection reference.

- [default_security](resources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params--tls_config--default_security.md): complete subsection reference.

- [low_security](resources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params--tls_config--low_security.md): complete subsection reference.

- [medium_security](resources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params--tls_config--medium_security.md): complete subsection reference.

## Next pages

- [proxy_config.https.tls_cert_params.tls_config.custom_security](resources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params--tls_config--custom_security.md)
- [proxy_config.https.tls_cert_params.tls_config.default_security](resources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params--tls_config--default_security.md)
- [proxy_config.https.tls_cert_params.tls_config.low_security](resources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params--tls_config--low_security.md)
- [proxy_config.https.tls_cert_params.tls_config.medium_security](resources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params--tls_config--medium_security.md)
- [proxy_config.https.tls_cert_params](resources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params.md)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md)
