---
page_title: "proxy_config.https.coalescing_options"
subcategory: ""
description: "TLS connection coalescing configuration (not compatible with mTLS)"
xcsh_docs: {"aliases": ["proxy config https coalescing options"], "body_bytes": 2606, "body_sha256": "sha256:01547ca6b27926963f14672a7d3acb3689767435ebab82094f58db995fa29545", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:coalescing_options:default_coalescing", "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:coalescing_options:strict_coalescing"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:coalescing_options", "parent_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https", "path": "documentation/resources/bigip_http_proxy/properties/proxy_config/https/coalescing_options/index.md", "product": "distributed-cloud", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0220013003110000-1011123121320123-0211100123130311-3133200103222013-0102323300313331-3332301003131101-2003012100230103-3101321311200333", "registry_path": "docs/guides/resources--bigip_http_proxy--reference--group-003.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "proxy_config.https.coalescing_options:ConflictingObjectAttributes:default_coalescing,strict_coalescing", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:coalescing_options:default_coalescing", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "proxy_config.https.coalescing_options:ConflictingObjectAttributes:default_coalescing,strict_coalescing", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:coalescing_options:strict_coalescing", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["proxy_config", "https", "coalescing_options"], "schema_version": 1, "sections": [{"aliases": ["proxy config https coalescing options default coalescing"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:coalescing_options:default_coalescing", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["proxy_config", "https", "coalescing_options", "default_coalescing"], "syntax": "attribute", "type": "object"}, {"aliases": ["proxy config https coalescing options strict coalescing"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:coalescing_options:strict_coalescing", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["proxy_config", "https", "coalescing_options", "strict_coalescing"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bigip_http_proxy/properties/proxy_config/https/coalescing_options/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "TLS connection coalescing configuration (not compatible with mTLS)", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# proxy_config.https.coalescing_options

Breadcrumbs:

- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/)
- [proxy_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/)
- [proxy_config.https](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https/)
- proxy_config.https.coalescing_options

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

TLS connection coalescing configuration (not compatible with mTLS).

Upstream description:

TLS connection coalescing configuration (not compatible with mTLS)

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_coalescing",
    "strict_coalescing")}
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
  "x-ves-oneof-field-coalescing_choice": "[\"default_coalescing\",\"strict_coalescing\"]"
}
```

Terraform syntax:

```terraform
coalescing_options {
  # Configure direct properties listed below.
}
```

## Direct properties

- [default_coalescing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https/coalescing_options/default_coalescing/): complete subsection reference.

- [strict_coalescing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https/coalescing_options/strict_coalescing/): complete subsection reference.

## Next pages

- [proxy_config.https.coalescing_options.default_coalescing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https/coalescing_options/default_coalescing/)
- [proxy_config.https.coalescing_options.strict_coalescing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https/coalescing_options/strict_coalescing/)
- [proxy_config.https](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https/)
- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/)
