---
page_title: "syslog.tls_server.volterra_ca"
subcategory: "Monitoring"
description: "syslog.tls_server.volterra_ca for xcsh_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 1051, "body_sha256": "sha256:e14e76d801b181a1fb4ac9e1bc0f96c3fd75c8e36835d71f47cad4075b2aed67", "canonical_id": "xcsh-docs:data-sources:log_receiver:properties:syslog:tls_server:volterra_ca", "child_ids": [], "collection_id": "xcsh-docs:data-sources:log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:log_receiver:properties:syslog:tls_server:volterra_ca", "parent_id": "xcsh-docs:data-sources:log_receiver:properties:syslog:tls_server", "path": "docs/guides/data-sources--log_receiver--properties--syslog--tls_server--volterra_ca.md", "provider_name": "log_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["syslog", "tls_server", "volterra_ca"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/log_receiver/properties/syslog/tls_server/volterra_ca/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "syslog.tls_server.volterra_ca for xcsh_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# syslog.tls_server.volterra_ca

Breadcrumbs:

- [xcsh_log_receiver](../data-sources/log_receiver.md)
- [Property reference](data-sources--log_receiver--reference.md)
- [syslog](data-sources--log_receiver--properties--syslog.md)
- [syslog.tls_server](data-sources--log_receiver--properties--syslog--tls_server.md)
- syslog.tls_server.volterra_ca

<a id="section"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for volterra ca.

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

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [syslog.tls_server](data-sources--log_receiver--properties--syslog--tls_server.md)
- [xcsh_log_receiver](../data-sources/log_receiver.md)
