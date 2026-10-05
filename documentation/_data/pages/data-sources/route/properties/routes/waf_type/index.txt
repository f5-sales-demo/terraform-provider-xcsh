---
page_title: "routes.waf_type"
subcategory: ""
description: "WAF instance will be pointing to an app_firewall object."
xcsh_docs: {"aliases": ["routes waf type"], "body_bytes": 2009, "body_sha256": "sha256:7fa8248106ab4f81c52282ae9c2498025b64b394582cf8c6dba8c1147eea47c2", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:route:properties:routes:waf_type:app_firewall", "xcsh-docs:data-sources:route:properties:routes:waf_type:disable_waf", "xcsh-docs:data-sources:route:properties:routes:waf_type:inherit_waf"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:route:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:route:properties:routes:waf_type", "parent_id": "xcsh-docs:data-sources:route:properties:routes", "path": "documentation/data-sources/route/properties/routes/waf_type/index.md", "product": "distributed-cloud", "provider_name": "route", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-1303120232330332-3020313221101000-0130321002101012-1313102131100000-3013131220122231-1130022213201220-0013210022302201-0010200133301120", "registry_path": "docs/guides/data-sources--route--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "waf_type"], "schema_version": 1, "sections": [{"aliases": ["routes waf type app firewall"], "anchor": "section", "description": "A list of references to the app_firewall configuration objects.", "document_id": "xcsh-docs:data-sources:route:properties:routes:waf_type:app_firewall", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["routes", "waf_type", "app_firewall"], "syntax": "attribute", "type": "object"}, {"aliases": ["routes waf type disable waf"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:route:properties:routes:waf_type:disable_waf", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "waf_type", "disable_waf"], "syntax": "attribute", "type": "object"}, {"aliases": ["routes waf type inherit waf"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:route:properties:routes:waf_type:inherit_waf", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "waf_type", "inherit_waf"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/route/properties/routes/waf_type/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "WAF instance will be pointing to an app_firewall object.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["routeCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.waf_type

Breadcrumbs:

- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/)
- routes.waf_type

<a id="section"></a>

Type: `"single"`. Computed.

WAF instance will be pointing to an app\_firewall object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ref_type": "[\"app_firewall\",\"disable_waf\",\"inherit_waf\"]"
}
```

## Direct properties

- [app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/waf_type/app_firewall/): complete subsection reference.

- [disable_waf](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/waf_type/disable_waf/): complete subsection reference.

- [inherit_waf](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/waf_type/inherit_waf/): complete subsection reference.

## Next pages

- [routes.waf_type.app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/waf_type/app_firewall/)
- [routes.waf_type.disable_waf](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/waf_type/disable_waf/)
- [routes.waf_type.inherit_waf](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/waf_type/inherit_waf/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/)
- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/)
