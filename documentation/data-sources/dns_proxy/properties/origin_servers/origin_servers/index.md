---
page_title: "origin_servers.origin_servers"
subcategory: ""
description: "List of origin servers for Proxy."
xcsh_docs: {"aliases": ["backend servers", "origin servers", "origin servers origin servers", "upstream servers"], "body_bytes": 3774, "body_sha256": "sha256:4c9bd001353bccfa08808f92e3ec505b9f0537b8bc7ef6eb64592285bcde7b38", "capabilities": ["dns", "load-balancing.backend-servers"], "category": "dns", "child_ids": ["xcsh-docs:data-sources:dns_proxy:properties:origin_servers:origin_servers:k8s_service", "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:origin_servers:no_preference", "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:origin_servers:public_ip", "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:origin_servers:public_name", "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:origin_servers:site_preferences"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:dns_proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:origin_servers", "parent_id": "xcsh-docs:data-sources:dns_proxy:properties:origin_servers", "path": "documentation/data-sources/dns_proxy/properties/origin_servers/origin_servers/index.md", "product": "distributed-cloud", "provider_name": "dns_proxy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1333131311220333-1231120200230111-1310003132330122-1302102303223232-1331131102023222-3003000012212311-2300033331021112-0011131102103213", "registry_path": "docs/guides/data-sources--dns_proxy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["origin_servers", "origin_servers"], "schema_version": 1, "sections": [{"aliases": ["k8s service"], "anchor": "section", "description": "Specify origin server with K8s service name and site information.", "document_id": "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:origin_servers:k8s_service", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["origin_servers", "origin_servers", "k8s_service"], "syntax": "attribute", "type": "object"}, {"aliases": ["no preference"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:origin_servers:no_preference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_servers", "origin_servers", "no_preference"], "syntax": "attribute", "type": "object"}, {"aliases": ["public ip"], "anchor": "section", "description": "Specify origin server with public IP address.", "document_id": "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:origin_servers:public_ip", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["origin_servers", "origin_servers", "public_ip"], "syntax": "attribute", "type": "object"}, {"aliases": ["public name"], "anchor": "section", "description": "Specify origin server with public DNS name.", "document_id": "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:origin_servers:public_name", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["origin_servers", "origin_servers", "public_name"], "syntax": "attribute", "type": "object"}, {"aliases": ["site preferences"], "anchor": "section", "description": "Carries the references to one or more sites.", "document_id": "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:origin_servers:site_preferences", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["origin_servers", "origin_servers", "site_preferences"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_proxy/properties/origin_servers/origin_servers/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "List of origin servers for Proxy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["dns_proxyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_servers.origin_servers

Breadcrumbs:

- [xcsh_dns_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/)
- [origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/)
- origin_servers.origin_servers

<a id="section"></a>

Type: `"list"`. Computed.

List Of Origin Servers. List of origin servers for Proxy.

Upstream description:

List of origin servers for Proxy.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Direct properties

- [k8s_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/k8s_service/): complete subsection reference.

- [no_preference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/no_preference/): complete subsection reference.

- [public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/public_ip/): complete subsection reference.

- [public_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/public_name/): complete subsection reference.

- [site_preferences](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/site_preferences/): complete subsection reference.

## Next pages

- [origin_servers.origin_servers.k8s_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/k8s_service/)
- [origin_servers.origin_servers.no_preference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/no_preference/)
- [origin_servers.origin_servers.public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/public_ip/)
- [origin_servers.origin_servers.public_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/public_name/)
- [origin_servers.origin_servers.site_preferences](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/site_preferences/)
- [origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/)
- [xcsh_dns_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/)
