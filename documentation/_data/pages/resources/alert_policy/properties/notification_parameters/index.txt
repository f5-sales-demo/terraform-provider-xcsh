---
page_title: "notification_parameters"
subcategory: "Monitoring"
description: "Set of notification parameters to decide how and when the alert notifications should be sent to the receivers."
xcsh_docs: {"aliases": ["notification parameters"], "body_bytes": 7611, "body_sha256": "sha256:0459da58dd2353bc2dbb059fb5b9647451f589d61933100578bb0cac2a3b13bc", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:resources:alert_policy:properties:notification_parameters:custom", "xcsh-docs:resources:alert_policy:properties:notification_parameters:default", "xcsh-docs:resources:alert_policy:properties:notification_parameters:individual", "xcsh-docs:resources:alert_policy:properties:notification_parameters:ves_io_group"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:alert_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:alert_policy:properties:notification_parameters", "parent_id": "xcsh-docs:resources:alert_policy:reference", "path": "documentation/resources/alert_policy/properties/notification_parameters/index.md", "product": "distributed-cloud", "provider_name": "alert_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2302223211220010-3133333223123323-0320311202331210-3230201211033133-0220001322010300-2323033221331022-2321110220123020-0201210012212021", "registry_path": "docs/guides/resources--alert_policy--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "notification_parameters:ConflictingObjectAttributes:custom,default", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:notification_parameters:custom", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "notification_parameters:ConflictingObjectAttributes:custom,individual", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:notification_parameters:custom", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "notification_parameters:ConflictingObjectAttributes:custom,ves_io_group", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:notification_parameters:custom", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "notification_parameters:ConflictingObjectAttributes:custom,default", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:notification_parameters:default", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "notification_parameters:ConflictingObjectAttributes:default,individual", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:notification_parameters:default", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "notification_parameters:ConflictingObjectAttributes:default,ves_io_group", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:notification_parameters:default", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "notification_parameters:ConflictingObjectAttributes:custom,individual", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:notification_parameters:individual", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "notification_parameters:ConflictingObjectAttributes:default,individual", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:notification_parameters:individual", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "notification_parameters:ConflictingObjectAttributes:individual,ves_io_group", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:notification_parameters:individual", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "notification_parameters:ConflictingObjectAttributes:custom,ves_io_group", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:notification_parameters:ves_io_group", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "notification_parameters:ConflictingObjectAttributes:default,ves_io_group", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:notification_parameters:ves_io_group", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "notification_parameters:ConflictingObjectAttributes:individual,ves_io_group", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:notification_parameters:ves_io_group", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["notification_parameters"], "schema_version": 1, "sections": [{"aliases": ["custom"], "anchor": "section", "description": "Specify list of custom labels to group/aggregate the alerts.", "document_id": "xcsh-docs:resources:alert_policy:properties:notification_parameters:custom", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["notification_parameters", "custom"], "syntax": "block", "type": "object"}, {"aliases": ["default"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:alert_policy:properties:notification_parameters:default", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["notification_parameters", "default"], "syntax": "attribute", "type": "object"}, {"aliases": ["group interval"], "anchor": "schema-notification_parameters--group_interval", "description": "Group Interval is used to specify how long to wait before sending a notification about new alerts that are added to the group for which an initial notification has already been sent. Format: , where s - seconds, m - minutes, h - hours, d - days If not specified, group_interval defaults to \"1m\"", "document_id": "xcsh-docs:resources:alert_policy:properties:notification_parameters", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["notification_parameters", "group_interval"], "syntax": "attribute", "type": "string"}, {"aliases": ["group wait"], "anchor": "schema-notification_parameters--group_wait", "description": "Time value used to specify how long to initially wait for an inhibiting alert to arrive or collect more alerts for the same group. Format: , where s - seconds, m - minutes, h - hours, d - days If not specified, group_wait defaults to \"30s\"", "document_id": "xcsh-docs:resources:alert_policy:properties:notification_parameters", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["notification_parameters", "group_wait"], "syntax": "attribute", "type": "string"}, {"aliases": ["individual"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:alert_policy:properties:notification_parameters:individual", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["notification_parameters", "individual"], "syntax": "attribute", "type": "object"}, {"aliases": ["login success", "repeat interval", "succeeded", "success", "successful"], "anchor": "schema-notification_parameters--repeat_interval", "description": "Repeat Interval is used to specify how long to wait before sending a notification again if it has already been sent successfully. Format: , where s - seconds, m - minutes, h - hours, d - days If not specified, group_interval defaults to \"4h\"", "document_id": "xcsh-docs:resources:alert_policy:properties:notification_parameters", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["notification_parameters", "repeat_interval"], "syntax": "attribute", "type": "string"}, {"aliases": ["ves io group"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:alert_policy:properties:notification_parameters:ves_io_group", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["notification_parameters", "ves_io_group"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/alert_policy/properties/notification_parameters/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Set of notification parameters to decide how and when the alert notifications should be sent to the receivers.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["alert_policyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# notification_parameters

Breadcrumbs:

- [xcsh_alert_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/)
- notification_parameters

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Set of notification parameters to decide how and when the alert notifications should be sent to the
receivers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom",
    "default"),
  validators.ConflictingObjectAttributes("custom",
    "individual"),
  validators.ConflictingObjectAttributes("custom",
    "ves_io_group"),
  validators.ConflictingObjectAttributes("default",
    "individual"),
  validators.ConflictingObjectAttributes("default",
    "ves_io_group"),
  validators.ConflictingObjectAttributes("individual",
    "ves_io_group")}
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
  "x-ves-oneof-field-group_by": "[\"custom\",\"default\",\"individual\",\"ves_io_group\"]"
}
```

Terraform syntax:

```terraform
notification_parameters {
  # Configure direct properties listed below.
}
```

## Direct properties

- [custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/notification_parameters/custom/): complete subsection reference.

- [default](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/notification_parameters/default/): complete subsection reference.

<a id="schema-notification_parameters--group_interval"></a>

### group_interval property

Type: `"string"`. Optional.

Group Interval is used to specify how long to wait before sending a notification about new alerts
that are added to the group for which an initial notification has already been sent. Format:
\[0-9\]\[smhd\], where s - seconds, m - minutes, h - hours, d - days If not specified,
group\_interval..

Upstream description:

Group Interval is used to specify how long to wait before sending a notification about new alerts
that are added to the group for which an initial notification has already been sent. Format:
\[0-9\]\[smhd\], where s - seconds, m - minutes, h - hours, d - days If not specified,
group\_interval defaults to "1m"

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_time_interval": "10m",
    "ves.io.schema.rules.string.min_time_interval": "30s",
    "ves.io.schema.rules.string.time_interval": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_time_interval": "10m",
    "ves.io.schema.rules.string.min_time_interval": "30s",
    "ves.io.schema.rules.string.time_interval": "true"
  }
}
```

