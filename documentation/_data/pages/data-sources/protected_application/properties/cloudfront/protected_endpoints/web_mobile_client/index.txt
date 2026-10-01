---
page_title: "cloudfront.protected_endpoints.web_mobile_client"
subcategory: ""
description: "cloudfront.protected_endpoints.web_mobile_client for xcsh_protected_application."
xcsh_docs: {"aliases": [], "body_bytes": 3785, "body_sha256": "sha256:43b3a88eb6b4077bf7e78b0a6ace883c274f70f9844f0f8b6fffd90a1c9206d5", "child_ids": ["xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:web_mobile_client:block_mobile", "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:web_mobile_client:block_web", "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:web_mobile_client:continue_mobile", "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:web_mobile_client:continue_web", "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:web_mobile_client:redirect_web"], "collection_id": "xcsh-docs:data-sources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:web_mobile_client", "parent_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints", "path": "documentation/data-sources/protected_application/properties/cloudfront/protected_endpoints/web_mobile_client/index.md", "provider_name": "protected_application", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "properties", "schema_path": ["cloudfront", "protected_endpoints", "web_mobile_client"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protected_application/properties/cloudfront/protected_endpoints/web_mobile_client/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cloudfront.protected_endpoints.web_mobile_client for xcsh_protected_application.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudfront.protected_endpoints.web_mobile_client

Breadcrumbs:

- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/)
- [cloudfront](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/)
- [cloudfront.protected_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/)
- cloudfront.protected_endpoints.web_mobile_client

<a id="section"></a>

Type: `"single"`. Computed.

Web and Mobile client configuration OPTIONS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-mobile_mitigation": "[\"block_mobile\",\"continue_mobile\"]",
  "x-ves-oneof-field-web_mitigation": "[\"block_web\",\"continue_web\",\"redirect_web\"]"
}
```

## Direct properties

- [block_mobile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/web_mobile_client/block_mobile/): complete subsection reference.

- [block_web](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/web_mobile_client/block_web/): complete subsection reference.

- [continue_mobile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/web_mobile_client/continue_mobile/): complete subsection reference.

- [continue_web](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/web_mobile_client/continue_web/): complete subsection reference.

- [redirect_web](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/web_mobile_client/redirect_web/): complete subsection reference.

## Next pages

- [cloudfront.protected_endpoints.web_mobile_client.block_mobile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/web_mobile_client/block_mobile/)
- [cloudfront.protected_endpoints.web_mobile_client.block_web](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/web_mobile_client/block_web/)
- [cloudfront.protected_endpoints.web_mobile_client.continue_mobile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/web_mobile_client/continue_mobile/)
- [cloudfront.protected_endpoints.web_mobile_client.continue_web](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/web_mobile_client/continue_web/)
- [cloudfront.protected_endpoints.web_mobile_client.redirect_web](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/web_mobile_client/redirect_web/)
- [cloudfront.protected_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/)
- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/)
