---
page_title: "cloudfront.protected_endpoints.web_mobile_client.continue_mobile"
subcategory: ""
description: "cloudfront.protected_endpoints.web_mobile_client.continue_mobile for xcsh_protected_application."
xcsh_docs: {"aliases": [], "body_bytes": 2927, "body_sha256": "sha256:d506b72f749a42561849078e2ee3e5c8908dacc36289e0da989fcf34578225de", "child_ids": ["xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:web_mobile_client:continue_mobile:add_header", "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:web_mobile_client:continue_mobile:no_header"], "collection_id": "xcsh-docs:resources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:web_mobile_client:continue_mobile", "parent_id": "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:web_mobile_client", "path": "documentation/resources/protected_application/properties/cloudfront/protected_endpoints/web_mobile_client/continue_mobile/index.md", "provider_name": "protected_application", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["cloudfront", "protected_endpoints", "web_mobile_client", "continue_mobile"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protected_application/properties/cloudfront/protected_endpoints/web_mobile_client/continue_mobile/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cloudfront.protected_endpoints.web_mobile_client.continue_mobile for xcsh_protected_application.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# cloudfront.protected_endpoints.web_mobile_client.continue_mobile

Breadcrumbs:

- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/)
- [cloudfront](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/)
- [cloudfront.protected_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/)
- [cloudfront.protected_endpoints.web_mobile_client](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/web_mobile_client/)
- cloudfront.protected_endpoints.web_mobile_client.continue_mobile

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

- [add_header](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/web_mobile_client/continue_mobile/add_header/): complete subsection reference.

- [no_header](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/web_mobile_client/continue_mobile/no_header/): complete subsection reference.

## Next pages

- [cloudfront.protected_endpoints.web_mobile_client.continue_mobile.add_header](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/web_mobile_client/continue_mobile/add_header/)
- [cloudfront.protected_endpoints.web_mobile_client.continue_mobile.no_header](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/web_mobile_client/continue_mobile/no_header/)
- [cloudfront.protected_endpoints.web_mobile_client](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/web_mobile_client/)
- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/)
