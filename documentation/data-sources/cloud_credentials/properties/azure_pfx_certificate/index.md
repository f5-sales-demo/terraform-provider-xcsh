---
page_title: "azure_pfx_certificate"
subcategory: "Infrastructure"
description: "Azure Credentials Client Certificate type."
xcsh_docs: {"aliases": ["authentication", "azure pfx certificate", "cert", "certificate", "credential setup", "credentials", "existing certificates", "tls certificates"], "body_bytes": 5492, "body_sha256": "sha256:df3fca7ed8c538a334d4cdaf79768c27dc5b6821c6e994bfc099e3b82092f92f", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:cloud_credentials:properties:azure_pfx_certificate:password"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:cloud_credentials:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cloud_credentials:properties:azure_pfx_certificate", "parent_id": "xcsh-docs:data-sources:cloud_credentials:reference", "path": "documentation/data-sources/cloud_credentials/properties/azure_pfx_certificate/index.md", "product": "distributed-cloud", "provider_name": "cloud_credentials", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2113031023203122-3103312223330121-2311021122221203-2130130001020303-1120000001300312-2133312023023223-2031132322002033-2131013330102232", "registry_path": "docs/guides/data-sources--cloud_credentials--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["azure_pfx_certificate"], "schema_version": 1, "sections": [{"aliases": ["cert", "certificate", "certificate url", "existing certificates", "tls certificates"], "anchor": "schema-azure_pfx_certificate--certificate_url", "description": "URL for Client Certificate in '.pfx' or '.p12' whose certificate is linked to service principal object Certificate URL can contain client certificate in string:///<Base64 of certificate> format. Here <Base64 of certificate> is base64 of '.pfx' or '.p12' binary file.", "document_id": "xcsh-docs:data-sources:cloud_credentials:properties:azure_pfx_certificate", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["azure_pfx_certificate", "certificate_url"], "syntax": "attribute", "type": "string"}, {"aliases": ["client id"], "anchor": "schema-azure_pfx_certificate--client_id", "description": "Client ID for your Azure service principal.", "document_id": "xcsh-docs:data-sources:cloud_credentials:properties:azure_pfx_certificate", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["azure_pfx_certificate", "client_id"], "syntax": "attribute", "type": "string"}, {"aliases": ["password"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:data-sources:cloud_credentials:properties:azure_pfx_certificate:password", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["azure_pfx_certificate", "password"], "syntax": "attribute", "type": "object"}, {"aliases": ["subscription id"], "anchor": "schema-azure_pfx_certificate--subscription_id", "description": "Subscription ID for your Azure service principal.", "document_id": "xcsh-docs:data-sources:cloud_credentials:properties:azure_pfx_certificate", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["azure_pfx_certificate", "subscription_id"], "syntax": "attribute", "type": "string"}, {"aliases": ["tenant id"], "anchor": "schema-azure_pfx_certificate--tenant_id", "description": "Tenant ID for your Azure service principal.", "document_id": "xcsh-docs:data-sources:cloud_credentials:properties:azure_pfx_certificate", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["azure_pfx_certificate", "tenant_id"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cloud_credentials/properties/azure_pfx_certificate/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Azure Credentials Client Certificate type.", "tasks": ["authentication", "configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_credentialsCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# azure_pfx_certificate

Breadcrumbs:

- [xcsh_cloud_credentials](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_credentials/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_credentials/properties/)
- azure_pfx_certificate

<a id="section"></a>

Type: `"single"`. Computed.

Azure Credentials Client Certificate type.

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

<a id="schema-azure_pfx_certificate--certificate_url"></a>

### certificate_url property

Type: `"string"`. Computed.

URL for Client Certificate in '.pfx' or '.p12' whose certificate is linked to service principal
object Certificate URL can contain client certificate in string:///&lt;Base64 of certificate&gt;
format. Here &lt;Base64 of certificate&gt; is base64 of '.pfx' or '.p12' binary file.

Upstream description:

URL for Client Certificate in '.pfx' or '.p12' whose certificate is linked to service principal
object Certificate URL can contain client certificate in string:///&lt;Base64 of certificate&gt;
format. Here &lt;Base64 of certificate&gt; is base64 of '.pfx' or '.p12' binary file.

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

- [password](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_credentials/properties/azure_pfx_certificate/password/): complete subsection reference.

<a id="schema-azure_pfx_certificate--subscription_id"></a>

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

<a id="schema-azure_pfx_certificate--tenant_id"></a>

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

- [azure_pfx_certificate.password](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_credentials/properties/azure_pfx_certificate/password/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_credentials/properties/)
- [xcsh_cloud_credentials](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_credentials/)
