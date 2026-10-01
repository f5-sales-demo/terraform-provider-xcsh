---
page_title: "items.object.status"
subcategory: ""
description: "items.object.status for xcsh_site_registrations_by_site."
xcsh_docs: {"aliases": [], "body_bytes": 4565, "body_sha256": "sha256:9ce33e3cd62e3ead2cd5dff44c718ff1aadadc54c924fc9269767c3aa17f487c", "canonical_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:status", "child_ids": ["xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:status:object_status"], "collection_id": "xcsh-docs:data-sources:site_registrations_by_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:status", "parent_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object", "path": "docs/guides/data-sources--site_registrations_by_site--properties--items--object--status.md", "provider_name": "site_registrations_by_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["items", "object", "status"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_site/properties/items/object/status/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "items.object.status for xcsh_site_registrations_by_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.object.status

Breadcrumbs:

- [xcsh_site_registrations_by_site](../data-sources/site_registrations_by_site.md)
- [Property reference](data-sources--site_registrations_by_site--reference.md)
- [items](data-sources--site_registrations_by_site--properties--items.md)
- [items.object](data-sources--site_registrations_by_site--properties--items--object.md)
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

- [object_status](data-sources--site_registrations_by_site--properties--items--object--status--object_status.md): complete subsection reference.

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

- [items.object.status.object_status](data-sources--site_registrations_by_site--properties--items--object--status--object_status.md)
- [items.object](data-sources--site_registrations_by_site--properties--items--object.md)
- [xcsh_site_registrations_by_site](../data-sources/site_registrations_by_site.md)
