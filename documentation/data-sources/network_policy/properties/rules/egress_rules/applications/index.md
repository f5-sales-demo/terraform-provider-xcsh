---
page_title: "rules.egress_rules.applications"
subcategory: "Security"
description: "Application protocols like HTTP, SNMP."
xcsh_docs: {"aliases": ["rules egress rules applications"], "body_bytes": 1948, "body_sha256": "sha256:d6a9efb457ad47b9acdcb0d6fb7cc3ccb38b3d877939462687da9647d27fcca7", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:network_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_policy:properties:rules:egress_rules:applications", "parent_id": "xcsh-docs:data-sources:network_policy:properties:rules:egress_rules", "path": "documentation/data-sources/network_policy/properties/rules/egress_rules/applications/index.md", "product": "distributed-cloud", "provider_name": "network_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3102123131233031-2120101113230113-0020021220223203-0301330100002131-3222120133321221-0220010023032013-0331332323310202-3320002100220031", "registry_path": "docs/guides/data-sources--network_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rules", "egress_rules", "applications"], "schema_version": 1, "sections": [{"aliases": ["applications"], "anchor": "schema-rules--egress_rules--applications--applications", "description": "Application protocols like HTTP, SNMP.", "document_id": "xcsh-docs:data-sources:network_policy:properties:rules:egress_rules:applications", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "egress_rules", "applications", "applications"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_policy/properties/rules/egress_rules/applications/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Application protocols like HTTP, SNMP.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.egress_rules.applications

Breadcrumbs:

- [xcsh_network_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy/properties/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy/properties/rules/)
- [rules.egress_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy/properties/rules/egress_rules/)
- rules.egress_rules.applications

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

<a id="schema-rules--egress_rules--applications--applications"></a>

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

- [rules.egress_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy/properties/rules/egress_rules/)
- [xcsh_network_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy/)
