---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_global_log_receiver."
xcsh_docs: {"aliases": ["global log receiver"], "body_bytes": 101336, "body_sha256": "sha256:1e080960e3b891e5c813ac891ce570268c362ff2df7fcc3ca768620f6ad3d26d", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:resources:global_log_receiver:properties:audit_logs", "xcsh-docs:resources:global_log_receiver:properties:aws_cloud_watch_receiver", "xcsh-docs:resources:global_log_receiver:properties:azure_event_hubs_receiver", "xcsh-docs:resources:global_log_receiver:properties:azure_receiver", "xcsh-docs:resources:global_log_receiver:properties:datadog_receiver", "xcsh-docs:resources:global_log_receiver:properties:dns_logs", "xcsh-docs:resources:global_log_receiver:properties:gcp_bucket_receiver", "xcsh-docs:resources:global_log_receiver:properties:http_receiver", "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver", "xcsh-docs:resources:global_log_receiver:properties:new_relic_receiver", "xcsh-docs:resources:global_log_receiver:properties:ns_all", "xcsh-docs:resources:global_log_receiver:properties:ns_current", "xcsh-docs:resources:global_log_receiver:properties:ns_list", "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver", "xcsh-docs:resources:global_log_receiver:properties:request_logs", "xcsh-docs:resources:global_log_receiver:properties:s3_receiver", "xcsh-docs:resources:global_log_receiver:properties:security_events", "xcsh-docs:resources:global_log_receiver:properties:splunk_receiver", "xcsh-docs:resources:global_log_receiver:properties:sumo_logic_receiver", "xcsh-docs:resources:global_log_receiver:properties:timeouts"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:reference", "parent_id": "xcsh-docs:resources:global_log_receiver:fundamentals", "path": "documentation/resources/global_log_receiver/properties/index.md", "product": "distributed-cloud", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020", "registry_path": "docs/guides/resources--global_log_receiver--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects.", "document_id": "xcsh-docs:resources:global_log_receiver:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["audit logs"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:audit_logs", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["audit_logs"], "syntax": "attribute", "type": "object"}, {"aliases": ["aws cloud watch receiver"], "anchor": "section", "description": "AWS Cloudwatch Logs Configuration for Global Log Receiver.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:aws_cloud_watch_receiver", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-aws_cloud_watch_receiver--aws_region", "enforcement": "provider-schema", "group": "aws_cloud_watch_receiver:RequiredObjectAttributes:aws_region,group_name,stream_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:aws_cloud_watch_receiver", "type": "requires"}, {"anchor": "schema-aws_cloud_watch_receiver--group_name", "enforcement": "provider-schema", "group": "aws_cloud_watch_receiver:RequiredObjectAttributes:aws_region,group_name,stream_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:aws_cloud_watch_receiver", "type": "requires"}, {"anchor": "schema-aws_cloud_watch_receiver--stream_name", "enforcement": "provider-schema", "group": "aws_cloud_watch_receiver:RequiredObjectAttributes:aws_region,group_name,stream_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:aws_cloud_watch_receiver", "type": "requires"}], "schema_path": ["aws_cloud_watch_receiver"], "syntax": "block", "type": "object"}, {"aliases": ["azure event hubs receiver"], "anchor": "section", "description": "Azure Event Hubs Configuration for Global Log Receiver.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:azure_event_hubs_receiver", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-azure_event_hubs_receiver--instance", "enforcement": "provider-schema", "group": "azure_event_hubs_receiver:RequiredObjectAttributes:instance,namespace", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:azure_event_hubs_receiver", "type": "requires"}, {"anchor": "schema-azure_event_hubs_receiver--namespace", "enforcement": "provider-schema", "group": "azure_event_hubs_receiver:RequiredObjectAttributes:instance,namespace", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:azure_event_hubs_receiver", "type": "requires"}], "schema_path": ["azure_event_hubs_receiver"], "syntax": "block", "type": "object"}, {"aliases": ["azure receiver"], "anchor": "section", "description": "Azure Blob Configuration for Global Log Receiver.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-azure_receiver--container_name", "enforcement": "provider-schema", "group": "azure_receiver:RequiredObjectAttributes:container_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver", "type": "requires"}], "schema_path": ["azure_receiver"], "syntax": "block", "type": "object"}, {"aliases": ["datadog receiver"], "anchor": "section", "description": "Configuration for Datadog endpoint.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:datadog_receiver", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-datadog_receiver--endpoint", "enforcement": "provider-schema", "group": "datadog_receiver:ConflictingObjectAttributes:endpoint,site", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:datadog_receiver", "type": "conflicts"}, {"anchor": "schema-datadog_receiver--site", "enforcement": "provider-schema", "group": "datadog_receiver:ConflictingObjectAttributes:endpoint,site", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:datadog_receiver", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "datadog_receiver:ConflictingObjectAttributes:no_tls,use_tls", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:datadog_receiver:no_tls", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "datadog_receiver:ConflictingObjectAttributes:no_tls,use_tls", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:datadog_receiver:use_tls", "type": "conflicts"}], "schema_path": ["datadog_receiver"], "syntax": "block", "type": "object"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Human readable description for the object.", "document_id": "xcsh-docs:resources:global_log_receiver:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["disable"], "anchor": "schema-disable", "description": "A value of true will administratively disable the object.", "document_id": "xcsh-docs:resources:global_log_receiver:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["disable"], "syntax": "attribute", "type": "bool"}, {"aliases": ["dns logs"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:dns_logs", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dns_logs"], "syntax": "attribute", "type": "object"}, {"aliases": ["gcp bucket receiver"], "anchor": "section", "description": "GCP Bucket Configuration for Global Log Receiver.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:gcp_bucket_receiver", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-gcp_bucket_receiver--bucket", "enforcement": "provider-schema", "group": "gcp_bucket_receiver:RequiredObjectAttributes:bucket", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:gcp_bucket_receiver", "type": "requires"}], "schema_path": ["gcp_bucket_receiver"], "syntax": "block", "type": "object"}, {"aliases": ["http receiver"], "anchor": "section", "description": "Configuration for HTTP endpoint.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "http_receiver:ConflictingObjectAttributes:auth_basic,auth_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver:auth_basic", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "http_receiver:ConflictingObjectAttributes:auth_basic,auth_token", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver:auth_basic", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "http_receiver:ConflictingObjectAttributes:auth_basic,auth_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver:auth_none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "http_receiver:ConflictingObjectAttributes:auth_none,auth_token", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver:auth_none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "http_receiver:ConflictingObjectAttributes:auth_basic,auth_token", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver:auth_token", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "http_receiver:ConflictingObjectAttributes:auth_none,auth_token", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver:auth_token", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "http_receiver:ConflictingObjectAttributes:no_tls,use_tls", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver:no_tls", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "http_receiver:ConflictingObjectAttributes:no_tls,use_tls", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver:use_tls", "type": "conflicts"}, {"anchor": "schema-http_receiver--uri", "enforcement": "provider-schema", "group": "http_receiver:RequiredObjectAttributes:uri", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver", "type": "requires"}], "schema_path": ["http_receiver"], "syntax": "block", "type": "object"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier for the resource.", "document_id": "xcsh-docs:resources:global_log_receiver:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["kafka receiver"], "anchor": "section", "description": "Kafka Configuration for Global Log Receiver.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "kafka_receiver:ConflictingObjectAttributes:no_tls,use_tls", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:no_tls", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "kafka_receiver:ConflictingObjectAttributes:no_tls,use_tls", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:use_tls", "type": "conflicts"}, {"anchor": "schema-kafka_receiver--bootstrap_servers", "enforcement": "provider-schema", "group": "kafka_receiver:RequiredObjectAttributes:bootstrap_servers,kafka_topic", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver", "type": "requires"}, {"anchor": "schema-kafka_receiver--kafka_topic", "enforcement": "provider-schema", "group": "kafka_receiver:RequiredObjectAttributes:bootstrap_servers,kafka_topic", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver", "type": "requires"}], "schema_path": ["kafka_receiver"], "syntax": "block", "type": "object"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects as chosen by the user. Values specified here will be used by selector expression.", "document_id": "xcsh-docs:resources:global_log_receiver:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["name"], "anchor": "schema-name", "description": "This is the name of configuration object. It has to be unique within the namespace. It can only be specified during create API and cannot be changed during replace API. The value of name has to follow DNS-1035 format.", "document_id": "xcsh-docs:resources:global_log_receiver:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "This defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be \"\"", "document_id": "xcsh-docs:resources:global_log_receiver:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["new relic receiver"], "anchor": "section", "description": "Configuration for NewRelic endpoint.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:new_relic_receiver", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "new_relic_receiver:ConflictingObjectAttributes:eu,us", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:new_relic_receiver:eu", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "new_relic_receiver:ConflictingObjectAttributes:eu,us", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:new_relic_receiver:us", "type": "conflicts"}], "schema_path": ["new_relic_receiver"], "syntax": "block", "type": "object"}, {"aliases": ["ns all"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:ns_all", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ns_all"], "syntax": "attribute", "type": "object"}, {"aliases": ["ns current"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:ns_current", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ns_current"], "syntax": "attribute", "type": "object"}, {"aliases": ["ns list"], "anchor": "section", "description": "Namespace List.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:ns_list", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-ns_list--namespaces", "enforcement": "provider-schema", "group": "ns_list:RequiredObjectAttributes:namespaces", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:ns_list", "type": "requires"}], "schema_path": ["ns_list"], "syntax": "block", "type": "object"}, {"aliases": ["qradar receiver"], "anchor": "section", "description": "Configuration for IBM QRadar endpoint.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "qradar_receiver:ConflictingObjectAttributes:no_tls,use_tls", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver:no_tls", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "qradar_receiver:ConflictingObjectAttributes:no_tls,use_tls", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver:use_tls", "type": "conflicts"}, {"anchor": "schema-qradar_receiver--uri", "enforcement": "provider-schema", "group": "qradar_receiver:RequiredObjectAttributes:uri", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver", "type": "requires"}], "schema_path": ["qradar_receiver"], "syntax": "block", "type": "object"}, {"aliases": ["request logs"], "anchor": "section", "description": "Configuration for request logs with sampling choice. Allows selection between sampled (default) or unsampled (full) request logs.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:request_logs", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "request_logs:ConflictingObjectAttributes:sampled,unsampled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:request_logs:sampled", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "request_logs:ConflictingObjectAttributes:sampled,unsampled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:request_logs:unsampled", "type": "conflicts"}], "schema_path": ["request_logs"], "syntax": "block", "type": "object"}, {"aliases": ["s3 receiver"], "anchor": "section", "description": "S3 Configuration for Global Log Receiver.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:s3_receiver", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-s3_receiver--aws_region", "enforcement": "provider-schema", "group": "s3_receiver:RequiredObjectAttributes:aws_region,bucket", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:s3_receiver", "type": "requires"}, {"anchor": "schema-s3_receiver--bucket", "enforcement": "provider-schema", "group": "s3_receiver:RequiredObjectAttributes:aws_region,bucket", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:s3_receiver", "type": "requires"}], "schema_path": ["s3_receiver"], "syntax": "block", "type": "object"}, {"aliases": ["security events"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:security_events", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["security_events"], "syntax": "attribute", "type": "object"}, {"aliases": ["splunk receiver"], "anchor": "section", "description": "Configuration for Splunk HEC Logs endpoint.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:splunk_receiver", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "splunk_receiver:ConflictingObjectAttributes:no_tls,use_tls", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:splunk_receiver:no_tls", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "splunk_receiver:ConflictingObjectAttributes:no_tls,use_tls", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:splunk_receiver:use_tls", "type": "conflicts"}, {"anchor": "schema-splunk_receiver--endpoint", "enforcement": "provider-schema", "group": "splunk_receiver:RequiredObjectAttributes:endpoint", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:splunk_receiver", "type": "requires"}], "schema_path": ["splunk_receiver"], "syntax": "block", "type": "object"}, {"aliases": ["sumo logic receiver"], "anchor": "section", "description": "Configuration for SumoLogic endpoint.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:sumo_logic_receiver", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["sumo_logic_receiver"], "syntax": "block", "type": "object"}, {"aliases": ["duration", "operation timeout", "timeouts"], "anchor": "section", "description": "timeouts", "document_id": "xcsh-docs:resources:global_log_receiver:properties:timeouts", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["timeouts"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_global_log_receiver.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/)
- Property reference

## Direct properties

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Optional.

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata.

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

- [audit_logs](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/audit_logs/): complete subsection reference.

- [aws_cloud_watch_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/aws_cloud_watch_receiver/): complete subsection reference.

- [azure_event_hubs_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/azure_event_hubs_receiver/): complete subsection reference.

- [azure_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/azure_receiver/): complete subsection reference.

- [datadog_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/datadog_receiver/): complete subsection reference.

<a id="schema-description"></a>

### description property

Type: `"string"`. Optional.

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

<a id="schema-disable"></a>

### disable property

Type: `"bool"`. Optional.

A value of true administratively disables the object.

Upstream description:

A value of true will administratively disable the object.

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

- [dns_logs](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/dns_logs/): complete subsection reference.

- [gcp_bucket_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/gcp_bucket_receiver/): complete subsection reference.

- [http_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/): complete subsection reference.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

- [kafka_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/): complete subsection reference.

<a id="schema-labels"></a>

### labels property

Type: `["map", "string"]`. Optional.

Labels is a user defined key value map that can be attached to resources for organization and
filtering.

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

Name of the Global Log Receiver. Must be unique within the namespace.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NameValidator(),
}
```

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

Namespace where the Global Log Receiver is created.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NamespaceValidator(),
}
```

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

