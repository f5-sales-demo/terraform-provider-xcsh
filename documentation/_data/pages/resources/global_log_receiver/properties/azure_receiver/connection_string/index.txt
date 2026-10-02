---
page_title: "azure_receiver.connection_string"
subcategory: ""
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["azure receiver connection string"], "body_bytes": 2391, "body_sha256": "sha256:96002cd7180eb318a320e74ea826f184ebca3294e87f30b3abdf43bbbbc6685b", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:resources:global_log_receiver:properties:azure_receiver:connection_string:blindfold_secret_info", "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:connection_string:clear_secret_info"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:connection_string", "parent_id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver", "path": "documentation/resources/global_log_receiver/properties/azure_receiver/connection_string/index.md", "product": "distributed-cloud", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3321122021101233-0323300002001002-2212120302013221-2231302210203201-3121310031113020-0111200312123321-2120030012123003-0212312323022131", "registry_path": "docs/guides/resources--global_log_receiver--reference--group-002.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "azure_receiver.connection_string:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:connection_string:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "azure_receiver.connection_string:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:connection_string:clear_secret_info", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["azure_receiver", "connection_string"], "schema_version": 1, "sections": [{"aliases": ["blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:connection_string:blindfold_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-azure_receiver--connection_string--blindfold_secret_info--location", "enforcement": "provider-schema", "group": "azure_receiver.connection_string.blindfold_secret_info:RequiredObjectAttributes:location", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:connection_string:blindfold_secret_info", "type": "requires"}], "schema_path": ["azure_receiver", "connection_string", "blindfold_secret_info"], "syntax": "block", "type": "object"}, {"aliases": ["clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:connection_string:clear_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-azure_receiver--connection_string--clear_secret_info--url", "enforcement": "provider-schema", "group": "azure_receiver.connection_string.clear_secret_info:RequiredObjectAttributes:url", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:connection_string:clear_secret_info", "type": "requires"}], "schema_path": ["azure_receiver", "connection_string", "clear_secret_info"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/azure_receiver/connection_string/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# azure_receiver.connection_string

Breadcrumbs:

- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/)
- [azure_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/azure_receiver/)
- azure_receiver.connection_string

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
connection_string {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/azure_receiver/connection_string/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/azure_receiver/connection_string/clear_secret_info/): complete subsection reference.

## Next pages

- [azure_receiver.connection_string.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/azure_receiver/connection_string/blindfold_secret_info/)
- [azure_receiver.connection_string.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/azure_receiver/connection_string/clear_secret_info/)
- [azure_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/azure_receiver/)
- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/)
