---
page_title: "dynamic_proxy"
subcategory: ""
description: "Configuration parameter for dynamic proxy."
xcsh_docs: {"aliases": ["dynamic proxy"], "body_bytes": 5619, "body_sha256": "sha256:9e63b9e69647260a559646bc9b5cab246d01ce3ec99e1b83983ebfad882ebbbc", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:proxy:properties:dynamic_proxy:disable_dns_masquerade", "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:enable_dns_masquerade", "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:http_proxy", "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:https_proxy", "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:sni_proxy"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:proxy:properties:dynamic_proxy", "parent_id": "xcsh-docs:data-sources:proxy:reference", "path": "documentation/data-sources/proxy/properties/dynamic_proxy/index.md", "product": "distributed-cloud", "provider_name": "proxy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313", "registry_path": "docs/guides/data-sources--proxy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["dynamic_proxy"], "schema_version": 1, "sections": [{"aliases": ["dynamic proxy disable dns masquerade"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:disable_dns_masquerade", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dynamic_proxy", "disable_dns_masquerade"], "syntax": "attribute", "type": "object"}, {"aliases": ["dynamic proxy domains"], "anchor": "schema-dynamic_proxy--domains", "description": "A list of Domains to be proxied. Wildcard hosts are supported in the suffix or prefix form Supported Domains and search order: 1. Exact Domain names: www.example.com. 2. Domains starting with a Wildcard: *.example.com. Not supported Domains: - Just a Wildcard: * - A Wildcard and TLD with no root Domain: *.com. - A", "document_id": "xcsh-docs:data-sources:proxy:properties:dynamic_proxy", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dynamic_proxy", "domains"], "syntax": "attribute", "type": "list"}, {"aliases": ["dynamic proxy enable dns masquerade"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:enable_dns_masquerade", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dynamic_proxy", "enable_dns_masquerade"], "syntax": "attribute", "type": "object"}, {"aliases": ["dynamic proxy http proxy"], "anchor": "section", "description": "Parameters for dynamic HTTP proxy.", "document_id": "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:http_proxy", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["dynamic_proxy", "http_proxy"], "syntax": "attribute", "type": "object"}, {"aliases": ["dynamic proxy https proxy"], "anchor": "section", "description": "Parameters for dynamic HTTPS proxy.", "document_id": "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:https_proxy", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["dynamic_proxy", "https_proxy"], "syntax": "attribute", "type": "object"}, {"aliases": ["dynamic proxy sni proxy"], "anchor": "section", "description": "Parameters for dynamic SNI proxy.", "document_id": "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:sni_proxy", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["dynamic_proxy", "sni_proxy"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/proxy/properties/dynamic_proxy/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Configuration parameter for dynamic proxy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["proxyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# dynamic_proxy

Breadcrumbs:

- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/)
- dynamic_proxy

<a id="section"></a>

Type: `"single"`. Computed.

\[OneOf: dynamic\_proxy, http\_proxy\] Configuration parameter for dynamic proxy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-dns_choice": "[\"disable_dns_masquerade\",\"enable_dns_masquerade\"]",
  "x-ves-oneof-field-proxy_choice": "[\"http_proxy\",\"https_proxy\",\"sni_proxy\"]"
}
```

OneOf alternatives in this subsection:

- [dynamic_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/dynamic_proxy/#section)
- [http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/http_proxy/#section)

Select alternatives according to the provider validators above.

## Direct properties

- [disable_dns_masquerade](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/dynamic_proxy/disable_dns_masquerade/): complete subsection reference.

<a id="schema-dynamic_proxy--domains"></a>

### domains property

Type: `["list", "string"]`. Computed.

List of Domains to be proxied. Wildcard hosts are supported in the suffix or prefix form Supported
Domains and search order: 1. Exact Domain names: www&#46;example.com. 2.

Upstream description:

A list of Domains to be proxied. Wildcard hosts are supported in the suffix or prefix form

Supported Domains and search order: &#8203;1. Exact Domain names: www&#46;example.com. &#8203;2.
Domains starting with a Wildcard: \*.example.com.

Not supported Domains: &#8203;- Just a Wildcard: \* &#8203;- A Wildcard and TLD with no root Domain:
\*.com. &#8203;- A Wildcard not matching a whole DNS label. E.g. \*.example.com and
\*.bar.example.com are valid Wildcards however \*bar.example.com, \*-bar.example.com, and
bar\*.example.com are all invalid.

Additional notes: A Wildcard will not match empty string. E.g. \*.example.com will match
bar.example.com and baz-bar.example.com but not .example.com. The longest Wildcards match first.
Only a single virtual host in the entire route configuration can match on \*. Also a Domain must be
unique across all virtual hosts within an advertise policy.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
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
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [enable_dns_masquerade](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/dynamic_proxy/enable_dns_masquerade/): complete subsection reference.

- [http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/dynamic_proxy/http_proxy/): complete subsection reference.

- [https_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/dynamic_proxy/https_proxy/): complete subsection reference.

- [sni_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/dynamic_proxy/sni_proxy/): complete subsection reference.

## Next pages

- [dynamic_proxy.disable_dns_masquerade](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/dynamic_proxy/disable_dns_masquerade/)
- [dynamic_proxy.enable_dns_masquerade](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/dynamic_proxy/enable_dns_masquerade/)
- [dynamic_proxy.http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/dynamic_proxy/http_proxy/)
- [dynamic_proxy.https_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/dynamic_proxy/https_proxy/)
- [dynamic_proxy.sni_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/dynamic_proxy/sni_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/)
- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/)
