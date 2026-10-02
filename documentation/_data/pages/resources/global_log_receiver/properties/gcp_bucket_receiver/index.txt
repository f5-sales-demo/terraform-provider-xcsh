---
page_title: "gcp_bucket_receiver"
subcategory: ""
description: "GCP Bucket Configuration for Global Log Receiver."
xcsh_docs: {"aliases": ["gcp bucket receiver"], "body_bytes": 3989, "body_sha256": "sha256:1a67e5d085c7d3a4eb2f22140625f9dced075bb2c057c4faa77e681ea414e010", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:resources:global_log_receiver:properties:gcp_bucket_receiver:batch", "xcsh-docs:resources:global_log_receiver:properties:gcp_bucket_receiver:compression", "xcsh-docs:resources:global_log_receiver:properties:gcp_bucket_receiver:filename_options", "xcsh-docs:resources:global_log_receiver:properties:gcp_bucket_receiver:gcp_cred"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:gcp_bucket_receiver", "parent_id": "xcsh-docs:resources:global_log_receiver:reference", "path": "documentation/resources/global_log_receiver/properties/gcp_bucket_receiver/index.md", "product": "distributed-cloud", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2020331032132223-1121103322222103-0231300112111122-3202221112310311-1212000231222030-2200012032311130-3011322122012001-1001013020113031", "registry_path": "docs/guides/resources--global_log_receiver--reference--group-002.md", "relationships": [{"anchor": "schema-gcp_bucket_receiver--bucket", "enforcement": "provider-schema", "group": "gcp_bucket_receiver:RequiredObjectAttributes:bucket", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:gcp_bucket_receiver", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["gcp_bucket_receiver"], "schema_version": 1, "sections": [{"aliases": ["batch"], "anchor": "section", "description": "Batch OPTIONS allow tuning for how batches of logs are sent to an endpoint.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:gcp_bucket_receiver:batch", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-gcp_bucket_receiver--batch--max_bytes", "enforcement": "provider-schema", "group": "gcp_bucket_receiver.batch:ConflictingObjectAttributes:max_bytes,max_bytes_disabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:gcp_bucket_receiver:batch", "type": "conflicts"}, {"anchor": "schema-gcp_bucket_receiver--batch--max_events", "enforcement": "provider-schema", "group": "gcp_bucket_receiver.batch:ConflictingObjectAttributes:max_events,max_events_disabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:gcp_bucket_receiver:batch", "type": "conflicts"}, {"anchor": "schema-gcp_bucket_receiver--batch--timeout_seconds", "enforcement": "provider-schema", "group": "gcp_bucket_receiver.batch:ConflictingObjectAttributes:timeout_seconds,timeout_seconds_default", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:gcp_bucket_receiver:batch", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "gcp_bucket_receiver.batch:ConflictingObjectAttributes:max_bytes,max_bytes_disabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:gcp_bucket_receiver:batch:max_bytes_disabled", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "gcp_bucket_receiver.batch:ConflictingObjectAttributes:max_events,max_events_disabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:gcp_bucket_receiver:batch:max_events_disabled", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "gcp_bucket_receiver.batch:ConflictingObjectAttributes:timeout_seconds,timeout_seconds_default", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:gcp_bucket_receiver:batch:timeout_seconds_default", "type": "conflicts"}], "schema_path": ["gcp_bucket_receiver", "batch"], "syntax": "block", "type": "object"}, {"aliases": ["bucket"], "anchor": "schema-gcp_bucket_receiver--bucket", "description": "GCP Bucket Name.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:gcp_bucket_receiver", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["gcp_bucket_receiver", "bucket"], "syntax": "attribute", "type": "string"}, {"aliases": ["compression"], "anchor": "section", "description": "Compression Type.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:gcp_bucket_receiver:compression", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "gcp_bucket_receiver.compression:ConflictingObjectAttributes:compression_default,compression_gzip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:gcp_bucket_receiver:compression:compression_default", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "gcp_bucket_receiver.compression:ConflictingObjectAttributes:compression_default,compression_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:gcp_bucket_receiver:compression:compression_default", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "gcp_bucket_receiver.compression:ConflictingObjectAttributes:compression_default,compression_gzip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:gcp_bucket_receiver:compression:compression_gzip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "gcp_bucket_receiver.compression:ConflictingObjectAttributes:compression_gzip,compression_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:gcp_bucket_receiver:compression:compression_gzip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "gcp_bucket_receiver.compression:ConflictingObjectAttributes:compression_default,compression_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:gcp_bucket_receiver:compression:compression_none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "gcp_bucket_receiver.compression:ConflictingObjectAttributes:compression_gzip,compression_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:gcp_bucket_receiver:compression:compression_none", "type": "conflicts"}], "schema_path": ["gcp_bucket_receiver", "compression"], "syntax": "block", "type": "object"}, {"aliases": ["filename options"], "anchor": "section", "description": "Filename OPTIONS allow customization of filename and folder paths used by a destination endpoint bucket or file.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:gcp_bucket_receiver:filename_options", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-gcp_bucket_receiver--filename_options--custom_folder", "enforcement": "provider-schema", "group": "gcp_bucket_receiver.filename_options:ConflictingObjectAttributes:custom_folder,log_type_folder", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:gcp_bucket_receiver:filename_options", "type": "conflicts"}, {"anchor": "schema-gcp_bucket_receiver--filename_options--custom_folder", "enforcement": "provider-schema", "group": "gcp_bucket_receiver.filename_options:ConflictingObjectAttributes:custom_folder,no_folder", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:gcp_bucket_receiver:filename_options", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "gcp_bucket_receiver.filename_options:ConflictingObjectAttributes:custom_folder,log_type_folder", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:gcp_bucket_receiver:filename_options:log_type_folder", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "gcp_bucket_receiver.filename_options:ConflictingObjectAttributes:log_type_folder,no_folder", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:gcp_bucket_receiver:filename_options:log_type_folder", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "gcp_bucket_receiver.filename_options:ConflictingObjectAttributes:custom_folder,no_folder", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:gcp_bucket_receiver:filename_options:no_folder", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "gcp_bucket_receiver.filename_options:ConflictingObjectAttributes:log_type_folder,no_folder", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:gcp_bucket_receiver:filename_options:no_folder", "type": "conflicts"}], "schema_path": ["gcp_bucket_receiver", "filename_options"], "syntax": "block", "type": "object"}, {"aliases": ["gcp cred"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:gcp_bucket_receiver:gcp_cred", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-gcp_bucket_receiver--gcp_cred--name", "enforcement": "provider-schema", "group": "gcp_bucket_receiver.gcp_cred:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:gcp_bucket_receiver:gcp_cred", "type": "requires"}], "schema_path": ["gcp_bucket_receiver", "gcp_cred"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/gcp_bucket_receiver/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "GCP Bucket Configuration for Global Log Receiver.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# gcp_bucket_receiver

Breadcrumbs:

- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/)
- gcp_bucket_receiver

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

