---
page_title: "code_base_integration.gitlab_enterprise"
subcategory: ""
description: "code_base_integration.gitlab_enterprise for xcsh_code_base_integration."
xcsh_docs: {"aliases": [], "body_bytes": 2663, "body_sha256": "sha256:2e6d6ee1cd88c34450c21cab8a8dedeb6dde1831d625524895da5d633b2d59f4", "canonical_id": "xcsh-docs:resources:code_base_integration:properties:code_base_integration:gitlab_enterprise", "child_ids": ["xcsh-docs:resources:code_base_integration:properties:code_base_integration:gitlab_enterprise:access_token"], "collection_id": "xcsh-docs:resources:code_base_integration:collection", "completeness": "complete", "id": "xcsh-docs:resources:code_base_integration:properties:code_base_integration:gitlab_enterprise", "parent_id": "xcsh-docs:resources:code_base_integration:properties:code_base_integration", "path": "docs/guides/resources--code_base_integration--properties--code_base_integration--gitlab_enterprise.md", "provider_name": "code_base_integration", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["code_base_integration", "gitlab_enterprise"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/code_base_integration/properties/code_base_integration/gitlab_enterprise/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "code_base_integration.gitlab_enterprise for xcsh_code_base_integration.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["code_base_integrationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# code_base_integration.gitlab_enterprise

Breadcrumbs:

- [xcsh_code_base_integration](../resources/code_base_integration.md)
- [Property reference](resources--code_base_integration--reference.md)
- [code_base_integration](resources--code_base_integration--properties--code_base_integration.md)
- code_base_integration.gitlab_enterprise

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for gitlab enterprise.

Provider validators and defaults (from schema source):

```go
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
gitlab_enterprise {
  # Configure direct properties listed below.
}
```

## Direct properties

- [access_token](resources--code_base_integration--properties--code_base_integration--gitlab_enterprise--access_token.md): complete subsection reference.

<a id="schema-code_base_integration--gitlab_enterprise--url"></a>

### url property

Type: `"string"`. Optional.

GitLab URL. URL or URI reference

Upstream description:

URL or URI reference

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
  stringvalidator.RegexMatches(regexp.MustCompile(`^(https?|ftp)://[^\s/$.?#].[^\s]*$`),
    ""),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "network",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.9,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

## Next pages

- [code_base_integration.gitlab_enterprise.access_token](resources--code_base_integration--properties--code_base_integration--gitlab_enterprise--access_token.md)
- [code_base_integration](resources--code_base_integration--properties--code_base_integration.md)
- [xcsh_code_base_integration](../resources/code_base_integration.md)
