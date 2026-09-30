---
page_title: "splunk_receiver.batch.max_events_disabled"
subcategory: ""
description: "splunk_receiver.batch.max_events_disabled for xcsh_global_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 1088, "body_sha256": "sha256:186b7c6ec411f389d94eb7107c9d0c95c8dcedde8c2cae8ce8048a6ad0831fa1", "canonical_id": "xcsh-docs:resources:global_log_receiver:properties:splunk_receiver:batch:max_events_disabled", "child_ids": [], "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:splunk_receiver:batch:max_events_disabled", "parent_id": "xcsh-docs:resources:global_log_receiver:properties:splunk_receiver:batch", "path": "docs/guides/resources--global_log_receiver--properties--splunk_receiver--batch--max_events_disabled.md", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["splunk_receiver", "batch", "max_events_disabled"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/splunk_receiver/batch/max_events_disabled/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "splunk_receiver.batch.max_events_disabled for xcsh_global_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# splunk_receiver.batch.max_events_disabled

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md)
- [Property reference](resources--global_log_receiver--reference.md)
- [splunk_receiver](resources--global_log_receiver--properties--splunk_receiver.md)
- [splunk_receiver.batch](resources--global_log_receiver--properties--splunk_receiver--batch.md)
- splunk_receiver.batch.max_events_disabled

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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
max_events_disabled = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [splunk_receiver.batch](resources--global_log_receiver--properties--splunk_receiver--batch.md)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md)
