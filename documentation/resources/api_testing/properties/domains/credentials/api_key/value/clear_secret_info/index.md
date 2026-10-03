---
page_title: "domains.credentials.api_key.value.clear_secret_info"
subcategory: ""
description: "ClearSecretInfoType specifies information about the Secret that is not encrypted."
xcsh_docs: {"aliases": ["domains credentials api key value clear secret info"], "body_bytes": 4077, "body_sha256": "sha256:3abcad1b5c3db6455bfaa7cbe93b7df9a114ba1c7cdf8df5e2eb803e22bb5efc", "capabilities": ["api-management"], "category": "api-management", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:api_testing:collection", "completeness": "complete", "id": "xcsh-docs:resources:api_testing:properties:domains:credentials:api_key:value:clear_secret_info", "parent_id": "xcsh-docs:resources:api_testing:properties:domains:credentials:api_key:value", "path": "documentation/resources/api_testing/properties/domains/credentials/api_key/value/clear_secret_info/index.md", "product": "distributed-cloud", "provider_name": "api_testing", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3001133012120300-1300303320332233-3102000022233033-2330210212201023-3032213011331222-0123332203013322-2000000000132111-2003233300001211", "registry_path": "docs/guides/resources--api_testing--reference--group-001.md", "relationships": [{"anchor": "schema-domains--credentials--api_key--value--clear_secret_info--url", "enforcement": "provider-schema", "group": "domains.credentials.api_key.value.clear_secret_info:RequiredObjectAttributes:url", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:api_testing:properties:domains:credentials:api_key:value:clear_secret_info", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["domains", "credentials", "api_key", "value", "clear_secret_info"], "schema_version": 1, "sections": [{"aliases": ["domains credentials api key value clear secret info provider ref"], "anchor": "schema-domains--credentials--api_key--value--clear_secret_info--provider_ref", "description": "Name of the Secret Management Access object that contains information about the store to GET encrypted bytes This field needs to be provided only if the URL scheme is not string:///.", "document_id": "xcsh-docs:resources:api_testing:properties:domains:credentials:api_key:value:clear_secret_info", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["domains", "credentials", "api_key", "value", "clear_secret_info", "provider_ref"], "syntax": "attribute", "type": "string"}, {"aliases": ["domains credentials api key value clear secret info url"], "anchor": "schema-domains--credentials--api_key--value--clear_secret_info--url", "description": "URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after Base64 decoding.", "document_id": "xcsh-docs:resources:api_testing:properties:domains:credentials:api_key:value:clear_secret_info", "flags": ["optional", "sensitive"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["domains", "credentials", "api_key", "value", "clear_secret_info", "url"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/api_testing/properties/domains/credentials/api_key/value/clear_secret_info/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["api_testingCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# domains.credentials.api_key.value.clear_secret_info

Breadcrumbs:

- [xcsh_api_testing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/properties/)
- [domains](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/properties/domains/)
- [domains.credentials](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/properties/domains/credentials/)
- [domains.credentials.api_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/properties/domains/credentials/api_key/)
- [domains.credentials.api_key.value](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/properties/domains/credentials/api_key/value/)
- domains.credentials.api_key.value.clear_secret_info

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

<a id="schema-domains--credentials--api_key--value--clear_secret_info--provider_ref"></a>

### provider_ref property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="schema-domains--credentials--api_key--value--clear_secret_info--url"></a>

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

- [domains.credentials.api_key.value](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/properties/domains/credentials/api_key/value/)
- [xcsh_api_testing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/)
