---
page_title: "routes.notification_parameters"
subcategory: "Monitoring"
description: "Set of notification parameters to decide how and when the alert notifications should be sent to the receivers."
xcsh_docs: {"aliases": ["routes notification parameters"], "body_bytes": 7220, "body_sha256": "sha256:2c976a53e28bea7a37254e106f6e131d25466d046b10240dc315c40127152c60", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:data-sources:alert_policy:properties:routes:notification_parameters:custom", "xcsh-docs:data-sources:alert_policy:properties:routes:notification_parameters:default", "xcsh-docs:data-sources:alert_policy:properties:routes:notification_parameters:individual", "xcsh-docs:data-sources:alert_policy:properties:routes:notification_parameters:ves_io_group"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:alert_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:alert_policy:properties:routes:notification_parameters", "parent_id": "xcsh-docs:data-sources:alert_policy:properties:routes", "path": "documentation/data-sources/alert_policy/properties/routes/notification_parameters/index.md", "product": "distributed-cloud", "provider_name": "alert_policy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3222233311320102-0030322313100000-1023230020323111-3023210212302030-3002100220323100-0111221013302013-0223102120002031-0233120202310322", "registry_path": "docs/guides/data-sources--alert_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "notification_parameters"], "schema_version": 1, "sections": [{"aliases": ["routes notification parameters custom"], "anchor": "section", "description": "Specify list of custom labels to group/aggregate the alerts.", "document_id": "xcsh-docs:data-sources:alert_policy:properties:routes:notification_parameters:custom", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["routes", "notification_parameters", "custom"], "syntax": "attribute", "type": "object"}, {"aliases": ["routes notification parameters default"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:alert_policy:properties:routes:notification_parameters:default", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "notification_parameters", "default"], "syntax": "attribute", "type": "object"}, {"aliases": ["routes notification parameters group interval"], "anchor": "schema-routes--notification_parameters--group_interval", "description": "Group Interval is used to specify how long to wait before sending a notification about new alerts that are added to the group for which an initial notification has already been sent. Format: , where s - seconds, m - minutes, h - hours, d - days If not specified, group_interval defaults to \"1m\"", "document_id": "xcsh-docs:data-sources:alert_policy:properties:routes:notification_parameters", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "notification_parameters", "group_interval"], "syntax": "attribute", "type": "string"}, {"aliases": ["routes notification parameters group wait"], "anchor": "schema-routes--notification_parameters--group_wait", "description": "Time value used to specify how long to initially wait for an inhibiting alert to arrive or collect more alerts for the same group. Format: , where s - seconds, m - minutes, h - hours, d - days If not specified, group_wait defaults to \"30s\"", "document_id": "xcsh-docs:data-sources:alert_policy:properties:routes:notification_parameters", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "notification_parameters", "group_wait"], "syntax": "attribute", "type": "string"}, {"aliases": ["routes notification parameters individual"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:alert_policy:properties:routes:notification_parameters:individual", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "notification_parameters", "individual"], "syntax": "attribute", "type": "object"}, {"aliases": ["routes notification parameters repeat interval", "succeeded", "success", "successful"], "anchor": "schema-routes--notification_parameters--repeat_interval", "description": "Repeat Interval is used to specify how long to wait before sending a notification again if it has already been sent successfully. Format: , where s - seconds, m - minutes, h - hours, d - days If not specified, group_interval defaults to \"4h\"", "document_id": "xcsh-docs:data-sources:alert_policy:properties:routes:notification_parameters", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "notification_parameters", "repeat_interval"], "syntax": "attribute", "type": "string"}, {"aliases": ["routes notification parameters ves io group"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:alert_policy:properties:routes:notification_parameters:ves_io_group", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "notification_parameters", "ves_io_group"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/alert_policy/properties/routes/notification_parameters/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Set of notification parameters to decide how and when the alert notifications should be sent to the receivers.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["alert_policyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.notification_parameters

Breadcrumbs:

- [xcsh_alert_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/routes/)
- routes.notification_parameters

<a id="section"></a>

Type: `"single"`. Computed.

Set of notification parameters to decide how and when the alert notifications should be sent to the
receivers.

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

## Direct properties

- [custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/routes/notification_parameters/custom/): complete subsection reference.

- [default](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/routes/notification_parameters/default/): complete subsection reference.

<a id="schema-routes--notification_parameters--group_interval"></a>

### group_interval property

Type: `"string"`. Computed.

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

<a id="schema-routes--notification_parameters--group_wait"></a>

### group_wait property

Type: `"string"`. Computed.

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

- [individual](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/routes/notification_parameters/individual/): complete subsection reference.

<a id="schema-routes--notification_parameters--repeat_interval"></a>

### repeat_interval property

Type: `"string"`. Computed.

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

- [ves_io_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/routes/notification_parameters/ves_io_group/): complete subsection reference.

## Next pages

- [routes.notification_parameters.custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/routes/notification_parameters/custom/)
- [routes.notification_parameters.default](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/routes/notification_parameters/default/)
- [routes.notification_parameters.individual](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/routes/notification_parameters/individual/)
- [routes.notification_parameters.ves_io_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/routes/notification_parameters/ves_io_group/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/routes/)
- [xcsh_alert_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/)
