---
page_title: "dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key"
subcategory: ""
description: "dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key for xcsh_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 2487, "body_sha256": "sha256:f5e75e195727af29da08ca946dcfa3c7333f70eb9b240babc636bed03e9a3749", "canonical_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:tls_params:tls_certificates:private_key", "child_ids": ["xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:tls_params:tls_certificates:private_key:blindfold_secret_info", "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:tls_params:tls_certificates:private_key:clear_secret_info"], "collection_id": "xcsh-docs:resources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:tls_params:tls_certificates:private_key", "parent_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:tls_params:tls_certificates", "path": "docs/guides/resources--proxy--properties--dynamic_proxy--https_proxy--tls_params--tls_certificates--private_key.md", "provider_name": "proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["dynamic_proxy", "https_proxy", "tls_params", "tls_certificates", "private_key"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/proxy/properties/dynamic_proxy/https_proxy/tls_params/tls_certificates/private_key/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key for xcsh_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md)
- [Property reference](resources--proxy--reference.md)
- [dynamic_proxy](resources--proxy--properties--dynamic_proxy.md)
- [dynamic_proxy.https_proxy](resources--proxy--properties--dynamic_proxy--https_proxy.md)
- [dynamic_proxy.https_proxy.tls_params](resources--proxy--properties--dynamic_proxy--https_proxy--tls_params.md)
- [dynamic_proxy.https_proxy.tls_params.tls_certificates](resources--proxy--properties--dynamic_proxy--https_proxy--tls_params--tls_certificates.md)
- dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
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
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
private_key {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blindfold_secret_info](resources--proxy--properties--dynamic_proxy--https_proxy--tls_params--tls_certificates--private_key--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](resources--proxy--properties--dynamic_proxy--https_proxy--tls_params--tls_certificates--private_key--clear_secret_info.md): complete subsection reference.

## Next pages

- [dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.blindfold_secret_info](resources--proxy--properties--dynamic_proxy--https_proxy--tls_params--tls_certificates--private_key--blindfold_secret_info.md)
- [dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.clear_secret_info](resources--proxy--properties--dynamic_proxy--https_proxy--tls_params--tls_certificates--private_key--clear_secret_info.md)
- [dynamic_proxy.https_proxy.tls_params.tls_certificates](resources--proxy--properties--dynamic_proxy--https_proxy--tls_params--tls_certificates.md)
- [xcsh_proxy](../resources/proxy.md)
