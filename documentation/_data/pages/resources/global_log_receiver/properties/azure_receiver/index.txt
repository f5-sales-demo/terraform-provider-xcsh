---
page_title: "azure_receiver"
subcategory: ""
description: "Azure Blob Configuration for Global Log Receiver."
xcsh_docs: {"aliases": ["azure receiver"], "body_bytes": 3139, "body_sha256": "sha256:485a619820e422b36d4530b3a6d1d0eae212bba0170af282d87f13ec3797bc4b", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:resources:global_log_receiver:properties:azure_receiver:batch", "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:compression", "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:connection_string", "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:filename_options"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver", "parent_id": "xcsh-docs:resources:global_log_receiver:reference", "path": "documentation/resources/global_log_receiver/properties/azure_receiver/index.md", "product": "distributed-cloud", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-3103303230212132-2010213002200222-1302002300121311-3323011111112102-2333203023000310-3220320102133222-1020130321112320-0303201103200031", "registry_path": "docs/guides/resources--global_log_receiver--reference--group-001.md", "relationships": [{"anchor": "schema-azure_receiver--container_name", "enforcement": "provider-schema", "group": "azure_receiver:RequiredObjectAttributes:container_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["azure_receiver"], "schema_version": 1, "sections": [{"aliases": ["azure receiver batch"], "anchor": "section", "description": "Batch OPTIONS allow tuning for how batches of logs are sent to an endpoint.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:batch", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-azure_receiver--batch--max_bytes", "enforcement": "provider-schema", "group": "azure_receiver.batch:ConflictingObjectAttributes:max_bytes,max_bytes_disabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:batch", "type": "conflicts"}, {"anchor": "schema-azure_receiver--batch--max_events", "enforcement": "provider-schema", "group": "azure_receiver.batch:ConflictingObjectAttributes:max_events,max_events_disabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:batch", "type": "conflicts"}, {"anchor": "schema-azure_receiver--batch--timeout_seconds", "enforcement": "provider-schema", "group": "azure_receiver.batch:ConflictingObjectAttributes:timeout_seconds,timeout_seconds_default", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:batch", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "azure_receiver.batch:ConflictingObjectAttributes:max_bytes,max_bytes_disabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:batch:max_bytes_disabled", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "azure_receiver.batch:ConflictingObjectAttributes:max_events,max_events_disabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:batch:max_events_disabled", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "azure_receiver.batch:ConflictingObjectAttributes:timeout_seconds,timeout_seconds_default", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:batch:timeout_seconds_default", "type": "conflicts"}], "schema_path": ["azure_receiver", "batch"], "syntax": "block", "type": "object"}, {"aliases": ["azure receiver compression"], "anchor": "section", "description": "Compression Type.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:compression", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "azure_receiver.compression:ConflictingObjectAttributes:compression_default,compression_gzip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:compression:compression_default", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "azure_receiver.compression:ConflictingObjectAttributes:compression_default,compression_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:compression:compression_default", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "azure_receiver.compression:ConflictingObjectAttributes:compression_default,compression_gzip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:compression:compression_gzip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "azure_receiver.compression:ConflictingObjectAttributes:compression_gzip,compression_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:compression:compression_gzip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "azure_receiver.compression:ConflictingObjectAttributes:compression_default,compression_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:compression:compression_none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "azure_receiver.compression:ConflictingObjectAttributes:compression_gzip,compression_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:compression:compression_none", "type": "conflicts"}], "schema_path": ["azure_receiver", "compression"], "syntax": "block", "type": "object"}, {"aliases": ["azure receiver connection string"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:connection_string", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "azure_receiver.connection_string:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:connection_string:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "azure_receiver.connection_string:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:connection_string:clear_secret_info", "type": "conflicts"}], "schema_path": ["azure_receiver", "connection_string"], "syntax": "block", "type": "object"}, {"aliases": ["azure receiver container name"], "anchor": "schema-azure_receiver--container_name", "description": "Container Name is the name of the container into which logs should be stored.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["azure_receiver", "container_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["azure receiver filename options"], "anchor": "section", "description": "Filename OPTIONS allow customization of filename and folder paths used by a destination endpoint bucket or file.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:filename_options", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-azure_receiver--filename_options--custom_folder", "enforcement": "provider-schema", "group": "azure_receiver.filename_options:ConflictingObjectAttributes:custom_folder,log_type_folder", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:filename_options", "type": "conflicts"}, {"anchor": "schema-azure_receiver--filename_options--custom_folder", "enforcement": "provider-schema", "group": "azure_receiver.filename_options:ConflictingObjectAttributes:custom_folder,no_folder", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:filename_options", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "azure_receiver.filename_options:ConflictingObjectAttributes:custom_folder,log_type_folder", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:filename_options:log_type_folder", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "azure_receiver.filename_options:ConflictingObjectAttributes:log_type_folder,no_folder", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:filename_options:log_type_folder", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "azure_receiver.filename_options:ConflictingObjectAttributes:custom_folder,no_folder", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:filename_options:no_folder", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "azure_receiver.filename_options:ConflictingObjectAttributes:log_type_folder,no_folder", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:filename_options:no_folder", "type": "conflicts"}], "schema_path": ["azure_receiver", "filename_options"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/azure_receiver/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Azure Blob Configuration for Global Log Receiver.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# azure_receiver

Breadcrumbs:

- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/)
- azure_receiver

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Azure Blob Configuration for Global Log Receiver.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("container_name")}
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
azure_receiver {
  # Configure direct properties listed below.
}
```

## Direct properties

- [batch](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/azure_receiver/batch/): complete subsection reference.

- [compression](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/azure_receiver/compression/): complete subsection reference.

- [connection_string](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/azure_receiver/connection_string/): complete subsection reference.

<a id="schema-azure_receiver--container_name"></a>

### container_name property

Type: `"string"`. Optional.

Container Name is the name of the container into which logs should be stored.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(3, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 63,
  "minLength": 3,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minLength": 3,
    "pattern": "^[A-Za-z0-9][A-Za-z0-9-]+[A-Za-z0-9]$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "63",
    "ves.io.schema.rules.string.min_len": "3",
    "ves.io.schema.rules.string.pattern": "^[A-Za-z0-9][A-Za-z0-9-]+[A-Za-z0-9]$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "63",
    "ves.io.schema.rules.string.min_len": "3",
    "ves.io.schema.rules.string.pattern": "^[A-Za-z0-9][A-Za-z0-9-]+[A-Za-z0-9]$"
  }
}
```

- [filename_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/azure_receiver/filename_options/): complete subsection reference.
