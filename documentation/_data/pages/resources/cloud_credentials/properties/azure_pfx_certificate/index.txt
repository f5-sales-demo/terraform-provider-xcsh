---
page_title: "azure_pfx_certificate"
subcategory: "Infrastructure"
description: "Azure Credentials Client Certificate type."
xcsh_docs: {"aliases": ["azure pfx certificate"], "body_bytes": 5796, "body_sha256": "sha256:d0c04356cc4bdf64c52cf8f3b418064824fe925186d47363ef66ddf273d85f3b", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:cloud_credentials:properties:azure_pfx_certificate:password"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cloud_credentials:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_credentials:properties:azure_pfx_certificate", "parent_id": "xcsh-docs:resources:cloud_credentials:reference", "path": "documentation/resources/cloud_credentials/properties/azure_pfx_certificate/index.md", "product": "distributed-cloud", "provider_name": "cloud_credentials", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "resources", "registry_anchor": "canonical-3121121300221320-0323123210301120-3332103121012010-0133333031111230-3231111231030100-0102030303212231-0002210222100033-0213220221123010", "registry_path": "docs/guides/resources--cloud_credentials--reference--group-001.md", "relationships": [{"anchor": "schema-azure_pfx_certificate--certificate_url", "enforcement": "provider-schema", "group": "azure_pfx_certificate:RequiredObjectAttributes:certificate_url,client_id,subscription_id,tenant_id", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cloud_credentials:properties:azure_pfx_certificate", "type": "requires"}, {"anchor": "schema-azure_pfx_certificate--client_id", "enforcement": "provider-schema", "group": "azure_pfx_certificate:RequiredObjectAttributes:certificate_url,client_id,subscription_id,tenant_id", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cloud_credentials:properties:azure_pfx_certificate", "type": "requires"}, {"anchor": "schema-azure_pfx_certificate--subscription_id", "enforcement": "provider-schema", "group": "azure_pfx_certificate:RequiredObjectAttributes:certificate_url,client_id,subscription_id,tenant_id", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cloud_credentials:properties:azure_pfx_certificate", "type": "requires"}, {"anchor": "schema-azure_pfx_certificate--tenant_id", "enforcement": "provider-schema", "group": "azure_pfx_certificate:RequiredObjectAttributes:certificate_url,client_id,subscription_id,tenant_id", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cloud_credentials:properties:azure_pfx_certificate", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["azure_pfx_certificate"], "schema_version": 1, "sections": [{"aliases": ["azure pfx certificate certificate url", "cert", "certificate", "existing certificates", "tls certificates"], "anchor": "schema-azure_pfx_certificate--certificate_url", "description": "URL for Client Certificate in '.pfx' or '.p12' whose certificate is linked to service principal object Certificate URL can contain client certificate in string:///<Base64 of certificate> format. Here <Base64 of certificate> is base64 of '.pfx' or '.p12' binary file.", "document_id": "xcsh-docs:resources:cloud_credentials:properties:azure_pfx_certificate", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["azure_pfx_certificate", "certificate_url"], "syntax": "attribute", "type": "string"}, {"aliases": ["azure pfx certificate client id"], "anchor": "schema-azure_pfx_certificate--client_id", "description": "Client ID for your Azure service principal.", "document_id": "xcsh-docs:resources:cloud_credentials:properties:azure_pfx_certificate", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["azure_pfx_certificate", "client_id"], "syntax": "attribute", "type": "string"}, {"aliases": ["azure pfx certificate password"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:resources:cloud_credentials:properties:azure_pfx_certificate:password", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "azure_pfx_certificate.password:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cloud_credentials:properties:azure_pfx_certificate:password:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "azure_pfx_certificate.password:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cloud_credentials:properties:azure_pfx_certificate:password:clear_secret_info", "type": "conflicts"}], "schema_path": ["azure_pfx_certificate", "password"], "syntax": "block", "type": "object"}, {"aliases": ["azure pfx certificate subscription id"], "anchor": "schema-azure_pfx_certificate--subscription_id", "description": "Subscription ID for your Azure service principal.", "document_id": "xcsh-docs:resources:cloud_credentials:properties:azure_pfx_certificate", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["azure_pfx_certificate", "subscription_id"], "syntax": "attribute", "type": "string"}, {"aliases": ["azure pfx certificate tenant id"], "anchor": "schema-azure_pfx_certificate--tenant_id", "description": "Tenant ID for your Azure service principal.", "document_id": "xcsh-docs:resources:cloud_credentials:properties:azure_pfx_certificate", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["azure_pfx_certificate", "tenant_id"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_credentials/properties/azure_pfx_certificate/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Azure Credentials Client Certificate type.", "tasks": ["authentication", "configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["cloud_credentialsCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# azure_pfx_certificate

Breadcrumbs:

- [xcsh_cloud_credentials](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/)
- azure_pfx_certificate

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Azure Credentials Client Certificate type.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("certificate_url",
    "client_id",
    "subscription_id",
    "tenant_id")}
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
azure_pfx_certificate {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-azure_pfx_certificate--certificate_url"></a>

### certificate_url property

Type: `"string"`. Optional.

URL for Client Certificate in '.pfx' or '.p12' whose certificate is linked to service principal
object Certificate URL can contain client certificate in string:///&lt;Base64 of certificate&gt;
format. Here &lt;Base64 of certificate&gt; is base64 of '.pfx' or '.p12' binary file.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(8192),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8192,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 8192,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "8192",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "8192",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="schema-azure_pfx_certificate--client_id"></a>

### client_id property

Type: `"string"`. Optional.

Client ID for your Azure service principal.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

- [password](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/azure_pfx_certificate/password/): complete subsection reference.

<a id="schema-azure_pfx_certificate--subscription_id"></a>

### subscription_id property

Type: `"string"`. Optional.

Subscription ID for your Azure service principal.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="schema-azure_pfx_certificate--tenant_id"></a>

### tenant_id property

Type: `"string"`. Optional.

Tenant ID for your Azure service principal.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```
