---
page_title: "authentication"
subcategory: ""
description: "Authentication related information. This allows to configure the URL to redirect after the authentication Authentication Object Reference, configuration of cookie params etc."
xcsh_docs: {"aliases": ["authentication", "credential setup", "credentials"], "body_bytes": 4681, "body_sha256": "sha256:eb0e49c71d1e92364b36509cd85e0428327cac5e12932ea9a85400eed869a34e", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:virtual_host:properties:authentication:auth_config", "xcsh-docs:data-sources:virtual_host:properties:authentication:cookie_params", "xcsh-docs:data-sources:virtual_host:properties:authentication:redirect_dynamic", "xcsh-docs:data-sources:virtual_host:properties:authentication:use_auth_object_config"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:virtual_host:properties:authentication", "parent_id": "xcsh-docs:data-sources:virtual_host:reference", "path": "documentation/data-sources/virtual_host/properties/authentication/index.md", "product": "distributed-cloud", "provider_name": "virtual_host", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0111212013333231-0233303000120330-1321320213133323-3312301102111320-1120003002132332-0223012232003100-0203021203333003-2011300212121200", "registry_path": "docs/guides/data-sources--virtual_host--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["authentication"], "schema_version": 1, "sections": [{"aliases": ["auth config", "authentication", "credential setup", "credentials"], "anchor": "section", "description": "Reference to Authentication Config Object.", "document_id": "xcsh-docs:data-sources:virtual_host:properties:authentication:auth_config", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["authentication", "auth_config"], "syntax": "attribute", "type": "object"}, {"aliases": ["authentication", "cookie params", "credential setup", "credentials"], "anchor": "section", "description": "Specifies different cookie related config parameters for authentication.", "document_id": "xcsh-docs:data-sources:virtual_host:properties:authentication:cookie_params", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["authentication", "cookie_params"], "syntax": "attribute", "type": "object"}, {"aliases": ["redirect dynamic"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:virtual_host:properties:authentication:redirect_dynamic", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["authentication", "redirect_dynamic"], "syntax": "attribute", "type": "object"}, {"aliases": ["redirect url"], "anchor": "schema-authentication--redirect_url", "description": "Exclusive with user can provide a URL for e.g https://abc.xyz.com where user gets redirected. This URL configured here must match with the redirect URL configured with the OIDC provider.", "document_id": "xcsh-docs:data-sources:virtual_host:properties:authentication", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["authentication", "redirect_url"], "syntax": "attribute", "type": "string"}, {"aliases": ["use auth object config"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:virtual_host:properties:authentication:use_auth_object_config", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["authentication", "use_auth_object_config"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/virtual_host/properties/authentication/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Authentication related information. This allows to configure the URL to redirect after the authentication Authentication Object Reference, configuration of cookie params etc.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# authentication

Breadcrumbs:

- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/)
- authentication

<a id="section"></a>

Type: `"single"`. Computed.

\[OneOf: authentication, no\_authentication; Default: no\_authentication\] Authentication related
information. This allows to configure the URL to redirect after the authentication Authentication
Object Reference, configuration of cookie params etc.

Upstream description:

Authentication related information. This allows to configure the URL to redirect after the
authentication Authentication Object Reference, configuration of cookie params etc.

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

- [authentication](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/authentication/#section)
- [no_authentication](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/no_authentication/#section)

Select alternatives according to the provider validators above.

## Direct properties

- [auth_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/authentication/auth_config/): complete subsection reference.

- [cookie_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/authentication/cookie_params/): complete subsection reference.

- [redirect_dynamic](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/authentication/redirect_dynamic/): complete subsection reference.

<a id="schema-authentication--redirect_url"></a>

### redirect_url property

Type: `"string"`. Computed.

Exclusive with \[redirect\_dynamic\] user can provide a URL for e.g https&#58;//abc.xyz.com where
user gets redirected. This URL configured here must match with the redirect URL configured with the
OIDC provider.

Upstream description:

Exclusive with \[redirect\_dynamic\]

user can provide a URL for e.g https&#58;//abc.xyz.com where user gets redirected. This URL
configured here must match with the redirect URL configured with the OIDC provider.

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

- [use_auth_object_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/authentication/use_auth_object_config/): complete subsection reference.

## Next pages

- [authentication.auth_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/authentication/auth_config/)
- [authentication.cookie_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/authentication/cookie_params/)
- [authentication.redirect_dynamic](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/authentication/redirect_dynamic/)
- [authentication.use_auth_object_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/authentication/use_auth_object_config/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/)
- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/)
