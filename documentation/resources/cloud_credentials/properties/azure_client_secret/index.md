---
page_title: "azure_client_secret"
subcategory: "Infrastructure"
description: "Azure Credentials Client Secret type."
xcsh_docs: {"aliases": ["authentication", "azure client secret", "credential setup", "credentials"], "body_bytes": 4728, "body_sha256": "sha256:ed158f38e8710152074e846414247eb98ffcaafa6d1d9316b3725fce80faed69", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:cloud_credentials:properties:azure_client_secret:client_secret"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cloud_credentials:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_credentials:properties:azure_client_secret", "parent_id": "xcsh-docs:resources:cloud_credentials:reference", "path": "documentation/resources/cloud_credentials/properties/azure_client_secret/index.md", "product": "distributed-cloud", "provider_name": "cloud_credentials", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0031003222133102-3220220230302231-3221203120103122-2132001113201132-2100013132100213-2210320123302223-3120332020133311-3031023322021323", "registry_path": "docs/guides/resources--cloud_credentials--reference--group-001.md", "relationships": [{"anchor": "schema-azure_client_secret--client_id", "enforcement": "provider-schema", "group": "azure_client_secret:RequiredObjectAttributes:client_id,subscription_id,tenant_id", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cloud_credentials:properties:azure_client_secret", "type": "requires"}, {"anchor": "schema-azure_client_secret--subscription_id", "enforcement": "provider-schema", "group": "azure_client_secret:RequiredObjectAttributes:client_id,subscription_id,tenant_id", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cloud_credentials:properties:azure_client_secret", "type": "requires"}, {"anchor": "schema-azure_client_secret--tenant_id", "enforcement": "provider-schema", "group": "azure_client_secret:RequiredObjectAttributes:client_id,subscription_id,tenant_id", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cloud_credentials:properties:azure_client_secret", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["azure_client_secret"], "schema_version": 1, "sections": [{"aliases": ["client id"], "anchor": "schema-azure_client_secret--client_id", "description": "Client ID for your Azure service principal.", "document_id": "xcsh-docs:resources:cloud_credentials:properties:azure_client_secret", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["azure_client_secret", "client_id"], "syntax": "attribute", "type": "string"}, {"aliases": ["client secret"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:resources:cloud_credentials:properties:azure_client_secret:client_secret", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "azure_client_secret.client_secret:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cloud_credentials:properties:azure_client_secret:client_secret:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "azure_client_secret.client_secret:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cloud_credentials:properties:azure_client_secret:client_secret:clear_secret_info", "type": "conflicts"}], "schema_path": ["azure_client_secret", "client_secret"], "syntax": "block", "type": "object"}, {"aliases": ["subscription id"], "anchor": "schema-azure_client_secret--subscription_id", "description": "Subscription ID for your Azure service principal.", "document_id": "xcsh-docs:resources:cloud_credentials:properties:azure_client_secret", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["azure_client_secret", "subscription_id"], "syntax": "attribute", "type": "string"}, {"aliases": ["tenant id"], "anchor": "schema-azure_client_secret--tenant_id", "description": "Tenant ID for your Azure service principal.", "document_id": "xcsh-docs:resources:cloud_credentials:properties:azure_client_secret", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["azure_client_secret", "tenant_id"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_credentials/properties/azure_client_secret/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Azure Credentials Client Secret type.", "tasks": ["authentication", "configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_credentialsCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
