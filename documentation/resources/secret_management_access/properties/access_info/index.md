---
page_title: "access_info"
subcategory: ""
description: "HostAccessInfoType contains the information about how to connect to the remote host."
xcsh_docs: {"aliases": ["access info"], "body_bytes": 4349, "body_sha256": "sha256:e2f69f01949fbc64dc24d2c0d7d82aae52789a494b21d0e672efd61fdea38c06", "capabilities": ["identity"], "category": "identity", "child_ids": ["xcsh-docs:resources:secret_management_access:properties:access_info:rest_auth_info", "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config", "xcsh-docs:resources:secret_management_access:properties:access_info:vault_auth_info"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:secret_management_access:collection", "completeness": "complete", "id": "xcsh-docs:resources:secret_management_access:properties:access_info", "parent_id": "xcsh-docs:resources:secret_management_access:reference", "path": "documentation/resources/secret_management_access/properties/access_info/index.md", "product": "distributed-cloud", "provider_name": "secret_management_access", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2302030002201031-3031010213110200-2013100112131001-2333012320201231-1232103301212112-3331320111021012-0031323320220221-1211023031200100", "registry_path": "docs/guides/resources--secret_management_access--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "access_info:ConflictingObjectAttributes:rest_auth_info,vault_auth_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:secret_management_access:properties:access_info:rest_auth_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "access_info:ConflictingObjectAttributes:rest_auth_info,vault_auth_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:secret_management_access:properties:access_info:vault_auth_info", "type": "conflicts"}, {"anchor": "schema-access_info--server_endpoint", "enforcement": "provider-schema", "group": "access_info:RequiredObjectAttributes:server_endpoint", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:secret_management_access:properties:access_info", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["access_info"], "schema_version": 1, "sections": [{"aliases": ["rest auth info"], "anchor": "section", "description": "Authentication parameters for REST based hosts.", "document_id": "xcsh-docs:resources:secret_management_access:properties:access_info:rest_auth_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "access_info.rest_auth_info:ConflictingObjectAttributes:basic_auth,headers_auth", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:secret_management_access:properties:access_info:rest_auth_info:basic_auth", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "access_info.rest_auth_info:ConflictingObjectAttributes:basic_auth,query_params_auth", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:secret_management_access:properties:access_info:rest_auth_info:basic_auth", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "access_info.rest_auth_info:ConflictingObjectAttributes:basic_auth,headers_auth", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:secret_management_access:properties:access_info:rest_auth_info:headers_auth", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "access_info.rest_auth_info:ConflictingObjectAttributes:headers_auth,query_params_auth", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:secret_management_access:properties:access_info:rest_auth_info:headers_auth", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "access_info.rest_auth_info:ConflictingObjectAttributes:basic_auth,query_params_auth", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:secret_management_access:properties:access_info:rest_auth_info:query_params_auth", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "access_info.rest_auth_info:ConflictingObjectAttributes:headers_auth,query_params_auth", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:secret_management_access:properties:access_info:rest_auth_info:query_params_auth", "type": "conflicts"}], "schema_path": ["access_info", "rest_auth_info"], "syntax": "block", "type": "object"}, {"aliases": ["scheme"], "anchor": "schema-access_info--scheme", "description": "SchemeType is used to indicate URL scheme HTTP:// scheme HTTPS:// scheme.", "document_id": "xcsh-docs:resources:secret_management_access:properties:access_info", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["access_info", "scheme"], "syntax": "attribute", "type": "string"}, {"aliases": ["server endpoint"], "anchor": "schema-access_info--server_endpoint", "description": "Endpoint to connect to, in host:port format.", "document_id": "xcsh-docs:resources:secret_management_access:properties:access_info", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["access_info", "server_endpoint"], "syntax": "attribute", "type": "string"}, {"aliases": ["tls config"], "anchor": "section", "description": "TLS configuration for upstream connections.", "document_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-access_info--tls_config--max_session_keys", "enforcement": "provider-schema", "group": "access_info.tls_config:ConflictingObjectAttributes:default_session_key_caching,max_session_keys", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config", "type": "conflicts"}, {"anchor": "schema-access_info--tls_config--max_session_keys", "enforcement": "provider-schema", "group": "access_info.tls_config:ConflictingObjectAttributes:disable_session_key_caching,max_session_keys", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config", "type": "conflicts"}, {"anchor": "schema-access_info--tls_config--sni", "enforcement": "provider-schema", "group": "access_info.tls_config:ConflictingObjectAttributes:disable_sni,sni", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config", "type": "conflicts"}, {"anchor": "schema-access_info--tls_config--sni", "enforcement": "provider-schema", "group": "access_info.tls_config:ConflictingObjectAttributes:sni,use_host_header_as_sni", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "access_info.tls_config:ConflictingObjectAttributes:cert_params,common_params", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:cert_params", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "access_info.tls_config:ConflictingObjectAttributes:cert_params,common_params", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:common_params", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "access_info.tls_config:ConflictingObjectAttributes:default_session_key_caching,disable_session_key_caching", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:default_session_key_caching", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "access_info.tls_config:ConflictingObjectAttributes:default_session_key_caching,max_session_keys", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:default_session_key_caching", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "access_info.tls_config:ConflictingObjectAttributes:default_session_key_caching,disable_session_key_caching", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:disable_session_key_caching", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "access_info.tls_config:ConflictingObjectAttributes:disable_session_key_caching,max_session_keys", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:disable_session_key_caching", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "access_info.tls_config:ConflictingObjectAttributes:disable_sni,sni", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:disable_sni", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "access_info.tls_config:ConflictingObjectAttributes:disable_sni,use_host_header_as_sni", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:disable_sni", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "access_info.tls_config:ConflictingObjectAttributes:disable_sni,use_host_header_as_sni", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:use_host_header_as_sni", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "access_info.tls_config:ConflictingObjectAttributes:sni,use_host_header_as_sni", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:use_host_header_as_sni", "type": "conflicts"}], "schema_path": ["access_info", "tls_config"], "syntax": "block", "type": "object"}, {"aliases": ["vault auth info"], "anchor": "section", "description": "Authentication parameters for Hashicorp Vault hosts.", "document_id": "xcsh-docs:resources:secret_management_access:properties:access_info:vault_auth_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "access_info.vault_auth_info:ConflictingObjectAttributes:app_role_auth,token", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:secret_management_access:properties:access_info:vault_auth_info:app_role_auth", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "access_info.vault_auth_info:ConflictingObjectAttributes:app_role_auth,token", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:secret_management_access:properties:access_info:vault_auth_info:token", "type": "conflicts"}], "schema_path": ["access_info", "vault_auth_info"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/secret_management_access/properties/access_info/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "HostAccessInfoType contains the information about how to connect to the remote host.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["secret_management_accessCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

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
