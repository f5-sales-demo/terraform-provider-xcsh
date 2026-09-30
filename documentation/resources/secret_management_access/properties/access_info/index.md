---
page_title: "access_info"
subcategory: ""
description: "access_info for xcsh_secret_management_access."
xcsh_docs: {"aliases": [], "body_bytes": 4250, "body_sha256": "sha256:51044f9c11aceadb45e1da6e1c7753a1b88f2dee32f82d357a34f724cc5997b6", "child_ids": ["xcsh-docs:resources:secret_management_access:properties:access_info:rest_auth_info", "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config", "xcsh-docs:resources:secret_management_access:properties:access_info:vault_auth_info"], "collection_id": "xcsh-docs:resources:secret_management_access:collection", "completeness": "complete", "id": "xcsh-docs:resources:secret_management_access:properties:access_info", "parent_id": "xcsh-docs:resources:secret_management_access:reference", "path": "documentation/resources/secret_management_access/properties/access_info/index.md", "provider_name": "secret_management_access", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["access_info"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/secret_management_access/properties/access_info/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "access_info for xcsh_secret_management_access.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["secret_management_accessCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# access_info

Breadcrumbs:

- [xcsh_secret_management_access](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/)
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

- [rest_auth_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/rest_auth_info/): complete subsection reference.

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

- [tls_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/tls_config/): complete subsection reference.

- [vault_auth_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/vault_auth_info/): complete subsection reference.

## Next pages

- [access_info.rest_auth_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/rest_auth_info/)
- [access_info.tls_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/tls_config/)
- [access_info.vault_auth_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/vault_auth_info/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/)
- [xcsh_secret_management_access](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/)
