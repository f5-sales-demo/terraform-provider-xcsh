---
page_title: "proxy_config.https.tls_cert_params"
subcategory: ""
description: "proxy_config.https.tls_cert_params for xcsh_bigip_http_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 2540, "body_sha256": "sha256:dd4baa8912ba4f1b9ac0b3ecc0ed0156fa94f8f0b9459ef25a06e7c2fdb36692", "canonical_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params", "child_ids": ["xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params:certificates", "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params:no_mtls", "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params:tls_config", "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params:use_mtls"], "collection_id": "xcsh-docs:resources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params", "parent_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https", "path": "docs/guides/resources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params.md", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["proxy_config", "https", "tls_cert_params"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bigip_http_proxy/properties/proxy_config/https/tls_cert_params/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "proxy_config.https.tls_cert_params for xcsh_bigip_http_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# proxy_config.https.tls_cert_params

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md)
- [Property reference](resources--bigip_http_proxy--reference.md)
- [proxy_config](resources--bigip_http_proxy--properties--proxy_config.md)
- [proxy_config.https](resources--bigip_http_proxy--properties--proxy_config--https.md)
- proxy_config.https.tls_cert_params

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for tls cert params.

Upstream description:

Select TLS Parameters and Certificates.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("certificates"),
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
tls_cert_params {
  # Configure direct properties listed below.
}
```

## Direct properties

- [certificates](resources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params--certificates.md): complete subsection reference.

- [no_mtls](resources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params--no_mtls.md): complete subsection reference.

- [tls_config](resources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params--tls_config.md): complete subsection reference.

- [use_mtls](resources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params--use_mtls.md): complete subsection reference.

## Next pages

- [proxy_config.https.tls_cert_params.certificates](resources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params--certificates.md)
- [proxy_config.https.tls_cert_params.no_mtls](resources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params--no_mtls.md)
- [proxy_config.https.tls_cert_params.tls_config](resources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params--tls_config.md)
- [proxy_config.https.tls_cert_params.use_mtls](resources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params--use_mtls.md)
- [proxy_config.https](resources--bigip_http_proxy--properties--proxy_config--https.md)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md)
