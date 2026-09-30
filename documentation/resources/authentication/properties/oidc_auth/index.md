---
page_title: "oidc_auth"
subcategory: ""
description: "oidc_auth for xcsh_authentication."
xcsh_docs: {"aliases": [], "body_bytes": 4500, "body_sha256": "sha256:301ebdd30380be6f12e33f5f76f73e09970e7fedf34c6f6d4eb894b077afa16b", "child_ids": ["xcsh-docs:resources:authentication:properties:oidc_auth:client_secret", "xcsh-docs:resources:authentication:properties:oidc_auth:oidc_auth_params"], "collection_id": "xcsh-docs:resources:authentication:collection", "completeness": "complete", "id": "xcsh-docs:resources:authentication:properties:oidc_auth", "parent_id": "xcsh-docs:resources:authentication:reference", "path": "documentation/resources/authentication/properties/oidc_auth/index.md", "provider_name": "authentication", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["oidc_auth"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/authentication/properties/oidc_auth/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "oidc_auth for xcsh_authentication.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["authenticationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# oidc_auth

Breadcrumbs:

- [xcsh_authentication](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authentication/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authentication/properties/)
- oidc_auth

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

OIDCAuthType.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("oidc_client_id"),
  validators.ConflictingObjectAttributes("oidc_auth_params",
    "oidc_well_known_config_url")}
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
  "x-ves-oneof-field-auth_params_choice": "[\"oidc_auth_params\",\"oidc_well_known_config_url\"]"
}
```

Terraform syntax:

```terraform
oidc_auth {
  # Configure direct properties listed below.
}
```

## Direct properties

- [client_secret](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authentication/properties/oidc_auth/client_secret/): complete subsection reference.

- [oidc_auth_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authentication/properties/oidc_auth/oidc_auth_params/): complete subsection reference.

<a id="schema-oidc_auth--oidc_client_id"></a>

### oidc_client_id property

Type: `"string"`. Optional.

Client ID used while sending the Authorization Request to OIDC server.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="schema-oidc_auth--oidc_well_known_config_url"></a>

### oidc_well_known_config_url property

Type: `"string"`. Optional.

Exclusive with \[oidc\_auth\_params\] An OIDC well-known configuration URL that will be used to
fetch authentication related endpoints.

Upstream description:

Exclusive with \[oidc\_auth\_params\] An OIDC well-known configuration URL that will be used to
fetch authentication related endpoints.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

## Next pages

- [oidc_auth.client_secret](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authentication/properties/oidc_auth/client_secret/)
- [oidc_auth.oidc_auth_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authentication/properties/oidc_auth/oidc_auth_params/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authentication/properties/)
- [xcsh_authentication](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authentication/)
