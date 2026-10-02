---
page_title: "local_ip.ip_address.virtual_network_type.site_local"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["local ip ip address virtual network type site local"], "body_bytes": 1592, "body_sha256": "sha256:f5a1f3b8a4f0b98234390b1385d344a904939408fbf5b0e7c59a8cfa7e2bd998", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:tunnel:collection", "completeness": "complete", "id": "xcsh-docs:resources:tunnel:properties:local_ip:ip_address:virtual_network_type:site_local", "parent_id": "xcsh-docs:resources:tunnel:properties:local_ip:ip_address:virtual_network_type", "path": "documentation/resources/tunnel/properties/local_ip/ip_address/virtual_network_type/site_local/index.md", "product": "distributed-cloud", "provider_name": "tunnel", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3311200211120100-2120200220132030-0303323022032102-2210030333200002-3300210100312310-0223102133023131-3220110222300230-1201202100310110", "registry_path": "docs/guides/resources--tunnel--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["local_ip", "ip_address", "virtual_network_type", "site_local"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tunnel/properties/local_ip/ip_address/virtual_network_type/site_local/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tunnelCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# local_ip.ip_address.virtual_network_type.site_local

Breadcrumbs:

- [xcsh_tunnel](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/)
- [local_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/local_ip/)
- [local_ip.ip_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/local_ip/ip_address/)
- [local_ip.ip_address.virtual_network_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/local_ip/ip_address/virtual_network_type/)
- local_ip.ip_address.virtual_network_type.site_local

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
site_local = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [local_ip.ip_address.virtual_network_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/local_ip/ip_address/virtual_network_type/)
- [xcsh_tunnel](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/)
