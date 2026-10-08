---
page_title: "authentication"
subcategory: ""
description: "Authentication related information. This allows to configure the URL to redirect after the authentication Authentication Object Reference, configuration of cookie params etc."
xcsh_docs: {"aliases": ["authentication", "credential setup", "credentials"], "body_bytes": 3356, "body_sha256": "sha256:fd469c6067ec53b0af2e6a1474c10e6953e891742dfc1ef8f54c00eb8b030554", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:virtual_host:properties:authentication:auth_config", "xcsh-docs:data-sources:virtual_host:properties:authentication:cookie_params", "xcsh-docs:data-sources:virtual_host:properties:authentication:redirect_dynamic", "xcsh-docs:data-sources:virtual_host:properties:authentication:use_auth_object_config"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:virtual_host:properties:authentication", "parent_id": "xcsh-docs:data-sources:virtual_host:reference", "path": "documentation/data-sources/virtual_host/properties/authentication/index.md", "product": "distributed-cloud", "provider_name": "virtual_host", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-0111212013333231-0233303000120330-1321320213133323-3312301102111320-1120003002132332-0223012232003100-0203021203333003-2011300212121200", "registry_path": "docs/guides/data-sources--virtual_host--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["authentication"], "schema_version": 1, "sections": [{"aliases": ["authentication auth config"], "anchor": "section", "description": "Reference to Authentication Config Object.", "document_id": "xcsh-docs:data-sources:virtual_host:properties:authentication:auth_config", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["authentication", "auth_config"], "syntax": "attribute", "type": "object"}, {"aliases": ["authentication cookie params"], "anchor": "section", "description": "Specifies different cookie related config parameters for authentication.", "document_id": "xcsh-docs:data-sources:virtual_host:properties:authentication:cookie_params", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["authentication", "cookie_params"], "syntax": "attribute", "type": "object"}, {"aliases": ["authentication redirect dynamic"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:virtual_host:properties:authentication:redirect_dynamic", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["authentication", "redirect_dynamic"], "syntax": "attribute", "type": "object"}, {"aliases": ["authentication redirect url"], "anchor": "schema-authentication--redirect_url", "description": "Exclusive with user can provide a URL for e.g https://abc.xyz.com where user gets redirected. This URL configured here must match with the redirect URL configured with the OIDC provider.", "document_id": "xcsh-docs:data-sources:virtual_host:properties:authentication", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["authentication", "redirect_url"], "syntax": "attribute", "type": "string"}, {"aliases": ["authentication use auth object config"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:virtual_host:properties:authentication:use_auth_object_config", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["authentication", "use_auth_object_config"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/virtual_host/properties/authentication/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Authentication related information. This allows to configure the URL to redirect after the authentication Authentication Object Reference, configuration of cookie params etc.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
