---
page_title: "authentication.cookie_params.auth_hmac.prim_key"
subcategory: ""
description: "authentication.cookie_params.auth_hmac.prim_key for xcsh_virtual_host."
xcsh_docs: {"aliases": [], "body_bytes": 2259, "body_sha256": "sha256:3de59640cddbbe64e747ca8f68a4c0ae3f0404d3a0bf9a301e4443dbf2a11910", "canonical_id": "xcsh-docs:resources:virtual_host:properties:authentication:cookie_params:auth_hmac:prim_key", "child_ids": ["xcsh-docs:resources:virtual_host:properties:authentication:cookie_params:auth_hmac:prim_key:blindfold_secret_info", "xcsh-docs:resources:virtual_host:properties:authentication:cookie_params:auth_hmac:prim_key:clear_secret_info"], "collection_id": "xcsh-docs:resources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_host:properties:authentication:cookie_params:auth_hmac:prim_key", "parent_id": "xcsh-docs:resources:virtual_host:properties:authentication:cookie_params:auth_hmac", "path": "docs/guides/resources--virtual_host--properties--authentication--cookie_params--auth_hmac--prim_key.md", "provider_name": "virtual_host", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["authentication", "cookie_params", "auth_hmac", "prim_key"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_host/properties/authentication/cookie_params/auth_hmac/prim_key/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "authentication.cookie_params.auth_hmac.prim_key for xcsh_virtual_host.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# authentication.cookie_params.auth_hmac.prim_key

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md)
- [Property reference](resources--virtual_host--reference.md)
- [authentication](resources--virtual_host--properties--authentication.md)
- [authentication.cookie_params](resources--virtual_host--properties--authentication--cookie_params.md)
- [authentication.cookie_params.auth_hmac](resources--virtual_host--properties--authentication--cookie_params--auth_hmac.md)
- authentication.cookie_params.auth_hmac.prim_key

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

- [blindfold_secret_info](resources--virtual_host--properties--authentication--cookie_params--auth_hmac--prim_key--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](resources--virtual_host--properties--authentication--cookie_params--auth_hmac--prim_key--clear_secret_info.md): complete subsection reference.

## Next pages

- [authentication.cookie_params.auth_hmac.prim_key.blindfold_secret_info](resources--virtual_host--properties--authentication--cookie_params--auth_hmac--prim_key--blindfold_secret_info.md)
- [authentication.cookie_params.auth_hmac.prim_key.clear_secret_info](resources--virtual_host--properties--authentication--cookie_params--auth_hmac--prim_key--clear_secret_info.md)
- [authentication.cookie_params.auth_hmac](resources--virtual_host--properties--authentication--cookie_params--auth_hmac.md)
- [xcsh_virtual_host](../resources/virtual_host.md)
