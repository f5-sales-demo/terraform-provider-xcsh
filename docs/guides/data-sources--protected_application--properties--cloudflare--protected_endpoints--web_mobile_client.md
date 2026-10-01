---
page_title: "cloudflare.protected_endpoints.web_mobile_client"
subcategory: ""
description: "cloudflare.protected_endpoints.web_mobile_client for xcsh_protected_application."
xcsh_docs: {"aliases": [], "body_bytes": 2999, "body_sha256": "sha256:00b57c5fb962a0f10859111a7fdbf48d883db840fde9afb17483578513956d7c", "canonical_id": "xcsh-docs:data-sources:protected_application:properties:cloudflare:protected_endpoints:web_mobile_client", "child_ids": ["xcsh-docs:data-sources:protected_application:properties:cloudflare:protected_endpoints:web_mobile_client:block_mobile", "xcsh-docs:data-sources:protected_application:properties:cloudflare:protected_endpoints:web_mobile_client:block_web", "xcsh-docs:data-sources:protected_application:properties:cloudflare:protected_endpoints:web_mobile_client:continue_mobile", "xcsh-docs:data-sources:protected_application:properties:cloudflare:protected_endpoints:web_mobile_client:continue_web", "xcsh-docs:data-sources:protected_application:properties:cloudflare:protected_endpoints:web_mobile_client:redirect_web"], "collection_id": "xcsh-docs:data-sources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:protected_application:properties:cloudflare:protected_endpoints:web_mobile_client", "parent_id": "xcsh-docs:data-sources:protected_application:properties:cloudflare:protected_endpoints", "path": "docs/guides/data-sources--protected_application--properties--cloudflare--protected_endpoints--web_mobile_client.md", "provider_name": "protected_application", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["cloudflare", "protected_endpoints", "web_mobile_client"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protected_application/properties/cloudflare/protected_endpoints/web_mobile_client/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cloudflare.protected_endpoints.web_mobile_client for xcsh_protected_application.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudflare.protected_endpoints.web_mobile_client

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md)
- [Property reference](data-sources--protected_application--reference.md)
- [cloudflare](data-sources--protected_application--properties--cloudflare.md)
- [cloudflare.protected_endpoints](data-sources--protected_application--properties--cloudflare--protected_endpoints.md)
- cloudflare.protected_endpoints.web_mobile_client

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

- [block_mobile](data-sources--protected_application--properties--cloudflare--protected_endpoints--web_mobile_client--block_mobile.md): complete subsection reference.

- [block_web](data-sources--protected_application--properties--cloudflare--protected_endpoints--web_mobile_client--block_web.md): complete subsection reference.

- [continue_mobile](data-sources--protected_application--properties--cloudflare--protected_endpoints--web_mobile_client--continue_mobile.md): complete subsection reference.

- [continue_web](data-sources--protected_application--properties--cloudflare--protected_endpoints--web_mobile_client--continue_web.md): complete subsection reference.

- [redirect_web](data-sources--protected_application--properties--cloudflare--protected_endpoints--web_mobile_client--redirect_web.md): complete subsection reference.

## Next pages

- [cloudflare.protected_endpoints.web_mobile_client.block_mobile](data-sources--protected_application--properties--cloudflare--protected_endpoints--web_mobile_client--block_mobile.md)
- [cloudflare.protected_endpoints.web_mobile_client.block_web](data-sources--protected_application--properties--cloudflare--protected_endpoints--web_mobile_client--block_web.md)
- [cloudflare.protected_endpoints.web_mobile_client.continue_mobile](data-sources--protected_application--properties--cloudflare--protected_endpoints--web_mobile_client--continue_mobile.md)
- [cloudflare.protected_endpoints.web_mobile_client.continue_web](data-sources--protected_application--properties--cloudflare--protected_endpoints--web_mobile_client--continue_web.md)
- [cloudflare.protected_endpoints.web_mobile_client.redirect_web](data-sources--protected_application--properties--cloudflare--protected_endpoints--web_mobile_client--redirect_web.md)
- [cloudflare.protected_endpoints](data-sources--protected_application--properties--cloudflare--protected_endpoints.md)
- [xcsh_protected_application](../data-sources/protected_application.md)
