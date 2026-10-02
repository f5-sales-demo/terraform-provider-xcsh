---
page_title: "user_id_client_ip"
subcategory: "Load Balancing"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["user id client ip"], "body_bytes": 1651, "body_sha256": "sha256:4bbf29ae4be9f643bf9ed1e58ad6efd7366cc671abbb088ec53952acd015b1f9", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:user_id_client_ip", "parent_id": "xcsh-docs:resources:http_loadbalancer:reference", "path": "documentation/resources/http_loadbalancer/properties/user_id_client_ip/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3332013311013213-1233202210102101-2202301332332033-2130302002311222-1231331220201121-2021122303232113-1311030310321233-1201033022300000", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-027.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["user_id_client_ip"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/user_id_client_ip/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# user_id_client_ip

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- user_id_client_ip

<a id="section"></a>

Type: `["object", {}]`. Optional, Computed.

\[OneOf: user\_id\_client\_ip, user\_identification\] Enable this option. Defaults to \`map\[\]\`.
Server applies default when omitted.

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

- [user_id_client_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/user_id_client_ip/#section)
- [user_identification](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/user_identification/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
user_id_client_ip = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
