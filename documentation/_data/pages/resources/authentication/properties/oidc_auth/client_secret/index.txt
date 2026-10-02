---
page_title: "oidc_auth.client_secret"
subcategory: ""
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["oidc auth client secret"], "body_bytes": 2240, "body_sha256": "sha256:eb975d45842238405ad6ea85bf6f2364e0134df284f792fc574d437926a948a8", "capabilities": ["identity"], "category": "identity", "child_ids": ["xcsh-docs:resources:authentication:properties:oidc_auth:client_secret:blindfold_secret_info", "xcsh-docs:resources:authentication:properties:oidc_auth:client_secret:clear_secret_info"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:authentication:collection", "completeness": "complete", "id": "xcsh-docs:resources:authentication:properties:oidc_auth:client_secret", "parent_id": "xcsh-docs:resources:authentication:properties:oidc_auth", "path": "documentation/resources/authentication/properties/oidc_auth/client_secret/index.md", "product": "distributed-cloud", "provider_name": "authentication", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1320230130103302-2003312321123101-3113322013121030-2001000113312133-1110211313130221-2210321120320211-2002022121312232-2333123113001213", "registry_path": "docs/guides/resources--authentication--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "oidc_auth.client_secret:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:authentication:properties:oidc_auth:client_secret:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "oidc_auth.client_secret:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:authentication:properties:oidc_auth:client_secret:clear_secret_info", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["oidc_auth", "client_secret"], "schema_version": 1, "sections": [{"aliases": ["blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:resources:authentication:properties:oidc_auth:client_secret:blindfold_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-oidc_auth--client_secret--blindfold_secret_info--location", "enforcement": "provider-schema", "group": "oidc_auth.client_secret.blindfold_secret_info:RequiredObjectAttributes:location", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:authentication:properties:oidc_auth:client_secret:blindfold_secret_info", "type": "requires"}], "schema_path": ["oidc_auth", "client_secret", "blindfold_secret_info"], "syntax": "block", "type": "object"}, {"aliases": ["clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:resources:authentication:properties:oidc_auth:client_secret:clear_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-oidc_auth--client_secret--clear_secret_info--url", "enforcement": "provider-schema", "group": "oidc_auth.client_secret.clear_secret_info:RequiredObjectAttributes:url", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:authentication:properties:oidc_auth:client_secret:clear_secret_info", "type": "requires"}], "schema_path": ["oidc_auth", "client_secret", "clear_secret_info"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/authentication/properties/oidc_auth/client_secret/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["authentication", "configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["authenticationCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# oidc_auth.client_secret

Breadcrumbs:

- [xcsh_authentication](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authentication/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authentication/properties/)
- [oidc_auth](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authentication/properties/oidc_auth/)
- oidc_auth.client_secret

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
client_secret {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authentication/properties/oidc_auth/client_secret/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authentication/properties/oidc_auth/client_secret/clear_secret_info/): complete subsection reference.

## Next pages

- [oidc_auth.client_secret.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authentication/properties/oidc_auth/client_secret/blindfold_secret_info/)
- [oidc_auth.client_secret.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authentication/properties/oidc_auth/client_secret/clear_secret_info/)
- [oidc_auth](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authentication/properties/oidc_auth/)
- [xcsh_authentication](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authentication/)
