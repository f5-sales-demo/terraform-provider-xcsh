---
page_title: "items.object.status"
subcategory: ""
description: "Status Type. Most recent observer status of object."
xcsh_docs: {"aliases": ["items object status"], "body_bytes": 6060, "body_sha256": "sha256:30dd9dbee66daaabb46f84dafc408b111eeac0d71389e32d1d7edf4b83b9fef4", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:status:object_status"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations_by_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:status", "parent_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object", "path": "documentation/data-sources/site_registrations_by_site/properties/items/object/status/index.md", "product": "distributed-cloud", "provider_name": "site_registrations_by_site", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-2022122012210233-2130330101221002-0312333301332132-2200233202230322-2122202001131323-1223332221102330-1330312033322030-3220020332323113", "registry_path": "docs/guides/data-sources--site_registrations_by_site--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["items", "object", "status"], "schema_version": 1, "sections": [{"aliases": ["items object status current state"], "anchor": "schema-items--object--status--current_state", "description": "Defines states for registration object State isn't set Object was created (registration request was received and object created) Registration was approved and waiting for configuration This state can be set by user only if current state is NEW Registration is approved and prepared for to connect.. Possible values are", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:status", "enum_extraction_complete": true, "enum_validators": [{"case_sensitive": true, "complete": true, "source": "ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf", "validator": "OneOf", "values": ["ADMITTED", "APPROVED", "DONE", "FAILED", "FAILED_INACTIVE", "MAINTENANCE", "NEW", "NOTSET", "ONLINE", "PENDING", "RETIRED", "UPGRADING"], "version": 1}], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "status", "current_state"], "syntax": "attribute", "type": "string"}, {"aliases": ["items object status object status"], "anchor": "section", "description": "Status is a return value for calls that don't return other objects.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:status:object_status", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["items", "object", "status", "object_status"], "syntax": "attribute", "type": "object"}, {"aliases": ["items object status parent current state"], "anchor": "schema-items--object--status--parent_current_state", "description": "State of Site defines in which operational state site itself is. Site is online and operational. Site is in provisioning state. For instance during site deployment or switching to different connected Regional Edge. Site is in process of upgrade. Possible values are `ONLINE`, `PROVISIONING`, `UPGRADING`, `STANDBY`,", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:status", "enum_extraction_complete": true, "enum_validators": [{"case_sensitive": true, "complete": true, "source": "ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf", "validator": "OneOf", "values": ["DECOMMISSIONING", "DELETED_CLOUD_RESOURCES", "DELETE_QUEUED", "DELETING_CLOUD_RESOURCES", "ERROR_DELETING_CLOUD_RESOURCES", "ERROR_IN_ORCHESTRATION", "ERROR_UPDATING_CLOUD_RESOURCES", "FAILED", "FAILED_INACTIVE", "ONLINE", "ORCHESTRATION_COMPLETE", "ORCHESTRATION_IN_PROGRESS", "ORCHESTRATION_QUEUED", "PROVISIONING", "REREGISTRATION", "STANDBY", "UPDATE_QUEUED", "UPDATING_CLOUD_RESOURCES", "UPGRADING", "VALIDATION_FAILED", "VALIDATION_IN_PROGRESS", "VALIDATION_SUCCESS", "WAITINGNODES", "WAITING_FOR_REGISTRATION"], "version": 1}], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "status", "parent_current_state"], "syntax": "attribute", "type": "string"}, {"aliases": ["items object status state update timestamp"], "anchor": "schema-items--object--status--state_update_timestamp", "description": "Registration state update timestamp. Time of last registration state update.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:status", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "status", "state_update_timestamp"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_site/properties/items/object/status/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Status Type. Most recent observer status of object.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": [], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.object.status

Breadcrumbs:

- [xcsh_site_registrations_by_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/)
- [items.object](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/object/)
- items.object.status

<a id="section"></a>

Type: `"single"`. Computed.

Status Type. Most recent observer status of object.

## Direct properties

<a id="schema-items--object--status--current_state"></a>

### current_state property

Type: `"string"`. Computed.

\[Enum:
NOTSET|NEW|APPROVED|ADMITTED|RETIRED|FAILED|DONE|PENDING|ONLINE|UPGRADING|MAINTENANCE|FAILED\_INACTIVE\]
Defines states for registration object State isn't set Object was created (registration request was
received and object created) Registration was approved and waiting for configuration This state can
be set by user only if current state is NEW Registration is approved and prepared for to connect..
Possible values are \`NOTSET\`, \`NEW\`, \`APPROVED\`, \`ADMITTED\`, \`RETIRED\`, \`FAILED\`,
\`DONE\`, \`PENDING\`, \`ONLINE\`, \`UPGRADING\`, \`MAINTENANCE\`, \`FAILED\_INACTIVE\`. Defaults to
\`NOTSET\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["ADMITTED","APPROVED","DONE","FAILED","FAILED_INACTIVE","MAINTENANCE","NEW","NOTSET","ONLINE","PENDING","RETIRED","UPGRADING"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("NOTSET",
    "NEW",
    "APPROVED",
    "ADMITTED",
    "RETIRED",
    "FAILED",
    "DONE",
    "PENDING",
    "ONLINE",
    "UPGRADING",
    "MAINTENANCE",
    "FAILED_INACTIVE"),
}
```

- [object_status](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/object/status/object_status/): complete subsection reference.

<a id="schema-items--object--status--parent_current_state"></a>

### parent_current_state property

Type: `"string"`. Computed.

\[Enum:
ONLINE|PROVISIONING|UPGRADING|STANDBY|FAILED|REREGISTRATION|WAITINGNODES|DECOMMISSIONING|WAITING\_FOR\_REGISTRATION|ORCHESTRATION\_IN\_PROGRESS|ORCHESTRATION\_COMPLETE|ERROR\_IN\_ORCHESTRATION|DELETING\_CLOUD\_RESOURCES|DELETED\_CLOUD\_RESOURCES|ERROR\_DELETING\_CLOUD\_RESOURCES|VALIDATION\_IN\_PROGRESS|VALIDATION\_SUCCESS|VALIDATION\_FAILED|FAILED\_INACTIVE|UPDATING\_CLOUD\_RESOURCES|ERROR\_UPDATING\_CLOUD\_RESOURCES|ORCHESTRATION\_QUEUED|UPDATE\_QUEUED|DELETE\_QUEUED\]
State of Site defines in which operational state site itself is. Site is online and operational.
Site is in provisioning state. For instance during site deployment or switching to different
connected Regional Edge. Site is in process of upgrade. Possible values are \`ONLINE\`,
\`PROVISIONING\`, \`UPGRADING\`, \`STANDBY\`, \`FAILED\`, \`REREGISTRATION\`, \`WAITINGNODES\`,
\`DECOMMISSIONING\`, \`WAITING\_FOR\_REGISTRATION\`, \`ORCHESTRATION\_IN\_PROGRESS\`,
\`ORCHESTRATION\_COMPLETE\`, \`ERROR\_IN\_ORCHESTRATION\`, \`DELETING\_CLOUD\_RESOURCES\`,
\`DELETED\_CLOUD\_RESOURCES\`, \`ERROR\_DELETING\_CLOUD\_RESOURCES\`, \`VALIDATION\_IN\_PROGRESS\`,
\`VALIDATION\_SUCCESS\`, \`VALIDATION\_FAILED\`, \`FAILED\_INACTIVE\`,
\`UPDATING\_CLOUD\_RESOURCES\`, \`ERROR\_UPDATING\_CLOUD\_RESOURCES\`, \`ORCHESTRATION\_QUEUED\`,
\`UPDATE\_QUEUED\`, \`DELETE\_QUEUED\`. Defaults to \`ONLINE\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["DECOMMISSIONING","DELETED_CLOUD_RESOURCES","DELETE_QUEUED","DELETING_CLOUD_RESOURCES","ERROR_DELETING_CLOUD_RESOURCES","ERROR_IN_ORCHESTRATION","ERROR_UPDATING_CLOUD_RESOURCES","FAILED","FAILED_INACTIVE","ONLINE","ORCHESTRATION_COMPLETE","ORCHESTRATION_IN_PROGRESS","ORCHESTRATION_QUEUED","PROVISIONING","REREGISTRATION","STANDBY","UPDATE_QUEUED","UPDATING_CLOUD_RESOURCES","UPGRADING","VALIDATION_FAILED","VALIDATION_IN_PROGRESS","VALIDATION_SUCCESS","WAITINGNODES","WAITING_FOR_REGISTRATION"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("ONLINE",
    "PROVISIONING",
    "UPGRADING",
    "STANDBY",
    "FAILED",
    "REREGISTRATION",
    "WAITINGNODES",
    "DECOMMISSIONING",
    "WAITING_FOR_REGISTRATION",
    "ORCHESTRATION_IN_PROGRESS",
    "ORCHESTRATION_COMPLETE",
    "ERROR_IN_ORCHESTRATION",
    "DELETING_CLOUD_RESOURCES",
    "DELETED_CLOUD_RESOURCES",
    "ERROR_DELETING_CLOUD_RESOURCES",
    "VALIDATION_IN_PROGRESS",
    "VALIDATION_SUCCESS",
    "VALIDATION_FAILED",
    "FAILED_INACTIVE",
    "UPDATING_CLOUD_RESOURCES",
    "ERROR_UPDATING_CLOUD_RESOURCES",
    "ORCHESTRATION_QUEUED",
    "UPDATE_QUEUED",
    "DELETE_QUEUED"),
}
```

<a id="schema-items--object--status--state_update_timestamp"></a>

### state_update_timestamp property

Type: `"string"`. Computed.

Registration state update timestamp. Time of last registration state update.

## Next pages

- [items.object.status.object_status](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/object/status/object_status/)
- [items.object](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/object/)
- [xcsh_site_registrations_by_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/)
