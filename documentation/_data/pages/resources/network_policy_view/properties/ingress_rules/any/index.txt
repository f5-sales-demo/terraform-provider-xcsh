---
page_title: "ingress_rules.any"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["ingress rules any"], "body_bytes": 983, "body_sha256": "sha256:052febf7180835a196d8c87a3ce0117881e0764eb829262b64f325e8637babd7", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:network_policy_view:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_policy_view:properties:ingress_rules:any", "parent_id": "xcsh-docs:resources:network_policy_view:properties:ingress_rules", "path": "documentation/resources/network_policy_view/properties/ingress_rules/any/index.md", "product": "distributed-cloud", "provider_name": "network_policy_view", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "resources", "registry_anchor": "canonical-0110023221122000-1132101133202201-3023120323102212-3020232013300203-2000221202212311-1311223103102230-2121331101302003-1333233203223211", "registry_path": "docs/guides/resources--network_policy_view--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ingress_rules", "any"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_policy_view/properties/ingress_rules/any/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": ["network_policy_viewCreateRequest"], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_rules.any

Breadcrumbs:

- [xcsh_network_policy_view](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy_view/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy_view/properties/)
- [ingress_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy_view/properties/ingress_rules/)
- ingress_rules.any

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
