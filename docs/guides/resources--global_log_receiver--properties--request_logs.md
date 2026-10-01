---
page_title: "request_logs"
subcategory: ""
description: "request_logs for xcsh_global_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 1748, "body_sha256": "sha256:7f12de1575ae7e9c01243b64d7b3799aab38c4190249e211efbcc36f9d6c8f2a", "canonical_id": "xcsh-docs:resources:global_log_receiver:properties:request_logs", "child_ids": ["xcsh-docs:resources:global_log_receiver:properties:request_logs:sampled", "xcsh-docs:resources:global_log_receiver:properties:request_logs:unsampled"], "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:request_logs", "parent_id": "xcsh-docs:resources:global_log_receiver:reference", "path": "docs/guides/resources--global_log_receiver--properties--request_logs.md", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["request_logs"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/request_logs/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "request_logs for xcsh_global_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# request_logs

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md)
- [Property reference](resources--global_log_receiver--reference.md)
- request_logs

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration for request logs with sampling choice. Allows selection between sampled (default) or
unsampled (full) request logs.

Upstream description:

Configuration for request logs with sampling choice. Allows selection between sampled (default) or
unsampled (full) request logs.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("sampled",
    "unsampled")}
```

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

Terraform syntax:

```terraform
request_logs {
  # Configure direct properties listed below.
}
```

## Direct properties

- [sampled](resources--global_log_receiver--properties--request_logs--sampled.md): complete subsection reference.

- [unsampled](resources--global_log_receiver--properties--request_logs--unsampled.md): complete subsection reference.

## Next pages

- [request_logs.sampled](resources--global_log_receiver--properties--request_logs--sampled.md)
- [request_logs.unsampled](resources--global_log_receiver--properties--request_logs--unsampled.md)
- [Property reference](resources--global_log_receiver--reference.md)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md)
