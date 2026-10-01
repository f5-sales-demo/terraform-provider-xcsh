---
page_title: "authentication"
subcategory: ""
description: "authentication for xcsh_virtual_host."
xcsh_docs: {"aliases": [], "body_bytes": 4515, "body_sha256": "sha256:c268fccd8adee1bf5e49cfc88611012609e4e0fa61f9f845ff0bfb95ef0ce74c", "canonical_id": "xcsh-docs:resources:virtual_host:properties:authentication", "child_ids": ["xcsh-docs:resources:virtual_host:properties:authentication:auth_config", "xcsh-docs:resources:virtual_host:properties:authentication:cookie_params", "xcsh-docs:resources:virtual_host:properties:authentication:redirect_dynamic", "xcsh-docs:resources:virtual_host:properties:authentication:use_auth_object_config"], "collection_id": "xcsh-docs:resources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_host:properties:authentication", "parent_id": "xcsh-docs:resources:virtual_host:reference", "path": "docs/guides/resources--virtual_host--properties--authentication.md", "provider_name": "virtual_host", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["authentication"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_host/properties/authentication/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "authentication for xcsh_virtual_host.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# authentication

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md)
- [Property reference](resources--virtual_host--reference.md)
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

- [authentication](resources--virtual_host--properties--authentication.md#section)
- [no_authentication](resources--virtual_host--properties--no_authentication.md#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
authentication {
  # Configure direct properties listed below.
}
```

## Direct properties

- [auth_config](resources--virtual_host--properties--authentication--auth_config.md): complete subsection reference.

- [cookie_params](resources--virtual_host--properties--authentication--cookie_params.md): complete subsection reference.

- [redirect_dynamic](resources--virtual_host--properties--authentication--redirect_dynamic.md): complete subsection reference.

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

- [use_auth_object_config](resources--virtual_host--properties--authentication--use_auth_object_config.md): complete subsection reference.

## Next pages

- [authentication.auth_config](resources--virtual_host--properties--authentication--auth_config.md)
- [authentication.cookie_params](resources--virtual_host--properties--authentication--cookie_params.md)
- [authentication.redirect_dynamic](resources--virtual_host--properties--authentication--redirect_dynamic.md)
- [authentication.use_auth_object_config](resources--virtual_host--properties--authentication--use_auth_object_config.md)
- [Property reference](resources--virtual_host--reference.md)
- [xcsh_virtual_host](../resources/virtual_host.md)
