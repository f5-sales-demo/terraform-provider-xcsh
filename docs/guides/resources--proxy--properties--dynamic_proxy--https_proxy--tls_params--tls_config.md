---
page_title: "dynamic_proxy.https_proxy.tls_params.tls_config"
subcategory: ""
description: "dynamic_proxy.https_proxy.tls_params.tls_config for xcsh_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 3177, "body_sha256": "sha256:88824f57ffa616fa0a9a66912e3e94311da02d0244e9ead8b8899625b66f5ca4", "canonical_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:tls_params:tls_config", "child_ids": ["xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:tls_params:tls_config:custom_security", "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:tls_params:tls_config:default_security", "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:tls_params:tls_config:low_security", "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:tls_params:tls_config:medium_security"], "collection_id": "xcsh-docs:resources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:tls_params:tls_config", "parent_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:tls_params", "path": "docs/guides/resources--proxy--properties--dynamic_proxy--https_proxy--tls_params--tls_config.md", "provider_name": "proxy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["dynamic_proxy", "https_proxy", "tls_params", "tls_config"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/proxy/properties/dynamic_proxy/https_proxy/tls_params/tls_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "dynamic_proxy.https_proxy.tls_params.tls_config for xcsh_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# dynamic_proxy.https_proxy.tls_params.tls_config

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md)
- [Property reference](resources--proxy--reference.md)
- [dynamic_proxy](resources--proxy--properties--dynamic_proxy.md)
- [dynamic_proxy.https_proxy](resources--proxy--properties--dynamic_proxy--https_proxy.md)
- [dynamic_proxy.https_proxy.tls_params](resources--proxy--properties--dynamic_proxy--https_proxy--tls_params.md)
- dynamic_proxy.https_proxy.tls_params.tls_config

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

- [custom_security](resources--proxy--properties--dynamic_proxy--https_proxy--tls_params--tls_config--custom_security.md): complete subsection reference.

- [default_security](resources--proxy--properties--dynamic_proxy--https_proxy--tls_params--tls_config--default_security.md): complete subsection reference.

- [low_security](resources--proxy--properties--dynamic_proxy--https_proxy--tls_params--tls_config--low_security.md): complete subsection reference.

- [medium_security](resources--proxy--properties--dynamic_proxy--https_proxy--tls_params--tls_config--medium_security.md): complete subsection reference.

## Next pages

- [dynamic_proxy.https_proxy.tls_params.tls_config.custom_security](resources--proxy--properties--dynamic_proxy--https_proxy--tls_params--tls_config--custom_security.md)
- [dynamic_proxy.https_proxy.tls_params.tls_config.default_security](resources--proxy--properties--dynamic_proxy--https_proxy--tls_params--tls_config--default_security.md)
- [dynamic_proxy.https_proxy.tls_params.tls_config.low_security](resources--proxy--properties--dynamic_proxy--https_proxy--tls_params--tls_config--low_security.md)
- [dynamic_proxy.https_proxy.tls_params.tls_config.medium_security](resources--proxy--properties--dynamic_proxy--https_proxy--tls_params--tls_config--medium_security.md)
- [dynamic_proxy.https_proxy.tls_params](resources--proxy--properties--dynamic_proxy--https_proxy--tls_params.md)
- [xcsh_proxy](../resources/proxy.md)
