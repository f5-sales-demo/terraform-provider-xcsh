---
page_title: "forward_proxy_pbr.forward_proxy_pbr_rules.all_destinations"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["forward proxy pbr forward proxy pbr rules all destinations"], "body_bytes": 1305, "body_sha256": "sha256:47757e2eaeb4667138b19471ada6d8f34f9aa77f1fe50f61847953e09d0c58b6", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:policy_based_routing:collection", "completeness": "complete", "id": "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:all_destinations", "parent_id": "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules", "path": "documentation/resources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/all_destinations/index.md", "product": "distributed-cloud", "provider_name": "policy_based_routing", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-2003230233203132-0311032232211232-3030021301211033-1002333012032333-0012220020202310-0310320331113312-2330330311023031-0203232311300220", "registry_path": "docs/guides/resources--policy_based_routing--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["forward_proxy_pbr", "forward_proxy_pbr_rules", "all_destinations"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/all_destinations/index.txt", "spec_pin_digest": "sha256:ea0d80973b5efe05eb3575be6beda9c96b86701d264dc1aee3972e6a23ebb3fb", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.3", "schema_components": ["policy_based_routingCreateRequest"], "target_commit": "6ee1e85b2e917d58078f38a270b790077b3fba12"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# forward_proxy_pbr.forward_proxy_pbr_rules.all_destinations

Breadcrumbs:

- [xcsh_policy_based_routing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/properties/)
- [forward_proxy_pbr](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/properties/forward_proxy_pbr/)
- [forward_proxy_pbr.forward_proxy_pbr_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/)
- forward_proxy_pbr.forward_proxy_pbr_rules.all_destinations

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for all destinations.

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
all_destinations = {}
```

This is an empty object or choice marker. It has no direct properties.
