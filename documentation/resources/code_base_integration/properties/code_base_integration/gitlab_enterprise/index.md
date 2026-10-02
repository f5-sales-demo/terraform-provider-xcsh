---
page_title: "code_base_integration.gitlab_enterprise"
subcategory: ""
description: "Configuration parameter for gitlab enterprise."
xcsh_docs: {"aliases": ["code base integration gitlab enterprise"], "body_bytes": 3117, "body_sha256": "sha256:0b9fd0768394556f47294567edd4d3b0b551f77ad454b4cc0c40d66b2850453f", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:code_base_integration:properties:code_base_integration:gitlab_enterprise:access_token"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:code_base_integration:collection", "completeness": "complete", "id": "xcsh-docs:resources:code_base_integration:properties:code_base_integration:gitlab_enterprise", "parent_id": "xcsh-docs:resources:code_base_integration:properties:code_base_integration", "path": "documentation/resources/code_base_integration/properties/code_base_integration/gitlab_enterprise/index.md", "product": "distributed-cloud", "provider_name": "code_base_integration", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2300101201031033-0233121100211322-0323002120210212-3020213213320323-2323322312133130-2103132210131111-2130333232002130-1220301121211123", "registry_path": "docs/guides/resources--code_base_integration--reference--group-001.md", "relationships": [{"anchor": "schema-code_base_integration--gitlab_enterprise--url", "enforcement": "provider-schema", "group": "code_base_integration.gitlab_enterprise:RequiredObjectAttributes:url", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:code_base_integration:properties:code_base_integration:gitlab_enterprise", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["code_base_integration", "gitlab_enterprise"], "schema_version": 1, "sections": [{"aliases": ["access token"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:resources:code_base_integration:properties:code_base_integration:gitlab_enterprise:access_token", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "code_base_integration.gitlab_enterprise.access_token:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:code_base_integration:properties:code_base_integration:gitlab_enterprise:access_token:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "code_base_integration.gitlab_enterprise.access_token:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:code_base_integration:properties:code_base_integration:gitlab_enterprise:access_token:clear_secret_info", "type": "conflicts"}], "schema_path": ["code_base_integration", "gitlab_enterprise", "access_token"], "syntax": "block", "type": "object"}, {"aliases": ["url"], "anchor": "schema-code_base_integration--gitlab_enterprise--url", "description": "URL or URI reference", "document_id": "xcsh-docs:resources:code_base_integration:properties:code_base_integration:gitlab_enterprise", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["code_base_integration", "gitlab_enterprise", "url"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/code_base_integration/properties/code_base_integration/gitlab_enterprise/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Configuration parameter for gitlab enterprise.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["code_base_integrationCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# code_base_integration.gitlab_enterprise

Breadcrumbs:

- [xcsh_code_base_integration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/code_base_integration/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/code_base_integration/properties/)
- [code_base_integration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/code_base_integration/properties/code_base_integration/)
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

- [access_token](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/code_base_integration/properties/code_base_integration/gitlab_enterprise/access_token/): complete subsection reference.

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

- [code_base_integration.gitlab_enterprise.access_token](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/code_base_integration/properties/code_base_integration/gitlab_enterprise/access_token/)
- [code_base_integration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/code_base_integration/properties/code_base_integration/)
- [xcsh_code_base_integration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/code_base_integration/)
