---
page_title: "custom_proxy_bypass"
subcategory: ""
description: "List of domains to bypass the proxy."
xcsh_docs: {"aliases": ["custom proxy bypass"], "body_bytes": 2726, "body_sha256": "sha256:66517c2d8b654b0409e6e543d99a2367d83f6c00e8348694719bf8a135f0ded6", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:custom_proxy_bypass", "parent_id": "xcsh-docs:resources:securemesh_site_v2:reference", "path": "documentation/resources/securemesh_site_v2/properties/custom_proxy_bypass/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:057968f86e4ef0ae0087dd4d6097131e285b60998d97655ee1a02c00315f6b0f", "provider_type": "resources", "registry_anchor": "canonical-1201120122021230-3013113003303010-0312222210310120-1032233221132010-0003011010201111-3212223322123210-1233330103221220-3111220102212230", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-006.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_proxy_bypass"], "schema_version": 1, "sections": [{"aliases": ["custom proxy bypass proxy bypass"], "anchor": "schema-custom_proxy_bypass--proxy_bypass", "description": "List of domains to bypass the proxy.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:custom_proxy_bypass", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_proxy_bypass", "proxy_bypass"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/custom_proxy_bypass/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "List of domains to bypass the proxy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_proxy_bypass

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- custom_proxy_bypass

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

- [custom_proxy_bypass](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/custom_proxy_bypass/#section)
- [no_proxy_bypass](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/no_proxy_bypass/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
custom_proxy_bypass {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-custom_proxy_bypass--proxy_bypass"></a>

### proxy_bypass property

Type: `["list", "string"]`. Optional.

Proxy Bypass. List of domains to bypass the proxy.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
