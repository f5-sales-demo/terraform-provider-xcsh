---
page_title: "cookie_params.auth_hmac.prim_key"
subcategory: ""
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["cookie params auth hmac prim key"], "body_bytes": 1842, "body_sha256": "sha256:08842642a69e84c3833590f02711fa7eeffa9cb509d77e0557107a368c3ab624", "capabilities": ["identity"], "category": "identity", "child_ids": ["xcsh-docs:resources:authentication:properties:cookie_params:auth_hmac:prim_key:blindfold_secret_info", "xcsh-docs:resources:authentication:properties:cookie_params:auth_hmac:prim_key:clear_secret_info"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:authentication:collection", "completeness": "complete", "id": "xcsh-docs:resources:authentication:properties:cookie_params:auth_hmac:prim_key", "parent_id": "xcsh-docs:resources:authentication:properties:cookie_params:auth_hmac", "path": "documentation/resources/authentication/properties/cookie_params/auth_hmac/prim_key/index.md", "product": "distributed-cloud", "provider_name": "authentication", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "resources", "registry_anchor": "canonical-3032121123310031-0232123133330200-3020103210122032-1020011331320323-3300310120322032-2321011320320102-0302123300213023-1221300110301223", "registry_path": "docs/guides/resources--authentication--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "cookie_params.auth_hmac.prim_key:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:authentication:properties:cookie_params:auth_hmac:prim_key:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cookie_params.auth_hmac.prim_key:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:authentication:properties:cookie_params:auth_hmac:prim_key:clear_secret_info", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["cookie_params", "auth_hmac", "prim_key"], "schema_version": 1, "sections": [{"aliases": ["cookie params auth hmac prim key blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:resources:authentication:properties:cookie_params:auth_hmac:prim_key:blindfold_secret_info", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-cookie_params--auth_hmac--prim_key--blindfold_secret_info--location", "enforcement": "provider-schema", "group": "cookie_params.auth_hmac.prim_key.blindfold_secret_info:RequiredObjectAttributes:location", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:authentication:properties:cookie_params:auth_hmac:prim_key:blindfold_secret_info", "type": "requires"}], "schema_path": ["cookie_params", "auth_hmac", "prim_key", "blindfold_secret_info"], "syntax": "block", "type": "object"}, {"aliases": ["cookie params auth hmac prim key clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:resources:authentication:properties:cookie_params:auth_hmac:prim_key:clear_secret_info", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-cookie_params--auth_hmac--prim_key--clear_secret_info--url", "enforcement": "provider-schema", "group": "cookie_params.auth_hmac.prim_key.clear_secret_info:RequiredObjectAttributes:url", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:authentication:properties:cookie_params:auth_hmac:prim_key:clear_secret_info", "type": "requires"}], "schema_path": ["cookie_params", "auth_hmac", "prim_key", "clear_secret_info"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/authentication/properties/cookie_params/auth_hmac/prim_key/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["authentication", "configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["authenticationCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cookie_params.auth_hmac.prim_key

Breadcrumbs:

- [xcsh_authentication](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authentication/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authentication/properties/)
- [cookie_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authentication/properties/cookie_params/)
- [cookie_params.auth_hmac](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authentication/properties/cookie_params/auth_hmac/)
- cookie_params.auth_hmac.prim_key

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

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authentication/properties/cookie_params/auth_hmac/prim_key/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authentication/properties/cookie_params/auth_hmac/prim_key/clear_secret_info/): complete subsection reference.
