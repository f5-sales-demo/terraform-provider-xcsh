---
page_title: "proxy_config.https.tls_parameters.tls_config"
subcategory: ""
description: "proxy_config.https.tls_parameters.tls_config for xcsh_bigip_http_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 3393, "body_sha256": "sha256:1992e0376510155591f3d3591bef14294f9a020ff21e7e7b215a02e430dc22f0", "canonical_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_parameters:tls_config", "child_ids": ["xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_parameters:tls_config:custom_security", "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_parameters:tls_config:default_security", "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_parameters:tls_config:low_security", "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_parameters:tls_config:medium_security"], "collection_id": "xcsh-docs:resources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_parameters:tls_config", "parent_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_parameters", "path": "docs/guides/resources--bigip_http_proxy--properties--proxy_config--https--tls_parameters--tls_config.md", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["proxy_config", "https", "tls_parameters", "tls_config"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bigip_http_proxy/properties/proxy_config/https/tls_parameters/tls_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "proxy_config.https.tls_parameters.tls_config for xcsh_bigip_http_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# proxy_config.https.tls_parameters.tls_config

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md)
- [Property reference](resources--bigip_http_proxy--reference.md)
- [proxy_config](resources--bigip_http_proxy--properties--proxy_config.md)
- [proxy_config.https](resources--bigip_http_proxy--properties--proxy_config--https.md)
- [proxy_config.https.tls_parameters](resources--bigip_http_proxy--properties--proxy_config--https--tls_parameters.md)
- proxy_config.https.tls_parameters.tls_config

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

- [custom_security](resources--bigip_http_proxy--properties--proxy_config--https--tls_parameters--tls_config--custom_security.md): complete subsection reference.

- [default_security](resources--bigip_http_proxy--properties--proxy_config--https--tls_parameters--tls_config--default_security.md): complete subsection reference.

- [low_security](resources--bigip_http_proxy--properties--proxy_config--https--tls_parameters--tls_config--low_security.md): complete subsection reference.

- [medium_security](resources--bigip_http_proxy--properties--proxy_config--https--tls_parameters--tls_config--medium_security.md): complete subsection reference.

## Next pages

- [proxy_config.https.tls_parameters.tls_config.custom_security](resources--bigip_http_proxy--properties--proxy_config--https--tls_parameters--tls_config--custom_security.md)
- [proxy_config.https.tls_parameters.tls_config.default_security](resources--bigip_http_proxy--properties--proxy_config--https--tls_parameters--tls_config--default_security.md)
- [proxy_config.https.tls_parameters.tls_config.low_security](resources--bigip_http_proxy--properties--proxy_config--https--tls_parameters--tls_config--low_security.md)
- [proxy_config.https.tls_parameters.tls_config.medium_security](resources--bigip_http_proxy--properties--proxy_config--https--tls_parameters--tls_config--medium_security.md)
- [proxy_config.https.tls_parameters](resources--bigip_http_proxy--properties--proxy_config--https--tls_parameters.md)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md)
