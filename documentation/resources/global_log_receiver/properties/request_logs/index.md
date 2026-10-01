---
page_title: "request_logs"
subcategory: ""
description: "request_logs for xcsh_global_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 2156, "body_sha256": "sha256:1fce68bfb1c5d02dd6669a2b2fc8662a4fe7efeb932b55373d5ea33adad40b03", "child_ids": ["xcsh-docs:resources:global_log_receiver:properties:request_logs:sampled", "xcsh-docs:resources:global_log_receiver:properties:request_logs:unsampled"], "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:request_logs", "parent_id": "xcsh-docs:resources:global_log_receiver:reference", "path": "documentation/resources/global_log_receiver/properties/request_logs/index.md", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["request_logs"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/request_logs/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "request_logs for xcsh_global_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# request_logs

Breadcrumbs:

- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/)
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

- [sampled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/request_logs/sampled/): complete subsection reference.

- [unsampled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/request_logs/unsampled/): complete subsection reference.

## Next pages

- [request_logs.sampled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/request_logs/sampled/)
- [request_logs.unsampled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/request_logs/unsampled/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/)
- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/)
