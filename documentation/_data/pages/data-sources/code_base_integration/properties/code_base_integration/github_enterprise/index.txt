---
page_title: "code_base_integration.github_enterprise"
subcategory: ""
description: "code_base_integration.github_enterprise for xcsh_code_base_integration."
xcsh_docs: {"aliases": [], "body_bytes": 3557, "body_sha256": "sha256:0b845cfbb2876dc6f170bd82815c3661f7434adae5af357044f5d2a3bb1404a0", "child_ids": ["xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:github_enterprise:access_token"], "collection_id": "xcsh-docs:data-sources:code_base_integration:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:github_enterprise", "parent_id": "xcsh-docs:data-sources:code_base_integration:properties:code_base_integration", "path": "documentation/data-sources/code_base_integration/properties/code_base_integration/github_enterprise/index.md", "provider_name": "code_base_integration", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["code_base_integration", "github_enterprise"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/code_base_integration/properties/code_base_integration/github_enterprise/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "code_base_integration.github_enterprise for xcsh_code_base_integration.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["code_base_integrationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# code_base_integration.github_enterprise

Breadcrumbs:

- [xcsh_code_base_integration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/)
- [code_base_integration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/)
- code_base_integration.github_enterprise

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for github enterprise.

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

- [access_token](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/github_enterprise/access_token/): complete subsection reference.

<a id="schema-code_base_integration--github_enterprise--hostname"></a>

### hostname property

Type: `"string"`. Computed.

GitHub Hostname. Human-readable name for the resource

Upstream description:

Human-readable name for the resource

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "network",
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.95,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1123"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="schema-code_base_integration--github_enterprise--username"></a>

### username property

Type: `"string"`. Computed.

GitHub Username. Human-readable name for the resource

Upstream description:

Human-readable name for the resource

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "identity",
    "characterSet": {
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-zA-Z0-9_.-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

## Next pages

- [code_base_integration.github_enterprise.access_token](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/github_enterprise/access_token/)
- [code_base_integration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/)
- [xcsh_code_base_integration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/)
