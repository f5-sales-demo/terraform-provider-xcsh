---
page_title: "syslog.tls_server.mtls_enable.key_url"
subcategory: "Monitoring"
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["syslog tls server mtls enable key url"], "body_bytes": 1977, "body_sha256": "sha256:783ab761c6c5c2eeb8f1dafc884f5d6072861c9d32677bcb31c4b6fda7121e27", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:resources:log_receiver:properties:syslog:tls_server:mtls_enable:key_url:blindfold_secret_info", "xcsh-docs:resources:log_receiver:properties:syslog:tls_server:mtls_enable:key_url:clear_secret_info"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:log_receiver:properties:syslog:tls_server:mtls_enable:key_url", "parent_id": "xcsh-docs:resources:log_receiver:properties:syslog:tls_server:mtls_enable", "path": "documentation/resources/log_receiver/properties/syslog/tls_server/mtls_enable/key_url/index.md", "product": "distributed-cloud", "provider_name": "log_receiver", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-3322210101013313-1312220332330211-2001102110333313-2202202111133301-1302332320331112-0231222021202210-2313322201113023-1213300111211310", "registry_path": "docs/guides/resources--log_receiver--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "syslog.tls_server.mtls_enable.key_url:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:log_receiver:properties:syslog:tls_server:mtls_enable:key_url:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "syslog.tls_server.mtls_enable.key_url:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:log_receiver:properties:syslog:tls_server:mtls_enable:key_url:clear_secret_info", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["syslog", "tls_server", "mtls_enable", "key_url"], "schema_version": 1, "sections": [{"aliases": ["syslog tls server mtls enable key url blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:resources:log_receiver:properties:syslog:tls_server:mtls_enable:key_url:blindfold_secret_info", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-syslog--tls_server--mtls_enable--key_url--blindfold_secret_info--location", "enforcement": "provider-schema", "group": "syslog.tls_server.mtls_enable.key_url.blindfold_secret_info:RequiredObjectAttributes:location", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:log_receiver:properties:syslog:tls_server:mtls_enable:key_url:blindfold_secret_info", "type": "requires"}], "schema_path": ["syslog", "tls_server", "mtls_enable", "key_url", "blindfold_secret_info"], "syntax": "block", "type": "object"}, {"aliases": ["syslog tls server mtls enable key url clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:resources:log_receiver:properties:syslog:tls_server:mtls_enable:key_url:clear_secret_info", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-syslog--tls_server--mtls_enable--key_url--clear_secret_info--url", "enforcement": "provider-schema", "group": "syslog.tls_server.mtls_enable.key_url.clear_secret_info:RequiredObjectAttributes:url", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:log_receiver:properties:syslog:tls_server:mtls_enable:key_url:clear_secret_info", "type": "requires"}], "schema_path": ["syslog", "tls_server", "mtls_enable", "key_url", "clear_secret_info"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/log_receiver/properties/syslog/tls_server/mtls_enable/key_url/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["log_receiverCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
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
EnumExtractionComplete: false
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
