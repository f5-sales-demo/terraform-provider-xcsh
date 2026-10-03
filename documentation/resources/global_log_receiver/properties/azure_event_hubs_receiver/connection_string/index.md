---
page_title: "azure_event_hubs_receiver.connection_string"
subcategory: ""
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["azure event hubs receiver connection string"], "body_bytes": 2523, "body_sha256": "sha256:51e855b393f655ee9336aa301a9a5c3baac4b6e3d78b5214a9ccaf7f2d150bc4", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:resources:global_log_receiver:properties:azure_event_hubs_receiver:connection_string:blindfold_secret_info", "xcsh-docs:resources:global_log_receiver:properties:azure_event_hubs_receiver:connection_string:clear_secret_info"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:azure_event_hubs_receiver:connection_string", "parent_id": "xcsh-docs:resources:global_log_receiver:properties:azure_event_hubs_receiver", "path": "documentation/resources/global_log_receiver/properties/azure_event_hubs_receiver/connection_string/index.md", "product": "distributed-cloud", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2301122011100222-3311101033203332-0002310332001210-3211221210133220-0301210231010033-1331133323130121-0223213230121233-1000101013232122", "registry_path": "docs/guides/resources--global_log_receiver--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "azure_event_hubs_receiver.connection_string:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:azure_event_hubs_receiver:connection_string:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "azure_event_hubs_receiver.connection_string:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:azure_event_hubs_receiver:connection_string:clear_secret_info", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["azure_event_hubs_receiver", "connection_string"], "schema_version": 1, "sections": [{"aliases": ["azure event hubs receiver connection string blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:azure_event_hubs_receiver:connection_string:blindfold_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-azure_event_hubs_receiver--connection_string--blindfold_secret_info--location", "enforcement": "provider-schema", "group": "azure_event_hubs_receiver.connection_string.blindfold_secret_info:RequiredObjectAttributes:location", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:azure_event_hubs_receiver:connection_string:blindfold_secret_info", "type": "requires"}], "schema_path": ["azure_event_hubs_receiver", "connection_string", "blindfold_secret_info"], "syntax": "block", "type": "object"}, {"aliases": ["azure event hubs receiver connection string clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:azure_event_hubs_receiver:connection_string:clear_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-azure_event_hubs_receiver--connection_string--clear_secret_info--url", "enforcement": "provider-schema", "group": "azure_event_hubs_receiver.connection_string.clear_secret_info:RequiredObjectAttributes:url", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:azure_event_hubs_receiver:connection_string:clear_secret_info", "type": "requires"}], "schema_path": ["azure_event_hubs_receiver", "connection_string", "clear_secret_info"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/azure_event_hubs_receiver/connection_string/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# azure_event_hubs_receiver.connection_string

Breadcrumbs:

- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/)
- [azure_event_hubs_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/azure_event_hubs_receiver/)
- azure_event_hubs_receiver.connection_string

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

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/azure_event_hubs_receiver/connection_string/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/azure_event_hubs_receiver/connection_string/clear_secret_info/): complete subsection reference.

## Next pages

- [azure_event_hubs_receiver.connection_string.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/azure_event_hubs_receiver/connection_string/blindfold_secret_info/)
- [azure_event_hubs_receiver.connection_string.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/azure_event_hubs_receiver/connection_string/clear_secret_info/)
- [azure_event_hubs_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/azure_event_hubs_receiver/)
- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/)
