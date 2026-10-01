---
page_title: "routes.severity"
subcategory: "Monitoring"
description: "routes.severity for xcsh_alert_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1366, "body_sha256": "sha256:ca857ce0d04589a9747db2851844d4fbef34843c2df039be090cd2d78f3b487b", "canonical_id": "xcsh-docs:resources:alert_policy:properties:routes:severity", "child_ids": [], "collection_id": "xcsh-docs:resources:alert_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:alert_policy:properties:routes:severity", "parent_id": "xcsh-docs:resources:alert_policy:properties:routes", "path": "docs/guides/resources--alert_policy--properties--routes--severity.md", "provider_name": "alert_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["routes", "severity"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/alert_policy/properties/routes/severity/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "routes.severity for xcsh_alert_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["alert_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.severity

Breadcrumbs:

- [xcsh_alert_policy](../resources/alert_policy.md)
- [Property reference](resources--alert_policy--reference.md)
- [routes](resources--alert_policy--properties--routes.md)
- routes.severity

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
severity {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-routes--severity--severities"></a>

### severities property

Type: `["list", "string"]`. Optional.

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

- [routes](resources--alert_policy--properties--routes.md)
- [xcsh_alert_policy](../resources/alert_policy.md)
