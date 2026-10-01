---
page_title: "syslog.tls_server.mtls_enable.key_url"
subcategory: "Monitoring"
description: "syslog.tls_server.mtls_enable.key_url for xcsh_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 2104, "body_sha256": "sha256:3bf746553ccb4a721c963b35e9a6e78e53b596a0f03f247fe7fc3eff9c828024", "canonical_id": "xcsh-docs:resources:log_receiver:properties:syslog:tls_server:mtls_enable:key_url", "child_ids": ["xcsh-docs:resources:log_receiver:properties:syslog:tls_server:mtls_enable:key_url:blindfold_secret_info", "xcsh-docs:resources:log_receiver:properties:syslog:tls_server:mtls_enable:key_url:clear_secret_info"], "collection_id": "xcsh-docs:resources:log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:log_receiver:properties:syslog:tls_server:mtls_enable:key_url", "parent_id": "xcsh-docs:resources:log_receiver:properties:syslog:tls_server:mtls_enable", "path": "docs/guides/resources--log_receiver--properties--syslog--tls_server--mtls_enable--key_url.md", "provider_name": "log_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["syslog", "tls_server", "mtls_enable", "key_url"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/log_receiver/properties/syslog/tls_server/mtls_enable/key_url/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "syslog.tls_server.mtls_enable.key_url for xcsh_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# syslog.tls_server.mtls_enable.key_url

Breadcrumbs:

- [xcsh_log_receiver](../resources/log_receiver.md)
- [Property reference](resources--log_receiver--reference.md)
- [syslog](resources--log_receiver--properties--syslog.md)
- [syslog.tls_server](resources--log_receiver--properties--syslog--tls_server.md)
- [syslog.tls_server.mtls_enable](resources--log_receiver--properties--syslog--tls_server--mtls_enable.md)
- syslog.tls_server.mtls_enable.key_url

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
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
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
key_url {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blindfold_secret_info](resources--log_receiver--properties--syslog--tls_server--mtls_enable--key_url--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](resources--log_receiver--properties--syslog--tls_server--mtls_enable--key_url--clear_secret_info.md): complete subsection reference.

## Next pages

- [syslog.tls_server.mtls_enable.key_url.blindfold_secret_info](resources--log_receiver--properties--syslog--tls_server--mtls_enable--key_url--blindfold_secret_info.md)
- [syslog.tls_server.mtls_enable.key_url.clear_secret_info](resources--log_receiver--properties--syslog--tls_server--mtls_enable--key_url--clear_secret_info.md)
- [syslog.tls_server.mtls_enable](resources--log_receiver--properties--syslog--tls_server--mtls_enable.md)
- [xcsh_log_receiver](../resources/log_receiver.md)
