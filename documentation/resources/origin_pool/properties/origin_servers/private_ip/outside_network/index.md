---
page_title: "origin_servers.private_ip.outside_network"
subcategory: "Load Balancing"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["backend servers", "origin servers", "origin servers private ip outside network", "upstream servers"], "body_bytes": 1465, "body_sha256": "sha256:12d03e76812a27be489962fe18eab0c4d06c6246dfb5147f69f85ff995d52a36", "capabilities": ["load-balancing", "load-balancing.backend-servers"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:resources:origin_pool:properties:origin_servers:private_ip:outside_network", "parent_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:private_ip", "path": "documentation/resources/origin_pool/properties/origin_servers/private_ip/outside_network/index.md", "product": "distributed-cloud", "provider_name": "origin_pool", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2332312302210220-3001332333131122-1100220121131321-2233230013312312-2320320220221301-1130102233012220-2120000130122303-2120113222003130", "registry_path": "docs/guides/resources--origin_pool--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["origin_servers", "private_ip", "outside_network"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/origin_pool/properties/origin_servers/private_ip/outside_network/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["origin_poolCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_servers.private_ip.outside_network

Breadcrumbs:

- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/)
- [origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/)
- [origin_servers.private_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/private_ip/)
- origin_servers.private_ip.outside_network

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for outside network.

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
outside_network = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [origin_servers.private_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/private_ip/)
- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/)
