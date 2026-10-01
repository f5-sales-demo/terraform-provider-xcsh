---
page_title: "use_tls.use_mtls.tls_certificates.private_key"
subcategory: "Load Balancing"
description: "use_tls.use_mtls.tls_certificates.private_key for xcsh_origin_pool."
xcsh_docs: {"aliases": [], "body_bytes": 2175, "body_sha256": "sha256:37b2b4710530af1dc9195e1919350a95926d0e34b31f9637789b49788634288b", "canonical_id": "xcsh-docs:resources:origin_pool:properties:use_tls:use_mtls:tls_certificates:private_key", "child_ids": ["xcsh-docs:resources:origin_pool:properties:use_tls:use_mtls:tls_certificates:private_key:blindfold_secret_info", "xcsh-docs:resources:origin_pool:properties:use_tls:use_mtls:tls_certificates:private_key:clear_secret_info"], "collection_id": "xcsh-docs:resources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:resources:origin_pool:properties:use_tls:use_mtls:tls_certificates:private_key", "parent_id": "xcsh-docs:resources:origin_pool:properties:use_tls:use_mtls:tls_certificates", "path": "docs/guides/resources--origin_pool--properties--use_tls--use_mtls--tls_certificates--private_key.md", "provider_name": "origin_pool", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["use_tls", "use_mtls", "tls_certificates", "private_key"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/origin_pool/properties/use_tls/use_mtls/tls_certificates/private_key/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "use_tls.use_mtls.tls_certificates.private_key for xcsh_origin_pool.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# use_tls.use_mtls.tls_certificates.private_key

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md)
- [Property reference](resources--origin_pool--reference.md)
- [use_tls](resources--origin_pool--properties--use_tls.md)
- [use_tls.use_mtls](resources--origin_pool--properties--use_tls--use_mtls.md)
- [use_tls.use_mtls.tls_certificates](resources--origin_pool--properties--use_tls--use_mtls--tls_certificates.md)
- use_tls.use_mtls.tls_certificates.private_key

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

- [blindfold_secret_info](resources--origin_pool--properties--use_tls--use_mtls--tls_certificates--private_key--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](resources--origin_pool--properties--use_tls--use_mtls--tls_certificates--private_key--clear_secret_info.md): complete subsection reference.

## Next pages

- [use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info](resources--origin_pool--properties--use_tls--use_mtls--tls_certificates--private_key--blindfold_secret_info.md)
- [use_tls.use_mtls.tls_certificates.private_key.clear_secret_info](resources--origin_pool--properties--use_tls--use_mtls--tls_certificates--private_key--clear_secret_info.md)
- [use_tls.use_mtls.tls_certificates](resources--origin_pool--properties--use_tls--use_mtls--tls_certificates.md)
- [xcsh_origin_pool](../resources/origin_pool.md)
