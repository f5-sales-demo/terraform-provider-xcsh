---
page_title: "dynamic_proxy.sni_proxy"
subcategory: ""
description: "dynamic_proxy.sni_proxy for xcsh_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 1794, "body_sha256": "sha256:1976344d7012cdd853305111a2175108b406828043f325a870e18293a8c3b633", "canonical_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:sni_proxy", "child_ids": [], "collection_id": "xcsh-docs:resources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:sni_proxy", "parent_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy", "path": "docs/guides/resources--proxy--properties--dynamic_proxy--sni_proxy.md", "provider_name": "proxy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["dynamic_proxy", "sni_proxy"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/proxy/properties/dynamic_proxy/sni_proxy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "dynamic_proxy.sni_proxy for xcsh_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# dynamic_proxy.sni_proxy

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md)
- [Property reference](resources--proxy--reference.md)
- [dynamic_proxy](resources--proxy--properties--dynamic_proxy.md)
- dynamic_proxy.sni_proxy

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Dynamic SNI Proxy Type. Parameters for dynamic SNI proxy.

Upstream description:

Parameters for dynamic SNI proxy.

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
sni_proxy {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-dynamic_proxy--sni_proxy--idle_timeout"></a>

### idle_timeout property

Type: `"number"`. Optional.

The amount of time that a stream can exist without upstream or downstream activity, in milliseconds.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(86400000),
}
```

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
    "ves.io.schema.rules.uint32.lte": "86400000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "86400000"
  }
}
```

## Next pages

- [dynamic_proxy](resources--proxy--properties--dynamic_proxy.md)
- [xcsh_proxy](../resources/proxy.md)
