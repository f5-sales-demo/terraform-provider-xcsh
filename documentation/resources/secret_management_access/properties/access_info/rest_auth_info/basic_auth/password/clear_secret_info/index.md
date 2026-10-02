---
page_title: "access_info.rest_auth_info.basic_auth.password.clear_secret_info"
subcategory: ""
description: "ClearSecretInfoType specifies information about the Secret that is not encrypted."
xcsh_docs: {"aliases": ["access info rest auth info basic auth password clear secret info"], "body_bytes": 4353, "body_sha256": "sha256:0a5f79228b8d5bdf661b886564add3ebee870880191f639ffaa4a81c371d6aa8", "capabilities": ["identity"], "category": "identity", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:secret_management_access:collection", "completeness": "complete", "id": "xcsh-docs:resources:secret_management_access:properties:access_info:rest_auth_info:basic_auth:password:clear_secret_info", "parent_id": "xcsh-docs:resources:secret_management_access:properties:access_info:rest_auth_info:basic_auth:password", "path": "documentation/resources/secret_management_access/properties/access_info/rest_auth_info/basic_auth/password/clear_secret_info/index.md", "product": "distributed-cloud", "provider_name": "secret_management_access", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0120133001312011-0211303012020131-0230100102301300-2131203232221300-0202303221121231-1133201102323311-3211100222313112-3222201012122211", "registry_path": "docs/guides/resources--secret_management_access--reference--group-001.md", "relationships": [{"anchor": "schema-access_info--rest_auth_info--basic_auth--password--clear_secret_info--url", "enforcement": "provider-schema", "group": "access_info.rest_auth_info.basic_auth.password.clear_secret_info:RequiredObjectAttributes:url", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:secret_management_access:properties:access_info:rest_auth_info:basic_auth:password:clear_secret_info", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["access_info", "rest_auth_info", "basic_auth", "password", "clear_secret_info"], "schema_version": 1, "sections": [{"aliases": ["provider ref"], "anchor": "schema-access_info--rest_auth_info--basic_auth--password--clear_secret_info--provider_ref", "description": "Name of the Secret Management Access object that contains information about the store to GET encrypted bytes This field needs to be provided only if the URL scheme is not string:///.", "document_id": "xcsh-docs:resources:secret_management_access:properties:access_info:rest_auth_info:basic_auth:password:clear_secret_info", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["access_info", "rest_auth_info", "basic_auth", "password", "clear_secret_info", "provider_ref"], "syntax": "attribute", "type": "string"}, {"aliases": ["url"], "anchor": "schema-access_info--rest_auth_info--basic_auth--password--clear_secret_info--url", "description": "URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after Base64 decoding.", "document_id": "xcsh-docs:resources:secret_management_access:properties:access_info:rest_auth_info:basic_auth:password:clear_secret_info", "flags": ["optional", "sensitive"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["access_info", "rest_auth_info", "basic_auth", "password", "clear_secret_info", "url"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/secret_management_access/properties/access_info/rest_auth_info/basic_auth/password/clear_secret_info/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["secret_management_accessCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# access_info.rest_auth_info.basic_auth.password.clear_secret_info

Breadcrumbs:

- [xcsh_secret_management_access](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/)
- [access_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/)
- [access_info.rest_auth_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/rest_auth_info/)
- [access_info.rest_auth_info.basic_auth](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/rest_auth_info/basic_auth/)
- [access_info.rest_auth_info.basic_auth.password](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/rest_auth_info/basic_auth/password/)
- access_info.rest_auth_info.basic_auth.password.clear_secret_info

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
clear_secret_info {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-access_info--rest_auth_info--basic_auth--password--clear_secret_info--provider_ref"></a>

### provider_ref property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="schema-access_info--rest_auth_info--basic_auth--password--clear_secret_info--url"></a>

### url property

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

## Next pages

- [access_info.rest_auth_info.basic_auth.password](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/rest_auth_info/basic_auth/password/)
- [xcsh_secret_management_access](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/)
