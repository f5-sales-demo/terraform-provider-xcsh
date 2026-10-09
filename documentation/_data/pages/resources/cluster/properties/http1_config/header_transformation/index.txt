---
page_title: "http1_config.header_transformation"
subcategory: ""
description: "Header Transformation OPTIONS for HTTP/1.1 request/response headers."
xcsh_docs: {"aliases": ["http1 config header transformation"], "body_bytes": 2284, "body_sha256": "sha256:809897b9a9491093fe197d26c7d43abf86a3439866880cc7c461d171a12709cf", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:cluster:properties:http1_config:header_transformation:default_header_transformation", "xcsh-docs:resources:cluster:properties:http1_config:header_transformation:preserve_case_header_transformation", "xcsh-docs:resources:cluster:properties:http1_config:header_transformation:proper_case_header_transformation"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cluster:collection", "completeness": "complete", "id": "xcsh-docs:resources:cluster:properties:http1_config:header_transformation", "parent_id": "xcsh-docs:resources:cluster:properties:http1_config", "path": "documentation/resources/cluster/properties/http1_config/header_transformation/index.md", "product": "distributed-cloud", "provider_name": "cluster", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-3203221021123323-3232201000310320-1220303231330220-2221303021202311-1322120131023313-3223020332000123-3210030230131311-2303313021022011", "registry_path": "docs/guides/resources--cluster--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "http1_config.header_transformation:ConflictingObjectAttributes:default_header_transformation,preserve_case_header_transformation", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cluster:properties:http1_config:header_transformation:default_header_transformation", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "http1_config.header_transformation:ConflictingObjectAttributes:default_header_transformation,proper_case_header_transformation", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cluster:properties:http1_config:header_transformation:default_header_transformation", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "http1_config.header_transformation:ConflictingObjectAttributes:default_header_transformation,preserve_case_header_transformation", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cluster:properties:http1_config:header_transformation:preserve_case_header_transformation", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "http1_config.header_transformation:ConflictingObjectAttributes:preserve_case_header_transformation,proper_case_header_transformation", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cluster:properties:http1_config:header_transformation:preserve_case_header_transformation", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "http1_config.header_transformation:ConflictingObjectAttributes:default_header_transformation,proper_case_header_transformation", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cluster:properties:http1_config:header_transformation:proper_case_header_transformation", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "http1_config.header_transformation:ConflictingObjectAttributes:preserve_case_header_transformation,proper_case_header_transformation", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cluster:properties:http1_config:header_transformation:proper_case_header_transformation", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["http1_config", "header_transformation"], "schema_version": 1, "sections": [{"aliases": ["http1 config header transformation default header transformation"], "anchor": "section", "description": "Use the platform's current default HTTP header transformation behavior.", "document_id": "xcsh-docs:resources:cluster:properties:http1_config:header_transformation:default_header_transformation", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["http1_config", "header_transformation", "default_header_transformation"], "syntax": "attribute", "type": "object"}, {"aliases": ["http1 config header transformation preserve case header transformation"], "anchor": "section", "description": "Preserve HTTP header-name case when upstream case must remain unchanged.", "document_id": "xcsh-docs:resources:cluster:properties:http1_config:header_transformation:preserve_case_header_transformation", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["http1_config", "header_transformation", "preserve_case_header_transformation"], "syntax": "attribute", "type": "object"}, {"aliases": ["http1 config header transformation proper case header transformation"], "anchor": "section", "description": "Transform HTTP header names to proper case when explicit transformation is required.", "document_id": "xcsh-docs:resources:cluster:properties:http1_config:header_transformation:proper_case_header_transformation", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["http1_config", "header_transformation", "proper_case_header_transformation"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cluster/properties/http1_config/header_transformation/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Header Transformation OPTIONS for HTTP/1.1 request/response headers.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["clusterCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# http1_config.header_transformation

Breadcrumbs:

- [xcsh_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/)
- [http1_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/http1_config/)
- http1_config.header_transformation

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Header Transformation OPTIONS for HTTP/1.1 request/response headers.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

- [default_header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/http1_config/header_transformation/default_header_transformation/): complete subsection reference.

- [preserve_case_header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/http1_config/header_transformation/preserve_case_header_transformation/): complete subsection reference.

- [proper_case_header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/http1_config/header_transformation/proper_case_header_transformation/): complete subsection reference.
