---
page_title: "virtual_server.source_port"
subcategory: ""
description: "Specifies whether the system preserves the source port of the connection. The default is Preserve."
xcsh_docs: {"aliases": ["virtual server source port"], "body_bytes": 3038, "body_sha256": "sha256:1bd976b1f371471ff26613b826fbd0208e60b6f78784afe074293ed0edae5453", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:application_profiles:properties:virtual_server:source_port:source_port_change", "xcsh-docs:resources:application_profiles:properties:virtual_server:source_port:source_port_preserve", "xcsh-docs:resources:application_profiles:properties:virtual_server:source_port:source_port_preserve_strict"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:resources:application_profiles:properties:virtual_server:source_port", "parent_id": "xcsh-docs:resources:application_profiles:properties:virtual_server", "path": "documentation/resources/application_profiles/properties/virtual_server/source_port/index.md", "product": "distributed-cloud", "provider_name": "application_profiles", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1232230330103113-1011323112020110-2112322333333013-3302112210103221-2002022032123220-0023200230233130-2102030102320103-0013123112003122", "registry_path": "docs/guides/resources--application_profiles--reference--group-003.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "virtual_server.source_port:ConflictingObjectAttributes:source_port_change,source_port_preserve", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:source_port:source_port_change", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "virtual_server.source_port:ConflictingObjectAttributes:source_port_change,source_port_preserve_strict", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:source_port:source_port_change", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "virtual_server.source_port:ConflictingObjectAttributes:source_port_change,source_port_preserve", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:source_port:source_port_preserve", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "virtual_server.source_port:ConflictingObjectAttributes:source_port_preserve,source_port_preserve_strict", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:source_port:source_port_preserve", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "virtual_server.source_port:ConflictingObjectAttributes:source_port_change,source_port_preserve_strict", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:source_port:source_port_preserve_strict", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "virtual_server.source_port:ConflictingObjectAttributes:source_port_preserve,source_port_preserve_strict", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:source_port:source_port_preserve_strict", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["virtual_server", "source_port"], "schema_version": 1, "sections": [{"aliases": ["source port change"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:source_port:source_port_change", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["virtual_server", "source_port", "source_port_change"], "syntax": "attribute", "type": "object"}, {"aliases": ["source port preserve"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:source_port:source_port_preserve", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["virtual_server", "source_port", "source_port_preserve"], "syntax": "attribute", "type": "object"}, {"aliases": ["source port preserve strict"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:source_port:source_port_preserve_strict", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["virtual_server", "source_port", "source_port_preserve_strict"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/application_profiles/properties/virtual_server/source_port/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Specifies whether the system preserves the source port of the connection. The default is Preserve.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["application_profilesCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

## Next pages

- [virtual_server.source_port.source_port_change](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/source_port/source_port_change/)
- [virtual_server.source_port.source_port_preserve](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/source_port/source_port_preserve/)
- [virtual_server.source_port.source_port_preserve_strict](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/source_port/source_port_preserve_strict/)
- [virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/)
- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/)
