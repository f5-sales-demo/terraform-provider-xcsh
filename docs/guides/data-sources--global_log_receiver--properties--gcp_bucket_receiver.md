---
page_title: "gcp_bucket_receiver"
subcategory: ""
description: "gcp_bucket_receiver for xcsh_global_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 2903, "body_sha256": "sha256:5d1829f4880f3ca2d0e369a8297aaa3ea2f83850bc8fcf5034c1ea8aee03b93c", "canonical_id": "xcsh-docs:data-sources:global_log_receiver:properties:gcp_bucket_receiver", "child_ids": ["xcsh-docs:data-sources:global_log_receiver:properties:gcp_bucket_receiver:batch", "xcsh-docs:data-sources:global_log_receiver:properties:gcp_bucket_receiver:compression", "xcsh-docs:data-sources:global_log_receiver:properties:gcp_bucket_receiver:filename_options", "xcsh-docs:data-sources:global_log_receiver:properties:gcp_bucket_receiver:gcp_cred"], "collection_id": "xcsh-docs:data-sources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:global_log_receiver:properties:gcp_bucket_receiver", "parent_id": "xcsh-docs:data-sources:global_log_receiver:reference", "path": "docs/guides/data-sources--global_log_receiver--properties--gcp_bucket_receiver.md", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["gcp_bucket_receiver"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/global_log_receiver/properties/gcp_bucket_receiver/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "gcp_bucket_receiver for xcsh_global_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# gcp_bucket_receiver

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md)
- [Property reference](data-sources--global_log_receiver--reference.md)
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

- [batch](data-sources--global_log_receiver--properties--gcp_bucket_receiver--batch.md): complete subsection reference.

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

- [compression](data-sources--global_log_receiver--properties--gcp_bucket_receiver--compression.md): complete subsection reference.

- [filename_options](data-sources--global_log_receiver--properties--gcp_bucket_receiver--filename_options.md): complete subsection reference.

- [gcp_cred](data-sources--global_log_receiver--properties--gcp_bucket_receiver--gcp_cred.md): complete subsection reference.

## Next pages

- [gcp_bucket_receiver.batch](data-sources--global_log_receiver--properties--gcp_bucket_receiver--batch.md)
- [gcp_bucket_receiver.compression](data-sources--global_log_receiver--properties--gcp_bucket_receiver--compression.md)
- [gcp_bucket_receiver.filename_options](data-sources--global_log_receiver--properties--gcp_bucket_receiver--filename_options.md)
- [gcp_bucket_receiver.gcp_cred](data-sources--global_log_receiver--properties--gcp_bucket_receiver--gcp_cred.md)
- [Property reference](data-sources--global_log_receiver--reference.md)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md)
