---
page_title: "authentication.cookie_params.auth_hmac.sec_key"
subcategory: ""
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["authentication", "authentication cookie params auth hmac sec key", "credential setup", "credentials"], "body_bytes": 2792, "body_sha256": "sha256:2f5db8ddf0297a543f7e1638f1bb390117edffa5974c737e7ce077744883d96b", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:virtual_host:properties:authentication:cookie_params:auth_hmac:sec_key:blindfold_secret_info", "xcsh-docs:resources:virtual_host:properties:authentication:cookie_params:auth_hmac:sec_key:clear_secret_info"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_host:properties:authentication:cookie_params:auth_hmac:sec_key", "parent_id": "xcsh-docs:resources:virtual_host:properties:authentication:cookie_params:auth_hmac", "path": "documentation/resources/virtual_host/properties/authentication/cookie_params/auth_hmac/sec_key/index.md", "product": "distributed-cloud", "provider_name": "virtual_host", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1020112220213021-2013010221201110-1201030001102333-3203321222232303-1312012322330201-2030010101010333-0233031002023333-0022232113021301", "registry_path": "docs/guides/resources--virtual_host--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "authentication.cookie_params.auth_hmac.sec_key:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:authentication:cookie_params:auth_hmac:sec_key:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "authentication.cookie_params.auth_hmac.sec_key:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:authentication:cookie_params:auth_hmac:sec_key:clear_secret_info", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["authentication", "cookie_params", "auth_hmac", "sec_key"], "schema_version": 1, "sections": [{"aliases": ["blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:resources:virtual_host:properties:authentication:cookie_params:auth_hmac:sec_key:blindfold_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-authentication--cookie_params--auth_hmac--sec_key--blindfold_secret_info--location", "enforcement": "provider-schema", "group": "authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info:RequiredObjectAttributes:location", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:authentication:cookie_params:auth_hmac:sec_key:blindfold_secret_info", "type": "requires"}], "schema_path": ["authentication", "cookie_params", "auth_hmac", "sec_key", "blindfold_secret_info"], "syntax": "block", "type": "object"}, {"aliases": ["clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:resources:virtual_host:properties:authentication:cookie_params:auth_hmac:sec_key:clear_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-authentication--cookie_params--auth_hmac--sec_key--clear_secret_info--url", "enforcement": "provider-schema", "group": "authentication.cookie_params.auth_hmac.sec_key.clear_secret_info:RequiredObjectAttributes:url", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:authentication:cookie_params:auth_hmac:sec_key:clear_secret_info", "type": "requires"}], "schema_path": ["authentication", "cookie_params", "auth_hmac", "sec_key", "clear_secret_info"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_host/properties/authentication/cookie_params/auth_hmac/sec_key/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# authentication.cookie_params.auth_hmac.sec_key

Breadcrumbs:

- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/)
- [authentication](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/)
- [authentication.cookie_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/cookie_params/)
- [authentication.cookie_params.auth_hmac](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/cookie_params/auth_hmac/)
- authentication.cookie_params.auth_hmac.sec_key

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
sec_key {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/cookie_params/auth_hmac/sec_key/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/cookie_params/auth_hmac/sec_key/clear_secret_info/): complete subsection reference.

## Next pages

- [authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/cookie_params/auth_hmac/sec_key/blindfold_secret_info/)
- [authentication.cookie_params.auth_hmac.sec_key.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/cookie_params/auth_hmac/sec_key/clear_secret_info/)
- [authentication.cookie_params.auth_hmac](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/cookie_params/auth_hmac/)
- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/)
