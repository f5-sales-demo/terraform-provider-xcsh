---
page_title: "gcp_bucket_receiver"
subcategory: ""
description: "gcp_bucket_receiver for xcsh_global_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 3610, "body_sha256": "sha256:00737596e8983f66adccc159df2df304f6dcb249119b6162d0c40d8eed763794", "child_ids": ["xcsh-docs:data-sources:global_log_receiver:properties:gcp_bucket_receiver:batch", "xcsh-docs:data-sources:global_log_receiver:properties:gcp_bucket_receiver:compression", "xcsh-docs:data-sources:global_log_receiver:properties:gcp_bucket_receiver:filename_options", "xcsh-docs:data-sources:global_log_receiver:properties:gcp_bucket_receiver:gcp_cred"], "collection_id": "xcsh-docs:data-sources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:global_log_receiver:properties:gcp_bucket_receiver", "parent_id": "xcsh-docs:data-sources:global_log_receiver:reference", "path": "documentation/data-sources/global_log_receiver/properties/gcp_bucket_receiver/index.md", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["gcp_bucket_receiver"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/global_log_receiver/properties/gcp_bucket_receiver/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "gcp_bucket_receiver for xcsh_global_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
