---
page_title: "origin_servers.origin_servers.public_ip"
subcategory: ""
description: "Specify origin server with public IP address."
xcsh_docs: {"aliases": ["origin servers origin servers public ip"], "body_bytes": 2172, "body_sha256": "sha256:285bc118206a465f60ba7e201e56f8f9ead41c95490c883ec67ee841ba366860", "capabilities": ["dns", "load-balancing.backend-servers"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:dns_proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:origin_servers:public_ip", "parent_id": "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:origin_servers", "path": "documentation/data-sources/dns_proxy/properties/origin_servers/origin_servers/public_ip/index.md", "product": "distributed-cloud", "provider_name": "dns_proxy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1130222201211312-0231111130203012-0031022133111331-0230322101311202-3230020113113213-0202213031223101-2223023003321102-2203102122311123", "registry_path": "docs/guides/data-sources--dns_proxy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["origin_servers", "origin_servers", "public_ip"], "schema_version": 1, "sections": [{"aliases": ["origin servers origin servers public ip ip"], "anchor": "schema-origin_servers--origin_servers--public_ip--ip", "description": "Exclusive with Public IPv4 address.", "document_id": "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:origin_servers:public_ip", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_servers", "origin_servers", "public_ip", "ip"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_proxy/properties/origin_servers/origin_servers/public_ip/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Specify origin server with public IP address.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["dns_proxyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_servers.origin_servers.public_ip

Breadcrumbs:

- [xcsh_dns_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/)
- [origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/)
- [origin_servers.origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/)
- origin_servers.origin_servers.public_ip

<a id="section"></a>

Type: `"single"`. Computed.

Specify origin server with public IP address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-public_ip_choice": "[\"ip\"]"
}
```

## Direct properties

<a id="schema-origin_servers--origin_servers--public_ip--ip"></a>

### ip property

Type: `"string"`. Computed.

Public IPv4. Exclusive with \[\] Public IPv4 address.

Upstream description:

Exclusive with \[\] Public IPv4 address.

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

## Next pages

- [origin_servers.origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/)
- [xcsh_dns_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/)
