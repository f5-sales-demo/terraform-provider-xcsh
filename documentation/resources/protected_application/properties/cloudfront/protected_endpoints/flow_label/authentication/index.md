---
page_title: "cloudfront.protected_endpoints.flow_label.authentication"
subcategory: ""
description: "cloudfront.protected_endpoints.flow_label.authentication for xcsh_protected_application."
xcsh_docs: {"aliases": [], "body_bytes": 4912, "body_sha256": "sha256:70fea27759df296477e7be0c6675f7396d72ee28e2109105fa867416f0dec64a", "child_ids": ["xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label:authentication:login", "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label:authentication:login_mfa", "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label:authentication:login_partner", "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label:authentication:logout", "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label:authentication:token_refresh"], "collection_id": "xcsh-docs:resources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label:authentication", "parent_id": "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label", "path": "documentation/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/authentication/index.md", "provider_name": "protected_application", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["cloudfront", "protected_endpoints", "flow_label", "authentication"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/authentication/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cloudfront.protected_endpoints.flow_label.authentication for xcsh_protected_application.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudfront.protected_endpoints.flow_label.authentication

Breadcrumbs:

- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/)
- [cloudfront](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/)
- [cloudfront.protected_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/)
- [cloudfront.protected_endpoints.flow_label](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/)
- cloudfront.protected_endpoints.flow_label.authentication

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Flow Label Authentication Category.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("login",
    "login_mfa"),
  validators.ConflictingObjectAttributes("login",
    "login_partner"),
  validators.ConflictingObjectAttributes("login",
    "logout"),
  validators.ConflictingObjectAttributes("login",
    "token_refresh"),
  validators.ConflictingObjectAttributes("login_mfa",
    "login_partner"),
  validators.ConflictingObjectAttributes("login_mfa",
    "logout"),
  validators.ConflictingObjectAttributes("login_mfa",
    "token_refresh"),
  validators.ConflictingObjectAttributes("login_partner",
    "logout"),
  validators.ConflictingObjectAttributes("login_partner",
    "token_refresh"),
  validators.ConflictingObjectAttributes("logout",
    "token_refresh")}
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
  "x-ves-oneof-field-label_choice": "[\"login\",\"login_mfa\",\"login_partner\",\"logout\",\"token_refresh\"]"
}
```

Terraform syntax:

```terraform
authentication {
  # Configure direct properties listed below.
}
```

## Direct properties

- [login](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/authentication/login/): complete subsection reference.

- [login_mfa](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/authentication/login_mfa/): complete subsection reference.

- [login_partner](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/authentication/login_partner/): complete subsection reference.

- [logout](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/authentication/logout/): complete subsection reference.

- [token_refresh](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/authentication/token_refresh/): complete subsection reference.

## Next pages

- [cloudfront.protected_endpoints.flow_label.authentication.login](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/authentication/login/)
- [cloudfront.protected_endpoints.flow_label.authentication.login_mfa](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/authentication/login_mfa/)
- [cloudfront.protected_endpoints.flow_label.authentication.login_partner](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/authentication/login_partner/)
- [cloudfront.protected_endpoints.flow_label.authentication.logout](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/authentication/logout/)
- [cloudfront.protected_endpoints.flow_label.authentication.token_refresh](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/authentication/token_refresh/)
- [cloudfront.protected_endpoints.flow_label](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/)
- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/)
