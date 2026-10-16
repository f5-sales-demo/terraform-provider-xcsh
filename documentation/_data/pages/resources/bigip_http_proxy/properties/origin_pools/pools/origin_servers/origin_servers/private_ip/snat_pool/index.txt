---
page_title: "origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool"
subcategory: ""
description: "SNAT Pool configuration."
xcsh_docs: {"aliases": ["origin pools pools origin servers origin servers private ip snat pool"], "body_bytes": 2269, "body_sha256": "sha256:7c20e16067c9706b4dcad67b07b5ef5044d6190902d01aabcf56512869f9d4be", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers:private_ip:snat_pool:no_snat_pool", "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers:private_ip:snat_pool:snat_pool"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers:private_ip:snat_pool", "parent_id": "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers:private_ip", "path": "documentation/resources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/private_ip/snat_pool/index.md", "product": "distributed-cloud", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "resources", "registry_anchor": "canonical-0012323013103221-2113331011300123-0022012222333113-1101000223003033-2320202111220110-1232200202020301-2020120203313201-0021111002012011", "registry_path": "docs/guides/resources--bigip_http_proxy--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["origin_pools", "pools", "origin_servers", "origin_servers", "private_ip", "snat_pool"], "schema_version": 1, "sections": [{"aliases": ["origin pools pools origin servers origin servers private ip snat pool no snat pool"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers:private_ip:snat_pool:no_snat_pool", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_pools", "pools", "origin_servers", "origin_servers", "private_ip", "snat_pool", "no_snat_pool"], "syntax": "attribute", "type": "object"}, {"aliases": ["origin pools pools origin servers origin servers private ip snat pool snat pool"], "anchor": "section", "description": "List of IPv4 prefixes that represent an endpoint.", "document_id": "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers:private_ip:snat_pool:snat_pool", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["origin_pools", "pools", "origin_servers", "origin_servers", "private_ip", "snat_pool", "snat_pool"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/private_ip/snat_pool/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "SNAT Pool configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool

Breadcrumbs:

- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/)
- [origin_pools](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/origin_pools/)
- [origin_pools.pools](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/origin_pools/pools/)
- [origin_pools.pools.origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/)
- [origin_pools.pools.origin_servers.origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/)
- [origin_pools.pools.origin_servers.origin_servers.private_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/private_ip/)
- origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

SNAT Pool. SNAT Pool configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-snat_pool_choice": "[\"no_snat_pool\",\"snat_pool\"]"
}
```

Terraform syntax:

```terraform
snat_pool {
  # Configure direct properties listed below.
}
```

## Direct properties

- [no_snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/private_ip/snat_pool/no_snat_pool/): complete subsection reference.

- [snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/private_ip/snat_pool/snat_pool/): complete subsection reference.
