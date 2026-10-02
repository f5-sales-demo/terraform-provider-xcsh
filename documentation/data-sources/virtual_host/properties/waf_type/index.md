---
page_title: "waf_type"
subcategory: ""
description: "WAF instance will be pointing to an app_firewall object."
xcsh_docs: {"aliases": ["waf type"], "body_bytes": 1915, "body_sha256": "sha256:dc374560cec0e07ab73bfd8a69da3eb6e43d3075fbf9c4d34a6d585d2ebd8875", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:virtual_host:properties:waf_type:app_firewall", "xcsh-docs:data-sources:virtual_host:properties:waf_type:disable_waf", "xcsh-docs:data-sources:virtual_host:properties:waf_type:inherit_waf"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:virtual_host:properties:waf_type", "parent_id": "xcsh-docs:data-sources:virtual_host:reference", "path": "documentation/data-sources/virtual_host/properties/waf_type/index.md", "product": "distributed-cloud", "provider_name": "virtual_host", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3313303333213033-3223023131200111-1302021310020211-2020230220202121-2013001333301310-3033012222123323-3123120210220011-2132111220233001", "registry_path": "docs/guides/data-sources--virtual_host--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["waf_type"], "schema_version": 1, "sections": [{"aliases": ["app firewall"], "anchor": "section", "description": "A list of references to the app_firewall configuration objects.", "document_id": "xcsh-docs:data-sources:virtual_host:properties:waf_type:app_firewall", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["waf_type", "app_firewall"], "syntax": "attribute", "type": "object"}, {"aliases": ["disable waf"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:virtual_host:properties:waf_type:disable_waf", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["waf_type", "disable_waf"], "syntax": "attribute", "type": "object"}, {"aliases": ["inherit waf"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:virtual_host:properties:waf_type:inherit_waf", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["waf_type", "inherit_waf"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/virtual_host/properties/waf_type/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "WAF instance will be pointing to an app_firewall object.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# waf_type

Breadcrumbs:

- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/)
- waf_type

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

- [app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/waf_type/app_firewall/): complete subsection reference.

- [disable_waf](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/waf_type/disable_waf/): complete subsection reference.

- [inherit_waf](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/waf_type/inherit_waf/): complete subsection reference.

## Next pages

- [waf_type.app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/waf_type/app_firewall/)
- [waf_type.disable_waf](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/waf_type/disable_waf/)
- [waf_type.inherit_waf](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/waf_type/inherit_waf/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/)
- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/)
