---
page_title: "service.volumes.persistent_volume.storage"
subcategory: "Container"
description: "Persistent storage configuration is used to configure Persistent Volume Claim (PVC)"
xcsh_docs: {"aliases": ["service volumes persistent volume storage"], "body_bytes": 4654, "body_sha256": "sha256:75e371943bb15319dd115eafa00253f33ae433d76a6cd2c120fee3ddb11a27b1", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:workload:properties:service:volumes:persistent_volume:storage:default"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:service:volumes:persistent_volume:storage", "parent_id": "xcsh-docs:resources:workload:properties:service:volumes:persistent_volume", "path": "documentation/resources/workload/properties/service/volumes/persistent_volume/storage/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-2211031333323330-2003301230021101-0212101233333302-2000330220323032-2030230333012232-0130320200232322-1112231131220111-2033313302022312", "registry_path": "docs/guides/resources--workload--reference--group-017.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["service", "volumes", "persistent_volume", "storage"], "schema_version": 1, "sections": [{"aliases": ["service volumes persistent volume storage access mode"], "anchor": "schema-service--volumes--persistent_volume--storage--access_mode", "description": "Persistence storage access mode is used to configure access mode for persistent storage - ACCESS_MODE_READ_WRITE_ONCE: Read Write Once Read Write Once is used to mount persistent storage in read/write mode to exactly 1 host - ACCESS_MODE_READ_WRITE_MANY: Read Write Many Read Write Many is used to mount persistent", "document_id": "xcsh-docs:resources:workload:properties:service:volumes:persistent_volume:storage", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "volumes", "persistent_volume", "storage", "access_mode"], "syntax": "attribute", "type": "string"}, {"aliases": ["service volumes persistent volume storage class name"], "anchor": "schema-service--volumes--persistent_volume--storage--class_name", "description": "Exclusive with Use the specified class name.", "document_id": "xcsh-docs:resources:workload:properties:service:volumes:persistent_volume:storage", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "volumes", "persistent_volume", "storage", "class_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["service volumes persistent volume storage default"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:workload:properties:service:volumes:persistent_volume:storage:default", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "volumes", "persistent_volume", "storage", "default"], "syntax": "attribute", "type": "object"}, {"aliases": ["service volumes persistent volume storage storage size"], "anchor": "schema-service--volumes--persistent_volume--storage--storage_size", "description": "Size in GiB of the persistent storage.", "document_id": "xcsh-docs:resources:workload:properties:service:volumes:persistent_volume:storage", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "volumes", "persistent_volume", "storage", "storage_size"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/service/volumes/persistent_volume/storage/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Persistent storage configuration is used to configure Persistent Volume Claim (PVC)", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["workloadCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.volumes.persistent_volume.storage

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/)
- [service.volumes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/volumes/)
- [service.volumes.persistent_volume](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/volumes/persistent_volume/)
- service.volumes.persistent_volume.storage

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
storage {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-service--volumes--persistent_volume--storage--access_mode"></a>

### access_mode property

Type: `"string"`. Optional.

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

<a id="schema-service--volumes--persistent_volume--storage--class_name"></a>

### class_name property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

- [default](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/volumes/persistent_volume/storage/default/): complete subsection reference.

<a id="schema-service--volumes--persistent_volume--storage--storage_size"></a>

### storage_size property

Type: `"number"`. Optional.

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
