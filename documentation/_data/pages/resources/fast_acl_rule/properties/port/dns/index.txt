---
page_title: "port.dns"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["port dns"], "body_bytes": 1166, "body_sha256": "sha256:63f1f37d503d1cfaab976de481dac398e4a8057c7eba2599cc1fe59097827781", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:fast_acl_rule:collection", "completeness": "complete", "id": "xcsh-docs:resources:fast_acl_rule:properties:port:dns", "parent_id": "xcsh-docs:resources:fast_acl_rule:properties:port", "path": "documentation/resources/fast_acl_rule/properties/port/dns/index.md", "product": "distributed-cloud", "provider_name": "fast_acl_rule", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-2321310033213222-1210123211133230-0310201203322011-0113122020022133-0311122321230130-3131223320320133-1311233320003010-3300220333302233", "registry_path": "docs/guides/resources--fast_acl_rule--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["port", "dns"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fast_acl_rule/properties/port/dns/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["fast_acl_ruleCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# port.dns

Breadcrumbs:

- [xcsh_fast_acl_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/properties/)
- [port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/properties/port/)
- port.dns

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

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
dns = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/properties/port/)
- [xcsh_fast_acl_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/)
