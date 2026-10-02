---
page_title: "default_pool.origin_servers.vn_private_ip"
subcategory: "Load Balancing"
description: "Specify origin server with IP on Virtual Network."
xcsh_docs: {"aliases": ["default pool origin servers vn private ip"], "body_bytes": 2920, "body_sha256": "sha256:ab24deebcc93dceb5f7fd60502942fb9c7ac5eb28d5e3a2922aae6249c374e3a", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:vn_private_ip:virtual_network"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:vn_private_ip", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers", "path": "documentation/resources/http_loadbalancer/properties/default_pool/origin_servers/vn_private_ip/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0022310013130030-2111303220210213-2020201222032233-1120103131300012-1111203033323012-1021113330333233-3121232112020030-0200301203202122", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-017.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["default_pool", "origin_servers", "vn_private_ip"], "schema_version": 1, "sections": [{"aliases": ["ip"], "anchor": "schema-default_pool--origin_servers--vn_private_ip--ip", "description": "Exclusive with IPv4 address.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:vn_private_ip", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["default_pool", "origin_servers", "vn_private_ip", "ip"], "syntax": "attribute", "type": "string"}, {"aliases": ["virtual network"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:vn_private_ip:virtual_network", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-default_pool--origin_servers--vn_private_ip--virtual_network--name", "enforcement": "provider-schema", "group": "default_pool.origin_servers.vn_private_ip.virtual_network:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:vn_private_ip:virtual_network", "type": "requires"}], "schema_path": ["default_pool", "origin_servers", "vn_private_ip", "virtual_network"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/default_pool/origin_servers/vn_private_ip/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Specify origin server with IP on Virtual Network.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_pool.origin_servers.vn_private_ip

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [default_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/)
- [default_pool.origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/origin_servers/)
- default_pool.origin_servers.vn_private_ip

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with IP on Virtual Network.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-virtual_network_ip_choice": "[\"ip\"]"
}
```

Terraform syntax:

```terraform
vn_private_ip {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-default_pool--origin_servers--vn_private_ip--ip"></a>

### ip property

Type: `"string"`. Optional.

IPv4. Exclusive with \[\] IPv4 address.

Upstream description:

Exclusive with \[\] IPv4 address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [virtual_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/origin_servers/vn_private_ip/virtual_network/): complete subsection reference.

## Next pages

- [default_pool.origin_servers.vn_private_ip.virtual_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/origin_servers/vn_private_ip/virtual_network/)
- [default_pool.origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/origin_servers/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
