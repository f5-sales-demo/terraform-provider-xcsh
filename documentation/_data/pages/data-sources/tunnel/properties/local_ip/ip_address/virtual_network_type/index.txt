---
page_title: "local_ip.ip_address.virtual_network_type"
subcategory: ""
description: "Different types of virtual networks understood by the system."
xcsh_docs: {"aliases": ["local ip ip address virtual network type"], "body_bytes": 2465, "body_sha256": "sha256:e036584528ce5adbfda79038419a3c85a48af73f55343222da19cc5ddf632a0d", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:tunnel:properties:local_ip:ip_address:virtual_network_type:public", "xcsh-docs:data-sources:tunnel:properties:local_ip:ip_address:virtual_network_type:site_local", "xcsh-docs:data-sources:tunnel:properties:local_ip:ip_address:virtual_network_type:site_local_inside"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:tunnel:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tunnel:properties:local_ip:ip_address:virtual_network_type", "parent_id": "xcsh-docs:data-sources:tunnel:properties:local_ip:ip_address", "path": "documentation/data-sources/tunnel/properties/local_ip/ip_address/virtual_network_type/index.md", "product": "distributed-cloud", "provider_name": "tunnel", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "data-sources", "registry_anchor": "canonical-3330000213100001-2310200332122010-2031312312233110-2322231331011002-0323311022322333-1123031011101001-0000010302321030-0221121113120030", "registry_path": "docs/guides/data-sources--tunnel--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["local_ip", "ip_address", "virtual_network_type"], "schema_version": 1, "sections": [{"aliases": ["public"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:tunnel:properties:local_ip:ip_address:virtual_network_type:public", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["local_ip", "ip_address", "virtual_network_type", "public"], "syntax": "attribute", "type": "object"}, {"aliases": ["site local"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:tunnel:properties:local_ip:ip_address:virtual_network_type:site_local", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["local_ip", "ip_address", "virtual_network_type", "site_local"], "syntax": "attribute", "type": "object"}, {"aliases": ["site local inside"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:tunnel:properties:local_ip:ip_address:virtual_network_type:site_local_inside", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["local_ip", "ip_address", "virtual_network_type", "site_local_inside"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tunnel/properties/local_ip/ip_address/virtual_network_type/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Different types of virtual networks understood by the system.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["tunnelCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# local_ip.ip_address.virtual_network_type

Breadcrumbs:

- [xcsh_tunnel](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/)
- [local_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/local_ip/)
- [local_ip.ip_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/local_ip/ip_address/)
- local_ip.ip_address.virtual_network_type

<a id="section"></a>

Type: `"single"`. Computed.

Different types of virtual networks understood by the system.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-vn_type_choice": "[\"public\",\"site_local\",\"site_local_inside\"]"
}
```

## Direct properties

- [public](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/local_ip/ip_address/virtual_network_type/public/): complete subsection reference.

- [site_local](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/local_ip/ip_address/virtual_network_type/site_local/): complete subsection reference.

- [site_local_inside](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/local_ip/ip_address/virtual_network_type/site_local_inside/): complete subsection reference.

## Next pages

- [local_ip.ip_address.virtual_network_type.public](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/local_ip/ip_address/virtual_network_type/public/)
- [local_ip.ip_address.virtual_network_type.site_local](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/local_ip/ip_address/virtual_network_type/site_local/)
- [local_ip.ip_address.virtual_network_type.site_local_inside](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/local_ip/ip_address/virtual_network_type/site_local_inside/)
- [local_ip.ip_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/local_ip/ip_address/)
- [xcsh_tunnel](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/)
