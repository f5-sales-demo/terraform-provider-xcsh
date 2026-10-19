---
page_title: "azure_receiver"
subcategory: ""
description: "Azure Blob Configuration for Global Log Receiver."
xcsh_docs: {"aliases": ["azure receiver"], "body_bytes": 2681, "body_sha256": "sha256:4d150fd410f498c3b430aac39008508b65a6ae8defdab99f66a1e5d7553ed023", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:data-sources:global_log_receiver:properties:azure_receiver:batch", "xcsh-docs:data-sources:global_log_receiver:properties:azure_receiver:compression", "xcsh-docs:data-sources:global_log_receiver:properties:azure_receiver:connection_string", "xcsh-docs:data-sources:global_log_receiver:properties:azure_receiver:filename_options"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:global_log_receiver:properties:azure_receiver", "parent_id": "xcsh-docs:data-sources:global_log_receiver:reference", "path": "documentation/data-sources/global_log_receiver/properties/azure_receiver/index.md", "product": "distributed-cloud", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-2121301220330030-1221312033213003-0320321302002303-2133212312111123-3330122101022123-0133210113001032-2201030003210013-3211030023023312", "registry_path": "docs/guides/data-sources--global_log_receiver--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["azure_receiver"], "schema_version": 1, "sections": [{"aliases": ["azure receiver batch"], "anchor": "section", "description": "Batch OPTIONS allow tuning for how batches of logs are sent to an endpoint.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:azure_receiver:batch", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["azure_receiver", "batch"], "syntax": "attribute", "type": "object"}, {"aliases": ["azure receiver compression"], "anchor": "section", "description": "Compression Type.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:azure_receiver:compression", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["azure_receiver", "compression"], "syntax": "attribute", "type": "object"}, {"aliases": ["azure receiver connection string"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:azure_receiver:connection_string", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["azure_receiver", "connection_string"], "syntax": "attribute", "type": "object"}, {"aliases": ["azure receiver container name"], "anchor": "schema-azure_receiver--container_name", "description": "Container Name is the name of the container into which logs should be stored.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:azure_receiver", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["azure_receiver", "container_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["azure receiver filename options"], "anchor": "section", "description": "Filename OPTIONS allow customization of filename and folder paths used by a destination endpoint bucket or file.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:azure_receiver:filename_options", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["azure_receiver", "filename_options"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/global_log_receiver/properties/azure_receiver/index.txt", "spec_pin_digest": "sha256:ea0d80973b5efe05eb3575be6beda9c96b86701d264dc1aee3972e6a23ebb3fb", "summary": "Azure Blob Configuration for Global Log Receiver.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.3", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "6ee1e85b2e917d58078f38a270b790077b3fba12"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# azure_receiver

Breadcrumbs:

- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/)
- azure_receiver

<a id="section"></a>

Type: `"single"`. Computed.

Azure Blob Configuration for Global Log Receiver.

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

- [batch](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/azure_receiver/batch/): complete subsection reference.

- [compression](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/azure_receiver/compression/): complete subsection reference.

- [connection_string](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/azure_receiver/connection_string/): complete subsection reference.

<a id="schema-azure_receiver--container_name"></a>

### container_name property

Type: `"string"`. Computed.

Container Name is the name of the container into which logs should be stored.

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
      "validatedAt": "2026-10-17T10:10:42+00:00"
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

- [filename_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/azure_receiver/filename_options/): complete subsection reference.
