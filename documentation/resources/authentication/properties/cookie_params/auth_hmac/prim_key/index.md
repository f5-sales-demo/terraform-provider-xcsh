---
page_title: "cookie_params.auth_hmac.prim_key"
subcategory: ""
description: "cookie_params.auth_hmac.prim_key for xcsh_authentication."
xcsh_docs: {"aliases": [], "body_bytes": 2489, "body_sha256": "sha256:ac27a1dd951b2f083e594436090ca051cc2db8653c7e371b0abbbeba7f433b00", "child_ids": ["xcsh-docs:resources:authentication:properties:cookie_params:auth_hmac:prim_key:blindfold_secret_info", "xcsh-docs:resources:authentication:properties:cookie_params:auth_hmac:prim_key:clear_secret_info"], "collection_id": "xcsh-docs:resources:authentication:collection", "completeness": "complete", "id": "xcsh-docs:resources:authentication:properties:cookie_params:auth_hmac:prim_key", "parent_id": "xcsh-docs:resources:authentication:properties:cookie_params:auth_hmac", "path": "documentation/resources/authentication/properties/cookie_params/auth_hmac/prim_key/index.md", "provider_name": "authentication", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["cookie_params", "auth_hmac", "prim_key"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/authentication/properties/cookie_params/auth_hmac/prim_key/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cookie_params.auth_hmac.prim_key for xcsh_authentication.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["authenticationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cookie_params.auth_hmac.prim_key

Breadcrumbs:

- [xcsh_authentication](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authentication/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authentication/properties/)
- [cookie_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authentication/properties/cookie_params/)
- [cookie_params.auth_hmac](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authentication/properties/cookie_params/auth_hmac/)
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

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authentication/properties/cookie_params/auth_hmac/prim_key/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authentication/properties/cookie_params/auth_hmac/prim_key/clear_secret_info/): complete subsection reference.

## Next pages

- [cookie_params.auth_hmac.prim_key.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authentication/properties/cookie_params/auth_hmac/prim_key/blindfold_secret_info/)
- [cookie_params.auth_hmac.prim_key.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authentication/properties/cookie_params/auth_hmac/prim_key/clear_secret_info/)
- [cookie_params.auth_hmac](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authentication/properties/cookie_params/auth_hmac/)
- [xcsh_authentication](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authentication/)
