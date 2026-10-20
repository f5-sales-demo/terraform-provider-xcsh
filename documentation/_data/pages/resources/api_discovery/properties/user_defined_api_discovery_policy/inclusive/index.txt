---
page_title: "user_defined_api_discovery_policy.inclusive"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["user defined api discovery policy inclusive"], "body_bytes": 1130, "body_sha256": "sha256:9b21719e92c8dc5f45cd944e03f858c9de4bf649a81da2b516dcdd0d5d010af0", "capabilities": ["api-management"], "category": "api-management", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:api_discovery:collection", "completeness": "complete", "id": "xcsh-docs:resources:api_discovery:properties:user_defined_api_discovery_policy:inclusive", "parent_id": "xcsh-docs:resources:api_discovery:properties:user_defined_api_discovery_policy", "path": "documentation/resources/api_discovery/properties/user_defined_api_discovery_policy/inclusive/index.md", "product": "distributed-cloud", "provider_name": "api_discovery", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-2213032120131311-3200112332220110-3021221201133002-1011013213312212-0033003103311013-0320112210022330-0300213232332322-2210313210122333", "registry_path": "docs/guides/resources--api_discovery--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["user_defined_api_discovery_policy", "inclusive"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/api_discovery/properties/user_defined_api_discovery_policy/inclusive/index.txt", "spec_pin_digest": "sha256:2a27992e2c450f8ea5ba537ebc7f5c26b0dfbc946db7ef0e7e86625dba9fb492", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.4", "schema_components": ["api_discoveryCreateRequest"], "target_commit": "a9149265e7737f6841c289490cb1a3ca946de83e"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# user_defined_api_discovery_policy.inclusive

Breadcrumbs:

- [xcsh_api_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/properties/)
- [user_defined_api_discovery_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/properties/user_defined_api_discovery_policy/)
- user_defined_api_discovery_policy.inclusive

<a id="section"></a>

Type: `["object", {}]`. Optional, Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

Additional upstream details:

This can be used for messages where no values are needed.

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
inclusive = {}
```

This is an empty object or choice marker. It has no direct properties.
