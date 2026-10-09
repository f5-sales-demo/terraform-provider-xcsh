---
page_title: "webhook.http_config.auth_token.token"
subcategory: ""
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["webhook http config auth token token"], "body_bytes": 1995, "body_sha256": "sha256:03bcd2004406b948f02868a1665ef3f68b41a4a8da8b65b0230db24ef178e50c", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:resources:alert_receiver:properties:webhook:http_config:auth_token:token:blindfold_secret_info", "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:auth_token:token:clear_secret_info"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:alert_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:auth_token:token", "parent_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:auth_token", "path": "documentation/resources/alert_receiver/properties/webhook/http_config/auth_token/token/index.md", "product": "distributed-cloud", "provider_name": "alert_receiver", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-3031221001332212-0302220212232310-0311223033332023-0310333300020220-0020001121203312-0132132131203011-3103301002202003-3320311220110022", "registry_path": "docs/guides/resources--alert_receiver--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "webhook.http_config.auth_token.token:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:auth_token:token:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "webhook.http_config.auth_token.token:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:auth_token:token:clear_secret_info", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["webhook", "http_config", "auth_token", "token"], "schema_version": 1, "sections": [{"aliases": ["webhook http config auth token token blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:auth_token:token:blindfold_secret_info", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-webhook--http_config--auth_token--token--blindfold_secret_info--location", "enforcement": "provider-schema", "group": "webhook.http_config.auth_token.token.blindfold_secret_info:RequiredObjectAttributes:location", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:auth_token:token:blindfold_secret_info", "type": "requires"}], "schema_path": ["webhook", "http_config", "auth_token", "token", "blindfold_secret_info"], "syntax": "block", "type": "object"}, {"aliases": ["webhook http config auth token token clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:auth_token:token:clear_secret_info", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-webhook--http_config--auth_token--token--clear_secret_info--url", "enforcement": "provider-schema", "group": "webhook.http_config.auth_token.token.clear_secret_info:RequiredObjectAttributes:url", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:auth_token:token:clear_secret_info", "type": "requires"}], "schema_path": ["webhook", "http_config", "auth_token", "token", "clear_secret_info"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/alert_receiver/properties/webhook/http_config/auth_token/token/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["alert_receiverCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# webhook.http_config.auth_token.token

Breadcrumbs:

- [xcsh_alert_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/)
- [webhook](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/)
- [webhook.http_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/)
- [webhook.http_config.auth_token](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/auth_token/)
- webhook.http_config.auth_token.token

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
token {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/auth_token/token/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/auth_token/token/clear_secret_info/): complete subsection reference.
