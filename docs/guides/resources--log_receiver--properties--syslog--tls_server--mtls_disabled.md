---
page_title: "syslog.tls_server.mtls_disabled"
subcategory: "Monitoring"
description: "syslog.tls_server.mtls_disabled for xcsh_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 1071, "body_sha256": "sha256:218c1b9ff797f4d7c5dba855f773e9d1135a496ddbb93db0228d03e3bef97dbf", "canonical_id": "xcsh-docs:resources:log_receiver:properties:syslog:tls_server:mtls_disabled", "child_ids": [], "collection_id": "xcsh-docs:resources:log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:log_receiver:properties:syslog:tls_server:mtls_disabled", "parent_id": "xcsh-docs:resources:log_receiver:properties:syslog:tls_server", "path": "docs/guides/resources--log_receiver--properties--syslog--tls_server--mtls_disabled.md", "provider_name": "log_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["syslog", "tls_server", "mtls_disabled"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/log_receiver/properties/syslog/tls_server/mtls_disabled/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "syslog.tls_server.mtls_disabled for xcsh_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# syslog.tls_server.mtls_disabled

Breadcrumbs:

- [xcsh_log_receiver](../resources/log_receiver.md)
- [Property reference](resources--log_receiver--reference.md)
- [syslog](resources--log_receiver--properties--syslog.md)
- [syslog.tls_server](resources--log_receiver--properties--syslog--tls_server.md)
- syslog.tls_server.mtls_disabled

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
mtls_disabled = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [syslog.tls_server](resources--log_receiver--properties--syslog--tls_server.md)
- [xcsh_log_receiver](../resources/log_receiver.md)
