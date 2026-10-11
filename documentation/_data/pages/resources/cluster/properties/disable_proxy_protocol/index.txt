---
page_title: "disable_proxy_protocol"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["disable proxy protocol"], "body_bytes": 1512, "body_sha256": "sha256:814c722319f6746c3bef348acfa53fe54e7bca7f9a7583ded5c269da75230ded", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cluster:collection", "completeness": "complete", "id": "xcsh-docs:resources:cluster:properties:disable_proxy_protocol", "parent_id": "xcsh-docs:resources:cluster:reference", "path": "documentation/resources/cluster/properties/disable_proxy_protocol/index.md", "product": "distributed-cloud", "provider_name": "cluster", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "resources", "registry_anchor": "canonical-1332231121213303-0300331221322310-1213312031322122-0103330310021233-0230233100333021-3101001230012121-0312003112200120-0231110302130011", "registry_path": "docs/guides/resources--cluster--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["disable_proxy_protocol"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cluster/properties/disable_proxy_protocol/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["clusterCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
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

OneOf alternatives in this subsection:

- [disable_proxy_protocol](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/disable_proxy_protocol/#section)
- [proxy_protocol_v1](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/proxy_protocol_v1/#section)
- [proxy_protocol_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/proxy_protocol_v2/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_proxy_protocol = {}
```

This is an empty object or choice marker. It has no direct properties.
