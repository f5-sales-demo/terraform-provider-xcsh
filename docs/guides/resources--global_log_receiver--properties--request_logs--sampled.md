---
page_title: "request_logs.sampled"
subcategory: ""
description: "request_logs.sampled for xcsh_global_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 985, "body_sha256": "sha256:c47354a2ef9ddc8620f07fb8f6cff765e357a7b50cae5642e863cd8b8066a338", "canonical_id": "xcsh-docs:resources:global_log_receiver:properties:request_logs:sampled", "child_ids": [], "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:request_logs:sampled", "parent_id": "xcsh-docs:resources:global_log_receiver:properties:request_logs", "path": "docs/guides/resources--global_log_receiver--properties--request_logs--sampled.md", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["request_logs", "sampled"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/request_logs/sampled/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "request_logs.sampled for xcsh_global_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# request_logs.sampled

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md)
- [Property reference](resources--global_log_receiver--reference.md)
- [request_logs](resources--global_log_receiver--properties--request_logs.md)
- request_logs.sampled

<a id="section"></a>

Type: `["object", {}]`. Optional, Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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
sampled = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [request_logs](resources--global_log_receiver--properties--request_logs.md)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md)
