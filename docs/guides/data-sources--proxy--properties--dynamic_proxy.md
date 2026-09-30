---
page_title: "dynamic_proxy"
subcategory: ""
description: "dynamic_proxy for xcsh_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 4710, "body_sha256": "sha256:12ff22c8e5584ea26c68349ed9190b654495c1adbbd2105ecbba9d040a8979ee", "canonical_id": "xcsh-docs:data-sources:proxy:properties:dynamic_proxy", "child_ids": ["xcsh-docs:data-sources:proxy:properties:dynamic_proxy:disable_dns_masquerade", "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:enable_dns_masquerade", "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:http_proxy", "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:https_proxy", "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:sni_proxy"], "collection_id": "xcsh-docs:data-sources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:proxy:properties:dynamic_proxy", "parent_id": "xcsh-docs:data-sources:proxy:reference", "path": "docs/guides/data-sources--proxy--properties--dynamic_proxy.md", "provider_name": "proxy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["dynamic_proxy"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/proxy/properties/dynamic_proxy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "dynamic_proxy for xcsh_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# dynamic_proxy

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md)
- [Property reference](data-sources--proxy--reference.md)
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

- [dynamic_proxy](data-sources--proxy--properties--dynamic_proxy.md#section)
- [http_proxy](data-sources--proxy--properties--http_proxy.md#section)

Select alternatives according to the provider validators above.

## Direct properties

- [disable_dns_masquerade](data-sources--proxy--properties--dynamic_proxy--disable_dns_masquerade.md): complete subsection reference.

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

- [enable_dns_masquerade](data-sources--proxy--properties--dynamic_proxy--enable_dns_masquerade.md): complete subsection reference.

- [http_proxy](data-sources--proxy--properties--dynamic_proxy--http_proxy.md): complete subsection reference.

- [https_proxy](data-sources--proxy--properties--dynamic_proxy--https_proxy.md): complete subsection reference.

- [sni_proxy](data-sources--proxy--properties--dynamic_proxy--sni_proxy.md): complete subsection reference.

## Next pages

- [dynamic_proxy.disable_dns_masquerade](data-sources--proxy--properties--dynamic_proxy--disable_dns_masquerade.md)
- [dynamic_proxy.enable_dns_masquerade](data-sources--proxy--properties--dynamic_proxy--enable_dns_masquerade.md)
- [dynamic_proxy.http_proxy](data-sources--proxy--properties--dynamic_proxy--http_proxy.md)
- [dynamic_proxy.https_proxy](data-sources--proxy--properties--dynamic_proxy--https_proxy.md)
- [dynamic_proxy.sni_proxy](data-sources--proxy--properties--dynamic_proxy--sni_proxy.md)
- [Property reference](data-sources--proxy--reference.md)
- [xcsh_proxy](../data-sources/proxy.md)
