---
page_title: "ipsec.ike_parameters.dpd_keep_alive_timer"
subcategory: ""
description: "Configuration parameter for dpd keep alive timer."
xcsh_docs: {"aliases": ["ipsec ike parameters dpd keep alive timer"], "body_bytes": 1891, "body_sha256": "sha256:e24ecda3970f1303d6a1ab9cd1711e7664178a541cd423ce40192e98a5406be9", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:external_connector:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:external_connector:properties:ipsec:ike_parameters:dpd_keep_alive_timer", "parent_id": "xcsh-docs:data-sources:external_connector:properties:ipsec:ike_parameters", "path": "documentation/data-sources/external_connector/properties/ipsec/ike_parameters/dpd_keep_alive_timer/index.md", "product": "distributed-cloud", "provider_name": "external_connector", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-2122201322103313-2010210312210111-0332212213122101-1323230223100201-1302222022030111-3001121120121131-1221130012133231-2002100220100112", "registry_path": "docs/guides/data-sources--external_connector--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ipsec", "ike_parameters", "dpd_keep_alive_timer"], "schema_version": 1, "sections": [{"aliases": ["duration", "ipsec ike parameters dpd keep alive timer timeout"], "anchor": "schema-ipsec--ike_parameters--dpd_keep_alive_timer--timeout", "description": "Operation timeout duration", "document_id": "xcsh-docs:data-sources:external_connector:properties:ipsec:ike_parameters:dpd_keep_alive_timer", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ipsec", "ike_parameters", "dpd_keep_alive_timer", "timeout"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/external_connector/properties/ipsec/ike_parameters/dpd_keep_alive_timer/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Configuration parameter for dpd keep alive timer.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["external_connectorCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ipsec.ike_parameters.dpd_keep_alive_timer

Breadcrumbs:

- [xcsh_external_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/)
- [ipsec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/)
- [ipsec.ike_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ike_parameters/)
- ipsec.ike_parameters.dpd_keep_alive_timer

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for dpd keep alive timer.

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

<a id="schema-ipsec--ike_parameters--dpd_keep_alive_timer--timeout"></a>

### timeout property

Type: `"number"`. Computed.

Keepalive Timer. Operation timeout duration

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 5,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "5"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "5"
  }
}
```
