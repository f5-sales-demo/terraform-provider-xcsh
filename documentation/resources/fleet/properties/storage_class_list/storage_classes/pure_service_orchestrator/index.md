---
page_title: "storage_class_list.storage_classes.pure_service_orchestrator"
subcategory: ""
description: "storage_class_list.storage_classes.pure_service_orchestrator for xcsh_fleet."
xcsh_docs: {"aliases": [], "body_bytes": 5222, "body_sha256": "sha256:8fde009524ae30beaabfdad262570132d87b0d6257d14d4b65d13a01e3abd711", "child_ids": [], "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:storage_class_list:storage_classes:pure_service_orchestrator", "parent_id": "xcsh-docs:resources:fleet:properties:storage_class_list:storage_classes", "path": "documentation/resources/fleet/properties/storage_class_list/storage_classes/pure_service_orchestrator/index.md", "provider_name": "fleet", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["storage_class_list", "storage_classes", "pure_service_orchestrator"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/storage_class_list/storage_classes/pure_service_orchestrator/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "storage_class_list.storage_classes.pure_service_orchestrator for xcsh_fleet.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# storage_class_list.storage_classes.pure_service_orchestrator

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/)
- [storage_class_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_class_list/)
- [storage_class_list.storage_classes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_class_list/storage_classes/)
- storage_class_list.storage_classes.pure_service_orchestrator

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Storage class Device configuration for Pure Service Orchestrator.

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
pure_service_orchestrator {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-storage_class_list--storage_classes--pure_service_orchestrator--backend"></a>

### backend property

Type: `"string"`. Optional.

\[Enum: block|file\] Defines type of Pure storage backend block or file. The volume will have the
aspects defined in the chosen virtual pool. Possible values are \`block\`, \`file\`.

Upstream description:

Defines type of Pure storage backend block or file. The volume will have the aspects defined in the
chosen virtual pool.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("block",
    "file"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "block",
    "file"
  ],
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
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
    "ves.io.schema.rules.string.in": "[\\\"block\\\",\\\"file\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"block\\\",\\\"file\\\"]"
  }
}
```

<a id="schema-storage_class_list--storage_classes--pure_service_orchestrator--bandwidth_limit"></a>

### bandwidth_limit property

Type: `"string"`. Optional.

It must be between 1 MB/s and 512 GB/s. Enter the size as a number (bytes must be multiple of 512)
or number with a single character unit symbol. Valid unit symbols are K, M, G, representing KiB,
MiB, and GiB.

Upstream description:

It must be between 1 MB/s and 512 GB/s. Enter the size as a number (bytes must be multiple of 512)
or number with a single character unit symbol. Valid unit symbols are K, M, G, representing KiB,
MiB, and GiB.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(12),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 12,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 12,
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
    "ves.io.schema.rules.string.max_len": "12"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "12"
  }
}
```

<a id="schema-storage_class_list--storage_classes--pure_service_orchestrator--iops_limit"></a>

### iops_limit property

Type: `"number"`. Optional.

Enable IOPS limitation. It must be between 100 and 100 million. If value is 0, IOPS limit is not
defined.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  validators.Int64RangeSetValidator(
    validators.Int64Range{Minimum: 0, Maximum: 0},
    validators.Int64Range{Minimum: 100, Maximum: 100000000},
  ),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 100000000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.ranges": "0,100-100000000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,100-100000000"
  }
}
```

## Next pages

- [storage_class_list.storage_classes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_class_list/storage_classes/)
- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
