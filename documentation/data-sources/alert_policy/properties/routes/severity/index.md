---
page_title: "routes.severity"
subcategory: "Monitoring"
description: "Select one or more severity levels to match the incoming alert."
xcsh_docs: {"aliases": ["routes severity"], "body_bytes": 1521, "body_sha256": "sha256:78adf52134c525fa41d5f428f9a6ffc0dee8b476a88c2992c7ef1545696a4eef", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:alert_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:alert_policy:properties:routes:severity", "parent_id": "xcsh-docs:data-sources:alert_policy:properties:routes", "path": "documentation/data-sources/alert_policy/properties/routes/severity/index.md", "product": "distributed-cloud", "provider_name": "alert_policy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0023021201013110-2032203003113120-0330000033303313-1323320123013030-1130013333022020-1220223110321132-0232223012001321-2122302330130311", "registry_path": "docs/guides/data-sources--alert_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "severity"], "schema_version": 1, "sections": [{"aliases": ["severities"], "anchor": "schema-routes--severity--severities", "description": "List of severity levels.", "document_id": "xcsh-docs:data-sources:alert_policy:properties:routes:severity", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "severity", "severities"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/alert_policy/properties/routes/severity/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Select one or more severity levels to match the incoming alert.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["alert_policyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.severity

Breadcrumbs:

- [xcsh_alert_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/routes/)
- routes.severity

<a id="section"></a>

Type: `"single"`. Computed.

Select one or more severity levels to match the incoming alert.

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

<a id="schema-routes--severity--severities"></a>

### severities property

Type: `["list", "string"]`. Computed.

\[Enum: MINOR|MAJOR|CRITICAL\] Severities. List of severity levels. Possible values are \`MINOR\`,
\`MAJOR\`, \`CRITICAL\`. Defaults to \`MINOR\`.

Upstream description:

List of severity levels.

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

## Next pages

- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/routes/)
- [xcsh_alert_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/)
