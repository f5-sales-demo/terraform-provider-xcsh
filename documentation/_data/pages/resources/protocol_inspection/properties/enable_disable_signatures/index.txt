---
page_title: "enable_disable_signatures"
subcategory: ""
description: "Enable Disable Signature Choice."
xcsh_docs: {"aliases": ["enable disable signatures"], "body_bytes": 1606, "body_sha256": "sha256:78f1ccaa78ebbf61875657e37d4d41b4042a9bf0805fb4ce11ad4daaafc66e5d", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:protocol_inspection:properties:enable_disable_signatures:disable_signature", "xcsh-docs:resources:protocol_inspection:properties:enable_disable_signatures:enable_signature"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:protocol_inspection:collection", "completeness": "complete", "id": "xcsh-docs:resources:protocol_inspection:properties:enable_disable_signatures", "parent_id": "xcsh-docs:resources:protocol_inspection:reference", "path": "documentation/resources/protocol_inspection/properties/enable_disable_signatures/index.md", "product": "distributed-cloud", "provider_name": "protocol_inspection", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-1103202120331001-3021223223230333-3311111113320300-0113233330323001-1310220331310130-3111230130303221-3033121302233302-0220122231222233", "registry_path": "docs/guides/resources--protocol_inspection--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "enable_disable_signatures:ConflictingObjectAttributes:disable_signature,enable_signature", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protocol_inspection:properties:enable_disable_signatures:disable_signature", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "enable_disable_signatures:ConflictingObjectAttributes:disable_signature,enable_signature", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protocol_inspection:properties:enable_disable_signatures:enable_signature", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["enable_disable_signatures"], "schema_version": 1, "sections": [{"aliases": ["enable disable signatures disable signature"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:protocol_inspection:properties:enable_disable_signatures:disable_signature", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_disable_signatures", "disable_signature"], "syntax": "attribute", "type": "object"}, {"aliases": ["enable disable signatures enable signature"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:protocol_inspection:properties:enable_disable_signatures:enable_signature", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_disable_signatures", "enable_signature"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protocol_inspection/properties/enable_disable_signatures/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Enable Disable Signature Choice.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["protocol_inspectionCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_disable_signatures

Breadcrumbs:

- [xcsh_protocol_inspection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_inspection/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_inspection/properties/)
- enable_disable_signatures

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for enable disable signatures.

Additional upstream details:

Enable Disable Signature Choice.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_signature",
    "enable_signature")}
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
  "x-ves-oneof-field-signature_choice": "[\"disable_signature\",\"enable_signature\"]"
}
```

Terraform syntax:

```terraform
enable_disable_signatures {
  # Configure direct properties listed below.
}
```

## Direct properties

- [disable_signature](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_inspection/properties/enable_disable_signatures/disable_signature/): complete subsection reference.

- [enable_signature](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_inspection/properties/enable_disable_signatures/enable_signature/): complete subsection reference.
