---
page_title: "virtual_server.virtual_server_state"
subcategory: ""
description: "Displays the current state on the object."
xcsh_docs: {"aliases": ["virtual server virtual server state"], "body_bytes": 2045, "body_sha256": "sha256:d062ae41c8c42ea7c38e97530f511c4e9e474fdbbbd80863f528b6898000c5de", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:application_profiles:properties:virtual_server:virtual_server_state:state_disabled", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:virtual_server_state:state_enabled"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:virtual_server_state", "parent_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server", "path": "documentation/data-sources/application_profiles/properties/virtual_server/virtual_server_state/index.md", "product": "distributed-cloud", "provider_name": "application_profiles", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0300001203202301-3231100101123312-2010220203023303-2030320120111312-3201010231210333-1312201331313222-1321111320032321-3200132220130130", "registry_path": "docs/guides/data-sources--application_profiles--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["virtual_server", "virtual_server_state"], "schema_version": 1, "sections": [{"aliases": ["state disabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:virtual_server_state:state_disabled", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["virtual_server", "virtual_server_state", "state_disabled"], "syntax": "attribute", "type": "object"}, {"aliases": ["state enabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:virtual_server_state:state_enabled", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["virtual_server", "virtual_server_state", "state_enabled"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/application_profiles/properties/virtual_server/virtual_server_state/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Displays the current state on the object.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["application_profilesCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# virtual_server.virtual_server_state

Breadcrumbs:

- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/)
- [virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/)
- virtual_server.virtual_server_state

<a id="section"></a>

Type: `"single"`. Computed.

Displays the current state on the object.

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

## Direct properties

- [state_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/virtual_server_state/state_disabled/): complete subsection reference.

- [state_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/virtual_server_state/state_enabled/): complete subsection reference.

## Next pages

- [virtual_server.virtual_server_state.state_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/virtual_server_state/state_disabled/)
- [virtual_server.virtual_server_state.state_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/virtual_server_state/state_enabled/)
- [virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/)
- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/)
