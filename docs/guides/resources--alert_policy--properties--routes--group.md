---
page_title: "routes.group"
subcategory: "Monitoring"
description: "routes.group for xcsh_alert_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1786, "body_sha256": "sha256:3151943dc887a2627cd97c1b077d47e415e02a0e3b0bb88a05aa668d538dc79a", "canonical_id": "xcsh-docs:resources:alert_policy:properties:routes:group", "child_ids": [], "collection_id": "xcsh-docs:resources:alert_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:alert_policy:properties:routes:group", "parent_id": "xcsh-docs:resources:alert_policy:properties:routes", "path": "docs/guides/resources--alert_policy--properties--routes--group.md", "provider_name": "alert_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["routes", "group"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/alert_policy/properties/routes/group/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "routes.group for xcsh_alert_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["alert_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.group

Breadcrumbs:

- [xcsh_alert_policy](../resources/alert_policy.md)
- [Property reference](resources--alert_policy--reference.md)
- [routes](resources--alert_policy--properties--routes.md)
- routes.group

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Select one or more known group names to match the incoming alert.

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
group {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-routes--group--groups"></a>

### groups property

Type: `["list", "string"]`. Optional.

\[Enum:
INFRASTRUCTURE|IAAS\_CAAS|VIRTUAL\_HOST|VOLT\_SHARE|UAM|SECURITY|TIMESERIES\_ANOMALY|SHAPE\_SECURITY|SECURITY\_CSD|CDN|SYNTHETIC\_MONITORS|TLS|SECURITY\_BOT\_DEFENSE|CLOUD\_LINK|DNS|ROUTED\_DDOS\]
Groups. Name of groups to match the alert. Possible values are \`INFRASTRUCTURE\`, \`IAAS\_CAAS\`,
\`VIRTUAL\_HOST\`, \`VOLT\_SHARE\`, \`UAM\`, \`SECURITY\`, \`TIMESERIES\_ANOMALY\`,
\`SHAPE\_SECURITY\`, \`SECURITY\_CSD\`, \`CDN\`, \`SYNTHETIC\_MONITORS\`, \`TLS\`,
\`SECURITY\_BOT\_DEFENSE\`, \`CLOUD\_LINK\`, \`DNS\`, \`ROUTED\_DDOS\`. Defaults to
\`INFRASTRUCTURE\`.

Upstream description:

Name of groups to match the alert.

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
