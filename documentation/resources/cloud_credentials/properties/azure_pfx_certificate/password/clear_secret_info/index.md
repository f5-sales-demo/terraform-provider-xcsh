---
page_title: "azure_pfx_certificate.password.clear_secret_info"
subcategory: "Infrastructure"
description: "ClearSecretInfoType specifies information about the Secret that is not encrypted."
xcsh_docs: {"aliases": ["azure pfx certificate password clear secret info"], "body_bytes": 3371, "body_sha256": "sha256:dfb1c9fd4c7ff074cb324a0a7db01d877f8e1ad158b1c2c847579d94f3519791", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cloud_credentials:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_credentials:properties:azure_pfx_certificate:password:clear_secret_info", "parent_id": "xcsh-docs:resources:cloud_credentials:properties:azure_pfx_certificate:password", "path": "documentation/resources/cloud_credentials/properties/azure_pfx_certificate/password/clear_secret_info/index.md", "product": "distributed-cloud", "provider_name": "cloud_credentials", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "resources", "registry_anchor": "canonical-0020203022020331-0102202103100332-3301232132332130-1331010013313123-0233022330321020-2210302021003012-3100021220023210-3002203232332002", "registry_path": "docs/guides/resources--cloud_credentials--reference--group-001.md", "relationships": [{"anchor": "schema-azure_pfx_certificate--password--clear_secret_info--url", "enforcement": "provider-schema", "group": "azure_pfx_certificate.password.clear_secret_info:RequiredObjectAttributes:url", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cloud_credentials:properties:azure_pfx_certificate:password:clear_secret_info", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["azure_pfx_certificate", "password", "clear_secret_info"], "schema_version": 1, "sections": [{"aliases": ["azure pfx certificate password clear secret info provider ref"], "anchor": "schema-azure_pfx_certificate--password--clear_secret_info--provider_ref", "description": "Name of the Secret Management Access object that contains information about the store to GET encrypted bytes This field needs to be provided only if the URL scheme is not string:///.", "document_id": "xcsh-docs:resources:cloud_credentials:properties:azure_pfx_certificate:password:clear_secret_info", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["azure_pfx_certificate", "password", "clear_secret_info", "provider_ref"], "syntax": "attribute", "type": "string"}, {"aliases": ["azure pfx certificate password clear secret info url"], "anchor": "schema-azure_pfx_certificate--password--clear_secret_info--url", "description": "URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after Base64 decoding.", "document_id": "xcsh-docs:resources:cloud_credentials:properties:azure_pfx_certificate:password:clear_secret_info", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional", "sensitive"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["azure_pfx_certificate", "password", "clear_secret_info", "url"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_credentials/properties/azure_pfx_certificate/password/clear_secret_info/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "tasks": ["authentication", "configuration"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["cloud_credentialsCreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# azure_pfx_certificate.password.clear_secret_info

Breadcrumbs:

- [xcsh_cloud_credentials](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/)
- [azure_pfx_certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/azure_pfx_certificate/)
- [azure_pfx_certificate.password](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/azure_pfx_certificate/password/)
- azure_pfx_certificate.password.clear_secret_info

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="schema-azure_pfx_certificate--password--clear_secret_info--provider_ref"></a>

### provider_ref property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="schema-azure_pfx_certificate--password--clear_secret_info--url"></a>

### url property

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-16T10:31:02+00:00"
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
