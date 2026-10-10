---
page_title: "custom_proxy_bypass"
subcategory: ""
description: "List of domains to bypass the proxy."
xcsh_docs: {"aliases": ["custom proxy bypass"], "body_bytes": 2449, "body_sha256": "sha256:2daec9a052374a44880316459e3cd184aaafc33669e6b8bee2d4dad9eb6c4041", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:custom_proxy_bypass", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:reference", "path": "documentation/data-sources/securemesh_site_v2/properties/custom_proxy_bypass/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-1031033310322022-0301132212030010-1323201221211321-3313320110313211-3210322323302033-2222331020312200-2202123022123110-3100133101310020", "registry_path": "docs/guides/data-sources--securemesh_site_v2--reference--group-006.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_proxy_bypass"], "schema_version": 1, "sections": [{"aliases": ["custom proxy bypass proxy bypass"], "anchor": "schema-custom_proxy_bypass--proxy_bypass", "description": "List of domains to bypass the proxy.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:custom_proxy_bypass", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_proxy_bypass", "proxy_bypass"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/custom_proxy_bypass/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "List of domains to bypass the proxy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_proxy_bypass

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/)
- custom_proxy_bypass

<a id="section"></a>

Type: `"single"`. Computed.

\[OneOf: custom\_proxy\_bypass, no\_proxy\_bypass; Default: no\_proxy\_bypass\] Configuration
parameter for custom proxy bypass.

Additional upstream details:

List of domains to bypass the proxy.

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

- [custom_proxy_bypass](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/custom_proxy_bypass/#section)
- [no_proxy_bypass](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/no_proxy_bypass/#section)

Select alternatives according to the provider validators above.

## Direct properties

<a id="schema-custom_proxy_bypass--proxy_bypass"></a>

### proxy_bypass property

Type: `["list", "string"]`. Computed.

Proxy Bypass. List of domains to bypass the proxy.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.items.string.hostname_or_ip": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.hostname_or_ip": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```
