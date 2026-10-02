---
page_title: "default_pool.advanced_options.http1_config.header_transformation"
subcategory: "Load Balancing"
description: "Header Transformation OPTIONS for HTTP/1.1 request/response headers."
xcsh_docs: {"aliases": ["default pool advanced options http1 config header transformation"], "body_bytes": 4031, "body_sha256": "sha256:fbacece3dc41ef99387165992970aed7b56ac26ab722d988662483b1f71acb05", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:http1_config:header_transformation:default_header_transformation", "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:http1_config:header_transformation:preserve_case_header_transformation", "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:http1_config:header_transformation:proper_case_header_transformation"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:http1_config:header_transformation", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:http1_config", "path": "documentation/resources/http_loadbalancer/properties/default_pool/advanced_options/http1_config/header_transformation/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1101330321000013-2212300010220002-1310000332322113-3003211002121311-0123331303233202-0131033311223202-2123300010020112-1321021330013301", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-016.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "default_pool.advanced_options.http1_config.header_transformation:ConflictingObjectAttributes:default_header_transformation,preserve_case_header_transformation", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:http1_config:header_transformation:default_header_transformation", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "default_pool.advanced_options.http1_config.header_transformation:ConflictingObjectAttributes:default_header_transformation,proper_case_header_transformation", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:http1_config:header_transformation:default_header_transformation", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "default_pool.advanced_options.http1_config.header_transformation:ConflictingObjectAttributes:default_header_transformation,preserve_case_header_transformation", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:http1_config:header_transformation:preserve_case_header_transformation", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "default_pool.advanced_options.http1_config.header_transformation:ConflictingObjectAttributes:preserve_case_header_transformation,proper_case_header_transformation", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:http1_config:header_transformation:preserve_case_header_transformation", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "default_pool.advanced_options.http1_config.header_transformation:ConflictingObjectAttributes:default_header_transformation,proper_case_header_transformation", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:http1_config:header_transformation:proper_case_header_transformation", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "default_pool.advanced_options.http1_config.header_transformation:ConflictingObjectAttributes:preserve_case_header_transformation,proper_case_header_transformation", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:http1_config:header_transformation:proper_case_header_transformation", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["default_pool", "advanced_options", "http1_config", "header_transformation"], "schema_version": 1, "sections": [{"aliases": ["default header transformation"], "anchor": "section", "description": "Use the platform's current default HTTP header transformation behavior.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:http1_config:header_transformation:default_header_transformation", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["default_pool", "advanced_options", "http1_config", "header_transformation", "default_header_transformation"], "syntax": "attribute", "type": "object"}, {"aliases": ["preserve case header transformation"], "anchor": "section", "description": "Preserve HTTP header-name case when upstream case must remain unchanged.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:http1_config:header_transformation:preserve_case_header_transformation", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["default_pool", "advanced_options", "http1_config", "header_transformation", "preserve_case_header_transformation"], "syntax": "attribute", "type": "object"}, {"aliases": ["proper case header transformation"], "anchor": "section", "description": "Transform HTTP header names to proper case when explicit transformation is required.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:http1_config:header_transformation:proper_case_header_transformation", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["default_pool", "advanced_options", "http1_config", "header_transformation", "proper_case_header_transformation"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/default_pool/advanced_options/http1_config/header_transformation/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Header Transformation OPTIONS for HTTP/1.1 request/response headers.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_pool.advanced_options.http1_config.header_transformation

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [default_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/)
- [default_pool.advanced_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/advanced_options/)
- [default_pool.advanced_options.http1_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/advanced_options/http1_config/)
- default_pool.advanced_options.http1_config.header_transformation

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Header Transformation OPTIONS for HTTP/1.1 request/response headers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_header_transformation",
    "preserve_case_header_transformation"),
  validators.ConflictingObjectAttributes("default_header_transformation",
    "proper_case_header_transformation"),
  validators.ConflictingObjectAttributes("preserve_case_header_transformation",
    "proper_case_header_transformation")}
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
  "x-ves-oneof-field-header_transformation_choice": "[\"default_header_transformation\",\"preserve_case_header_transformation\",\"proper_case_header_transformation\"]"
}
```

Terraform syntax:

```terraform
header_transformation {
  # Configure direct properties listed below.
}
```

## Direct properties

- [default_header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/advanced_options/http1_config/header_transformation/default_header_transformation/): complete subsection reference.

- [preserve_case_header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/advanced_options/http1_config/header_transformation/preserve_case_header_transformation/): complete subsection reference.

- [proper_case_header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/advanced_options/http1_config/header_transformation/proper_case_header_transformation/): complete subsection reference.

## Next pages

- [default_pool.advanced_options.http1_config.header_transformation.default_header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/advanced_options/http1_config/header_transformation/default_header_transformation/)
- [default_pool.advanced_options.http1_config.header_transformation.preserve_case_header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/advanced_options/http1_config/header_transformation/preserve_case_header_transformation/)
- [default_pool.advanced_options.http1_config.header_transformation.proper_case_header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/advanced_options/http1_config/header_transformation/proper_case_header_transformation/)
- [default_pool.advanced_options.http1_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/advanced_options/http1_config/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
