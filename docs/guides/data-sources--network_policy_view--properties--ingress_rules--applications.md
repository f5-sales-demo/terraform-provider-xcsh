---
page_title: "ingress_rules.applications"
subcategory: ""
description: "ingress_rules.applications for xcsh_network_policy_view."
xcsh_docs: {"aliases": [], "body_bytes": 1577, "body_sha256": "sha256:f9d80e4cc4acadd45975ed68a89a8af3e0e31d46fe8fdfc4bb3c6537178edbcb", "canonical_id": "xcsh-docs:data-sources:network_policy_view:properties:ingress_rules:applications", "child_ids": [], "collection_id": "xcsh-docs:data-sources:network_policy_view:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_policy_view:properties:ingress_rules:applications", "parent_id": "xcsh-docs:data-sources:network_policy_view:properties:ingress_rules", "path": "docs/guides/data-sources--network_policy_view--properties--ingress_rules--applications.md", "provider_name": "network_policy_view", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ingress_rules", "applications"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_policy_view/properties/ingress_rules/applications/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_rules.applications for xcsh_network_policy_view.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_policy_viewCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_rules.applications

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md)
- [Property reference](data-sources--network_policy_view--reference.md)
- [ingress_rules](data-sources--network_policy_view--properties--ingress_rules.md)
- ingress_rules.applications

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for applications.

Upstream description:

Application protocols like HTTP, SNMP.

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

<a id="schema-ingress_rules--applications--applications"></a>

### applications property

Type: `["list", "string"]`. Computed.

\[Enum: APPLICATION\_HTTP|APPLICATION\_HTTPS|APPLICATION\_SNMP|APPLICATION\_DNS\] Application
Protocols. Application protocols like HTTP, SNMP. Possible values are \`APPLICATION\_HTTP\`,
\`APPLICATION\_HTTPS\`, \`APPLICATION\_SNMP\`, \`APPLICATION\_DNS\`. Defaults to
\`APPLICATION\_HTTP\`.

Upstream description:

Application protocols like HTTP, SNMP.

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

- [ingress_rules](data-sources--network_policy_view--properties--ingress_rules.md)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md)
