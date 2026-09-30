---
page_title: "audit_logs"
subcategory: ""
description: "audit_logs for xcsh_global_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 1290, "body_sha256": "sha256:d2a8706111db5af2c9ff9480bddf681bc23dcc8706a7bf997c28098051bbf804", "canonical_id": "xcsh-docs:data-sources:global_log_receiver:properties:audit_logs", "child_ids": [], "collection_id": "xcsh-docs:data-sources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:global_log_receiver:properties:audit_logs", "parent_id": "xcsh-docs:data-sources:global_log_receiver:reference", "path": "docs/guides/data-sources--global_log_receiver--properties--audit_logs.md", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["audit_logs"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/global_log_receiver/properties/audit_logs/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "audit_logs for xcsh_global_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# audit_logs

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md)
- [Property reference](data-sources--global_log_receiver--reference.md)
- audit_logs

<a id="section"></a>

Type: `["object", {}]`. Computed.

\[OneOf: audit\_logs, dns\_logs, request\_logs, security\_events\] Enable this option

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

OneOf alternatives in this subsection:

- [audit_logs](data-sources--global_log_receiver--properties--audit_logs.md#section)
- [dns_logs](data-sources--global_log_receiver--properties--dns_logs.md#section)
- [request_logs](data-sources--global_log_receiver--properties--request_logs.md#section)
- [security_events](data-sources--global_log_receiver--properties--security_events.md#section)

Select alternatives according to the provider validators above.

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](data-sources--global_log_receiver--reference.md)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md)
