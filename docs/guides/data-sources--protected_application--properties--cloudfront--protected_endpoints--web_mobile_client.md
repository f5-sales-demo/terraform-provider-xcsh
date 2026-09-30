---
page_title: "cloudfront.protected_endpoints.web_mobile_client"
subcategory: ""
description: "cloudfront.protected_endpoints.web_mobile_client for xcsh_protected_application."
xcsh_docs: {"aliases": [], "body_bytes": 2900, "body_sha256": "sha256:1e9ba5ed960432bf98fdb1b12391b5995e6ed8410646e0f52014b764548b2f4a", "canonical_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:web_mobile_client", "child_ids": ["xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:web_mobile_client:block_mobile", "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:web_mobile_client:block_web", "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:web_mobile_client:continue_mobile", "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:web_mobile_client:continue_web", "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:web_mobile_client:redirect_web"], "collection_id": "xcsh-docs:data-sources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:web_mobile_client", "parent_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints", "path": "docs/guides/data-sources--protected_application--properties--cloudfront--protected_endpoints--web_mobile_client.md", "provider_name": "protected_application", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["cloudfront", "protected_endpoints", "web_mobile_client"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protected_application/properties/cloudfront/protected_endpoints/web_mobile_client/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cloudfront.protected_endpoints.web_mobile_client for xcsh_protected_application.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# cloudfront.protected_endpoints.web_mobile_client

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md)
- [Property reference](data-sources--protected_application--reference.md)
- [cloudfront](data-sources--protected_application--properties--cloudfront.md)
- [cloudfront.protected_endpoints](data-sources--protected_application--properties--cloudfront--protected_endpoints.md)
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

- [block_mobile](data-sources--protected_application--properties--cloudfront--protected_endpoints--web_mobile_client--block_mobile.md): complete subsection reference.

- [block_web](data-sources--protected_application--properties--cloudfront--protected_endpoints--web_mobile_client--block_web.md): complete subsection reference.

- [continue_mobile](data-sources--protected_application--properties--cloudfront--protected_endpoints--web_mobile_client--continue_mobile.md): complete subsection reference.

- [continue_web](data-sources--protected_application--properties--cloudfront--protected_endpoints--web_mobile_client--continue_web.md): complete subsection reference.

- [redirect_web](data-sources--protected_application--properties--cloudfront--protected_endpoints--web_mobile_client--redirect_web.md): complete subsection reference.

## Next pages

- [cloudfront.protected_endpoints.web_mobile_client.block_mobile](data-sources--protected_application--properties--cloudfront--protected_endpoints--web_mobile_client--block_mobile.md)
- [cloudfront.protected_endpoints.web_mobile_client.block_web](data-sources--protected_application--properties--cloudfront--protected_endpoints--web_mobile_client--block_web.md)
- [cloudfront.protected_endpoints.web_mobile_client.continue_mobile](data-sources--protected_application--properties--cloudfront--protected_endpoints--web_mobile_client--continue_mobile.md)
- [cloudfront.protected_endpoints.web_mobile_client.continue_web](data-sources--protected_application--properties--cloudfront--protected_endpoints--web_mobile_client--continue_web.md)
- [cloudfront.protected_endpoints.web_mobile_client.redirect_web](data-sources--protected_application--properties--cloudfront--protected_endpoints--web_mobile_client--redirect_web.md)
- [cloudfront.protected_endpoints](data-sources--protected_application--properties--cloudfront--protected_endpoints.md)
- [xcsh_protected_application](../data-sources/protected_application.md)
