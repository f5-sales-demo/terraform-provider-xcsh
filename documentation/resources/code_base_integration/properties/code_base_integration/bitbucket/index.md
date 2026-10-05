---
page_title: "code_base_integration.bitbucket"
subcategory: ""
description: "BitBucket Cloud Integration."
xcsh_docs: {"aliases": ["code base integration bitbucket"], "body_bytes": 2777, "body_sha256": "sha256:3482249131f2923d6ce12fb908bbfc31f1f5b0ea84355798ed6ec771f11acdd6", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:code_base_integration:properties:code_base_integration:bitbucket:passwd"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:code_base_integration:collection", "completeness": "complete", "id": "xcsh-docs:resources:code_base_integration:properties:code_base_integration:bitbucket", "parent_id": "xcsh-docs:resources:code_base_integration:properties:code_base_integration", "path": "documentation/resources/code_base_integration/properties/code_base_integration/bitbucket/index.md", "product": "distributed-cloud", "provider_name": "code_base_integration", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-3331021320323120-0021121133102300-3331011033013300-0131212023232003-3131102311220212-2030030132023320-0333230023300002-0133212120120111", "registry_path": "docs/guides/resources--code_base_integration--reference--group-001.md", "relationships": [{"anchor": "schema-code_base_integration--bitbucket--username", "enforcement": "provider-schema", "group": "code_base_integration.bitbucket:RequiredObjectAttributes:username", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:code_base_integration:properties:code_base_integration:bitbucket", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["code_base_integration", "bitbucket"], "schema_version": 1, "sections": [{"aliases": ["code base integration bitbucket passwd"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:resources:code_base_integration:properties:code_base_integration:bitbucket:passwd", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "code_base_integration.bitbucket.passwd:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:code_base_integration:properties:code_base_integration:bitbucket:passwd:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "code_base_integration.bitbucket.passwd:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:code_base_integration:properties:code_base_integration:bitbucket:passwd:clear_secret_info", "type": "conflicts"}], "schema_path": ["code_base_integration", "bitbucket", "passwd"], "syntax": "block", "type": "object"}, {"aliases": ["code base integration bitbucket username"], "anchor": "schema-code_base_integration--bitbucket--username", "description": "Human-readable name for the resource", "document_id": "xcsh-docs:resources:code_base_integration:properties:code_base_integration:bitbucket", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["code_base_integration", "bitbucket", "username"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/code_base_integration/properties/code_base_integration/bitbucket/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "BitBucket Cloud Integration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["code_base_integrationCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# code_base_integration.bitbucket

Breadcrumbs:

- [xcsh_code_base_integration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/code_base_integration/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/code_base_integration/properties/)
- [code_base_integration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/code_base_integration/properties/code_base_integration/)
- code_base_integration.bitbucket

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

BitBucket Cloud Integration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("username")}
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
bitbucket {
  # Configure direct properties listed below.
}
```

## Direct properties

- [passwd](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/code_base_integration/properties/code_base_integration/bitbucket/passwd/): complete subsection reference.

<a id="schema-code_base_integration--bitbucket--username"></a>

### username property

Type: `"string"`. Optional.

BitBucket Username. Human-readable name for the resource

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [code_base_integration.bitbucket.passwd](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/code_base_integration/properties/code_base_integration/bitbucket/passwd/)
- [code_base_integration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/code_base_integration/properties/code_base_integration/)
- [xcsh_code_base_integration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/code_base_integration/)
