---
page_title: "origin_servers.consul_service.snat_pool"
subcategory: "Load Balancing"
description: "origin_servers.consul_service.snat_pool for xcsh_origin_pool."
xcsh_docs: {"aliases": [], "body_bytes": 2084, "body_sha256": "sha256:ffea2c85272aed07e18640dde67c271dc0bffe7ebc4f678c38497d2116a6ab75", "child_ids": ["xcsh-docs:data-sources:origin_pool:properties:origin_servers:consul_service:snat_pool:no_snat_pool", "xcsh-docs:data-sources:origin_pool:properties:origin_servers:consul_service:snat_pool:snat_pool"], "collection_id": "xcsh-docs:data-sources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:origin_pool:properties:origin_servers:consul_service:snat_pool", "parent_id": "xcsh-docs:data-sources:origin_pool:properties:origin_servers:consul_service", "path": "documentation/data-sources/origin_pool/properties/origin_servers/consul_service/snat_pool/index.md", "provider_name": "origin_pool", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["origin_servers", "consul_service", "snat_pool"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/origin_pool/properties/origin_servers/consul_service/snat_pool/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "origin_servers.consul_service.snat_pool for xcsh_origin_pool.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# origin_servers.consul_service.snat_pool

Breadcrumbs:

- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/)
- [origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/origin_servers/)
- [origin_servers.consul_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/origin_servers/consul_service/)
- origin_servers.consul_service.snat_pool

<a id="section"></a>

Type: `"single"`. Computed.

SNAT Pool. SNAT Pool configuration.

Upstream description:

SNAT Pool configuration.

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

## Direct properties

- [no_snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/origin_servers/consul_service/snat_pool/no_snat_pool/): complete subsection reference.

- [snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/origin_servers/consul_service/snat_pool/snat_pool/): complete subsection reference.

## Next pages

- [origin_servers.consul_service.snat_pool.no_snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/origin_servers/consul_service/snat_pool/no_snat_pool/)
- [origin_servers.consul_service.snat_pool.snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/origin_servers/consul_service/snat_pool/snat_pool/)
- [origin_servers.consul_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/origin_servers/consul_service/)
- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/)
