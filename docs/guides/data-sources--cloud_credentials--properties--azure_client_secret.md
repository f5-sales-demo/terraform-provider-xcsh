---
page_title: "azure_client_secret"
subcategory: "Infrastructure"
description: "azure_client_secret for xcsh_cloud_credentials."
xcsh_docs: {"aliases": [], "body_bytes": 3711, "body_sha256": "sha256:17d0941e728e21ce94b4e9a13d70816220c5d009d08e55f4785d1e8c6194fe28", "canonical_id": "xcsh-docs:data-sources:cloud_credentials:properties:azure_client_secret", "child_ids": ["xcsh-docs:data-sources:cloud_credentials:properties:azure_client_secret:client_secret"], "collection_id": "xcsh-docs:data-sources:cloud_credentials:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cloud_credentials:properties:azure_client_secret", "parent_id": "xcsh-docs:data-sources:cloud_credentials:reference", "path": "docs/guides/data-sources--cloud_credentials--properties--azure_client_secret.md", "provider_name": "cloud_credentials", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["azure_client_secret"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cloud_credentials/properties/azure_client_secret/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "azure_client_secret for xcsh_cloud_credentials.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_credentialsCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# azure_client_secret

Breadcrumbs:

- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md)
- [Property reference](data-sources--cloud_credentials--reference.md)
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

- [client_secret](data-sources--cloud_credentials--properties--azure_client_secret--client_secret.md): complete subsection reference.

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

- [azure_client_secret.client_secret](data-sources--cloud_credentials--properties--azure_client_secret--client_secret.md)
- [Property reference](data-sources--cloud_credentials--reference.md)
- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md)
