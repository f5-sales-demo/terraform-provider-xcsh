---
page_title: "routes.service_policy"
subcategory: ""
description: "ServicePolicy configuration details at route level."
xcsh_docs: {"aliases": ["routes service policy"], "body_bytes": 1263, "body_sha256": "sha256:b0c4b4d80645cdeb6b39d29c9da61f2026def05ae5a20a5f56157a069b3b5963", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:route:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:route:properties:routes:service_policy", "parent_id": "xcsh-docs:data-sources:route:properties:routes", "path": "documentation/data-sources/route/properties/routes/service_policy/index.md", "product": "distributed-cloud", "provider_name": "route", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0122003031122022-2312121320313011-3232203000312111-2300213223121003-3333300031112101-3031102112131312-1230313311203332-3230301311131320", "registry_path": "docs/guides/data-sources--route--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "service_policy"], "schema_version": 1, "sections": [{"aliases": ["disable spec"], "anchor": "schema-routes--service_policy--disable_spec", "description": "Exclusive with disable service policy at route level, if it is configured at virtual-host level.", "document_id": "xcsh-docs:data-sources:route:properties:routes:service_policy", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "service_policy", "disable_spec"], "syntax": "attribute", "type": "bool"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/route/properties/routes/service_policy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ServicePolicy configuration details at route level.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["routeCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.service_policy

Breadcrumbs:

- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/)
- routes.service_policy

<a id="section"></a>

Type: `"single"`. Computed.

ServicePolicy configuration details at route level.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-service_policy_choice": "[\"disable\"]"
}
```

## Direct properties

<a id="schema-routes--service_policy--disable_spec"></a>

### disable_spec property

Type: `"bool"`. Computed.

Exclusive with \[\] disable service policy at route level, if it is configured at virtual-host
level.

## Next pages

- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/)
- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/)
