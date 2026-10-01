---
page_title: "cookie_params.auth_hmac.sec_key"
subcategory: ""
description: "cookie_params.auth_hmac.sec_key for xcsh_authentication."
xcsh_docs: {"aliases": [], "body_bytes": 1982, "body_sha256": "sha256:75ed42e08a1555dfb0f94079407de6a3462459bdd396d9d841df066321b3d89d", "canonical_id": "xcsh-docs:resources:authentication:properties:cookie_params:auth_hmac:sec_key", "child_ids": ["xcsh-docs:resources:authentication:properties:cookie_params:auth_hmac:sec_key:blindfold_secret_info", "xcsh-docs:resources:authentication:properties:cookie_params:auth_hmac:sec_key:clear_secret_info"], "collection_id": "xcsh-docs:resources:authentication:collection", "completeness": "complete", "id": "xcsh-docs:resources:authentication:properties:cookie_params:auth_hmac:sec_key", "parent_id": "xcsh-docs:resources:authentication:properties:cookie_params:auth_hmac", "path": "docs/guides/resources--authentication--properties--cookie_params--auth_hmac--sec_key.md", "provider_name": "authentication", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["cookie_params", "auth_hmac", "sec_key"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/authentication/properties/cookie_params/auth_hmac/sec_key/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cookie_params.auth_hmac.sec_key for xcsh_authentication.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["authenticationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cookie_params.auth_hmac.sec_key

Breadcrumbs:

- [xcsh_authentication](../resources/authentication.md)
- [Property reference](resources--authentication--reference.md)
- [cookie_params](resources--authentication--properties--cookie_params.md)
- [cookie_params.auth_hmac](resources--authentication--properties--cookie_params--auth_hmac.md)
- cookie_params.auth_hmac.sec_key

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
sec_key {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blindfold_secret_info](resources--authentication--properties--cookie_params--auth_hmac--sec_key--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](resources--authentication--properties--cookie_params--auth_hmac--sec_key--clear_secret_info.md): complete subsection reference.

## Next pages

- [cookie_params.auth_hmac.sec_key.blindfold_secret_info](resources--authentication--properties--cookie_params--auth_hmac--sec_key--blindfold_secret_info.md)
- [cookie_params.auth_hmac.sec_key.clear_secret_info](resources--authentication--properties--cookie_params--auth_hmac--sec_key--clear_secret_info.md)
- [cookie_params.auth_hmac](resources--authentication--properties--cookie_params--auth_hmac.md)
- [xcsh_authentication](../resources/authentication.md)
