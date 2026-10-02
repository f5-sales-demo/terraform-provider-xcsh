---
page_title: "waf_type.app_firewall"
subcategory: ""
description: "A list of references to the app_firewall configuration objects."
xcsh_docs: {"aliases": ["waf type app firewall"], "body_bytes": 1492, "body_sha256": "sha256:f995bfc9e5b126e8462f3accbb31bbb9fc4585b8ced9099a8dd99c2a4a7d2d23", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:virtual_host:properties:waf_type:app_firewall:app_firewall"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:virtual_host:properties:waf_type:app_firewall", "parent_id": "xcsh-docs:data-sources:virtual_host:properties:waf_type", "path": "documentation/data-sources/virtual_host/properties/waf_type/app_firewall/index.md", "product": "distributed-cloud", "provider_name": "virtual_host", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1333300322323113-1232131101000031-2233210011203232-3322023300230033-1233133023310221-1233001022120331-2003112011233330-2313020200000013", "registry_path": "docs/guides/data-sources--virtual_host--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["waf_type", "app_firewall"], "schema_version": 1, "sections": [{"aliases": ["app firewall"], "anchor": "section", "description": "References to an Application Firewall configuration object.", "document_id": "xcsh-docs:data-sources:virtual_host:properties:waf_type:app_firewall:app_firewall", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["waf_type", "app_firewall", "app_firewall"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/virtual_host/properties/waf_type/app_firewall/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "A list of references to the app_firewall configuration objects.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# waf_type.app_firewall

Breadcrumbs:

- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/)
- [waf_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/waf_type/)
- waf_type.app_firewall

<a id="section"></a>

Type: `"single"`. Computed.

List of references to the app\_firewall configuration objects.

Upstream description:

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

- [app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/waf_type/app_firewall/app_firewall/): complete subsection reference.

## Next pages

- [waf_type.app_firewall.app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/waf_type/app_firewall/app_firewall/)
- [waf_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/waf_type/)
- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/)
