---
page_title: "cloudflare.protected_endpoints.web_mobile_client.continue_web"
subcategory: ""
description: "cloudflare.protected_endpoints.web_mobile_client.continue_web for xcsh_protected_application."
xcsh_docs: {"aliases": [], "body_bytes": 2200, "body_sha256": "sha256:9fa9a6d25acf9d6adaa79da11635811ef4fc274e1db3d6be014f03a9740ae47b", "canonical_id": "xcsh-docs:data-sources:protected_application:properties:cloudflare:protected_endpoints:web_mobile_client:continue_web", "child_ids": ["xcsh-docs:data-sources:protected_application:properties:cloudflare:protected_endpoints:web_mobile_client:continue_web:add_header", "xcsh-docs:data-sources:protected_application:properties:cloudflare:protected_endpoints:web_mobile_client:continue_web:no_header"], "collection_id": "xcsh-docs:data-sources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:protected_application:properties:cloudflare:protected_endpoints:web_mobile_client:continue_web", "parent_id": "xcsh-docs:data-sources:protected_application:properties:cloudflare:protected_endpoints:web_mobile_client", "path": "docs/guides/data-sources--protected_application--properties--cloudflare--protected_endpoints--web_mobile_client--continue_web.md", "provider_name": "protected_application", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["cloudflare", "protected_endpoints", "web_mobile_client", "continue_web"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protected_application/properties/cloudflare/protected_endpoints/web_mobile_client/continue_web/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cloudflare.protected_endpoints.web_mobile_client.continue_web for xcsh_protected_application.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudflare.protected_endpoints.web_mobile_client.continue_web

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md)
- [Property reference](data-sources--protected_application--reference.md)
- [cloudflare](data-sources--protected_application--properties--cloudflare.md)
- [cloudflare.protected_endpoints](data-sources--protected_application--properties--cloudflare--protected_endpoints.md)
- [cloudflare.protected_endpoints.web_mobile_client](data-sources--protected_application--properties--cloudflare--protected_endpoints--web_mobile_client.md)
- cloudflare.protected_endpoints.web_mobile_client.continue_web

<a id="section"></a>

Type: `"single"`. Computed.

Select Continue Bot Mitigation Action. Continue mitigation action.

Upstream description:

Continue mitigation action.

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

## Direct properties

- [add_header](data-sources--protected_application--properties--cloudflare--protected_endpoints--web_mobile_client--continue_web--add_header.md): complete subsection reference.

- [no_header](data-sources--protected_application--properties--cloudflare--protected_endpoints--web_mobile_client--continue_web--no_header.md): complete subsection reference.

## Next pages

- [cloudflare.protected_endpoints.web_mobile_client.continue_web.add_header](data-sources--protected_application--properties--cloudflare--protected_endpoints--web_mobile_client--continue_web--add_header.md)
- [cloudflare.protected_endpoints.web_mobile_client.continue_web.no_header](data-sources--protected_application--properties--cloudflare--protected_endpoints--web_mobile_client--continue_web--no_header.md)
- [cloudflare.protected_endpoints.web_mobile_client](data-sources--protected_application--properties--cloudflare--protected_endpoints--web_mobile_client.md)
- [xcsh_protected_application](../data-sources/protected_application.md)
