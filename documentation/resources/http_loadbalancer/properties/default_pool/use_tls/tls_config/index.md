---
page_title: "default_pool.use_tls.tls_config"
subcategory: "Load Balancing"
description: "This defines various OPTIONS to configure TLS configuration parameters."
xcsh_docs: {"aliases": ["default pool use tls tls config"], "body_bytes": 4112, "body_sha256": "sha256:0d8452a53514f5c947d413c9593fccd66a9a24f75b78d790472d2f078c4ebe3a", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:tls_config:custom_security", "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:tls_config:default_security", "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:tls_config:low_security", "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:tls_config:medium_security"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:tls_config", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls", "path": "documentation/resources/http_loadbalancer/properties/default_pool/use_tls/tls_config/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3022001312010311-2033321103031000-0230211202123211-3032221333033333-2023331303010222-3103112113331032-3303302020113030-0113122302220112", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-017.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "default_pool.use_tls.tls_config:ConflictingObjectAttributes:custom_security,default_security", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:tls_config:custom_security", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "default_pool.use_tls.tls_config:ConflictingObjectAttributes:custom_security,low_security", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:tls_config:custom_security", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "default_pool.use_tls.tls_config:ConflictingObjectAttributes:custom_security,medium_security", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:tls_config:custom_security", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "default_pool.use_tls.tls_config:ConflictingObjectAttributes:custom_security,default_security", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:tls_config:default_security", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "default_pool.use_tls.tls_config:ConflictingObjectAttributes:default_security,low_security", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:tls_config:default_security", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "default_pool.use_tls.tls_config:ConflictingObjectAttributes:default_security,medium_security", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:tls_config:default_security", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "default_pool.use_tls.tls_config:ConflictingObjectAttributes:custom_security,low_security", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:tls_config:low_security", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "default_pool.use_tls.tls_config:ConflictingObjectAttributes:default_security,low_security", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:tls_config:low_security", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "default_pool.use_tls.tls_config:ConflictingObjectAttributes:low_security,medium_security", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:tls_config:low_security", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "default_pool.use_tls.tls_config:ConflictingObjectAttributes:custom_security,medium_security", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:tls_config:medium_security", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "default_pool.use_tls.tls_config:ConflictingObjectAttributes:default_security,medium_security", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:tls_config:medium_security", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "default_pool.use_tls.tls_config:ConflictingObjectAttributes:low_security,medium_security", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:tls_config:medium_security", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["default_pool", "use_tls", "tls_config"], "schema_version": 1, "sections": [{"aliases": ["custom security"], "anchor": "section", "description": "This defines TLS protocol config including min/max versions and allowed ciphers.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:tls_config:custom_security", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-default_pool--use_tls--tls_config--custom_security--cipher_suites", "enforcement": "provider-schema", "group": "default_pool.use_tls.tls_config.custom_security:RequiredObjectAttributes:cipher_suites", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:tls_config:custom_security", "type": "requires"}], "schema_path": ["default_pool", "use_tls", "tls_config", "custom_security"], "syntax": "block", "type": "object"}, {"aliases": ["default security"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:tls_config:default_security", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["default_pool", "use_tls", "tls_config", "default_security"], "syntax": "attribute", "type": "object"}, {"aliases": ["low security"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:tls_config:low_security", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["default_pool", "use_tls", "tls_config", "low_security"], "syntax": "attribute", "type": "object"}, {"aliases": ["medium security"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:tls_config:medium_security", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["default_pool", "use_tls", "tls_config", "medium_security"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/default_pool/use_tls/tls_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This defines various OPTIONS to configure TLS configuration parameters.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_pool.use_tls.tls_config

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [default_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/)
- [default_pool.use_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/use_tls/)
- default_pool.use_tls.tls_config

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

- [custom_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/use_tls/tls_config/custom_security/): complete subsection reference.

- [default_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/use_tls/tls_config/default_security/): complete subsection reference.

- [low_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/use_tls/tls_config/low_security/): complete subsection reference.

- [medium_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/use_tls/tls_config/medium_security/): complete subsection reference.

## Next pages

- [default_pool.use_tls.tls_config.custom_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/use_tls/tls_config/custom_security/)
- [default_pool.use_tls.tls_config.default_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/use_tls/tls_config/default_security/)
- [default_pool.use_tls.tls_config.low_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/use_tls/tls_config/low_security/)
- [default_pool.use_tls.tls_config.medium_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/use_tls/tls_config/medium_security/)
- [default_pool.use_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/use_tls/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
