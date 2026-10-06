---
page_title: "virtual_server.port_translation"
subcategory: ""
description: "Specifies, when checked (enabled), that the system translates the port of the virtual server. When cleared (disabled), specifies that the system uses the port without translation. Turning off port translation for a virtual server is useful if you want to use the virtual server to load balance connections to any"
xcsh_docs: {"aliases": ["virtual server port translation"], "body_bytes": 2050, "body_sha256": "sha256:b59407154734b2bf7570611791305f9ae0382aa0b0b92b21aa9aa9bf16aa9580", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:application_profiles:properties:virtual_server:port_translation:port_translation_disable", "xcsh-docs:resources:application_profiles:properties:virtual_server:port_translation:port_translation_enable"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:resources:application_profiles:properties:virtual_server:port_translation", "parent_id": "xcsh-docs:resources:application_profiles:properties:virtual_server", "path": "documentation/resources/application_profiles/properties/virtual_server/port_translation/index.md", "product": "distributed-cloud", "provider_name": "application_profiles", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-3121121231001031-3330300000133002-1031230222111031-0110330122312033-2031202031313323-3102103200300201-2230013301021233-3121002202231030", "registry_path": "docs/guides/resources--application_profiles--reference--group-004.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "virtual_server.port_translation:ConflictingObjectAttributes:port_translation_disable,port_translation_enable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:port_translation:port_translation_disable", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "virtual_server.port_translation:ConflictingObjectAttributes:port_translation_disable,port_translation_enable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:port_translation:port_translation_enable", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["virtual_server", "port_translation"], "schema_version": 1, "sections": [{"aliases": ["virtual server port translation port translation disable"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:port_translation:port_translation_disable", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["virtual_server", "port_translation", "port_translation_disable"], "syntax": "attribute", "type": "object"}, {"aliases": ["virtual server port translation port translation enable"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:port_translation:port_translation_enable", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["virtual_server", "port_translation", "port_translation_enable"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/application_profiles/properties/virtual_server/port_translation/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Specifies, when checked (enabled), that the system translates the port of the virtual server. When cleared (disabled), specifies that the system uses the port without translation. Turning off port translation for a virtual server is useful if you want to use the virtual server to load balance connections to any", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["application_profilesCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# virtual_server.port_translation

Breadcrumbs:

- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/)
- [virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/)
- virtual_server.port_translation

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Specifies, when checked (enabled), that the system translates the port of the virtual server. When
cleared (disabled), specifies that the system uses the port without translation. Turning off port
translation for a virtual server is useful if you want to use the virtual server to load balance
connections to any service. The default is enabled.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("port_translation_disable",
    "port_translation_enable")}
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
  "x-ves-oneof-field-port_translation_choice": "[\"port_translation_disable\",\"port_translation_enable\"]"
}
```

Terraform syntax:

```terraform
port_translation {
  # Configure direct properties listed below.
}
```

## Direct properties

- [port_translation_disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/port_translation/port_translation_disable/): complete subsection reference.

- [port_translation_enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/port_translation/port_translation_enable/): complete subsection reference.
