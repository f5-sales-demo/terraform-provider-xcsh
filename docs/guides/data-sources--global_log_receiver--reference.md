---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_global_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 81879, "body_sha256": "sha256:3d7b164ff3ff7ec559ec995c27ddf2fd59cf7ae8711963092a37a1e5812c1613", "canonical_id": "xcsh-docs:data-sources:global_log_receiver:reference", "child_ids": ["xcsh-docs:data-sources:global_log_receiver:properties:audit_logs", "xcsh-docs:data-sources:global_log_receiver:properties:aws_cloud_watch_receiver", "xcsh-docs:data-sources:global_log_receiver:properties:azure_event_hubs_receiver", "xcsh-docs:data-sources:global_log_receiver:properties:azure_receiver", "xcsh-docs:data-sources:global_log_receiver:properties:datadog_receiver", "xcsh-docs:data-sources:global_log_receiver:properties:dns_logs", "xcsh-docs:data-sources:global_log_receiver:properties:gcp_bucket_receiver", "xcsh-docs:data-sources:global_log_receiver:properties:http_receiver", "xcsh-docs:data-sources:global_log_receiver:properties:kafka_receiver", "xcsh-docs:data-sources:global_log_receiver:properties:new_relic_receiver", "xcsh-docs:data-sources:global_log_receiver:properties:ns_all", "xcsh-docs:data-sources:global_log_receiver:properties:ns_current", "xcsh-docs:data-sources:global_log_receiver:properties:ns_list", "xcsh-docs:data-sources:global_log_receiver:properties:qradar_receiver", "xcsh-docs:data-sources:global_log_receiver:properties:request_logs", "xcsh-docs:data-sources:global_log_receiver:properties:s3_receiver", "xcsh-docs:data-sources:global_log_receiver:properties:security_events", "xcsh-docs:data-sources:global_log_receiver:properties:splunk_receiver", "xcsh-docs:data-sources:global_log_receiver:properties:sumo_logic_receiver"], "collection_id": "xcsh-docs:data-sources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:global_log_receiver:reference", "parent_id": "xcsh-docs:data-sources:global_log_receiver:fundamentals", "path": "docs/guides/data-sources--global_log_receiver--reference.md", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/global_log_receiver/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_global_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Property reference

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md)
- Property reference

## Direct properties

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

Upstream description:

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

- [audit_logs](data-sources--global_log_receiver--properties--audit_logs.md): complete subsection reference.

- [aws_cloud_watch_receiver](data-sources--global_log_receiver--properties--aws_cloud_watch_receiver.md): complete subsection reference.

- [azure_event_hubs_receiver](data-sources--global_log_receiver--properties--azure_event_hubs_receiver.md): complete subsection reference.

- [azure_receiver](data-sources--global_log_receiver--properties--azure_receiver.md): complete subsection reference.

- [datadog_receiver](data-sources--global_log_receiver--properties--datadog_receiver.md): complete subsection reference.

<a id="schema-description"></a>

### description property

Type: `"string"`. Computed.

Description of the GlobalLogReceiver.

Upstream description:

