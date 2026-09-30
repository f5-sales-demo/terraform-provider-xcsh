---
page_title: "virtual_server.immediate_action_on_service_down"
subcategory: ""
description: "virtual_server.immediate_action_on_service_down for xcsh_application_profiles."
xcsh_docs: {"aliases": [], "body_bytes": 3851, "body_sha256": "sha256:b63f3bb7d00ab6e23f386f0d365d8ff14651fb0636d1558933926948f890e44f", "canonical_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:immediate_action_on_service_down", "child_ids": ["xcsh-docs:resources:application_profiles:properties:virtual_server:immediate_action_on_service_down:immediate_action_on_service_down_drop", "xcsh-docs:resources:application_profiles:properties:virtual_server:immediate_action_on_service_down:immediate_action_on_service_down_none", "xcsh-docs:resources:application_profiles:properties:virtual_server:immediate_action_on_service_down:immediate_action_on_service_down_reset"], "collection_id": "xcsh-docs:resources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:resources:application_profiles:properties:virtual_server:immediate_action_on_service_down", "parent_id": "xcsh-docs:resources:application_profiles:properties:virtual_server", "path": "docs/guides/resources--application_profiles--properties--virtual_server--immediate_action_on_service_down.md", "provider_name": "application_profiles", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["virtual_server", "immediate_action_on_service_down"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/application_profiles/properties/virtual_server/immediate_action_on_service_down/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "virtual_server.immediate_action_on_service_down for xcsh_application_profiles.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["application_profilesCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# virtual_server.immediate_action_on_service_down

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md)
- [Property reference](resources--application_profiles--reference.md)
- [virtual_server](resources--application_profiles--properties--virtual_server.md)
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

- [immediate_action_on_service_down_drop](resources--application_profiles--properties--virtual_server--immediate_action_on_service_down--immediate_action_on_service_down_drop.md): complete subsection reference.

- [immediate_action_on_service_down_none](resources--application_profiles--properties--virtual_server--immediate_action_on_service_down--immediate_action_on_service_down_none.md): complete subsection reference.

- [immediate_action_on_service_down_reset](resources--application_profiles--properties--virtual_server--immediate_action_on_service_down--immediate_action_on_service_down_reset.md): complete subsection reference.

## Next pages

- [virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_drop](resources--application_profiles--properties--virtual_server--immediate_action_on_service_down--immediate_action_on_service_down_drop.md)
- [virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_none](resources--application_profiles--properties--virtual_server--immediate_action_on_service_down--immediate_action_on_service_down_none.md)
- [virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_reset](resources--application_profiles--properties--virtual_server--immediate_action_on_service_down--immediate_action_on_service_down_reset.md)
- [virtual_server](resources--application_profiles--properties--virtual_server.md)
- [xcsh_application_profiles](../resources/application_profiles.md)
