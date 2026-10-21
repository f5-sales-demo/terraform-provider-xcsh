---
page_title: "egress_rules.applications"
subcategory: ""
description: "Application protocols like HTTP, SNMP."
xcsh_docs: {"aliases": ["egress rules applications"], "body_bytes": 1615, "body_sha256": "sha256:665e4a1f8fdf6a163f027358993e9a82227aeb39b67f9a128dfa6f44b8b669df", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:network_policy_view:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:applications", "parent_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules", "path": "documentation/resources/network_policy_view/properties/egress_rules/applications/index.md", "product": "distributed-cloud", "provider_name": "network_policy_view", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-2231120103111332-1031333103211212-2112303213120211-3301032110010030-3003202221103300-1231111330101233-3101302210133303-3313002010201213", "registry_path": "docs/guides/resources--network_policy_view--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["egress_rules", "applications"], "schema_version": 1, "sections": [{"aliases": ["egress rules applications applications"], "anchor": "schema-egress_rules--applications--applications", "description": "Application protocols like HTTP, SNMP.", "document_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:applications", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["egress_rules", "applications", "applications"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_policy_view/properties/egress_rules/applications/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "Application protocols like HTTP, SNMP.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["network_policy_viewCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# egress_rules.applications

Breadcrumbs:

- [xcsh_network_policy_view](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy_view/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy_view/properties/)
- [egress_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy_view/properties/egress_rules/)
- egress_rules.applications

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for applications.

Additional upstream details:

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

Terraform syntax:

```terraform
applications {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-egress_rules--applications--applications"></a>

### applications property

Type: `["list", "string"]`. Optional.

\[Enum: APPLICATION\_HTTP|APPLICATION\_HTTPS|APPLICATION\_SNMP|APPLICATION\_DNS\] Application
Protocols. Application protocols like HTTP, SNMP. Possible values are \`APPLICATION\_HTTP\`,
\`APPLICATION\_HTTPS\`, \`APPLICATION\_SNMP\`, \`APPLICATION\_DNS\`. Defaults to
\`APPLICATION\_HTTP\`.

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
