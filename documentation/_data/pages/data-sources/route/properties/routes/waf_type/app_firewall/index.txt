---
page_title: "routes.waf_type.app_firewall"
subcategory: ""
description: "A list of references to the app_firewall configuration objects."
xcsh_docs: {"aliases": ["routes waf type app firewall"], "body_bytes": 1103, "body_sha256": "sha256:cd4afe14382405b0874f56f36ef1ac4650eabe655eb873eda26d69f689bc449b", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:route:properties:routes:waf_type:app_firewall:app_firewall"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:route:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:route:properties:routes:waf_type:app_firewall", "parent_id": "xcsh-docs:data-sources:route:properties:routes:waf_type", "path": "documentation/data-sources/route/properties/routes/waf_type/app_firewall/index.md", "product": "distributed-cloud", "provider_name": "route", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-2020000300332202-2313022220012333-0031313002211002-1203003321210130-2133210011020131-2022331332130112-3122033002213300-3020310233103332", "registry_path": "docs/guides/data-sources--route--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "waf_type", "app_firewall"], "schema_version": 1, "sections": [{"aliases": ["routes waf type app firewall app firewall"], "anchor": "section", "description": "References to an Application Firewall configuration object.", "document_id": "xcsh-docs:data-sources:route:properties:routes:waf_type:app_firewall:app_firewall", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["routes", "waf_type", "app_firewall", "app_firewall"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/route/properties/routes/waf_type/app_firewall/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "A list of references to the app_firewall configuration objects.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["routeCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.waf_type.app_firewall

Breadcrumbs:

- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/)
- [routes.waf_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/waf_type/)
- routes.waf_type.app_firewall

<a id="section"></a>

Type: `"single"`. Computed.

A list of references to the app\_firewall configuration objects.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

## Direct properties

- [app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/waf_type/app_firewall/app_firewall/): complete subsection reference.
