---
page_title: "routes.custom"
subcategory: "Monitoring"
description: "A set of matchers an alert has to fulfill to match the route."
xcsh_docs: {"aliases": ["routes custom"], "body_bytes": 2349, "body_sha256": "sha256:3612878d264111fd1acc1971e92ae3d3f14b4d55ef666c149f62a382eca3b410", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:data-sources:alert_policy:properties:routes:custom:alertlabel", "xcsh-docs:data-sources:alert_policy:properties:routes:custom:alertname", "xcsh-docs:data-sources:alert_policy:properties:routes:custom:group", "xcsh-docs:data-sources:alert_policy:properties:routes:custom:severity"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:alert_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:alert_policy:properties:routes:custom", "parent_id": "xcsh-docs:data-sources:alert_policy:properties:routes", "path": "documentation/data-sources/alert_policy/properties/routes/custom/index.md", "product": "distributed-cloud", "provider_name": "alert_policy", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "data-sources", "registry_anchor": "canonical-2111032222302223-0233011322210313-1123202120313233-2330130321123330-0231130212120033-2213022223222302-1301321103330130-3002100333101000", "registry_path": "docs/guides/data-sources--alert_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "custom"], "schema_version": 1, "sections": [{"aliases": ["alertlabel"], "anchor": "section", "description": "AlertLabel to configure the alert policy rule.", "document_id": "xcsh-docs:data-sources:alert_policy:properties:routes:custom:alertlabel", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["routes", "custom", "alertlabel"], "syntax": "attribute", "type": "object"}, {"aliases": ["alertname"], "anchor": "section", "description": "Label Matcher.", "document_id": "xcsh-docs:data-sources:alert_policy:properties:routes:custom:alertname", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["routes", "custom", "alertname"], "syntax": "attribute", "type": "object"}, {"aliases": ["group"], "anchor": "section", "description": "Label Matcher.", "document_id": "xcsh-docs:data-sources:alert_policy:properties:routes:custom:group", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["routes", "custom", "group"], "syntax": "attribute", "type": "object"}, {"aliases": ["severity"], "anchor": "section", "description": "Label Matcher.", "document_id": "xcsh-docs:data-sources:alert_policy:properties:routes:custom:severity", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["routes", "custom", "severity"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/alert_policy/properties/routes/custom/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "A set of matchers an alert has to fulfill to match the route.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["alert_policyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

Set of matchers an alert has to fulfill to match the route.

Upstream description:

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

## Next pages

- [routes.custom.alertlabel](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/routes/custom/alertlabel/)
- [routes.custom.alertname](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/routes/custom/alertname/)
- [routes.custom.group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/routes/custom/group/)
- [routes.custom.severity](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/routes/custom/severity/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/routes/)
- [xcsh_alert_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/)