- [new_relic_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/new_relic_receiver/): complete subsection reference.

- [ns_all](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/ns_all/): complete subsection reference.

- [ns_current](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/ns_current/): complete subsection reference.

- [ns_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/ns_list/): complete subsection reference.

- [qradar_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/qradar_receiver/): complete subsection reference.

- [request_logs](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/request_logs/): complete subsection reference.

- [s3_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/s3_receiver/): complete subsection reference.

- [security_events](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/security_events/): complete subsection reference.

- [splunk_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/splunk_receiver/): complete subsection reference.

- [sumo_logic_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/sumo_logic_receiver/): complete subsection reference.

- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/timeouts/): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/#schema-annotations) |
| `audit_logs` | [audit_logs](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/audit_logs/#section) |
| `aws_cloud_watch_receiver` | [aws_cloud_watch_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/aws_cloud_watch_receiver/#section) |
| `aws_cloud_watch_receiver.aws_cred` | [aws_cloud_watch_receiver.aws_cred](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/aws_cloud_watch_receiver/aws_cred/#section) |
| `aws_cloud_watch_receiver.aws_cred.name` | [aws_cloud_watch_receiver.aws_cred.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/aws_cloud_watch_receiver/aws_cred/#schema-aws_cloud_watch_receiver--aws_cred--name) |
| `aws_cloud_watch_receiver.aws_cred.namespace` | [aws_cloud_watch_receiver.aws_cred.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/aws_cloud_watch_receiver/aws_cred/#schema-aws_cloud_watch_receiver--aws_cred--namespace) |
| `aws_cloud_watch_receiver.aws_cred.tenant` | [aws_cloud_watch_receiver.aws_cred.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/aws_cloud_watch_receiver/aws_cred/#schema-aws_cloud_watch_receiver--aws_cred--tenant) |
| `aws_cloud_watch_receiver.aws_region` | [aws_cloud_watch_receiver.aws_region](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/aws_cloud_watch_receiver/#schema-aws_cloud_watch_receiver--aws_region) |
| `aws_cloud_watch_receiver.batch` | [aws_cloud_watch_receiver.batch](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/aws_cloud_watch_receiver/batch/#section) |
| `aws_cloud_watch_receiver.batch.max_bytes` | [aws_cloud_watch_receiver.batch.max_bytes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/aws_cloud_watch_receiver/batch/#schema-aws_cloud_watch_receiver--batch--max_bytes) |
| `aws_cloud_watch_receiver.batch.max_bytes_disabled` | [aws_cloud_watch_receiver.batch.max_bytes_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/aws_cloud_watch_receiver/batch/max_bytes_disabled/#section) |
| `aws_cloud_watch_receiver.batch.max_events` | [aws_cloud_watch_receiver.batch.max_events](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/aws_cloud_watch_receiver/batch/#schema-aws_cloud_watch_receiver--batch--max_events) |
| `aws_cloud_watch_receiver.batch.max_events_disabled` | [aws_cloud_watch_receiver.batch.max_events_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/aws_cloud_watch_receiver/batch/max_events_disabled/#section) |
| `aws_cloud_watch_receiver.batch.timeout_seconds` | [aws_cloud_watch_receiver.batch.timeout_seconds](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/aws_cloud_watch_receiver/batch/#schema-aws_cloud_watch_receiver--batch--timeout_seconds) |
| `aws_cloud_watch_receiver.batch.timeout_seconds_default` | [aws_cloud_watch_receiver.batch.timeout_seconds_default](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/aws_cloud_watch_receiver/batch/timeout_seconds_default/#section) |
| `aws_cloud_watch_receiver.compression` | [aws_cloud_watch_receiver.compression](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/aws_cloud_watch_receiver/compression/#section) |
| `aws_cloud_watch_receiver.compression.compression_default` | [aws_cloud_watch_receiver.compression.compression_default](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/aws_cloud_watch_receiver/compression/compression_default/#section) |
| `aws_cloud_watch_receiver.compression.compression_gzip` | [aws_cloud_watch_receiver.compression.compression_gzip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/aws_cloud_watch_receiver/compression/compression_gzip/#section) |
| `aws_cloud_watch_receiver.compression.compression_none` | [aws_cloud_watch_receiver.compression.compression_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/aws_cloud_watch_receiver/compression/compression_none/#section) |
| `aws_cloud_watch_receiver.group_name` | [aws_cloud_watch_receiver.group_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/aws_cloud_watch_receiver/#schema-aws_cloud_watch_receiver--group_name) |
| `aws_cloud_watch_receiver.stream_name` | [aws_cloud_watch_receiver.stream_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/aws_cloud_watch_receiver/#schema-aws_cloud_watch_receiver--stream_name) |
| `azure_event_hubs_receiver` | [azure_event_hubs_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/azure_event_hubs_receiver/#section) |
| `azure_event_hubs_receiver.connection_string` | [azure_event_hubs_receiver.connection_string](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/azure_event_hubs_receiver/connection_string/#section) |
| `azure_event_hubs_receiver.connection_string.blindfold_secret_info` | [azure_event_hubs_receiver.connection_string.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/azure_event_hubs_receiver/connection_string/blindfold_secret_info/#section) |
| `azure_event_hubs_receiver.connection_string.blindfold_secret_info.decryption_provider` | [azure_event_hubs_receiver.connection_string.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/azure_event_hubs_receiver/connection_string/blindfold_secret_info/#schema-azure_event_hubs_receiver--connection_string--blindfold_secret_info--decryption_provider) |
| `azure_event_hubs_receiver.connection_string.blindfold_secret_info.location` | [azure_event_hubs_receiver.connection_string.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/azure_event_hubs_receiver/connection_string/blindfold_secret_info/#schema-azure_event_hubs_receiver--connection_string--blindfold_secret_info--location) |
| `azure_event_hubs_receiver.connection_string.blindfold_secret_info.store_provider` | [azure_event_hubs_receiver.connection_string.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/azure_event_hubs_receiver/connection_string/blindfold_secret_info/#schema-azure_event_hubs_receiver--connection_string--blindfold_secret_info--store_provider) |
| `azure_event_hubs_receiver.connection_string.clear_secret_info` | [azure_event_hubs_receiver.connection_string.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/azure_event_hubs_receiver/connection_string/clear_secret_info/#section) |
| `azure_event_hubs_receiver.connection_string.clear_secret_info.provider_ref` | [azure_event_hubs_receiver.connection_string.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/azure_event_hubs_receiver/connection_string/clear_secret_info/#schema-azure_event_hubs_receiver--connection_string--clear_secret_info--provider_ref) |
| `azure_event_hubs_receiver.connection_string.clear_secret_info.url` | [azure_event_hubs_receiver.connection_string.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/azure_event_hubs_receiver/connection_string/clear_secret_info/#schema-azure_event_hubs_receiver--connection_string--clear_secret_info--url) |
| `azure_event_hubs_receiver.instance` | [azure_event_hubs_receiver.instance](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/azure_event_hubs_receiver/#schema-azure_event_hubs_receiver--instance) |
| `azure_event_hubs_receiver.namespace` | [azure_event_hubs_receiver.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/azure_event_hubs_receiver/#schema-azure_event_hubs_receiver--namespace) |
| `azure_receiver` | [azure_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/azure_receiver/#section) |
| `azure_receiver.batch` | [azure_receiver.batch](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/azure_receiver/batch/#section) |
| `azure_receiver.batch.max_bytes` | [azure_receiver.batch.max_bytes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/azure_receiver/batch/#schema-azure_receiver--batch--max_bytes) |
| `azure_receiver.batch.max_bytes_disabled` | [azure_receiver.batch.max_bytes_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/azure_receiver/batch/max_bytes_disabled/#section) |
| `azure_receiver.batch.max_events` | [azure_receiver.batch.max_events](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/azure_receiver/batch/#schema-azure_receiver--batch--max_events) |
| `azure_receiver.batch.max_events_disabled` | [azure_receiver.batch.max_events_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/azure_receiver/batch/max_events_disabled/#section) |
| `azure_receiver.batch.timeout_seconds` | [azure_receiver.batch.timeout_seconds](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/azure_receiver/batch/#schema-azure_receiver--batch--timeout_seconds) |
| `azure_receiver.batch.timeout_seconds_default` | [azure_receiver.batch.timeout_seconds_default](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/azure_receiver/batch/timeout_seconds_default/#section) |
| `azure_receiver.compression` | [azure_receiver.compression](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/azure_receiver/compression/#section) |
| `azure_receiver.compression.compression_default` | [azure_receiver.compression.compression_default](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/azure_receiver/compression/compression_default/#section) |
| `azure_receiver.compression.compression_gzip` | [azure_receiver.compression.compression_gzip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/azure_receiver/compression/compression_gzip/#section) |
| `azure_receiver.compression.compression_none` | [azure_receiver.compression.compression_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/azure_receiver/compression/compression_none/#section) |
| `azure_receiver.connection_string` | [azure_receiver.connection_string](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/azure_receiver/connection_string/#section) |
| `azure_receiver.connection_string.blindfold_secret_info` | [azure_receiver.connection_string.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/azure_receiver/connection_string/blindfold_secret_info/#section) |
| `azure_receiver.connection_string.blindfold_secret_info.decryption_provider` | [azure_receiver.connection_string.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/azure_receiver/connection_string/blindfold_secret_info/#schema-azure_receiver--connection_string--blindfold_secret_info--decryption_provider) |
| `azure_receiver.connection_string.blindfold_secret_info.location` | [azure_receiver.connection_string.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/azure_receiver/connection_string/blindfold_secret_info/#schema-azure_receiver--connection_string--blindfold_secret_info--location) |
| `azure_receiver.connection_string.blindfold_secret_info.store_provider` | [azure_receiver.connection_string.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/azure_receiver/connection_string/blindfold_secret_info/#schema-azure_receiver--connection_string--blindfold_secret_info--store_provider) |
| `azure_receiver.connection_string.clear_secret_info` | [azure_receiver.connection_string.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/azure_receiver/connection_string/clear_secret_info/#section) |
| `azure_receiver.connection_string.clear_secret_info.provider_ref` | [azure_receiver.connection_string.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/azure_receiver/connection_string/clear_secret_info/#schema-azure_receiver--connection_string--clear_secret_info--provider_ref) |
| `azure_receiver.connection_string.clear_secret_info.url` | [azure_receiver.connection_string.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/azure_receiver/connection_string/clear_secret_info/#schema-azure_receiver--connection_string--clear_secret_info--url) |
| `azure_receiver.container_name` | [azure_receiver.container_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/azure_receiver/#schema-azure_receiver--container_name) |
| `azure_receiver.filename_options` | [azure_receiver.filename_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/azure_receiver/filename_options/#section) |
| `azure_receiver.filename_options.custom_folder` | [azure_receiver.filename_options.custom_folder](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/azure_receiver/filename_options/#schema-azure_receiver--filename_options--custom_folder) |
| `azure_receiver.filename_options.log_type_folder` | [azure_receiver.filename_options.log_type_folder](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/azure_receiver/filename_options/log_type_folder/#section) |
| `azure_receiver.filename_options.no_folder` | [azure_receiver.filename_options.no_folder](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/azure_receiver/filename_options/no_folder/#section) |
| `datadog_receiver` | [datadog_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/datadog_receiver/#section) |
| `datadog_receiver.batch` | [datadog_receiver.batch](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/datadog_receiver/batch/#section) |
| `datadog_receiver.batch.max_bytes` | [datadog_receiver.batch.max_bytes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/datadog_receiver/batch/#schema-datadog_receiver--batch--max_bytes) |
| `datadog_receiver.batch.max_bytes_disabled` | [datadog_receiver.batch.max_bytes_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/datadog_receiver/batch/max_bytes_disabled/#section) |
| `datadog_receiver.batch.max_events` | [datadog_receiver.batch.max_events](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/datadog_receiver/batch/#schema-datadog_receiver--batch--max_events) |
| `datadog_receiver.batch.max_events_disabled` | [datadog_receiver.batch.max_events_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/datadog_receiver/batch/max_events_disabled/#section) |
| `datadog_receiver.batch.timeout_seconds` | [datadog_receiver.batch.timeout_seconds](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/datadog_receiver/batch/#schema-datadog_receiver--batch--timeout_seconds) |
| `datadog_receiver.batch.timeout_seconds_default` | [datadog_receiver.batch.timeout_seconds_default](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/datadog_receiver/batch/timeout_seconds_default/#section) |
| `datadog_receiver.compression` | [datadog_receiver.compression](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/datadog_receiver/compression/#section) |
| `datadog_receiver.compression.compression_default` | [datadog_receiver.compression.compression_default](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/datadog_receiver/compression/compression_default/#section) |
| `datadog_receiver.compression.compression_gzip` | [datadog_receiver.compression.compression_gzip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/datadog_receiver/compression/compression_gzip/#section) |
| `datadog_receiver.compression.compression_none` | [datadog_receiver.compression.compression_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/datadog_receiver/compression/compression_none/#section) |
| `datadog_receiver.datadog_api_key` | [datadog_receiver.datadog_api_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/datadog_receiver/datadog_api_key/#section) |
| `datadog_receiver.datadog_api_key.blindfold_secret_info` | [datadog_receiver.datadog_api_key.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/datadog_receiver/datadog_api_key/blindfold_secret_info/#section) |
| `datadog_receiver.datadog_api_key.blindfold_secret_info.decryption_provider` | [datadog_receiver.datadog_api_key.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/datadog_receiver/datadog_api_key/blindfold_secret_info/#schema-datadog_receiver--datadog_api_key--blindfold_secret_info--decryption_provider) |
| `datadog_receiver.datadog_api_key.blindfold_secret_info.location` | [datadog_receiver.datadog_api_key.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/datadog_receiver/datadog_api_key/blindfold_secret_info/#schema-datadog_receiver--datadog_api_key--blindfold_secret_info--location) |
| `datadog_receiver.datadog_api_key.blindfold_secret_info.store_provider` | [datadog_receiver.datadog_api_key.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/datadog_receiver/datadog_api_key/blindfold_secret_info/#schema-datadog_receiver--datadog_api_key--blindfold_secret_info--store_provider) |
| `datadog_receiver.datadog_api_key.clear_secret_info` | [datadog_receiver.datadog_api_key.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/datadog_receiver/datadog_api_key/clear_secret_info/#section) |
| `datadog_receiver.datadog_api_key.clear_secret_info.provider_ref` | [datadog_receiver.datadog_api_key.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/datadog_receiver/datadog_api_key/clear_secret_info/#schema-datadog_receiver--datadog_api_key--clear_secret_info--provider_ref) |
| `datadog_receiver.datadog_api_key.clear_secret_info.url` | [datadog_receiver.datadog_api_key.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/datadog_receiver/datadog_api_key/clear_secret_info/#schema-datadog_receiver--datadog_api_key--clear_secret_info--url) |
| `datadog_receiver.endpoint` | [datadog_receiver.endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/datadog_receiver/#schema-datadog_receiver--endpoint) |
| `datadog_receiver.no_tls` | [datadog_receiver.no_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/datadog_receiver/no_tls/#section) |
| `datadog_receiver.site` | [datadog_receiver.site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/datadog_receiver/#schema-datadog_receiver--site) |
| `datadog_receiver.use_tls` | [datadog_receiver.use_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/datadog_receiver/use_tls/#section) |
| `datadog_receiver.use_tls.disable_verify_certificate` | [datadog_receiver.use_tls.disable_verify_certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/datadog_receiver/use_tls/disable_verify_certificate/#section) |
| `datadog_receiver.use_tls.disable_verify_hostname` | [datadog_receiver.use_tls.disable_verify_hostname](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/datadog_receiver/use_tls/disable_verify_hostname/#section) |
| `datadog_receiver.use_tls.enable_verify_certificate` | [datadog_receiver.use_tls.enable_verify_certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/datadog_receiver/use_tls/enable_verify_certificate/#section) |
| `datadog_receiver.use_tls.enable_verify_hostname` | [datadog_receiver.use_tls.enable_verify_hostname](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/datadog_receiver/use_tls/enable_verify_hostname/#section) |
| `datadog_receiver.use_tls.mtls_disabled` | [datadog_receiver.use_tls.mtls_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/datadog_receiver/use_tls/mtls_disabled/#section) |
| `datadog_receiver.use_tls.mtls_enable` | [datadog_receiver.use_tls.mtls_enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/datadog_receiver/use_tls/mtls_enable/#section) |
| `datadog_receiver.use_tls.mtls_enable.certificate` | [datadog_receiver.use_tls.mtls_enable.certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/datadog_receiver/use_tls/mtls_enable/#schema-datadog_receiver--use_tls--mtls_enable--certificate) |
| `datadog_receiver.use_tls.mtls_enable.key_url` | [datadog_receiver.use_tls.mtls_enable.key_url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/datadog_receiver/use_tls/mtls_enable/key_url/#section) |
| `datadog_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info` | [datadog_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/datadog_receiver/use_tls/mtls_enable/key_url/blindfold_secret_info/#section) |
| `datadog_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.decryption_provider` | [datadog_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/datadog_receiver/use_tls/mtls_enable/key_url/blindfold_secret_info/#schema-datadog_receiver--use_tls--mtls_enable--key_url--blindfold_secret_info--decryption_provider) |
| `datadog_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.location` | [datadog_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/datadog_receiver/use_tls/mtls_enable/key_url/blindfold_secret_info/#schema-datadog_receiver--use_tls--mtls_enable--key_url--blindfold_secret_info--location) |
| `datadog_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.store_provider` | [datadog_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/datadog_receiver/use_tls/mtls_enable/key_url/blindfold_secret_info/#schema-datadog_receiver--use_tls--mtls_enable--key_url--blindfold_secret_info--store_provider) |
| `datadog_receiver.use_tls.mtls_enable.key_url.clear_secret_info` | [datadog_receiver.use_tls.mtls_enable.key_url.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/datadog_receiver/use_tls/mtls_enable/key_url/clear_secret_info/#section) |
| `datadog_receiver.use_tls.mtls_enable.key_url.clear_secret_info.provider_ref` | [datadog_receiver.use_tls.mtls_enable.key_url.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/datadog_receiver/use_tls/mtls_enable/key_url/clear_secret_info/#schema-datadog_receiver--use_tls--mtls_enable--key_url--clear_secret_info--provider_ref) |
| `datadog_receiver.use_tls.mtls_enable.key_url.clear_secret_info.url` | [datadog_receiver.use_tls.mtls_enable.key_url.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/datadog_receiver/use_tls/mtls_enable/key_url/clear_secret_info/#schema-datadog_receiver--use_tls--mtls_enable--key_url--clear_secret_info--url) |
| `datadog_receiver.use_tls.no_ca` | [datadog_receiver.use_tls.no_ca](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/datadog_receiver/use_tls/no_ca/#section) |
| `datadog_receiver.use_tls.trusted_ca_url` | [datadog_receiver.use_tls.trusted_ca_url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/datadog_receiver/use_tls/#schema-datadog_receiver--use_tls--trusted_ca_url) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/#schema-description) |
| `disable` | [disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/#schema-disable) |
| `dns_logs` | [dns_logs](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/dns_logs/#section) |
| `gcp_bucket_receiver` | [gcp_bucket_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/gcp_bucket_receiver/#section) |
| `gcp_bucket_receiver.batch` | [gcp_bucket_receiver.batch](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/gcp_bucket_receiver/batch/#section) |
| `gcp_bucket_receiver.batch.max_bytes` | [gcp_bucket_receiver.batch.max_bytes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/gcp_bucket_receiver/batch/#schema-gcp_bucket_receiver--batch--max_bytes) |
| `gcp_bucket_receiver.batch.max_bytes_disabled` | [gcp_bucket_receiver.batch.max_bytes_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/gcp_bucket_receiver/batch/max_bytes_disabled/#section) |
| `gcp_bucket_receiver.batch.max_events` | [gcp_bucket_receiver.batch.max_events](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/gcp_bucket_receiver/batch/#schema-gcp_bucket_receiver--batch--max_events) |
| `gcp_bucket_receiver.batch.max_events_disabled` | [gcp_bucket_receiver.batch.max_events_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/gcp_bucket_receiver/batch/max_events_disabled/#section) |
| `gcp_bucket_receiver.batch.timeout_seconds` | [gcp_bucket_receiver.batch.timeout_seconds](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/gcp_bucket_receiver/batch/#schema-gcp_bucket_receiver--batch--timeout_seconds) |
| `gcp_bucket_receiver.batch.timeout_seconds_default` | [gcp_bucket_receiver.batch.timeout_seconds_default](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/gcp_bucket_receiver/batch/timeout_seconds_default/#section) |
| `gcp_bucket_receiver.bucket` | [gcp_bucket_receiver.bucket](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/gcp_bucket_receiver/#schema-gcp_bucket_receiver--bucket) |
| `gcp_bucket_receiver.compression` | [gcp_bucket_receiver.compression](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/gcp_bucket_receiver/compression/#section) |
| `gcp_bucket_receiver.compression.compression_default` | [gcp_bucket_receiver.compression.compression_default](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/gcp_bucket_receiver/compression/compression_default/#section) |
| `gcp_bucket_receiver.compression.compression_gzip` | [gcp_bucket_receiver.compression.compression_gzip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/gcp_bucket_receiver/compression/compression_gzip/#section) |
| `gcp_bucket_receiver.compression.compression_none` | [gcp_bucket_receiver.compression.compression_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/gcp_bucket_receiver/compression/compression_none/#section) |
| `gcp_bucket_receiver.filename_options` | [gcp_bucket_receiver.filename_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/gcp_bucket_receiver/filename_options/#section) |
| `gcp_bucket_receiver.filename_options.custom_folder` | [gcp_bucket_receiver.filename_options.custom_folder](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/gcp_bucket_receiver/filename_options/#schema-gcp_bucket_receiver--filename_options--custom_folder) |
| `gcp_bucket_receiver.filename_options.log_type_folder` | [gcp_bucket_receiver.filename_options.log_type_folder](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/gcp_bucket_receiver/filename_options/log_type_folder/#section) |
| `gcp_bucket_receiver.filename_options.no_folder` | [gcp_bucket_receiver.filename_options.no_folder](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/gcp_bucket_receiver/filename_options/no_folder/#section) |
| `gcp_bucket_receiver.gcp_cred` | [gcp_bucket_receiver.gcp_cred](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/gcp_bucket_receiver/gcp_cred/#section) |
| `gcp_bucket_receiver.gcp_cred.name` | [gcp_bucket_receiver.gcp_cred.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/gcp_bucket_receiver/gcp_cred/#schema-gcp_bucket_receiver--gcp_cred--name) |
| `gcp_bucket_receiver.gcp_cred.namespace` | [gcp_bucket_receiver.gcp_cred.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/gcp_bucket_receiver/gcp_cred/#schema-gcp_bucket_receiver--gcp_cred--namespace) |
| `gcp_bucket_receiver.gcp_cred.tenant` | [gcp_bucket_receiver.gcp_cred.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/gcp_bucket_receiver/gcp_cred/#schema-gcp_bucket_receiver--gcp_cred--tenant) |
| `http_receiver` | [http_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/#section) |
| `http_receiver.auth_basic` | [http_receiver.auth_basic](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/auth_basic/#section) |
| `http_receiver.auth_basic.password` | [http_receiver.auth_basic.password](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/auth_basic/password/#section) |
| `http_receiver.auth_basic.password.blindfold_secret_info` | [http_receiver.auth_basic.password.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/auth_basic/password/blindfold_secret_info/#section) |
| `http_receiver.auth_basic.password.blindfold_secret_info.decryption_provider` | [http_receiver.auth_basic.password.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/auth_basic/password/blindfold_secret_info/#schema-http_receiver--auth_basic--password--blindfold_secret_info--decryption_provider) |
| `http_receiver.auth_basic.password.blindfold_secret_info.location` | [http_receiver.auth_basic.password.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/auth_basic/password/blindfold_secret_info/#schema-http_receiver--auth_basic--password--blindfold_secret_info--location) |
| `http_receiver.auth_basic.password.blindfold_secret_info.store_provider` | [http_receiver.auth_basic.password.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/auth_basic/password/blindfold_secret_info/#schema-http_receiver--auth_basic--password--blindfold_secret_info--store_provider) |
| `http_receiver.auth_basic.password.clear_secret_info` | [http_receiver.auth_basic.password.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/auth_basic/password/clear_secret_info/#section) |
| `http_receiver.auth_basic.password.clear_secret_info.provider_ref` | [http_receiver.auth_basic.password.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/auth_basic/password/clear_secret_info/#schema-http_receiver--auth_basic--password--clear_secret_info--provider_ref) |
| `http_receiver.auth_basic.password.clear_secret_info.url` | [http_receiver.auth_basic.password.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/auth_basic/password/clear_secret_info/#schema-http_receiver--auth_basic--password--clear_secret_info--url) |
| `http_receiver.auth_basic.user_name` | [http_receiver.auth_basic.user_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/auth_basic/#schema-http_receiver--auth_basic--user_name) |
| `http_receiver.auth_none` | [http_receiver.auth_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/auth_none/#section) |
| `http_receiver.auth_token` | [http_receiver.auth_token](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/auth_token/#section) |
| `http_receiver.auth_token.token` | [http_receiver.auth_token.token](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/auth_token/token/#section) |
| `http_receiver.auth_token.token.blindfold_secret_info` | [http_receiver.auth_token.token.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/auth_token/token/blindfold_secret_info/#section) |
| `http_receiver.auth_token.token.blindfold_secret_info.decryption_provider` | [http_receiver.auth_token.token.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/auth_token/token/blindfold_secret_info/#schema-http_receiver--auth_token--token--blindfold_secret_info--decryption_provider) |
| `http_receiver.auth_token.token.blindfold_secret_info.location` | [http_receiver.auth_token.token.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/auth_token/token/blindfold_secret_info/#schema-http_receiver--auth_token--token--blindfold_secret_info--location) |
| `http_receiver.auth_token.token.blindfold_secret_info.store_provider` | [http_receiver.auth_token.token.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/auth_token/token/blindfold_secret_info/#schema-http_receiver--auth_token--token--blindfold_secret_info--store_provider) |
| `http_receiver.auth_token.token.clear_secret_info` | [http_receiver.auth_token.token.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/auth_token/token/clear_secret_info/#section) |
| `http_receiver.auth_token.token.clear_secret_info.provider_ref` | [http_receiver.auth_token.token.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/auth_token/token/clear_secret_info/#schema-http_receiver--auth_token--token--clear_secret_info--provider_ref) |
| `http_receiver.auth_token.token.clear_secret_info.url` | [http_receiver.auth_token.token.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/auth_token/token/clear_secret_info/#schema-http_receiver--auth_token--token--clear_secret_info--url) |
| `http_receiver.batch` | [http_receiver.batch](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/batch/#section) |
| `http_receiver.batch.max_bytes` | [http_receiver.batch.max_bytes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/batch/#schema-http_receiver--batch--max_bytes) |
| `http_receiver.batch.max_bytes_disabled` | [http_receiver.batch.max_bytes_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/batch/max_bytes_disabled/#section) |
| `http_receiver.batch.max_events` | [http_receiver.batch.max_events](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/batch/#schema-http_receiver--batch--max_events) |
| `http_receiver.batch.max_events_disabled` | [http_receiver.batch.max_events_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/batch/max_events_disabled/#section) |
| `http_receiver.batch.timeout_seconds` | [http_receiver.batch.timeout_seconds](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/batch/#schema-http_receiver--batch--timeout_seconds) |
| `http_receiver.batch.timeout_seconds_default` | [http_receiver.batch.timeout_seconds_default](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/batch/timeout_seconds_default/#section) |
| `http_receiver.compression` | [http_receiver.compression](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/compression/#section) |
| `http_receiver.compression.compression_default` | [http_receiver.compression.compression_default](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/compression/compression_default/#section) |
| `http_receiver.compression.compression_gzip` | [http_receiver.compression.compression_gzip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/compression/compression_gzip/#section) |
| `http_receiver.compression.compression_none` | [http_receiver.compression.compression_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/compression/compression_none/#section) |
| `http_receiver.no_tls` | [http_receiver.no_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/no_tls/#section) |
| `http_receiver.uri` | [http_receiver.uri](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/#schema-http_receiver--uri) |
| `http_receiver.use_tls` | [http_receiver.use_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/use_tls/#section) |
| `http_receiver.use_tls.disable_verify_certificate` | [http_receiver.use_tls.disable_verify_certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/use_tls/disable_verify_certificate/#section) |
| `http_receiver.use_tls.disable_verify_hostname` | [http_receiver.use_tls.disable_verify_hostname](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/use_tls/disable_verify_hostname/#section) |
| `http_receiver.use_tls.enable_verify_certificate` | [http_receiver.use_tls.enable_verify_certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/use_tls/enable_verify_certificate/#section) |
| `http_receiver.use_tls.enable_verify_hostname` | [http_receiver.use_tls.enable_verify_hostname](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/use_tls/enable_verify_hostname/#section) |
| `http_receiver.use_tls.mtls_disabled` | [http_receiver.use_tls.mtls_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/use_tls/mtls_disabled/#section) |
| `http_receiver.use_tls.mtls_enable` | [http_receiver.use_tls.mtls_enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/use_tls/mtls_enable/#section) |
| `http_receiver.use_tls.mtls_enable.certificate` | [http_receiver.use_tls.mtls_enable.certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/use_tls/mtls_enable/#schema-http_receiver--use_tls--mtls_enable--certificate) |
| `http_receiver.use_tls.mtls_enable.key_url` | [http_receiver.use_tls.mtls_enable.key_url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/use_tls/mtls_enable/key_url/#section) |
| `http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info` | [http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/use_tls/mtls_enable/key_url/blindfold_secret_info/#section) |
| `http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.decryption_provider` | [http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/use_tls/mtls_enable/key_url/blindfold_secret_info/#schema-http_receiver--use_tls--mtls_enable--key_url--blindfold_secret_info--decryption_provider) |
| `http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.location` | [http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/use_tls/mtls_enable/key_url/blindfold_secret_info/#schema-http_receiver--use_tls--mtls_enable--key_url--blindfold_secret_info--location) |
| `http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.store_provider` | [http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/use_tls/mtls_enable/key_url/blindfold_secret_info/#schema-http_receiver--use_tls--mtls_enable--key_url--blindfold_secret_info--store_provider) |
| `http_receiver.use_tls.mtls_enable.key_url.clear_secret_info` | [http_receiver.use_tls.mtls_enable.key_url.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/use_tls/mtls_enable/key_url/clear_secret_info/#section) |
| `http_receiver.use_tls.mtls_enable.key_url.clear_secret_info.provider_ref` | [http_receiver.use_tls.mtls_enable.key_url.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/use_tls/mtls_enable/key_url/clear_secret_info/#schema-http_receiver--use_tls--mtls_enable--key_url--clear_secret_info--provider_ref) |
| `http_receiver.use_tls.mtls_enable.key_url.clear_secret_info.url` | [http_receiver.use_tls.mtls_enable.key_url.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/use_tls/mtls_enable/key_url/clear_secret_info/#schema-http_receiver--use_tls--mtls_enable--key_url--clear_secret_info--url) |
| `http_receiver.use_tls.no_ca` | [http_receiver.use_tls.no_ca](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/use_tls/no_ca/#section) |
| `http_receiver.use_tls.trusted_ca_url` | [http_receiver.use_tls.trusted_ca_url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/use_tls/#schema-http_receiver--use_tls--trusted_ca_url) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/#schema-id) |
| `kafka_receiver` | [kafka_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/#section) |
| `kafka_receiver.batch` | [kafka_receiver.batch](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/batch/#section) |
| `kafka_receiver.batch.max_bytes` | [kafka_receiver.batch.max_bytes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/batch/#schema-kafka_receiver--batch--max_bytes) |
| `kafka_receiver.batch.max_bytes_disabled` | [kafka_receiver.batch.max_bytes_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/batch/max_bytes_disabled/#section) |
| `kafka_receiver.batch.max_events` | [kafka_receiver.batch.max_events](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/batch/#schema-kafka_receiver--batch--max_events) |
| `kafka_receiver.batch.max_events_disabled` | [kafka_receiver.batch.max_events_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/batch/max_events_disabled/#section) |
| `kafka_receiver.batch.timeout_seconds` | [kafka_receiver.batch.timeout_seconds](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/batch/#schema-kafka_receiver--batch--timeout_seconds) |
| `kafka_receiver.batch.timeout_seconds_default` | [kafka_receiver.batch.timeout_seconds_default](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/batch/timeout_seconds_default/#section) |
| `kafka_receiver.bootstrap_servers` | [kafka_receiver.bootstrap_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/#schema-kafka_receiver--bootstrap_servers) |
| `kafka_receiver.compression` | [kafka_receiver.compression](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/compression/#section) |
| `kafka_receiver.compression.compression_default` | [kafka_receiver.compression.compression_default](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/compression/compression_default/#section) |
| `kafka_receiver.compression.compression_gzip` | [kafka_receiver.compression.compression_gzip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/compression/compression_gzip/#section) |
| `kafka_receiver.compression.compression_none` | [kafka_receiver.compression.compression_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/compression/compression_none/#section) |
| `kafka_receiver.kafka_topic` | [kafka_receiver.kafka_topic](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/#schema-kafka_receiver--kafka_topic) |
| `kafka_receiver.no_tls` | [kafka_receiver.no_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/no_tls/#section) |
| `kafka_receiver.use_tls` | [kafka_receiver.use_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/use_tls/#section) |
| `kafka_receiver.use_tls.disable_verify_certificate` | [kafka_receiver.use_tls.disable_verify_certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/use_tls/disable_verify_certificate/#section) |
| `kafka_receiver.use_tls.disable_verify_hostname` | [kafka_receiver.use_tls.disable_verify_hostname](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/use_tls/disable_verify_hostname/#section) |
| `kafka_receiver.use_tls.enable_verify_certificate` | [kafka_receiver.use_tls.enable_verify_certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/use_tls/enable_verify_certificate/#section) |
| `kafka_receiver.use_tls.enable_verify_hostname` | [kafka_receiver.use_tls.enable_verify_hostname](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/use_tls/enable_verify_hostname/#section) |
| `kafka_receiver.use_tls.mtls_disabled` | [kafka_receiver.use_tls.mtls_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/use_tls/mtls_disabled/#section) |
| `kafka_receiver.use_tls.mtls_enable` | [kafka_receiver.use_tls.mtls_enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/use_tls/mtls_enable/#section) |
| `kafka_receiver.use_tls.mtls_enable.certificate` | [kafka_receiver.use_tls.mtls_enable.certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/use_tls/mtls_enable/#schema-kafka_receiver--use_tls--mtls_enable--certificate) |
| `kafka_receiver.use_tls.mtls_enable.key_url` | [kafka_receiver.use_tls.mtls_enable.key_url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/use_tls/mtls_enable/key_url/#section) |
| `kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info` | [kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/use_tls/mtls_enable/key_url/blindfold_secret_info/#section) |
| `kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.decryption_provider` | [kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/use_tls/mtls_enable/key_url/blindfold_secret_info/#schema-kafka_receiver--use_tls--mtls_enable--key_url--blindfold_secret_info--decryption_provider) |
| `kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.location` | [kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/use_tls/mtls_enable/key_url/blindfold_secret_info/#schema-kafka_receiver--use_tls--mtls_enable--key_url--blindfold_secret_info--location) |
| `kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.store_provider` | [kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/use_tls/mtls_enable/key_url/blindfold_secret_info/#schema-kafka_receiver--use_tls--mtls_enable--key_url--blindfold_secret_info--store_provider) |
| `kafka_receiver.use_tls.mtls_enable.key_url.clear_secret_info` | [kafka_receiver.use_tls.mtls_enable.key_url.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/use_tls/mtls_enable/key_url/clear_secret_info/#section) |
| `kafka_receiver.use_tls.mtls_enable.key_url.clear_secret_info.provider_ref` | [kafka_receiver.use_tls.mtls_enable.key_url.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/use_tls/mtls_enable/key_url/clear_secret_info/#schema-kafka_receiver--use_tls--mtls_enable--key_url--clear_secret_info--provider_ref) |
| `kafka_receiver.use_tls.mtls_enable.key_url.clear_secret_info.url` | [kafka_receiver.use_tls.mtls_enable.key_url.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/use_tls/mtls_enable/key_url/clear_secret_info/#schema-kafka_receiver--use_tls--mtls_enable--key_url--clear_secret_info--url) |
| `kafka_receiver.use_tls.no_ca` | [kafka_receiver.use_tls.no_ca](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/use_tls/no_ca/#section) |
| `kafka_receiver.use_tls.trusted_ca_url` | [kafka_receiver.use_tls.trusted_ca_url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/use_tls/#schema-kafka_receiver--use_tls--trusted_ca_url) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/#schema-namespace) |
| `new_relic_receiver` | [new_relic_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/new_relic_receiver/#section) |
| `new_relic_receiver.api_key` | [new_relic_receiver.api_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/new_relic_receiver/api_key/#section) |
| `new_relic_receiver.api_key.blindfold_secret_info` | [new_relic_receiver.api_key.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/new_relic_receiver/api_key/blindfold_secret_info/#section) |
| `new_relic_receiver.api_key.blindfold_secret_info.decryption_provider` | [new_relic_receiver.api_key.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/new_relic_receiver/api_key/blindfold_secret_info/#schema-new_relic_receiver--api_key--blindfold_secret_info--decryption_provider) |
| `new_relic_receiver.api_key.blindfold_secret_info.location` | [new_relic_receiver.api_key.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/new_relic_receiver/api_key/blindfold_secret_info/#schema-new_relic_receiver--api_key--blindfold_secret_info--location) |
| `new_relic_receiver.api_key.blindfold_secret_info.store_provider` | [new_relic_receiver.api_key.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/new_relic_receiver/api_key/blindfold_secret_info/#schema-new_relic_receiver--api_key--blindfold_secret_info--store_provider) |
| `new_relic_receiver.api_key.clear_secret_info` | [new_relic_receiver.api_key.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/new_relic_receiver/api_key/clear_secret_info/#section) |
| `new_relic_receiver.api_key.clear_secret_info.provider_ref` | [new_relic_receiver.api_key.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/new_relic_receiver/api_key/clear_secret_info/#schema-new_relic_receiver--api_key--clear_secret_info--provider_ref) |
| `new_relic_receiver.api_key.clear_secret_info.url` | [new_relic_receiver.api_key.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/new_relic_receiver/api_key/clear_secret_info/#schema-new_relic_receiver--api_key--clear_secret_info--url) |
| `new_relic_receiver.eu` | [new_relic_receiver.eu](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/new_relic_receiver/eu/#section) |
| `new_relic_receiver.us` | [new_relic_receiver.us](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/new_relic_receiver/us/#section) |
| `ns_all` | [ns_all](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/ns_all/#section) |
| `ns_current` | [ns_current](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/ns_current/#section) |
| `ns_list` | [ns_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/ns_list/#section) |
| `ns_list.namespaces` | [ns_list.namespaces](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/ns_list/#schema-ns_list--namespaces) |
| `qradar_receiver` | [qradar_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/qradar_receiver/#section) |
| `qradar_receiver.batch` | [qradar_receiver.batch](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/qradar_receiver/batch/#section) |
| `qradar_receiver.batch.max_bytes` | [qradar_receiver.batch.max_bytes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/qradar_receiver/batch/#schema-qradar_receiver--batch--max_bytes) |
| `qradar_receiver.batch.max_bytes_disabled` | [qradar_receiver.batch.max_bytes_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/qradar_receiver/batch/max_bytes_disabled/#section) |
| `qradar_receiver.batch.max_events` | [qradar_receiver.batch.max_events](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/qradar_receiver/batch/#schema-qradar_receiver--batch--max_events) |
| `qradar_receiver.batch.max_events_disabled` | [qradar_receiver.batch.max_events_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/qradar_receiver/batch/max_events_disabled/#section) |
| `qradar_receiver.batch.timeout_seconds` | [qradar_receiver.batch.timeout_seconds](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/qradar_receiver/batch/#schema-qradar_receiver--batch--timeout_seconds) |
| `qradar_receiver.batch.timeout_seconds_default` | [qradar_receiver.batch.timeout_seconds_default](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/qradar_receiver/batch/timeout_seconds_default/#section) |
| `qradar_receiver.compression` | [qradar_receiver.compression](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/qradar_receiver/compression/#section) |
| `qradar_receiver.compression.compression_default` | [qradar_receiver.compression.compression_default](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/qradar_receiver/compression/compression_default/#section) |
| `qradar_receiver.compression.compression_gzip` | [qradar_receiver.compression.compression_gzip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/qradar_receiver/compression/compression_gzip/#section) |
| `qradar_receiver.compression.compression_none` | [qradar_receiver.compression.compression_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/qradar_receiver/compression/compression_none/#section) |
| `qradar_receiver.no_tls` | [qradar_receiver.no_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/qradar_receiver/no_tls/#section) |
| `qradar_receiver.uri` | [qradar_receiver.uri](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/qradar_receiver/#schema-qradar_receiver--uri) |
| `qradar_receiver.use_tls` | [qradar_receiver.use_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/qradar_receiver/use_tls/#section) |
| `qradar_receiver.use_tls.disable_verify_certificate` | [qradar_receiver.use_tls.disable_verify_certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/qradar_receiver/use_tls/disable_verify_certificate/#section) |
| `qradar_receiver.use_tls.disable_verify_hostname` | [qradar_receiver.use_tls.disable_verify_hostname](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/qradar_receiver/use_tls/disable_verify_hostname/#section) |
| `qradar_receiver.use_tls.enable_verify_certificate` | [qradar_receiver.use_tls.enable_verify_certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/qradar_receiver/use_tls/enable_verify_certificate/#section) |
| `qradar_receiver.use_tls.enable_verify_hostname` | [qradar_receiver.use_tls.enable_verify_hostname](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/qradar_receiver/use_tls/enable_verify_hostname/#section) |
| `qradar_receiver.use_tls.mtls_disabled` | [qradar_receiver.use_tls.mtls_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/qradar_receiver/use_tls/mtls_disabled/#section) |
| `qradar_receiver.use_tls.mtls_enable` | [qradar_receiver.use_tls.mtls_enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/qradar_receiver/use_tls/mtls_enable/#section) |
| `qradar_receiver.use_tls.mtls_enable.certificate` | [qradar_receiver.use_tls.mtls_enable.certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/qradar_receiver/use_tls/mtls_enable/#schema-qradar_receiver--use_tls--mtls_enable--certificate) |
| `qradar_receiver.use_tls.mtls_enable.key_url` | [qradar_receiver.use_tls.mtls_enable.key_url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/qradar_receiver/use_tls/mtls_enable/key_url/#section) |
| `qradar_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info` | [qradar_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/qradar_receiver/use_tls/mtls_enable/key_url/blindfold_secret_info/#section) |
| `qradar_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.decryption_provider` | [qradar_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/qradar_receiver/use_tls/mtls_enable/key_url/blindfold_secret_info/#schema-qradar_receiver--use_tls--mtls_enable--key_url--blindfold_secret_info--decryption_provider) |
| `qradar_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.location` | [qradar_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/qradar_receiver/use_tls/mtls_enable/key_url/blindfold_secret_info/#schema-qradar_receiver--use_tls--mtls_enable--key_url--blindfold_secret_info--location) |
| `qradar_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.store_provider` | [qradar_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/qradar_receiver/use_tls/mtls_enable/key_url/blindfold_secret_info/#schema-qradar_receiver--use_tls--mtls_enable--key_url--blindfold_secret_info--store_provider) |
| `qradar_receiver.use_tls.mtls_enable.key_url.clear_secret_info` | [qradar_receiver.use_tls.mtls_enable.key_url.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/qradar_receiver/use_tls/mtls_enable/key_url/clear_secret_info/#section) |
| `qradar_receiver.use_tls.mtls_enable.key_url.clear_secret_info.provider_ref` | [qradar_receiver.use_tls.mtls_enable.key_url.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/qradar_receiver/use_tls/mtls_enable/key_url/clear_secret_info/#schema-qradar_receiver--use_tls--mtls_enable--key_url--clear_secret_info--provider_ref) |
| `qradar_receiver.use_tls.mtls_enable.key_url.clear_secret_info.url` | [qradar_receiver.use_tls.mtls_enable.key_url.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/qradar_receiver/use_tls/mtls_enable/key_url/clear_secret_info/#schema-qradar_receiver--use_tls--mtls_enable--key_url--clear_secret_info--url) |
| `qradar_receiver.use_tls.no_ca` | [qradar_receiver.use_tls.no_ca](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/qradar_receiver/use_tls/no_ca/#section) |
| `qradar_receiver.use_tls.trusted_ca_url` | [qradar_receiver.use_tls.trusted_ca_url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/qradar_receiver/use_tls/#schema-qradar_receiver--use_tls--trusted_ca_url) |
| `request_logs` | [request_logs](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/request_logs/#section) |
| `request_logs.sampled` | [request_logs.sampled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/request_logs/sampled/#section) |
| `request_logs.unsampled` | [request_logs.unsampled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/request_logs/unsampled/#section) |
| `s3_receiver` | [s3_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/s3_receiver/#section) |
| `s3_receiver.aws_cred` | [s3_receiver.aws_cred](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/s3_receiver/aws_cred/#section) |
| `s3_receiver.aws_cred.name` | [s3_receiver.aws_cred.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/s3_receiver/aws_cred/#schema-s3_receiver--aws_cred--name) |
| `s3_receiver.aws_cred.namespace` | [s3_receiver.aws_cred.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/s3_receiver/aws_cred/#schema-s3_receiver--aws_cred--namespace) |
| `s3_receiver.aws_cred.tenant` | [s3_receiver.aws_cred.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/s3_receiver/aws_cred/#schema-s3_receiver--aws_cred--tenant) |
| `s3_receiver.aws_region` | [s3_receiver.aws_region](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/s3_receiver/#schema-s3_receiver--aws_region) |
| `s3_receiver.batch` | [s3_receiver.batch](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/s3_receiver/batch/#section) |
| `s3_receiver.batch.max_bytes` | [s3_receiver.batch.max_bytes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/s3_receiver/batch/#schema-s3_receiver--batch--max_bytes) |
| `s3_receiver.batch.max_bytes_disabled` | [s3_receiver.batch.max_bytes_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/s3_receiver/batch/max_bytes_disabled/#section) |
| `s3_receiver.batch.max_events` | [s3_receiver.batch.max_events](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/s3_receiver/batch/#schema-s3_receiver--batch--max_events) |
| `s3_receiver.batch.max_events_disabled` | [s3_receiver.batch.max_events_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/s3_receiver/batch/max_events_disabled/#section) |
| `s3_receiver.batch.timeout_seconds` | [s3_receiver.batch.timeout_seconds](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/s3_receiver/batch/#schema-s3_receiver--batch--timeout_seconds) |
| `s3_receiver.batch.timeout_seconds_default` | [s3_receiver.batch.timeout_seconds_default](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/s3_receiver/batch/timeout_seconds_default/#section) |
| `s3_receiver.bucket` | [s3_receiver.bucket](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/s3_receiver/#schema-s3_receiver--bucket) |
| `s3_receiver.compression` | [s3_receiver.compression](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/s3_receiver/compression/#section) |
| `s3_receiver.compression.compression_default` | [s3_receiver.compression.compression_default](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/s3_receiver/compression/compression_default/#section) |
| `s3_receiver.compression.compression_gzip` | [s3_receiver.compression.compression_gzip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/s3_receiver/compression/compression_gzip/#section) |
| `s3_receiver.compression.compression_none` | [s3_receiver.compression.compression_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/s3_receiver/compression/compression_none/#section) |
| `s3_receiver.filename_options` | [s3_receiver.filename_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/s3_receiver/filename_options/#section) |
| `s3_receiver.filename_options.custom_folder` | [s3_receiver.filename_options.custom_folder](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/s3_receiver/filename_options/#schema-s3_receiver--filename_options--custom_folder) |
| `s3_receiver.filename_options.log_type_folder` | [s3_receiver.filename_options.log_type_folder](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/s3_receiver/filename_options/log_type_folder/#section) |
| `s3_receiver.filename_options.no_folder` | [s3_receiver.filename_options.no_folder](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/s3_receiver/filename_options/no_folder/#section) |
| `security_events` | [security_events](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/security_events/#section) |
| `splunk_receiver` | [splunk_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/splunk_receiver/#section) |
| `splunk_receiver.batch` | [splunk_receiver.batch](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/splunk_receiver/batch/#section) |
| `splunk_receiver.batch.max_bytes` | [splunk_receiver.batch.max_bytes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/splunk_receiver/batch/#schema-splunk_receiver--batch--max_bytes) |
| `splunk_receiver.batch.max_bytes_disabled` | [splunk_receiver.batch.max_bytes_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/splunk_receiver/batch/max_bytes_disabled/#section) |
| `splunk_receiver.batch.max_events` | [splunk_receiver.batch.max_events](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/splunk_receiver/batch/#schema-splunk_receiver--batch--max_events) |
| `splunk_receiver.batch.max_events_disabled` | [splunk_receiver.batch.max_events_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/splunk_receiver/batch/max_events_disabled/#section) |
| `splunk_receiver.batch.timeout_seconds` | [splunk_receiver.batch.timeout_seconds](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/splunk_receiver/batch/#schema-splunk_receiver--batch--timeout_seconds) |
| `splunk_receiver.batch.timeout_seconds_default` | [splunk_receiver.batch.timeout_seconds_default](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/splunk_receiver/batch/timeout_seconds_default/#section) |
| `splunk_receiver.compression` | [splunk_receiver.compression](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/splunk_receiver/compression/#section) |
| `splunk_receiver.compression.compression_default` | [splunk_receiver.compression.compression_default](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/splunk_receiver/compression/compression_default/#section) |
| `splunk_receiver.compression.compression_gzip` | [splunk_receiver.compression.compression_gzip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/splunk_receiver/compression/compression_gzip/#section) |
| `splunk_receiver.compression.compression_none` | [splunk_receiver.compression.compression_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/splunk_receiver/compression/compression_none/#section) |
| `splunk_receiver.endpoint` | [splunk_receiver.endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/splunk_receiver/#schema-splunk_receiver--endpoint) |
| `splunk_receiver.no_tls` | [splunk_receiver.no_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/splunk_receiver/no_tls/#section) |
| `splunk_receiver.splunk_hec_token` | [splunk_receiver.splunk_hec_token](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/splunk_receiver/splunk_hec_token/#section) |
| `splunk_receiver.splunk_hec_token.blindfold_secret_info` | [splunk_receiver.splunk_hec_token.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/splunk_receiver/splunk_hec_token/blindfold_secret_info/#section) |
| `splunk_receiver.splunk_hec_token.blindfold_secret_info.decryption_provider` | [splunk_receiver.splunk_hec_token.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/splunk_receiver/splunk_hec_token/blindfold_secret_info/#schema-splunk_receiver--splunk_hec_token--blindfold_secret_info--decryption_provider) |
| `splunk_receiver.splunk_hec_token.blindfold_secret_info.location` | [splunk_receiver.splunk_hec_token.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/splunk_receiver/splunk_hec_token/blindfold_secret_info/#schema-splunk_receiver--splunk_hec_token--blindfold_secret_info--location) |
| `splunk_receiver.splunk_hec_token.blindfold_secret_info.store_provider` | [splunk_receiver.splunk_hec_token.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/splunk_receiver/splunk_hec_token/blindfold_secret_info/#schema-splunk_receiver--splunk_hec_token--blindfold_secret_info--store_provider) |
| `splunk_receiver.splunk_hec_token.clear_secret_info` | [splunk_receiver.splunk_hec_token.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/splunk_receiver/splunk_hec_token/clear_secret_info/#section) |
| `splunk_receiver.splunk_hec_token.clear_secret_info.provider_ref` | [splunk_receiver.splunk_hec_token.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/splunk_receiver/splunk_hec_token/clear_secret_info/#schema-splunk_receiver--splunk_hec_token--clear_secret_info--provider_ref) |
| `splunk_receiver.splunk_hec_token.clear_secret_info.url` | [splunk_receiver.splunk_hec_token.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/splunk_receiver/splunk_hec_token/clear_secret_info/#schema-splunk_receiver--splunk_hec_token--clear_secret_info--url) |
| `splunk_receiver.use_tls` | [splunk_receiver.use_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/splunk_receiver/use_tls/#section) |
| `splunk_receiver.use_tls.disable_verify_certificate` | [splunk_receiver.use_tls.disable_verify_certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/splunk_receiver/use_tls/disable_verify_certificate/#section) |
| `splunk_receiver.use_tls.disable_verify_hostname` | [splunk_receiver.use_tls.disable_verify_hostname](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/splunk_receiver/use_tls/disable_verify_hostname/#section) |
| `splunk_receiver.use_tls.enable_verify_certificate` | [splunk_receiver.use_tls.enable_verify_certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/splunk_receiver/use_tls/enable_verify_certificate/#section) |
| `splunk_receiver.use_tls.enable_verify_hostname` | [splunk_receiver.use_tls.enable_verify_hostname](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/splunk_receiver/use_tls/enable_verify_hostname/#section) |
| `splunk_receiver.use_tls.mtls_disabled` | [splunk_receiver.use_tls.mtls_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/splunk_receiver/use_tls/mtls_disabled/#section) |
| `splunk_receiver.use_tls.mtls_enable` | [splunk_receiver.use_tls.mtls_enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/splunk_receiver/use_tls/mtls_enable/#section) |
| `splunk_receiver.use_tls.mtls_enable.certificate` | [splunk_receiver.use_tls.mtls_enable.certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/splunk_receiver/use_tls/mtls_enable/#schema-splunk_receiver--use_tls--mtls_enable--certificate) |
| `splunk_receiver.use_tls.mtls_enable.key_url` | [splunk_receiver.use_tls.mtls_enable.key_url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/splunk_receiver/use_tls/mtls_enable/key_url/#section) |
| `splunk_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info` | [splunk_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/splunk_receiver/use_tls/mtls_enable/key_url/blindfold_secret_info/#section) |
| `splunk_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.decryption_provider` | [splunk_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/splunk_receiver/use_tls/mtls_enable/key_url/blindfold_secret_info/#schema-splunk_receiver--use_tls--mtls_enable--key_url--blindfold_secret_info--decryption_provider) |
| `splunk_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.location` | [splunk_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/splunk_receiver/use_tls/mtls_enable/key_url/blindfold_secret_info/#schema-splunk_receiver--use_tls--mtls_enable--key_url--blindfold_secret_info--location) |
| `splunk_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.store_provider` | [splunk_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/splunk_receiver/use_tls/mtls_enable/key_url/blindfold_secret_info/#schema-splunk_receiver--use_tls--mtls_enable--key_url--blindfold_secret_info--store_provider) |
| `splunk_receiver.use_tls.mtls_enable.key_url.clear_secret_info` | [splunk_receiver.use_tls.mtls_enable.key_url.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/splunk_receiver/use_tls/mtls_enable/key_url/clear_secret_info/#section) |
| `splunk_receiver.use_tls.mtls_enable.key_url.clear_secret_info.provider_ref` | [splunk_receiver.use_tls.mtls_enable.key_url.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/splunk_receiver/use_tls/mtls_enable/key_url/clear_secret_info/#schema-splunk_receiver--use_tls--mtls_enable--key_url--clear_secret_info--provider_ref) |
| `splunk_receiver.use_tls.mtls_enable.key_url.clear_secret_info.url` | [splunk_receiver.use_tls.mtls_enable.key_url.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/splunk_receiver/use_tls/mtls_enable/key_url/clear_secret_info/#schema-splunk_receiver--use_tls--mtls_enable--key_url--clear_secret_info--url) |
| `splunk_receiver.use_tls.no_ca` | [splunk_receiver.use_tls.no_ca](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/splunk_receiver/use_tls/no_ca/#section) |
| `splunk_receiver.use_tls.trusted_ca_url` | [splunk_receiver.use_tls.trusted_ca_url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/splunk_receiver/use_tls/#schema-splunk_receiver--use_tls--trusted_ca_url) |
| `sumo_logic_receiver` | [sumo_logic_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/sumo_logic_receiver/#section) |
| `sumo_logic_receiver.url` | [sumo_logic_receiver.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/sumo_logic_receiver/url/#section) |
| `sumo_logic_receiver.url.blindfold_secret_info` | [sumo_logic_receiver.url.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/sumo_logic_receiver/url/blindfold_secret_info/#section) |
| `sumo_logic_receiver.url.blindfold_secret_info.decryption_provider` | [sumo_logic_receiver.url.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/sumo_logic_receiver/url/blindfold_secret_info/#schema-sumo_logic_receiver--url--blindfold_secret_info--decryption_provider) |
| `sumo_logic_receiver.url.blindfold_secret_info.location` | [sumo_logic_receiver.url.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/sumo_logic_receiver/url/blindfold_secret_info/#schema-sumo_logic_receiver--url--blindfold_secret_info--location) |
| `sumo_logic_receiver.url.blindfold_secret_info.store_provider` | [sumo_logic_receiver.url.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/sumo_logic_receiver/url/blindfold_secret_info/#schema-sumo_logic_receiver--url--blindfold_secret_info--store_provider) |
| `sumo_logic_receiver.url.clear_secret_info` | [sumo_logic_receiver.url.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/sumo_logic_receiver/url/clear_secret_info/#section) |
| `sumo_logic_receiver.url.clear_secret_info.provider_ref` | [sumo_logic_receiver.url.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/sumo_logic_receiver/url/clear_secret_info/#schema-sumo_logic_receiver--url--clear_secret_info--provider_ref) |
| `sumo_logic_receiver.url.clear_secret_info.url` | [sumo_logic_receiver.url.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/sumo_logic_receiver/url/clear_secret_info/#schema-sumo_logic_receiver--url--clear_secret_info--url) |
| `timeouts` | [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/timeouts/#section) |
| `timeouts.create` | [timeouts.create](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/timeouts/#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/timeouts/#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/timeouts/#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/timeouts/#schema-timeouts--update) |

## Next pages

- [audit_logs](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/audit_logs/)
- [aws_cloud_watch_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/aws_cloud_watch_receiver/)
- [azure_event_hubs_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/azure_event_hubs_receiver/)
- [azure_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/azure_receiver/)
- [datadog_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/datadog_receiver/)
- [dns_logs](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/dns_logs/)
- [gcp_bucket_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/gcp_bucket_receiver/)
- [http_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/)
- [kafka_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/)
- [new_relic_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/new_relic_receiver/)
- [ns_all](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/ns_all/)
- [ns_current](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/ns_current/)
- [ns_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/ns_list/)
- [qradar_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/qradar_receiver/)
- [request_logs](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/request_logs/)
- [s3_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/s3_receiver/)
- [security_events](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/security_events/)
- [splunk_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/splunk_receiver/)
- [sumo_logic_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/sumo_logic_receiver/)
- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/timeouts/)
- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/)
