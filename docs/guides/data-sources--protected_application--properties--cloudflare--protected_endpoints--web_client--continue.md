---
page_title: "cloudflare.protected_endpoints.web_client.continue"
subcategory: ""
description: "cloudflare.protected_endpoints.web_client.continue for xcsh_protected_application."
xcsh_docs: {"aliases": [], "body_bytes": 1985, "body_sha256": "sha256:5c610f97549d3415672ed0a1085f44e92d0c4d5e53e3a83ede6272bd04b2795e", "canonical_id": "xcsh-docs:data-sources:protected_application:properties:cloudflare:protected_endpoints:web_client:continue", "child_ids": ["xcsh-docs:data-sources:protected_application:properties:cloudflare:protected_endpoints:web_client:continue:add_header", "xcsh-docs:data-sources:protected_application:properties:cloudflare:protected_endpoints:web_client:continue:no_header"], "collection_id": "xcsh-docs:data-sources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:protected_application:properties:cloudflare:protected_endpoints:web_client:continue", "parent_id": "xcsh-docs:data-sources:protected_application:properties:cloudflare:protected_endpoints:web_client", "path": "docs/guides/data-sources--protected_application--properties--cloudflare--protected_endpoints--web_client--continue.md", "provider_name": "protected_application", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["cloudflare", "protected_endpoints", "web_client", "continue"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protected_application/properties/cloudflare/protected_endpoints/web_client/continue/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cloudflare.protected_endpoints.web_client.continue for xcsh_protected_application.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# cloudflare.protected_endpoints.web_client.continue

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md)
- [Property reference](data-sources--protected_application--reference.md)
- [cloudflare](data-sources--protected_application--properties--cloudflare.md)
- [cloudflare.protected_endpoints](data-sources--protected_application--properties--cloudflare--protected_endpoints.md)
- [cloudflare.protected_endpoints.web_client](data-sources--protected_application--properties--cloudflare--protected_endpoints--web_client.md)
- cloudflare.protected_endpoints.web_client.continue

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

- [add_header](data-sources--protected_application--properties--cloudflare--protected_endpoints--web_client--continue--add_header.md): complete subsection reference.

- [no_header](data-sources--protected_application--properties--cloudflare--protected_endpoints--web_client--continue--no_header.md): complete subsection reference.

## Next pages

- [cloudflare.protected_endpoints.web_client.continue.add_header](data-sources--protected_application--properties--cloudflare--protected_endpoints--web_client--continue--add_header.md)
- [cloudflare.protected_endpoints.web_client.continue.no_header](data-sources--protected_application--properties--cloudflare--protected_endpoints--web_client--continue--no_header.md)
- [cloudflare.protected_endpoints.web_client](data-sources--protected_application--properties--cloudflare--protected_endpoints--web_client.md)
- [xcsh_protected_application](../data-sources/protected_application.md)
