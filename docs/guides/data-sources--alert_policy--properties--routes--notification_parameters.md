---
page_title: "routes.notification_parameters"
subcategory: "Monitoring"
description: "routes.notification_parameters for xcsh_alert_policy."
xcsh_docs: {"aliases": [], "body_bytes": 6571, "body_sha256": "sha256:b5cded3578f59a99784ce27aa7201aac34d0cf45a8bedb15a39fd60219958b05", "canonical_id": "xcsh-docs:data-sources:alert_policy:properties:routes:notification_parameters", "child_ids": ["xcsh-docs:data-sources:alert_policy:properties:routes:notification_parameters:custom", "xcsh-docs:data-sources:alert_policy:properties:routes:notification_parameters:default", "xcsh-docs:data-sources:alert_policy:properties:routes:notification_parameters:individual", "xcsh-docs:data-sources:alert_policy:properties:routes:notification_parameters:ves_io_group"], "collection_id": "xcsh-docs:data-sources:alert_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:alert_policy:properties:routes:notification_parameters", "parent_id": "xcsh-docs:data-sources:alert_policy:properties:routes", "path": "docs/guides/data-sources--alert_policy--properties--routes--notification_parameters.md", "provider_name": "alert_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["routes", "notification_parameters"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/alert_policy/properties/routes/notification_parameters/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "routes.notification_parameters for xcsh_alert_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["alert_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.notification_parameters

Breadcrumbs:

- [xcsh_alert_policy](../data-sources/alert_policy.md)
- [Property reference](data-sources--alert_policy--reference.md)
- [routes](data-sources--alert_policy--properties--routes.md)
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

- [custom](data-sources--alert_policy--properties--routes--notification_parameters--custom.md): complete subsection reference.

- [default](data-sources--alert_policy--properties--routes--notification_parameters--default.md): complete subsection reference.

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

- [individual](data-sources--alert_policy--properties--routes--notification_parameters--individual.md): complete subsection reference.

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

- [ves_io_group](data-sources--alert_policy--properties--routes--notification_parameters--ves_io_group.md): complete subsection reference.

## Next pages

- [routes.notification_parameters.custom](data-sources--alert_policy--properties--routes--notification_parameters--custom.md)
- [routes.notification_parameters.default](data-sources--alert_policy--properties--routes--notification_parameters--default.md)
- [routes.notification_parameters.individual](data-sources--alert_policy--properties--routes--notification_parameters--individual.md)
- [routes.notification_parameters.ves_io_group](data-sources--alert_policy--properties--routes--notification_parameters--ves_io_group.md)
- [routes](data-sources--alert_policy--properties--routes.md)
- [xcsh_alert_policy](../data-sources/alert_policy.md)
