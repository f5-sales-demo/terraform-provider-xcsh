---
page_title: "authentication"
subcategory: ""
description: "Authentication related information. This allows to configure the URL to redirect after the authentication Authentication Object Reference, configuration of cookie params etc."
xcsh_docs: {"aliases": ["authentication", "credential setup", "credentials"], "body_bytes": 5225, "body_sha256": "sha256:146450352219c92c14f470c157ac6a1832bae28c0c9647d5bdef2ef3da2791e9", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:virtual_host:properties:authentication:auth_config", "xcsh-docs:resources:virtual_host:properties:authentication:cookie_params", "xcsh-docs:resources:virtual_host:properties:authentication:redirect_dynamic", "xcsh-docs:resources:virtual_host:properties:authentication:use_auth_object_config"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_host:properties:authentication", "parent_id": "xcsh-docs:resources:virtual_host:reference", "path": "documentation/resources/virtual_host/properties/authentication/index.md", "product": "distributed-cloud", "provider_name": "virtual_host", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2003102221320033-2002121330313023-0331120011300333-0213101002200303-3003330033121012-0330110020200302-3233233301001020-2111210130323123", "registry_path": "docs/guides/resources--virtual_host--reference--group-001.md", "relationships": [{"anchor": "schema-authentication--redirect_url", "enforcement": "provider-schema", "group": "authentication:ConflictingObjectAttributes:redirect_dynamic,redirect_url", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:authentication", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "authentication:ConflictingObjectAttributes:cookie_params,use_auth_object_config", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:authentication:cookie_params", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "authentication:ConflictingObjectAttributes:redirect_dynamic,redirect_url", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:authentication:redirect_dynamic", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "authentication:ConflictingObjectAttributes:cookie_params,use_auth_object_config", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:authentication:use_auth_object_config", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "authentication:RequiredObjectAttributes:auth_config", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:authentication:auth_config", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["authentication"], "schema_version": 1, "sections": [{"aliases": ["authentication auth config"], "anchor": "section", "description": "Reference to Authentication Config Object.", "document_id": "xcsh-docs:resources:virtual_host:properties:authentication:auth_config", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["authentication", "auth_config"], "syntax": "block", "type": "object"}, {"aliases": ["authentication cookie params"], "anchor": "section", "description": "Specifies different cookie related config parameters for authentication.", "document_id": "xcsh-docs:resources:virtual_host:properties:authentication:cookie_params", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "authentication.cookie_params:ConflictingObjectAttributes:auth_hmac,kms_key_hmac", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:authentication:cookie_params:auth_hmac", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "authentication.cookie_params:ConflictingObjectAttributes:auth_hmac,kms_key_hmac", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:authentication:cookie_params:kms_key_hmac", "type": "conflicts"}], "schema_path": ["authentication", "cookie_params"], "syntax": "block", "type": "object"}, {"aliases": ["authentication redirect dynamic"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:virtual_host:properties:authentication:redirect_dynamic", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["authentication", "redirect_dynamic"], "syntax": "attribute", "type": "object"}, {"aliases": ["authentication redirect url"], "anchor": "schema-authentication--redirect_url", "description": "Exclusive with user can provide a URL for e.g https://abc.xyz.com where user gets redirected. This URL configured here must match with the redirect URL configured with the OIDC provider.", "document_id": "xcsh-docs:resources:virtual_host:properties:authentication", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["authentication", "redirect_url"], "syntax": "attribute", "type": "string"}, {"aliases": ["authentication use auth object config"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:virtual_host:properties:authentication:use_auth_object_config", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["authentication", "use_auth_object_config"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_host/properties/authentication/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Authentication related information. This allows to configure the URL to redirect after the authentication Authentication Object Reference, configuration of cookie params etc.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# authentication

Breadcrumbs:

- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/)
- authentication

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: authentication, no\_authentication; Default: no\_authentication\] Authentication related
information. This allows to configure the URL to redirect after the authentication Authentication
Object Reference, configuration of cookie params etc.

Upstream description:

Authentication related information. This allows to configure the URL to redirect after the
authentication Authentication Object Reference, configuration of cookie params etc.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("auth_config"),
  validators.ConflictingObjectAttributes("cookie_params",
    "use_auth_object_config"),
  validators.ConflictingObjectAttributes("redirect_dynamic",
    "redirect_url")}
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
  "x-ves-oneof-field-cookie_params_choice": "[\"cookie_params\",\"use_auth_object_config\"]",
  "x-ves-oneof-field-redirect_url_choice": "[\"redirect_dynamic\",\"redirect_url\"]"
}
```

OneOf alternatives in this subsection:

- [authentication](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/#section)
- [no_authentication](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/no_authentication/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
authentication {
  # Configure direct properties listed below.
}
```

## Direct properties

- [auth_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/auth_config/): complete subsection reference.

- [cookie_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/cookie_params/): complete subsection reference.

- [redirect_dynamic](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/redirect_dynamic/): complete subsection reference.

<a id="schema-authentication--redirect_url"></a>

### redirect_url property

Type: `"string"`. Optional.

Exclusive with \[redirect\_dynamic\] user can provide a URL for e.g https&#58;//abc.xyz.com where
user gets redirected. This URL configured here must match with the redirect URL configured with the
OIDC provider.

Upstream description:

Exclusive with \[redirect\_dynamic\]

user can provide a URL for e.g https&#58;//abc.xyz.com where user gets redirected. This URL
configured here must match with the redirect URL configured with the OIDC provider.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

- [use_auth_object_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/use_auth_object_config/): complete subsection reference.

## Next pages

- [authentication.auth_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/auth_config/)
- [authentication.cookie_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/cookie_params/)
- [authentication.redirect_dynamic](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/redirect_dynamic/)
- [authentication.use_auth_object_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/use_auth_object_config/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/)
- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/)
