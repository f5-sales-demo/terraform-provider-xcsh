---
page_title: "lb_algorithm.round_robin"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["lb algorithm round robin"], "body_bytes": 1245, "body_sha256": "sha256:fd0bb5970384895465fa3a117d1ecfc94f0b014513bde318a308dda6a747f938", "capabilities": ["dns"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:dns_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_proxy:properties:lb_algorithm:round_robin", "parent_id": "xcsh-docs:resources:dns_proxy:properties:lb_algorithm", "path": "documentation/resources/dns_proxy/properties/lb_algorithm/round_robin/index.md", "product": "distributed-cloud", "provider_name": "dns_proxy", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-2332223102321133-1211332210000331-0313300331013111-2112332101333331-2330231002323332-2023323330023123-2110102021103332-2321321311022232", "registry_path": "docs/guides/resources--dns_proxy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["lb_algorithm", "round_robin"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_proxy/properties/lb_algorithm/round_robin/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["dns_proxyCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# lb_algorithm.round_robin

Breadcrumbs:

- [xcsh_dns_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/)
- [lb_algorithm](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/lb_algorithm/)
- lb_algorithm.round_robin

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for round robin.

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
round_robin {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [lb_algorithm](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/lb_algorithm/)
- [xcsh_dns_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/)
