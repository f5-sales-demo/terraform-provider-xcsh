---
page_title: "endpoint.any"
subcategory: "Security"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["endpoint any"], "body_bytes": 943, "body_sha256": "sha256:bee0f6930ebd4628788682472705abdc7a5689d192f391c2ed6ba961d32e7021", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:network_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_policy:properties:endpoint:any", "parent_id": "xcsh-docs:resources:network_policy:properties:endpoint", "path": "documentation/resources/network_policy/properties/endpoint/any/index.md", "product": "distributed-cloud", "provider_name": "network_policy", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-0121302023111133-0330100131013110-3210002030111101-1133332330132321-2312031223131003-3323232221023130-3132222310232331-3213102030300113", "registry_path": "docs/guides/resources--network_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["endpoint", "any"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_policy/properties/endpoint/any/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["network_policyCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# endpoint.any

Breadcrumbs:

- [xcsh_network_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/)
- [endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/endpoint/)
- endpoint.any

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

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
any = {}
```

This is an empty object or choice marker. It has no direct properties.
