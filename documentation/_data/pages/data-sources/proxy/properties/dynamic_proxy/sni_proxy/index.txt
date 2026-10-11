---
page_title: "dynamic_proxy.sni_proxy"
subcategory: ""
description: "Parameters for dynamic SNI proxy."
xcsh_docs: {"aliases": ["dynamic proxy sni proxy"], "body_bytes": 1626, "body_sha256": "sha256:4f4e268a81a0cd19962f5636dc81db898163f7c17efe697ed866c020a51683d6", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:sni_proxy", "parent_id": "xcsh-docs:data-sources:proxy:properties:dynamic_proxy", "path": "documentation/data-sources/proxy/properties/dynamic_proxy/sni_proxy/index.md", "product": "distributed-cloud", "provider_name": "proxy", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "data-sources", "registry_anchor": "canonical-3201203013323203-2030201011013320-2110012213112021-0333333012303323-1113131112130002-1333330220013232-0212100100120203-2000200313330302", "registry_path": "docs/guides/data-sources--proxy--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["dynamic_proxy", "sni_proxy"], "schema_version": 1, "sections": [{"aliases": ["duration", "dynamic proxy sni proxy idle timeout"], "anchor": "schema-dynamic_proxy--sni_proxy--idle_timeout", "description": "The amount of time that a stream can exist without upstream or downstream activity, in milliseconds.", "document_id": "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:sni_proxy", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dynamic_proxy", "sni_proxy", "idle_timeout"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/proxy/properties/dynamic_proxy/sni_proxy/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Parameters for dynamic SNI proxy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["proxyCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# dynamic_proxy.sni_proxy

Breadcrumbs:

- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/)
- [dynamic_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/dynamic_proxy/)
- dynamic_proxy.sni_proxy

<a id="section"></a>

Type: `"single"`. Computed.

Dynamic SNI Proxy Type. Parameters for dynamic SNI proxy.

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

## Direct properties

<a id="schema-dynamic_proxy--sni_proxy--idle_timeout"></a>

### idle_timeout property

Type: `"number"`. Computed.

The amount of time that a stream can exist without upstream or downstream activity, in milliseconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 86400000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "86400000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "86400000"
  }
}
```
