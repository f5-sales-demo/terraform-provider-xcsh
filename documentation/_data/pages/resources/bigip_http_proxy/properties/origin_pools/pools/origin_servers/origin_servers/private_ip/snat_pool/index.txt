---
page_title: "origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool"
subcategory: ""
description: "SNAT Pool configuration."
xcsh_docs: {"aliases": ["origin pools pools origin servers origin servers private ip snat pool"], "body_bytes": 3362, "body_sha256": "sha256:74fc4e98f631a3dbfa24d45fd92c4d84f46b7576c6d1b7cece3ce6271c89b137", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers:private_ip:snat_pool:no_snat_pool", "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers:private_ip:snat_pool:snat_pool"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers:private_ip:snat_pool", "parent_id": "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers:private_ip", "path": "documentation/resources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/private_ip/snat_pool/index.md", "product": "distributed-cloud", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-0012323013103221-2113331011300123-0022012222333113-1101000223003033-2320202111220110-1232200202020301-2020120203313201-0021111002012011", "registry_path": "docs/guides/resources--bigip_http_proxy--reference--group-002.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool:ConflictingObjectAttributes:no_snat_pool,snat_pool", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers:private_ip:snat_pool:no_snat_pool", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool:ConflictingObjectAttributes:no_snat_pool,snat_pool", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers:private_ip:snat_pool:snat_pool", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["origin_pools", "pools", "origin_servers", "origin_servers", "private_ip", "snat_pool"], "schema_version": 1, "sections": [{"aliases": ["origin pools pools origin servers origin servers private ip snat pool no snat pool"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers:private_ip:snat_pool:no_snat_pool", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_pools", "pools", "origin_servers", "origin_servers", "private_ip", "snat_pool", "no_snat_pool"], "syntax": "attribute", "type": "object"}, {"aliases": ["origin pools pools origin servers origin servers private ip snat pool snat pool"], "anchor": "section", "description": "List of IPv4 prefixes that represent an endpoint.", "document_id": "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers:private_ip:snat_pool:snat_pool", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["origin_pools", "pools", "origin_servers", "origin_servers", "private_ip", "snat_pool", "snat_pool"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/private_ip/snat_pool/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "SNAT Pool configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
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

Upstream description:

SNAT Pool configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("no_snat_pool",
    "snat_pool")}
```

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

## Next pages

- [origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool.no_snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/private_ip/snat_pool/no_snat_pool/)
- [origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool.snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/private_ip/snat_pool/snat_pool/)
- [origin_pools.pools.origin_servers.origin_servers.private_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/private_ip/)
- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/)
