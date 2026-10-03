---
page_title: "custom_storage_config.storage_class_list.storage_classes.pure_service_orchestrator"
subcategory: ""
description: "Storage class Device configuration for Pure Service Orchestrator."
xcsh_docs: {"aliases": ["custom storage config storage class list storage classes pure service orchestrator"], "body_bytes": 5030, "body_sha256": "sha256:cf270846224bfe4db25432620972b653d946695718a8b84eac2ac9615058a605", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_class_list:storage_classes:pure_service_orchestrator", "parent_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_class_list:storage_classes", "path": "documentation/data-sources/voltstack_site/properties/custom_storage_config/storage_class_list/storage_classes/pure_service_orchestrator/index.md", "product": "distributed-cloud", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2213301022312101-0220001232311302-3103321123100103-2333332331303311-1213112123211302-0323330130321032-0333202133301210-1223202222332031", "registry_path": "docs/guides/data-sources--voltstack_site--reference--group-006.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_storage_config", "storage_class_list", "storage_classes", "pure_service_orchestrator"], "schema_version": 1, "sections": [{"aliases": ["backend servers", "custom storage config storage class list storage classes pure service orchestrator backend", "origin servers", "upstream servers"], "anchor": "schema-custom_storage_config--storage_class_list--storage_classes--pure_service_orchestrator--backend", "description": "Defines type of Pure storage backend block or file. The volume will have the aspects defined in the chosen virtual pool.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_class_list:storage_classes:pure_service_orchestrator", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_class_list", "storage_classes", "pure_service_orchestrator", "backend"], "syntax": "attribute", "type": "string"}, {"aliases": ["custom storage config storage class list storage classes pure service orchestrator bandwidth limit"], "anchor": "schema-custom_storage_config--storage_class_list--storage_classes--pure_service_orchestrator--bandwidth_limit", "description": "It must be between 1 MB/s and 512 GB/s. Enter the size as a number (bytes must be multiple of 512) or number with a single character unit symbol. Valid unit symbols are K, M, G, representing KiB, MiB, and GiB.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_class_list:storage_classes:pure_service_orchestrator", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_class_list", "storage_classes", "pure_service_orchestrator", "bandwidth_limit"], "syntax": "attribute", "type": "string"}, {"aliases": ["custom storage config storage class list storage classes pure service orchestrator iops limit"], "anchor": "schema-custom_storage_config--storage_class_list--storage_classes--pure_service_orchestrator--iops_limit", "description": "Enable IOPS limitation. It must be between 100 and 100 million. If value is 0, IOPS limit is not defined.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_class_list:storage_classes:pure_service_orchestrator", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_class_list", "storage_classes", "pure_service_orchestrator", "iops_limit"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/voltstack_site/properties/custom_storage_config/storage_class_list/storage_classes/pure_service_orchestrator/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Storage class Device configuration for Pure Service Orchestrator.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_storage_config.storage_class_list.storage_classes.pure_service_orchestrator

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/)
- [custom_storage_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/)
- [custom_storage_config.storage_class_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/storage_class_list/)
- [custom_storage_config.storage_class_list.storage_classes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/storage_class_list/storage_classes/)
- custom_storage_config.storage_class_list.storage_classes.pure_service_orchestrator

<a id="section"></a>

Type: `"single"`. Computed.

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

## Direct properties

<a id="schema-custom_storage_config--storage_class_list--storage_classes--pure_service_orchestrator--backend"></a>

### backend property

Type: `"string"`. Computed.

\[Enum: block|file\] Defines type of Pure storage backend block or file. The volume will have the
aspects defined in the chosen virtual pool. Possible values are \`block\`, \`file\`.

Upstream description:

Defines type of Pure storage backend block or file. The volume will have the aspects defined in the
chosen virtual pool.

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

Type: `"string"`. Computed.

It must be between 1 MB/s and 512 GB/s. Enter the size as a number (bytes must be multiple of 512)
or number with a single character unit symbol. Valid unit symbols are K, M, G, representing KiB,
MiB, and GiB.

Upstream description:

It must be between 1 MB/s and 512 GB/s. Enter the size as a number (bytes must be multiple of 512)
or number with a single character unit symbol. Valid unit symbols are K, M, G, representing KiB,
MiB, and GiB.

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

Type: `"number"`. Computed.

Enable IOPS limitation. It must be between 100 and 100 million. If value is 0, IOPS limit is not
defined.

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

- [custom_storage_config.storage_class_list.storage_classes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/storage_class_list/storage_classes/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/)
