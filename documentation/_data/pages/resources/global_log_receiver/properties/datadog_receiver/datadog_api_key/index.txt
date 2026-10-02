---
page_title: "datadog_receiver.datadog_api_key"
subcategory: ""
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["datadog receiver datadog api key"], "body_bytes": 2397, "body_sha256": "sha256:6137f1437d822411b4cacab5c04d51b6321b4e2266b9b880773b6ed0c186167f", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:resources:global_log_receiver:properties:datadog_receiver:datadog_api_key:blindfold_secret_info", "xcsh-docs:resources:global_log_receiver:properties:datadog_receiver:datadog_api_key:clear_secret_info"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:datadog_receiver:datadog_api_key", "parent_id": "xcsh-docs:resources:global_log_receiver:properties:datadog_receiver", "path": "documentation/resources/global_log_receiver/properties/datadog_receiver/datadog_api_key/index.md", "product": "distributed-cloud", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3003331032300003-3113111102303300-0022111111131031-2220230330110113-1120122230101110-2030202100202320-3232330113023022-0231001111213310", "registry_path": "docs/guides/resources--global_log_receiver--reference--group-002.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "datadog_receiver.datadog_api_key:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:datadog_receiver:datadog_api_key:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "datadog_receiver.datadog_api_key:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:datadog_receiver:datadog_api_key:clear_secret_info", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["datadog_receiver", "datadog_api_key"], "schema_version": 1, "sections": [{"aliases": ["blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:datadog_receiver:datadog_api_key:blindfold_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-datadog_receiver--datadog_api_key--blindfold_secret_info--location", "enforcement": "provider-schema", "group": "datadog_receiver.datadog_api_key.blindfold_secret_info:RequiredObjectAttributes:location", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:datadog_receiver:datadog_api_key:blindfold_secret_info", "type": "requires"}], "schema_path": ["datadog_receiver", "datadog_api_key", "blindfold_secret_info"], "syntax": "block", "type": "object"}, {"aliases": ["clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:datadog_receiver:datadog_api_key:clear_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-datadog_receiver--datadog_api_key--clear_secret_info--url", "enforcement": "provider-schema", "group": "datadog_receiver.datadog_api_key.clear_secret_info:RequiredObjectAttributes:url", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:datadog_receiver:datadog_api_key:clear_secret_info", "type": "requires"}], "schema_path": ["datadog_receiver", "datadog_api_key", "clear_secret_info"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/datadog_receiver/datadog_api_key/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# datadog_receiver.datadog_api_key

Breadcrumbs:

- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/)
- [datadog_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/datadog_receiver/)
- datadog_receiver.datadog_api_key

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
datadog_api_key {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/datadog_receiver/datadog_api_key/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/datadog_receiver/datadog_api_key/clear_secret_info/): complete subsection reference.

## Next pages

- [datadog_receiver.datadog_api_key.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/datadog_receiver/datadog_api_key/blindfold_secret_info/)
- [datadog_receiver.datadog_api_key.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/datadog_receiver/datadog_api_key/clear_secret_info/)
- [datadog_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/datadog_receiver/)
- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/)
