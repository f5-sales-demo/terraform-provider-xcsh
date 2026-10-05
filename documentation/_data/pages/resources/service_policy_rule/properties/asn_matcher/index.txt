---
page_title: "asn_matcher"
subcategory: ""
description: "Match any AS number contained in the list of bgp_asn_sets."
xcsh_docs: {"aliases": ["asn matcher"], "body_bytes": 1524, "body_sha256": "sha256:8db40defc7dd6d57cdc2382991742d34f4f711627fe59a345050097cc7d31b65", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:service_policy_rule:properties:asn_matcher:asn_sets"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:service_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy_rule:properties:asn_matcher", "parent_id": "xcsh-docs:resources:service_policy_rule:reference", "path": "documentation/resources/service_policy_rule/properties/asn_matcher/index.md", "product": "distributed-cloud", "provider_name": "service_policy_rule", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-1303131313113301-2002011311201202-0113230121212323-0202022331303111-2010020122331301-0002311130331320-1111101133131202-0133331233122210", "registry_path": "docs/guides/resources--service_policy_rule--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "asn_matcher:RequiredObjectAttributes:asn_sets", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:asn_matcher:asn_sets", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["asn_matcher"], "schema_version": 1, "sections": [{"aliases": ["asn matcher asn sets"], "anchor": "section", "description": "A list of references to bgp_asn_set objects.", "document_id": "xcsh-docs:resources:service_policy_rule:properties:asn_matcher:asn_sets", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["asn_matcher", "asn_sets"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy_rule/properties/asn_matcher/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Match any AS number contained in the list of bgp_asn_sets.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["service_policy_ruleCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# asn_matcher

Breadcrumbs:

- [xcsh_service_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/)
- asn_matcher

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Match any AS number contained in the list of bgp\_asn\_sets.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("asn_sets")}
```

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
asn_matcher {
  # Configure direct properties listed below.
}
```

## Direct properties

- [asn_sets](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/asn_matcher/asn_sets/): complete subsection reference.

## Next pages

- [asn_matcher.asn_sets](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/asn_matcher/asn_sets/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/)
- [xcsh_service_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/)
