---
page_title: "stateful_service.persistent_volumes.persistent_volume.storage"
subcategory: "Container"
description: "Persistent storage configuration is used to configure Persistent Volume Claim (PVC)"
xcsh_docs: {"aliases": ["stateful service persistent volumes persistent volume storage"], "body_bytes": 5842, "body_sha256": "sha256:1baf8c0892950b58c9a777a613a1b444bbd646bf5e8976654431c1624fb0e128", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:workload:properties:stateful_service:persistent_volumes:persistent_volume:storage:default"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:stateful_service:persistent_volumes:persistent_volume:storage", "parent_id": "xcsh-docs:resources:workload:properties:stateful_service:persistent_volumes:persistent_volume", "path": "documentation/resources/workload/properties/stateful_service/persistent_volumes/persistent_volume/storage/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-2303112111002111-0301211130322101-1303203032331132-1303120113203310-1303311030303232-2300030233330020-3201133032112123-2222323020330130", "registry_path": "docs/guides/resources--workload--reference--group-028.md", "relationships": [{"anchor": "schema-stateful_service--persistent_volumes--persistent_volume--storage--class_name", "enforcement": "provider-schema", "group": "stateful_service.persistent_volumes.persistent_volume.storage:ConflictingObjectAttributes:class_name,default", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:persistent_volumes:persistent_volume:storage", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "stateful_service.persistent_volumes.persistent_volume.storage:ConflictingObjectAttributes:class_name,default", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:persistent_volumes:persistent_volume:storage:default", "type": "conflicts"}, {"anchor": "schema-stateful_service--persistent_volumes--persistent_volume--storage--storage_size", "enforcement": "provider-schema", "group": "stateful_service.persistent_volumes.persistent_volume.storage:RequiredObjectAttributes:storage_size", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:persistent_volumes:persistent_volume:storage", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["stateful_service", "persistent_volumes", "persistent_volume", "storage"], "schema_version": 1, "sections": [{"aliases": ["stateful service persistent volumes persistent volume storage access mode"], "anchor": "schema-stateful_service--persistent_volumes--persistent_volume--storage--access_mode", "description": "Persistence storage access mode is used to configure access mode for persistent storage - ACCESS_MODE_READ_WRITE_ONCE: Read Write Once Read Write Once is used to mount persistent storage in read/write mode to exactly 1 host - ACCESS_MODE_READ_WRITE_MANY: Read Write Many Read Write Many is used to mount persistent", "document_id": "xcsh-docs:resources:workload:properties:stateful_service:persistent_volumes:persistent_volume:storage", "enum_extraction_complete": true, "enum_validators": [{"case_sensitive": true, "complete": true, "source": "ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf", "validator": "OneOf", "values": ["ACCESS_MODE_READ_ONLY_MANY", "ACCESS_MODE_READ_WRITE_MANY", "ACCESS_MODE_READ_WRITE_ONCE"], "version": 1}], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["stateful_service", "persistent_volumes", "persistent_volume", "storage", "access_mode"], "syntax": "attribute", "type": "string"}, {"aliases": ["stateful service persistent volumes persistent volume storage class name"], "anchor": "schema-stateful_service--persistent_volumes--persistent_volume--storage--class_name", "description": "Exclusive with Use the specified class name.", "document_id": "xcsh-docs:resources:workload:properties:stateful_service:persistent_volumes:persistent_volume:storage", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["stateful_service", "persistent_volumes", "persistent_volume", "storage", "class_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["stateful service persistent volumes persistent volume storage default"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:workload:properties:stateful_service:persistent_volumes:persistent_volume:storage:default", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["stateful_service", "persistent_volumes", "persistent_volume", "storage", "default"], "syntax": "attribute", "type": "object"}, {"aliases": ["stateful service persistent volumes persistent volume storage storage size"], "anchor": "schema-stateful_service--persistent_volumes--persistent_volume--storage--storage_size", "description": "Size in GiB of the persistent storage.", "document_id": "xcsh-docs:resources:workload:properties:stateful_service:persistent_volumes:persistent_volume:storage", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["stateful_service", "persistent_volumes", "persistent_volume", "storage", "storage_size"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/stateful_service/persistent_volumes/persistent_volume/storage/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Persistent storage configuration is used to configure Persistent Volume Claim (PVC)", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# stateful_service.persistent_volumes.persistent_volume.storage

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [stateful_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/)
- [stateful_service.persistent_volumes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/persistent_volumes/)
- [stateful_service.persistent_volumes.persistent_volume](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/persistent_volumes/persistent_volume/)
- stateful_service.persistent_volumes.persistent_volume.storage

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Persistent storage configuration is used to configure Persistent Volume Claim (PVC).

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("storage_size"),
  validators.ConflictingObjectAttributes("class_name",
    "default")}
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

<a id="schema-stateful_service--persistent_volumes--persistent_volume--storage--access_mode"></a>

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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["ACCESS_MODE_READ_ONLY_MANY","ACCESS_MODE_READ_WRITE_MANY","ACCESS_MODE_READ_WRITE_ONCE"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("ACCESS_MODE_READ_WRITE_ONCE",
    "ACCESS_MODE_READ_WRITE_MANY",
    "ACCESS_MODE_READ_ONLY_MANY"),
}
```

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

<a id="schema-stateful_service--persistent_volumes--persistent_volume--storage--class_name"></a>

### class_name property

Type: `"string"`. Optional.

Exclusive with \[default\] Use the specified class name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [default](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/persistent_volumes/persistent_volume/storage/default/): complete subsection reference.

<a id="schema-stateful_service--persistent_volumes--persistent_volume--storage--storage_size"></a>

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
