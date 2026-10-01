---
page_title: "cloudflare.protected_endpoints.web_client"
subcategory: ""
description: "cloudflare.protected_endpoints.web_client for xcsh_protected_application."
xcsh_docs: {"aliases": [], "body_bytes": 2118, "body_sha256": "sha256:2a842e2c4b07eb3b742524705f9f92d14fa7f55765431fa2e09cca7c07b13240", "canonical_id": "xcsh-docs:data-sources:protected_application:properties:cloudflare:protected_endpoints:web_client", "child_ids": ["xcsh-docs:data-sources:protected_application:properties:cloudflare:protected_endpoints:web_client:block", "xcsh-docs:data-sources:protected_application:properties:cloudflare:protected_endpoints:web_client:continue", "xcsh-docs:data-sources:protected_application:properties:cloudflare:protected_endpoints:web_client:redirect"], "collection_id": "xcsh-docs:data-sources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:protected_application:properties:cloudflare:protected_endpoints:web_client", "parent_id": "xcsh-docs:data-sources:protected_application:properties:cloudflare:protected_endpoints", "path": "docs/guides/data-sources--protected_application--properties--cloudflare--protected_endpoints--web_client.md", "provider_name": "protected_application", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["cloudflare", "protected_endpoints", "web_client"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protected_application/properties/cloudflare/protected_endpoints/web_client/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cloudflare.protected_endpoints.web_client for xcsh_protected_application.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudflare.protected_endpoints.web_client

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md)
- [Property reference](data-sources--protected_application--reference.md)
- [cloudflare](data-sources--protected_application--properties--cloudflare.md)
- [cloudflare.protected_endpoints](data-sources--protected_application--properties--cloudflare--protected_endpoints.md)
- cloudflare.protected_endpoints.web_client

<a id="section"></a>

Type: `"single"`. Computed.

Web Client. Web client configuration OPTIONS.

Upstream description:

Web client configuration OPTIONS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-mitigation": "[\"block\",\"continue\",\"redirect\"]"
}
```

## Direct properties

- [block](data-sources--protected_application--properties--cloudflare--protected_endpoints--web_client--block.md): complete subsection reference.

- [continue](data-sources--protected_application--properties--cloudflare--protected_endpoints--web_client--continue.md): complete subsection reference.

- [redirect](data-sources--protected_application--properties--cloudflare--protected_endpoints--web_client--redirect.md): complete subsection reference.

## Next pages

- [cloudflare.protected_endpoints.web_client.block](data-sources--protected_application--properties--cloudflare--protected_endpoints--web_client--block.md)
- [cloudflare.protected_endpoints.web_client.continue](data-sources--protected_application--properties--cloudflare--protected_endpoints--web_client--continue.md)
- [cloudflare.protected_endpoints.web_client.redirect](data-sources--protected_application--properties--cloudflare--protected_endpoints--web_client--redirect.md)
- [cloudflare.protected_endpoints](data-sources--protected_application--properties--cloudflare--protected_endpoints.md)
- [xcsh_protected_application](../data-sources/protected_application.md)
