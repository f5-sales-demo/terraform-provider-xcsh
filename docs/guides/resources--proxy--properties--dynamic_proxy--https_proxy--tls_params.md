---
page_title: "dynamic_proxy.https_proxy.tls_params"
subcategory: ""
description: "dynamic_proxy.https_proxy.tls_params for xcsh_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 2421, "body_sha256": "sha256:d8597afc534f1f06b73383609064bbb79bb57cee469e777d06f11f1a261d83f5", "canonical_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:tls_params", "child_ids": ["xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:tls_params:no_mtls", "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:tls_params:tls_certificates", "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:tls_params:tls_config", "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:tls_params:use_mtls"], "collection_id": "xcsh-docs:resources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:tls_params", "parent_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy", "path": "docs/guides/resources--proxy--properties--dynamic_proxy--https_proxy--tls_params.md", "provider_name": "proxy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["dynamic_proxy", "https_proxy", "tls_params"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/proxy/properties/dynamic_proxy/https_proxy/tls_params/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "dynamic_proxy.https_proxy.tls_params for xcsh_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# dynamic_proxy.https_proxy.tls_params

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md)
- [Property reference](resources--proxy--reference.md)
- [dynamic_proxy](resources--proxy--properties--dynamic_proxy.md)
- [dynamic_proxy.https_proxy](resources--proxy--properties--dynamic_proxy--https_proxy.md)
- dynamic_proxy.https_proxy.tls_params

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Inline TLS Parameters. Inline TLS parameters.

Upstream description:

Inline TLS parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("tls_certificates"),
  validators.ConflictingObjectAttributes("no_mtls",
    "use_mtls")}
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
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]"
}
```

Terraform syntax:

```terraform
tls_params {
  # Configure direct properties listed below.
}
```

## Direct properties

- [no_mtls](resources--proxy--properties--dynamic_proxy--https_proxy--tls_params--no_mtls.md): complete subsection reference.

- [tls_certificates](resources--proxy--properties--dynamic_proxy--https_proxy--tls_params--tls_certificates.md): complete subsection reference.

- [tls_config](resources--proxy--properties--dynamic_proxy--https_proxy--tls_params--tls_config.md): complete subsection reference.

- [use_mtls](resources--proxy--properties--dynamic_proxy--https_proxy--tls_params--use_mtls.md): complete subsection reference.

## Next pages

- [dynamic_proxy.https_proxy.tls_params.no_mtls](resources--proxy--properties--dynamic_proxy--https_proxy--tls_params--no_mtls.md)
- [dynamic_proxy.https_proxy.tls_params.tls_certificates](resources--proxy--properties--dynamic_proxy--https_proxy--tls_params--tls_certificates.md)
- [dynamic_proxy.https_proxy.tls_params.tls_config](resources--proxy--properties--dynamic_proxy--https_proxy--tls_params--tls_config.md)
- [dynamic_proxy.https_proxy.tls_params.use_mtls](resources--proxy--properties--dynamic_proxy--https_proxy--tls_params--use_mtls.md)
- [dynamic_proxy.https_proxy](resources--proxy--properties--dynamic_proxy--https_proxy.md)
- [xcsh_proxy](../resources/proxy.md)
