---
page_title: "cloudflare.protected_endpoints.web_mobile_client.continue_mobile"
subcategory: ""
description: "cloudflare.protected_endpoints.web_mobile_client.continue_mobile for xcsh_protected_application."
xcsh_docs: {"aliases": [], "body_bytes": 3026, "body_sha256": "sha256:69bb89116081a9fa2109980f3a967ba9eab83e006e37a414842fa9c09ddceb1a", "child_ids": ["xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:web_mobile_client:continue_mobile:add_header", "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:web_mobile_client:continue_mobile:no_header"], "collection_id": "xcsh-docs:resources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:web_mobile_client:continue_mobile", "parent_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:web_mobile_client", "path": "documentation/resources/protected_application/properties/cloudflare/protected_endpoints/web_mobile_client/continue_mobile/index.md", "provider_name": "protected_application", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["cloudflare", "protected_endpoints", "web_mobile_client", "continue_mobile"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protected_application/properties/cloudflare/protected_endpoints/web_mobile_client/continue_mobile/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cloudflare.protected_endpoints.web_mobile_client.continue_mobile for xcsh_protected_application.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudflare.protected_endpoints.web_mobile_client.continue_mobile

Breadcrumbs:

- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/)
- [cloudflare](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/)
- [cloudflare.protected_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/protected_endpoints/)
- [cloudflare.protected_endpoints.web_mobile_client](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/protected_endpoints/web_mobile_client/)
- cloudflare.protected_endpoints.web_mobile_client.continue_mobile

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Select Continue Bot Mitigation Action. Continue mitigation action.

Upstream description:

Continue mitigation action.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("add_header",
    "no_header")}
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
  "x-ves-oneof-field-add_header_choice": "[\"add_header\",\"no_header\"]"
}
```

Terraform syntax:

```terraform
continue_mobile {
  # Configure direct properties listed below.
}
```

## Direct properties

- [add_header](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/protected_endpoints/web_mobile_client/continue_mobile/add_header/): complete subsection reference.

- [no_header](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/protected_endpoints/web_mobile_client/continue_mobile/no_header/): complete subsection reference.

## Next pages

- [cloudflare.protected_endpoints.web_mobile_client.continue_mobile.add_header](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/protected_endpoints/web_mobile_client/continue_mobile/add_header/)
- [cloudflare.protected_endpoints.web_mobile_client.continue_mobile.no_header](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/protected_endpoints/web_mobile_client/continue_mobile/no_header/)
- [cloudflare.protected_endpoints.web_mobile_client](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/protected_endpoints/web_mobile_client/)
- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/)
