---
page_title: "cloudfront.mobile_sdk_config.mobile_identifier"
subcategory: ""
description: "Mobile traffic identifier type."
xcsh_docs: {"aliases": ["cloudfront mobile sdk config mobile identifier"], "body_bytes": 1953, "body_sha256": "sha256:b2639ba8e1901f6d015ad967b481528319f0550ce07c1300d41b93275681c37f", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:protected_application:properties:cloudfront:mobile_sdk_config:mobile_identifier:headers"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:resources:protected_application:properties:cloudfront:mobile_sdk_config:mobile_identifier", "parent_id": "xcsh-docs:resources:protected_application:properties:cloudfront:mobile_sdk_config", "path": "documentation/resources/protected_application/properties/cloudfront/mobile_sdk_config/mobile_identifier/index.md", "product": "distributed-cloud", "provider_name": "protected_application", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0213223320221112-2122101302230222-1313301132220012-2212300132233123-3031030123331203-3003133332121301-1031000123123233-1033210032003123", "registry_path": "docs/guides/resources--protected_application--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["cloudfront", "mobile_sdk_config", "mobile_identifier"], "schema_version": 1, "sections": [{"aliases": ["headers"], "anchor": "section", "description": "A list of headers that can be used to identify mobile traffic.", "document_id": "xcsh-docs:resources:protected_application:properties:cloudfront:mobile_sdk_config:mobile_identifier:headers", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-cloudfront--mobile_sdk_config--mobile_identifier--headers--exact", "enforcement": "provider-schema", "group": "cloudfront.mobile_sdk_config.mobile_identifier.headers:ConflictingListObjectAttributes:exact,regex", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudfront:mobile_sdk_config:mobile_identifier:headers", "type": "conflicts"}, {"anchor": "schema-cloudfront--mobile_sdk_config--mobile_identifier--headers--regex", "enforcement": "provider-schema", "group": "cloudfront.mobile_sdk_config.mobile_identifier.headers:ConflictingListObjectAttributes:exact,regex", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudfront:mobile_sdk_config:mobile_identifier:headers", "type": "conflicts"}, {"anchor": "schema-cloudfront--mobile_sdk_config--mobile_identifier--headers--name", "enforcement": "provider-schema", "group": "cloudfront.mobile_sdk_config.mobile_identifier.headers:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudfront:mobile_sdk_config:mobile_identifier:headers", "type": "requires"}], "schema_path": ["cloudfront", "mobile_sdk_config", "mobile_identifier", "headers"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protected_application/properties/cloudfront/mobile_sdk_config/mobile_identifier/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Mobile traffic identifier type.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudfront.mobile_sdk_config.mobile_identifier

Breadcrumbs:

- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/)
- [cloudfront](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/)
- [cloudfront.mobile_sdk_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/mobile_sdk_config/)
- cloudfront.mobile_sdk_config.mobile_identifier

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Mobile Traffic Identifier. Mobile traffic identifier type.

Upstream description:

Mobile traffic identifier type.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
mobile_identifier {
  # Configure direct properties listed below.
}
```

## Direct properties

- [headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/mobile_sdk_config/mobile_identifier/headers/): complete subsection reference.

## Next pages

- [cloudfront.mobile_sdk_config.mobile_identifier.headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/mobile_sdk_config/mobile_identifier/headers/)
- [cloudfront.mobile_sdk_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/mobile_sdk_config/)
- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/)
