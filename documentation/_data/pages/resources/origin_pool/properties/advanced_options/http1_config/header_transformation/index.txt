---
page_title: "advanced_options.http1_config.header_transformation"
subcategory: "Load Balancing"
description: "Header Transformation OPTIONS for HTTP/1.1 request/response headers."
xcsh_docs: {"aliases": ["advanced options http1 config header transformation"], "body_bytes": 2078, "body_sha256": "sha256:e9e14d0c2528b5eeb93abd4a40a8cab059efc5a1987047c391a662d461ecac2a", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:origin_pool:properties:advanced_options:http1_config:header_transformation:default_header_transformation", "xcsh-docs:resources:origin_pool:properties:advanced_options:http1_config:header_transformation:preserve_case_header_transformation", "xcsh-docs:resources:origin_pool:properties:advanced_options:http1_config:header_transformation:proper_case_header_transformation"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:resources:origin_pool:properties:advanced_options:http1_config:header_transformation", "parent_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:http1_config", "path": "documentation/resources/origin_pool/properties/advanced_options/http1_config/header_transformation/index.md", "product": "distributed-cloud", "provider_name": "origin_pool", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-2303031223333310-0232131331321322-3223230333233112-3131212303323012-3113232223013321-1010312321321233-0032313121131111-1033120330313113", "registry_path": "docs/guides/resources--origin_pool--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["advanced_options", "http1_config", "header_transformation"], "schema_version": 1, "sections": [{"aliases": ["advanced options http1 config header transformation default header transformation"], "anchor": "section", "description": "Use the platform's current default HTTP header transformation behavior.", "document_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:http1_config:header_transformation:default_header_transformation", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["advanced_options", "http1_config", "header_transformation", "default_header_transformation"], "syntax": "attribute", "type": "object"}, {"aliases": ["advanced options http1 config header transformation preserve case header transformation"], "anchor": "section", "description": "Preserve HTTP header-name case when upstream case must remain unchanged.", "document_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:http1_config:header_transformation:preserve_case_header_transformation", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["advanced_options", "http1_config", "header_transformation", "preserve_case_header_transformation"], "syntax": "attribute", "type": "object"}, {"aliases": ["advanced options http1 config header transformation proper case header transformation"], "anchor": "section", "description": "Transform HTTP header names to proper case when explicit transformation is required.", "document_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:http1_config:header_transformation:proper_case_header_transformation", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["advanced_options", "http1_config", "header_transformation", "proper_case_header_transformation"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/origin_pool/properties/advanced_options/http1_config/header_transformation/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Header Transformation OPTIONS for HTTP/1.1 request/response headers.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["origin_poolCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# advanced_options.http1_config.header_transformation

Breadcrumbs:

- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/)
- [advanced_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/)
- [advanced_options.http1_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/http1_config/)
- advanced_options.http1_config.header_transformation

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Header Transformation OPTIONS for HTTP/1.1 request/response headers.

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

- [default_header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/http1_config/header_transformation/default_header_transformation/): complete subsection reference.

- [preserve_case_header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/http1_config/header_transformation/preserve_case_header_transformation/): complete subsection reference.

- [proper_case_header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/http1_config/header_transformation/proper_case_header_transformation/): complete subsection reference.
