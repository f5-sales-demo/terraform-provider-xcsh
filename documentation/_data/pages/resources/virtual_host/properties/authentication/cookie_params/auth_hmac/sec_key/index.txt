---
page_title: "authentication.cookie_params.auth_hmac.sec_key"
subcategory: ""
description: "authentication.cookie_params.auth_hmac.sec_key for xcsh_virtual_host."
xcsh_docs: {"aliases": [], "body_bytes": 2693, "body_sha256": "sha256:41aae707e6a277c8d095ffea9a838eb6846868c92020c80b539ed568ef7bf1df", "child_ids": ["xcsh-docs:resources:virtual_host:properties:authentication:cookie_params:auth_hmac:sec_key:blindfold_secret_info", "xcsh-docs:resources:virtual_host:properties:authentication:cookie_params:auth_hmac:sec_key:clear_secret_info"], "collection_id": "xcsh-docs:resources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_host:properties:authentication:cookie_params:auth_hmac:sec_key", "parent_id": "xcsh-docs:resources:virtual_host:properties:authentication:cookie_params:auth_hmac", "path": "documentation/resources/virtual_host/properties/authentication/cookie_params/auth_hmac/sec_key/index.md", "provider_name": "virtual_host", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["authentication", "cookie_params", "auth_hmac", "sec_key"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_host/properties/authentication/cookie_params/auth_hmac/sec_key/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "authentication.cookie_params.auth_hmac.sec_key for xcsh_virtual_host.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# authentication.cookie_params.auth_hmac.sec_key

Breadcrumbs:

- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/)
- [authentication](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/)
- [authentication.cookie_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/cookie_params/)
- [authentication.cookie_params.auth_hmac](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/cookie_params/auth_hmac/)
- authentication.cookie_params.auth_hmac.sec_key

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

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/cookie_params/auth_hmac/sec_key/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/cookie_params/auth_hmac/sec_key/clear_secret_info/): complete subsection reference.

## Next pages

- [authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/cookie_params/auth_hmac/sec_key/blindfold_secret_info/)
- [authentication.cookie_params.auth_hmac.sec_key.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/cookie_params/auth_hmac/sec_key/clear_secret_info/)
- [authentication.cookie_params.auth_hmac](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/cookie_params/auth_hmac/)
- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/)
