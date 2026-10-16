---
page_title: "authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info"
subcategory: ""
description: "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management."
xcsh_docs: {"aliases": ["authentication cookie params auth hmac sec key blindfold secret info"], "body_bytes": 4413, "body_sha256": "sha256:8ef5ec8590057df87fb81a214e6788f15fc31d8153ba8ece03e53d99e4b3c563", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_host:properties:authentication:cookie_params:auth_hmac:sec_key:blindfold_secret_info", "parent_id": "xcsh-docs:resources:virtual_host:properties:authentication:cookie_params:auth_hmac:sec_key", "path": "documentation/resources/virtual_host/properties/authentication/cookie_params/auth_hmac/sec_key/blindfold_secret_info/index.md", "product": "distributed-cloud", "provider_name": "virtual_host", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "resources", "registry_anchor": "canonical-3113131201300102-2023030131011101-3132321133033312-2211303211003103-2233313203010322-1112021303203003-3103330300320111-3331000131123222", "registry_path": "docs/guides/resources--virtual_host--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["authentication", "cookie_params", "auth_hmac", "sec_key", "blindfold_secret_info"], "schema_version": 1, "sections": [{"aliases": ["authentication cookie params auth hmac sec key blindfold secret info decryption provider"], "anchor": "schema-authentication--cookie_params--auth_hmac--sec_key--blindfold_secret_info--decryption_provider", "description": "Name of the Secret Management Access object that contains information about the backend Secret Management service.", "document_id": "xcsh-docs:resources:virtual_host:properties:authentication:cookie_params:auth_hmac:sec_key:blindfold_secret_info", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["authentication", "cookie_params", "auth_hmac", "sec_key", "blindfold_secret_info", "decryption_provider"], "syntax": "attribute", "type": "string"}, {"aliases": ["authentication cookie params auth hmac sec key blindfold secret info location"], "anchor": "schema-authentication--cookie_params--auth_hmac--sec_key--blindfold_secret_info--location", "description": "Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location.", "document_id": "xcsh-docs:resources:virtual_host:properties:authentication:cookie_params:auth_hmac:sec_key:blindfold_secret_info", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional", "sensitive"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["authentication", "cookie_params", "auth_hmac", "sec_key", "blindfold_secret_info", "location"], "syntax": "attribute", "type": "string"}, {"aliases": ["authentication cookie params auth hmac sec key blindfold secret info store provider"], "anchor": "schema-authentication--cookie_params--auth_hmac--sec_key--blindfold_secret_info--store_provider", "description": "Name of the Secret Management Access object that contains information about the store to GET encrypted bytes This field needs to be provided only if the URL scheme is not string:///.", "document_id": "xcsh-docs:resources:virtual_host:properties:authentication:cookie_params:auth_hmac:sec_key:blindfold_secret_info", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["authentication", "cookie_params", "auth_hmac", "sec_key", "blindfold_secret_info", "store_provider"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_host/properties/authentication/cookie_params/auth_hmac/sec_key/blindfold_secret_info/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info

Breadcrumbs:

- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/)
- [authentication](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/)
- [authentication.cookie_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/cookie_params/)
- [authentication.cookie_params.auth_hmac](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/cookie_params/auth_hmac/)
- [authentication.cookie_params.auth_hmac.sec_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/cookie_params/auth_hmac/sec_key/)
- authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

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
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-authentication--cookie_params--auth_hmac--sec_key--blindfold_secret_info--decryption_provider"></a>

### decryption_provider property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the backend Secret
Management service.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-15T10:49:30+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-authentication--cookie_params--auth_hmac--sec_key--blindfold_secret_info--location"></a>

### location property

Type: `"string"`. Optional, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
    "metadata": {
      "category": "content",
      "confidence": 1.0,
      "note": "Blindfold envelope encryption (AES-256-GCM + RSA-OAEP) of an RSA-2048 TLS private key produces ~3700 char string:/// URL. 128KB max secret size = ~175KB base64. Discovery reported 1024 which is incorrect.",
      "source": "manual-override",
      "validatedAt": "2026-10-15T10:49:30+00:00"
    },
    "minLength": 4
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
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="schema-authentication--cookie_params--auth_hmac--sec_key--blindfold_secret_info--store_provider"></a>

### store_provider property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-15T10:49:30+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```
