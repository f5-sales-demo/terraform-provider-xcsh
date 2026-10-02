---
page_title: "virtual_server.virtual_server_state"
subcategory: ""
description: "Displays the current state on the object."
xcsh_docs: {"aliases": ["virtual server virtual server state"], "body_bytes": 2324, "body_sha256": "sha256:d8cebd8de5f7ea2e06efb183bae6a031971fdce7b1de31f28a3b5177af472065", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:application_profiles:properties:virtual_server:virtual_server_state:state_disabled", "xcsh-docs:resources:application_profiles:properties:virtual_server:virtual_server_state:state_enabled"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:resources:application_profiles:properties:virtual_server:virtual_server_state", "parent_id": "xcsh-docs:resources:application_profiles:properties:virtual_server", "path": "documentation/resources/application_profiles/properties/virtual_server/virtual_server_state/index.md", "product": "distributed-cloud", "provider_name": "application_profiles", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0323021232302203-3330212203221312-0021222111110121-2230020130302203-1311230313200231-3310002022201103-2331300211333032-3021322223300011", "registry_path": "docs/guides/resources--application_profiles--reference--group-004.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "virtual_server.virtual_server_state:ConflictingObjectAttributes:state_disabled,state_enabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:virtual_server_state:state_disabled", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "virtual_server.virtual_server_state:ConflictingObjectAttributes:state_disabled,state_enabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:virtual_server_state:state_enabled", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["virtual_server", "virtual_server_state"], "schema_version": 1, "sections": [{"aliases": ["state disabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:virtual_server_state:state_disabled", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["virtual_server", "virtual_server_state", "state_disabled"], "syntax": "attribute", "type": "object"}, {"aliases": ["state enabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:virtual_server_state:state_enabled", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["virtual_server", "virtual_server_state", "state_enabled"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/application_profiles/properties/virtual_server/virtual_server_state/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Displays the current state on the object.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["application_profilesCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

## Next pages

- [virtual_server.virtual_server_state.state_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/virtual_server_state/state_disabled/)
- [virtual_server.virtual_server_state.state_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/virtual_server_state/state_enabled/)
- [virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/)
- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/)