Human readable description for the object.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1200,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1200
    },
    "category": "discovery",
    "characterSet": {
      "description": "Free text with UTF-8 support"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1200,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  }
}
```

- [dns_logs](data-sources--global_log_receiver--properties--dns_logs.md): complete subsection reference.

- [gcp_bucket_receiver](data-sources--global_log_receiver--properties--gcp_bucket_receiver.md): complete subsection reference.

- [http_receiver](data-sources--global_log_receiver--properties--http_receiver.md): complete subsection reference.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

- [kafka_receiver](data-sources--global_log_receiver--properties--kafka_receiver.md): complete subsection reference.

<a id="schema-labels"></a>

### labels property

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

Upstream description:

Map of string keys and values that can be used to organize and categorize (scope and select) objects
as chosen by the user. Values specified here will be used by selector expression.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the GlobalLogReceiver.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Required.

Namespace where the GlobalLogReceiver exists.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [new_relic_receiver](data-sources--global_log_receiver--properties--new_relic_receiver.md): complete subsection reference.

- [ns_all](data-sources--global_log_receiver--properties--ns_all.md): complete subsection reference.

- [ns_current](data-sources--global_log_receiver--properties--ns_current.md): complete subsection reference.

- [ns_list](data-sources--global_log_receiver--properties--ns_list.md): complete subsection reference.

- [qradar_receiver](data-sources--global_log_receiver--properties--qradar_receiver.md): complete subsection reference.

- [request_logs](data-sources--global_log_receiver--properties--request_logs.md): complete subsection reference.

- [s3_receiver](data-sources--global_log_receiver--properties--s3_receiver.md): complete subsection reference.

- [security_events](data-sources--global_log_receiver--properties--security_events.md): complete subsection reference.

- [splunk_receiver](data-sources--global_log_receiver--properties--splunk_receiver.md): complete subsection reference.

- [sumo_logic_receiver](data-sources--global_log_receiver--properties--sumo_logic_receiver.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--global_log_receiver--reference.md#schema-annotations) |
| `audit_logs` | [audit_logs](data-sources--global_log_receiver--properties--audit_logs.md#section) |
| `aws_cloud_watch_receiver` | [aws_cloud_watch_receiver](data-sources--global_log_receiver--properties--aws_cloud_watch_receiver.md#section) |
| `aws_cloud_watch_receiver.aws_cred` | [aws_cloud_watch_receiver.aws_cred](data-sources--global_log_receiver--properties--aws_cloud_watch_receiver--aws_cred.md#section) |
| `aws_cloud_watch_receiver.aws_cred.name` | [aws_cloud_watch_receiver.aws_cred.name](data-sources--global_log_receiver--properties--aws_cloud_watch_receiver--aws_cred.md#schema-aws_cloud_watch_receiver--aws_cred--name) |
| `aws_cloud_watch_receiver.aws_cred.namespace` | [aws_cloud_watch_receiver.aws_cred.namespace](data-sources--global_log_receiver--properties--aws_cloud_watch_receiver--aws_cred.md#schema-aws_cloud_watch_receiver--aws_cred--namespace) |
| `aws_cloud_watch_receiver.aws_cred.tenant` | [aws_cloud_watch_receiver.aws_cred.tenant](data-sources--global_log_receiver--properties--aws_cloud_watch_receiver--aws_cred.md#schema-aws_cloud_watch_receiver--aws_cred--tenant) |
| `aws_cloud_watch_receiver.aws_region` | [aws_cloud_watch_receiver.aws_region](data-sources--global_log_receiver--properties--aws_cloud_watch_receiver.md#schema-aws_cloud_watch_receiver--aws_region) |
| `aws_cloud_watch_receiver.batch` | [aws_cloud_watch_receiver.batch](data-sources--global_log_receiver--properties--aws_cloud_watch_receiver--batch.md#section) |
| `aws_cloud_watch_receiver.batch.max_bytes` | [aws_cloud_watch_receiver.batch.max_bytes](data-sources--global_log_receiver--properties--aws_cloud_watch_receiver--batch.md#schema-aws_cloud_watch_receiver--batch--max_bytes) |
| `aws_cloud_watch_receiver.batch.max_bytes_disabled` | [aws_cloud_watch_receiver.batch.max_bytes_disabled](data-sources--global_log_receiver--properties--aws_cloud_watch_receiver--batch--max_bytes_disabled.md#section) |
| `aws_cloud_watch_receiver.batch.max_events` | [aws_cloud_watch_receiver.batch.max_events](data-sources--global_log_receiver--properties--aws_cloud_watch_receiver--batch.md#schema-aws_cloud_watch_receiver--batch--max_events) |
| `aws_cloud_watch_receiver.batch.max_events_disabled` | [aws_cloud_watch_receiver.batch.max_events_disabled](data-sources--global_log_receiver--properties--aws_cloud_watch_receiver--batch--max_events_disabled.md#section) |
| `aws_cloud_watch_receiver.batch.timeout_seconds` | [aws_cloud_watch_receiver.batch.timeout_seconds](data-sources--global_log_receiver--properties--aws_cloud_watch_receiver--batch.md#schema-aws_cloud_watch_receiver--batch--timeout_seconds) |
| `aws_cloud_watch_receiver.batch.timeout_seconds_default` | [aws_cloud_watch_receiver.batch.timeout_seconds_default](data-sources--global_log_receiver--properties--aws_cloud_watch_receiver--batch--timeout_seconds_default.md#section) |
| `aws_cloud_watch_receiver.compression` | [aws_cloud_watch_receiver.compression](data-sources--global_log_receiver--properties--aws_cloud_watch_receiver--compression.md#section) |
| `aws_cloud_watch_receiver.compression.compression_default` | [aws_cloud_watch_receiver.compression.compression_default](data-sources--global_log_receiver--properties--aws_cloud_watch_receiver--compression--compression_default.md#section) |
| `aws_cloud_watch_receiver.compression.compression_gzip` | [aws_cloud_watch_receiver.compression.compression_gzip](data-sources--global_log_receiver--properties--aws_cloud_watch_receiver--compression--compression_gzip.md#section) |
| `aws_cloud_watch_receiver.compression.compression_none` | [aws_cloud_watch_receiver.compression.compression_none](data-sources--global_log_receiver--properties--aws_cloud_watch_receiver--compression--compression_none.md#section) |
| `aws_cloud_watch_receiver.group_name` | [aws_cloud_watch_receiver.group_name](data-sources--global_log_receiver--properties--aws_cloud_watch_receiver.md#schema-aws_cloud_watch_receiver--group_name) |
| `aws_cloud_watch_receiver.stream_name` | [aws_cloud_watch_receiver.stream_name](data-sources--global_log_receiver--properties--aws_cloud_watch_receiver.md#schema-aws_cloud_watch_receiver--stream_name) |
| `azure_event_hubs_receiver` | [azure_event_hubs_receiver](data-sources--global_log_receiver--properties--azure_event_hubs_receiver.md#section) |
| `azure_event_hubs_receiver.connection_string` | [azure_event_hubs_receiver.connection_string](data-sources--global_log_receiver--properties--azure_event_hubs_receiver--connection_string.md#section) |
| `azure_event_hubs_receiver.connection_string.blindfold_secret_info` | [azure_event_hubs_receiver.connection_string.blindfold_secret_info](data-sources--global_log_receiver--properties--azure_event_hubs_receiver--connection_string--blindfold_secret_info.md#section) |
| `azure_event_hubs_receiver.connection_string.blindfold_secret_info.decryption_provider` | [azure_event_hubs_receiver.connection_string.blindfold_secret_info.decryption_provider](data-sources--global_log_receiver--properties--azure_event_hubs_receiver--connection_string--blindfold_secret_info.md#schema-azure_event_hubs_receiver--connection_string--blindfold_secret_info--decryption_provider) |
| `azure_event_hubs_receiver.connection_string.blindfold_secret_info.location` | [azure_event_hubs_receiver.connection_string.blindfold_secret_info.location](data-sources--global_log_receiver--properties--azure_event_hubs_receiver--connection_string--blindfold_secret_info.md#schema-azure_event_hubs_receiver--connection_string--blindfold_secret_info--location) |
| `azure_event_hubs_receiver.connection_string.blindfold_secret_info.store_provider` | [azure_event_hubs_receiver.connection_string.blindfold_secret_info.store_provider](data-sources--global_log_receiver--properties--azure_event_hubs_receiver--connection_string--blindfold_secret_info.md#schema-azure_event_hubs_receiver--connection_string--blindfold_secret_info--store_provider) |
| `azure_event_hubs_receiver.connection_string.clear_secret_info` | [azure_event_hubs_receiver.connection_string.clear_secret_info](data-sources--global_log_receiver--properties--azure_event_hubs_receiver--connection_string--clear_secret_info.md#section) |
| `azure_event_hubs_receiver.connection_string.clear_secret_info.provider_ref` | [azure_event_hubs_receiver.connection_string.clear_secret_info.provider_ref](data-sources--global_log_receiver--properties--azure_event_hubs_receiver--connection_string--clear_secret_info.md#schema-azure_event_hubs_receiver--connection_string--clear_secret_info--provider_ref) |
| `azure_event_hubs_receiver.connection_string.clear_secret_info.url` | [azure_event_hubs_receiver.connection_string.clear_secret_info.url](data-sources--global_log_receiver--properties--azure_event_hubs_receiver--connection_string--clear_secret_info.md#schema-azure_event_hubs_receiver--connection_string--clear_secret_info--url) |
| `azure_event_hubs_receiver.instance` | [azure_event_hubs_receiver.instance](data-sources--global_log_receiver--properties--azure_event_hubs_receiver.md#schema-azure_event_hubs_receiver--instance) |
| `azure_event_hubs_receiver.namespace` | [azure_event_hubs_receiver.namespace](data-sources--global_log_receiver--properties--azure_event_hubs_receiver.md#schema-azure_event_hubs_receiver--namespace) |
| `azure_receiver` | [azure_receiver](data-sources--global_log_receiver--properties--azure_receiver.md#section) |
| `azure_receiver.batch` | [azure_receiver.batch](data-sources--global_log_receiver--properties--azure_receiver--batch.md#section) |
| `azure_receiver.batch.max_bytes` | [azure_receiver.batch.max_bytes](data-sources--global_log_receiver--properties--azure_receiver--batch.md#schema-azure_receiver--batch--max_bytes) |
| `azure_receiver.batch.max_bytes_disabled` | [azure_receiver.batch.max_bytes_disabled](data-sources--global_log_receiver--properties--azure_receiver--batch--max_bytes_disabled.md#section) |
| `azure_receiver.batch.max_events` | [azure_receiver.batch.max_events](data-sources--global_log_receiver--properties--azure_receiver--batch.md#schema-azure_receiver--batch--max_events) |
| `azure_receiver.batch.max_events_disabled` | [azure_receiver.batch.max_events_disabled](data-sources--global_log_receiver--properties--azure_receiver--batch--max_events_disabled.md#section) |
| `azure_receiver.batch.timeout_seconds` | [azure_receiver.batch.timeout_seconds](data-sources--global_log_receiver--properties--azure_receiver--batch.md#schema-azure_receiver--batch--timeout_seconds) |
| `azure_receiver.batch.timeout_seconds_default` | [azure_receiver.batch.timeout_seconds_default](data-sources--global_log_receiver--properties--azure_receiver--batch--timeout_seconds_default.md#section) |
| `azure_receiver.compression` | [azure_receiver.compression](data-sources--global_log_receiver--properties--azure_receiver--compression.md#section) |
| `azure_receiver.compression.compression_default` | [azure_receiver.compression.compression_default](data-sources--global_log_receiver--properties--azure_receiver--compression--compression_default.md#section) |
| `azure_receiver.compression.compression_gzip` | [azure_receiver.compression.compression_gzip](data-sources--global_log_receiver--properties--azure_receiver--compression--compression_gzip.md#section) |
| `azure_receiver.compression.compression_none` | [azure_receiver.compression.compression_none](data-sources--global_log_receiver--properties--azure_receiver--compression--compression_none.md#section) |
| `azure_receiver.connection_string` | [azure_receiver.connection_string](data-sources--global_log_receiver--properties--azure_receiver--connection_string.md#section) |
| `azure_receiver.connection_string.blindfold_secret_info` | [azure_receiver.connection_string.blindfold_secret_info](data-sources--global_log_receiver--properties--azure_receiver--connection_string--blindfold_secret_info.md#section) |
| `azure_receiver.connection_string.blindfold_secret_info.decryption_provider` | [azure_receiver.connection_string.blindfold_secret_info.decryption_provider](data-sources--global_log_receiver--properties--azure_receiver--connection_string--blindfold_secret_info.md#schema-azure_receiver--connection_string--blindfold_secret_info--decryption_provider) |
| `azure_receiver.connection_string.blindfold_secret_info.location` | [azure_receiver.connection_string.blindfold_secret_info.location](data-sources--global_log_receiver--properties--azure_receiver--connection_string--blindfold_secret_info.md#schema-azure_receiver--connection_string--blindfold_secret_info--location) |
| `azure_receiver.connection_string.blindfold_secret_info.store_provider` | [azure_receiver.connection_string.blindfold_secret_info.store_provider](data-sources--global_log_receiver--properties--azure_receiver--connection_string--blindfold_secret_info.md#schema-azure_receiver--connection_string--blindfold_secret_info--store_provider) |
| `azure_receiver.connection_string.clear_secret_info` | [azure_receiver.connection_string.clear_secret_info](data-sources--global_log_receiver--properties--azure_receiver--connection_string--clear_secret_info.md#section) |
| `azure_receiver.connection_string.clear_secret_info.provider_ref` | [azure_receiver.connection_string.clear_secret_info.provider_ref](data-sources--global_log_receiver--properties--azure_receiver--connection_string--clear_secret_info.md#schema-azure_receiver--connection_string--clear_secret_info--provider_ref) |
| `azure_receiver.connection_string.clear_secret_info.url` | [azure_receiver.connection_string.clear_secret_info.url](data-sources--global_log_receiver--properties--azure_receiver--connection_string--clear_secret_info.md#schema-azure_receiver--connection_string--clear_secret_info--url) |
| `azure_receiver.container_name` | [azure_receiver.container_name](data-sources--global_log_receiver--properties--azure_receiver.md#schema-azure_receiver--container_name) |
| `azure_receiver.filename_options` | [azure_receiver.filename_options](data-sources--global_log_receiver--properties--azure_receiver--filename_options.md#section) |
| `azure_receiver.filename_options.custom_folder` | [azure_receiver.filename_options.custom_folder](data-sources--global_log_receiver--properties--azure_receiver--filename_options.md#schema-azure_receiver--filename_options--custom_folder) |
| `azure_receiver.filename_options.log_type_folder` | [azure_receiver.filename_options.log_type_folder](data-sources--global_log_receiver--properties--azure_receiver--filename_options--log_type_folder.md#section) |
| `azure_receiver.filename_options.no_folder` | [azure_receiver.filename_options.no_folder](data-sources--global_log_receiver--properties--azure_receiver--filename_options--no_folder.md#section) |
| `datadog_receiver` | [datadog_receiver](data-sources--global_log_receiver--properties--datadog_receiver.md#section) |
| `datadog_receiver.batch` | [datadog_receiver.batch](data-sources--global_log_receiver--properties--datadog_receiver--batch.md#section) |
| `datadog_receiver.batch.max_bytes` | [datadog_receiver.batch.max_bytes](data-sources--global_log_receiver--properties--datadog_receiver--batch.md#schema-datadog_receiver--batch--max_bytes) |
| `datadog_receiver.batch.max_bytes_disabled` | [datadog_receiver.batch.max_bytes_disabled](data-sources--global_log_receiver--properties--datadog_receiver--batch--max_bytes_disabled.md#section) |
| `datadog_receiver.batch.max_events` | [datadog_receiver.batch.max_events](data-sources--global_log_receiver--properties--datadog_receiver--batch.md#schema-datadog_receiver--batch--max_events) |
| `datadog_receiver.batch.max_events_disabled` | [datadog_receiver.batch.max_events_disabled](data-sources--global_log_receiver--properties--datadog_receiver--batch--max_events_disabled.md#section) |
| `datadog_receiver.batch.timeout_seconds` | [datadog_receiver.batch.timeout_seconds](data-sources--global_log_receiver--properties--datadog_receiver--batch.md#schema-datadog_receiver--batch--timeout_seconds) |
| `datadog_receiver.batch.timeout_seconds_default` | [datadog_receiver.batch.timeout_seconds_default](data-sources--global_log_receiver--properties--datadog_receiver--batch--timeout_seconds_default.md#section) |
| `datadog_receiver.compression` | [datadog_receiver.compression](data-sources--global_log_receiver--properties--datadog_receiver--compression.md#section) |
| `datadog_receiver.compression.compression_default` | [datadog_receiver.compression.compression_default](data-sources--global_log_receiver--properties--datadog_receiver--compression--compression_default.md#section) |
| `datadog_receiver.compression.compression_gzip` | [datadog_receiver.compression.compression_gzip](data-sources--global_log_receiver--properties--datadog_receiver--compression--compression_gzip.md#section) |
| `datadog_receiver.compression.compression_none` | [datadog_receiver.compression.compression_none](data-sources--global_log_receiver--properties--datadog_receiver--compression--compression_none.md#section) |
| `datadog_receiver.datadog_api_key` | [datadog_receiver.datadog_api_key](data-sources--global_log_receiver--properties--datadog_receiver--datadog_api_key.md#section) |
| `datadog_receiver.datadog_api_key.blindfold_secret_info` | [datadog_receiver.datadog_api_key.blindfold_secret_info](data-sources--global_log_receiver--properties--datadog_receiver--datadog_api_key--blindfold_secret_info.md#section) |
| `datadog_receiver.datadog_api_key.blindfold_secret_info.decryption_provider` | [datadog_receiver.datadog_api_key.blindfold_secret_info.decryption_provider](data-sources--global_log_receiver--properties--datadog_receiver--datadog_api_key--blindfold_secret_info.md#schema-datadog_receiver--datadog_api_key--blindfold_secret_info--decryption_provider) |
| `datadog_receiver.datadog_api_key.blindfold_secret_info.location` | [datadog_receiver.datadog_api_key.blindfold_secret_info.location](data-sources--global_log_receiver--properties--datadog_receiver--datadog_api_key--blindfold_secret_info.md#schema-datadog_receiver--datadog_api_key--blindfold_secret_info--location) |
| `datadog_receiver.datadog_api_key.blindfold_secret_info.store_provider` | [datadog_receiver.datadog_api_key.blindfold_secret_info.store_provider](data-sources--global_log_receiver--properties--datadog_receiver--datadog_api_key--blindfold_secret_info.md#schema-datadog_receiver--datadog_api_key--blindfold_secret_info--store_provider) |
| `datadog_receiver.datadog_api_key.clear_secret_info` | [datadog_receiver.datadog_api_key.clear_secret_info](data-sources--global_log_receiver--properties--datadog_receiver--datadog_api_key--clear_secret_info.md#section) |
| `datadog_receiver.datadog_api_key.clear_secret_info.provider_ref` | [datadog_receiver.datadog_api_key.clear_secret_info.provider_ref](data-sources--global_log_receiver--properties--datadog_receiver--datadog_api_key--clear_secret_info.md#schema-datadog_receiver--datadog_api_key--clear_secret_info--provider_ref) |
| `datadog_receiver.datadog_api_key.clear_secret_info.url` | [datadog_receiver.datadog_api_key.clear_secret_info.url](data-sources--global_log_receiver--properties--datadog_receiver--datadog_api_key--clear_secret_info.md#schema-datadog_receiver--datadog_api_key--clear_secret_info--url) |
| `datadog_receiver.endpoint` | [datadog_receiver.endpoint](data-sources--global_log_receiver--properties--datadog_receiver.md#schema-datadog_receiver--endpoint) |
| `datadog_receiver.no_tls` | [datadog_receiver.no_tls](data-sources--global_log_receiver--properties--datadog_receiver--no_tls.md#section) |
| `datadog_receiver.site` | [datadog_receiver.site](data-sources--global_log_receiver--properties--datadog_receiver.md#schema-datadog_receiver--site) |
| `datadog_receiver.use_tls` | [datadog_receiver.use_tls](data-sources--global_log_receiver--properties--datadog_receiver--use_tls.md#section) |
| `datadog_receiver.use_tls.disable_verify_certificate` | [datadog_receiver.use_tls.disable_verify_certificate](data-sources--global_log_receiver--properties--datadog_receiver--use_tls--disable_verify_certificate.md#section) |
| `datadog_receiver.use_tls.disable_verify_hostname` | [datadog_receiver.use_tls.disable_verify_hostname](data-sources--global_log_receiver--properties--datadog_receiver--use_tls--disable_verify_hostname.md#section) |
| `datadog_receiver.use_tls.enable_verify_certificate` | [datadog_receiver.use_tls.enable_verify_certificate](data-sources--global_log_receiver--properties--datadog_receiver--use_tls--enable_verify_certificate.md#section) |
| `datadog_receiver.use_tls.enable_verify_hostname` | [datadog_receiver.use_tls.enable_verify_hostname](data-sources--global_log_receiver--properties--datadog_receiver--use_tls--enable_verify_hostname.md#section) |
| `datadog_receiver.use_tls.mtls_disabled` | [datadog_receiver.use_tls.mtls_disabled](data-sources--global_log_receiver--properties--datadog_receiver--use_tls--mtls_disabled.md#section) |
| `datadog_receiver.use_tls.mtls_enable` | [datadog_receiver.use_tls.mtls_enable](data-sources--global_log_receiver--properties--datadog_receiver--use_tls--mtls_enable.md#section) |
| `datadog_receiver.use_tls.mtls_enable.certificate` | [datadog_receiver.use_tls.mtls_enable.certificate](data-sources--global_log_receiver--properties--datadog_receiver--use_tls--mtls_enable.md#schema-datadog_receiver--use_tls--mtls_enable--certificate) |
| `datadog_receiver.use_tls.mtls_enable.key_url` | [datadog_receiver.use_tls.mtls_enable.key_url](data-sources--global_log_receiver--properties--datadog_receiver--use_tls--mtls_enable--key_url.md#section) |
| `datadog_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info` | [datadog_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info](data-sources--global_log_receiver--properties--datadog_receiver--use_tls--mtls_enable--key_url--blindfold_secret_info.md#section) |
| `datadog_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.decryption_provider` | [datadog_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.decryption_provider](data-sources--global_log_receiver--properties--datadog_receiver--use_tls--mtls_enable--key_url--blindfold_secret_info.md#schema-datadog_receiver--use_tls--mtls_enable--key_url--blindfold_secret_info--decryption_provider) |
| `datadog_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.location` | [datadog_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.location](data-sources--global_log_receiver--properties--datadog_receiver--use_tls--mtls_enable--key_url--blindfold_secret_info.md#schema-datadog_receiver--use_tls--mtls_enable--key_url--blindfold_secret_info--location) |
| `datadog_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.store_provider` | [datadog_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.store_provider](data-sources--global_log_receiver--properties--datadog_receiver--use_tls--mtls_enable--key_url--blindfold_secret_info.md#schema-datadog_receiver--use_tls--mtls_enable--key_url--blindfold_secret_info--store_provider) |
| `datadog_receiver.use_tls.mtls_enable.key_url.clear_secret_info` | [datadog_receiver.use_tls.mtls_enable.key_url.clear_secret_info](data-sources--global_log_receiver--properties--datadog_receiver--use_tls--mtls_enable--key_url--clear_secret_info.md#section) |
| `datadog_receiver.use_tls.mtls_enable.key_url.clear_secret_info.provider_ref` | [datadog_receiver.use_tls.mtls_enable.key_url.clear_secret_info.provider_ref](data-sources--global_log_receiver--properties--datadog_receiver--use_tls--mtls_enable--key_url--clear_secret_info.md#schema-datadog_receiver--use_tls--mtls_enable--key_url--clear_secret_info--provider_ref) |
| `datadog_receiver.use_tls.mtls_enable.key_url.clear_secret_info.url` | [datadog_receiver.use_tls.mtls_enable.key_url.clear_secret_info.url](data-sources--global_log_receiver--properties--datadog_receiver--use_tls--mtls_enable--key_url--clear_secret_info.md#schema-datadog_receiver--use_tls--mtls_enable--key_url--clear_secret_info--url) |
| `datadog_receiver.use_tls.no_ca` | [datadog_receiver.use_tls.no_ca](data-sources--global_log_receiver--properties--datadog_receiver--use_tls--no_ca.md#section) |
| `datadog_receiver.use_tls.trusted_ca_url` | [datadog_receiver.use_tls.trusted_ca_url](data-sources--global_log_receiver--properties--datadog_receiver--use_tls.md#schema-datadog_receiver--use_tls--trusted_ca_url) |
| `description` | [description](data-sources--global_log_receiver--reference.md#schema-description) |
| `dns_logs` | [dns_logs](data-sources--global_log_receiver--properties--dns_logs.md#section) |
| `gcp_bucket_receiver` | [gcp_bucket_receiver](data-sources--global_log_receiver--properties--gcp_bucket_receiver.md#section) |
| `gcp_bucket_receiver.batch` | [gcp_bucket_receiver.batch](data-sources--global_log_receiver--properties--gcp_bucket_receiver--batch.md#section) |
| `gcp_bucket_receiver.batch.max_bytes` | [gcp_bucket_receiver.batch.max_bytes](data-sources--global_log_receiver--properties--gcp_bucket_receiver--batch.md#schema-gcp_bucket_receiver--batch--max_bytes) |
| `gcp_bucket_receiver.batch.max_bytes_disabled` | [gcp_bucket_receiver.batch.max_bytes_disabled](data-sources--global_log_receiver--properties--gcp_bucket_receiver--batch--max_bytes_disabled.md#section) |
| `gcp_bucket_receiver.batch.max_events` | [gcp_bucket_receiver.batch.max_events](data-sources--global_log_receiver--properties--gcp_bucket_receiver--batch.md#schema-gcp_bucket_receiver--batch--max_events) |
| `gcp_bucket_receiver.batch.max_events_disabled` | [gcp_bucket_receiver.batch.max_events_disabled](data-sources--global_log_receiver--properties--gcp_bucket_receiver--batch--max_events_disabled.md#section) |
| `gcp_bucket_receiver.batch.timeout_seconds` | [gcp_bucket_receiver.batch.timeout_seconds](data-sources--global_log_receiver--properties--gcp_bucket_receiver--batch.md#schema-gcp_bucket_receiver--batch--timeout_seconds) |
| `gcp_bucket_receiver.batch.timeout_seconds_default` | [gcp_bucket_receiver.batch.timeout_seconds_default](data-sources--global_log_receiver--properties--gcp_bucket_receiver--batch--timeout_seconds_default.md#section) |
| `gcp_bucket_receiver.bucket` | [gcp_bucket_receiver.bucket](data-sources--global_log_receiver--properties--gcp_bucket_receiver.md#schema-gcp_bucket_receiver--bucket) |
| `gcp_bucket_receiver.compression` | [gcp_bucket_receiver.compression](data-sources--global_log_receiver--properties--gcp_bucket_receiver--compression.md#section) |
| `gcp_bucket_receiver.compression.compression_default` | [gcp_bucket_receiver.compression.compression_default](data-sources--global_log_receiver--properties--gcp_bucket_receiver--compression--compression_default.md#section) |
| `gcp_bucket_receiver.compression.compression_gzip` | [gcp_bucket_receiver.compression.compression_gzip](data-sources--global_log_receiver--properties--gcp_bucket_receiver--compression--compression_gzip.md#section) |
| `gcp_bucket_receiver.compression.compression_none` | [gcp_bucket_receiver.compression.compression_none](data-sources--global_log_receiver--properties--gcp_bucket_receiver--compression--compression_none.md#section) |
| `gcp_bucket_receiver.filename_options` | [gcp_bucket_receiver.filename_options](data-sources--global_log_receiver--properties--gcp_bucket_receiver--filename_options.md#section) |
| `gcp_bucket_receiver.filename_options.custom_folder` | [gcp_bucket_receiver.filename_options.custom_folder](data-sources--global_log_receiver--properties--gcp_bucket_receiver--filename_options.md#schema-gcp_bucket_receiver--filename_options--custom_folder) |
| `gcp_bucket_receiver.filename_options.log_type_folder` | [gcp_bucket_receiver.filename_options.log_type_folder](data-sources--global_log_receiver--properties--gcp_bucket_receiver--filename_options--log_type_folder.md#section) |
| `gcp_bucket_receiver.filename_options.no_folder` | [gcp_bucket_receiver.filename_options.no_folder](data-sources--global_log_receiver--properties--gcp_bucket_receiver--filename_options--no_folder.md#section) |
| `gcp_bucket_receiver.gcp_cred` | [gcp_bucket_receiver.gcp_cred](data-sources--global_log_receiver--properties--gcp_bucket_receiver--gcp_cred.md#section) |
| `gcp_bucket_receiver.gcp_cred.name` | [gcp_bucket_receiver.gcp_cred.name](data-sources--global_log_receiver--properties--gcp_bucket_receiver--gcp_cred.md#schema-gcp_bucket_receiver--gcp_cred--name) |
| `gcp_bucket_receiver.gcp_cred.namespace` | [gcp_bucket_receiver.gcp_cred.namespace](data-sources--global_log_receiver--properties--gcp_bucket_receiver--gcp_cred.md#schema-gcp_bucket_receiver--gcp_cred--namespace) |
| `gcp_bucket_receiver.gcp_cred.tenant` | [gcp_bucket_receiver.gcp_cred.tenant](data-sources--global_log_receiver--properties--gcp_bucket_receiver--gcp_cred.md#schema-gcp_bucket_receiver--gcp_cred--tenant) |
| `http_receiver` | [http_receiver](data-sources--global_log_receiver--properties--http_receiver.md#section) |
| `http_receiver.auth_basic` | [http_receiver.auth_basic](data-sources--global_log_receiver--properties--http_receiver--auth_basic.md#section) |
| `http_receiver.auth_basic.password` | [http_receiver.auth_basic.password](data-sources--global_log_receiver--properties--http_receiver--auth_basic--password.md#section) |
| `http_receiver.auth_basic.password.blindfold_secret_info` | [http_receiver.auth_basic.password.blindfold_secret_info](data-sources--global_log_receiver--properties--http_receiver--auth_basic--password--blindfold_secret_info.md#section) |
| `http_receiver.auth_basic.password.blindfold_secret_info.decryption_provider` | [http_receiver.auth_basic.password.blindfold_secret_info.decryption_provider](data-sources--global_log_receiver--properties--http_receiver--auth_basic--password--blindfold_secret_info.md#schema-http_receiver--auth_basic--password--blindfold_secret_info--decryption_provider) |
| `http_receiver.auth_basic.password.blindfold_secret_info.location` | [http_receiver.auth_basic.password.blindfold_secret_info.location](data-sources--global_log_receiver--properties--http_receiver--auth_basic--password--blindfold_secret_info.md#schema-http_receiver--auth_basic--password--blindfold_secret_info--location) |
| `http_receiver.auth_basic.password.blindfold_secret_info.store_provider` | [http_receiver.auth_basic.password.blindfold_secret_info.store_provider](data-sources--global_log_receiver--properties--http_receiver--auth_basic--password--blindfold_secret_info.md#schema-http_receiver--auth_basic--password--blindfold_secret_info--store_provider) |
| `http_receiver.auth_basic.password.clear_secret_info` | [http_receiver.auth_basic.password.clear_secret_info](data-sources--global_log_receiver--properties--http_receiver--auth_basic--password--clear_secret_info.md#section) |
| `http_receiver.auth_basic.password.clear_secret_info.provider_ref` | [http_receiver.auth_basic.password.clear_secret_info.provider_ref](data-sources--global_log_receiver--properties--http_receiver--auth_basic--password--clear_secret_info.md#schema-http_receiver--auth_basic--password--clear_secret_info--provider_ref) |
| `http_receiver.auth_basic.password.clear_secret_info.url` | [http_receiver.auth_basic.password.clear_secret_info.url](data-sources--global_log_receiver--properties--http_receiver--auth_basic--password--clear_secret_info.md#schema-http_receiver--auth_basic--password--clear_secret_info--url) |
| `http_receiver.auth_basic.user_name` | [http_receiver.auth_basic.user_name](data-sources--global_log_receiver--properties--http_receiver--auth_basic.md#schema-http_receiver--auth_basic--user_name) |
| `http_receiver.auth_none` | [http_receiver.auth_none](data-sources--global_log_receiver--properties--http_receiver--auth_none.md#section) |
| `http_receiver.auth_token` | [http_receiver.auth_token](data-sources--global_log_receiver--properties--http_receiver--auth_token.md#section) |
| `http_receiver.auth_token.token` | [http_receiver.auth_token.token](data-sources--global_log_receiver--properties--http_receiver--auth_token--token.md#section) |
| `http_receiver.auth_token.token.blindfold_secret_info` | [http_receiver.auth_token.token.blindfold_secret_info](data-sources--global_log_receiver--properties--http_receiver--auth_token--token--blindfold_secret_info.md#section) |
| `http_receiver.auth_token.token.blindfold_secret_info.decryption_provider` | [http_receiver.auth_token.token.blindfold_secret_info.decryption_provider](data-sources--global_log_receiver--properties--http_receiver--auth_token--token--blindfold_secret_info.md#schema-http_receiver--auth_token--token--blindfold_secret_info--decryption_provider) |
| `http_receiver.auth_token.token.blindfold_secret_info.location` | [http_receiver.auth_token.token.blindfold_secret_info.location](data-sources--global_log_receiver--properties--http_receiver--auth_token--token--blindfold_secret_info.md#schema-http_receiver--auth_token--token--blindfold_secret_info--location) |
| `http_receiver.auth_token.token.blindfold_secret_info.store_provider` | [http_receiver.auth_token.token.blindfold_secret_info.store_provider](data-sources--global_log_receiver--properties--http_receiver--auth_token--token--blindfold_secret_info.md#schema-http_receiver--auth_token--token--blindfold_secret_info--store_provider) |
| `http_receiver.auth_token.token.clear_secret_info` | [http_receiver.auth_token.token.clear_secret_info](data-sources--global_log_receiver--properties--http_receiver--auth_token--token--clear_secret_info.md#section) |
| `http_receiver.auth_token.token.clear_secret_info.provider_ref` | [http_receiver.auth_token.token.clear_secret_info.provider_ref](data-sources--global_log_receiver--properties--http_receiver--auth_token--token--clear_secret_info.md#schema-http_receiver--auth_token--token--clear_secret_info--provider_ref) |
| `http_receiver.auth_token.token.clear_secret_info.url` | [http_receiver.auth_token.token.clear_secret_info.url](data-sources--global_log_receiver--properties--http_receiver--auth_token--token--clear_secret_info.md#schema-http_receiver--auth_token--token--clear_secret_info--url) |
| `http_receiver.batch` | [http_receiver.batch](data-sources--global_log_receiver--properties--http_receiver--batch.md#section) |
| `http_receiver.batch.max_bytes` | [http_receiver.batch.max_bytes](data-sources--global_log_receiver--properties--http_receiver--batch.md#schema-http_receiver--batch--max_bytes) |
| `http_receiver.batch.max_bytes_disabled` | [http_receiver.batch.max_bytes_disabled](data-sources--global_log_receiver--properties--http_receiver--batch--max_bytes_disabled.md#section) |
| `http_receiver.batch.max_events` | [http_receiver.batch.max_events](data-sources--global_log_receiver--properties--http_receiver--batch.md#schema-http_receiver--batch--max_events) |
| `http_receiver.batch.max_events_disabled` | [http_receiver.batch.max_events_disabled](data-sources--global_log_receiver--properties--http_receiver--batch--max_events_disabled.md#section) |
| `http_receiver.batch.timeout_seconds` | [http_receiver.batch.timeout_seconds](data-sources--global_log_receiver--properties--http_receiver--batch.md#schema-http_receiver--batch--timeout_seconds) |
| `http_receiver.batch.timeout_seconds_default` | [http_receiver.batch.timeout_seconds_default](data-sources--global_log_receiver--properties--http_receiver--batch--timeout_seconds_default.md#section) |
| `http_receiver.compression` | [http_receiver.compression](data-sources--global_log_receiver--properties--http_receiver--compression.md#section) |
| `http_receiver.compression.compression_default` | [http_receiver.compression.compression_default](data-sources--global_log_receiver--properties--http_receiver--compression--compression_default.md#section) |
| `http_receiver.compression.compression_gzip` | [http_receiver.compression.compression_gzip](data-sources--global_log_receiver--properties--http_receiver--compression--compression_gzip.md#section) |
| `http_receiver.compression.compression_none` | [http_receiver.compression.compression_none](data-sources--global_log_receiver--properties--http_receiver--compression--compression_none.md#section) |
| `http_receiver.no_tls` | [http_receiver.no_tls](data-sources--global_log_receiver--properties--http_receiver--no_tls.md#section) |
| `http_receiver.uri` | [http_receiver.uri](data-sources--global_log_receiver--properties--http_receiver.md#schema-http_receiver--uri) |
| `http_receiver.use_tls` | [http_receiver.use_tls](data-sources--global_log_receiver--properties--http_receiver--use_tls.md#section) |
| `http_receiver.use_tls.disable_verify_certificate` | [http_receiver.use_tls.disable_verify_certificate](data-sources--global_log_receiver--properties--http_receiver--use_tls--disable_verify_certificate.md#section) |
| `http_receiver.use_tls.disable_verify_hostname` | [http_receiver.use_tls.disable_verify_hostname](data-sources--global_log_receiver--properties--http_receiver--use_tls--disable_verify_hostname.md#section) |
| `http_receiver.use_tls.enable_verify_certificate` | [http_receiver.use_tls.enable_verify_certificate](data-sources--global_log_receiver--properties--http_receiver--use_tls--enable_verify_certificate.md#section) |
| `http_receiver.use_tls.enable_verify_hostname` | [http_receiver.use_tls.enable_verify_hostname](data-sources--global_log_receiver--properties--http_receiver--use_tls--enable_verify_hostname.md#section) |
| `http_receiver.use_tls.mtls_disabled` | [http_receiver.use_tls.mtls_disabled](data-sources--global_log_receiver--properties--http_receiver--use_tls--mtls_disabled.md#section) |
| `http_receiver.use_tls.mtls_enable` | [http_receiver.use_tls.mtls_enable](data-sources--global_log_receiver--properties--http_receiver--use_tls--mtls_enable.md#section) |
| `http_receiver.use_tls.mtls_enable.certificate` | [http_receiver.use_tls.mtls_enable.certificate](data-sources--global_log_receiver--properties--http_receiver--use_tls--mtls_enable.md#schema-http_receiver--use_tls--mtls_enable--certificate) |
| `http_receiver.use_tls.mtls_enable.key_url` | [http_receiver.use_tls.mtls_enable.key_url](data-sources--global_log_receiver--properties--http_receiver--use_tls--mtls_enable--key_url.md#section) |
| `http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info` | [http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info](data-sources--global_log_receiver--properties--http_receiver--use_tls--mtls_enable--key_url--blindfold_secret_info.md#section) |
| `http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.decryption_provider` | [http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.decryption_provider](data-sources--global_log_receiver--properties--http_receiver--use_tls--mtls_enable--key_url--blindfold_secret_info.md#schema-http_receiver--use_tls--mtls_enable--key_url--blindfold_secret_info--decryption_provider) |
| `http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.location` | [http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.location](data-sources--global_log_receiver--properties--http_receiver--use_tls--mtls_enable--key_url--blindfold_secret_info.md#schema-http_receiver--use_tls--mtls_enable--key_url--blindfold_secret_info--location) |
| `http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.store_provider` | [http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.store_provider](data-sources--global_log_receiver--properties--http_receiver--use_tls--mtls_enable--key_url--blindfold_secret_info.md#schema-http_receiver--use_tls--mtls_enable--key_url--blindfold_secret_info--store_provider) |
| `http_receiver.use_tls.mtls_enable.key_url.clear_secret_info` | [http_receiver.use_tls.mtls_enable.key_url.clear_secret_info](data-sources--global_log_receiver--properties--http_receiver--use_tls--mtls_enable--key_url--clear_secret_info.md#section) |
| `http_receiver.use_tls.mtls_enable.key_url.clear_secret_info.provider_ref` | [http_receiver.use_tls.mtls_enable.key_url.clear_secret_info.provider_ref](data-sources--global_log_receiver--properties--http_receiver--use_tls--mtls_enable--key_url--clear_secret_info.md#schema-http_receiver--use_tls--mtls_enable--key_url--clear_secret_info--provider_ref) |
| `http_receiver.use_tls.mtls_enable.key_url.clear_secret_info.url` | [http_receiver.use_tls.mtls_enable.key_url.clear_secret_info.url](data-sources--global_log_receiver--properties--http_receiver--use_tls--mtls_enable--key_url--clear_secret_info.md#schema-http_receiver--use_tls--mtls_enable--key_url--clear_secret_info--url) |
| `http_receiver.use_tls.no_ca` | [http_receiver.use_tls.no_ca](data-sources--global_log_receiver--properties--http_receiver--use_tls--no_ca.md#section) |
| `http_receiver.use_tls.trusted_ca_url` | [http_receiver.use_tls.trusted_ca_url](data-sources--global_log_receiver--properties--http_receiver--use_tls.md#schema-http_receiver--use_tls--trusted_ca_url) |
| `id` | [id](data-sources--global_log_receiver--reference.md#schema-id) |
| `kafka_receiver` | [kafka_receiver](data-sources--global_log_receiver--properties--kafka_receiver.md#section) |
| `kafka_receiver.batch` | [kafka_receiver.batch](data-sources--global_log_receiver--properties--kafka_receiver--batch.md#section) |
| `kafka_receiver.batch.max_bytes` | [kafka_receiver.batch.max_bytes](data-sources--global_log_receiver--properties--kafka_receiver--batch.md#schema-kafka_receiver--batch--max_bytes) |
| `kafka_receiver.batch.max_bytes_disabled` | [kafka_receiver.batch.max_bytes_disabled](data-sources--global_log_receiver--properties--kafka_receiver--batch--max_bytes_disabled.md#section) |
| `kafka_receiver.batch.max_events` | [kafka_receiver.batch.max_events](data-sources--global_log_receiver--properties--kafka_receiver--batch.md#schema-kafka_receiver--batch--max_events) |
| `kafka_receiver.batch.max_events_disabled` | [kafka_receiver.batch.max_events_disabled](data-sources--global_log_receiver--properties--kafka_receiver--batch--max_events_disabled.md#section) |
| `kafka_receiver.batch.timeout_seconds` | [kafka_receiver.batch.timeout_seconds](data-sources--global_log_receiver--properties--kafka_receiver--batch.md#schema-kafka_receiver--batch--timeout_seconds) |
| `kafka_receiver.batch.timeout_seconds_default` | [kafka_receiver.batch.timeout_seconds_default](data-sources--global_log_receiver--properties--kafka_receiver--batch--timeout_seconds_default.md#section) |
| `kafka_receiver.bootstrap_servers` | [kafka_receiver.bootstrap_servers](data-sources--global_log_receiver--properties--kafka_receiver.md#schema-kafka_receiver--bootstrap_servers) |
| `kafka_receiver.compression` | [kafka_receiver.compression](data-sources--global_log_receiver--properties--kafka_receiver--compression.md#section) |
| `kafka_receiver.compression.compression_default` | [kafka_receiver.compression.compression_default](data-sources--global_log_receiver--properties--kafka_receiver--compression--compression_default.md#section) |
| `kafka_receiver.compression.compression_gzip` | [kafka_receiver.compression.compression_gzip](data-sources--global_log_receiver--properties--kafka_receiver--compression--compression_gzip.md#section) |
| `kafka_receiver.compression.compression_none` | [kafka_receiver.compression.compression_none](data-sources--global_log_receiver--properties--kafka_receiver--compression--compression_none.md#section) |
| `kafka_receiver.kafka_topic` | [kafka_receiver.kafka_topic](data-sources--global_log_receiver--properties--kafka_receiver.md#schema-kafka_receiver--kafka_topic) |
| `kafka_receiver.no_tls` | [kafka_receiver.no_tls](data-sources--global_log_receiver--properties--kafka_receiver--no_tls.md#section) |
| `kafka_receiver.use_tls` | [kafka_receiver.use_tls](data-sources--global_log_receiver--properties--kafka_receiver--use_tls.md#section) |
| `kafka_receiver.use_tls.disable_verify_certificate` | [kafka_receiver.use_tls.disable_verify_certificate](data-sources--global_log_receiver--properties--kafka_receiver--use_tls--disable_verify_certificate.md#section) |
| `kafka_receiver.use_tls.disable_verify_hostname` | [kafka_receiver.use_tls.disable_verify_hostname](data-sources--global_log_receiver--properties--kafka_receiver--use_tls--disable_verify_hostname.md#section) |
| `kafka_receiver.use_tls.enable_verify_certificate` | [kafka_receiver.use_tls.enable_verify_certificate](data-sources--global_log_receiver--properties--kafka_receiver--use_tls--enable_verify_certificate.md#section) |
| `kafka_receiver.use_tls.enable_verify_hostname` | [kafka_receiver.use_tls.enable_verify_hostname](data-sources--global_log_receiver--properties--kafka_receiver--use_tls--enable_verify_hostname.md#section) |
| `kafka_receiver.use_tls.mtls_disabled` | [kafka_receiver.use_tls.mtls_disabled](data-sources--global_log_receiver--properties--kafka_receiver--use_tls--mtls_disabled.md#section) |
| `kafka_receiver.use_tls.mtls_enable` | [kafka_receiver.use_tls.mtls_enable](data-sources--global_log_receiver--properties--kafka_receiver--use_tls--mtls_enable.md#section) |
| `kafka_receiver.use_tls.mtls_enable.certificate` | [kafka_receiver.use_tls.mtls_enable.certificate](data-sources--global_log_receiver--properties--kafka_receiver--use_tls--mtls_enable.md#schema-kafka_receiver--use_tls--mtls_enable--certificate) |
| `kafka_receiver.use_tls.mtls_enable.key_url` | [kafka_receiver.use_tls.mtls_enable.key_url](data-sources--global_log_receiver--properties--kafka_receiver--use_tls--mtls_enable--key_url.md#section) |
| `kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info` | [kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info](data-sources--global_log_receiver--properties--kafka_receiver--use_tls--mtls_enable--key_url--blindfold_secret_info.md#section) |
| `kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.decryption_provider` | [kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.decryption_provider](data-sources--global_log_receiver--properties--kafka_receiver--use_tls--mtls_enable--key_url--blindfold_secret_info.md#schema-kafka_receiver--use_tls--mtls_enable--key_url--blindfold_secret_info--decryption_provider) |
| `kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.location` | [kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.location](data-sources--global_log_receiver--properties--kafka_receiver--use_tls--mtls_enable--key_url--blindfold_secret_info.md#schema-kafka_receiver--use_tls--mtls_enable--key_url--blindfold_secret_info--location) |
| `kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.store_provider` | [kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.store_provider](data-sources--global_log_receiver--properties--kafka_receiver--use_tls--mtls_enable--key_url--blindfold_secret_info.md#schema-kafka_receiver--use_tls--mtls_enable--key_url--blindfold_secret_info--store_provider) |
| `kafka_receiver.use_tls.mtls_enable.key_url.clear_secret_info` | [kafka_receiver.use_tls.mtls_enable.key_url.clear_secret_info](data-sources--global_log_receiver--properties--kafka_receiver--use_tls--mtls_enable--key_url--clear_secret_info.md#section) |
| `kafka_receiver.use_tls.mtls_enable.key_url.clear_secret_info.provider_ref` | [kafka_receiver.use_tls.mtls_enable.key_url.clear_secret_info.provider_ref](data-sources--global_log_receiver--properties--kafka_receiver--use_tls--mtls_enable--key_url--clear_secret_info.md#schema-kafka_receiver--use_tls--mtls_enable--key_url--clear_secret_info--provider_ref) |
| `kafka_receiver.use_tls.mtls_enable.key_url.clear_secret_info.url` | [kafka_receiver.use_tls.mtls_enable.key_url.clear_secret_info.url](data-sources--global_log_receiver--properties--kafka_receiver--use_tls--mtls_enable--key_url--clear_secret_info.md#schema-kafka_receiver--use_tls--mtls_enable--key_url--clear_secret_info--url) |
| `kafka_receiver.use_tls.no_ca` | [kafka_receiver.use_tls.no_ca](data-sources--global_log_receiver--properties--kafka_receiver--use_tls--no_ca.md#section) |
| `kafka_receiver.use_tls.trusted_ca_url` | [kafka_receiver.use_tls.trusted_ca_url](data-sources--global_log_receiver--properties--kafka_receiver--use_tls.md#schema-kafka_receiver--use_tls--trusted_ca_url) |
| `labels` | [labels](data-sources--global_log_receiver--reference.md#schema-labels) |
| `name` | [name](data-sources--global_log_receiver--reference.md#schema-name) |
| `namespace` | [namespace](data-sources--global_log_receiver--reference.md#schema-namespace) |
| `new_relic_receiver` | [new_relic_receiver](data-sources--global_log_receiver--properties--new_relic_receiver.md#section) |
| `new_relic_receiver.api_key` | [new_relic_receiver.api_key](data-sources--global_log_receiver--properties--new_relic_receiver--api_key.md#section) |
| `new_relic_receiver.api_key.blindfold_secret_info` | [new_relic_receiver.api_key.blindfold_secret_info](data-sources--global_log_receiver--properties--new_relic_receiver--api_key--blindfold_secret_info.md#section) |
| `new_relic_receiver.api_key.blindfold_secret_info.decryption_provider` | [new_relic_receiver.api_key.blindfold_secret_info.decryption_provider](data-sources--global_log_receiver--properties--new_relic_receiver--api_key--blindfold_secret_info.md#schema-new_relic_receiver--api_key--blindfold_secret_info--decryption_provider) |
| `new_relic_receiver.api_key.blindfold_secret_info.location` | [new_relic_receiver.api_key.blindfold_secret_info.location](data-sources--global_log_receiver--properties--new_relic_receiver--api_key--blindfold_secret_info.md#schema-new_relic_receiver--api_key--blindfold_secret_info--location) |
| `new_relic_receiver.api_key.blindfold_secret_info.store_provider` | [new_relic_receiver.api_key.blindfold_secret_info.store_provider](data-sources--global_log_receiver--properties--new_relic_receiver--api_key--blindfold_secret_info.md#schema-new_relic_receiver--api_key--blindfold_secret_info--store_provider) |
| `new_relic_receiver.api_key.clear_secret_info` | [new_relic_receiver.api_key.clear_secret_info](data-sources--global_log_receiver--properties--new_relic_receiver--api_key--clear_secret_info.md#section) |
| `new_relic_receiver.api_key.clear_secret_info.provider_ref` | [new_relic_receiver.api_key.clear_secret_info.provider_ref](data-sources--global_log_receiver--properties--new_relic_receiver--api_key--clear_secret_info.md#schema-new_relic_receiver--api_key--clear_secret_info--provider_ref) |
| `new_relic_receiver.api_key.clear_secret_info.url` | [new_relic_receiver.api_key.clear_secret_info.url](data-sources--global_log_receiver--properties--new_relic_receiver--api_key--clear_secret_info.md#schema-new_relic_receiver--api_key--clear_secret_info--url) |
| `new_relic_receiver.eu` | [new_relic_receiver.eu](data-sources--global_log_receiver--properties--new_relic_receiver--eu.md#section) |
| `new_relic_receiver.us` | [new_relic_receiver.us](data-sources--global_log_receiver--properties--new_relic_receiver--us.md#section) |
| `ns_all` | [ns_all](data-sources--global_log_receiver--properties--ns_all.md#section) |
| `ns_current` | [ns_current](data-sources--global_log_receiver--properties--ns_current.md#section) |
| `ns_list` | [ns_list](data-sources--global_log_receiver--properties--ns_list.md#section) |
| `ns_list.namespaces` | [ns_list.namespaces](data-sources--global_log_receiver--properties--ns_list.md#schema-ns_list--namespaces) |
| `qradar_receiver` | [qradar_receiver](data-sources--global_log_receiver--properties--qradar_receiver.md#section) |
| `qradar_receiver.batch` | [qradar_receiver.batch](data-sources--global_log_receiver--properties--qradar_receiver--batch.md#section) |
| `qradar_receiver.batch.max_bytes` | [qradar_receiver.batch.max_bytes](data-sources--global_log_receiver--properties--qradar_receiver--batch.md#schema-qradar_receiver--batch--max_bytes) |
| `qradar_receiver.batch.max_bytes_disabled` | [qradar_receiver.batch.max_bytes_disabled](data-sources--global_log_receiver--properties--qradar_receiver--batch--max_bytes_disabled.md#section) |
| `qradar_receiver.batch.max_events` | [qradar_receiver.batch.max_events](data-sources--global_log_receiver--properties--qradar_receiver--batch.md#schema-qradar_receiver--batch--max_events) |
| `qradar_receiver.batch.max_events_disabled` | [qradar_receiver.batch.max_events_disabled](data-sources--global_log_receiver--properties--qradar_receiver--batch--max_events_disabled.md#section) |
| `qradar_receiver.batch.timeout_seconds` | [qradar_receiver.batch.timeout_seconds](data-sources--global_log_receiver--properties--qradar_receiver--batch.md#schema-qradar_receiver--batch--timeout_seconds) |
| `qradar_receiver.batch.timeout_seconds_default` | [qradar_receiver.batch.timeout_seconds_default](data-sources--global_log_receiver--properties--qradar_receiver--batch--timeout_seconds_default.md#section) |
| `qradar_receiver.compression` | [qradar_receiver.compression](data-sources--global_log_receiver--properties--qradar_receiver--compression.md#section) |
| `qradar_receiver.compression.compression_default` | [qradar_receiver.compression.compression_default](data-sources--global_log_receiver--properties--qradar_receiver--compression--compression_default.md#section) |
| `qradar_receiver.compression.compression_gzip` | [qradar_receiver.compression.compression_gzip](data-sources--global_log_receiver--properties--qradar_receiver--compression--compression_gzip.md#section) |
| `qradar_receiver.compression.compression_none` | [qradar_receiver.compression.compression_none](data-sources--global_log_receiver--properties--qradar_receiver--compression--compression_none.md#section) |
| `qradar_receiver.no_tls` | [qradar_receiver.no_tls](data-sources--global_log_receiver--properties--qradar_receiver--no_tls.md#section) |
| `qradar_receiver.uri` | [qradar_receiver.uri](data-sources--global_log_receiver--properties--qradar_receiver.md#schema-qradar_receiver--uri) |
| `qradar_receiver.use_tls` | [qradar_receiver.use_tls](data-sources--global_log_receiver--properties--qradar_receiver--use_tls.md#section) |
| `qradar_receiver.use_tls.disable_verify_certificate` | [qradar_receiver.use_tls.disable_verify_certificate](data-sources--global_log_receiver--properties--qradar_receiver--use_tls--disable_verify_certificate.md#section) |
| `qradar_receiver.use_tls.disable_verify_hostname` | [qradar_receiver.use_tls.disable_verify_hostname](data-sources--global_log_receiver--properties--qradar_receiver--use_tls--disable_verify_hostname.md#section) |
| `qradar_receiver.use_tls.enable_verify_certificate` | [qradar_receiver.use_tls.enable_verify_certificate](data-sources--global_log_receiver--properties--qradar_receiver--use_tls--enable_verify_certificate.md#section) |
| `qradar_receiver.use_tls.enable_verify_hostname` | [qradar_receiver.use_tls.enable_verify_hostname](data-sources--global_log_receiver--properties--qradar_receiver--use_tls--enable_verify_hostname.md#section) |
| `qradar_receiver.use_tls.mtls_disabled` | [qradar_receiver.use_tls.mtls_disabled](data-sources--global_log_receiver--properties--qradar_receiver--use_tls--mtls_disabled.md#section) |
| `qradar_receiver.use_tls.mtls_enable` | [qradar_receiver.use_tls.mtls_enable](data-sources--global_log_receiver--properties--qradar_receiver--use_tls--mtls_enable.md#section) |
| `qradar_receiver.use_tls.mtls_enable.certificate` | [qradar_receiver.use_tls.mtls_enable.certificate](data-sources--global_log_receiver--properties--qradar_receiver--use_tls--mtls_enable.md#schema-qradar_receiver--use_tls--mtls_enable--certificate) |
| `qradar_receiver.use_tls.mtls_enable.key_url` | [qradar_receiver.use_tls.mtls_enable.key_url](data-sources--global_log_receiver--properties--qradar_receiver--use_tls--mtls_enable--key_url.md#section) |
| `qradar_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info` | [qradar_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info](data-sources--global_log_receiver--properties--qradar_receiver--use_tls--mtls_enable--key_url--blindfold_secret_info.md#section) |
| `qradar_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.decryption_provider` | [qradar_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.decryption_provider](data-sources--global_log_receiver--properties--qradar_receiver--use_tls--mtls_enable--key_url--blindfold_secret_info.md#schema-qradar_receiver--use_tls--mtls_enable--key_url--blindfold_secret_info--decryption_provider) |
| `qradar_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.location` | [qradar_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.location](data-sources--global_log_receiver--properties--qradar_receiver--use_tls--mtls_enable--key_url--blindfold_secret_info.md#schema-qradar_receiver--use_tls--mtls_enable--key_url--blindfold_secret_info--location) |
| `qradar_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.store_provider` | [qradar_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.store_provider](data-sources--global_log_receiver--properties--qradar_receiver--use_tls--mtls_enable--key_url--blindfold_secret_info.md#schema-qradar_receiver--use_tls--mtls_enable--key_url--blindfold_secret_info--store_provider) |
| `qradar_receiver.use_tls.mtls_enable.key_url.clear_secret_info` | [qradar_receiver.use_tls.mtls_enable.key_url.clear_secret_info](data-sources--global_log_receiver--properties--qradar_receiver--use_tls--mtls_enable--key_url--clear_secret_info.md#section) |
| `qradar_receiver.use_tls.mtls_enable.key_url.clear_secret_info.provider_ref` | [qradar_receiver.use_tls.mtls_enable.key_url.clear_secret_info.provider_ref](data-sources--global_log_receiver--properties--qradar_receiver--use_tls--mtls_enable--key_url--clear_secret_info.md#schema-qradar_receiver--use_tls--mtls_enable--key_url--clear_secret_info--provider_ref) |
| `qradar_receiver.use_tls.mtls_enable.key_url.clear_secret_info.url` | [qradar_receiver.use_tls.mtls_enable.key_url.clear_secret_info.url](data-sources--global_log_receiver--properties--qradar_receiver--use_tls--mtls_enable--key_url--clear_secret_info.md#schema-qradar_receiver--use_tls--mtls_enable--key_url--clear_secret_info--url) |
| `qradar_receiver.use_tls.no_ca` | [qradar_receiver.use_tls.no_ca](data-sources--global_log_receiver--properties--qradar_receiver--use_tls--no_ca.md#section) |
| `qradar_receiver.use_tls.trusted_ca_url` | [qradar_receiver.use_tls.trusted_ca_url](data-sources--global_log_receiver--properties--qradar_receiver--use_tls.md#schema-qradar_receiver--use_tls--trusted_ca_url) |
| `request_logs` | [request_logs](data-sources--global_log_receiver--properties--request_logs.md#section) |
| `request_logs.sampled` | [request_logs.sampled](data-sources--global_log_receiver--properties--request_logs--sampled.md#section) |
| `request_logs.unsampled` | [request_logs.unsampled](data-sources--global_log_receiver--properties--request_logs--unsampled.md#section) |
| `s3_receiver` | [s3_receiver](data-sources--global_log_receiver--properties--s3_receiver.md#section) |
| `s3_receiver.aws_cred` | [s3_receiver.aws_cred](data-sources--global_log_receiver--properties--s3_receiver--aws_cred.md#section) |
| `s3_receiver.aws_cred.name` | [s3_receiver.aws_cred.name](data-sources--global_log_receiver--properties--s3_receiver--aws_cred.md#schema-s3_receiver--aws_cred--name) |
| `s3_receiver.aws_cred.namespace` | [s3_receiver.aws_cred.namespace](data-sources--global_log_receiver--properties--s3_receiver--aws_cred.md#schema-s3_receiver--aws_cred--namespace) |
| `s3_receiver.aws_cred.tenant` | [s3_receiver.aws_cred.tenant](data-sources--global_log_receiver--properties--s3_receiver--aws_cred.md#schema-s3_receiver--aws_cred--tenant) |
| `s3_receiver.aws_region` | [s3_receiver.aws_region](data-sources--global_log_receiver--properties--s3_receiver.md#schema-s3_receiver--aws_region) |
| `s3_receiver.batch` | [s3_receiver.batch](data-sources--global_log_receiver--properties--s3_receiver--batch.md#section) |
| `s3_receiver.batch.max_bytes` | [s3_receiver.batch.max_bytes](data-sources--global_log_receiver--properties--s3_receiver--batch.md#schema-s3_receiver--batch--max_bytes) |
| `s3_receiver.batch.max_bytes_disabled` | [s3_receiver.batch.max_bytes_disabled](data-sources--global_log_receiver--properties--s3_receiver--batch--max_bytes_disabled.md#section) |
| `s3_receiver.batch.max_events` | [s3_receiver.batch.max_events](data-sources--global_log_receiver--properties--s3_receiver--batch.md#schema-s3_receiver--batch--max_events) |
| `s3_receiver.batch.max_events_disabled` | [s3_receiver.batch.max_events_disabled](data-sources--global_log_receiver--properties--s3_receiver--batch--max_events_disabled.md#section) |
| `s3_receiver.batch.timeout_seconds` | [s3_receiver.batch.timeout_seconds](data-sources--global_log_receiver--properties--s3_receiver--batch.md#schema-s3_receiver--batch--timeout_seconds) |
| `s3_receiver.batch.timeout_seconds_default` | [s3_receiver.batch.timeout_seconds_default](data-sources--global_log_receiver--properties--s3_receiver--batch--timeout_seconds_default.md#section) |
| `s3_receiver.bucket` | [s3_receiver.bucket](data-sources--global_log_receiver--properties--s3_receiver.md#schema-s3_receiver--bucket) |
| `s3_receiver.compression` | [s3_receiver.compression](data-sources--global_log_receiver--properties--s3_receiver--compression.md#section) |
| `s3_receiver.compression.compression_default` | [s3_receiver.compression.compression_default](data-sources--global_log_receiver--properties--s3_receiver--compression--compression_default.md#section) |
| `s3_receiver.compression.compression_gzip` | [s3_receiver.compression.compression_gzip](data-sources--global_log_receiver--properties--s3_receiver--compression--compression_gzip.md#section) |
| `s3_receiver.compression.compression_none` | [s3_receiver.compression.compression_none](data-sources--global_log_receiver--properties--s3_receiver--compression--compression_none.md#section) |
| `s3_receiver.filename_options` | [s3_receiver.filename_options](data-sources--global_log_receiver--properties--s3_receiver--filename_options.md#section) |
| `s3_receiver.filename_options.custom_folder` | [s3_receiver.filename_options.custom_folder](data-sources--global_log_receiver--properties--s3_receiver--filename_options.md#schema-s3_receiver--filename_options--custom_folder) |
| `s3_receiver.filename_options.log_type_folder` | [s3_receiver.filename_options.log_type_folder](data-sources--global_log_receiver--properties--s3_receiver--filename_options--log_type_folder.md#section) |
| `s3_receiver.filename_options.no_folder` | [s3_receiver.filename_options.no_folder](data-sources--global_log_receiver--properties--s3_receiver--filename_options--no_folder.md#section) |
| `security_events` | [security_events](data-sources--global_log_receiver--properties--security_events.md#section) |
| `splunk_receiver` | [splunk_receiver](data-sources--global_log_receiver--properties--splunk_receiver.md#section) |
| `splunk_receiver.batch` | [splunk_receiver.batch](data-sources--global_log_receiver--properties--splunk_receiver--batch.md#section) |
| `splunk_receiver.batch.max_bytes` | [splunk_receiver.batch.max_bytes](data-sources--global_log_receiver--properties--splunk_receiver--batch.md#schema-splunk_receiver--batch--max_bytes) |
| `splunk_receiver.batch.max_bytes_disabled` | [splunk_receiver.batch.max_bytes_disabled](data-sources--global_log_receiver--properties--splunk_receiver--batch--max_bytes_disabled.md#section) |
| `splunk_receiver.batch.max_events` | [splunk_receiver.batch.max_events](data-sources--global_log_receiver--properties--splunk_receiver--batch.md#schema-splunk_receiver--batch--max_events) |
| `splunk_receiver.batch.max_events_disabled` | [splunk_receiver.batch.max_events_disabled](data-sources--global_log_receiver--properties--splunk_receiver--batch--max_events_disabled.md#section) |
| `splunk_receiver.batch.timeout_seconds` | [splunk_receiver.batch.timeout_seconds](data-sources--global_log_receiver--properties--splunk_receiver--batch.md#schema-splunk_receiver--batch--timeout_seconds) |
| `splunk_receiver.batch.timeout_seconds_default` | [splunk_receiver.batch.timeout_seconds_default](data-sources--global_log_receiver--properties--splunk_receiver--batch--timeout_seconds_default.md#section) |
| `splunk_receiver.compression` | [splunk_receiver.compression](data-sources--global_log_receiver--properties--splunk_receiver--compression.md#section) |
| `splunk_receiver.compression.compression_default` | [splunk_receiver.compression.compression_default](data-sources--global_log_receiver--properties--splunk_receiver--compression--compression_default.md#section) |
| `splunk_receiver.compression.compression_gzip` | [splunk_receiver.compression.compression_gzip](data-sources--global_log_receiver--properties--splunk_receiver--compression--compression_gzip.md#section) |
| `splunk_receiver.compression.compression_none` | [splunk_receiver.compression.compression_none](data-sources--global_log_receiver--properties--splunk_receiver--compression--compression_none.md#section) |
| `splunk_receiver.endpoint` | [splunk_receiver.endpoint](data-sources--global_log_receiver--properties--splunk_receiver.md#schema-splunk_receiver--endpoint) |
| `splunk_receiver.no_tls` | [splunk_receiver.no_tls](data-sources--global_log_receiver--properties--splunk_receiver--no_tls.md#section) |
| `splunk_receiver.splunk_hec_token` | [splunk_receiver.splunk_hec_token](data-sources--global_log_receiver--properties--splunk_receiver--splunk_hec_token.md#section) |
| `splunk_receiver.splunk_hec_token.blindfold_secret_info` | [splunk_receiver.splunk_hec_token.blindfold_secret_info](data-sources--global_log_receiver--properties--splunk_receiver--splunk_hec_token--blindfold_secret_info.md#section) |
| `splunk_receiver.splunk_hec_token.blindfold_secret_info.decryption_provider` | [splunk_receiver.splunk_hec_token.blindfold_secret_info.decryption_provider](data-sources--global_log_receiver--properties--splunk_receiver--splunk_hec_token--blindfold_secret_info.md#schema-splunk_receiver--splunk_hec_token--blindfold_secret_info--decryption_provider) |
| `splunk_receiver.splunk_hec_token.blindfold_secret_info.location` | [splunk_receiver.splunk_hec_token.blindfold_secret_info.location](data-sources--global_log_receiver--properties--splunk_receiver--splunk_hec_token--blindfold_secret_info.md#schema-splunk_receiver--splunk_hec_token--blindfold_secret_info--location) |
| `splunk_receiver.splunk_hec_token.blindfold_secret_info.store_provider` | [splunk_receiver.splunk_hec_token.blindfold_secret_info.store_provider](data-sources--global_log_receiver--properties--splunk_receiver--splunk_hec_token--blindfold_secret_info.md#schema-splunk_receiver--splunk_hec_token--blindfold_secret_info--store_provider) |
| `splunk_receiver.splunk_hec_token.clear_secret_info` | [splunk_receiver.splunk_hec_token.clear_secret_info](data-sources--global_log_receiver--properties--splunk_receiver--splunk_hec_token--clear_secret_info.md#section) |
| `splunk_receiver.splunk_hec_token.clear_secret_info.provider_ref` | [splunk_receiver.splunk_hec_token.clear_secret_info.provider_ref](data-sources--global_log_receiver--properties--splunk_receiver--splunk_hec_token--clear_secret_info.md#schema-splunk_receiver--splunk_hec_token--clear_secret_info--provider_ref) |
| `splunk_receiver.splunk_hec_token.clear_secret_info.url` | [splunk_receiver.splunk_hec_token.clear_secret_info.url](data-sources--global_log_receiver--properties--splunk_receiver--splunk_hec_token--clear_secret_info.md#schema-splunk_receiver--splunk_hec_token--clear_secret_info--url) |
| `splunk_receiver.use_tls` | [splunk_receiver.use_tls](data-sources--global_log_receiver--properties--splunk_receiver--use_tls.md#section) |
| `splunk_receiver.use_tls.disable_verify_certificate` | [splunk_receiver.use_tls.disable_verify_certificate](data-sources--global_log_receiver--properties--splunk_receiver--use_tls--disable_verify_certificate.md#section) |
| `splunk_receiver.use_tls.disable_verify_hostname` | [splunk_receiver.use_tls.disable_verify_hostname](data-sources--global_log_receiver--properties--splunk_receiver--use_tls--disable_verify_hostname.md#section) |
| `splunk_receiver.use_tls.enable_verify_certificate` | [splunk_receiver.use_tls.enable_verify_certificate](data-sources--global_log_receiver--properties--splunk_receiver--use_tls--enable_verify_certificate.md#section) |
| `splunk_receiver.use_tls.enable_verify_hostname` | [splunk_receiver.use_tls.enable_verify_hostname](data-sources--global_log_receiver--properties--splunk_receiver--use_tls--enable_verify_hostname.md#section) |
| `splunk_receiver.use_tls.mtls_disabled` | [splunk_receiver.use_tls.mtls_disabled](data-sources--global_log_receiver--properties--splunk_receiver--use_tls--mtls_disabled.md#section) |
| `splunk_receiver.use_tls.mtls_enable` | [splunk_receiver.use_tls.mtls_enable](data-sources--global_log_receiver--properties--splunk_receiver--use_tls--mtls_enable.md#section) |
| `splunk_receiver.use_tls.mtls_enable.certificate` | [splunk_receiver.use_tls.mtls_enable.certificate](data-sources--global_log_receiver--properties--splunk_receiver--use_tls--mtls_enable.md#schema-splunk_receiver--use_tls--mtls_enable--certificate) |
| `splunk_receiver.use_tls.mtls_enable.key_url` | [splunk_receiver.use_tls.mtls_enable.key_url](data-sources--global_log_receiver--properties--splunk_receiver--use_tls--mtls_enable--key_url.md#section) |
| `splunk_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info` | [splunk_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info](data-sources--global_log_receiver--properties--splunk_receiver--use_tls--mtls_enable--key_url--blindfold_secret_info.md#section) |
| `splunk_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.decryption_provider` | [splunk_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.decryption_provider](data-sources--global_log_receiver--properties--splunk_receiver--use_tls--mtls_enable--key_url--blindfold_secret_info.md#schema-splunk_receiver--use_tls--mtls_enable--key_url--blindfold_secret_info--decryption_provider) |
| `splunk_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.location` | [splunk_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.location](data-sources--global_log_receiver--properties--splunk_receiver--use_tls--mtls_enable--key_url--blindfold_secret_info.md#schema-splunk_receiver--use_tls--mtls_enable--key_url--blindfold_secret_info--location) |
| `splunk_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.store_provider` | [splunk_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.store_provider](data-sources--global_log_receiver--properties--splunk_receiver--use_tls--mtls_enable--key_url--blindfold_secret_info.md#schema-splunk_receiver--use_tls--mtls_enable--key_url--blindfold_secret_info--store_provider) |
| `splunk_receiver.use_tls.mtls_enable.key_url.clear_secret_info` | [splunk_receiver.use_tls.mtls_enable.key_url.clear_secret_info](data-sources--global_log_receiver--properties--splunk_receiver--use_tls--mtls_enable--key_url--clear_secret_info.md#section) |
| `splunk_receiver.use_tls.mtls_enable.key_url.clear_secret_info.provider_ref` | [splunk_receiver.use_tls.mtls_enable.key_url.clear_secret_info.provider_ref](data-sources--global_log_receiver--properties--splunk_receiver--use_tls--mtls_enable--key_url--clear_secret_info.md#schema-splunk_receiver--use_tls--mtls_enable--key_url--clear_secret_info--provider_ref) |
| `splunk_receiver.use_tls.mtls_enable.key_url.clear_secret_info.url` | [splunk_receiver.use_tls.mtls_enable.key_url.clear_secret_info.url](data-sources--global_log_receiver--properties--splunk_receiver--use_tls--mtls_enable--key_url--clear_secret_info.md#schema-splunk_receiver--use_tls--mtls_enable--key_url--clear_secret_info--url) |
| `splunk_receiver.use_tls.no_ca` | [splunk_receiver.use_tls.no_ca](data-sources--global_log_receiver--properties--splunk_receiver--use_tls--no_ca.md#section) |
| `splunk_receiver.use_tls.trusted_ca_url` | [splunk_receiver.use_tls.trusted_ca_url](data-sources--global_log_receiver--properties--splunk_receiver--use_tls.md#schema-splunk_receiver--use_tls--trusted_ca_url) |
| `sumo_logic_receiver` | [sumo_logic_receiver](data-sources--global_log_receiver--properties--sumo_logic_receiver.md#section) |
| `sumo_logic_receiver.url` | [sumo_logic_receiver.url](data-sources--global_log_receiver--properties--sumo_logic_receiver--url.md#section) |
| `sumo_logic_receiver.url.blindfold_secret_info` | [sumo_logic_receiver.url.blindfold_secret_info](data-sources--global_log_receiver--properties--sumo_logic_receiver--url--blindfold_secret_info.md#section) |
| `sumo_logic_receiver.url.blindfold_secret_info.decryption_provider` | [sumo_logic_receiver.url.blindfold_secret_info.decryption_provider](data-sources--global_log_receiver--properties--sumo_logic_receiver--url--blindfold_secret_info.md#schema-sumo_logic_receiver--url--blindfold_secret_info--decryption_provider) |
| `sumo_logic_receiver.url.blindfold_secret_info.location` | [sumo_logic_receiver.url.blindfold_secret_info.location](data-sources--global_log_receiver--properties--sumo_logic_receiver--url--blindfold_secret_info.md#schema-sumo_logic_receiver--url--blindfold_secret_info--location) |
| `sumo_logic_receiver.url.blindfold_secret_info.store_provider` | [sumo_logic_receiver.url.blindfold_secret_info.store_provider](data-sources--global_log_receiver--properties--sumo_logic_receiver--url--blindfold_secret_info.md#schema-sumo_logic_receiver--url--blindfold_secret_info--store_provider) |
| `sumo_logic_receiver.url.clear_secret_info` | [sumo_logic_receiver.url.clear_secret_info](data-sources--global_log_receiver--properties--sumo_logic_receiver--url--clear_secret_info.md#section) |
| `sumo_logic_receiver.url.clear_secret_info.provider_ref` | [sumo_logic_receiver.url.clear_secret_info.provider_ref](data-sources--global_log_receiver--properties--sumo_logic_receiver--url--clear_secret_info.md#schema-sumo_logic_receiver--url--clear_secret_info--provider_ref) |
| `sumo_logic_receiver.url.clear_secret_info.url` | [sumo_logic_receiver.url.clear_secret_info.url](data-sources--global_log_receiver--properties--sumo_logic_receiver--url--clear_secret_info.md#schema-sumo_logic_receiver--url--clear_secret_info--url) |

## Next pages

- [audit_logs](data-sources--global_log_receiver--properties--audit_logs.md)
- [aws_cloud_watch_receiver](data-sources--global_log_receiver--properties--aws_cloud_watch_receiver.md)
- [azure_event_hubs_receiver](data-sources--global_log_receiver--properties--azure_event_hubs_receiver.md)
- [azure_receiver](data-sources--global_log_receiver--properties--azure_receiver.md)
- [datadog_receiver](data-sources--global_log_receiver--properties--datadog_receiver.md)
- [dns_logs](data-sources--global_log_receiver--properties--dns_logs.md)
- [gcp_bucket_receiver](data-sources--global_log_receiver--properties--gcp_bucket_receiver.md)
- [http_receiver](data-sources--global_log_receiver--properties--http_receiver.md)
- [kafka_receiver](data-sources--global_log_receiver--properties--kafka_receiver.md)
- [new_relic_receiver](data-sources--global_log_receiver--properties--new_relic_receiver.md)
- [ns_all](data-sources--global_log_receiver--properties--ns_all.md)
- [ns_current](data-sources--global_log_receiver--properties--ns_current.md)
- [ns_list](data-sources--global_log_receiver--properties--ns_list.md)
- [qradar_receiver](data-sources--global_log_receiver--properties--qradar_receiver.md)
- [request_logs](data-sources--global_log_receiver--properties--request_logs.md)
- [s3_receiver](data-sources--global_log_receiver--properties--s3_receiver.md)
- [security_events](data-sources--global_log_receiver--properties--security_events.md)
- [splunk_receiver](data-sources--global_log_receiver--properties--splunk_receiver.md)
- [sumo_logic_receiver](data-sources--global_log_receiver--properties--sumo_logic_receiver.md)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md)
