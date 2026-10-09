---
page_title: "cloudfront.protected_endpoints.mobile_client"
subcategory: ""
description: "Mobile client configuration OPTIONS."
xcsh_docs: {"aliases": ["cloudfront protected endpoints mobile client"], "body_bytes": 1821, "body_sha256": "sha256:b58a39e595710223bce0e027f83289a37f4794c4489656e19d2380e2fdc45e4b", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:mobile_client:block", "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:mobile_client:continue"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:mobile_client", "parent_id": "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints", "path": "documentation/resources/protected_application/properties/cloudfront/protected_endpoints/mobile_client/index.md", "product": "distributed-cloud", "provider_name": "protected_application", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-0213232003133201-2213121300112120-1230023220211230-1331323331121131-3322331313113311-3231221012003322-0013220102232300-1333301123112210", "registry_path": "docs/guides/resources--protected_application--reference--group-004.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "cloudfront.protected_endpoints.mobile_client:ConflictingObjectAttributes:block,continue", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:mobile_client:block", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cloudfront.protected_endpoints.mobile_client:ConflictingObjectAttributes:block,continue", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:mobile_client:continue", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["cloudfront", "protected_endpoints", "mobile_client"], "schema_version": 1, "sections": [{"aliases": ["cloudfront protected endpoints mobile client block"], "anchor": "section", "description": "Block Response.", "document_id": "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:mobile_client:block", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["cloudfront", "protected_endpoints", "mobile_client", "block"], "syntax": "block", "type": "object"}, {"aliases": ["cloudfront protected endpoints mobile client continue"], "anchor": "section", "description": "Continue mitigation action.", "document_id": "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:mobile_client:continue", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "cloudfront.protected_endpoints.mobile_client.continue:ConflictingObjectAttributes:add_header,no_header", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:mobile_client:continue:add_header", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cloudfront.protected_endpoints.mobile_client.continue:ConflictingObjectAttributes:add_header,no_header", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:mobile_client:continue:no_header", "type": "conflicts"}], "schema_path": ["cloudfront", "protected_endpoints", "mobile_client", "continue"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protected_application/properties/cloudfront/protected_endpoints/mobile_client/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Mobile client configuration OPTIONS.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudfront.protected_endpoints.mobile_client

Breadcrumbs:

- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/)
- [cloudfront](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/)
- [cloudfront.protected_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/)
- cloudfront.protected_endpoints.mobile_client

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Mobile Client. Mobile client configuration OPTIONS.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("block",
    "continue")}
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
  "x-ves-oneof-field-mitigation": "[\"block\",\"continue\"]"
}
```

Terraform syntax:

```terraform
mobile_client {
  # Configure direct properties listed below.
}
```

## Direct properties

- [block](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/mobile_client/block/): complete subsection reference.

- [continue](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/mobile_client/continue/): complete subsection reference.
