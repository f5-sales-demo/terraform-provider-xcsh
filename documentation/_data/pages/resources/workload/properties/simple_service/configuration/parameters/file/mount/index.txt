---
page_title: "simple_service.configuration.parameters.file.mount"
subcategory: "Container"
description: "Volume mount describes how volume is mounted inside a workload."
xcsh_docs: {"aliases": ["simple service configuration parameters file mount"], "body_bytes": 5554, "body_sha256": "sha256:f8f2695d1ba323be3f28ca0a0503a91f5fdd0d6ec27302343f2e345e0a48d61b", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:simple_service:configuration:parameters:file:mount", "parent_id": "xcsh-docs:resources:workload:properties:simple_service:configuration:parameters:file", "path": "documentation/resources/workload/properties/simple_service/configuration/parameters/file/mount/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2101021012222331-3232010123232032-3303103311113010-2030322211032012-0103002110333102-2031001130221330-3231010210311002-0223201022223102", "registry_path": "docs/guides/resources--workload--reference--group-016.md", "relationships": [{"anchor": "schema-simple_service--configuration--parameters--file--mount--mount_path", "enforcement": "provider-schema", "group": "simple_service.configuration.parameters.file.mount:RequiredObjectAttributes:mount_path", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:simple_service:configuration:parameters:file:mount", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["simple_service", "configuration", "parameters", "file", "mount"], "schema_version": 1, "sections": [{"aliases": ["mode"], "anchor": "schema-simple_service--configuration--parameters--file--mount--mode", "description": "Mode in which the volume should be mounted to the workload - VOLUME_MOUNT_READ_ONLY: ReadOnly Mount the volume in read-only mode - VOLUME_MOUNT_READ_WRITE: Read Write Mount the volume in read-write mode.", "document_id": "xcsh-docs:resources:workload:properties:simple_service:configuration:parameters:file:mount", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["simple_service", "configuration", "parameters", "file", "mount", "mode"], "syntax": "attribute", "type": "string"}, {"aliases": ["mount path"], "anchor": "schema-simple_service--configuration--parameters--file--mount--mount_path", "description": "Path within the workload container at which the volume should be mounted. Must not contain ':'.", "document_id": "xcsh-docs:resources:workload:properties:simple_service:configuration:parameters:file:mount", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["simple_service", "configuration", "parameters", "file", "mount", "mount_path"], "syntax": "attribute", "type": "string"}, {"aliases": ["sub path"], "anchor": "schema-simple_service--configuration--parameters--file--mount--sub_path", "description": "Path within the volume from which the workload's volume should be mounted. Defaults to \"\" (volume's root).", "document_id": "xcsh-docs:resources:workload:properties:simple_service:configuration:parameters:file:mount", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["simple_service", "configuration", "parameters", "file", "mount", "sub_path"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/simple_service/configuration/parameters/file/mount/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Volume mount describes how volume is mounted inside a workload.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# simple_service.configuration.parameters.file.mount

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [simple_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/simple_service/)
- [simple_service.configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/simple_service/configuration/)
- [simple_service.configuration.parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/simple_service/configuration/parameters/)
- [simple_service.configuration.parameters.file](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/simple_service/configuration/parameters/file/)
- simple_service.configuration.parameters.file.mount

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

<a id="schema-simple_service--configuration--parameters--file--mount--mode"></a>

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

<a id="schema-simple_service--configuration--parameters--file--mount--mount_path"></a>

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

<a id="schema-simple_service--configuration--parameters--file--mount--sub_path"></a>

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

- [simple_service.configuration.parameters.file](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/simple_service/configuration/parameters/file/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
