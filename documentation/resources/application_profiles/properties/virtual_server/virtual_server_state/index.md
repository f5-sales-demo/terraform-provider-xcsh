---
page_title: "virtual_server.virtual_server_state"
subcategory: ""
description: "Displays the current state on the object."
xcsh_docs: {"aliases": ["virtual server virtual server state"], "body_bytes": 1675, "body_sha256": "sha256:4f7da02d51c0c1579a2fd51bc0fcb900f52b65617174a27ab0b8edc91f9419ce", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:application_profiles:properties:virtual_server:virtual_server_state:state_disabled", "xcsh-docs:resources:application_profiles:properties:virtual_server:virtual_server_state:state_enabled"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:resources:application_profiles:properties:virtual_server:virtual_server_state", "parent_id": "xcsh-docs:resources:application_profiles:properties:virtual_server", "path": "documentation/resources/application_profiles/properties/virtual_server/virtual_server_state/index.md", "product": "distributed-cloud", "provider_name": "application_profiles", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-0323021232302203-3330212203221312-0021222111110121-2230020130302203-1311230313200231-3310002022201103-2331300211333032-3021322223300011", "registry_path": "docs/guides/resources--application_profiles--reference--group-004.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "virtual_server.virtual_server_state:ConflictingObjectAttributes:state_disabled,state_enabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:virtual_server_state:state_disabled", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "virtual_server.virtual_server_state:ConflictingObjectAttributes:state_disabled,state_enabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:virtual_server_state:state_enabled", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["virtual_server", "virtual_server_state"], "schema_version": 1, "sections": [{"aliases": ["virtual server virtual server state state disabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:virtual_server_state:state_disabled", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["virtual_server", "virtual_server_state", "state_disabled"], "syntax": "attribute", "type": "object"}, {"aliases": ["virtual server virtual server state state enabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:virtual_server_state:state_enabled", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["virtual_server", "virtual_server_state", "state_enabled"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/application_profiles/properties/virtual_server/virtual_server_state/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Displays the current state on the object.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["application_profilesCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# virtual_server.virtual_server_state

Breadcrumbs:

- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/)
- [virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/)
- virtual_server.virtual_server_state

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Displays the current state on the object.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("state_disabled",
    "state_enabled")}
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
  "x-ves-oneof-field-state_choice": "[\"state_disabled\",\"state_enabled\"]"
}
```

Terraform syntax:

```terraform
virtual_server_state {
  # Configure direct properties listed below.
}
```

## Direct properties

- [state_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/virtual_server_state/state_disabled/): complete subsection reference.

- [state_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/virtual_server_state/state_enabled/): complete subsection reference.
