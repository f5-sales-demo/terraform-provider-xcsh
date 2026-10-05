---
page_title: "ddos_profile.disable_ddos_mitigation"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["ddos profile disable ddos mitigation"], "body_bytes": 1246, "body_sha256": "sha256:75ea60342bb498bd94cce72a754e678c6aa679bb6efa524d31617d024af0173b", "capabilities": ["dns"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:dns_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_proxy:properties:ddos_profile:disable_ddos_mitigation", "parent_id": "xcsh-docs:resources:dns_proxy:properties:ddos_profile", "path": "documentation/resources/dns_proxy/properties/ddos_profile/disable_ddos_mitigation/index.md", "product": "distributed-cloud", "provider_name": "dns_proxy", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-2223032100103321-3003333220221323-0312212120210010-1010111021120133-3320220312033101-1331331110112102-3203221130110220-0332132303020203", "registry_path": "docs/guides/resources--dns_proxy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ddos_profile", "disable_ddos_mitigation"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_proxy/properties/ddos_profile/disable_ddos_mitigation/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["dns_proxyCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ddos_profile.disable_ddos_mitigation

Breadcrumbs:

- [xcsh_dns_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/)
- [ddos_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/ddos_profile/)
- ddos_profile.disable_ddos_mitigation

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
disable_ddos_mitigation = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [ddos_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/ddos_profile/)
- [xcsh_dns_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/)
