---
page_title: "azure_client_secret"
subcategory: "Infrastructure"
description: "Azure Credentials Client Secret type."
xcsh_docs: {"aliases": ["azure client secret"], "body_bytes": 3530, "body_sha256": "sha256:5538eafd217dfa327f2acce009e164a073141ccb86ecc8762f3a0efd04471204", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:cloud_credentials:properties:azure_client_secret:client_secret"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:cloud_credentials:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cloud_credentials:properties:azure_client_secret", "parent_id": "xcsh-docs:data-sources:cloud_credentials:reference", "path": "documentation/data-sources/cloud_credentials/properties/azure_client_secret/index.md", "product": "distributed-cloud", "provider_name": "cloud_credentials", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-0013331103030311-2121230021003010-3231120220132023-3122311310132133-3033001220020200-2102012312032210-1312112332323100-1313001301023203", "registry_path": "docs/guides/data-sources--cloud_credentials--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["azure_client_secret"], "schema_version": 1, "sections": [{"aliases": ["azure client secret client id"], "anchor": "schema-azure_client_secret--client_id", "description": "Client ID for your Azure service principal.", "document_id": "xcsh-docs:data-sources:cloud_credentials:properties:azure_client_secret", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["azure_client_secret", "client_id"], "syntax": "attribute", "type": "string"}, {"aliases": ["azure client secret client secret"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:data-sources:cloud_credentials:properties:azure_client_secret:client_secret", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["azure_client_secret", "client_secret"], "syntax": "attribute", "type": "object"}, {"aliases": ["azure client secret subscription id"], "anchor": "schema-azure_client_secret--subscription_id", "description": "Subscription ID for your Azure service principal.", "document_id": "xcsh-docs:data-sources:cloud_credentials:properties:azure_client_secret", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["azure_client_secret", "subscription_id"], "syntax": "attribute", "type": "string"}, {"aliases": ["azure client secret tenant id"], "anchor": "schema-azure_client_secret--tenant_id", "description": "Tenant ID for your Azure service principal.", "document_id": "xcsh-docs:data-sources:cloud_credentials:properties:azure_client_secret", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["azure_client_secret", "tenant_id"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cloud_credentials/properties/azure_client_secret/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Azure Credentials Client Secret type.", "tasks": ["authentication", "configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["cloud_credentialsCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# azure_client_secret

Breadcrumbs:

- [xcsh_cloud_credentials](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_credentials/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_credentials/properties/)
- azure_client_secret

<a id="section"></a>

Type: `"single"`. Computed.

Azure Client Secret. Azure Credentials Client Secret type.

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

## Direct properties

<a id="schema-azure_client_secret--client_id"></a>

### client_id property

Type: `"string"`. Computed.

Client ID for your Azure service principal.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [client_secret](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_credentials/properties/azure_client_secret/client_secret/): complete subsection reference.

<a id="schema-azure_client_secret--subscription_id"></a>

### subscription_id property

Type: `"string"`. Computed.

Subscription ID for your Azure service principal.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

Type: `"string"`. Computed.

Tenant ID for your Azure service principal.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
