---
page_title: "virtual_server.immediate_action_on_service_down"
subcategory: ""
description: "Specifies the immediate action the BIG-IP system should respond with upon the receipt of the initial client's SYN packet, if the availability status of the virtual server is Offline or Unavailable. This is supported for the virtual server of Standard type and TCP protocol. The default is None. None: Specifies that the"
xcsh_docs: {"aliases": ["virtual server immediate action on service down"], "body_bytes": 3113, "body_sha256": "sha256:bdf36d6a557a6acd7274968f734f15a0955912ceb46fe814909b00d9dcda11a0", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:application_profiles:properties:virtual_server:immediate_action_on_service_down:immediate_action_on_service_down_drop", "xcsh-docs:resources:application_profiles:properties:virtual_server:immediate_action_on_service_down:immediate_action_on_service_down_none", "xcsh-docs:resources:application_profiles:properties:virtual_server:immediate_action_on_service_down:immediate_action_on_service_down_reset"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:resources:application_profiles:properties:virtual_server:immediate_action_on_service_down", "parent_id": "xcsh-docs:resources:application_profiles:properties:virtual_server", "path": "documentation/resources/application_profiles/properties/virtual_server/immediate_action_on_service_down/index.md", "product": "distributed-cloud", "provider_name": "application_profiles", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "resources", "registry_anchor": "canonical-2120313011001333-0332132133112211-2032113202213200-1133333320210120-0120003121121103-1110233000210030-2313321130012321-0301012230001201", "registry_path": "docs/guides/resources--application_profiles--reference--group-004.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "virtual_server.immediate_action_on_service_down:ConflictingObjectAttributes:immediate_action_on_service_down_drop,immediate_action_on_service_down_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:immediate_action_on_service_down:immediate_action_on_service_down_drop", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "virtual_server.immediate_action_on_service_down:ConflictingObjectAttributes:immediate_action_on_service_down_drop,immediate_action_on_service_down_reset", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:immediate_action_on_service_down:immediate_action_on_service_down_drop", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "virtual_server.immediate_action_on_service_down:ConflictingObjectAttributes:immediate_action_on_service_down_drop,immediate_action_on_service_down_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:immediate_action_on_service_down:immediate_action_on_service_down_none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "virtual_server.immediate_action_on_service_down:ConflictingObjectAttributes:immediate_action_on_service_down_none,immediate_action_on_service_down_reset", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:immediate_action_on_service_down:immediate_action_on_service_down_none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "virtual_server.immediate_action_on_service_down:ConflictingObjectAttributes:immediate_action_on_service_down_drop,immediate_action_on_service_down_reset", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:immediate_action_on_service_down:immediate_action_on_service_down_reset", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "virtual_server.immediate_action_on_service_down:ConflictingObjectAttributes:immediate_action_on_service_down_none,immediate_action_on_service_down_reset", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:immediate_action_on_service_down:immediate_action_on_service_down_reset", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["virtual_server", "immediate_action_on_service_down"], "schema_version": 1, "sections": [{"aliases": ["virtual server immediate action on service down immediate action on service down drop"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:immediate_action_on_service_down:immediate_action_on_service_down_drop", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["virtual_server", "immediate_action_on_service_down", "immediate_action_on_service_down_drop"], "syntax": "attribute", "type": "object"}, {"aliases": ["virtual server immediate action on service down immediate action on service down none"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:immediate_action_on_service_down:immediate_action_on_service_down_none", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["virtual_server", "immediate_action_on_service_down", "immediate_action_on_service_down_none"], "syntax": "attribute", "type": "object"}, {"aliases": ["virtual server immediate action on service down immediate action on service down reset"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:immediate_action_on_service_down:immediate_action_on_service_down_reset", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["virtual_server", "immediate_action_on_service_down", "immediate_action_on_service_down_reset"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/application_profiles/properties/virtual_server/immediate_action_on_service_down/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Specifies the immediate action the BIG-IP system should respond with upon the receipt of the initial client's SYN packet, if the availability status of the virtual server is Offline or Unavailable. This is supported for the virtual server of Standard type and TCP protocol. The default is None. None: Specifies that the", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["application_profilesCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# virtual_server.immediate_action_on_service_down

Breadcrumbs:

- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/)
- [virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/)
- virtual_server.immediate_action_on_service_down

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Specifies the immediate action the BIG-IP system should respond with upon the receipt of the initial
client's SYN packet, if the availability status of the virtual server is Offline or Unavailable.
This is supported for the virtual server of Standard type and TCP protocol. The default is None.
None: Specifies that the system takes no immediate action if the virtual server is reported Offline
or Unavailable. Reset: Specifies that the system resets the connections when the virtual server is
reported Offline or Unavailable. Drop: Specifies that the system drops the connections when the
virtual server is reported Offline or Unavailable.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("immediate_action_on_service_down_drop",
    "immediate_action_on_service_down_none"),
  validators.ConflictingObjectAttributes("immediate_action_on_service_down_drop",
    "immediate_action_on_service_down_reset"),
  validators.ConflictingObjectAttributes("immediate_action_on_service_down_none",
    "immediate_action_on_service_down_reset")}
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
  "x-ves-oneof-field-immediate_action_on_service_down_choice": "[\"immediate_action_on_service_down_drop\",\"immediate_action_on_service_down_none\",\"immediate_action_on_service_down_reset\"]"
}
```

Terraform syntax:

```terraform
immediate_action_on_service_down {
  # Configure direct properties listed below.
}
```

## Direct properties

- [immediate_action_on_service_down_drop](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/immediate_action_on_service_down/immediate_action_on_service_down_drop/): complete subsection reference.

- [immediate_action_on_service_down_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/immediate_action_on_service_down/immediate_action_on_service_down_none/): complete subsection reference.

- [immediate_action_on_service_down_reset](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/immediate_action_on_service_down/immediate_action_on_service_down_reset/): complete subsection reference.
