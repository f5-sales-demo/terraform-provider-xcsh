---
page_title: "routes.severity"
subcategory: "Monitoring"
description: "Select one or more severity levels to match the incoming alert."
xcsh_docs: {"aliases": ["routes severity"], "body_bytes": 1521, "body_sha256": "sha256:78adf52134c525fa41d5f428f9a6ffc0dee8b476a88c2992c7ef1545696a4eef", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": [], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:alert_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:alert_policy:properties:routes:severity", "parent_id": "xcsh-docs:data-sources:alert_policy:properties:routes", "path": "documentation/data-sources/alert_policy/properties/routes/severity/index.md", "product": "distributed-cloud", "provider_name": "alert_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0023021201013110-2032203003113120-0330000033303313-1323320123013030-1130013333022020-1220223110321132-0232223012001321-2122302330130311", "registry_path": "docs/guides/data-sources--alert_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "severity"], "schema_version": 1, "sections": [{"aliases": ["severities"], "anchor": "schema-routes--severity--severities", "description": "List of severity levels.", "document_id": "xcsh-docs:data-sources:alert_policy:properties:routes:severity", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "severity", "severities"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/alert_policy/properties/routes/severity/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Select one or more severity levels to match the incoming alert.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["alert_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
