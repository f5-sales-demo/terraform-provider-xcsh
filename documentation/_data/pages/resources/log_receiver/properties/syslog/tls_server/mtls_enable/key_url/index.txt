---
page_title: "syslog.tls_server.mtls_enable.key_url"
subcategory: "Monitoring"
description: "syslog.tls_server.mtls_enable.key_url for xcsh_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 2646, "body_sha256": "sha256:b0e66f149a73afde747af9ce8dccf50bdac1d810bf386ea0e64d2b81bbcc30fa", "child_ids": ["xcsh-docs:resources:log_receiver:properties:syslog:tls_server:mtls_enable:key_url:blindfold_secret_info", "xcsh-docs:resources:log_receiver:properties:syslog:tls_server:mtls_enable:key_url:clear_secret_info"], "collection_id": "xcsh-docs:resources:log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:log_receiver:properties:syslog:tls_server:mtls_enable:key_url", "parent_id": "xcsh-docs:resources:log_receiver:properties:syslog:tls_server:mtls_enable", "path": "documentation/resources/log_receiver/properties/syslog/tls_server/mtls_enable/key_url/index.md", "provider_name": "log_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["syslog", "tls_server", "mtls_enable", "key_url"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/log_receiver/properties/syslog/tls_server/mtls_enable/key_url/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "syslog.tls_server.mtls_enable.key_url for xcsh_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# syslog.tls_server.mtls_enable.key_url

Breadcrumbs:

- [xcsh_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/log_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/log_receiver/properties/)
- [syslog](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/log_receiver/properties/syslog/)
- [syslog.tls_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/log_receiver/properties/syslog/tls_server/)
- [syslog.tls_server.mtls_enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/log_receiver/properties/syslog/tls_server/mtls_enable/)
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

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/log_receiver/properties/syslog/tls_server/mtls_enable/key_url/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/log_receiver/properties/syslog/tls_server/mtls_enable/key_url/clear_secret_info/): complete subsection reference.

## Next pages

- [syslog.tls_server.mtls_enable.key_url.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/log_receiver/properties/syslog/tls_server/mtls_enable/key_url/blindfold_secret_info/)
- [syslog.tls_server.mtls_enable.key_url.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/log_receiver/properties/syslog/tls_server/mtls_enable/key_url/clear_secret_info/)
- [syslog.tls_server.mtls_enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/log_receiver/properties/syslog/tls_server/mtls_enable/)
- [xcsh_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/log_receiver/)
