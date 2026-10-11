---
page_title: "routes.service_policy"
subcategory: ""
description: "ServicePolicy configuration details at route level."
xcsh_docs: {"aliases": ["routes service policy"], "body_bytes": 1049, "body_sha256": "sha256:3a9c46909c4fb95290c77bca198268f1fa7304c8dc084a8f692acdbdf3975e99", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:route:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:route:properties:routes:service_policy", "parent_id": "xcsh-docs:data-sources:route:properties:routes", "path": "documentation/data-sources/route/properties/routes/service_policy/index.md", "product": "distributed-cloud", "provider_name": "route", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "data-sources", "registry_anchor": "canonical-0122003031122022-2312121320313011-3232203000312111-2300213223121003-3333300031112101-3031102112131312-1230313311203332-3230301311131320", "registry_path": "docs/guides/data-sources--route--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "service_policy"], "schema_version": 1, "sections": [{"aliases": ["routes service policy disable spec"], "anchor": "schema-routes--service_policy--disable_spec", "description": "Exclusive with disable service policy at route level, if it is configured at virtual-host level.", "document_id": "xcsh-docs:data-sources:route:properties:routes:service_policy", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "service_policy", "disable_spec"], "syntax": "attribute", "type": "bool"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/route/properties/routes/service_policy/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "ServicePolicy configuration details at route level.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["routeCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
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
