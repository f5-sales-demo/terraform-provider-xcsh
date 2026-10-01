---
page_title: "request_logs.unsampled"
subcategory: ""
description: "request_logs.unsampled for xcsh_global_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 1017, "body_sha256": "sha256:e4812fc0d599ed5b1c83ffb857a3c198738151bb393ab287ea2413a8821d4dc1", "canonical_id": "xcsh-docs:resources:global_log_receiver:properties:request_logs:unsampled", "child_ids": [], "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:request_logs:unsampled", "parent_id": "xcsh-docs:resources:global_log_receiver:properties:request_logs", "path": "docs/guides/resources--global_log_receiver--properties--request_logs--unsampled.md", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["request_logs", "unsampled"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/request_logs/unsampled/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "request_logs.unsampled for xcsh_global_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# request_logs.unsampled

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md)
- [Property reference](resources--global_log_receiver--reference.md)
- [request_logs](resources--global_log_receiver--properties--request_logs.md)
- request_logs.unsampled

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
unsampled = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [request_logs](resources--global_log_receiver--properties--request_logs.md)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md)
