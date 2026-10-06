---
page_title: "authentication.cookie_params.auth_hmac.prim_key"
subcategory: ""
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["authentication cookie params auth hmac prim key"], "body_bytes": 2074, "body_sha256": "sha256:5d76fe0d76f4627c1021610ceae49ee7778256e7188ddc51003fe3efb4496857", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:virtual_host:properties:authentication:cookie_params:auth_hmac:prim_key:blindfold_secret_info", "xcsh-docs:resources:virtual_host:properties:authentication:cookie_params:auth_hmac:prim_key:clear_secret_info"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_host:properties:authentication:cookie_params:auth_hmac:prim_key", "parent_id": "xcsh-docs:resources:virtual_host:properties:authentication:cookie_params:auth_hmac", "path": "documentation/resources/virtual_host/properties/authentication/cookie_params/auth_hmac/prim_key/index.md", "product": "distributed-cloud", "provider_name": "virtual_host", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-3200010310233203-1210331030223203-1011131031030201-0000131311322302-3132300210320223-0302302113112202-3012111312112213-3221330010132221", "registry_path": "docs/guides/resources--virtual_host--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "authentication.cookie_params.auth_hmac.prim_key:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:authentication:cookie_params:auth_hmac:prim_key:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "authentication.cookie_params.auth_hmac.prim_key:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:authentication:cookie_params:auth_hmac:prim_key:clear_secret_info", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["authentication", "cookie_params", "auth_hmac", "prim_key"], "schema_version": 1, "sections": [{"aliases": ["authentication cookie params auth hmac prim key blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:resources:virtual_host:properties:authentication:cookie_params:auth_hmac:prim_key:blindfold_secret_info", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-authentication--cookie_params--auth_hmac--prim_key--blindfold_secret_info--location", "enforcement": "provider-schema", "group": "authentication.cookie_params.auth_hmac.prim_key.blindfold_secret_info:RequiredObjectAttributes:location", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:authentication:cookie_params:auth_hmac:prim_key:blindfold_secret_info", "type": "requires"}], "schema_path": ["authentication", "cookie_params", "auth_hmac", "prim_key", "blindfold_secret_info"], "syntax": "block", "type": "object"}, {"aliases": ["authentication cookie params auth hmac prim key clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:resources:virtual_host:properties:authentication:cookie_params:auth_hmac:prim_key:clear_secret_info", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-authentication--cookie_params--auth_hmac--prim_key--clear_secret_info--url", "enforcement": "provider-schema", "group": "authentication.cookie_params.auth_hmac.prim_key.clear_secret_info:RequiredObjectAttributes:url", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:authentication:cookie_params:auth_hmac:prim_key:clear_secret_info", "type": "requires"}], "schema_path": ["authentication", "cookie_params", "auth_hmac", "prim_key", "clear_secret_info"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_host/properties/authentication/cookie_params/auth_hmac/prim_key/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# authentication.cookie_params.auth_hmac.prim_key

Breadcrumbs:

- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/)
- [authentication](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/)
- [authentication.cookie_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/cookie_params/)
- [authentication.cookie_params.auth_hmac](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/cookie_params/auth_hmac/)
- authentication.cookie_params.auth_hmac.prim_key

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
prim_key {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/cookie_params/auth_hmac/prim_key/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/cookie_params/auth_hmac/prim_key/clear_secret_info/): complete subsection reference.
