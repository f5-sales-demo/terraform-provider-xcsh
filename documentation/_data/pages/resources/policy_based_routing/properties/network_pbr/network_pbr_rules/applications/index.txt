---
page_title: "network_pbr.network_pbr_rules.applications"
subcategory: ""
description: "Application protocols like HTTP, SNMP."
xcsh_docs: {"aliases": ["network pbr network pbr rules applications"], "body_bytes": 1833, "body_sha256": "sha256:2b639f9f805289f14ce4b8db9c6ad968462298d98448abd2864d29c913a5c1cb", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:policy_based_routing:collection", "completeness": "complete", "id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:applications", "parent_id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules", "path": "documentation/resources/policy_based_routing/properties/network_pbr/network_pbr_rules/applications/index.md", "product": "distributed-cloud", "provider_name": "policy_based_routing", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-0032303031332332-1013122323122030-1030322033202310-1221123210103022-0021100013323122-3233123102222330-0220321113111211-0113232201210330", "registry_path": "docs/guides/resources--policy_based_routing--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["network_pbr", "network_pbr_rules", "applications"], "schema_version": 1, "sections": [{"aliases": ["network pbr network pbr rules applications applications"], "anchor": "schema-network_pbr--network_pbr_rules--applications--applications", "description": "Application protocols like HTTP, SNMP.", "document_id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:applications", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["network_pbr", "network_pbr_rules", "applications", "applications"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/policy_based_routing/properties/network_pbr/network_pbr_rules/applications/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "Application protocols like HTTP, SNMP.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["policy_based_routingCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# network_pbr.network_pbr_rules.applications

Breadcrumbs:

- [xcsh_policy_based_routing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/properties/)
- [network_pbr](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/properties/network_pbr/)
- [network_pbr.network_pbr_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/properties/network_pbr/network_pbr_rules/)
- network_pbr.network_pbr_rules.applications

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

<a id="schema-network_pbr--network_pbr_rules--applications--applications"></a>

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
