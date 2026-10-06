---
page_title: "code_base_integration.gitlab_enterprise.access_token"
subcategory: ""
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["code base integration gitlab enterprise access token"], "body_bytes": 2023, "body_sha256": "sha256:247161eaf47e246e1980eacc632113306a119542a6c556f1bed12f9eb0361487", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:code_base_integration:properties:code_base_integration:gitlab_enterprise:access_token:blindfold_secret_info", "xcsh-docs:resources:code_base_integration:properties:code_base_integration:gitlab_enterprise:access_token:clear_secret_info"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:code_base_integration:collection", "completeness": "complete", "id": "xcsh-docs:resources:code_base_integration:properties:code_base_integration:gitlab_enterprise:access_token", "parent_id": "xcsh-docs:resources:code_base_integration:properties:code_base_integration:gitlab_enterprise", "path": "documentation/resources/code_base_integration/properties/code_base_integration/gitlab_enterprise/access_token/index.md", "product": "distributed-cloud", "provider_name": "code_base_integration", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-3212032121013322-0112030301122103-3330010000103200-1013100110132233-2322022303312123-1230311212332232-3222301012303332-1021320033031032", "registry_path": "docs/guides/resources--code_base_integration--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "code_base_integration.gitlab_enterprise.access_token:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:code_base_integration:properties:code_base_integration:gitlab_enterprise:access_token:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "code_base_integration.gitlab_enterprise.access_token:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:code_base_integration:properties:code_base_integration:gitlab_enterprise:access_token:clear_secret_info", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["code_base_integration", "gitlab_enterprise", "access_token"], "schema_version": 1, "sections": [{"aliases": ["code base integration gitlab enterprise access token blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:resources:code_base_integration:properties:code_base_integration:gitlab_enterprise:access_token:blindfold_secret_info", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-code_base_integration--gitlab_enterprise--access_token--blindfold_secret_info--location", "enforcement": "provider-schema", "group": "code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info:RequiredObjectAttributes:location", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:code_base_integration:properties:code_base_integration:gitlab_enterprise:access_token:blindfold_secret_info", "type": "requires"}], "schema_path": ["code_base_integration", "gitlab_enterprise", "access_token", "blindfold_secret_info"], "syntax": "block", "type": "object"}, {"aliases": ["code base integration gitlab enterprise access token clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:resources:code_base_integration:properties:code_base_integration:gitlab_enterprise:access_token:clear_secret_info", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-code_base_integration--gitlab_enterprise--access_token--clear_secret_info--url", "enforcement": "provider-schema", "group": "code_base_integration.gitlab_enterprise.access_token.clear_secret_info:RequiredObjectAttributes:url", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:code_base_integration:properties:code_base_integration:gitlab_enterprise:access_token:clear_secret_info", "type": "requires"}], "schema_path": ["code_base_integration", "gitlab_enterprise", "access_token", "clear_secret_info"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/code_base_integration/properties/code_base_integration/gitlab_enterprise/access_token/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["code_base_integrationCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# code_base_integration.gitlab_enterprise.access_token

Breadcrumbs:

- [xcsh_code_base_integration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/code_base_integration/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/code_base_integration/properties/)
- [code_base_integration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/code_base_integration/properties/code_base_integration/)
- [code_base_integration.gitlab_enterprise](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/code_base_integration/properties/code_base_integration/gitlab_enterprise/)
- code_base_integration.gitlab_enterprise.access_token

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
access_token {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/code_base_integration/properties/code_base_integration/gitlab_enterprise/access_token/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/code_base_integration/properties/code_base_integration/gitlab_enterprise/access_token/clear_secret_info/): complete subsection reference.
