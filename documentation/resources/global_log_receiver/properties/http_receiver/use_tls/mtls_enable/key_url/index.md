---
page_title: "http_receiver.use_tls.mtls_enable.key_url"
subcategory: ""
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["http receiver use tls mtls enable key url"], "body_bytes": 2079, "body_sha256": "sha256:f53dd6aab6065926292f8bd3725ac25695d4dceebb2bbf0cd8f444a3b37b23b6", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:resources:global_log_receiver:properties:http_receiver:use_tls:mtls_enable:key_url:blindfold_secret_info", "xcsh-docs:resources:global_log_receiver:properties:http_receiver:use_tls:mtls_enable:key_url:clear_secret_info"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver:use_tls:mtls_enable:key_url", "parent_id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver:use_tls:mtls_enable", "path": "documentation/resources/global_log_receiver/properties/http_receiver/use_tls/mtls_enable/key_url/index.md", "product": "distributed-cloud", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-0322200101111022-3312322200113313-0003312032113332-3310112300030021-2322101032013231-1000311330012110-0100023213131132-1023201201202002", "registry_path": "docs/guides/resources--global_log_receiver--reference--group-003.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "http_receiver.use_tls.mtls_enable.key_url:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver:use_tls:mtls_enable:key_url:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "http_receiver.use_tls.mtls_enable.key_url:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver:use_tls:mtls_enable:key_url:clear_secret_info", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["http_receiver", "use_tls", "mtls_enable", "key_url"], "schema_version": 1, "sections": [{"aliases": ["http receiver use tls mtls enable key url blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver:use_tls:mtls_enable:key_url:blindfold_secret_info", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-http_receiver--use_tls--mtls_enable--key_url--blindfold_secret_info--location", "enforcement": "provider-schema", "group": "http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info:RequiredObjectAttributes:location", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver:use_tls:mtls_enable:key_url:blindfold_secret_info", "type": "requires"}], "schema_path": ["http_receiver", "use_tls", "mtls_enable", "key_url", "blindfold_secret_info"], "syntax": "block", "type": "object"}, {"aliases": ["http receiver use tls mtls enable key url clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver:use_tls:mtls_enable:key_url:clear_secret_info", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-http_receiver--use_tls--mtls_enable--key_url--clear_secret_info--url", "enforcement": "provider-schema", "group": "http_receiver.use_tls.mtls_enable.key_url.clear_secret_info:RequiredObjectAttributes:url", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver:use_tls:mtls_enable:key_url:clear_secret_info", "type": "requires"}], "schema_path": ["http_receiver", "use_tls", "mtls_enable", "key_url", "clear_secret_info"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/http_receiver/use_tls/mtls_enable/key_url/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# http_receiver.use_tls.mtls_enable.key_url

Breadcrumbs:

- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/)
- [http_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/)
- [http_receiver.use_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/use_tls/)
- [http_receiver.use_tls.mtls_enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/use_tls/mtls_enable/)
- http_receiver.use_tls.mtls_enable.key_url

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

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/use_tls/mtls_enable/key_url/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/use_tls/mtls_enable/key_url/clear_secret_info/): complete subsection reference.
