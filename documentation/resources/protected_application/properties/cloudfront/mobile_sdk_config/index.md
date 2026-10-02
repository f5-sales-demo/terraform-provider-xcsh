---
page_title: "cloudfront.mobile_sdk_config"
subcategory: ""
description: "Mobile SDK configuration."
xcsh_docs: {"aliases": ["cloudfront mobile sdk config"], "body_bytes": 1691, "body_sha256": "sha256:08fe6b1400edbe1c56b01a3782a31a3d1c877cd29dbe7e19420eff86db9ff2e7", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:protected_application:properties:cloudfront:mobile_sdk_config:mobile_identifier"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:resources:protected_application:properties:cloudfront:mobile_sdk_config", "parent_id": "xcsh-docs:resources:protected_application:properties:cloudfront", "path": "documentation/resources/protected_application/properties/cloudfront/mobile_sdk_config/index.md", "product": "distributed-cloud", "provider_name": "protected_application", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3033020132301333-0003231320203312-0230003322013233-2133203312003202-3300000233233110-3001200323232120-3021132321131110-0111230321020021", "registry_path": "docs/guides/resources--protected_application--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["cloudfront", "mobile_sdk_config"], "schema_version": 1, "sections": [{"aliases": ["mobile identifier"], "anchor": "section", "description": "Mobile traffic identifier type.", "document_id": "xcsh-docs:resources:protected_application:properties:cloudfront:mobile_sdk_config:mobile_identifier", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["cloudfront", "mobile_sdk_config", "mobile_identifier"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protected_application/properties/cloudfront/mobile_sdk_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Mobile SDK configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudfront.mobile_sdk_config

Breadcrumbs:

- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/)
- [cloudfront](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/)
- cloudfront.mobile_sdk_config

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Mobile SDK Configuration. Mobile SDK configuration.

Upstream description:

Mobile SDK configuration.

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
mobile_sdk_config {
  # Configure direct properties listed below.
}
```

## Direct properties

- [mobile_identifier](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/mobile_sdk_config/mobile_identifier/): complete subsection reference.

## Next pages

- [cloudfront.mobile_sdk_config.mobile_identifier](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/mobile_sdk_config/mobile_identifier/)
- [cloudfront](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/)
- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/)
