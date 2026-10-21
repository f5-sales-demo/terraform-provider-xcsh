---
page_title: "routes.custom"
subcategory: "Monitoring"
description: "A set of matchers an alert has to fulfill to match the route."
xcsh_docs: {"aliases": ["routes custom"], "body_bytes": 1450, "body_sha256": "sha256:884c6432493f62b6cc8c24d52ee6595a8775414c7a9e5a7532648d03c607ec83", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:data-sources:alert_policy:properties:routes:custom:alertlabel", "xcsh-docs:data-sources:alert_policy:properties:routes:custom:alertname", "xcsh-docs:data-sources:alert_policy:properties:routes:custom:group", "xcsh-docs:data-sources:alert_policy:properties:routes:custom:severity"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:alert_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:alert_policy:properties:routes:custom", "parent_id": "xcsh-docs:data-sources:alert_policy:properties:routes", "path": "documentation/data-sources/alert_policy/properties/routes/custom/index.md", "product": "distributed-cloud", "provider_name": "alert_policy", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-2111032222302223-0233011322210313-1123202120313233-2330130321123330-0231130212120033-2213022223222302-1301321103330130-3002100333101000", "registry_path": "docs/guides/data-sources--alert_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "custom"], "schema_version": 1, "sections": [{"aliases": ["routes custom alertlabel"], "anchor": "section", "description": "AlertLabel to configure the alert policy rule.", "document_id": "xcsh-docs:data-sources:alert_policy:properties:routes:custom:alertlabel", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["routes", "custom", "alertlabel"], "syntax": "attribute", "type": "object"}, {"aliases": ["routes custom alertname"], "anchor": "section", "description": "Label Matcher.", "document_id": "xcsh-docs:data-sources:alert_policy:properties:routes:custom:alertname", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["routes", "custom", "alertname"], "syntax": "attribute", "type": "object"}, {"aliases": ["routes custom group"], "anchor": "section", "description": "Label Matcher.", "document_id": "xcsh-docs:data-sources:alert_policy:properties:routes:custom:group", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["routes", "custom", "group"], "syntax": "attribute", "type": "object"}, {"aliases": ["routes custom severity"], "anchor": "section", "description": "Label Matcher.", "document_id": "xcsh-docs:data-sources:alert_policy:properties:routes:custom:severity", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["routes", "custom", "severity"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/alert_policy/properties/routes/custom/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "A set of matchers an alert has to fulfill to match the route.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["alert_policyCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.custom

Breadcrumbs:

- [xcsh_alert_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/routes/)
- routes.custom

<a id="section"></a>

Type: `"single"`. Computed.

A set of matchers an alert has to fulfill to match the route.

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

- [alertlabel](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/routes/custom/alertlabel/): complete subsection reference.

- [alertname](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/routes/custom/alertname/): complete subsection reference.

- [group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/routes/custom/group/): complete subsection reference.

- [severity](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/routes/custom/severity/): complete subsection reference.
