---
page_title: "origin_servers.origin_servers.k8s_service.snat_pool.snat_pool"
subcategory: ""
description: "List of IPv4 prefixes that represent an endpoint."
xcsh_docs: {"aliases": ["origin servers origin servers k8s service snat pool snat pool"], "body_bytes": 3066, "body_sha256": "sha256:6ddd28519258101b734ee41b3a7742c0923092fd6926f59013b1f9d204cf3761", "capabilities": ["dns", "load-balancing.backend-servers"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:dns_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:origin_servers:k8s_service:snat_pool:snat_pool", "parent_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:origin_servers:k8s_service:snat_pool", "path": "documentation/resources/dns_proxy/properties/origin_servers/origin_servers/k8s_service/snat_pool/snat_pool/index.md", "product": "distributed-cloud", "provider_name": "dns_proxy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0310023313013120-0013121001322102-3111232210211221-1102223131023201-0232121120323020-0312332301003003-2123130201121330-0230300010333303", "registry_path": "docs/guides/resources--dns_proxy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["origin_servers", "origin_servers", "k8s_service", "snat_pool", "snat_pool"], "schema_version": 1, "sections": [{"aliases": ["prefixes"], "anchor": "schema-origin_servers--origin_servers--k8s_service--snat_pool--snat_pool--prefixes", "description": "List of IPv4 prefixes that represent an endpoint.", "document_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:origin_servers:k8s_service:snat_pool:snat_pool", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_servers", "origin_servers", "k8s_service", "snat_pool", "snat_pool", "prefixes"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_proxy/properties/origin_servers/origin_servers/k8s_service/snat_pool/snat_pool/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "List of IPv4 prefixes that represent an endpoint.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["dns_proxyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_servers.origin_servers.k8s_service.snat_pool.snat_pool

Breadcrumbs:

- [xcsh_dns_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/)
- [origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/origin_servers/)
- [origin_servers.origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/origin_servers/origin_servers/)
- [origin_servers.origin_servers.k8s_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/origin_servers/origin_servers/k8s_service/)
- [origin_servers.origin_servers.k8s_service.snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/origin_servers/origin_servers/k8s_service/snat_pool/)
- origin_servers.origin_servers.k8s_service.snat_pool.snat_pool

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

List of IPv4 prefixes that represent an endpoint.

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
snat_pool {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-origin_servers--origin_servers--k8s_service--snat_pool--snat_pool--prefixes"></a>

### prefixes property

Type: `["list", "string"]`. Optional.

List of IPv4 prefixes that represent an endpoint.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Next pages

- [origin_servers.origin_servers.k8s_service.snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/origin_servers/origin_servers/k8s_service/snat_pool/)
- [xcsh_dns_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/)
