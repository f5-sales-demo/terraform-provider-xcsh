---
page_title: "business_logic_markup_setting.discovered_api_settings"
subcategory: ""
description: "Configure Discovered API Settings."
xcsh_docs: {"aliases": ["business logic markup setting discovered api settings"], "body_bytes": 2199, "body_sha256": "sha256:64c2320ac6282589326ad04634fd6433e45e9a11b9d9ff7f1c82cc3766cf899a", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:app_type:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:app_type:properties:business_logic_markup_setting:discovered_api_settings", "parent_id": "xcsh-docs:data-sources:app_type:properties:business_logic_markup_setting", "path": "documentation/data-sources/app_type/properties/business_logic_markup_setting/discovered_api_settings/index.md", "product": "distributed-cloud", "provider_name": "app_type", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-2032222223210300-2133102122221101-0302332213001120-0000102200020222-2012301022111102-0013310120101020-2013121102320313-0300133021102221", "registry_path": "docs/guides/data-sources--app_type--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["business_logic_markup_setting", "discovered_api_settings"], "schema_version": 1, "sections": [{"aliases": ["business logic markup setting discovered api settings purge duration for inactive discovered apis"], "anchor": "schema-business_logic_markup_setting--discovered_api_settings--purge_duration_for_inactive_discovered_apis", "description": "Inactive discovered API will be deleted after configured duration.", "document_id": "xcsh-docs:data-sources:app_type:properties:business_logic_markup_setting:discovered_api_settings", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["business_logic_markup_setting", "discovered_api_settings", "purge_duration_for_inactive_discovered_apis"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_type/properties/business_logic_markup_setting/discovered_api_settings/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Configure Discovered API Settings.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["app_typeCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# business_logic_markup_setting.discovered_api_settings

Breadcrumbs:

- [xcsh_app_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_type/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_type/properties/)
- [business_logic_markup_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_type/properties/business_logic_markup_setting/)
- business_logic_markup_setting.discovered_api_settings

<a id="section"></a>

Type: `"single"`. Computed.

Discovered API Settings. Configure Discovered API Settings.

Upstream description:

Configure Discovered API Settings.

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

<a id="schema-business_logic_markup_setting--discovered_api_settings--purge_duration_for_inactive_discovered_apis"></a>

### purge_duration_for_inactive_discovered_apis property

Type: `"number"`. Computed.

Inactive discovered API will be deleted after configured duration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 7,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "7"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "7"
  }
}
```

## Next pages

- [business_logic_markup_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_type/properties/business_logic_markup_setting/)
- [xcsh_app_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_type/)
