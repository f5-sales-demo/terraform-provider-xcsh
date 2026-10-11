---
page_title: "virtual_server.source_port"
subcategory: ""
description: "Specifies whether the system preserves the source port of the connection. The default is Preserve."
xcsh_docs: {"aliases": ["virtual server source port"], "body_bytes": 2189, "body_sha256": "sha256:e3fe36367dc63c82fa32cd2056f5de6c49b39f0152dca18bb04d83545bbf7271", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:application_profiles:properties:virtual_server:source_port:source_port_change", "xcsh-docs:resources:application_profiles:properties:virtual_server:source_port:source_port_preserve", "xcsh-docs:resources:application_profiles:properties:virtual_server:source_port:source_port_preserve_strict"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:resources:application_profiles:properties:virtual_server:source_port", "parent_id": "xcsh-docs:resources:application_profiles:properties:virtual_server", "path": "documentation/resources/application_profiles/properties/virtual_server/source_port/index.md", "product": "distributed-cloud", "provider_name": "application_profiles", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "resources", "registry_anchor": "canonical-1232230330103113-1011323112020110-2112322333333013-3302112210103221-2002022032123220-0023200230233130-2102030102320103-0013123112003122", "registry_path": "docs/guides/resources--application_profiles--reference--group-004.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "virtual_server.source_port:ConflictingObjectAttributes:source_port_change,source_port_preserve", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:source_port:source_port_change", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "virtual_server.source_port:ConflictingObjectAttributes:source_port_change,source_port_preserve_strict", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:source_port:source_port_change", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "virtual_server.source_port:ConflictingObjectAttributes:source_port_change,source_port_preserve", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:source_port:source_port_preserve", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "virtual_server.source_port:ConflictingObjectAttributes:source_port_preserve,source_port_preserve_strict", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:source_port:source_port_preserve", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "virtual_server.source_port:ConflictingObjectAttributes:source_port_change,source_port_preserve_strict", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:source_port:source_port_preserve_strict", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "virtual_server.source_port:ConflictingObjectAttributes:source_port_preserve,source_port_preserve_strict", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:source_port:source_port_preserve_strict", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["virtual_server", "source_port"], "schema_version": 1, "sections": [{"aliases": ["virtual server source port source port change"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:source_port:source_port_change", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["virtual_server", "source_port", "source_port_change"], "syntax": "attribute", "type": "object"}, {"aliases": ["virtual server source port source port preserve"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:source_port:source_port_preserve", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["virtual_server", "source_port", "source_port_preserve"], "syntax": "attribute", "type": "object"}, {"aliases": ["virtual server source port source port preserve strict"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:source_port:source_port_preserve_strict", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["virtual_server", "source_port", "source_port_preserve_strict"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/application_profiles/properties/virtual_server/source_port/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Specifies whether the system preserves the source port of the connection. The default is Preserve.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["application_profilesCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# virtual_server.source_port

Breadcrumbs:

- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/)
- [virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/)
- virtual_server.source_port

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Specifies whether the system preserves the source port of the connection. The default is Preserve.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("source_port_change",
    "source_port_preserve"),
  validators.ConflictingObjectAttributes("source_port_change",
    "source_port_preserve_strict"),
  validators.ConflictingObjectAttributes("source_port_preserve",
    "source_port_preserve_strict")}
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
  "x-ves-oneof-field-source_port_choice": "[\"source_port_change\",\"source_port_preserve\",\"source_port_preserve_strict\"]"
}
```

Terraform syntax:

```terraform
source_port {
  # Configure direct properties listed below.
}
```

## Direct properties

- [source_port_change](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/source_port/source_port_change/): complete subsection reference.

- [source_port_preserve](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/source_port/source_port_preserve/): complete subsection reference.

- [source_port_preserve_strict](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/source_port/source_port_preserve_strict/): complete subsection reference.