<a id="schema-notification_parameters--group_wait"></a>

### group_wait property

Type: `"string"`. Optional.

Time value used to specify how long to initially wait for an inhibiting alert to arrive or collect
more alerts for the same group. Format: \[0-9\]\[smhd\], where s - seconds, m - minutes, h - hours,
d - days If not specified, group\_wait defaults to '30s'.

Upstream description:

Time value used to specify how long to initially wait for an inhibiting alert to arrive or collect
more alerts for the same group. Format: \[0-9\]\[smhd\], where s - seconds, m - minutes, h - hours,
d - days If not specified, group\_wait defaults to "30s"

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_time_interval": "5m",
    "ves.io.schema.rules.string.min_time_interval": "0s",
    "ves.io.schema.rules.string.time_interval": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_time_interval": "5m",
    "ves.io.schema.rules.string.min_time_interval": "0s",
    "ves.io.schema.rules.string.time_interval": "true"
  }
}
```

- [individual](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/notification_parameters/individual/): complete subsection reference.

<a id="schema-notification_parameters--repeat_interval"></a>

### repeat_interval property

Type: `"string"`. Optional.

Repeat Interval is used to specify how long to wait before sending a notification again if it has
already been sent successfully. Format: \[0-9\]\[smhd\], where s - seconds, m - minutes, h - hours,
d - days If not specified, group\_interval defaults to '4h'.

Upstream description:

Repeat Interval is used to specify how long to wait before sending a notification again if it has
already been sent successfully. Format: \[0-9\]\[smhd\], where s - seconds, m - minutes, h - hours,
d - days If not specified, group\_interval defaults to "4h"

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_time_interval": "10d",
    "ves.io.schema.rules.string.min_time_interval": "30m",
    "ves.io.schema.rules.string.time_interval": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_time_interval": "10d",
    "ves.io.schema.rules.string.min_time_interval": "30m",
    "ves.io.schema.rules.string.time_interval": "true"
  }
}
```

- [ves_io_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/notification_parameters/ves_io_group/): complete subsection reference.

## Next pages

- [notification_parameters.custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/notification_parameters/custom/)
- [notification_parameters.default](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/notification_parameters/default/)
- [notification_parameters.individual](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/notification_parameters/individual/)
- [notification_parameters.ves_io_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/notification_parameters/ves_io_group/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/)
- [xcsh_alert_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/)
