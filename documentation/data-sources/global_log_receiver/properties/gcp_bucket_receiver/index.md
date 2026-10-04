---
page_title: "gcp_bucket_receiver"
subcategory: ""
description: "GCP Bucket Configuration for Global Log Receiver."
xcsh_docs: {"aliases": ["gcp bucket receiver"], "body_bytes": 3610, "body_sha256": "sha256:bb455834613bfb97eff36429e9699c727ca2743dcc841605d415494e725cc518", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:data-sources:global_log_receiver:properties:gcp_bucket_receiver:batch", "xcsh-docs:data-sources:global_log_receiver:properties:gcp_bucket_receiver:compression", "xcsh-docs:data-sources:global_log_receiver:properties:gcp_bucket_receiver:filename_options", "xcsh-docs:data-sources:global_log_receiver:properties:gcp_bucket_receiver:gcp_cred"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:global_log_receiver:properties:gcp_bucket_receiver", "parent_id": "xcsh-docs:data-sources:global_log_receiver:reference", "path": "documentation/data-sources/global_log_receiver/properties/gcp_bucket_receiver/index.md", "product": "distributed-cloud", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1120122010230312-1032002320321100-0101031101303121-2303120002011020-2310102000121031-0120030130012202-1331001231331332-1200111011022302", "registry_path": "docs/guides/data-sources--global_log_receiver--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["gcp_bucket_receiver"], "schema_version": 1, "sections": [{"aliases": ["gcp bucket receiver batch"], "anchor": "section", "description": "Batch OPTIONS allow tuning for how batches of logs are sent to an endpoint.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:gcp_bucket_receiver:batch", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["gcp_bucket_receiver", "batch"], "syntax": "attribute", "type": "object"}, {"aliases": ["gcp bucket receiver bucket"], "anchor": "schema-gcp_bucket_receiver--bucket", "description": "GCP Bucket Name.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:gcp_bucket_receiver", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["gcp_bucket_receiver", "bucket"], "syntax": "attribute", "type": "string"}, {"aliases": ["gcp bucket receiver compression"], "anchor": "section", "description": "Compression Type.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:gcp_bucket_receiver:compression", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["gcp_bucket_receiver", "compression"], "syntax": "attribute", "type": "object"}, {"aliases": ["gcp bucket receiver filename options"], "anchor": "section", "description": "Filename OPTIONS allow customization of filename and folder paths used by a destination endpoint bucket or file.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:gcp_bucket_receiver:filename_options", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["gcp_bucket_receiver", "filename_options"], "syntax": "attribute", "type": "object"}, {"aliases": ["gcp bucket receiver gcp cred"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:gcp_bucket_receiver:gcp_cred", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["gcp_bucket_receiver", "gcp_cred"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/global_log_receiver/properties/gcp_bucket_receiver/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "GCP Bucket Configuration for Global Log Receiver.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# gcp_bucket_receiver

Breadcrumbs:

- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/)
- gcp_bucket_receiver

<a id="section"></a>

Type: `"single"`. Computed.

GCP Bucket Configuration for Global Log Receiver.

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

- [batch](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/gcp_bucket_receiver/batch/): complete subsection reference.

<a id="schema-gcp_bucket_receiver--bucket"></a>

### bucket property

Type: `"string"`. Computed.

GCP Bucket Name. GCP Bucket Name.

Upstream description:

GCP Bucket Name.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [compression](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/gcp_bucket_receiver/compression/): complete subsection reference.

- [filename_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/gcp_bucket_receiver/filename_options/): complete subsection reference.

- [gcp_cred](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/gcp_bucket_receiver/gcp_cred/): complete subsection reference.

## Next pages

- [gcp_bucket_receiver.batch](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/gcp_bucket_receiver/batch/)
- [gcp_bucket_receiver.compression](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/gcp_bucket_receiver/compression/)
- [gcp_bucket_receiver.filename_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/gcp_bucket_receiver/filename_options/)
- [gcp_bucket_receiver.gcp_cred](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/gcp_bucket_receiver/gcp_cred/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/)
- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/)
