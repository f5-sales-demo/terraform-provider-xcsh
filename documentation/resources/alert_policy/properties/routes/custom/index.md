---
page_title: "routes.custom"
subcategory: "Monitoring"
description: "A set of matchers an alert has to fulfill to match the route."
xcsh_docs: {"aliases": ["routes custom"], "body_bytes": 2425, "body_sha256": "sha256:f4a9e5afa15b6430bc9e517bd123553019410803364ef15f861f85d9b0817436", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:resources:alert_policy:properties:routes:custom:alertlabel", "xcsh-docs:resources:alert_policy:properties:routes:custom:alertname", "xcsh-docs:resources:alert_policy:properties:routes:custom:group", "xcsh-docs:resources:alert_policy:properties:routes:custom:severity"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:alert_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:alert_policy:properties:routes:custom", "parent_id": "xcsh-docs:resources:alert_policy:properties:routes", "path": "documentation/resources/alert_policy/properties/routes/custom/index.md", "product": "distributed-cloud", "provider_name": "alert_policy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3230201200300110-3010211222100103-3201230231102031-1213011103213132-0310132223101002-0010233123332210-2030021202212202-3031020303123101", "registry_path": "docs/guides/resources--alert_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "custom"], "schema_version": 1, "sections": [{"aliases": ["routes custom alertlabel"], "anchor": "section", "description": "AlertLabel to configure the alert policy rule.", "document_id": "xcsh-docs:resources:alert_policy:properties:routes:custom:alertlabel", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["routes", "custom", "alertlabel"], "syntax": "block", "type": "object"}, {"aliases": ["routes custom alertname"], "anchor": "section", "description": "Label Matcher.", "document_id": "xcsh-docs:resources:alert_policy:properties:routes:custom:alertname", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-routes--custom--alertname--exact_match", "enforcement": "provider-schema", "group": "routes.custom.alertname:ConflictingObjectAttributes:exact_match,regex_match", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:routes:custom:alertname", "type": "conflicts"}, {"anchor": "schema-routes--custom--alertname--regex_match", "enforcement": "provider-schema", "group": "routes.custom.alertname:ConflictingObjectAttributes:exact_match,regex_match", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:routes:custom:alertname", "type": "conflicts"}], "schema_path": ["routes", "custom", "alertname"], "syntax": "block", "type": "object"}, {"aliases": ["routes custom group"], "anchor": "section", "description": "Label Matcher.", "document_id": "xcsh-docs:resources:alert_policy:properties:routes:custom:group", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-routes--custom--group--exact_match", "enforcement": "provider-schema", "group": "routes.custom.group:ConflictingObjectAttributes:exact_match,regex_match", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:routes:custom:group", "type": "conflicts"}, {"anchor": "schema-routes--custom--group--regex_match", "enforcement": "provider-schema", "group": "routes.custom.group:ConflictingObjectAttributes:exact_match,regex_match", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:routes:custom:group", "type": "conflicts"}], "schema_path": ["routes", "custom", "group"], "syntax": "block", "type": "object"}, {"aliases": ["routes custom severity"], "anchor": "section", "description": "Label Matcher.", "document_id": "xcsh-docs:resources:alert_policy:properties:routes:custom:severity", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-routes--custom--severity--exact_match", "enforcement": "provider-schema", "group": "routes.custom.severity:ConflictingObjectAttributes:exact_match,regex_match", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:routes:custom:severity", "type": "conflicts"}, {"anchor": "schema-routes--custom--severity--regex_match", "enforcement": "provider-schema", "group": "routes.custom.severity:ConflictingObjectAttributes:exact_match,regex_match", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:routes:custom:severity", "type": "conflicts"}], "schema_path": ["routes", "custom", "severity"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/alert_policy/properties/routes/custom/index.txt", "spec_pin_digest": "sha256:0e278b4afb598ac8628ac74dca6a1b2831cd8607de1d26f56e72e3461a7440c7", "summary": "A set of matchers an alert has to fulfill to match the route.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.0", "schema_components": ["alert_policyCreateRequest"], "target_commit": "ac024ccbfb8b9f84412813f3e2ab9f5821937ef2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.custom

Breadcrumbs:

- [xcsh_alert_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/routes/)
- routes.custom

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
custom {
  # Configure direct properties listed below.
}
```

## Direct properties

- [alertlabel](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/routes/custom/alertlabel/): complete subsection reference.

- [alertname](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/routes/custom/alertname/): complete subsection reference.

- [group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/routes/custom/group/): complete subsection reference.

- [severity](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/routes/custom/severity/): complete subsection reference.

## Next pages

- [routes.custom.alertlabel](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/routes/custom/alertlabel/)
- [routes.custom.alertname](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/routes/custom/alertname/)
- [routes.custom.group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/routes/custom/group/)
- [routes.custom.severity](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/routes/custom/severity/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/routes/)
- [xcsh_alert_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/)
