---
page_title: "azure_client_secret"
subcategory: "Infrastructure"
description: "Azure Credentials Client Secret type."
xcsh_docs: {"aliases": ["authentication", "azure client secret", "credential setup", "credentials"], "body_bytes": 4019, "body_sha256": "sha256:489d206e67059448244cff396892e98d6ba84539314e7c5517b262be5d23cfe6", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:cloud_credentials:properties:azure_client_secret:client_secret"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:cloud_credentials:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cloud_credentials:properties:azure_client_secret", "parent_id": "xcsh-docs:data-sources:cloud_credentials:reference", "path": "documentation/data-sources/cloud_credentials/properties/azure_client_secret/index.md", "product": "distributed-cloud", "provider_name": "cloud_credentials", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0013331103030311-2121230021003010-3231120220132023-3122311310132133-3033001220020200-2102012312032210-1312112332323100-1313001301023203", "registry_path": "docs/guides/data-sources--cloud_credentials--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["azure_client_secret"], "schema_version": 1, "sections": [{"aliases": ["client id"], "anchor": "schema-azure_client_secret--client_id", "description": "Client ID for your Azure service principal.", "document_id": "xcsh-docs:data-sources:cloud_credentials:properties:azure_client_secret", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["azure_client_secret", "client_id"], "syntax": "attribute", "type": "string"}, {"aliases": ["client secret"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:data-sources:cloud_credentials:properties:azure_client_secret:client_secret", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["azure_client_secret", "client_secret"], "syntax": "attribute", "type": "object"}, {"aliases": ["subscription id"], "anchor": "schema-azure_client_secret--subscription_id", "description": "Subscription ID for your Azure service principal.", "document_id": "xcsh-docs:data-sources:cloud_credentials:properties:azure_client_secret", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["azure_client_secret", "subscription_id"], "syntax": "attribute", "type": "string"}, {"aliases": ["tenant id"], "anchor": "schema-azure_client_secret--tenant_id", "description": "Tenant ID for your Azure service principal.", "document_id": "xcsh-docs:data-sources:cloud_credentials:properties:azure_client_secret", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["azure_client_secret", "tenant_id"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cloud_credentials/properties/azure_client_secret/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Azure Credentials Client Secret type.", "tasks": ["authentication", "configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_credentialsCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

Upstream description:

Azure Credentials Client Secret type.

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

- [azure_client_secret.client_secret](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_credentials/properties/azure_client_secret/client_secret/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_credentials/properties/)
- [xcsh_cloud_credentials](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_credentials/)
