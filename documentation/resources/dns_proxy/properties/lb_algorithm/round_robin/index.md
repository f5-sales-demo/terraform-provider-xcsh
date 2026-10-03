---
page_title: "lb_algorithm.round_robin"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["lb algorithm round robin"], "body_bytes": 1245, "body_sha256": "sha256:fd0bb5970384895465fa3a117d1ecfc94f0b014513bde318a308dda6a747f938", "capabilities": ["dns"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:dns_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_proxy:properties:lb_algorithm:round_robin", "parent_id": "xcsh-docs:resources:dns_proxy:properties:lb_algorithm", "path": "documentation/resources/dns_proxy/properties/lb_algorithm/round_robin/index.md", "product": "distributed-cloud", "provider_name": "dns_proxy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2332223102321133-1211332210000331-0313300331013111-2112332101333331-2330231002323332-2023323330023123-2110102021103332-2321321311022232", "registry_path": "docs/guides/resources--dns_proxy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["lb_algorithm", "round_robin"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_proxy/properties/lb_algorithm/round_robin/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["dns_proxyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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
