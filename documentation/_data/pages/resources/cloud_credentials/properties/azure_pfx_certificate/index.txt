---
page_title: "azure_pfx_certificate"
subcategory: "Infrastructure"
description: "Azure Credentials Client Certificate type."
xcsh_docs: {"aliases": ["authentication", "azure pfx certificate", "cert", "certificate", "credential setup", "credentials", "existing certificates", "tls certificates"], "body_bytes": 6365, "body_sha256": "sha256:ece7b41ec9be305ed5c1de84a3f8c89f2641237f69a33ddb84cfbb72cf1d3103", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:cloud_credentials:properties:azure_pfx_certificate:password"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cloud_credentials:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_credentials:properties:azure_pfx_certificate", "parent_id": "xcsh-docs:resources:cloud_credentials:reference", "path": "documentation/resources/cloud_credentials/properties/azure_pfx_certificate/index.md", "product": "distributed-cloud", "provider_name": "cloud_credentials", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3121121300221320-0323123210301120-3332103121012010-0133333031111230-3231111231030100-0102030303212231-0002210222100033-0213220221123010", "registry_path": "docs/guides/resources--cloud_credentials--reference--group-001.md", "relationships": [{"anchor": "schema-azure_pfx_certificate--certificate_url", "enforcement": "provider-schema", "group": "azure_pfx_certificate:RequiredObjectAttributes:certificate_url,client_id,subscription_id,tenant_id", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cloud_credentials:properties:azure_pfx_certificate", "type": "requires"}, {"anchor": "schema-azure_pfx_certificate--client_id", "enforcement": "provider-schema", "group": "azure_pfx_certificate:RequiredObjectAttributes:certificate_url,client_id,subscription_id,tenant_id", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cloud_credentials:properties:azure_pfx_certificate", "type": "requires"}, {"anchor": "schema-azure_pfx_certificate--subscription_id", "enforcement": "provider-schema", "group": "azure_pfx_certificate:RequiredObjectAttributes:certificate_url,client_id,subscription_id,tenant_id", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cloud_credentials:properties:azure_pfx_certificate", "type": "requires"}, {"anchor": "schema-azure_pfx_certificate--tenant_id", "enforcement": "provider-schema", "group": "azure_pfx_certificate:RequiredObjectAttributes:certificate_url,client_id,subscription_id,tenant_id", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cloud_credentials:properties:azure_pfx_certificate", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["azure_pfx_certificate"], "schema_version": 1, "sections": [{"aliases": ["cert", "certificate", "certificate url", "existing certificates", "tls certificates"], "anchor": "schema-azure_pfx_certificate--certificate_url", "description": "URL for Client Certificate in '.pfx' or '.p12' whose certificate is linked to service principal object Certificate URL can contain client certificate in string:///<Base64 of certificate> format. Here <Base64 of certificate> is base64 of '.pfx' or '.p12' binary file.", "document_id": "xcsh-docs:resources:cloud_credentials:properties:azure_pfx_certificate", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["azure_pfx_certificate", "certificate_url"], "syntax": "attribute", "type": "string"}, {"aliases": ["client id"], "anchor": "schema-azure_pfx_certificate--client_id", "description": "Client ID for your Azure service principal.", "document_id": "xcsh-docs:resources:cloud_credentials:properties:azure_pfx_certificate", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["azure_pfx_certificate", "client_id"], "syntax": "attribute", "type": "string"}, {"aliases": ["password"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:resources:cloud_credentials:properties:azure_pfx_certificate:password", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "azure_pfx_certificate.password:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cloud_credentials:properties:azure_pfx_certificate:password:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "azure_pfx_certificate.password:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cloud_credentials:properties:azure_pfx_certificate:password:clear_secret_info", "type": "conflicts"}], "schema_path": ["azure_pfx_certificate", "password"], "syntax": "block", "type": "object"}, {"aliases": ["subscription id"], "anchor": "schema-azure_pfx_certificate--subscription_id", "description": "Subscription ID for your Azure service principal.", "document_id": "xcsh-docs:resources:cloud_credentials:properties:azure_pfx_certificate", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["azure_pfx_certificate", "subscription_id"], "syntax": "attribute", "type": "string"}, {"aliases": ["tenant id"], "anchor": "schema-azure_pfx_certificate--tenant_id", "description": "Tenant ID for your Azure service principal.", "document_id": "xcsh-docs:resources:cloud_credentials:properties:azure_pfx_certificate", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["azure_pfx_certificate", "tenant_id"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_credentials/properties/azure_pfx_certificate/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Azure Credentials Client Certificate type.", "tasks": ["authentication", "configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_credentialsCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

Upstream description:

URL for Client Certificate in '.pfx' or '.p12' whose certificate is linked to service principal
object Certificate URL can contain client certificate in string:///&lt;Base64 of certificate&gt;
format. Here &lt;Base64 of certificate&gt; is base64 of '.pfx' or '.p12' binary file.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

## Next pages

- [azure_pfx_certificate.password](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/azure_pfx_certificate/password/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/)
- [xcsh_cloud_credentials](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/)
