---
page_title: "default_pool.origin_servers.private_ip"
subcategory: "Load Balancing"
description: "Specify origin server with private or public IP address and site information."
xcsh_docs: {"aliases": ["default pool origin servers private ip"], "body_bytes": 2987, "body_sha256": "sha256:f8cab8457f7268d935ef1a1a120e711a0302f64939dcd72a3c5ba4289e7498b2", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:private_ip:inside_network", "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:private_ip:outside_network", "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:private_ip:segment", "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:private_ip:site_locator", "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:private_ip:snat_pool"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:private_ip", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers", "path": "documentation/data-sources/http_loadbalancer/properties/default_pool/origin_servers/private_ip/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-0011033200022011-2330321200020232-3033333103222033-2212313221112010-3303213020232130-1303311022310010-0112020221103100-3223101011022033", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-016.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["default_pool", "origin_servers", "private_ip"], "schema_version": 1, "sections": [{"aliases": ["default pool origin servers private ip inside network"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:private_ip:inside_network", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["default_pool", "origin_servers", "private_ip", "inside_network"], "syntax": "attribute", "type": "object"}, {"aliases": ["default pool origin servers private ip ip"], "anchor": "schema-default_pool--origin_servers--private_ip--ip", "description": "Exclusive with Private IPv4 address.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:private_ip", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["default_pool", "origin_servers", "private_ip", "ip"], "syntax": "attribute", "type": "string"}, {"aliases": ["default pool origin servers private ip outside network"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:private_ip:outside_network", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["default_pool", "origin_servers", "private_ip", "outside_network"], "syntax": "attribute", "type": "object"}, {"aliases": ["default pool origin servers private ip segment"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:private_ip:segment", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["default_pool", "origin_servers", "private_ip", "segment"], "syntax": "attribute", "type": "object"}, {"aliases": ["default pool origin servers private ip site locator"], "anchor": "section", "description": "This message defines a reference to a site or virtual site object.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:private_ip:site_locator", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["default_pool", "origin_servers", "private_ip", "site_locator"], "syntax": "attribute", "type": "object"}, {"aliases": ["default pool origin servers private ip snat pool"], "anchor": "section", "description": "SNAT Pool configuration.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:private_ip:snat_pool", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["default_pool", "origin_servers", "private_ip", "snat_pool"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/default_pool/origin_servers/private_ip/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Specify origin server with private or public IP address and site information.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_pool.origin_servers.private_ip

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [default_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/)
- [default_pool.origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/origin_servers/)
- default_pool.origin_servers.private_ip

<a id="section"></a>

Type: `"single"`. Computed.

Specify origin server with private or public IP address and site information.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-network_choice": "[\"inside_network\",\"outside_network\",\"segment\"]",
  "x-ves-oneof-field-private_ip_choice": "[\"ip\"]"
}
```

## Direct properties

- [inside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/origin_servers/private_ip/inside_network/): complete subsection reference.

<a id="schema-default_pool--origin_servers--private_ip--ip"></a>

### ip property

Type: `"string"`. Computed.

IP. Exclusive with \[\] Private IPv4 address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

- [outside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/origin_servers/private_ip/outside_network/): complete subsection reference.

- [segment](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/origin_servers/private_ip/segment/): complete subsection reference.

- [site_locator](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/origin_servers/private_ip/site_locator/): complete subsection reference.

- [snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/origin_servers/private_ip/snat_pool/): complete subsection reference.
