---
page_title: "proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation"
subcategory: ""
description: "Header Transformation OPTIONS for HTTP/1.1 request/response headers."
xcsh_docs: {"aliases": ["proxy config https http protocol options http protocol enable v1 only header transformation"], "body_bytes": 4581, "body_sha256": "sha256:358fc679d03cf01c494ea7c2a3ea7532a742fb9446c2c2f3afb39a708926a60a", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:http_protocol_options:http_protocol_enable_v1_only:header_transformation:default_header_transformation", "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:http_protocol_options:http_protocol_enable_v1_only:header_transformation:preserve_case_header_transformation", "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:http_protocol_options:http_protocol_enable_v1_only:header_transformation:proper_case_header_transformation"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:http_protocol_options:http_protocol_enable_v1_only:header_transformation", "parent_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:http_protocol_options:http_protocol_enable_v1_only", "path": "documentation/resources/bigip_http_proxy/properties/proxy_config/https/http_protocol_options/http_protocol_enable_v1_only/header_transformation/index.md", "product": "distributed-cloud", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2232331113101101-1303111213011013-2133032203200333-2113230322210220-2102030302011300-3312003131232323-2010132203213013-3302123331122130", "registry_path": "docs/guides/resources--bigip_http_proxy--reference--group-003.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation:ConflictingObjectAttributes:default_header_transformation,preserve_case_header_transformation", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:http_protocol_options:http_protocol_enable_v1_only:header_transformation:default_header_transformation", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation:ConflictingObjectAttributes:default_header_transformation,proper_case_header_transformation", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:http_protocol_options:http_protocol_enable_v1_only:header_transformation:default_header_transformation", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation:ConflictingObjectAttributes:default_header_transformation,preserve_case_header_transformation", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:http_protocol_options:http_protocol_enable_v1_only:header_transformation:preserve_case_header_transformation", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation:ConflictingObjectAttributes:preserve_case_header_transformation,proper_case_header_transformation", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:http_protocol_options:http_protocol_enable_v1_only:header_transformation:preserve_case_header_transformation", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation:ConflictingObjectAttributes:default_header_transformation,proper_case_header_transformation", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:http_protocol_options:http_protocol_enable_v1_only:header_transformation:proper_case_header_transformation", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation:ConflictingObjectAttributes:preserve_case_header_transformation,proper_case_header_transformation", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:http_protocol_options:http_protocol_enable_v1_only:header_transformation:proper_case_header_transformation", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["proxy_config", "https", "http_protocol_options", "http_protocol_enable_v1_only", "header_transformation"], "schema_version": 1, "sections": [{"aliases": ["proxy config https http protocol options http protocol enable v1 only header transformation default header transformation"], "anchor": "section", "description": "Use the platform's current default HTTP header transformation behavior.", "document_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:http_protocol_options:http_protocol_enable_v1_only:header_transformation:default_header_transformation", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["proxy_config", "https", "http_protocol_options", "http_protocol_enable_v1_only", "header_transformation", "default_header_transformation"], "syntax": "attribute", "type": "object"}, {"aliases": ["proxy config https http protocol options http protocol enable v1 only header transformation preserve case header transformation"], "anchor": "section", "description": "Preserve HTTP header-name case when upstream case must remain unchanged.", "document_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:http_protocol_options:http_protocol_enable_v1_only:header_transformation:preserve_case_header_transformation", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["proxy_config", "https", "http_protocol_options", "http_protocol_enable_v1_only", "header_transformation", "preserve_case_header_transformation"], "syntax": "attribute", "type": "object"}, {"aliases": ["proxy config https http protocol options http protocol enable v1 only header transformation proper case header transformation"], "anchor": "section", "description": "Transform HTTP header names to proper case when explicit transformation is required.", "document_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:http_protocol_options:http_protocol_enable_v1_only:header_transformation:proper_case_header_transformation", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["proxy_config", "https", "http_protocol_options", "http_protocol_enable_v1_only", "header_transformation", "proper_case_header_transformation"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bigip_http_proxy/properties/proxy_config/https/http_protocol_options/http_protocol_enable_v1_only/header_transformation/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Header Transformation OPTIONS for HTTP/1.1 request/response headers.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation

Breadcrumbs:

- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/)
- [proxy_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/)
- [proxy_config.https](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https/)
- [proxy_config.https.http_protocol_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https/http_protocol_options/)
- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https/http_protocol_options/http_protocol_enable_v1_only/)
- proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation

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

- [default_header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https/http_protocol_options/http_protocol_enable_v1_only/header_transformation/default_header_transformation/): complete subsection reference.

- [preserve_case_header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https/http_protocol_options/http_protocol_enable_v1_only/header_transformation/preserve_case_header_transformation/): complete subsection reference.

- [proper_case_header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https/http_protocol_options/http_protocol_enable_v1_only/header_transformation/proper_case_header_transformation/): complete subsection reference.

## Next pages

- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https/http_protocol_options/http_protocol_enable_v1_only/header_transformation/default_header_transformation/)
- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https/http_protocol_options/http_protocol_enable_v1_only/header_transformation/preserve_case_header_transformation/)
- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https/http_protocol_options/http_protocol_enable_v1_only/header_transformation/proper_case_header_transformation/)
- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https/http_protocol_options/http_protocol_enable_v1_only/)
- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/)
