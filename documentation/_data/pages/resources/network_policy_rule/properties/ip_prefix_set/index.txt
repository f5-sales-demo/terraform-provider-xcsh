---
page_title: "ip_prefix_set"
subcategory: ""
description: "A list of references to ip_prefix_set objects."
xcsh_docs: {"aliases": ["ip prefix set"], "body_bytes": 1607, "body_sha256": "sha256:5bfbf252fbe964aec059ecf2c3493238db5c79b54a667575d9fe9b52efa9c02d", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:network_policy_rule:properties:ip_prefix_set:ref"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:network_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_policy_rule:properties:ip_prefix_set", "parent_id": "xcsh-docs:resources:network_policy_rule:reference", "path": "documentation/resources/network_policy_rule/properties/ip_prefix_set/index.md", "product": "distributed-cloud", "provider_name": "network_policy_rule", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-1330023003330311-0120320212023000-3011113312022212-3321212113003322-1322030310230311-0020213122332002-2032202202130303-0101011123223132", "registry_path": "docs/guides/resources--network_policy_rule--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ip_prefix_set"], "schema_version": 1, "sections": [{"aliases": ["ip prefix set ref"], "anchor": "section", "description": "A list of references to ip_prefix_set objects.", "document_id": "xcsh-docs:resources:network_policy_rule:properties:ip_prefix_set:ref", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["ip_prefix_set", "ref"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_policy_rule/properties/ip_prefix_set/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "A list of references to ip_prefix_set objects.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["network_policy_ruleCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ip_prefix_set

Breadcrumbs:

- [xcsh_network_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy_rule/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy_rule/properties/)
- ip_prefix_set

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: ip\_prefix\_set, prefix, prefix\_selector\] List of references to ip\_prefix\_set objects.

Additional upstream details:

A list of references to ip\_prefix\_set objects.

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

OneOf alternatives in this subsection:

- [ip_prefix_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy_rule/properties/ip_prefix_set/#section)
- [prefix](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy_rule/properties/prefix/#section)
- [prefix_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy_rule/properties/prefix_selector/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
ip_prefix_set {
  # Configure direct properties listed below.
}
```

## Direct properties

- [ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy_rule/properties/ip_prefix_set/ref/): complete subsection reference.
