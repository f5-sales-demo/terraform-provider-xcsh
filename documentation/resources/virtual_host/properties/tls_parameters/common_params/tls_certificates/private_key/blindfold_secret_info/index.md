---
page_title: "tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info"
subcategory: ""
description: "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management."
xcsh_docs: {"aliases": ["cert", "certificate", "existing certificates", "tls certificates", "tls parameters common params tls certificates private key blindfold secret info"], "body_bytes": 5496, "body_sha256": "sha256:e8da15a2ea35dcbd198f60b5189218711cc40f394fd63cc5980eae8f719625ec", "capabilities": ["load-balancing", "load-balancing.tls"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_host:properties:tls_parameters:common_params:tls_certificates:private_key:blindfold_secret_info", "parent_id": "xcsh-docs:resources:virtual_host:properties:tls_parameters:common_params:tls_certificates:private_key", "path": "documentation/resources/virtual_host/properties/tls_parameters/common_params/tls_certificates/private_key/blindfold_secret_info/index.md", "product": "distributed-cloud", "provider_name": "virtual_host", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2303223212323122-2111120020132110-2303311132222211-3122020120123201-1301200220013030-0121301312210220-3202033203132021-2010113331023333", "registry_path": "docs/guides/resources--virtual_host--reference--group-003.md", "relationships": [{"anchor": "schema-tls_parameters--common_params--tls_certificates--private_key--blindfold_secret_info--location", "enforcement": "provider-schema", "group": "tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info:RequiredObjectAttributes:location", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:tls_parameters:common_params:tls_certificates:private_key:blindfold_secret_info", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["tls_parameters", "common_params", "tls_certificates", "private_key", "blindfold_secret_info"], "schema_version": 1, "sections": [{"aliases": ["backend servers", "decryption provider", "origin servers", "upstream servers"], "anchor": "schema-tls_parameters--common_params--tls_certificates--private_key--blindfold_secret_info--decryption_provider", "description": "Name of the Secret Management Access object that contains information about the backend Secret Management service.", "document_id": "xcsh-docs:resources:virtual_host:properties:tls_parameters:common_params:tls_certificates:private_key:blindfold_secret_info", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tls_parameters", "common_params", "tls_certificates", "private_key", "blindfold_secret_info", "decryption_provider"], "syntax": "attribute", "type": "string"}, {"aliases": ["location"], "anchor": "schema-tls_parameters--common_params--tls_certificates--private_key--blindfold_secret_info--location", "description": "Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location.", "document_id": "xcsh-docs:resources:virtual_host:properties:tls_parameters:common_params:tls_certificates:private_key:blindfold_secret_info", "flags": ["optional", "sensitive"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tls_parameters", "common_params", "tls_certificates", "private_key", "blindfold_secret_info", "location"], "syntax": "attribute", "type": "string"}, {"aliases": ["store provider"], "anchor": "schema-tls_parameters--common_params--tls_certificates--private_key--blindfold_secret_info--store_provider", "description": "Name of the Secret Management Access object that contains information about the store to GET encrypted bytes This field needs to be provided only if the URL scheme is not string:///.", "document_id": "xcsh-docs:resources:virtual_host:properties:tls_parameters:common_params:tls_certificates:private_key:blindfold_secret_info", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tls_parameters", "common_params", "tls_certificates", "private_key", "blindfold_secret_info", "store_provider"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_host/properties/tls_parameters/common_params/tls_certificates/private_key/blindfold_secret_info/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info

Breadcrumbs:

- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/)
- [tls_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_parameters/)
- [tls_parameters.common_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_parameters/common_params/)
- [tls_parameters.common_params.tls_certificates](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_parameters/common_params/tls_certificates/)
- [tls_parameters.common_params.tls_certificates.private_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_parameters/common_params/tls_certificates/private_key/)
- tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("location")}
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
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-tls_parameters--common_params--tls_certificates--private_key--blindfold_secret_info--decryption_provider"></a>

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="schema-tls_parameters--common_params--tls_certificates--private_key--blindfold_secret_info--location"></a>

### location property

Type: `"string"`. Optional, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 131072),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="schema-tls_parameters--common_params--tls_certificates--private_key--blindfold_secret_info--store_provider"></a>

### store_provider property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

## Next pages

- [tls_parameters.common_params.tls_certificates.private_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_parameters/common_params/tls_certificates/private_key/)
- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/)
