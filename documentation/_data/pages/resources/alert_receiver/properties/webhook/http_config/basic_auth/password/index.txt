---
page_title: "webhook.http_config.basic_auth.password"
subcategory: ""
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["webhook http config basic auth password"], "body_bytes": 2699, "body_sha256": "sha256:139498def713c844ff93c9caa279c244a3ca1d24baaf162a2f2319013c84fe23", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:resources:alert_receiver:properties:webhook:http_config:basic_auth:password:blindfold_secret_info", "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:basic_auth:password:clear_secret_info"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:alert_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:basic_auth:password", "parent_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:basic_auth", "path": "documentation/resources/alert_receiver/properties/webhook/http_config/basic_auth/password/index.md", "product": "distributed-cloud", "provider_name": "alert_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1213010003022202-2021203100323000-3210013101220001-0210232033030323-0213011001010312-3230000023103311-2021210133023200-0211113303000233", "registry_path": "docs/guides/resources--alert_receiver--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "webhook.http_config.basic_auth.password:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:basic_auth:password:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "webhook.http_config.basic_auth.password:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:basic_auth:password:clear_secret_info", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["webhook", "http_config", "basic_auth", "password"], "schema_version": 1, "sections": [{"aliases": ["blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:basic_auth:password:blindfold_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-webhook--http_config--basic_auth--password--blindfold_secret_info--location", "enforcement": "provider-schema", "group": "webhook.http_config.basic_auth.password.blindfold_secret_info:RequiredObjectAttributes:location", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:basic_auth:password:blindfold_secret_info", "type": "requires"}], "schema_path": ["webhook", "http_config", "basic_auth", "password", "blindfold_secret_info"], "syntax": "block", "type": "object"}, {"aliases": ["clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:basic_auth:password:clear_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-webhook--http_config--basic_auth--password--clear_secret_info--url", "enforcement": "provider-schema", "group": "webhook.http_config.basic_auth.password.clear_secret_info:RequiredObjectAttributes:url", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:basic_auth:password:clear_secret_info", "type": "requires"}], "schema_path": ["webhook", "http_config", "basic_auth", "password", "clear_secret_info"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/alert_receiver/properties/webhook/http_config/basic_auth/password/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["alert_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# webhook.http_config.basic_auth.password

Breadcrumbs:

- [xcsh_alert_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/)
- [webhook](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/)
- [webhook.http_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/)
- [webhook.http_config.basic_auth](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/basic_auth/)
- webhook.http_config.basic_auth.password

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
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
password {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/basic_auth/password/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/basic_auth/password/clear_secret_info/): complete subsection reference.

## Next pages

- [webhook.http_config.basic_auth.password.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/basic_auth/password/blindfold_secret_info/)
- [webhook.http_config.basic_auth.password.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/basic_auth/password/clear_secret_info/)
- [webhook.http_config.basic_auth](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/basic_auth/)
- [xcsh_alert_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/)
