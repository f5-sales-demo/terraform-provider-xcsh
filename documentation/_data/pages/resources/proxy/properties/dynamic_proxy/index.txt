---
page_title: "dynamic_proxy"
subcategory: ""
description: "Configuration parameter for dynamic proxy."
xcsh_docs: {"aliases": ["dynamic proxy"], "body_bytes": 6293, "body_sha256": "sha256:ae38722a6f0d0f09d10502443dc05e7b26da1e6d91640f41bea2ada6c9dc869a", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:proxy:properties:dynamic_proxy:disable_dns_masquerade", "xcsh-docs:resources:proxy:properties:dynamic_proxy:enable_dns_masquerade", "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy", "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy", "xcsh-docs:resources:proxy:properties:dynamic_proxy:sni_proxy"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:proxy:properties:dynamic_proxy", "parent_id": "xcsh-docs:resources:proxy:reference", "path": "documentation/resources/proxy/properties/dynamic_proxy/index.md", "product": "distributed-cloud", "provider_name": "proxy", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "resources", "registry_anchor": "canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330", "registry_path": "docs/guides/resources--proxy--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "dynamic_proxy:ConflictingObjectAttributes:disable_dns_masquerade,enable_dns_masquerade", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:disable_dns_masquerade", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "dynamic_proxy:ConflictingObjectAttributes:disable_dns_masquerade,enable_dns_masquerade", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:enable_dns_masquerade", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "dynamic_proxy:ConflictingObjectAttributes:http_proxy,https_proxy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "dynamic_proxy:ConflictingObjectAttributes:http_proxy,sni_proxy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "dynamic_proxy:ConflictingObjectAttributes:http_proxy,https_proxy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "dynamic_proxy:ConflictingObjectAttributes:https_proxy,sni_proxy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "dynamic_proxy:ConflictingObjectAttributes:http_proxy,sni_proxy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:sni_proxy", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "dynamic_proxy:ConflictingObjectAttributes:https_proxy,sni_proxy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:sni_proxy", "type": "conflicts"}, {"anchor": "schema-dynamic_proxy--domains", "enforcement": "provider-schema", "group": "dynamic_proxy:RequiredObjectAttributes:domains", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["dynamic_proxy"], "schema_version": 1, "sections": [{"aliases": ["disable dns masquerade"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:disable_dns_masquerade", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dynamic_proxy", "disable_dns_masquerade"], "syntax": "attribute", "type": "object"}, {"aliases": ["domains"], "anchor": "schema-dynamic_proxy--domains", "description": "A list of Domains to be proxied. Wildcard hosts are supported in the suffix or prefix form Supported Domains and search order: 1. Exact Domain names: www.example.com. 2. Domains starting with a Wildcard: *.example.com. Not supported Domains: - Just a Wildcard: * - A Wildcard and TLD with no root Domain: *.com. - A", "document_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dynamic_proxy", "domains"], "syntax": "attribute", "type": "list"}, {"aliases": ["enable dns masquerade"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:enable_dns_masquerade", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dynamic_proxy", "enable_dns_masquerade"], "syntax": "attribute", "type": "object"}, {"aliases": ["http proxy"], "anchor": "section", "description": "Parameters for dynamic HTTP proxy.", "document_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["dynamic_proxy", "http_proxy"], "syntax": "block", "type": "object"}, {"aliases": ["https proxy"], "anchor": "section", "description": "Parameters for dynamic HTTPS proxy.", "document_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["dynamic_proxy", "https_proxy"], "syntax": "block", "type": "object"}, {"aliases": ["sni proxy"], "anchor": "section", "description": "Parameters for dynamic SNI proxy.", "document_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:sni_proxy", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["dynamic_proxy", "sni_proxy"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/proxy/properties/dynamic_proxy/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Configuration parameter for dynamic proxy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["proxyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# dynamic_proxy

Breadcrumbs:

- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/)
- dynamic_proxy

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: dynamic\_proxy, http\_proxy\] Configuration parameter for dynamic proxy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("domains"),
  validators.ConflictingObjectAttributes("disable_dns_masquerade",
    "enable_dns_masquerade"),
  validators.ConflictingObjectAttributes("http_proxy",
    "https_proxy"),
  validators.ConflictingObjectAttributes("http_proxy",
    "sni_proxy"),
  validators.ConflictingObjectAttributes("https_proxy",
    "sni_proxy")}
```

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

- [dynamic_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/#section)
- [http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/http_proxy/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
dynamic_proxy {
  # Configure direct properties listed below.
}
```

## Direct properties

- [disable_dns_masquerade](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/disable_dns_masquerade/): complete subsection reference.

<a id="schema-dynamic_proxy--domains"></a>

### domains property

Type: `["list", "string"]`. Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 32),
}
```

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

- [enable_dns_masquerade](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/enable_dns_masquerade/): complete subsection reference.

- [http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/http_proxy/): complete subsection reference.

- [https_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/https_proxy/): complete subsection reference.

- [sni_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/sni_proxy/): complete subsection reference.

## Next pages

- [dynamic_proxy.disable_dns_masquerade](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/disable_dns_masquerade/)
- [dynamic_proxy.enable_dns_masquerade](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/enable_dns_masquerade/)
- [dynamic_proxy.http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/http_proxy/)
- [dynamic_proxy.https_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/https_proxy/)
- [dynamic_proxy.sni_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/sni_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/)
- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/)
