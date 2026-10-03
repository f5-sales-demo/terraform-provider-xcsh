---
page_title: "use_tls.tls_config"
subcategory: "Load Balancing"
description: "This defines various OPTIONS to configure TLS configuration parameters."
xcsh_docs: {"aliases": ["use tls tls config"], "body_bytes": 3661, "body_sha256": "sha256:03863f2965b00125c0810e67d0e3edef90b900c587556d3847b2391fbe7a41ec", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:origin_pool:properties:use_tls:tls_config:custom_security", "xcsh-docs:resources:origin_pool:properties:use_tls:tls_config:default_security", "xcsh-docs:resources:origin_pool:properties:use_tls:tls_config:low_security", "xcsh-docs:resources:origin_pool:properties:use_tls:tls_config:medium_security"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:resources:origin_pool:properties:use_tls:tls_config", "parent_id": "xcsh-docs:resources:origin_pool:properties:use_tls", "path": "documentation/resources/origin_pool/properties/use_tls/tls_config/index.md", "product": "distributed-cloud", "provider_name": "origin_pool", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1112222132222202-3113101010112012-2301033031100332-1302203012210011-1202110221111220-2012033212332201-1202112331303003-3201200232221131", "registry_path": "docs/guides/resources--origin_pool--reference--group-003.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "use_tls.tls_config:ConflictingObjectAttributes:custom_security,default_security", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:use_tls:tls_config:custom_security", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "use_tls.tls_config:ConflictingObjectAttributes:custom_security,low_security", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:use_tls:tls_config:custom_security", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "use_tls.tls_config:ConflictingObjectAttributes:custom_security,medium_security", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:use_tls:tls_config:custom_security", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "use_tls.tls_config:ConflictingObjectAttributes:custom_security,default_security", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:use_tls:tls_config:default_security", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "use_tls.tls_config:ConflictingObjectAttributes:default_security,low_security", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:use_tls:tls_config:default_security", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "use_tls.tls_config:ConflictingObjectAttributes:default_security,medium_security", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:use_tls:tls_config:default_security", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "use_tls.tls_config:ConflictingObjectAttributes:custom_security,low_security", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:use_tls:tls_config:low_security", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "use_tls.tls_config:ConflictingObjectAttributes:default_security,low_security", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:use_tls:tls_config:low_security", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "use_tls.tls_config:ConflictingObjectAttributes:low_security,medium_security", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:use_tls:tls_config:low_security", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "use_tls.tls_config:ConflictingObjectAttributes:custom_security,medium_security", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:use_tls:tls_config:medium_security", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "use_tls.tls_config:ConflictingObjectAttributes:default_security,medium_security", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:use_tls:tls_config:medium_security", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "use_tls.tls_config:ConflictingObjectAttributes:low_security,medium_security", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:use_tls:tls_config:medium_security", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["use_tls", "tls_config"], "schema_version": 1, "sections": [{"aliases": ["use tls tls config custom security"], "anchor": "section", "description": "This defines TLS protocol config including min/max versions and allowed ciphers.", "document_id": "xcsh-docs:resources:origin_pool:properties:use_tls:tls_config:custom_security", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-use_tls--tls_config--custom_security--cipher_suites", "enforcement": "provider-schema", "group": "use_tls.tls_config.custom_security:RequiredObjectAttributes:cipher_suites", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:use_tls:tls_config:custom_security", "type": "requires"}], "schema_path": ["use_tls", "tls_config", "custom_security"], "syntax": "block", "type": "object"}, {"aliases": ["use tls tls config default security"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:origin_pool:properties:use_tls:tls_config:default_security", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["use_tls", "tls_config", "default_security"], "syntax": "attribute", "type": "object"}, {"aliases": ["use tls tls config low security"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:origin_pool:properties:use_tls:tls_config:low_security", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["use_tls", "tls_config", "low_security"], "syntax": "attribute", "type": "object"}, {"aliases": ["use tls tls config medium security"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:origin_pool:properties:use_tls:tls_config:medium_security", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["use_tls", "tls_config", "medium_security"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/origin_pool/properties/use_tls/tls_config/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This defines various OPTIONS to configure TLS configuration parameters.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["origin_poolCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# use_tls.tls_config

Breadcrumbs:

- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/)
- [use_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/use_tls/)
- use_tls.tls_config

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Defines various OPTIONS to configure TLS configuration parameters.

Upstream description:

This defines various OPTIONS to configure TLS configuration parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_security",
    "default_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "low_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("default_security",
    "low_security"),
  validators.ConflictingObjectAttributes("default_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("low_security",
    "medium_security")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "tls",
    "constraintType": "object",
    "deterministic": true,
    "metadata": {
      "category": "tls",
      "confidence": 0.99,
      "note": "Required when use_tls is selected — API returns 400 if nil",
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
  "x-ves-oneof-field-choice": "[\"custom_security\",\"default_security\",\"low_security\",\"medium_security\"]"
}
```

Terraform syntax:

```terraform
tls_config {
  # Configure direct properties listed below.
}
```

## Direct properties

- [custom_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/use_tls/tls_config/custom_security/): complete subsection reference.

- [default_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/use_tls/tls_config/default_security/): complete subsection reference.

- [low_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/use_tls/tls_config/low_security/): complete subsection reference.

- [medium_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/use_tls/tls_config/medium_security/): complete subsection reference.

## Next pages

- [use_tls.tls_config.custom_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/use_tls/tls_config/custom_security/)
- [use_tls.tls_config.default_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/use_tls/tls_config/default_security/)
- [use_tls.tls_config.low_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/use_tls/tls_config/low_security/)
- [use_tls.tls_config.medium_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/use_tls/tls_config/medium_security/)
- [use_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/use_tls/)
- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/)
