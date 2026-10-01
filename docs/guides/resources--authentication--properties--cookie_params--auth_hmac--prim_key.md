---
page_title: "cookie_params.auth_hmac.prim_key"
subcategory: ""
description: "cookie_params.auth_hmac.prim_key for xcsh_authentication."
xcsh_docs: {"aliases": [], "body_bytes": 1991, "body_sha256": "sha256:ada7410953b018e37e8cf0b119ef0de7a2c35a8f73a88891f4d7836683bc0fa9", "canonical_id": "xcsh-docs:resources:authentication:properties:cookie_params:auth_hmac:prim_key", "child_ids": ["xcsh-docs:resources:authentication:properties:cookie_params:auth_hmac:prim_key:blindfold_secret_info", "xcsh-docs:resources:authentication:properties:cookie_params:auth_hmac:prim_key:clear_secret_info"], "collection_id": "xcsh-docs:resources:authentication:collection", "completeness": "complete", "id": "xcsh-docs:resources:authentication:properties:cookie_params:auth_hmac:prim_key", "parent_id": "xcsh-docs:resources:authentication:properties:cookie_params:auth_hmac", "path": "docs/guides/resources--authentication--properties--cookie_params--auth_hmac--prim_key.md", "provider_name": "authentication", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["cookie_params", "auth_hmac", "prim_key"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/authentication/properties/cookie_params/auth_hmac/prim_key/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cookie_params.auth_hmac.prim_key for xcsh_authentication.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["authenticationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cookie_params.auth_hmac.prim_key

Breadcrumbs:

- [xcsh_authentication](../resources/authentication.md)
- [Property reference](resources--authentication--reference.md)
- [cookie_params](resources--authentication--properties--cookie_params.md)
- [cookie_params.auth_hmac](resources--authentication--properties--cookie_params--auth_hmac.md)
- cookie_params.auth_hmac.prim_key

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
prim_key {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blindfold_secret_info](resources--authentication--properties--cookie_params--auth_hmac--prim_key--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](resources--authentication--properties--cookie_params--auth_hmac--prim_key--clear_secret_info.md): complete subsection reference.

## Next pages

- [cookie_params.auth_hmac.prim_key.blindfold_secret_info](resources--authentication--properties--cookie_params--auth_hmac--prim_key--blindfold_secret_info.md)
- [cookie_params.auth_hmac.prim_key.clear_secret_info](resources--authentication--properties--cookie_params--auth_hmac--prim_key--clear_secret_info.md)
- [cookie_params.auth_hmac](resources--authentication--properties--cookie_params--auth_hmac.md)
- [xcsh_authentication](../resources/authentication.md)
