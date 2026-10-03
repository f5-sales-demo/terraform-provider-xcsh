---
page_title: "azure_client_secret"
subcategory: "Infrastructure"
description: "Azure Credentials Client Secret type."
xcsh_docs: {"aliases": ["azure client secret"], "body_bytes": 4728, "body_sha256": "sha256:ed158f38e8710152074e846414247eb98ffcaafa6d1d9316b3725fce80faed69", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:cloud_credentials:properties:azure_client_secret:client_secret"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cloud_credentials:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_credentials:properties:azure_client_secret", "parent_id": "xcsh-docs:resources:cloud_credentials:reference", "path": "documentation/resources/cloud_credentials/properties/azure_client_secret/index.md", "product": "distributed-cloud", "provider_name": "cloud_credentials", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0031003222133102-3220220230302231-3221203120103122-2132001113201132-2100013132100213-2210320123302223-3120332020133311-3031023322021323", "registry_path": "docs/guides/resources--cloud_credentials--reference--group-001.md", "relationships": [{"anchor": "schema-azure_client_secret--client_id", "enforcement": "provider-schema", "group": "azure_client_secret:RequiredObjectAttributes:client_id,subscription_id,tenant_id", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cloud_credentials:properties:azure_client_secret", "type": "requires"}, {"anchor": "schema-azure_client_secret--subscription_id", "enforcement": "provider-schema", "group": "azure_client_secret:RequiredObjectAttributes:client_id,subscription_id,tenant_id", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cloud_credentials:properties:azure_client_secret", "type": "requires"}, {"anchor": "schema-azure_client_secret--tenant_id", "enforcement": "provider-schema", "group": "azure_client_secret:RequiredObjectAttributes:client_id,subscription_id,tenant_id", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cloud_credentials:properties:azure_client_secret", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["azure_client_secret"], "schema_version": 1, "sections": [{"aliases": ["azure client secret client id"], "anchor": "schema-azure_client_secret--client_id", "description": "Client ID for your Azure service principal.", "document_id": "xcsh-docs:resources:cloud_credentials:properties:azure_client_secret", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["azure_client_secret", "client_id"], "syntax": "attribute", "type": "string"}, {"aliases": ["azure client secret client secret"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:resources:cloud_credentials:properties:azure_client_secret:client_secret", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "azure_client_secret.client_secret:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cloud_credentials:properties:azure_client_secret:client_secret:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "azure_client_secret.client_secret:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cloud_credentials:properties:azure_client_secret:client_secret:clear_secret_info", "type": "conflicts"}], "schema_path": ["azure_client_secret", "client_secret"], "syntax": "block", "type": "object"}, {"aliases": ["azure client secret subscription id"], "anchor": "schema-azure_client_secret--subscription_id", "description": "Subscription ID for your Azure service principal.", "document_id": "xcsh-docs:resources:cloud_credentials:properties:azure_client_secret", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["azure_client_secret", "subscription_id"], "syntax": "attribute", "type": "string"}, {"aliases": ["azure client secret tenant id"], "anchor": "schema-azure_client_secret--tenant_id", "description": "Tenant ID for your Azure service principal.", "document_id": "xcsh-docs:resources:cloud_credentials:properties:azure_client_secret", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["azure_client_secret", "tenant_id"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_credentials/properties/azure_client_secret/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Azure Credentials Client Secret type.", "tasks": ["authentication", "configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["cloud_credentialsCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# azure_client_secret

Breadcrumbs:

- [xcsh_cloud_credentials](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/)
- azure_client_secret

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Azure Client Secret. Azure Credentials Client Secret type.

Upstream description:

Azure Credentials Client Secret type.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("client_id",
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
azure_client_secret {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-azure_client_secret--client_id"></a>

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

- [client_secret](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/azure_client_secret/client_secret/): complete subsection reference.

<a id="schema-azure_client_secret--subscription_id"></a>

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

<a id="schema-azure_client_secret--tenant_id"></a>

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

- [azure_client_secret.client_secret](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/azure_client_secret/client_secret/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/)
- [xcsh_cloud_credentials](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/)
