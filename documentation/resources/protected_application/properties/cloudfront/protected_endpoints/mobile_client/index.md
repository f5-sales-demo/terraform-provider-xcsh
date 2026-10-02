---
page_title: "cloudfront.protected_endpoints.mobile_client"
subcategory: ""
description: "Mobile client configuration OPTIONS."
xcsh_docs: {"aliases": ["cloudfront protected endpoints mobile client"], "body_bytes": 2576, "body_sha256": "sha256:5c7807896d313e9b991585296f70cb69218b490336e5276b10c674b5eb021c58", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:mobile_client:block", "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:mobile_client:continue"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:mobile_client", "parent_id": "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints", "path": "documentation/resources/protected_application/properties/cloudfront/protected_endpoints/mobile_client/index.md", "product": "distributed-cloud", "provider_name": "protected_application", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0213232003133201-2213121300112120-1230023220211230-1331323331121131-3322331313113311-3231221012003322-0013220102232300-1333301123112210", "registry_path": "docs/guides/resources--protected_application--reference--group-004.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "cloudfront.protected_endpoints.mobile_client:ConflictingObjectAttributes:block,continue", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:mobile_client:block", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cloudfront.protected_endpoints.mobile_client:ConflictingObjectAttributes:block,continue", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:mobile_client:continue", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["cloudfront", "protected_endpoints", "mobile_client"], "schema_version": 1, "sections": [{"aliases": ["block"], "anchor": "section", "description": "Block Response.", "document_id": "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:mobile_client:block", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["cloudfront", "protected_endpoints", "mobile_client", "block"], "syntax": "block", "type": "object"}, {"aliases": ["continue"], "anchor": "section", "description": "Continue mitigation action.", "document_id": "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:mobile_client:continue", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "cloudfront.protected_endpoints.mobile_client.continue:ConflictingObjectAttributes:add_header,no_header", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:mobile_client:continue:add_header", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cloudfront.protected_endpoints.mobile_client.continue:ConflictingObjectAttributes:add_header,no_header", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:mobile_client:continue:no_header", "type": "conflicts"}], "schema_path": ["cloudfront", "protected_endpoints", "mobile_client", "continue"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protected_application/properties/cloudfront/protected_endpoints/mobile_client/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Mobile client configuration OPTIONS.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

Upstream description:

Mobile client configuration OPTIONS.

Provider validators and defaults (from schema source):

```go
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

## Next pages

- [cloudfront.protected_endpoints.mobile_client.block](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/mobile_client/block/)
- [cloudfront.protected_endpoints.mobile_client.continue](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/mobile_client/continue/)
- [cloudfront.protected_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/)
- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/)
