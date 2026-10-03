---
page_title: "service.configuration.parameters.file.mount"
subcategory: "Container"
description: "Volume mount describes how volume is mounted inside a workload."
xcsh_docs: {"aliases": ["service configuration parameters file mount"], "body_bytes": 5449, "body_sha256": "sha256:95473b344e07980e0e1b8e175e3574acbca0f21e448c2cf89e873dccd1416a89", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:service:configuration:parameters:file:mount", "parent_id": "xcsh-docs:resources:workload:properties:service:configuration:parameters:file", "path": "documentation/resources/workload/properties/service/configuration/parameters/file/mount/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2031313210013212-2213211123312132-1101031213300321-1021132030122012-2322101220331311-2330202332012212-2201121213103233-2113120200323020", "registry_path": "docs/guides/resources--workload--reference--group-015.md", "relationships": [{"anchor": "schema-service--configuration--parameters--file--mount--mount_path", "enforcement": "provider-schema", "group": "service.configuration.parameters.file.mount:RequiredObjectAttributes:mount_path", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:configuration:parameters:file:mount", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["service", "configuration", "parameters", "file", "mount"], "schema_version": 1, "sections": [{"aliases": ["service configuration parameters file mount mode"], "anchor": "schema-service--configuration--parameters--file--mount--mode", "description": "Mode in which the volume should be mounted to the workload - VOLUME_MOUNT_READ_ONLY: ReadOnly Mount the volume in read-only mode - VOLUME_MOUNT_READ_WRITE: Read Write Mount the volume in read-write mode.", "document_id": "xcsh-docs:resources:workload:properties:service:configuration:parameters:file:mount", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "configuration", "parameters", "file", "mount", "mode"], "syntax": "attribute", "type": "string"}, {"aliases": ["service configuration parameters file mount mount path"], "anchor": "schema-service--configuration--parameters--file--mount--mount_path", "description": "Path within the workload container at which the volume should be mounted. Must not contain ':'.", "document_id": "xcsh-docs:resources:workload:properties:service:configuration:parameters:file:mount", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "configuration", "parameters", "file", "mount", "mount_path"], "syntax": "attribute", "type": "string"}, {"aliases": ["service configuration parameters file mount sub path"], "anchor": "schema-service--configuration--parameters--file--mount--sub_path", "description": "Path within the volume from which the workload's volume should be mounted. Defaults to \"\" (volume's root).", "document_id": "xcsh-docs:resources:workload:properties:service:configuration:parameters:file:mount", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "configuration", "parameters", "file", "mount", "sub_path"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/service/configuration/parameters/file/mount/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Volume mount describes how volume is mounted inside a workload.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["workloadCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.configuration.parameters.file.mount

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/)
- [service.configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/configuration/)
- [service.configuration.parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/configuration/parameters/)
- [service.configuration.parameters.file](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/configuration/parameters/file/)
- service.configuration.parameters.file.mount

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Volume mount describes how volume is mounted inside a workload.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("mount_path")}
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
mount {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-service--configuration--parameters--file--mount--mode"></a>

### mode property

Type: `"string"`. Optional.

\[Enum: VOLUME\_MOUNT\_READ\_ONLY|VOLUME\_MOUNT\_READ\_WRITE\] Mode in which the volume should be
mounted to the workload - VOLUME\_MOUNT\_READ\_ONLY: ReadOnly Mount the volume in read-only mode -
VOLUME\_MOUNT\_READ\_WRITE: Read Write Mount the volume in read-write mode. Possible values are
\`VOLUME\_MOUNT\_READ\_ONLY\`, \`VOLUME\_MOUNT\_READ\_WRITE\`. Defaults to
\`VOLUME\_MOUNT\_READ\_ONLY\`.

Upstream description:

Mode in which the volume should be mounted to the workload

&#8203;- VOLUME\_MOUNT\_READ\_ONLY: ReadOnly

Mount the volume in read-only mode &#8203;- VOLUME\_MOUNT\_READ\_WRITE: Read Write

Mount the volume in read-write mode.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("VOLUME_MOUNT_READ_ONLY",
    "VOLUME_MOUNT_READ_WRITE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "VOLUME_MOUNT_READ_ONLY",
  "enum": [
    "VOLUME_MOUNT_READ_ONLY",
    "VOLUME_MOUNT_READ_WRITE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-service--configuration--parameters--file--mount--mount_path"></a>

### mount_path property

Type: `"string"`. Optional.

Path within the workload container at which the volume should be mounted. Must not contain ':'.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": "^[^:]*$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.pattern": "^[^:]*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.pattern": "^[^:]*$"
  }
}
```

<a id="schema-service--configuration--parameters--file--mount--sub_path"></a>

### sub_path property

Type: `"string"`. Optional.

Path within the volume from which the workload's volume should be mounted. Defaults to '' (volume's
root).

Upstream description:

Path within the volume from which the workload's volume should be mounted. Defaults to "" (volume's
root).

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

## Next pages

- [service.configuration.parameters.file](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/configuration/parameters/file/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
