---
page_title: "https_auto_cert.tls_config"
subcategory: "Load Balancing"
description: "This defines various OPTIONS to configure TLS configuration parameters."
xcsh_docs: {"aliases": ["https auto cert tls config"], "body_bytes": 2287, "body_sha256": "sha256:48be77bc1f6c488b746ae4bcf3e3042f1ed1facd0a28b6338b29e8b80d0ae9ec", "capabilities": ["cdn", "load-balancing.tls"], "category": "cdn", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:https_auto_cert:tls_config:tls_11_plus", "xcsh-docs:resources:cdn_loadbalancer:properties:https_auto_cert:tls_config:tls_12_plus"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:https_auto_cert:tls_config", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:https_auto_cert", "path": "documentation/resources/cdn_loadbalancer/properties/https_auto_cert/tls_config/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3330112310003221-2210231203321133-0110232120313131-0321132311312322-0013000323120230-1231130303233303-0030111013100230-0032311220323110", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-011.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "https_auto_cert.tls_config:ConflictingObjectAttributes:tls_11_plus,tls_12_plus", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:https_auto_cert:tls_config:tls_11_plus", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "https_auto_cert.tls_config:ConflictingObjectAttributes:tls_11_plus,tls_12_plus", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:https_auto_cert:tls_config:tls_12_plus", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["https_auto_cert", "tls_config"], "schema_version": 1, "sections": [{"aliases": ["https auto cert tls config tls 11 plus"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:https_auto_cert:tls_config:tls_11_plus", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["https_auto_cert", "tls_config", "tls_11_plus"], "syntax": "attribute", "type": "object"}, {"aliases": ["https auto cert tls config tls 12 plus"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:https_auto_cert:tls_config:tls_12_plus", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["https_auto_cert", "tls_config", "tls_12_plus"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/https_auto_cert/tls_config/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This defines various OPTIONS to configure TLS configuration parameters.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# https_auto_cert.tls_config

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [https_auto_cert](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/https_auto_cert/)
- https_auto_cert.tls_config

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Defines various OPTIONS to configure TLS configuration parameters.

Upstream description:

This defines various OPTIONS to configure TLS configuration parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("tls_11_plus",
    "tls_12_plus")}
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
  "x-ves-oneof-field-choice": "[\"tls_11_plus\",\"tls_12_plus\"]"
}
```

Terraform syntax:

```terraform
tls_config {
  # Configure direct properties listed below.
}
```

## Direct properties

- [tls_11_plus](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/https_auto_cert/tls_config/tls_11_plus/): complete subsection reference.

- [tls_12_plus](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/https_auto_cert/tls_config/tls_12_plus/): complete subsection reference.

## Next pages

- [https_auto_cert.tls_config.tls_11_plus](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/https_auto_cert/tls_config/tls_11_plus/)
- [https_auto_cert.tls_config.tls_12_plus](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/https_auto_cert/tls_config/tls_12_plus/)
- [https_auto_cert](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/https_auto_cert/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