GCP Bucket Configuration for Global Log Receiver.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("bucket")}
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
gcp_bucket_receiver {
  # Configure direct properties listed below.
}
```

## Direct properties

- [batch](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/gcp_bucket_receiver/batch/): complete subsection reference.

<a id="schema-gcp_bucket_receiver--bucket"></a>

### bucket property

Type: `"string"`. Optional.

GCP Bucket Name. GCP Bucket Name.

Upstream description:

GCP Bucket Name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(3, 128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 3,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 3,
    "pattern": "^[a-z0-9]+[a-z0-9_\\\\.-]+[a-z0-9]$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "3",
    "ves.io.schema.rules.string.pattern": "^[a-z0-9]+[a-z0-9_\\\\.-]+[a-z0-9]$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "3",
    "ves.io.schema.rules.string.pattern": "^[a-z0-9]+[a-z0-9_\\\\.-]+[a-z0-9]$"
  }
}
```

- [compression](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/gcp_bucket_receiver/compression/): complete subsection reference.

- [filename_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/gcp_bucket_receiver/filename_options/): complete subsection reference.

- [gcp_cred](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/gcp_bucket_receiver/gcp_cred/): complete subsection reference.

## Next pages

- [gcp_bucket_receiver.batch](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/gcp_bucket_receiver/batch/)
- [gcp_bucket_receiver.compression](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/gcp_bucket_receiver/compression/)
- [gcp_bucket_receiver.filename_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/gcp_bucket_receiver/filename_options/)
- [gcp_bucket_receiver.gcp_cred](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/gcp_bucket_receiver/gcp_cred/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/)
- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/)
