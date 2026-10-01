---
page_title: "access_info"
subcategory: ""
description: "access_info for xcsh_secret_management_access."
xcsh_docs: {"aliases": [], "body_bytes": 3841, "body_sha256": "sha256:36073e3e6d251e9f7d10bb5b8ca8b7a7a6eac2a6b807b8e1771a1496958d4637", "canonical_id": "xcsh-docs:resources:secret_management_access:properties:access_info", "child_ids": ["xcsh-docs:resources:secret_management_access:properties:access_info:rest_auth_info", "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config", "xcsh-docs:resources:secret_management_access:properties:access_info:vault_auth_info"], "collection_id": "xcsh-docs:resources:secret_management_access:collection", "completeness": "complete", "id": "xcsh-docs:resources:secret_management_access:properties:access_info", "parent_id": "xcsh-docs:resources:secret_management_access:reference", "path": "docs/guides/resources--secret_management_access--properties--access_info.md", "provider_name": "secret_management_access", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["access_info"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/secret_management_access/properties/access_info/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "access_info for xcsh_secret_management_access.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["secret_management_accessCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# access_info

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md)
- [Property reference](resources--secret_management_access--reference.md)
- access_info

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

HostAccessInfoType contains the information about how to connect to the remote host.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("server_endpoint"),
  validators.ConflictingObjectAttributes("rest_auth_info",
    "vault_auth_info")}
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
  "x-ves-oneof-field-auth_params": "[\"rest_auth_info\",\"vault_auth_info\"]"
}
```

Terraform syntax:

```terraform
access_info {
  # Configure direct properties listed below.
}
```

## Direct properties

- [rest_auth_info](resources--secret_management_access--properties--access_info--rest_auth_info.md): complete subsection reference.

<a id="schema-access_info--scheme"></a>

### scheme property

Type: `"string"`. Optional.

\[Enum: HTTP|HTTPS\] SchemeType is used to indicate URL scheme HTTP:// scheme HTTPS:// scheme.
Possible values are \`HTTP\`, \`HTTPS\`. Defaults to \`HTTP\`.

Upstream description:

SchemeType is used to indicate URL scheme

HTTP:// scheme HTTPS:// scheme.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("HTTP",
    "HTTPS"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "HTTP",
  "enum": [
    "HTTP",
    "HTTPS"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-access_info--server_endpoint"></a>

### server_endpoint property

Type: `"string"`. Optional.

Endpoint to connect to, in host:port format.

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

- [tls_config](resources--secret_management_access--properties--access_info--tls_config.md): complete subsection reference.

- [vault_auth_info](resources--secret_management_access--properties--access_info--vault_auth_info.md): complete subsection reference.

## Next pages

- [access_info.rest_auth_info](resources--secret_management_access--properties--access_info--rest_auth_info.md)
- [access_info.tls_config](resources--secret_management_access--properties--access_info--tls_config.md)
- [access_info.vault_auth_info](resources--secret_management_access--properties--access_info--vault_auth_info.md)
- [Property reference](resources--secret_management_access--reference.md)
- [xcsh_secret_management_access](../resources/secret_management_access.md)
