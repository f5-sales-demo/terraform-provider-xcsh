---
page_title: "virtual_server.immediate_action_on_service_down"
subcategory: ""
description: "Specifies the immediate action the BIG-IP system should respond with upon the receipt of the initial client's SYN packet, if the availability status of the virtual server is Offline or Unavailable. This is supported for the virtual server of Standard type and TCP protocol. The default is None. None: Specifies that the"
xcsh_docs: {"aliases": ["virtual server immediate action on service down"], "body_bytes": 4501, "body_sha256": "sha256:c85ce23cabd6151f97e5fd49560102a2bc286fd1886199e76d3ee20ffc6157d2", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:application_profiles:properties:virtual_server:immediate_action_on_service_down:immediate_action_on_service_down_drop", "xcsh-docs:resources:application_profiles:properties:virtual_server:immediate_action_on_service_down:immediate_action_on_service_down_none", "xcsh-docs:resources:application_profiles:properties:virtual_server:immediate_action_on_service_down:immediate_action_on_service_down_reset"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:resources:application_profiles:properties:virtual_server:immediate_action_on_service_down", "parent_id": "xcsh-docs:resources:application_profiles:properties:virtual_server", "path": "documentation/resources/application_profiles/properties/virtual_server/immediate_action_on_service_down/index.md", "product": "distributed-cloud", "provider_name": "application_profiles", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "resources", "registry_anchor": "canonical-2120313011001333-0332132133112211-2032113202213200-1133333320210120-0120003121121103-1110233000210030-2313321130012321-0301012230001201", "registry_path": "docs/guides/resources--application_profiles--reference--group-003.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "virtual_server.immediate_action_on_service_down:ConflictingObjectAttributes:immediate_action_on_service_down_drop,immediate_action_on_service_down_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:immediate_action_on_service_down:immediate_action_on_service_down_drop", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "virtual_server.immediate_action_on_service_down:ConflictingObjectAttributes:immediate_action_on_service_down_drop,immediate_action_on_service_down_reset", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:immediate_action_on_service_down:immediate_action_on_service_down_drop", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "virtual_server.immediate_action_on_service_down:ConflictingObjectAttributes:immediate_action_on_service_down_drop,immediate_action_on_service_down_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:immediate_action_on_service_down:immediate_action_on_service_down_none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "virtual_server.immediate_action_on_service_down:ConflictingObjectAttributes:immediate_action_on_service_down_none,immediate_action_on_service_down_reset", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:immediate_action_on_service_down:immediate_action_on_service_down_none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "virtual_server.immediate_action_on_service_down:ConflictingObjectAttributes:immediate_action_on_service_down_drop,immediate_action_on_service_down_reset", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:immediate_action_on_service_down:immediate_action_on_service_down_reset", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "virtual_server.immediate_action_on_service_down:ConflictingObjectAttributes:immediate_action_on_service_down_none,immediate_action_on_service_down_reset", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:immediate_action_on_service_down:immediate_action_on_service_down_reset", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["virtual_server", "immediate_action_on_service_down"], "schema_version": 1, "sections": [{"aliases": ["immediate action on service down drop"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:immediate_action_on_service_down:immediate_action_on_service_down_drop", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["virtual_server", "immediate_action_on_service_down", "immediate_action_on_service_down_drop"], "syntax": "attribute", "type": "object"}, {"aliases": ["immediate action on service down none"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:immediate_action_on_service_down:immediate_action_on_service_down_none", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["virtual_server", "immediate_action_on_service_down", "immediate_action_on_service_down_none"], "syntax": "attribute", "type": "object"}, {"aliases": ["immediate action on service down reset"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:immediate_action_on_service_down:immediate_action_on_service_down_reset", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["virtual_server", "immediate_action_on_service_down", "immediate_action_on_service_down_reset"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/application_profiles/properties/virtual_server/immediate_action_on_service_down/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Specifies the immediate action the BIG-IP system should respond with upon the receipt of the initial client's SYN packet, if the availability status of the virtual server is Offline or Unavailable. This is supported for the virtual server of Standard type and TCP protocol. The default is None. None: Specifies that the", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["application_profilesCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

Upstream description:

Specifies the immediate action the BIG-IP system should respond with upon the receipt of the initial
client's SYN packet, if the availability status of the virtual server is Offline or Unavailable.
This is supported for the virtual server of Standard type and TCP protocol. The default is None.
None: Specifies that the system takes no immediate action if the virtual server is reported Offline
or Unavailable. Reset: Specifies that the system resets the connections when the virtual server is
reported Offline or Unavailable. Drop: Specifies that the system drops the connections when the
virtual server is reported Offline or Unavailable.

Provider validators and defaults (from schema source):

```go
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

## Next pages

- [virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_drop](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/immediate_action_on_service_down/immediate_action_on_service_down_drop/)
- [virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/immediate_action_on_service_down/immediate_action_on_service_down_none/)
- [virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_reset](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/immediate_action_on_service_down/immediate_action_on_service_down_reset/)
- [virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/)
- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/)
