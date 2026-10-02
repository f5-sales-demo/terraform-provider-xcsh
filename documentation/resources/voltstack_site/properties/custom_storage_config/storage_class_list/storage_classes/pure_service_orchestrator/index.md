---
page_title: "custom_storage_config.storage_class_list.storage_classes.pure_service_orchestrator"
subcategory: ""
description: "Storage class Device configuration for Pure Service Orchestrator."
xcsh_docs: {"aliases": ["custom storage config storage class list storage classes pure service orchestrator"], "body_bytes": 5681, "body_sha256": "sha256:65f5627d91fa79df04dd018432633acc5aa8d50f500dce67a36e39cc004bfba8", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_class_list:storage_classes:pure_service_orchestrator", "parent_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_class_list:storage_classes", "path": "documentation/resources/voltstack_site/properties/custom_storage_config/storage_class_list/storage_classes/pure_service_orchestrator/index.md", "product": "distributed-cloud", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2013033101131001-2020332001123121-0301320113203232-1310233032132303-0031122121121110-0300100302023122-3200301322122133-2320202010323200", "registry_path": "docs/guides/resources--voltstack_site--reference--group-006.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_storage_config", "storage_class_list", "storage_classes", "pure_service_orchestrator"], "schema_version": 1, "sections": [{"aliases": ["backend", "backend servers", "origin servers", "upstream servers"], "anchor": "schema-custom_storage_config--storage_class_list--storage_classes--pure_service_orchestrator--backend", "description": "Defines type of Pure storage backend block or file. The volume will have the aspects defined in the chosen virtual pool.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_class_list:storage_classes:pure_service_orchestrator", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_class_list", "storage_classes", "pure_service_orchestrator", "backend"], "syntax": "attribute", "type": "string"}, {"aliases": ["bandwidth limit"], "anchor": "schema-custom_storage_config--storage_class_list--storage_classes--pure_service_orchestrator--bandwidth_limit", "description": "It must be between 1 MB/s and 512 GB/s. Enter the size as a number (bytes must be multiple of 512) or number with a single character unit symbol. Valid unit symbols are K, M, G, representing KiB, MiB, and GiB.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_class_list:storage_classes:pure_service_orchestrator", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_class_list", "storage_classes", "pure_service_orchestrator", "bandwidth_limit"], "syntax": "attribute", "type": "string"}, {"aliases": ["iops limit"], "anchor": "schema-custom_storage_config--storage_class_list--storage_classes--pure_service_orchestrator--iops_limit", "description": "Enable IOPS limitation. It must be between 100 and 100 million. If value is 0, IOPS limit is not defined.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_class_list:storage_classes:pure_service_orchestrator", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_class_list", "storage_classes", "pure_service_orchestrator", "iops_limit"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/custom_storage_config/storage_class_list/storage_classes/pure_service_orchestrator/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Storage class Device configuration for Pure Service Orchestrator.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_storage_config.storage_class_list.storage_classes.pure_service_orchestrator

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/)
- [custom_storage_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/)
- [custom_storage_config.storage_class_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_class_list/)
- [custom_storage_config.storage_class_list.storage_classes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_class_list/storage_classes/)
- custom_storage_config.storage_class_list.storage_classes.pure_service_orchestrator

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

<a id="schema-custom_storage_config--storage_class_list--storage_classes--pure_service_orchestrator--backend"></a>

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

<a id="schema-custom_storage_config--storage_class_list--storage_classes--pure_service_orchestrator--bandwidth_limit"></a>

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

<a id="schema-custom_storage_config--storage_class_list--storage_classes--pure_service_orchestrator--iops_limit"></a>

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

- [custom_storage_config.storage_class_list.storage_classes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_class_list/storage_classes/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
