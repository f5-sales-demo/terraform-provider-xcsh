---
page_title: "request_logs"
subcategory: ""
description: "request_logs for xcsh_global_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 1485, "body_sha256": "sha256:21152efb4f541b5069c144981d6f4b49e5543d9510d472351333aea0a841d9be", "canonical_id": "xcsh-docs:data-sources:global_log_receiver:properties:request_logs", "child_ids": ["xcsh-docs:data-sources:global_log_receiver:properties:request_logs:sampled", "xcsh-docs:data-sources:global_log_receiver:properties:request_logs:unsampled"], "collection_id": "xcsh-docs:data-sources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:global_log_receiver:properties:request_logs", "parent_id": "xcsh-docs:data-sources:global_log_receiver:reference", "path": "docs/guides/data-sources--global_log_receiver--properties--request_logs.md", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["request_logs"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/global_log_receiver/properties/request_logs/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "request_logs for xcsh_global_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# request_logs

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md)
- [Property reference](data-sources--global_log_receiver--reference.md)
- request_logs

<a id="section"></a>

Type: `"single"`. Computed.

Configuration for request logs with sampling choice. Allows selection between sampled (default) or
unsampled (full) request logs.

Upstream description:

Configuration for request logs with sampling choice. Allows selection between sampled (default) or
unsampled (full) request logs.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-sampling_choice": "[\"sampled\",\"unsampled\"]"
}
```

## Direct properties

- [sampled](data-sources--global_log_receiver--properties--request_logs--sampled.md): complete subsection reference.

- [unsampled](data-sources--global_log_receiver--properties--request_logs--unsampled.md): complete subsection reference.

## Next pages

- [request_logs.sampled](data-sources--global_log_receiver--properties--request_logs--sampled.md)
- [request_logs.unsampled](data-sources--global_log_receiver--properties--request_logs--unsampled.md)
- [Property reference](data-sources--global_log_receiver--reference.md)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md)
