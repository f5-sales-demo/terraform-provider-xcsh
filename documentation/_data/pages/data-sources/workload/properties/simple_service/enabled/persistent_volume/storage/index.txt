---
page_title: "simple_service.enabled.persistent_volume.storage"
subcategory: "Container"
description: "Persistent storage configuration is used to configure Persistent Volume Claim (PVC)"
xcsh_docs: {"aliases": ["simple service enabled persistent volume storage"], "body_bytes": 4640, "body_sha256": "sha256:1ed469cef3a58b9e9bb613b403adadd71434dc26bbda95f1311dd7044cc335e7", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:workload:properties:simple_service:enabled:persistent_volume:storage:default"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:simple_service:enabled:persistent_volume:storage", "parent_id": "xcsh-docs:data-sources:workload:properties:simple_service:enabled:persistent_volume", "path": "documentation/data-sources/workload/properties/simple_service/enabled/persistent_volume/storage/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-3222112122312030-0100131331212110-3311100330320223-2221122302332032-2221333323323122-0210301032111313-0001233130123103-3111000232322233", "registry_path": "docs/guides/data-sources--workload--reference--group-016.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["simple_service", "enabled", "persistent_volume", "storage"], "schema_version": 1, "sections": [{"aliases": ["simple service enabled persistent volume storage access mode"], "anchor": "schema-simple_service--enabled--persistent_volume--storage--access_mode", "description": "Persistence storage access mode is used to configure access mode for persistent storage - ACCESS_MODE_READ_WRITE_ONCE: Read Write Once Read Write Once is used to mount persistent storage in read/write mode to exactly 1 host - ACCESS_MODE_READ_WRITE_MANY: Read Write Many Read Write Many is used to mount persistent", "document_id": "xcsh-docs:data-sources:workload:properties:simple_service:enabled:persistent_volume:storage", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["simple_service", "enabled", "persistent_volume", "storage", "access_mode"], "syntax": "attribute", "type": "string"}, {"aliases": ["simple service enabled persistent volume storage class name"], "anchor": "schema-simple_service--enabled--persistent_volume--storage--class_name", "description": "Exclusive with Use the specified class name.", "document_id": "xcsh-docs:data-sources:workload:properties:simple_service:enabled:persistent_volume:storage", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["simple_service", "enabled", "persistent_volume", "storage", "class_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["simple service enabled persistent volume storage default"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:workload:properties:simple_service:enabled:persistent_volume:storage:default", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["simple_service", "enabled", "persistent_volume", "storage", "default"], "syntax": "attribute", "type": "object"}, {"aliases": ["simple service enabled persistent volume storage storage size"], "anchor": "schema-simple_service--enabled--persistent_volume--storage--storage_size", "description": "Size in GiB of the persistent storage.", "document_id": "xcsh-docs:data-sources:workload:properties:simple_service:enabled:persistent_volume:storage", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["simple_service", "enabled", "persistent_volume", "storage", "storage_size"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/simple_service/enabled/persistent_volume/storage/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Persistent storage configuration is used to configure Persistent Volume Claim (PVC)", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["workloadCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# simple_service.enabled.persistent_volume.storage

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/)
- [simple_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/simple_service/)
- [simple_service.enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/simple_service/enabled/)
- [simple_service.enabled.persistent_volume](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/simple_service/enabled/persistent_volume/)
- simple_service.enabled.persistent_volume.storage

<a id="section"></a>

Type: `"single"`. Computed.

Persistent storage configuration is used to configure Persistent Volume Claim (PVC).

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-class_name_choice": "[\"class_name\",\"default\"]"
}
```

## Direct properties

<a id="schema-simple_service--enabled--persistent_volume--storage--access_mode"></a>

### access_mode property

Type: `"string"`. Computed.

\[Enum:
ACCESS\_MODE\_READ\_WRITE\_ONCE|ACCESS\_MODE\_READ\_WRITE\_MANY|ACCESS\_MODE\_READ\_ONLY\_MANY\]
Persistence storage access mode is used to configure access mode for persistent storage -
ACCESS\_MODE\_READ\_WRITE\_ONCE: Read Write Once Read Write Once is used to mount persistent storage
in read/write mode to exactly 1 host - ACCESS\_MODE\_READ\_WRITE\_MANY: Read Write Many Read Write
Many is used.. Possible values are \`ACCESS\_MODE\_READ\_WRITE\_ONCE\`,
\`ACCESS\_MODE\_READ\_WRITE\_MANY\`, \`ACCESS\_MODE\_READ\_ONLY\_MANY\`. Defaults to
\`ACCESS\_MODE\_READ\_WRITE\_ONCE\`.

Additional upstream details:

Persistence storage access mode is used to configure access mode for persistent storage

&#8203;- ACCESS\_MODE\_READ\_WRITE\_ONCE: Read Write Once

Read Write Once is used to mount persistent storage in read/write mode to exactly 1 host &#8203;-
ACCESS\_MODE\_READ\_WRITE\_MANY: Read Write Many

Read Write Many is used to mount persistent storage in read/write mode to many hosts &#8203;-
ACCESS\_MODE\_READ\_ONLY\_MANY: Read Only Many

Read Only Many is used to mount persistent storage in read-only mode to many hosts.

Receipt-pinned upstream constraints:

```json
{
  "default": "ACCESS_MODE_READ_WRITE_ONCE",
  "enum": [
    "ACCESS_MODE_READ_WRITE_ONCE",
    "ACCESS_MODE_READ_WRITE_MANY",
    "ACCESS_MODE_READ_ONLY_MANY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-simple_service--enabled--persistent_volume--storage--class_name"></a>

### class_name property

Type: `"string"`. Computed.

Exclusive with \[default\] Use the specified class name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [default](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/simple_service/enabled/persistent_volume/storage/default/): complete subsection reference.

<a id="schema-simple_service--enabled--persistent_volume--storage--storage_size"></a>

### storage_size property

Type: `"number"`. Computed.

Size (in GiB). Size in GiB of the persistent storage.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.double.gte": "0.004",
    "ves.io.schema.rules.double.lte": "256",
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.double.gte": "0.004",
    "ves.io.schema.rules.double.lte": "256",
    "ves.io.schema.rules.message.required": "true"
  }
}
```
