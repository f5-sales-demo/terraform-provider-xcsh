---
page_title: "cloudflare.mobile_sdk_config"
subcategory: ""
description: "Mobile SDK configuration."
xcsh_docs: {"aliases": ["cloudflare mobile sdk config"], "body_bytes": 1064, "body_sha256": "sha256:64084552a8990cabfce885acc7ff8c79ad7c7c4f3634033908e1e2030edb47a4", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:protected_application:properties:cloudflare:mobile_sdk_config:mobile_identifier"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:protected_application:properties:cloudflare:mobile_sdk_config", "parent_id": "xcsh-docs:data-sources:protected_application:properties:cloudflare", "path": "documentation/data-sources/protected_application/properties/cloudflare/mobile_sdk_config/index.md", "product": "distributed-cloud", "provider_name": "protected_application", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-1012330233031112-2222321301033031-3302233131022100-2332032300003202-2320003111223303-0031323032112130-2020330111011333-2103332111101231", "registry_path": "docs/guides/data-sources--protected_application--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["cloudflare", "mobile_sdk_config"], "schema_version": 1, "sections": [{"aliases": ["cloudflare mobile sdk config mobile identifier"], "anchor": "section", "description": "Mobile traffic identifier type.", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudflare:mobile_sdk_config:mobile_identifier", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["cloudflare", "mobile_sdk_config", "mobile_identifier"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protected_application/properties/cloudflare/mobile_sdk_config/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Mobile SDK configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudflare.mobile_sdk_config

Breadcrumbs:

- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/)
- [cloudflare](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudflare/)
- cloudflare.mobile_sdk_config

<a id="section"></a>

Type: `"single"`. Computed.

Mobile SDK Configuration. Mobile SDK configuration.

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

## Direct properties

- [mobile_identifier](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudflare/mobile_sdk_config/mobile_identifier/): complete subsection reference.
