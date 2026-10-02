---
page_title: "disable_proxy_protocol"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["disable proxy protocol"], "body_bytes": 1746, "body_sha256": "sha256:8e3fb9558db4695706d421796d6596bcd5cf1ae23f8e874d10f7165d7bfe485e", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cluster:collection", "completeness": "complete", "id": "xcsh-docs:resources:cluster:properties:disable_proxy_protocol", "parent_id": "xcsh-docs:resources:cluster:reference", "path": "documentation/resources/cluster/properties/disable_proxy_protocol/index.md", "product": "distributed-cloud", "provider_name": "cluster", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1332231121213303-0300331221322310-1213312031322122-0103330310021233-0230233100333021-3101001230012121-0312003112200120-0231110302130011", "registry_path": "docs/guides/resources--cluster--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["disable_proxy_protocol"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cluster/properties/disable_proxy_protocol/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["clusterCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# disable_proxy_protocol

Breadcrumbs:

- [xcsh_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/)
- disable_proxy_protocol

<a id="section"></a>

Type: `["object", {}]`. Optional.

\[OneOf: disable\_proxy\_protocol, proxy\_protocol\_v1, proxy\_protocol\_v2; Default:
disable\_proxy\_protocol\] Configuration parameter for disable proxy protocol.

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

OneOf alternatives in this subsection:

- [disable_proxy_protocol](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/disable_proxy_protocol/#section)
- [proxy_protocol_v1](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/proxy_protocol_v1/#section)
- [proxy_protocol_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/proxy_protocol_v2/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_proxy_protocol = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/)
- [xcsh_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/)
