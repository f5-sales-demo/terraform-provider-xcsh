---
page_title: "response_headers_to_add.secret_value"
subcategory: ""
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["response headers to add secret value"], "body_bytes": 2377, "body_sha256": "sha256:b7a60a8e80ff7f80cecd1ce47cdbe675d6e897bdc5b80e55ee2f952499f787ac", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:virtual_host:properties:response_headers_to_add:secret_value:blindfold_secret_info", "xcsh-docs:resources:virtual_host:properties:response_headers_to_add:secret_value:clear_secret_info"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_host:properties:response_headers_to_add:secret_value", "parent_id": "xcsh-docs:resources:virtual_host:properties:response_headers_to_add", "path": "documentation/resources/virtual_host/properties/response_headers_to_add/secret_value/index.md", "product": "distributed-cloud", "provider_name": "virtual_host", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2111212032020003-3033322032323303-2321321231321000-1310211200131101-1030333200010020-3101110203003230-1013011320132113-3312210222133312", "registry_path": "docs/guides/resources--virtual_host--reference--group-003.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "response_headers_to_add.secret_value:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:response_headers_to_add:secret_value:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "response_headers_to_add.secret_value:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:response_headers_to_add:secret_value:clear_secret_info", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["response_headers_to_add", "secret_value"], "schema_version": 1, "sections": [{"aliases": ["blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:resources:virtual_host:properties:response_headers_to_add:secret_value:blindfold_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-response_headers_to_add--secret_value--blindfold_secret_info--location", "enforcement": "provider-schema", "group": "response_headers_to_add.secret_value.blindfold_secret_info:RequiredObjectAttributes:location", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:response_headers_to_add:secret_value:blindfold_secret_info", "type": "requires"}], "schema_path": ["response_headers_to_add", "secret_value", "blindfold_secret_info"], "syntax": "block", "type": "object"}, {"aliases": ["clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:resources:virtual_host:properties:response_headers_to_add:secret_value:clear_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-response_headers_to_add--secret_value--clear_secret_info--url", "enforcement": "provider-schema", "group": "response_headers_to_add.secret_value.clear_secret_info:RequiredObjectAttributes:url", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:response_headers_to_add:secret_value:clear_secret_info", "type": "requires"}], "schema_path": ["response_headers_to_add", "secret_value", "clear_secret_info"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_host/properties/response_headers_to_add/secret_value/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# response_headers_to_add.secret_value

Breadcrumbs:

- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/)
- [response_headers_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/response_headers_to_add/)
- response_headers_to_add.secret_value

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
secret_value {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/response_headers_to_add/secret_value/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/response_headers_to_add/secret_value/clear_secret_info/): complete subsection reference.

## Next pages

- [response_headers_to_add.secret_value.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/response_headers_to_add/secret_value/blindfold_secret_info/)
- [response_headers_to_add.secret_value.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/response_headers_to_add/secret_value/clear_secret_info/)
- [response_headers_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/response_headers_to_add/)
- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/)
