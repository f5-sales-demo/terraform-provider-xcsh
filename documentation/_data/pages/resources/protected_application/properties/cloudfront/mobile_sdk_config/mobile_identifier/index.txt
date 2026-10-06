---
page_title: "cloudfront.mobile_sdk_config.mobile_identifier"
subcategory: ""
description: "Mobile traffic identifier type."
xcsh_docs: {"aliases": ["cloudfront mobile sdk config mobile identifier"], "body_bytes": 1382, "body_sha256": "sha256:9cec3a6d9587042d2a254591bca5a81167a67516b37ef5643243f03f1abb308d", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:protected_application:properties:cloudfront:mobile_sdk_config:mobile_identifier:headers"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:resources:protected_application:properties:cloudfront:mobile_sdk_config:mobile_identifier", "parent_id": "xcsh-docs:resources:protected_application:properties:cloudfront:mobile_sdk_config", "path": "documentation/resources/protected_application/properties/cloudfront/mobile_sdk_config/mobile_identifier/index.md", "product": "distributed-cloud", "provider_name": "protected_application", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-0213223320221112-2122101302230222-1313301132220012-2212300132233123-3031030123331203-3003133332121301-1031000123123233-1033210032003123", "registry_path": "docs/guides/resources--protected_application--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["cloudfront", "mobile_sdk_config", "mobile_identifier"], "schema_version": 1, "sections": [{"aliases": ["cloudfront mobile sdk config mobile identifier headers"], "anchor": "section", "description": "A list of headers that can be used to identify mobile traffic.", "document_id": "xcsh-docs:resources:protected_application:properties:cloudfront:mobile_sdk_config:mobile_identifier:headers", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-cloudfront--mobile_sdk_config--mobile_identifier--headers--exact", "enforcement": "provider-schema", "group": "cloudfront.mobile_sdk_config.mobile_identifier.headers:ConflictingListObjectAttributes:exact,regex", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudfront:mobile_sdk_config:mobile_identifier:headers", "type": "conflicts"}, {"anchor": "schema-cloudfront--mobile_sdk_config--mobile_identifier--headers--regex", "enforcement": "provider-schema", "group": "cloudfront.mobile_sdk_config.mobile_identifier.headers:ConflictingListObjectAttributes:exact,regex", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudfront:mobile_sdk_config:mobile_identifier:headers", "type": "conflicts"}, {"anchor": "schema-cloudfront--mobile_sdk_config--mobile_identifier--headers--name", "enforcement": "provider-schema", "group": "cloudfront.mobile_sdk_config.mobile_identifier.headers:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudfront:mobile_sdk_config:mobile_identifier:headers", "type": "requires"}], "schema_path": ["cloudfront", "mobile_sdk_config", "mobile_identifier", "headers"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protected_application/properties/cloudfront/mobile_sdk_config/mobile_identifier/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Mobile traffic identifier type.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
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
