---
page_title: "cloudflare.mobile_sdk_config.mobile_identifier"
subcategory: ""
description: "Mobile traffic identifier type."
xcsh_docs: {"aliases": ["cloudflare mobile sdk config mobile identifier"], "body_bytes": 1382, "body_sha256": "sha256:849a64a905ccaaad6f38f01821398a352e2cbb107d314b87efbd4144dbcfd75d", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:protected_application:properties:cloudflare:mobile_sdk_config:mobile_identifier:headers"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:resources:protected_application:properties:cloudflare:mobile_sdk_config:mobile_identifier", "parent_id": "xcsh-docs:resources:protected_application:properties:cloudflare:mobile_sdk_config", "path": "documentation/resources/protected_application/properties/cloudflare/mobile_sdk_config/mobile_identifier/index.md", "product": "distributed-cloud", "provider_name": "protected_application", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-1110022210110133-0302122033301331-1232032102212313-1122232331103031-3032202212222021-2000030110212110-0020331021132012-0111121110230301", "registry_path": "docs/guides/resources--protected_application--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["cloudflare", "mobile_sdk_config", "mobile_identifier"], "schema_version": 1, "sections": [{"aliases": ["cloudflare mobile sdk config mobile identifier headers"], "anchor": "section", "description": "A list of headers that can be used to identify mobile traffic.", "document_id": "xcsh-docs:resources:protected_application:properties:cloudflare:mobile_sdk_config:mobile_identifier:headers", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-cloudflare--mobile_sdk_config--mobile_identifier--headers--exact", "enforcement": "provider-schema", "group": "cloudflare.mobile_sdk_config.mobile_identifier.headers:ConflictingListObjectAttributes:exact,regex", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:mobile_sdk_config:mobile_identifier:headers", "type": "conflicts"}, {"anchor": "schema-cloudflare--mobile_sdk_config--mobile_identifier--headers--regex", "enforcement": "provider-schema", "group": "cloudflare.mobile_sdk_config.mobile_identifier.headers:ConflictingListObjectAttributes:exact,regex", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:mobile_sdk_config:mobile_identifier:headers", "type": "conflicts"}, {"anchor": "schema-cloudflare--mobile_sdk_config--mobile_identifier--headers--name", "enforcement": "provider-schema", "group": "cloudflare.mobile_sdk_config.mobile_identifier.headers:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:mobile_sdk_config:mobile_identifier:headers", "type": "requires"}], "schema_path": ["cloudflare", "mobile_sdk_config", "mobile_identifier", "headers"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protected_application/properties/cloudflare/mobile_sdk_config/mobile_identifier/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Mobile traffic identifier type.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudflare.mobile_sdk_config.mobile_identifier

Breadcrumbs:

- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/)
- [cloudflare](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/)
- [cloudflare.mobile_sdk_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/mobile_sdk_config/)
- cloudflare.mobile_sdk_config.mobile_identifier

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

- [headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/mobile_sdk_config/mobile_identifier/headers/): complete subsection reference.
