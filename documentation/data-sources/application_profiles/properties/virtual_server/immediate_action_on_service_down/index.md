---
page_title: "virtual_server.immediate_action_on_service_down"
subcategory: ""
description: "Specifies the immediate action the BIG-IP system should respond with upon the receipt of the initial client's SYN packet, if the availability status of the virtual server is Offline or Unavailable. This is supported for the virtual server of Standard type and TCP protocol. The default is None. None: Specifies that the"
xcsh_docs: {"aliases": ["virtual server immediate action on service down"], "body_bytes": 3911, "body_sha256": "sha256:2b81dd6baf73b7297db15f0f8078d8ee0d44f05801fed00e9bef33535445591a", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:application_profiles:properties:virtual_server:immediate_action_on_service_down:immediate_action_on_service_down_drop", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:immediate_action_on_service_down:immediate_action_on_service_down_none", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:immediate_action_on_service_down:immediate_action_on_service_down_reset"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:immediate_action_on_service_down", "parent_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server", "path": "documentation/data-sources/application_profiles/properties/virtual_server/immediate_action_on_service_down/index.md", "product": "distributed-cloud", "provider_name": "application_profiles", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2300123112133033-1303100333113302-3231112332130213-3210001021212130-0303103023020132-3211232333100212-1030102323131320-3023101112213320", "registry_path": "docs/guides/data-sources--application_profiles--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["virtual_server", "immediate_action_on_service_down"], "schema_version": 1, "sections": [{"aliases": ["immediate action on service down drop"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:immediate_action_on_service_down:immediate_action_on_service_down_drop", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["virtual_server", "immediate_action_on_service_down", "immediate_action_on_service_down_drop"], "syntax": "attribute", "type": "object"}, {"aliases": ["immediate action on service down none"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:immediate_action_on_service_down:immediate_action_on_service_down_none", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["virtual_server", "immediate_action_on_service_down", "immediate_action_on_service_down_none"], "syntax": "attribute", "type": "object"}, {"aliases": ["immediate action on service down reset"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:immediate_action_on_service_down:immediate_action_on_service_down_reset", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["virtual_server", "immediate_action_on_service_down", "immediate_action_on_service_down_reset"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/application_profiles/properties/virtual_server/immediate_action_on_service_down/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Specifies the immediate action the BIG-IP system should respond with upon the receipt of the initial client's SYN packet, if the availability status of the virtual server is Offline or Unavailable. This is supported for the virtual server of Standard type and TCP protocol. The default is None. None: Specifies that the", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["application_profilesCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# virtual_server.immediate_action_on_service_down

Breadcrumbs:

- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/)
- [virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/)
- virtual_server.immediate_action_on_service_down

<a id="section"></a>

Type: `"single"`. Computed.

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

## Direct properties

- [immediate_action_on_service_down_drop](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/immediate_action_on_service_down/immediate_action_on_service_down_drop/): complete subsection reference.

- [immediate_action_on_service_down_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/immediate_action_on_service_down/immediate_action_on_service_down_none/): complete subsection reference.

- [immediate_action_on_service_down_reset](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/immediate_action_on_service_down/immediate_action_on_service_down_reset/): complete subsection reference.

## Next pages

- [virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_drop](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/immediate_action_on_service_down/immediate_action_on_service_down_drop/)
- [virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/immediate_action_on_service_down/immediate_action_on_service_down_none/)
- [virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_reset](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/immediate_action_on_service_down/immediate_action_on_service_down_reset/)
- [virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/)
- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/)
